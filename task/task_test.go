package task

import (
	"fmt"
	"log"
	"testing"

	"github.com/fghwett/pupu/config"
)

var conf *config.Conf

func init() {
	var err error
	conf, err = config.Init("../config.yml")
	if err != nil {
		log.Println("read config err: ", err)
		return
	}
}

func TestNew(t *testing.T) {
	task := New(conf.Config)
	task.Do()
	fmt.Println(task.GetResult())
}
