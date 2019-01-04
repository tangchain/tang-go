package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/tang/go/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("tang_address", govalidator.CustomTypeValidator(isTangAddress))
}

func isTangAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}
