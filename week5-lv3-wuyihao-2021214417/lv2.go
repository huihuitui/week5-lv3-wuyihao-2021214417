package main

import "fmt"

func UserInfo() {
	var key, value string
	Usermap = make(map[string]User)
	fmt.Println("输入用户名")
	fmt.Scan(&key)
	fmt.Println("输入密码")
	fmt.Scan(&value)
	Usermap[key] = User{
		Name:     key,
		Password: value,
	}
	str = fmt.Sprintf("%s:%s", key, value)
}
