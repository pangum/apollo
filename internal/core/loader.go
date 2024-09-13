package core

import (
	"context"
	"fmt"

	"github.com/pangum/config"
	"github.com/philchia/agollo/v4"
)

var _ config.Loader = (*Loader)(nil)

type Loader struct {
	apollo agollo.Client
}

func NewLoader(apollo agollo.Client) *Loader {
	return &Loader{
		apollo: apollo,
	}
}

func (l *Loader) Local() bool {
	return false
}

func (l *Loader) Load(ctx context.Context, target any) (loaded bool, err error) {
	fmt.Println(l.apollo.GetContent(agollo.WithNamespace("logging")))
	fmt.Println(l.apollo.GetContent())
	fmt.Println(l.apollo.GetString("level", agollo.WithNamespace("logging")))

	return
}
