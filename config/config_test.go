package config

import "testing"

func TestInit(t *testing.T) {
	conf, err := New("../config.yml")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(conf.ServerChan.SecretKey)

	conf.Config.ExpiredAt = 2000

	if err := conf.save(); err != nil {
		t.Error(err)
	}
}
