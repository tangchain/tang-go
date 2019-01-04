package helpers

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/tang/go/address"
	"github.com/tang/go/amount"
	"github.com/tang/go/strkey"
)

func init() {
	govalidator.CustomTypeTagMap.Set("tang_accountid", govalidator.CustomTypeValidator(isTangAccountID))
	govalidator.CustomTypeTagMap.Set("tang_seed", govalidator.CustomTypeValidator(isTangSeed))
	govalidator.CustomTypeTagMap.Set("tang_asset_code", govalidator.CustomTypeValidator(isTangAssetCode))
	govalidator.CustomTypeTagMap.Set("tang_address", govalidator.CustomTypeValidator(isTangAddress))
	govalidator.CustomTypeTagMap.Set("tang_amount", govalidator.CustomTypeValidator(isTangAmount))
	govalidator.CustomTypeTagMap.Set("tang_destination", govalidator.CustomTypeValidator(isTangDestination))

}

func Validate(request Request, params ...interface{}) error {
	valid, err := govalidator.ValidateStruct(request)

	if !valid {
		fields := govalidator.ErrorsByField(err)
		for field, errorValue := range fields {
			switch {
			case errorValue == "non zero value required":
				return NewMissingParameter(field)
			case strings.HasSuffix(errorValue, "does not validate as tang_accountid"):
				return NewInvalidParameterError(field, "Account ID must start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as tang_seed"):
				return NewInvalidParameterError(field, "Account secret must start with `S` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as tang_asset_code"):
				return NewInvalidParameterError(field, "Asset code must be 1-12 alphanumeric characters.")
			case strings.HasSuffix(errorValue, "does not validate as tang_address"):
				return NewInvalidParameterError(field, "Tang address must be of form user*domain.com")
			case strings.HasSuffix(errorValue, "does not validate as tang_destination"):
				return NewInvalidParameterError(field, "Tang destination must be of form user*domain.com or start with `G` and contain 56 alphanum characters.")
			case strings.HasSuffix(errorValue, "does not validate as tang_amount"):
				return NewInvalidParameterError(field, "Amount must be positive and have up to 7 decimal places.")
			default:
				return NewInvalidParameterError(field, errorValue)
			}
		}
	}

	return request.Validate(params...)
}

// These are copied from support/config. Should we move them to /strkey maybe?
func isTangAccountID(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteAccountID, enc)

	if err == nil {
		return true
	}

	return false
}

func isTangSeed(i interface{}, context interface{}) bool {
	enc, ok := i.(string)

	if !ok {
		return false
	}

	_, err := strkey.Decode(strkey.VersionByteSeed, enc)

	if err == nil {
		return true
	}

	return false
}

func isTangAssetCode(i interface{}, context interface{}) bool {
	code, ok := i.(string)

	if !ok {
		return false
	}

	if !govalidator.IsByteLength(code, 1, 12) {
		return false
	}

	if !govalidator.IsAlphanumeric(code) {
		return false
	}

	return true
}

func isTangAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)
	if err != nil {
		return false
	}

	return true
}

func isTangAmount(i interface{}, context interface{}) bool {
	am, ok := i.(string)

	if !ok {
		return false
	}

	_, err := amount.Parse(am)
	if err != nil {
		return false
	}

	return true
}

// isTangDestination checks if `i` is either account public key or Tang address.
func isTangDestination(i interface{}, context interface{}) bool {
	dest, ok := i.(string)

	if !ok {
		return false
	}

	_, err1 := strkey.Decode(strkey.VersionByteAccountID, dest)
	_, _, err2 := address.Split(dest)

	if err1 != nil && err2 != nil {
		return false
	}

	return true
}
