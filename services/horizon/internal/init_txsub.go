package horizon

import (
	"net/http"

	"github.com/tang/go/services/horizon/internal/db2/core"
	"github.com/tang/go/services/horizon/internal/db2/history"
	"github.com/tang/go/services/horizon/internal/txsub"
	results "github.com/tang/go/services/horizon/internal/txsub/results/db"
	"github.com/tang/go/services/horizon/internal/txsub/sequence"
)

func initSubmissionSystem(app *App) {
	cq := &core.Q{Session: app.CoreSession(nil)}

	app.submitter = &txsub.System{
		Pending:         txsub.NewDefaultSubmissionList(),
		Submitter:       txsub.NewDefaultSubmitter(http.DefaultClient, app.config.TangCoreURL),
		SubmissionQueue: sequence.NewManager(),
		Results: &results.DB{
			Core:    cq,
			History: &history.Q{Session: app.HorizonSession(nil)},
		},
		Sequences:         cq.SequenceProvider(),
		NetworkPassphrase: app.networkPassphrase,
	}
}

func init() {
	appInit.Add("txsub", initSubmissionSystem, "app-context", "log", "horizon-db", "core-db")
}
