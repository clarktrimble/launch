// Package svc demonstrates configuring a service layer with launch.
package svc

import (
	"fmt"

	"github.com/pkg/errors"
)

// Config represents configurables for an Svc.
type Config struct {
	Important string `json:"important" desc:"an important value" required:"true"`
	NotSoMuch int    `json:"not_so_much" desc:"a less important one" default:"42"`
}

// Svc is an example service layer in need of some config.
type Svc struct {
	imp string
	nsm int
}

// New creates an Svc.
func New(imp string, nsm int) (svc *Svc, err error) {

	if nsm < 0 {
		err = errors.Errorf("nsm may not be negative, got: %d", nsm)
		return
	}

	svc = &Svc{
		imp: imp,
		nsm: nsm,
	}

	return
}

// New creates an Svc from Config.
func (cfg *Config) New() (*Svc, error) {

	return New(cfg.Important, cfg.NotSoMuch)
}

// Disintermediate prints some stuff.
func (svc *Svc) Disintermediate() {
	fmt.Printf("Important: %s\n", svc.imp)
	fmt.Printf("Not as important: %d\n", svc.nsm)
}
