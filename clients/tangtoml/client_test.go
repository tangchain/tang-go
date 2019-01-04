package tangtoml

import (
	"net/http"
	"strings"
	"testing"

	"github.com/tang/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://tang.org/.well-known/tang.toml", c.url("tang.org"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://tang.org/.well-known/tang.toml", c.url("tang.org"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://tang.org/.well-known/tang.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetTangToml("tang.org")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// tang.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/tang.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", TangTomlMaxSize)+`"`,
		)
	stoml, err = c.GetTangToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "tang.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing.org/.well-known/tang.toml").
		ReturnNotFound()
	stoml, err = c.GetTangToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json.org/.well-known/tang.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetTangToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
