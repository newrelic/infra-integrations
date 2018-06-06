package main

import (
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"

	"github.com/newrelic/infra-integrations-sdk/data/inventory"
)

func getInventory() (map[string]interface{}, error) {
	rawYamlFile, err := ioutil.ReadFile(args.ConfigPath)
	if err != nil {
		return nil, err
	}

	i := make(map[string]interface{})
	err = yaml.Unmarshal(rawYamlFile, &i)

	return i, err
}

func populateInventory(i *inventory.Inventory, rawInventory map[string]interface{}) error {
	for k, v := range rawInventory {
		switch value := v.(type) {
		case map[interface{}]interface{}:
			for subk, subv := range value {
				switch subVal := subv.(type) {
				case []interface{}:
					//TODO: Do not include lists for now
				default:
					setValue(i, k, subk.(string), subVal)
				}
			}
		case []interface{}:
			//TODO: Do not include lists for now
		default:
			setValue(i, k, "value", value)
		}
	}
	return nil
}

func setValue(i *inventory.Inventory, key string, field string, value interface{}) {
	re, _ := regexp.Compile("(?i)password")

	if re.MatchString(key) || re.MatchString(field) {
		value = "(omitted value)"
	}
	i.SetItem(key, field, value)
}
