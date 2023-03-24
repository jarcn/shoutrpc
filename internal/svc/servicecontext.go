package svc

import (
	"github.com/jarcn/shoutrpc/internal/config"
	model "github.com/jarcn/shoutrpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
