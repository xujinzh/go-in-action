package main

import "fmt"

type user struct {
	name  string
	email string
}

// u user 称为接收者
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// u *user 称为接收者
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify() // notify 操作的是一个 bill 副本

	lisa := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify() // 底层先调整为 (*lisa).notify，notify 操作的是一个 *lisa 解引用后的副本

	// 底层先调整为 (&bill).changeEmail()，然后在调用
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// changeEmail 会共享调用方法时接收者所指向的值，它的修改会反映到接收者指针的所指向的值上
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
