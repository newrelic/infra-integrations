package main

import (
	"io/ioutil"
	"regexp"

	yaml "gopkg.in/yaml.v2"

	"github.com/newrelic/infra-integrations-sdk/sdk"
)

func getInventory() (map[string]interface{}, error) {
	rawYamlFile, err := ioutil.ReadFile(args.ConfigPath)
	if err != nil {
		return nil, err
	}

	inventory := make(map[string]interface{})
	err = yaml.Unmarshal(rawYamlFile, &inventory)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func populateInventory(inventory sdk.Inventory, rawInventory map[string]interface{}) error {
	for k, v := range rawInventory {
		switch value := v.(type) {
		case map[interface{}]interface{}:
			for subk, subv := range value {
				switch subVal := subv.(type) {
				case []interface{}:
					//TODO: Do not include lists for now
				default:
					setValue(inventory, k, subk.(string), subVal)
				}
			}
		case []interface{}:
			//TODO: Do not include lists for now
		default:
			setValue(inventory, k, "value", value)
		}
	}
	return nil
}

func setValue(inventory sdk.Inventory, key string, field string, value interface{}) {
	re, _ := regexp.Compile("(?i)password")

	if re.MatchString(key) || re.MatchString(field) {
		value = "(omitted value)"
	}
	inventory.SetItem(key, field, value)
}
