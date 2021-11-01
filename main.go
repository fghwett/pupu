package main

import (
	"flag"
	"fmt"
	"github.com/fghwett/pupu/config"
	"github.com/fghwett/pupu/notify"
	"github.com/fghwett/pupu/task"
	"github.com/fghwett/pupu/util"
	"log"
	"os"
	"runtime"
)

var path = flag.String("path", "./config.yml", "配置文件地址")

var (
	conf *config.Conf

	t *task.Task

	err error
)

func main() {
	flag.Parse()

	conf, err = config.New(*path)
	if err != nil {
		fmt.Printf("读取配置文件失败 err: %s", err)
		os.Exit(-1)
	}

	util.BigSleep(5, 20)

	t = task.New(conf)
	WithRecover(func() {
		t.Do()
	})

	if err = notify.Send(conf.ServerChan.SecretKey, conf.Name, t.GetResult()); err != nil {
		log.Printf("通知发送失败 %s\n", err)
	}
}

func WithRecover(f func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("捕获到系统异常：%v\n", r)
			buf := make([]byte, 1<<16)
			runtime.Stack(buf, true)
			log.Printf("系统异常内容：%s \n", string(buf))

			if err = notify.Send(conf.ServerChan.SecretKey, conf.Name, t.GetResult()); err != nil {
				log.Printf("通知发送失败 %s\n", err)
			}
		}
	}()
	f()
}
