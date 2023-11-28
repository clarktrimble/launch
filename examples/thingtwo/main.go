package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/clarktrimble/launch"
	"github.com/clarktrimble/launch/examples/thingone/minlog"
	"github.com/clarktrimble/launch/examples/thingone/svc"
	"github.com/clarktrimble/launch/spinner"
)

var (
	version string
)

const (
	cfgPrefix string = "demo"
)

type Config struct {
	Version  string        `json:"version" ignored:"true"`
	ThingTwo string        `json:"thing_two" desc:"the second thing" default:"bargle"`
	Token    launch.Redact `json:"token" desc:"secret for auth" required:"true"`
	Svc      *svc.Config   `json:"demo_svc"`
}

// Todo: check that sourcing env for demo is mentioned somewheres

func main() {

	// load config from env

	cfg := &Config{Version: version}
	launch.UsagePreamble = "thingone demonstrates use of the launch pkg."
	launch.Load(cfg, cfgPrefix)

	// log config

	lgr := &minlog.MinLog{}
	ctx := context.Background()
	lgr.Info(ctx, "starting", "config", cfg)

	// demonstrate that dependency is configured

	svc, err := cfg.Svc.New()
	launch.Check(ctx, lgr, err)

	svc.Disintermediate()

	// workout spinner

	sp := spinner.New()
	for i := 0; i < 99; i++ {
		sp.Spin()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(99)))
	}

	fmt.Printf("%d operations in %.2f seconds\n", sp.Count, sp.Elapsed())
}
