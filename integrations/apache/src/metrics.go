package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/metric"
)

const (
	headerContentType = "Content-Type"
	expectedMimeType  = "text/plain"
)

var metricsDefinition = map[string][]interface{}{
	"software.version":                     {"Server version", metric.ATTRIBUTE},
	"net.requestsPerSecond":                {"Total Accesses", metric.RATE},
	"net.bytesPerSecond":                   {getBytes, metric.RATE},
	"server.idleWorkers":                   {"IdleWorkers", metric.GAUGE},
	"server.busyWorkers":                   {"BusyWorkers", metric.GAUGE},
	"server.scoreboard.writingWorkers":     {getWorkerStatus("W"), metric.GAUGE},
	"server.scoreboard.loggingWorkers":     {getWorkerStatus("L"), metric.GAUGE},
	"server.scoreboard.finishingWorkers":   {getWorkerStatus("G"), metric.GAUGE},
	"server.scoreboard.readingWorkers":     {getWorkerStatus("R"), metric.GAUGE},
	"server.scoreboard.closingWorkers":     {getWorkerStatus("C"), metric.GAUGE},
	"server.scoreboard.keepAliveWorkers":   {getWorkerStatus("K"), metric.GAUGE},
	"server.scoreboard.dnsLookupWorkers":   {getWorkerStatus("D"), metric.GAUGE},
	"server.scoreboard.idleCleanupWorkers": {getWorkerStatus("I"), metric.GAUGE},
	"server.scoreboard.startingWorkers":    {getWorkerStatus("S"), metric.GAUGE},
	"server.scoreboard.totalWorkers":       {getTotalWorkers, metric.GAUGE},
}

func asValue(value string) interface{} {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}

	if b, err := strconv.ParseBool(value); err == nil {
		return b
	}
	return value
}

func populateMetrics(sample *metric.MetricSet, metrics map[string]interface{}, metricsDefinition map[string][]interface{}) error {
	for metricName, metricInfo := range metricsDefinition {
		rawSource := metricInfo[0]
		metricType := metricInfo[1].(metric.SourceType)

		var rawMetric interface{}
		var ok bool

		switch source := rawSource.(type) {
		case string:
			rawMetric, ok = metrics[source]
		case func(map[string]interface{}) (float64, bool):
			rawMetric, ok = source(metrics)
		default:
			log.Warn("Invalid raw source metric for %s", metricName)
			continue
		}

		if !ok {
			log.Warn("Can't find raw metrics in results for %s", metricName)
			continue
		}
		err := sample.SetMetric(metricName, rawMetric, metricType)

		if err != nil {
			log.Warn("Error setting value: %s", err)
			continue
		}
	}

	if len(*sample) < 2 {
		return fmt.Errorf("no metrics were found on the status response. Probably caused by a wrong response format")
	}
	return nil
}

// getWorkerStatus counts occurence of a given letter, which means status of a worker
// (i.e. "W" corresponds to writing status of the worker)
func getWorkerStatus(status string) func(metrics map[string]interface{}) (float64, bool) {
	return func(metrics map[string]interface{}) (float64, bool) {
		scoreboard, ok := metrics["Scoreboard"].(string)
		if ok {
			return float64(strings.Count(scoreboard, status)), true
		}
		return 0, false
	}
}

// getTotalWorkers counts number of characters for Scoreboard key, which means total number of workers
func getTotalWorkers(metrics map[string]interface{}) (float64, bool) {
	scoreboard, ok := metrics["Scoreboard"].(string)
	if ok {
		return float64(len(scoreboard)), true
	}
	return 0, false
}

//getBytes converts value of Total kBytes into bytes
func getBytes(metrics map[string]interface{}) (float64, bool) {
	totalkBytes, ok := metrics["Total kBytes"].(int)
	if ok {
		return float64(totalkBytes * 1024), true
	}
	return 0, false
}

// getRawMetrics reads an Apache status message and transforms its
// contents into a map of metrics with the keys and values extracted from the
// status endpoint.
func getRawMetrics(reader *bufio.Reader) (map[string]interface{}, error) {
	metrics := make(map[string]interface{})

	_, err := reader.Peek(1)
	if err != nil {
		return nil, fmt.Errorf("Empty result")
	}

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		splitedLine := strings.Split(line, ":")
		if len(splitedLine) != 2 {
			continue
		}
		metrics[splitedLine[0]] = asValue(strings.TrimSpace(splitedLine[1]))
	}
	if len(metrics) == 0 {
		log.Debug("Metrics data from status module not found")
	}
	return metrics, nil
}

func getMetricsData(status *Status, sample *metric.MetricSet) error {
	netClient := status.NewClient()

	log.Debug("retrieving Apache Status from %s", args.StatusURL)
	resp, err := netClient.Get(args.StatusURL)
	if err != nil {
		log.Fatal(fmt.Errorf("Error trying to connect to '%s'. Got error: %v", args.StatusURL, err.Error()))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned code %d (%s). Expecting 200", resp.StatusCode, resp.Status)
	}

	if !strings.Contains(resp.Header.Get(headerContentType), expectedMimeType) {
		return fmt.Errorf("apache Status endpoint returned an invalid content type. Expected '%s'. Got '%s'",
			expectedMimeType, resp.Header.Get(headerContentType))
	}

	var rawMetrics map[string]interface{}
	rawMetrics, err = getRawMetrics(bufio.NewReader(resp.Body))
	if err != nil {
		return err
	}

	version := resp.Header.Get("Server")
	if version != "" {
		rawMetrics["Server version"] = version
	} else {
		log.Debug("Software version not found")
	}

	return populateMetrics(sample, rawMetrics, metricsDefinition)
}
