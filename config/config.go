package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	path       string
	Name       string      `yaml:"name"`
	Config     *Config     `yaml:"config"`
	ServerChan *ServerChan `yaml:"serverChan"`
}

type Config struct {
	RefreshToken string `yaml:"refreshToken"`
	AccessToken  string `yaml:"accessToken"`
	ExpiredAt    int64  `yaml:"expiredAt"`
	Ua           string `yaml:"ua"`
	UserId       string `yaml:"userid"`
	PpOs         string `yaml:"ppos"`
	PpStoreId    string `yaml:"ppstoreid"`
}

type ServerChan struct {
	SecretKey string `yaml:"secretKey"`
}

func New(path string) (*Conf, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Conf{
		path: path,
	}
	err = yaml.Unmarshal(data, conf)

	return conf, err
}

func (c *Conf) Update(config *Config) error {
	c.Config = config
	return c.save()
}

func (c *Conf) save() error {
	content, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("配置yaml序列化出错 %v", err)
	}

	f, err := os.OpenFile(c.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("复写配置文件 打开失败 err: " + err.Error())
	}

	if _, err := io.WriteString(f, string(content)); err != nil {
		return fmt.Errorf("复写配置文件 写入失败 err: %v", err)
	}
	defer f.Close()

	return nil
}
