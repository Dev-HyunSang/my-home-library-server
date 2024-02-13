package db

import (
	"fmt"
	"log"

	"github.com/dev-hyunsang/my-home-library-server/config"
	"github.com/dev-hyunsang/my-home-library-server/ent"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*ent.Client, error) {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_HOST"), config.GetEnv("DB_PORT"), config.GetEnv("DB_NAME"))
	client, err := ent.Open("mysql", addr)
	if err != nil {
		log.Println("[ERROR] Failed Conntion MySQL")
		return nil, err
	}

	return client, nil
}
