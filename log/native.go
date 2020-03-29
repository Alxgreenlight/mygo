package log

import (
	"go.uber.org/zap"
)

//L provide native zap logger for getting
//functionality that nopt provided by this package
func L() *zap.Logger {
	return log.logger
}

//S provide native zap sugaredLogger for getting
//functionality that not provided by this package
func S() *zap.SugaredLogger {
	return log.sugar
}
