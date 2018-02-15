// Package status encapsulates the instantiation and configuration of the Apache status client
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"path/filepath"
	"strings"

	"time"

	"github.com/newrelic/infra-integrations-sdk/log"
)

// Status will create new HTTP client based on its configuration
type Status struct {
	CABundleFile string
	CABundleDir  string
	HTTPTimeout  time.Duration
}

// NewClient creates a new http.Client based on the provider's configuration
func (p Status) NewClient() *http.Client {
	return httpClient(p.CABundleFile, p.CABundleDir, p.HTTPTimeout)
}

func httpClient(certFile string, certDirectory string, httpTimeout time.Duration) *http.Client {
	// go default http transport settings
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if certFile != "" || certDirectory != "" {
		transport.TLSClientConfig = &tls.Config{RootCAs: getCertPool(certFile, certDirectory)}
	}

	return &http.Client{
		Timeout:   httpTimeout * time.Second,
		Transport: transport,
	}
}

func getCertPool(certFile string, certDirectory string) *x509.CertPool {
	caCertPool := x509.NewCertPool()
	if certFile != "" {
		caCert, err := ioutil.ReadFile(certFile)
		if err != nil {
			log.Fatal(err)
		}

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			log.Debug("Cert %q could not be appended", certFile)
		}
	}
	if certDirectory != "" {
		files, err := ioutil.ReadDir(certDirectory)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if strings.Contains(f.Name(), ".pem") {
				caCertFilePath := filepath.Join(certDirectory + "/" + f.Name())
				caCert, err := ioutil.ReadFile(caCertFilePath)
				if err != nil {
					log.Fatal(err)
				}
				ok := caCertPool.AppendCertsFromPEM(caCert)
				if !ok {
					log.Debug("Cert %q could not be appended", caCertFilePath)
				}
			}
		}
	}
	return caCertPool
}
