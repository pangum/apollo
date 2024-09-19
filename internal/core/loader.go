package core

import (
	"context"

	"github.com/goexl/apollo"
	"github.com/pangum/config"
)

var _ config.Loader = (*Loader)(nil)

type Loader struct {
	client *apollo.Client
}

func NewLoader(client *apollo.Client) *Loader {
	return &Loader{
		client: client,
	}
}

func (*Loader) Local() bool {
	return false
}

func (*Loader) Extensions() []string {
	return []string{
		// !不限制扩展名
	}
}

func (l *Loader) Load(ctx context.Context, target *map[string]any) (loaded bool, err error) {
	loader := l.client.Loader().Context(ctx)
	loader.Namespace("application") // 默认命令空间
	loader.Namespace("logging.yaml")
	loader.Namespace("logging.yml")
	loader.Namespace("logging.json")
	if le := loader.Build().Load(target); nil != le {
		err = le
	} else {
		loaded = true
	}

	return
}
