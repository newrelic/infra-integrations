package main

import (
	"regexp"

	"github.com/newrelic/infra-integrations-sdk/jmx"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/metric"
)

const columnFamiliesLimit = 20

// getMetrics will gather all node and keyspace level metrics and return them as two maps
// The main metrics map will contain all the keys got from JMX and the keyspace metrics map
// Will contain maps for each <keyspace>.<columnFamily> found while inspecting JMX metrics.
func getMetrics() (map[string]interface{}, map[string]map[string]interface{}, error) {
	internalKeyspaces := map[string]struct{}{
		"OpsCenter":          struct{}{},
		"system":             struct{}{},
		"system_auth":        struct{}{},
		"system_distributed": struct{}{},
		"system_schema":      struct{}{},
		"system_traces":      struct{}{},
	}
	metrics := make(map[string]interface{})
	columnFamilyMetrics := make(map[string]map[string]interface{})
	visitedColumnFamilies := make(map[string]struct{})

	re, err := regexp.Compile("keyspace=(.*),scope=(.*?),")
	if err != nil {
		return nil, nil, err
	}

	for _, query := range jmxPatterns {
		results, err := jmx.Query(query)
		if err != nil {
			return nil, nil, err
		}
		for key, value := range results {
			matches := re.FindStringSubmatch(key)
			key = re.ReplaceAllString(key, "")

			if len(matches) != 3 {
				metrics[key] = value
			} else {
				columnfamily := matches[2]
				keyspace := matches[1]
				eventkey := keyspace + "." + columnfamily

				_, found := internalKeyspaces[keyspace]
				if !found {
					_, found := visitedColumnFamilies[eventkey]
					if !found {
						if len(visitedColumnFamilies) < columnFamiliesLimit {
							visitedColumnFamilies[eventkey] = struct{}{}
						} else {
							continue
						}
					}

					_, ok := columnFamilyMetrics[eventkey]
					if !ok {
						columnFamilyMetrics[eventkey] = make(map[string]interface{})
						columnFamilyMetrics[eventkey]["keyspace"] = keyspace
						columnFamilyMetrics[eventkey]["columnFamily"] = columnfamily
						columnFamilyMetrics[eventkey]["keyspaceAndColumnFamily"] = eventkey
					}
					columnFamilyMetrics[eventkey][key] = value
				}

			}
		}
	}

	return metrics, columnFamilyMetrics, nil
}

func populateMetrics(sample *metric.MetricSet, metrics map[string]interface{}, definition map[string][]interface{}) {
	notFoundMetrics := make([]string, 0)
	for metricName, metricConf := range definition {
		rawSource := metricConf[0]
		metricType := metricConf[1].(metric.SourceType)

		var rawMetric interface{}
		var ok bool

		switch source := rawSource.(type) {
		case string:
			rawMetric, ok = metrics[source]
			percentileRe, err := regexp.Compile("attr=.*Percentile")
			if err != nil {
				continue
			}
			if rawMetric != nil && percentileRe.MatchString(source) {
				// Convert percentiles from microseconds to milliseconds
				rawMetric = rawMetric.(float64) / 1000.0
			}
		case func(map[string]interface{}) (float64, bool):
			rawMetric, ok = source(metrics)
		default:
			log.Debug("Invalid raw source metric for %s", metricName)
			continue
		}

		if !ok {
			notFoundMetrics = append(notFoundMetrics, metricName)

			continue
		}

		err := sample.SetMetric(metricName, rawMetric, metricType)
		if err != nil {
			log.Warn("Error setting value: %s", err)
			continue
		}
	}
	if len(notFoundMetrics) > 0 {
		log.Debug("Can't find raw metrics in results for keys: %v", notFoundMetrics)
	}

}
