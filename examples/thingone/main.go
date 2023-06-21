package main

import (
	"context"

	lh "github.com/clarktrimble/launch"
	"github.com/clarktrimble/launch/examples/thingone/minlog"
	"github.com/clarktrimble/launch/examples/thingone/svclayer"
)

var (
	version string
)

const (
	cfgPrefix string = "demo"
)

type Config struct {
	Version  string           `json:"version" ignored:"true"`
	ThingTwo string           `json:"thing_two" default:"bargle"`
	Token    lh.Redact        `json:"token" required:"true"`
	Svc      *svclayer.Config `json:"demo_svc"`
}

func main() {

	cfg := &Config{Version: version}
	lh.Load(cfg, cfgPrefix)

	lgr := &minlog.MinLog{}

	svc, err := cfg.Svc.New()
	lh.Check(context.Background(), lgr, err)

	svc.Disintermediate()
}
