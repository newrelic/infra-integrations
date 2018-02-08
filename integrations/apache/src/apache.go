package main

import (
	"time"

	sdk_args "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/sdk"
)

type argumentList struct {
	sdk_args.DefaultArgumentList
	StatusURL  string `default:"http://127.0.0.1/server-status?auto" help:"Apache status-server URL."`
	CACertFile string `default:"" help:"Apache server alternative Certificate Authority certificate file"`
	CACertDir  string `default:"" help:"Apache server alternative Certificate Authority certificate directory"`
}

const (
	integrationName    = "com.newrelic.apache"
	integrationVersion = "1.1.0"

	defaultHTTPTimeout = time.Second * 1
)

var args argumentList

func main() {
	log.Debug("Starting Apache integration")
	defer log.Debug("Apache integration exited")

	integration, err := sdk.NewIntegration(integrationName, integrationVersion, &args)
	fatalIfErr(err)

	if args.All || args.Inventory {
		log.Debug("Fetching data for '%s' integration", integrationName+"-inventory")
		fatalIfErr(setInventory(integration.Inventory))
	}

	if args.All || args.Metrics {
		log.Debug("Fetching data for '%s' integration", integrationName+"-metrics")
		ms := integration.NewMetricSet("ApacheSample")
		provider := &Status{
			CACertDir:   args.CACertDir,
			CACertFile:  args.CACertFile,
			HTTPTimeout: defaultHTTPTimeout,
		}
		fatalIfErr(getMetricsData(provider, ms))
	}

	fatalIfErr(integration.Publish())
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
