package main

import "fmt"

type Users struct {
	Users []*User
}
type User struct {
	Name string
	Age  int
}

// 错误用法
func IncorrectUsageOfForrangeAndPointer() {
	userList := []User{
		{Name: "aa", Age: 1},
		{Name: "bb", Age: 1},
	}

	var newUser []*User
	for _, u := range userList {
		newUser = append(newUser, &u)
	}

	for _, nu := range newUser {
		fmt.Printf("%+v\n", nu.Name)
	}
}

// 正确用法一
func CorrectUsageOfForrangeAndPointer() {
	usersList := Users{
		Users: []*User{
			{Name: "aa", Age: 1},
			{Name: "bb", Age: 1},
		},
	}

	var newUser []*User
	for _, u := range usersList.Users {
		newUser = append(newUser, u)
	}

	for _, nu := range newUser {
		fmt.Printf("%+v\n", nu.Name)
	}
}

// 正确用法二
func CorrectUsageOfForAndPointer() {
	userList := []User{
		{Name: "aa", Age: 1},
		{Name: "bb", Age: 1},
	}

	var newUser []*User
	for i := 0; i < len(userList); i++ {
		newUser = append(newUser, &userList[i])
	}

	for _, nu := range newUser {
		fmt.Printf("%+v\n", nu.Name)
	}
}
