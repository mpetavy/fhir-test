package main

import (
	"embed"
	"fmt"
	"github.com/mpetavy/common"
	"os"
	"time"
)

//go:embed go.mod
var resources embed.FS

func init() {
	common.Init("", "", "", "", "test", "", "", "", &resources, nil, nil, run, 0)
}

func run() error {
	src, err := os.ReadFile("./test.js")
	if common.Error(err) {
		return err
	}

	se, err := common.NewScriptEngine(string(src), "/home/ransom/app/node-v20.9.0-linux-x64/lib/node_modules")
	if common.Error(err) {
		return err
	}

	output, err := se.Run(time.Hour*3, "", "")
	if common.Error(err) {
		return err
	}

	fmt.Printf("%s\n", output)

	return nil
}

func main() {
	common.Run(nil)
}
