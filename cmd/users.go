// 2024.02.12 HyunSang Park <me@hyunsang.dev>
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dev-hyunsang/my-home-library-server/auth"
	"github.com/dev-hyunsang/my-home-library-server/db"
	"github.com/dev-hyunsang/my-home-library-server/dto"
	"github.com/dev-hyunsang/my-home-library-server/ent/user"
	"github.com/dev-hyunsang/my-home-library-server/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func JoinUserHandler(ctx *fiber.Ctx) error {
	req := new(dto.RequestJoinUser)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("")
	}

	/*회원가입 절차
	0. 사용자의 입력값 받기
	1. 메일, 닉네임 중복 확인
	2. 패스워드 난독화
	3. 데이터베이스 저장
	*/

	client, err := db.ConnectMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	result, err := client.User.Create().
		SetID(uuid.New()).
		SetNickname(req.NickName).
		SetEmail(req.Email).
		SetPassword(string(hashedPw)).
		SetCreatedAt(time.Now()).
		SetEditedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		log.Println("Failed Create User")
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseJoinUser{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "성공적으로 회원가입 했어요! 오늘부터 열심히 책을 읽어봐요!",
		},
		Data:        result,
		RespondedAt: time.Now(),
	})
}

func LoginUserHandler(ctx *fiber.Ctx) error {
	req := new(dto.RequestLoginUser)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("")
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	result, err := client.User.Query().Where(user.Email(req.Email)).First(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("")
	}

	ts, err := auth.CreateToken(result.ID)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON("")
	}

	err = auth.CreateAuth(result.ID, ts)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusNonAuthoritativeInformation).JSON("")
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseLoginUser{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "어서와요! 오늘도 열심히 책을 읽어봐요 :)",
		},
		Data: dto.JwtData{
			AccessUUID:   ts.AceessUUID,
			AccessToken:  ts.AceessToken,
			RefreshUUID:  ts.RefreshUUID,
			RefreshToken: ts.RefreshToken,
		},
		RespondedAt: time.Now(),
	})
}

func CheckinUserDataHandler(ctx *fiber.Ctx) error {
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	result, err := client.User.Query().Where(user.ID(token.UserID)).Only(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseUserData{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "성공적으로 사용자의 정보를 가지고 왔습니다.",
		},
		Data:        result,
		RespondedAt: time.Now(),
	})

}

func LogOutUserHandler(ctx *fiber.Ctx) error {
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{})
	}

	deleted, err := auth.DeleteAuth(token.AccessUUID)
	if err != nil || deleted == 0 {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	logger.Info(fmt.Sprintf("Deleted By %d", uint64(deleted)))

	return ctx.Status(fiber.StatusOK).JSON("")
}
