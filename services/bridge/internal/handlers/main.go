package handlers

import (
	"github.com/tang/go/clients/federation"
	"github.com/tang/go/clients/horizon"
	"github.com/tang/go/clients/tangtoml"
	"github.com/tang/go/services/bridge/internal/config"
	"github.com/tang/go/services/bridge/internal/db"
	"github.com/tang/go/services/bridge/internal/listener"
	"github.com/tang/go/services/bridge/internal/submitter"
	"github.com/tang/go/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              horizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	TangTomlResolver  tangtoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
