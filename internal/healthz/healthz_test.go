package healthz

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthz(t *testing.T) {
	assert := require.New(t)

	healthcheck := httptest.NewServer(http.HandlerFunc(Handler()))

	res, err := http.Get(healthcheck.URL)
	assert.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(err)
	assert.Equal("200 OK", res.Status)
	assert.Equal(`{"message":"ok"}`, string(body))
}
