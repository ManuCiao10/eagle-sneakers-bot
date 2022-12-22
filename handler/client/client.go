package client

import (
	"fmt"
	"net/url"
	"time"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/httputil"
	"golang.org/x/net/proxy"
)

func NewRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

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
