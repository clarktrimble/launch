
# Launch

![launch](https://github.com/clarktrimble/launch/assets/5055161/e22d6779-3ef2-459a-a3d6-f22eaebd5eee)

Lightly Wrapped envconfig for Golang

## Why?

Why wrap the excellent envconfig module?

 - "-h" flag for a list of environment variables available for configuration
 - "-c" flag to show what would be loaded from environment
 - redact type for non-disclosure via json.Marshal
 - usage blerb for a gentle reminder as to purpose
 - top-level error checking
 - simple spinner for Herculean command-line utils
 - demonstrate encapsulated Config for dependencies
 - but as much as anything to test and reuse surprisingly fiddly code

## Help Flag

```bash
~/proj/launch$ go run examples/thingone/main.go -h

'thingone' demonstrates use of the launch pkg.

The following environment variables are available for configuration:

KEY                   TYPE       DEFAULT    REQUIRED    DESCRIPTION
DEMO_THINGTWO         String     bargle                 the second thing
DEMO_TOKEN            Redact                true        secret for auth
DEMO_SVC_IMPORTANT    String                true        an important value
DEMO_SVC_NOTSOMUCH    Integer    42                     a less important one
```

Copy and paste from here into env file!

## Config Flag and Redact

```bash
~/proj/launch$ make
go generate ./...
golangci-lint run ./...
go test -count 1 github.com/clarktrimble/launch github.com/clarktrimble/launch/spinner
ok      github.com/clarktrimble/launch  0.005s
ok      github.com/clarktrimble/launch/spinner  0.014s
:: Building thingone
go build -ldflags '-X main.version=spin.14.d94b8a6' -o bin/thingone examples/thingone/main.go
:: Done

~/proj/launch$ . examples/thingone/env.sh ## source env file

~/proj/launch$ bin/thingone -c
{
  "version": "spin.14.d94b8a6",
  "thing_two": "thingone",
  "token": "--redacted--",
  "demo_svc": {
    "important": "Brush and floss every day!",
    "not_so_much": 42
  }
}
```

Nice for a quick sanity check!

Notice how `version` sneaks in with the build and `token` is redacted.

## Encapsulated Config

In some package:

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

Voila!
The package's configuration requirements are encapsulated.

Check out the [post](https://clarktrimble.online/blog/encapsulated-env-cfg/#encapsulation) over on _Ba Blog_ for more.

## Check

Triggering an error in `thingone`:

```bash
~/proj/launch$ DEMO_SVC_NOTSOMUCH=-1 bin/thingone
msg > starting up
kvs > ::config::{"version":"spin.14.d94b8a6","thing_two":"thingone","token":"--redacted--","demo_svc":{"important":"Brush and floss every day!","not_so_much":-1}}

err > fatal top-level error nsm may not be negative, got: -1
github.com/clarktrimble/launch/examples/thingone/svc.New
        /home/trimble/proj/launch/examples/thingone/svc/svclayer.go:26
github.com/clarktrimble/launch/examples/thingone/svc.(*Config).New
        /home/trimble/proj/launch/examples/thingone/svc/svclayer.go:41
main.main
        /home/trimble/proj/launch/examples/thingone/main.go:48
runtime.main
        /home/trimble/go1211/src/runtime/proc.go:267
runtime.goexit
        /home/trimble/go1211/src/runtime/asm_amd64.s:1650
```

We'll want to keep this sort of thing to a minimum in main of course.
Nice to capture the error in the logs though when it's expeditious.

Notice token is still redacted!

Hat tip to the legendary [Dave Cheney](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) for the stack trace.

## Spinner

A little fluffy perhaps, but nice to have when wielding a sluggish util:

```go
  sp := spinner.New()
  for i := 0; i < 99; i++ {
    sp.Spin()
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(99)))
  }

  fmt.Printf("%d operations in %.2f seconds\n", sp.Count, sp.Elapsed())
```
