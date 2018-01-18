package nosqldb

import (
	"github.com/jabong/floRest/src/common/logger"
)

type NoSqlDbError struct {
	ErrCode          string
	DeveloperMessage string
}

const (
	ERR_NO_PLATFORM          = "Platform not found"
	ERR_INITIALIZATION       = "Initialization failed"
	ERR_GET_FAILURE          = "Failure in Get() method"
	ERR_SET_FAILURE          = "Failure in Set() method"
	ERR_GET_BATCH_FAILURE    = "Failure in GetBatch() method"
	ERR_DELETE_FAILURE       = "Failure in Delete() method"
	ERR_DELETE_BATCH_FAILURE = "Failure in DeleteBatch() method"
)

// getErrObj returns error object with given details
func getErrObj(errCode string, developerMessage string) (ret *NoSqlDbError) {
	logger.Error("ErrorCode: " + errCode + ", developerMessage : " + developerMessage)
	ret = new(NoSqlDbError)
	ret.ErrCode = errCode
	ret.DeveloperMessage = developerMessage
	return ret
}
