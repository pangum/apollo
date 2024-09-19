package get

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/pangum/apollo/internal/config"
	"github.com/pangum/pangu"
)

type Client struct {
	pangu.Get

	Config *config.Apollo
	Logger log.Logger   `optional:"true"`
	Http   *http.Client `optional:"true"`
}
