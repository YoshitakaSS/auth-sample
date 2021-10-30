package main

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/src/config"
	"github.com/joho/godotenv"
)

func main() {
	// 明示的に環境変数を読み込み
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbMap, err := config.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	redis, err := config.InitRedis()
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		dbMap.Db.Close()
		redis.Close()
	}()
}
