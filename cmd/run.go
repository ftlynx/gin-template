package cmd

import (
	"gin-template/api/http/route"
	"gin-template/internal/config/file"
	"gin-template/internal/global"
	"gin-template/pkg"
	"github.com/ftlynx/tsx/mysqlx"
	"github.com/ftlynx/tsx/redisx"
	"github.com/spf13/cobra"
)

var confFile = "gin-template.conf"

var runCmd = &cobra.Command{
	Use:   "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := file.NewFileConf(confFile).GetConf()
		if err != nil {
			return err
		}

		global.SqlxDB, err = mysqlx.SqlxDB(c.Mysql.DataSource, c.Mysql.MaxOpenConn, c.Mysql.MaxIdleConn, c.Mysql.MaxLifeTime)
		if err != nil {
			return err
		}

		global.Redis, err = redisx.InitRedis(c.Redis.Addr, c.Redis.Passwd, c.Redis.DB)
		if err != nil {
			return err
		}

		global.Conf = c

		if err := pkg.InitConfig(); err != nil {
			return err
		}

		return route.MyRoute()
	},
}

func init() {
	runCmd.Flags().StringVarP(&confFile, "config", "f", confFile, "the service config from file")
	RootCmd.AddCommand(runCmd)
}
