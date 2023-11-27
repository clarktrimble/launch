package main

import (
	"context"

	"github.com/clarktrimble/launch"
	"github.com/clarktrimble/launch/examples/thingone/minlog"
	"github.com/clarktrimble/launch/examples/thingone/svc"
)

var (
	version string
)

const (
	cfgPrefix string = "demo"
	usage     string = "'thingone' demonstrates use of the launch pkg."
)

type Config struct {
	Version  string        `json:"version" ignored:"true"`
	ThingTwo string        `json:"thing_two" desc:"the second thing" default:"bargle"`
	Token    launch.Redact `json:"token" desc:"secret for auth" required:"true"`
	Svc      *svc.Config   `json:"demo_svc"`
}

func main() {

	cfg := &Config{Version: version}
	launch.Load(cfg, cfgPrefix, usage)

	lgr := &minlog.MinLog{}
	ctx := context.Background()
	lgr.Info(ctx, "starting", "config", cfg)

	svc, err := cfg.Svc.New()
	launch.Check(ctx, lgr, err)

	svc.Disintermediate()
}
