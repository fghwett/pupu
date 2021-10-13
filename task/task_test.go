package task

import (
	"fmt"
	"log"
	"testing"

	"github.com/fghwett/pupu/config"
)

var con *config.Conf

func init() {
	var err error
	con, err = config.New("../config.yml")
	if err != nil {
		log.Println("read config err: ", err)
		return
	}
}

func TestNew(t *testing.T) {
	task := New(con)
	task.Do()
	fmt.Println(task.GetResult())
}
