package horizon

import (
	"net/http"

	hProblem "github.com/tang/go/services/horizon/internal/render/problem"
	"github.com/tang/go/support/render/problem"
)

// RateLimitExceededAction renders a 429 response
type RateLimitExceededAction struct {
	Action
	App *App
}

func (action RateLimitExceededAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ap := &action.Action
	ap.Prepare(w, r)
	ap.App = action.App
	problem.Render(action.R.Context(), action.W, hProblem.RateLimitExceeded)
}
