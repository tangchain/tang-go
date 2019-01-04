package resourceadapter

import (
	"context"

	"github.com/tang/go/xdr"
	. "github.com/tang/go/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
