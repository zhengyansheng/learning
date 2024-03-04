package main

import (
	"os"

	"github.com/zhengyansheng/sample-kubelet/app"
	"k8s.io/component-base/cli"
)

func main() {
	command := app.NewKubeletCommand()
	code := cli.Run(command)
	os.Exit(code)
}
