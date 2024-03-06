package main

import (
	"fmt"
	"os"

	"github.com/pulkit-tanwar/temporal-workflow/app"
)

func main() {

	name := "stranger"

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ans := app.GreetSomeone(name)
	fmt.Println(ans)
}
