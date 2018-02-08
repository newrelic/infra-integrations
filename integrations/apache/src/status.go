// Package status encapsulates the instantiation and configuration of the Apache status client
package main

import (
	"net/http"

	"time"

	"go.datanerd.us/p/meatballs/backend/shared/svccommon/http/client"
)

// Status will create new HTTP clients based on its configuration
type Status struct {
	CABundleFile string
	CABundleDir  string
	HTTPTimeout  time.Duration
}

// NewClient creates a new http.Client based on the provider's configuration
func (p Status) NewClient() *http.Client {
	return svccommon_http_client.GetHttpClient(p.CABundleFile, p.CABundleDir, p.HTTPTimeout)
}
