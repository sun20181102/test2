package main

import (
	"fmt"

	"github.com/test2/pkg/echo"
	hellowword "github.com/test2/pkg/helloword"
)

func main() {
	m := hellowword.HelloWorld()
	fmt.Println(m)
	message := echo.Echo("asdf")
	fmt.Println(message)
}
