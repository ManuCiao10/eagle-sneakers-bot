package client

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/eagle/handler/utils"
)

//add clinet in task
//add task.DoneTask status

// NewClient creates a new client
func NewClient(parameters ...string) (*Client, error) {
	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: true,
		VerifyConnection: func(state tls.ConnectionState) error {
			for _, peercert := range state.PeerCertificates {
				der, err := x509.MarshalPKIXPublicKey(peercert.PublicKey)
				if err != nil {
					log.Println("Failed to get public key (https).")
				}

				var DNSName string
				if len(peercert.DNSNames) > 0 {
					DNSName = peercert.DNSNames[0]
				} else {
					DNSName = "Unknown Site"
				}

				hash := sha256.Sum256(der)
				stringHash := fmt.Sprintf("%x", hash)

				if utils.Debug {
					fmt.Printf("%s: %s", DNSName, stringHash)
				}

				if fingerprints[stringHash] == 1 {
					return nil
				} else {
					fmt.Println(DNSName + ": SSL mismatch.")
				}
			}
			return fmt.Errorf("invalid certificate")
		},
	}

	// parameters[0] = proxy
	// parameters[1] = sni
	if len(parameters) > 1 && len(parameters[1]) > 0 {
		tlsClientConfig.ServerName = parameters[1]
	}

	transport := &http.Transport{
		ForceAttemptHTTP2: true,
		TLSClientConfig:   tlsClientConfig,
	}

	if len(parameters) > 0 && len(parameters[0]) > 0 {
		proxyUrl, _ := url.Parse(parameters[0])

		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	return &Client{
		client: &http.Client{
			Transport: transport,
		},
		LatestResponse: &Response{},
	}, nil

}
