package main

import (
	"fmt"
)

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type Duration2 int64

func main() {
	var bill user
	fmt.Printf("bill: %#v\n", bill)

	lisa := user{
		name:       "Lisa",
		email:      "lisa@gmail.com",
		ext:        123,
		privileged: true,
	}
	fmt.Printf("lisa: %v\n", lisa)

	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@gmail.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Printf("fred: %v\n", fred)

	var dur Duration2
	dur = Duration2(1000)
	fmt.Printf("dur: %v\n", dur)
}
