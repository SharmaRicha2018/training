package formatter

import (
	"github.com/jabong/floRest/src/common/logger/message"
)

// FormatInterface interface methods for formatterss
type FormatInterface interface {
	GetFormattedLog(msg *message.LogMsg) interface{}
}
