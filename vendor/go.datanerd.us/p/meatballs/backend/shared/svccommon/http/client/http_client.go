package svccommon_http_client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func GetHttpClient(certFile string, certDirectory string, httpTimeout time.Duration) *http.Client {
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
			log.Print("Cert could not be appended")
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
					log.Print("Cert could not be appended")
				}
			}
		}
	}
	return caCertPool
}
