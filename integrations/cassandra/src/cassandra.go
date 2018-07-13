package main

import (
	"strconv"

	"os"

	sdk_args "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/jmx"
	"github.com/newrelic/infra-integrations-sdk/log"
)

type argumentList struct {
	sdk_args.DefaultArgumentList

	Hostname            string `default:"localhost" help:"Hostname or IP where Cassandra is running."`
	Port                int    `default:"7199" help:"Port on which JMX server is listening."`
	Username            string `default:"" help:"Username for accessing JMX."`
	Password            string `default:"" help:"Password for the given user."`
	ConfigPath          string `default:"/etc/cassandra/cassandra.yaml" help:"Cassandra configuration file."`
	Timeout             int    `default:"2000" help:"Timeout in milliseconds per single JMX query."`
	ColumnFamiliesLimit int    `default:"20" help:"Limit on number of Cassandra Column Families."`
}

const (
	integrationName    = "com.newrelic.cassandra"
	integrationVersion = "2.0.0-beta"
)

var (
	args argumentList
)

func main() {
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		panic(err)
	}

	e := i.LocalEntity()
	l := i.Logger()

	fatalIfErr(l, jmx.Open(args.Hostname, strconv.Itoa(args.Port), args.Username, args.Password))
	defer jmx.Close()

	if args.All() || args.Metrics {
		rawMetrics, allColumnFamilies, err := getMetrics()
		fatalIfErr(l, err)

		s := e.NewMetricSet("CassandraSample", metricSetAttributes...)

		populateMetrics(l, s, rawMetrics, metricsDefinition)
		populateMetrics(l, s, rawMetrics, commonDefinition)

		for _, columnFamilyMetrics := range allColumnFamilies {
			s := e.NewMetricSet("CassandraColumnFamilySample", metricSetAttributes...)
			populateMetrics(l, s, columnFamilyMetrics, columnFamilyDefinition)
			populateMetrics(l, s, rawMetrics, commonDefinition)
		}
	}

	if args.All() || args.Inventory {
		rawInventory, err := getInventory()
		fatalIfErr(l, err)
		populateInventory(e.Inventory, rawInventory)
	}

	fatalIfErr(l, i.Publish())
}

func fatalIfErr(l log.Logger, err error) {
	if err != nil {
		l.Errorf(err.Error())
		os.Exit(1)
	}
}
