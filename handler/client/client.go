// package client

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/http/cookiejar"
// 	"net/url"
// 	"time"

// 	"github.com/bogdanfinn/fhttp/http2"
// 	tls "github.com/bogdanfinn/utls"
// 	"github.com/eagle/handler/utils"
// 	"golang.org/x/net/proxy"
// )

// var (
// 	ErrFoo                = errors.New("no cookie jar in client")
// 	DefaultTimeoutSeconds = 30
// )

// type TransportOptions struct {
// 	DisableKeepAlives      bool
// 	DisableCompression     bool
// 	MaxIdleConns           int
// 	MaxIdleConnsPerHost    int
// 	MaxConnsPerHost        int
// 	MaxResponseHeaderBytes int64 // Zero means to use a default limit.
// 	WriteBufferSize        int   // If zero, a default (currently 4KB) is used.
// 	ReadBufferSize         int   // If zero, a default (currently 4KB) is used.
// }

// type httpClientConfig struct {
// 	debug                       bool
// 	followRedirects             bool
// 	ServerName                  string
// 	insecureSkipVerify          bool
// 	proxyUrl                    string
// 	serverNameOverwrite         string
// 	transportOptions            *TransportOptions
// 	cookieJar                   http.CookieJar
// 	clientProfile               ClientProfile
// 	withRandomTlsExtensionOrder bool
// 	forceHttp1                  bool
// 	timeout                     time.Duration
// }

// type directDialer struct {
// 	dialer net.Dialer
// }

// type ClientProfile struct {
// 	clientHelloId     tls.ClientHelloID
// 	settings          map[http2.SettingID]uint32
// 	settingsOrder     []http2.SettingID
// 	pseudoHeaderOrder []string
// 	connectionFlow    uint32
// 	priorities        []http2.Priority
// 	headerPriority    *http2.PriorityParam
// }

// func validateConfig(config *httpClientConfig) error {
// 	return nil
// }

// func newDirectDialer(timeout time.Duration) proxy.ContextDialer {
// 	return &directDialer{
// 		dialer: net.Dialer{
// 			Timeout: timeout,
// 		},
// 	}
// }

// func buildFromConfig(config *httpClientConfig) (*http.Client, ClientProfile, error) {
// 	var dialer proxy.ContextDialer
// 	dialer = newDirectDialer(config.timeout)

// 	if config.proxyUrl != "" {
// 		proxyDialer, err := newConnectDialer(config.proxyUrl, config.timeout)
// 		if err != nil {
// 			return nil, ClientProfile{}, err
// 		}

// 		dialer = proxyDialer
// 	}

// 	var redirectFunc func(req *http.Request, via []*http.Request) error
// 	if !config.followRedirects {
// 		redirectFunc = defaultRedirectFunc
// 	} else {
// 		redirectFunc = nil
// 	}

// 	clientProfile := config.clientProfile

// 	client := &http.Client{
// 		Timeout:       config.timeout,
// 		Transport:     newRoundTripper(clientProfile, config.transportOptions, config.serverNameOverwrite, config.insecureSkipVerify, config.withRandomTlsExtensionOrder, config.forceHttp1, dialer),
// 		CheckRedirect: redirectFunc,
// 	}

// 	if config.cookieJar != nil {
// 		client.Jar = config.cookieJar
// 	}

// 	return client, clientProfile, nil

// }

// type HttpClientOption func(config *httpClientConfig)

// // NewClient creates a new http client
// // Takes in the optional arguments: proxy, servername
// func NewClient(options ...HttpClientOption) (*Client, error) {
// 	config := &httpClientConfig{
// 		followRedirects: true,
// 		timeout:         time.Duration(DefaultTimeoutSeconds) * time.Second,
// 	}

// 	err := validateConfig(config)

// 	if err != nil {
// 		return nil, err
// 	}

// 	client, clientProfile, err := buildFromConfig(config)

// 	if err != nil {
// 		return nil, err
// 	}

// 	config.clientProfile = clientProfile
// 	// parameters[0] = proxy | parameters[1] = sni
// 	// if len(parameters) > 1 && len(parameters[1]) > 0 {
// 	// 	config.ServerName = parameters[1]
// 	// }

// 	// transport := &http.Transport{
// 	// 	Proxy:                 http.ProxyFromEnvironment,
// 	// 	DialContext:           utils.DialContext,
// 	// 	MaxIdleConns:          100,
// 	// 	IdleConnTimeout:       90 * time.Second,
// 	// 	TLSHandshakeTimeout:   10 * time.Second,
// 	// 	ExpectContinueTimeout: 1 * time.Second,
// 	// }

// 	// if len(parameters) > 0 && len(parameters[0]) > 0 {
// 	// 	proxyUrl, _ := url.Parse(parameters[0])

// 	// 	transport.Proxy = http.ProxyURL(proxyUrl)
// 	// }

// 	return &Client{
// 		Client:         *client,
// 		LatestResponse: &Response{},
// 	}, nil
// }

// // NewRequest creates a new request under a specified http client
// func (c *Client) NewRequest() *Request {
// 	return &Request{
// 		client: c,
// 		header: make(http.Header),
// 	}
// }

// func (c *Client) InitCookieJar() {
// 	if c.client.Jar == nil {
// 		c.client.Jar, _ = cookiejar.New(nil)
// 	}
// }

// // InitSessionJar creates session jar, returns if it already existed or not
// // func (c *Client) InitSessionJar(account *account.Account) bool {
// // 	// didExist := sessions.DoesSessionExist(account)

// // 	jar, err := sessionjar.New(&sessionjar.Options{
// // 		Filename: fmt.Sprintf("../.sessions/%s/%s.sessions", strings.Replace(utils.SiteIDtoSiteString[account.SiteId], "@", "", -1), account.Email),
// // 	})

// // 	if err != nil {
// // 		fmt.Println("Failed to initialize session. ", err)
// // 		return false
// // 	}

// // 	c.jar = jar
// // 	c.client.Jar = jar
// // 	return didExist
// // }

// func (c *Client) SaveCookies() {
// 	if c.client.Jar != nil {
// 		err := c.jar.Save()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// // AddCookie adds a new cookie to the request client cookie jar
// func (c *Client) AddCookie(u *url.URL, cookie *http.Cookie) error {
// 	if c.client.Jar == nil {
// 		c.client.Jar, _ = cookiejar.New(nil)
// 	}

// 	currentCookies := c.client.Jar.Cookies(u)
// 	currentCookies = append(currentCookies, cookie)
// 	c.client.Jar.SetCookies(u, currentCookies)

// 	return nil
// }

// // RemoveCookie removes the specified cookie from the request client cookie jar
// func (c *Client) RemoveCookie(u *url.URL, cookie string) error {
// 	if c.client.Jar == nil {
// 		c.client.Jar, _ = cookiejar.New(nil)
// 	}

// 	newCookie := &http.Cookie{
// 		Name:  cookie,
// 		Value: "",
// 	}

// 	c.client.Jar.SetCookies(u, []*http.Cookie{newCookie})

// 	return nil
// }

// func (c *Client) AddCookieByName(r *Response, u *url.URL, name string) error {
// 	cookie := r.GetCookieByName(name)
// 	if cookie != nil {
// 		err := c.AddCookie(u, cookie)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Do will send the specified request
// func (c *Client) Do(r *http.Request) (*Response, error) {
// 	resp, err := c.client.Do(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	_ = resp.Body.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// https://help.socketlabs.com/docs/how-to-fix-error-only-one-usage-of-each-socket-address-protocolnetwork-addressport-is-normally-permitted
// 	// https://www.geeksforgeeks.org/http-headers-connection/#:~:text=close%20This%20close%20connection%20directive,want%20your%20connection%20to%20close.
// 	r.Close = true // perhaps set this to false?

// 	response := &Response{
// 		headers:    resp.Header,
// 		body:       body,
// 		status:     resp.Status,
// 		statusCode: resp.StatusCode,
// 		cookies:    resp.Cookies(),
// 	}

// 	c.LatestResponse = response
// 	if utils.Debug {
// 		fmt.Println(fmt.Sprintf("%s %s", r.Method, r.URL.String()))
// 		fmt.Println(fmt.Sprintf("Response Body: %s", response.BodyAsString()))
// 	}

// 	return response, nil
// }

package client

import (
	"fmt"
	"net/url"
	"time"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/httputil"
	"golang.org/x/net/proxy"
)

func NewHttpClient(logger Logger, options ...HttpClientOption) (HttpClient, error) {
	config := &httpClientConfig{
		followRedirects: true,
		timeout:         time.Duration(DefaultTimeoutSeconds) * time.Second,
	}

	for _, opt := range options {
		opt(config)
	}

	err := validateConfig(config)

	if err != nil {
		return nil, err
	}

	client, clientProfile, err := buildFromConfig(config)

	if err != nil {
		return nil, err
	}

	config.clientProfile = clientProfile

	if config.debug {
		if logger == nil {
			logger = NewLogger()
		}

		logger = NewDebugLogger(logger)
	}

	if logger == nil {
		logger = NewNoopLogger()
	}

	return &HTTPClient{
		Client: *client,
		logger: logger,
		config: config,
	}, nil
}

func validateConfig(config *httpClientConfig) error {
	return nil
}

func buildFromConfig(config *httpClientConfig) (*http.Client, ClientProfile, error) {
	var dialer proxy.ContextDialer
	dialer = newDirectDialer(config.timeout)

	if config.proxyUrl != "" {
		proxyDialer, err := newConnectDialer(config.proxyUrl, config.timeout)
		if err != nil {
			return nil, ClientProfile{}, err
		}

		dialer = proxyDialer
	}

	var redirectFunc func(req *http.Request, via []*http.Request) error
	if !config.followRedirects {
		redirectFunc = defaultRedirectFunc
	} else {
		redirectFunc = nil
	}

	clientProfile := config.clientProfile

	client := &http.Client{
		Timeout:       config.timeout,
		Transport:     newRoundTripper(clientProfile, config.transportOptions, config.serverNameOverwrite, config.insecureSkipVerify, config.withRandomTlsExtensionOrder, config.forceHttp1, dialer),
		CheckRedirect: redirectFunc,
	}

	if config.cookieJar != nil {
		client.Jar = config.cookieJar
	}

	return client, clientProfile, nil
}

func (c *HTTPClient) SetFollowRedirect(followRedirect bool) {
	c.logger.Debug("set follow redirect from %v to %v", c.config.followRedirects, followRedirect)

	c.config.followRedirects = followRedirect
	c.applyFollowRedirect()
}

func (c *HTTPClient) GetFollowRedirect() bool {
	return c.config.followRedirects
}

func (c *HTTPClient) applyFollowRedirect() {
	if c.config.followRedirects {
		c.logger.Info("automatic redirect following is enabled")
		c.CheckRedirect = nil
	} else {
		c.logger.Info("automatic redirect following is disabled")
		c.CheckRedirect = defaultRedirectFunc
	}
}

func (c *HTTPClient) SetProxy(proxyUrl string) error {
	c.logger.Debug("set proxy from %s to %s", c.config.proxyUrl, proxyUrl)
	c.config.proxyUrl = proxyUrl
	c.logger.Info(fmt.Sprintf("set proxy to: %s", proxyUrl))

	return c.applyProxy()
}

func (c *HTTPClient) GetProxy() string {
	return c.config.proxyUrl
}

func (c *HTTPClient) applyProxy() error {
	var dialer proxy.ContextDialer
	dialer = proxy.Direct

	if c.config.proxyUrl != "" {
		c.logger.Debug("proxy url %s supplied - using proxy connect dialer", c.config.proxyUrl)
		proxyDialer, err := newConnectDialer(c.config.proxyUrl, c.config.timeout)
		if err != nil {
			c.logger.Error("failed to create proxy connect dialer: %s", err.Error())
			return err
		}

		dialer = proxyDialer
	}

	c.Transport = newRoundTripper(c.config.clientProfile, c.config.transportOptions, c.config.serverNameOverwrite, c.config.insecureSkipVerify, c.config.withRandomTlsExtensionOrder, c.config.forceHttp1, dialer)

	return nil
}

func (c *HTTPClient) GetCookies(u *url.URL) []*http.Cookie {
	c.logger.Info(fmt.Sprintf("get cookies for url: %s", u.String()))
	if c.Jar == nil {
		c.logger.Warn("you did not setup a cookie jar")
		return nil
	}

	return c.Jar.Cookies(u)
}

func (c *HTTPClient) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c.logger.Info(fmt.Sprintf("set cookies for url: %s", u.String()))

	if c.Jar == nil {
		c.logger.Warn("you did not setup a cookie jar")
		return
	}

	c.Jar.SetCookies(u, cookies)
}

func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.config.debug {
		requestBytes, err := httputil.DumpRequestOut(req, req.ContentLength > 0)

		if err != nil {
			return nil, err
		}

		c.logger.Debug("raw request bytes sent over wire: %d (%d kb)", len(requestBytes), len(requestBytes)/1024)
	}

	resp, err := c.Client.Do(req)

	if err != nil {
		c.logger.Debug("failed to do request: %s", err.Error())
		return nil, err
	}

	c.logger.Debug("cookies on request: %v", resp.Request.Cookies())
	c.logger.Debug("requested %s : status %d", req.URL.String(), resp.StatusCode)

	if c.config.debug {
		responseBytes, err := httputil.DumpResponse(resp, resp.ContentLength > 0)

		if err != nil {
			return nil, err
		}

		c.logger.Debug("raw response bytes received over wire: %d (%d kb)", len(responseBytes), len(responseBytes)/1024)
	}

	return resp, nil
}

func (c *HTTPClient) StatusCode(resp *http.Response) int {
	return resp.StatusCode
}

func (c *HTTPClient) Status(resp *http.Response) string {
	return resp.Status
}
