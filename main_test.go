package headers_test

import (
	"errors"
	headers "github.com/Napas/go-default-headers-round-tripper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func Test_defaultHeaders(t *testing.T) {
	t.Parallel()

	t.Run("Return response from the parent round tripper", func(t *testing.T) {
		expectedResp := &http.Response{}
		expectedErr := errors.New("error")
		req := &http.Request{}

		roundTripperMock := roundTripper{}
		roundTripperMock.
			On("RoundTrip", req).
			Once().
			Return(expectedResp, expectedErr)

		defaultHeaders := headers.NewDefaultHeaders(roundTripperMock, nil)
		resp, err := defaultHeaders.RoundTrip(req)

		assert.Equal(t, expectedResp, resp)
		assert.Equal(t, expectedErr, err)
		roundTripperMock.AssertExpectations(t)
	})

	t.Run("Sets headers", func(t *testing.T) {
		req := &http.Request{
			Header: http.Header{},
		}

		roundTripperMock := roundTripper{}
		roundTripperMock.
			On("RoundTrip", req).
			Return(nil, nil)

		defaultHeaders := headers.NewDefaultHeaders(roundTripperMock, map[string]string{
			"HeaderA": "value1",
			"HeaderB": "value2",
		})
		_, err := defaultHeaders.RoundTrip(req)

		assert.NoError(t, err)
		assert.Equal(t, "value1", req.Header.Get("HeaderA"))
		assert.Equal(t, "value2", req.Header.Get("HeaderB"))
	})
}

type roundTripper struct {
	mock.Mock
}

func (m roundTripper) RoundTrip(
	req *http.Request,
) (resp *http.Response, err error) {
	args := m.Called(req)

	if args.Get(0) != nil {
		resp = args.Get(0).(*http.Response)
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return resp, err
}
