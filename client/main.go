package main

import (
	"chatroom/redis"
	"fmt"
	"os"
)


const (
	Network         = "tcp"
	Address         = "127.0.0.1:8081"
)

var (
	userId  int
	userPwd string
)

func main() {
	initChatroom()
}

func initChatroom() {
	var choice int
	var loop = true

	for loop {
		fmt.Println("==============欢迎来到多人聊天室================")
		fmt.Println("1. 登录聊天室")
		fmt.Println("2. 注册用户")
		fmt.Println("3. 退出聊天室")
		fmt.Println("请选择1-3:")

		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册新用户")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

	if choice == 1 {
		fmt.Println("请输入您的用户id:")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入您的用户密码:")
		fmt.Scanf("%s\n", &userPwd)
		login(userId, userPwd)
	} else if choice == 2 {
		fmt.Println("执行注册新用户步骤")
		redis.RedisPingpong()
	} else {
		fmt.Println("执行退出聊天室")
		os.Exit(0)
	}
}
