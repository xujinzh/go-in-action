package main

/*
如果外部类型实现了内部类型同样的方法，那么内部类型的实现就不会被提升。

如果外部类型没有实现内部类型同样的方法，那么内部类型的实现就会被提升。

*/
import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // 嵌入类型
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>-%s\n", a.name, a.email, a.level)
}

func main() {
	ad := admin{
		user: user{
			name:  "John smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 直接访问内部类型的方法
	ad.user.notify()
	// 内部类型的方法也被提升到外部类型，当外部类型没有实现 notify 方法时
	ad.notify()
}
