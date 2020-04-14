package file

import (
	"fmt"
	"gin-template/internal/config"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var (
	cfg *config.Config
	once   sync.Once
)

func NewFileConf(filePath string) config.Configer {
	return &fileConfig{filePath: filePath}
}

type fileConfig struct {
	filePath string
}

func (f *fileConfig) GetConf() (*config.Config, error) {
	var err error

	once.Do(func() {
		err = parseConfig(f.filePath)
	})

	if err != nil {
		return nil, err
	}

	if err = cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseConfig(configpath string) error {
	configPath, err := filepath.Abs(configpath)
	if err != nil {
		return fmt.Errorf("get config file absolute path failed, %s", err.Error())
	}

	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("open config file error, %s", err.Error())
	}

	fd, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("read config file error, %s", err.Error())
	}

	cfg = new(config.Config)
	cfg.App = new(config.AppConf)
	cfg.Mysql = new(config.MysqlConf)
	cfg.Redis=new(config.RedisConf)

	if err := toml.Unmarshal(fd, cfg); err != nil {
		return fmt.Errorf("load config file error, %s", err.Error())
	}

	return nil
}
