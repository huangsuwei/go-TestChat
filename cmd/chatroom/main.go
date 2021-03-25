package main

import (
	"fmt"
)

var (
	addr   = ":2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

Go 语言编程之旅 —— 一起用 Go 做项目：ChatRoom，start on：%s
`
)

func main() {
	/*fmt.Printf(banner+"\n", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))*/

	CompareWithStruct()
}

type Dog struct {
	Name string
}

func CompareWithStruct() {
	a := Dog{Name: "A"}
	b := Dog{Name: "A"}
	c := &a
	fmt.Println(c.Name)
	fmt.Println(a == b)
	a.Name = "B"
	fmt.Println(a.Name)
	fmt.Println(a == b)
}
