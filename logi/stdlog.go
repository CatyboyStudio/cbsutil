package logi

import (
	"fmt"
	"log"
)

type stdlog int

// Error implements Logi.
func (stdlog) Error(err error, msg string) {
	log.Println(msg, err)
}

// Errorf implements Lolog
func (stdlog) Errorf(err error, format string, ps ...any) {
	msg := fmt.Sprintf(format, ps...)
	log.Println(msg, err)
}

// Print implements Logi.
func (stdlog) Print(msg string) {
	log.Println(msg)
}

// Printf implements Logi.
func (stdlog) Printf(format string, ps ...any) {
	msg := fmt.Sprintf(format, ps...)
	log.Println(msg)
}

var Stdlog Logi = stdlog(0)
