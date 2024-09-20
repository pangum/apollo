package core

import (
	"context"
	"fmt"

	"github.com/goexl/apollo"
	"github.com/goexl/gox"
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

func (l *Loader) Load(ctx context.Context, target *map[string]any, modules []string) (loaded bool, err error) {
	loader := l.client.Loader().Context(ctx)
	loader.Namespace("application") // 默认命令空间

	namespaces := make(map[string]*gox.Empty)
	for _, module := range modules {
		namespaces[module] = nil
		namespaces[fmt.Sprintf("%s.yaml", module)] = nil
		namespaces[fmt.Sprintf("%s.yml", module)] = nil
		namespaces[fmt.Sprintf("%s.json", module)] = nil
	}
	for namespace := range namespaces {
		loader.Namespace(namespace)
	}
	if le := loader.Build().Load(target); nil != le {
		err = le
	} else {
		loaded = true
	}

	return
}
