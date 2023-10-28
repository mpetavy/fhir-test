package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"os"
	"time"
)

func init() {
	common.Init("test", "", "", "", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
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
