package main

import (
	_ "expvar"
	"flag"
	"net/http"
	"runtime"
	"time"

	"github.com/robertkrimen/otto"

	log "github.com/Sirupsen/logrus"

	"go.uber.org/zap"

	newrelic "github.com/newrelic/go-agent"
	// _ "github.com/yuin/gopher-lua"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// or kingpin for nicer CLI

// VERSION gets overwritten by release target
var VERSION = "dev"
var _ = VERSION

var ip = flag.Int("flagname", 1234, "help message for flagname")

var debug = kingpin.Flag("debug", "Enable debug mode.").Bool()

func main() {
	kingpin.Parse()

	logger, _ := zap.NewProduction()
	config := newrelic.NewConfig("Your Application Name", "__YOUR_NEW_RELIC_LICENSE_KEY__")
	newrelic.NewApplication(config)
	vm := otto.New()

	go func() {
		http.ListenAndServe(":5160", nil)
	}()

	for {
		f := NewFarble(&Counter{})

		f.Bumple()
		vm.Run(`
			abc = 2 + 2;
			console.log("\nThe value of abc is " + abc); // 4
		`)
		//kingpin.Parse()
		logger.Info("OK", zap.Int("ip", *ip))
		log.Info("OK")

		runtime.GC()
		time.Sleep(time.Second * 1)
	}

}
