package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fghwett/pupu/config"
	"github.com/fghwett/pupu/notify"
	"github.com/fghwett/pupu/task"
)

var path = flag.String("path", "./config.yml", "配置文件地址")

func main() {
	flag.Parse()

	conf, err := config.New(*path)
	if err != nil {
		fmt.Printf("读取配置文件失败 err: %s", err)
		os.Exit(-1)
	}

	t := task.New(conf)
	t.Do()

	if err := notify.Send(conf.ServerChan.SecretKey, "朴朴超市任务", t.GetResult()); err != nil {
		log.Printf("通知发送失败 %s\n", err)
	}
}
