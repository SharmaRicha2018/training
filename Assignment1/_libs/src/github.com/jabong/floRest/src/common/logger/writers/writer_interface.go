package writers

import (
	"github.com/jabong/floRest/src/common/logger/formatter"
	"github.com/jabong/floRest/src/common/logger/message"
)

type LogWriter interface {
	Write(msg *message.LogMsg)
	SetFormatter(formatter.FormatInterface)
}
