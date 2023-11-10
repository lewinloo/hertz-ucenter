package main

import (
	"encoding/gob"
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz-ucenter/internal/dal"
	"hertz-ucenter/internal/handler"
	"hertz-ucenter/internal/models/vo"
	"hertz-ucenter/internal/mw"
	_ "hertz-ucenter/internal/service/logic"
	"hertz-ucenter/pkg/cfg"
)

func Init() error {
	// 加载配置
	if err := cfg.Load(); err != nil {
		return err
	}

	// 初始化数据访问层
	if err := dal.Init(); err != nil {
		return err
	}

	gob.Register(&vo.UserVO{})

	return nil
}

func main() {
	if err := Init(); err != nil {
		panic(err)
	}

	engine := server.Default(
		server.WithHostPorts(cfg.MustGet[string]("server.port")),
	)

	mw.InitSession(engine)

	// 注册路由
	handler.RegisterRoutes(engine)

	engine.Spin()
}
