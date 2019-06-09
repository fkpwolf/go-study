package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Send email to %s<%s>\n", u.name, u.email)
}
func (u *user) changeEmail(email string) {
	u.email = email
}

// 本方法不会修改值
func (u user) changeEmailAgain(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()
	bill.changeEmail("bill@163.com")
	bill.notify()
	bill.changeEmailAgain("bill@qq.com")
	bill.notify()
}
