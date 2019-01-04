package test

import (
	"github.com/tang/go/services/horizon/internal/test/scenarios"
)

func loadScenario(scenarioName string, includeHorizon bool) {
	tangCorePath := scenarioName + "-core.sql"
	horizonPath := scenarioName + "-horizon.sql"

	if !includeHorizon {
		horizonPath = "blank-horizon.sql"
	}

	scenarios.Load(TangCoreDatabaseURL(), tangCorePath)
	scenarios.Load(DatabaseURL(), horizonPath)
}
