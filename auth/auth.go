package auth

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dev-hyunsang/my-home-library-server/config"
	"github.com/dev-hyunsang/my-home-library-server/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AceessToken  string
	RefreshToken string
	AceessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUUID string
	UserID     uuid.UUID
}

func CreateToken(userID uuid.UUID) (*TokenDetails, error) {
	td := new(TokenDetails)
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AceessUUID = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New().String()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AceessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AceessToken, err = at.SignedString([]byte(config.GetEnv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.GetEnv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func CreateAuth(userID uuid.UUID, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	client := db.ConnectRedis()

	result, err := client.Set(context.Background(), td.AceessUUID, userID.String(), at.Sub(now)).Result()
	if err != nil {
		return err
	}

	log.Println(result)

	result, err = client.Set(context.Background(), td.RefreshUUID, userID.String(), rt.Sub(now)).Result()
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}

func ExtractToken(bearToken string) string {
	strArr := strings.Split(bearToken, "Bearer ")[1]

	return strArr
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(tokenString string) error {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {
	removeToken := ExtractToken(tokenString)

	token, err := VerifyToken(removeToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		// String to UUID => USER UUID
		parseUserUUID, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return nil, err
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     parseUserUUID,
		}, nil
	}
	return nil, err
}

func FetchAuth(authD *AccessDetails) (uuid.UUID, error) {
	client := db.ConnectRedis()

	userID, err := client.Get(context.Background(), authD.AccessUUID).Result()
	if err != nil {
		return uuid.Nil, err
	}

	paserUserUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, err
	}

	return paserUserUUID, nil
}

func DeleteAuth(removeUUID string) (int64, error) {
	client := db.ConnectRedis()

	result, err := client.Del(context.Background(), removeUUID).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}
