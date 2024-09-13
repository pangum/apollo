package plugin

import (
	"github.com/pangum/apollo/internal/config"
	"github.com/pangum/apollo/internal/core"
	"github.com/pangum/pangu"
	"github.com/philchia/agollo/v4"
)

type Constructor struct {
	// 构造方法
}

func (*Constructor) NewClient(config *config.Apollo) (client agollo.Client, err error) {
	client = agollo.NewClient(&agollo.Conf{
		AppID:          config.Appid,
		Cluster:        config.Cluster,
		NameSpaceNames: config.Namespaces,
		MetaAddr:       config.Endpoint,
	})
	err = client.Start()

	return
}

func (*Constructor) RegisterLoader(apollo agollo.Client) {
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
