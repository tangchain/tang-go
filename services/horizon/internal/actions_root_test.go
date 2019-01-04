package horizon

import (
	"encoding/json"
	"testing"

	"github.com/tang/go/services/horizon/internal/test"
	"github.com/tang/go/protocols/horizon"
)

func TestRootAction(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	server := test.NewStaticMockServer(`{
			"info": {
				"network": "test",
				"build": "test-core",
				"protocol_version": 4
			}
		}`)
	defer server.Close()

	ht.App.horizonVersion = "test-horizon"
	ht.App.config.TangCoreURL = server.URL
	ht.App.UpdateTangCoreInfo()

	w := ht.Get("/")
	if ht.Assert.Equal(200, w.Code) {
		var actual horizon.Root
		err := json.Unmarshal(w.Body.Bytes(), &actual)
		ht.Require.NoError(err)
		ht.Assert.Equal("test-horizon", actual.HorizonVersion)
		ht.Assert.Equal("test-core", actual.TangCoreVersion)
	}
}
