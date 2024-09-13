package plugin

import (
	"github.com/pangum/apollo/internal/config"
)

type Wrapper struct {
	Apollo *config.Apollo `json:"apollo,omitempty" yaml:"apollo,omitempty" xml:"apollo,omitempty" toml:"apollo,omitempty"`
}
