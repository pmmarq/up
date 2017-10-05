package main

import (
	"os"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"

	"github.com/apex/up"
	"github.com/apex/up/handler"
	"github.com/apex/up/internal/proxy"
	"github.com/apex/up/platform/lambda/runtime"
)

func main() {
	if s := os.Getenv("LOG_LEVEL"); s != "" {
		log.SetLevelFromString(s)
	}

	log.SetHandler(json.Default)
	stage := os.Getenv("UP_STAGE")
	log.WithField("stage", stage).Info("initialize")

	// read config
	c, err := up.ReadConfig("up.json")
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	// init project
	p := runtime.New(c)

	// init runtime
	if err := p.Init(stage); err != nil {
		log.Fatalf("error initializing: %s", err)
	}

	// init handler
	h, err := handler.New()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// serve
	apex.Handle(proxy.NewHandler(h))
}
