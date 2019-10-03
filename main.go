package headers

import "net/http"

type defaultHeaders struct {
	parent  http.RoundTripper
	headers map[string]string
}

func NewDefaultHeaders(parent http.RoundTripper, headers map[string]string) http.RoundTripper {
	return &defaultHeaders{parent, headers}
}

func (d defaultHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range d.headers {
		req.Header.Set(key, value)
	}

	return d.parent.RoundTrip(req)
}
