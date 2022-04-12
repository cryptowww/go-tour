// 通过一个工程来演示如何构建模块，我们通过构建hello和greeting两个模块来演示，hello模块是主调模块，greeting是被调用模块
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// 调用方法
//      messages, err := greetings.hellos(names)
func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	//message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
