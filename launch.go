// Package launch lightly wraps the lovely envconfig for main
package launch

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var (
	// UseagePreamble is prepended to usage output if defined.
	UsagePreamble string
)

// Load looks for help flags and loads the config from env
//
// -h for help (handly for writing env file!)
// -c to show config that would be loaded from the current env
func Load(cfg any, prefix string) {

	ctx := context.Background()
	help(ctx, cfg, prefix)

	err := envconfig.Process(prefix, cfg)
	check(ctx, nil, err)
}

// Check exits when error is not nil
func Check(ctx context.Context, lgr Logger, err error) {

	check(ctx, lgr, err)
}

// unexported

func help(ctx context.Context, cfg any, configPrefix string) {

	// Todo: rearrange so that: 'flag.Set("help", "true")' can be used for unit plz

	h := flag.Bool("h", false, "show help message")
	help := flag.Bool("help", false, "show help message")
	c := flag.Bool("c", false, "show configuration loaded from env")
	conf := flag.Bool("conf", false, "show configuration loaded from env")

	flag.Parse()

	switch {
	case *h || *help:
		format := customFormat
		if UsagePreamble != "" {
			format = fmt.Sprintf("\n%s%s", UsagePreamble, customFormat)
		}

		tabs := tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)
		_ = envconfig.Usagef(configPrefix, cfg, tabs, format)
		tabs.Flush()

		os.Exit(0)
	case *c || *conf:
		err := envconfig.Process(configPrefix, cfg)
		check(ctx, nil, err)
		fmt.Println(pp(cfg))
		os.Exit(0)
	}
}

func check(ctx context.Context, lgr Logger, err error) {

	if err == nil {
		return
	}

	if lgr != nil {
		lgr.Error(ctx, "fatal top-level error", err)
		os.Exit(1)
	}

	// Todo: optionally don't exit for unit?

	panic(err)
}

func pp(cfg any) string {

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		err = errors.Wrap(err, "failed to marshal config")
		return string(fmt.Sprintf("%+v", err))
	}

	return string(data)
}

var customFormat = `
The following environment variables are available for configuration:

KEY	TYPE	DEFAULT	REQUIRED	DESCRIPTION
{{range .}}{{usage_key .}}	{{usage_type .}}	{{usage_default .}}	{{usage_required .}}	{{usage_description .}}
{{end}}
`
