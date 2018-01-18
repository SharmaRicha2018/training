package message

import ()

/*
LogMsg struct is used for bundling message/requestcontext attributes
for dumping into log
*/
type LogMsg struct {
	Level         string   `json:"level"`
	Message       string   `json:"message"`
	TransactionId string   `json:"tId,omitempty"`
	RequestId     string   `json:"reqId,omitempty"`
	AppId         string   `json:"appId,omitempty"`
	SessionId     string   `json:"sessionId,omitempty"`
	UserId        string   `json:"userId,omitempty"`
	StackTraces   []string `json:"stackTraces,omitempty"`
	TimeStamp     string   `json:"timestamp"`
	Uri           string   `json:"uri,omitempty"`
}
