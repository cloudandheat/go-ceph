// Package log allows to enable go-ceph logging and integrate it with the
// logging of the go-ceph consumer.
package log

import (
	"sync/atomic"
	"unsafe"

	intLog "github.com/ceph/go-ceph/internal/log"
)

// Log levels
const (
	NoneLvl  = Level(intLog.NoneLvl)
	ErrorLvl = Level(intLog.ErrorLvl)
	WarnLvl  = Level(intLog.WarnLvl)
	InfoLvl  = Level(intLog.InfoLvl)
	DebugLvl = Level(intLog.DebugLvl)
)

// Level is the type for log levels.
type Level int32

// SetLevel sets the log level.
func SetLevel(lvl Level) {
	atomic.StoreInt32(&intLog.Level, int32(lvl))
}

// Outputer must be implemented by the log output destination.
type Outputer interface {
	Output(calldepth int, s string) error
}

// SetOutput sets the output destination for the logging.
func SetOutput(o Outputer) {
	var p unsafe.Pointer
	if o != nil {
		outFunc := o.Output
		p = unsafe.Pointer(&outFunc)
	}
	atomic.StorePointer(&intLog.OutputPtr, p)
}

// assert that type internal.Output matches Outputer interface
var _ = func() { _ = intLog.Output(Outputer(nil).Output) }
