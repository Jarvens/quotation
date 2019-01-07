// auth: kunlun
// date: 2019-01-07
package utils

import log "github.com/alecthomas/log4go"

// initial log4go
func init() {
	log.LoadConfiguration("log4go.xml")
}

// debug
func Debug(arg0 interface{}, args ...interface{}) {
	log.Debug(arg0, args)
}

// info
func Info(arg0 interface{}, args ...interface{}) {
	log.Info(arg0, args)
}

// war
func Warn(arg0 interface{}, args ...interface{}) {
	log.Warn(arg0, args)
}
