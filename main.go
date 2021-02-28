package main

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/config"
	"github.com/YoshitakaSS/go_auth/infra"
	"github.com/YoshitakaSS/go_auth/web"
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
	defer dbMap.Db.Close()

	redis, err := infra.InitRedis()
	if err != nil {
		fmt.Println(err)
	}
	defer redis.Close()

	s := web.NewServer()
	s.Start(fmt.Sprintf("%s:%s", config.HostName(), config.Port()))
}
