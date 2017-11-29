package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/metric"
)

var testApacheStatus = `Total Accesses: 66
Total kBytes: 73
Uptime: 31006
ReqPerSec: .00212862
BytesPerSec: 2.41089
BytesPerReq: 1132.61
BusyWorkers: 1
IdleWorkers: 4
Scoreboard: _W___......_CDCDII.II......KKKKKGG................__R_W.....S.....LS
`

var testApacheStatusWrongLinesFormat = `
Random text
Random text

`

var testApacheStatusEmpty = ``

func TestGetRawMetrics(t *testing.T) {
	rawMetrics, err := getRawMetrics(bufio.NewReader(strings.NewReader(testApacheStatus)))

	if len(rawMetrics) != 9 {
		t.Error()
	}

	if rawMetrics["Total Accesses"] != 66 {
		t.Error()
	}
	if rawMetrics["Uptime"] != 31006 {
		t.Error()
	}
	if rawMetrics["ReqPerSec"] != 0.00212862 {
		t.Error()
	}
	if rawMetrics["BytesPerSec"] != 2.41089 {
		t.Error()
	}
	if rawMetrics["IdleWorkers"] != 4 {
		t.Error()
	}
	if rawMetrics["BusyWorkers"] != 1 {
		t.Error()
	}
	if rawMetrics["Total kBytes"] != 73 {
		t.Error()
	}
	if rawMetrics["BytesPerReq"] != 1132.61 {
		t.Error()
	}
	if rawMetrics["Scoreboard"] != "_W___......_CDCDII.II......KKKKKGG................__R_W.....S.....LS" {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestGetMetricsInvalidData(t *testing.T) {
	rawMetrics, err := getRawMetrics(bufio.NewReader(strings.NewReader(testApacheStatusWrongLinesFormat)))

	if err != nil {
		t.Error()
	}
	if len(rawMetrics) != 0 {
		t.Error()
	}
}

func TestGetMetricsEmptyData(t *testing.T) {
	rawMetrics, err := getRawMetrics(bufio.NewReader(strings.NewReader(testApacheStatusEmpty)))

	if !reflect.DeepEqual(err.Error(), "Empty result") {
		t.Error()
	}
	if len(rawMetrics) != 0 {
		t.Error()
	}
}

func TestGetWorkerStatus(t *testing.T) {
	metrics := map[string]interface{}{
		"Scoreboard": "_W___......_DDII.II......KKKKKGG................__R_W.....S.....LS",
	}
	writingWorkersNumber, ok := getWorkerStatus("W")(metrics)
	if !ok {
		t.Error()
	}
	if writingWorkersNumber != float64(2) {
		t.Error()
	}

	closingWorkersNumber, ok := getWorkerStatus("C")(metrics)
	if !ok {
		t.Error()
	}
	if closingWorkersNumber != float64(0) {
		t.Error()
	}
}

func TestGetWorkerStatusInvalidDataKey(t *testing.T) {
	metrics := map[string]interface{}{
		"Total kBytes": "_W___......_CDCDII.II......KKKKKGG................__R_W.....S.....LS",
	}
	closingWorkersNumber, ok := getWorkerStatus("C")(metrics)
	if ok {
		t.Error()
	}
	if closingWorkersNumber != float64(0) {
		t.Error()
	}
}

func TestGetWorkerStatusInvalidDataType(t *testing.T) {
	metrics := map[string]interface{}{
		"Scoreboard": 5,
	}
	closingWorkersNumber, ok := getWorkerStatus("C")(metrics)
	if ok {
		t.Error()
	}
	if closingWorkersNumber != float64(0) {
		t.Error()
	}
}

func TestGetTotalWorkers(t *testing.T) {
	metrics := map[string]interface{}{
		"Scoreboard": "_W___......_DDII.II......KKKKKGG................__R_W.....S.....LS",
	}
	totalWorkersNumber, ok := getTotalWorkers(metrics)
	if !ok {
		t.Error()
	}
	if totalWorkersNumber != float64(66) {
		t.Error()
	}
}

func TestGetTotalWorkersInvalidDataKey(t *testing.T) {
	metrics := map[string]interface{}{
		"Total kBytes": "_W___......_CDCDII.II......KKKKKGG................__R_W.....S.....LS",
	}
	totalWorkersNumber, ok := getTotalWorkers(metrics)
	if ok {
		t.Error()
	}
	if totalWorkersNumber != float64(0) {
		t.Error()
	}
}

func TestGetTotalWorkersInvalidDataType(t *testing.T) {
	metrics := map[string]interface{}{
		"Scoreboard": 5,
	}
	totalWorkersNumber, ok := getTotalWorkers(metrics)
	if ok {
		t.Error()
	}
	if totalWorkersNumber != float64(0) {
		t.Error()
	}
}

func TestGetBytes_IntData(t *testing.T) {
	metrics := map[string]interface{}{
		"Total kBytes": 67,
	}
	totalBytes, ok := getBytes(metrics)
	if !ok {
		t.Error()
	}
	if totalBytes != float64(68608) {
		t.Error()
	}
}

func TestGetBytes_InvalidDataType(t *testing.T) {
	metrics := map[string]interface{}{
		"Total kBytes": 67.4,
	}
	totalBytes, ok := getBytes(metrics)
	if ok {
		t.Error()
	}
	if totalBytes != float64(0) {
		t.Error()
	}
}

func TestGetBytes_InvalidDataKey(t *testing.T) {
	metrics := map[string]interface{}{
		"TotalkBytes": 67,
	}
	totalBytes, ok := getBytes(metrics)
	if ok {
		t.Error()
	}
	if totalBytes != float64(0) {
		t.Error()
	}
}

func TestPopulateMetrics(t *testing.T) {
	// Given an Apache Status
	rawMetrics, _ := getRawMetrics(bufio.NewReader(strings.NewReader(testApacheStatus)))

	// When the system populates the metrics from the Apache Status
	populatedMetrics := metric.MetricSet{}
	err := populateMetrics(&populatedMetrics, rawMetrics, metricsDefinition)

	metricsSet := map[string]interface{}(populatedMetrics)

	// They populated metrics values are correct
	if err != nil {
		t.Error(err)
	}

	// TODO: use assertions library for tests
	if len(metricsSet) != 14 {
		t.Errorf("metricsSet length = %d. Expected 14", len(metricsSet))
	}
	if bytes, _ := getBytes(rawMetrics); bytes != float64(73*1024) {
		t.Errorf("getBytes = %f. Expected 73*1024", bytes)
	}
	if metricsSet["server.idleWorkers"] != 4 {
		t.Errorf("server.idleWorkers = %d. Expected 4", metricsSet["server.idleWorkers"])
	}
	if metricsSet["server.busyWorkers"] != 1 {
		t.Errorf("server.busyWorkers = %d. Expected 1", metricsSet["server.busyWorkers"])
	}
	const expectedScoreBoard = "_W___......_CDCDII.II......KKKKKGG................__R_W.....S.....LS"
	if metricsSet["server.scoreboard.totalWorkers"] != float64(len(expectedScoreBoard)) {
		t.Errorf("server.scoreboard.totalWorkers = %f. Expected %d",
			metricsSet["server.scoreboard.totalWorkers"], len(expectedScoreBoard))
	}
	if metricsSet["server.scoreboard.writingWorkers"] != float64(2) {
		t.Errorf("server.scoreboard.writingWorkers = %f. Expected 2", metricsSet["server.scoreboard.writingWorkers"])
	}

	if metricsSet["server.scoreboard.loggingWorkers"] != float64(1) {
		t.Errorf("server.scoreboard.loggingWorkers = %f. Expected 1", metricsSet["server.scoreboard.loggingWorkers"])
	}

	if metricsSet["server.scoreboard.finishingWorkers"] != float64(2) {
		t.Errorf("server.scoreboard.finishingWorkers = %f. Expected 2", metricsSet["server.scoreboard.finishingWorkers"])
	}

	if metricsSet["server.scoreboard.readingWorkers"] != float64(1) {
		t.Errorf("server.scoreboard.readingWorkers = %f. Expected 1", metricsSet["server.scoreboard.readingWorkers"])
	}

	if metricsSet["server.scoreboard.closingWorkers"] != float64(2) {
		t.Errorf("server.scoreboard.closingWorkers = %f. Expected 2", metricsSet["server.scoreboard.closingWorkers"])
	}

	if metricsSet["server.scoreboard.keepAliveWorkers"] != float64(5) {
		t.Errorf("server.scoreboard.keepAliveWorkers = %f. Expected 5", metricsSet["server.scoreboard.keepAliveWorkers"])
	}

	if metricsSet["server.scoreboard.dnsLookupWorkers"] != float64(2) {
		t.Errorf("server.scoreboard.dnsLookupWorkers = %f. Expected 2", metricsSet["server.scoreboard.dnsLookupWorkers"])
	}

	if metricsSet["server.scoreboard.idleCleanupWorkers"] != float64(4) {
		t.Errorf("server.scoreboard.idleCleanupWorkers = %f. Expected 4", metricsSet["server.scoreboard.idleCleanupWorkers"])
	}

	if metricsSet["server.scoreboard.startingWorkers"] != float64(2) {
		t.Errorf("server.scoreboard.startingWorkers = %f. Expected 2", metricsSet["server.scoreboard.startingWorkers"])
	}
}

func TestPopulateInvalidMetricsFormat(t *testing.T) {
	// Given an invalid format for the Apache Status
	rawMetrics, _ := getRawMetrics(bufio.NewReader(strings.NewReader("some invalid\nstring is here:\nhello!")))

	// When the system populates the metrics from the Apache Status
	populatedMetrics := metric.MetricSet{}
	err := populateMetrics(&populatedMetrics, rawMetrics, metricsDefinition)

	metricsSet := map[string]interface{}(populatedMetrics)

	// Then an error is returned
	if err == nil {
		t.Error("populateMetrics should return an error")
	}

	// As no metrics are set
	if len(metricsSet) != 0 {
		t.Errorf("no metrics should be set and there are %d", len(metricsSet))
	}
}
