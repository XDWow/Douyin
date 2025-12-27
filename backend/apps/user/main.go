package main

import (
	userv1 "github.com/XDWow/DouyinMall/backend/rpc_gen/kitex_gen/user/v1/userservice"
	"log"
)

func main() {
	svr := userv1.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
