package common

import (
	"github.com/tang/go/support/log"
)

const TangAmountPrecision = 7

func CreateLogger(serviceName string) *log.Entry {
	return log.DefaultLogger.WithField("service", serviceName)
}
