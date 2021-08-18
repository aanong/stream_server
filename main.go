package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"stream_server/pkg/util"
)

func main() {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		//  日志处理
		util.Errorf("读取文件失败：v", err)
		os.Exit(0)
	}
	fmt.Print(cfg)

}
