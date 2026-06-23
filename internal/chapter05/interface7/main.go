package main

/*
内部嵌入类型 user 实现了接口 notifier，外部类型并没有实现这个接口。

但是，因为内部类型的提升，内部类型实现的接口会自动提升到外部类型。
这意味着由于内部类型的实现，外部类型也同样实现了这个接口。
*/

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "John smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
