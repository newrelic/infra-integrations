package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/sdk"
)

var testGetModulesCorrectFormat = `Loaded Modules:
 http_module (static)
 access_compat_module (shared)
 actions_module (shared)
 alias_module (shared)
 core_module (static)
 so_module (static)
`
var testGetModulesDifferentLinesFormat = `Loaded Modules:
 core_module (static)
 so_module (static)
 http_module (static)
 random text
 alias_module (shared)
 :
`
var testGetVersionCorrectFormat = `Server version: Apache/2.4.6 (CentOS)
Server built:   Nov 14 2016 18:04:44
Server's Module Magic Number: 20120211:24
Server loaded:  APR 1.4.8, APR-UTIL 1.5.2
Compiled using: APR 1.4.8, APR-UTIL 1.5.2
Architecture:   64-bit
Server MPM:     prefork
  threaded:     no
    forked:     yes (variable process count)
Server compiled with....
 -D APR_HAS_SENDFILE
`
var testGetVersionDifferentLinesFormat = `Random text
Server built:   Nov 14 2016 18:04:44
Server version: Apache/2.4.6 (CentOS)
Server MPM:     prefork
  threaded:     no
    forked:     yes (variable process count)
 :
`
var testWrongLinesFormat = `
Random text
Random text
:
`
var testEmptyInput = ``

func TestParseGetModuleCorrectFormat(t *testing.T) {
	inventory := make(sdk.Inventory)

	err := getModules(bufio.NewReader(strings.NewReader(testGetModulesCorrectFormat)), inventory)

	if len(inventory) != 6 {
		t.Error()
	}
	if inventory["modules/access_compat"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/alias"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/actions"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/http"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/so"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/core"]["value"] != "enabled" {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetModuleDifferentLinesFormat(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getModules(bufio.NewReader(strings.NewReader(testGetModulesDifferentLinesFormat)), inventory)

	if len(inventory) != 4 {
		t.Error()
	}
	if inventory["modules/alias"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/http"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/so"]["value"] != "enabled" {
		t.Error()
	}
	if inventory["modules/core"]["value"] != "enabled" {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetModulesWrongLinesFormat(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getModules(bufio.NewReader(strings.NewReader(testWrongLinesFormat)), inventory)

	if len(inventory) != 0 {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetModulesEmptyInput(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getModules(bufio.NewReader(strings.NewReader(testEmptyInput)), inventory)

	if len(inventory) != 0 {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetVersionCorrectFormat(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getVersion(bufio.NewReader(strings.NewReader(testGetVersionCorrectFormat)), inventory)

	if len(inventory) != 1 {
		t.Error()
	}
	if inventory["version"]["value"] != "Apache/2.4.6 (CentOS)" {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetVersionDifferentLinesFormat(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getVersion(bufio.NewReader(strings.NewReader(testGetVersionDifferentLinesFormat)), inventory)

	if len(inventory) != 1 {
		t.Error()
	}
	if inventory["version"]["value"] != "Apache/2.4.6 (CentOS)" {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetVersionWrongLinesFormat(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getVersion(bufio.NewReader(strings.NewReader(testWrongLinesFormat)), inventory)

	if len(inventory) != 0 {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}

func TestParseGetVersionEmptyInput(t *testing.T) {
	inventory := make(sdk.Inventory)
	err := getVersion(bufio.NewReader(strings.NewReader(testEmptyInput)), inventory)

	if len(inventory) != 0 {
		t.Error()
	}
	if err != nil {
		t.Error()
	}
}
