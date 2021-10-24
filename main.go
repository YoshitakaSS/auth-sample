package main

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/infra"
	"github.com/joho/godotenv"
)

func main() {
	// 明示的に環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbMap, err := infra.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	redis, err := infra.InitRedis()
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		dbMap.Db.Close()
		redis.Close()
	}()
}
