
# Launch

Lightly Wrapped Envconfig for your Golang main

## Why?

Why wrap the excellent envconfig module?

 - "-h" flag for a list of environment variables of interest
 - "-c" flag to show what would be loaded from environment
 - Redact type for not disclosing sensitive strings via json.Marshal
 - demonstrate super handy Config pattern for packages
 - top-level error checking
 - b-but mostly to test and reuse this surprisingly fiddly code

## Help Flag

    launch % go run github.com/clarktrimble/launch/examples/thingone -h
    This application is configured via the environment. The following environment
    variables can be used:

    KEY                   TYPE       DEFAULT    REQUIRED    DESCRIPTION
    DEMO_THINGTWO         String     bargle
    DEMO_TOKEN            Redact                true
    DEMO_SVC_IMPORTANT    String                true
    DEMO_SVC_NOTSOMUCH    Integer    42

Copy and paste from here into env file!

## Config Flag and Redact

    launch % source examples/thingone/demo.env
    launch % go run github.com/clarktrimble/launch/examples/thingone -c
    {
      "version": "",
      "thing_two": "thingone",
      "token": "--redacted--",
      "svc_layer": {
        "important": "Brush and floss every day!",
        "not_so_much": 42
      }
    }

Nice for a quick sanity check!

## Config Pattern

From package:

    type Config struct {
      Important string `json:"important" required:"true"`
      NotSoMuch int    `json:"not_so_much" default:"42"`
    }

    func (cfg *Config) New() (svc *SvcLayer, err error) {
      // ...
      return
    }

In main:

    type Config struct {
      // ...
      Version  string           `json:"version" ignored:"true"`
      Svc      *svclayer.Config `json:"demo_svc"`
    }
    // ...

    cfg := &Config{Version: version}
    launch.Load(cfg, cfgPrefix)
    // ...

    svc, err := cfg.Svc.New()
    launch.Check(context.Background(), lgr, err)
    // ...

So yeah, such packages are "dependent" on the pattern, but feels like a reasonable 80/20 trade-off
so long as confined to the service-layer (ala Clean Architecture).

## Check

Heresy perhaps, but at the top one checks for errors.  Nice to capture this in the logs where possible.

## Test, Build, and/or Run

    launch % go test --count=1 ./...
    launch % go run github.com/clarktrimble/launch/examples/thingone

or

    proj/launch % make build
    :: Done
    proj/launch % make test
    CGO_ENABLED=0 go test -count 1 ./...
    ?   	github.com/clarktrimble/launch/examples/thingone	[no test files]
    ?   	github.com/clarktrimble/launch/examples/thingone/minlog	[no test files]
    ?   	github.com/clarktrimble/launch/examples/thingone/svclayer	[no test files]
    ok  	github.com/clarktrimble/launch	0.130s

## Golang (Anti) Idioms

I dig the Golang community, but I might be a touch rouge with:

  - multi-char variable names
  - named return parameters
  - BDD/DSL testing

Todo: spellcheck my man!

All in the name of readability, which of course, tends towards the subjective.

## License

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>

