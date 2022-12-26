package hclient

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/eagle/handler/utils"
)

var (
	ErrNoCertificates = errors.New("no certificates in client")
)

// func GetDNSName(parameters []string) []string {
// 	for _, peercert := range state.PeerCertificates {
// 		der, err := x509.MarshalPKIXPublicKey(peercert.PublicKey)
// 		if err != nil {
// 			log.Println("Failed to get public key (https).")
// 		}
// 		if len(peercert.DNSNames) > 0 {
// 			return peercert.DNSNames[0]
// 		}
// 	}
// 	return "Unknown Site"
// }

// NewClient creates a new http client
// Takes in the optional arguments: proxy, servername
func NewClient(parameters ...string) (*Client, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}
	//certificate template
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}

	//import cerficate sni.cloudflaressl.com.cer
	// cert, err := tls.LoadX509KeyPair("sni.cloudflaressl.com.cer", "sni.cloudflaressl.com.key")

	// dnsName := GetDNSName(parameters)
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Eagle-bot"},
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		log.Fatal("Failed to encode certificate to PEM")
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if pemKey == nil {
		log.Fatal("Failed to encode key to PEM")
	}

	cert, err := tls.X509KeyPair(pemCert, pemKey)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(pemCert); !ok {
		log.Fatal("Failed to append certificate to pool")
	}

	clientCert := []tls.Certificate{cert}

	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:            certPool,
					Certificates:       clientCert,
					InsecureSkipVerify: true,
				},
			},
		},
		LatestResponse: &Response{},
	}, nil

}

// NewRequest creates a new request under a specified http client
func (c *Client) NewRequest() *Request {
	return &Request{
		client: c,
		header: make(http.Header),
	}
}

func (c *Client) InitCookieJar() {
	if c.client.Jar == nil {
		c.client.Jar, _ = cookiejar.New(nil)
	}
}

// InitSessionJar creates session jar, returns if it already existed or not
// func (c *Client) InitSessionJar(account *account.Account) bool {
// 	didExist := sessions.DoesSessionExist(account)

// 	jar, err := sessionjar.New(&sessionjar.Options{
// 		Filename: fmt.Sprintf("../.sessions/%s/%s.sessions", strings.Replace(utils.SiteIDtoSiteString[account.SiteId], "@", "", -1), account.Email),
// 	})

// 	if err != nil {
// 		fmt.Println("Failed to initialize session. ", err)
// 		return false
// 	}

// 	c.jar = jar
// 	c.client.Jar = jar
// 	return didExist
// }

func (c *Client) SaveCookies() {
	if c.client.Jar != nil {
		err := c.jar.Save()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// AddCookie adds a new cookie to the request client cookie jar
func (c *Client) AddCookie(u *url.URL, cookie *http.Cookie) error {
	if c.client.Jar == nil {
		c.client.Jar, _ = cookiejar.New(nil)
	}

	currentCookies := c.client.Jar.Cookies(u)
	currentCookies = append(currentCookies, cookie)
	c.client.Jar.SetCookies(u, currentCookies)

	return nil
}

// RemoveCookie removes the specified cookie from the request client cookie jar
func (c *Client) RemoveCookie(u *url.URL, cookie string) error {
	if c.client.Jar == nil {
		c.client.Jar, _ = cookiejar.New(nil)
	}

	newCookie := &http.Cookie{
		Name:  cookie,
		Value: "",
	}

	c.client.Jar.SetCookies(u, []*http.Cookie{newCookie})

	return nil
}

func (c *Client) AddCookieByName(r *Response, u *url.URL, name string) error {
	cookie := r.GetCookieByName(name)
	if cookie != nil {
		err := c.AddCookie(u, cookie)
		if err != nil {
			return err
		}
	}

	return nil
}

// Do will send the specified request
func (c *Client) Do(r *http.Request) (*Response, error) {
	fmt.Print("Sending request: ", r.URL.String())
	resp, err := c.client.Do(r)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// https://help.socketlabs.com/docs/how-to-fix-error-only-one-usage-of-each-socket-address-protocolnetwork-addressport-is-normally-permitted
	// https://www.geeksforgeeks.org/http-headers-connection/#:~:text=close%20This%20close%20connection%20directive,want%20your%20connection%20to%20close.
	r.Close = true // perhaps set this to false?

	response := &Response{
		headers:    resp.Header,
		body:       body,
		status:     resp.Status,
		statusCode: resp.StatusCode,
		cookies:    resp.Cookies(),
	}
	c.LatestResponse = response
	if utils.Debug {
		fmt.Println(fmt.Sprintf("%s %s", r.Method, r.URL.String()))
		fmt.Println(fmt.Sprintf("Response Body: %s", response.BodyAsString()))
	}

	return response, nil
}
