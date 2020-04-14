package global

import (
	"database/sql"
	"gin-template/internal/config"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var (
	//Conf app config
	Conf *config.Config

	// mysql config
	SqlDB  *sql.DB
	SqlxDB *sqlx.DB

	// redis conn
	Redis *redis.Client
)
