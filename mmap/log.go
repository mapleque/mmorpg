package mmap

import (
	"fmt"
	syslog "log"
)

type logInterface interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Error(msg ...interface{})
	Warn(msg ...interface{})
	Debugf(str string, message ...interface{})
	Errorf(str string, message ...interface{})
	Warnf(str string, message ...interface{})
	Infof(str string, message ...interface{})
}

type logger struct{}

var log logInterface

func init() {
	log = &logger{}
}

func SetLogger(logger logInterface) {
	log = logger
}

func (lg *logger) Debug(msg ...interface{}) {
	syslog.Println(msg...)
}

func (lg *logger) Info(msg ...interface{}) {
	syslog.Println(msg...)
}

func (lg *logger) Warn(msg ...interface{}) {
	syslog.Println(msg...)
}

func (lg *logger) Error(msg ...interface{}) {
	syslog.Println(msg...)
}

func (lg *logger) Debugf(str string, msg ...interface{}) {
	syslog.Println(fmt.Printf(str, msg...))
}

func (lg *logger) Infof(str string, msg ...interface{}) {
	syslog.Println(fmt.Printf(str, msg...))
}

func (lg *logger) Warnf(str string, msg ...interface{}) {
	syslog.Println(fmt.Printf(str, msg...))
}

func (lg *logger) Errorf(str string, msg ...interface{}) {
	syslog.Println(fmt.Printf(str, msg...))
}
