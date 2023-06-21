// Package svclayer demonstrates configuring a service layer with launch
package svclayer

import (
	"fmt"

	"github.com/pkg/errors"
)

type Config struct {
	Important string `json:"important" required:"true"`
	NotSoMuch int    `json:"not_so_much" default:"42"`
}

type SvcLayer struct {
	imp string
	nsm int
}

func (cfg *Config) New() (svc *SvcLayer, err error) {

	if cfg.NotSoMuch < 0 {
		err = errors.Errorf("configurable 'not_so_much' may not be negative, got: %d", cfg.NotSoMuch)
		return
	}

	svc = &SvcLayer{
		imp: cfg.Important,
		nsm: cfg.NotSoMuch,
	}

	return
}

func (svc *SvcLayer) Disintermediate() {
	fmt.Printf("Important: %s\n", svc.imp)
	fmt.Printf("Not as important: %d\n", svc.nsm)
}
