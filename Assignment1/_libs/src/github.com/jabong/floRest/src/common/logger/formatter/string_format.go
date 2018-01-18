package formatter

import (
	"fmt"
	"github.com/jabong/floRest/src/common/logger/message"
)

type stringFormat struct {
}

//formatString is format of the log in string formattype configuration
var formatString string = "[level : %s, message : %s, tId : %s, reqId : %s, appId : %s, sessionId : %s, userId : %s, stackTraces : %s, timestamp : %s, uri : %s]"

//GetFormattedLog returns formatted log
func (sf *stringFormat) GetFormattedLog(msg *message.LogMsg) interface{} {
	return fmt.Sprintf(formatString, msg.Level,
		msg.Message,
		msg.TransactionId,
		msg.RequestId,
		msg.AppId,
		msg.SessionId,
		msg.UserId,
		msg.StackTraces,
		msg.TimeStamp,
		msg.Uri)
}
