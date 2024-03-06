package app

import (
	"go.temporal.io/sdk/workflow"
)

func GreetSomeone(ctx workflow.Context, str string) (string, error) {
	return "Hello " + str + "!", nil
}
