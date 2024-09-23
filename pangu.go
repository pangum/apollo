package apollo

import (
	"github.com/pangum/apollo/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Puts(
		ctor.LoadConfig, // 加载配置
		ctor.NewClient,  // 创建客户端
	).Get(
		ctor.RegisterLoader, // 注册加载器
	).Build().Build().Apply()
}
