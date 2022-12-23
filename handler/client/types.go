package client

import (
	"io"
	"net/url"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/http2"
	tls "github.com/bogdanfinn/utls"
)

type Response struct {
	headers http.Header

	body []byte

	status     string
	statusCode int
	cookies    []*http.Cookie
}

var defaultRedirectFunc = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

type HttpClient interface {
	GetCookies(u *url.URL) []*http.Cookie
	SetCookies(u *url.URL, cookies []*http.Cookie)
	SetProxy(proxyUrl string) error
	GetProxy() string
	SetFollowRedirect(followRedirect bool)
	GetFollowRedirect() bool
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)

	// StatusCode(resp *http.Response) int
	// PostForm(url string, data url.Values) (resp *http.Response, err error)
	NewRequest() *Request
	// SetURL(req *http.Request, url string) error
	// SetHeader(req *http.Request, key, value string)
	// SetHeaders(req *http.Request, headers map[string]string)
	// SetBody(req *http.Request, body io.Reader)
	// SetBodyString(req *http.Request, body string)
	// SetBodyBytes(req *http.Request, body []byte)
	// SetBodyJSON(req *http.Request, body interface{})
	// SetBodyForm(req *http.Request, data url.Values)
}

type Request struct {
	client *Client

	method, url, host string

	header http.Header

	body io.Reader

	cookies []*http.Cookie
}

type Client struct {
	http.Client
	logger         Logger
	config         *httpClientConfig
	LatestResponse *Response
}

// NewRequest implements HttpClient
func (c *Client) NewRequest() *Request {
	return &Request{
		client: c,
		header: make(http.Header),
	}
}

type ClientProfile struct {
	clientHelloId     tls.ClientHelloID
	settings          map[http2.SettingID]uint32
	settingsOrder     []http2.SettingID
	pseudoHeaderOrder []string
	connectionFlow    uint32
	priorities        []http2.Priority
	headerPriority    *http2.PriorityParam
}

var Chrome_108 = ClientProfile{
	clientHelloId: tls.HelloChrome_108,
	settings: map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:      65536,
		http2.SettingEnablePush:           0,
		http2.SettingMaxConcurrentStreams: 1000,
		http2.SettingInitialWindowSize:    6291456,
		http2.SettingMaxHeaderListSize:    262144,
	},
	settingsOrder: []http2.SettingID{
		http2.SettingHeaderTableSize,
		http2.SettingEnablePush,
		http2.SettingMaxConcurrentStreams,
		http2.SettingInitialWindowSize,
		http2.SettingMaxHeaderListSize,
	},
	pseudoHeaderOrder: []string{
		":method",
		":authority",
		":scheme",
		":path",
	},
	connectionFlow: 15663105,
}

var DefaultClientProfile = Chrome_108

var DefaultTimeoutSeconds = 30

var DefaultOptions = []HttpClientOption{
	WithTimeout(DefaultTimeoutSeconds),
	WithClientProfile(DefaultClientProfile),
	WithRandomTLSExtensionOrder(),
	WithNotFollowRedirects(),
}
