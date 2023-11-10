package dal

import (
	"hertz-ucenter/pkg/cfg"
	"hertz-ucenter/pkg/db"
	"time"
)

func Init() error {
	mysql, err := db.NewMySQL(&db.MySQLOptions{
		Host:                  cfg.MustGet[string]("db.host"),
		Username:              cfg.MustGet[string]("db.username"),
		Password:              cfg.MustGet[string]("db.password"),
		Database:              cfg.MustGet[string]("db.database"),
		MaxIdleConnections:    cfg.MustGet[int]("db.max-idle-connections"),
		MaxOpenConnections:    cfg.MustGet[int]("db.max-open-connections"),
		MaxConnectionLifeTime: time.Duration(cfg.MustGet[int]("db.max-connection-life-time")) * time.Second,
		LogLevel:              cfg.MustGet[int]("db.log-level"),
	})

	if err != nil {
		return err
	}

	SetDefault(mysql)

	return nil
}
