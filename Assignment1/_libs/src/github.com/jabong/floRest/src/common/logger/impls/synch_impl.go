package impls

import (
	"github.com/jabong/floRest/src/common/logger/message"
	"github.com/jabong/floRest/src/common/logger/writers"
)

type synchLogger struct {
	writer writers.LogWriter
}

func GetSynchLogger(writer writers.LogWriter) *synchLogger {
	ret := new(synchLogger)
	ret.writer = writer
	return ret
}

func (logr *synchLogger) Debug(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

func (logr *synchLogger) Info(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

func (logr *synchLogger) Trace(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

func (logr *synchLogger) Warning(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

func (logr *synchLogger) Error(msg message.LogMsg) {
	logr.writer.Write(&msg)
}

func (logr *synchLogger) Profile(msg message.LogMsg) {
	logr.writer.Write(&msg)
}
