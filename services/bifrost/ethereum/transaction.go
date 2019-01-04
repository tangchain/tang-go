package ethereum

import (
	"math/big"

	"github.com/tang/go/services/bifrost/common"
)

func (t Transaction) ValueToTang() string {
	valueEth := new(big.Rat)
	valueEth.Quo(new(big.Rat).SetInt(t.ValueWei), weiInEth)
	return valueEth.FloatString(common.TangAmountPrecision)
}
