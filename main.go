package main

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/infra"
)

func main() {
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

}
