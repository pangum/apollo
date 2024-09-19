package config

import (
	"github.com/goexl/config"
)

type Apollo struct {
	// 配置端点
	Meta string `json:"endpoint,omitempty" validate:"required,url"`
	// 应用
	Appid string `json:"appid,omitempty" validate:"required"`
	// 集群
	Cluster string `json:"cluster,omitempty"`
	// 命名空间
	Namespaces config.Slice[string] `json:"namespaces,omitempty"`
}
