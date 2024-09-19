package plugin

import (
	"github.com/goexl/apollo"
	"github.com/pangum/apollo/internal/config"
	"github.com/pangum/apollo/internal/core"
	"github.com/pangum/apollo/internal/plugin/internal/get"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewClient(get get.Client) *apollo.Client {
	builder := apollo.New().Appid(get.Config.Appid).Meta(get.Config.Meta)
	if nil != get.Logger {
		builder.Logger(get.Logger)
	}
	if nil != get.Http {
		builder.Http(get.Http)
	}
	if "" != get.Config.Cluster {
		builder.Cluster(get.Config.Cluster)
	}
	if 0 != len(get.Config.Namespaces) {
		builder.Namespaces(get.Config.Namespaces[0], get.Config.Namespaces[1:]...)
	}

	return builder.Build()
}

func (*Constructor) RegisterLoader(apollo *apollo.Client) {
	pangu.New().Config().Loader(core.NewLoader(apollo)).Build() // !注册加载器
}

func (*Constructor) LoadConfig(config *pangu.Config) (apollo *config.Apollo, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		apollo = wrapper.Apollo
	}

	return
}
