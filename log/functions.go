/*
Package log implements a wrapper over zap logger package
use it like zap, but refer via log.

Package supports different logging levels,
some widely used configuration options,
and fast switching between sugar/desugar behavior.

For more details see zap, on which this package based: https://github.com/uber-go/zap
*/
package log

import (
	"go.uber.org/zap"
)

const (
	defaultLevel    = zap.DebugLevel
	defaultEncoding = "console"
	defaultMKey     = ""
	defaultLKey     = ""
	defaultLEncoder = "uppercase"
)

var defaultOutputPath = []string{"stdout"}
var defaultErrorPath = []string{"stderr"}

type base struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	conf   zap.Config
}

var log base

//ForExample makes logger from example config,
//as it descibed in zap package
func ForExample() {
	log.logger = zap.NewExample()
	log.sugar = log.logger.Sugar()
}

//ForDevelopment makes logger from development config,
//as it descibed in zap package
func ForDevelopment() {
	log.logger, _ = zap.NewDevelopment()
	log.sugar = log.logger.Sugar()
}

//ForProduction makes logger from production template,
//as it descibed in zap package
func ForProduction() {
	log.logger, _ = zap.NewProduction()
	log.sugar = log.logger.Sugar()
}

//FromScratch function for those, who wants to set all
//options manually
//
//For available options learn docs for zap package
func FromScratch(options map[string]interface{}) {
	if val, ok := options["level"]; ok {
		switch val.(string) {
		case "debug":
			log.conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info":
			log.conf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warning":
			log.conf.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error":
			log.conf.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		default:
			log.conf.Level = zap.NewAtomicLevelAt(defaultLevel)
		}
	} else {
		log.conf.Level = zap.NewAtomicLevelAt(defaultLevel)
	}
}

//Sync used as standart zap behavior:
//defer log.Sync()
//It ensures writing all messages before application finished.
func Sync() {
	log.logger.Sync()
}
