package main

import (
	"fmt"

	"github.com/sun20181102/test2/pkg/echo"
	hellowword "github.com/sun20181102/test2/pkg/helloword"
)

func main() {
	m := hellowword.HelloWorld()
	fmt.Println(m)
	message := echo.Echo("asdf")
	fmt.Println(message)
}
