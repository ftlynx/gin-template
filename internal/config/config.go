package config

import (
	"errors"
	"fmt"
	"time"
)

//定义一个配置的interface
type Configer interface {
	GetConf() (*Config, error)
}

type Config struct {
	App     *AppConf     `toml:"app"`
	Mysql   *MysqlConf   `toml:"mysql"`
	Redis   *RedisConf   `toml:"redis"`
}

type AppConf struct {
	Listen string `toml:"listen"`
	Key    string `toml:"key"` //密码加密的key
}

type MysqlConf struct {
	DataSource  string        `toml:"data_source_url"`
	MaxOpenConn int           `toml:"max_open_conn"`
	MaxIdleConn int           `toml:"max_idle_conn"`
	MaxLifeTime time.Duration `toml:"max_life_time"`
}

type RedisConf struct {
	Addr   string `toml:"addr"`
	Passwd string `toml:"passwd"`
	DB     int    `toml:"db"`
}


func (c *Config) Validate() error {
	if c.App.Listen == "" {
		c.App.Listen = "0.0.0.0:3000"
	}
	if c.App.Key == "" {
		return fmt.Errorf("app key required")
	}

	if c.Mysql.MaxIdleConn == 0 {
		c.Mysql.MaxIdleConn = 6
	}

	if c.Mysql.MaxLifeTime == 0 {
		c.Mysql.MaxLifeTime = 3600
	}

	if c.Mysql.MaxOpenConn == 0 {
		c.Mysql.MaxOpenConn = 12
	}
	if c.Mysql.DataSource == "" {
		return errors.New("mysql data_source_url required")
	}

	return nil
}
