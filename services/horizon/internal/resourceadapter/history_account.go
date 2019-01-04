package resourceadapter

import (
	"context"

	"github.com/tang/go/services/horizon/internal/db2/history"
	. "github.com/tang/go/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
