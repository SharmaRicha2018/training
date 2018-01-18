package exectime

import (
	"github.com/jabong/floRest/src/common/monitor"
	"github.com/jabong/floRest/src/common/orchestrator"
	"time"
)

func Start(io orchestrator.WorkFlowData) {
	io.IOData.Set("apiStartTime", time.Now())
}

func End(io orchestrator.WorkFlowData, key string, tags []string) {
	dataDogAgent := monitor.GetInstance()
	startTimeIntf, err := io.IOData.Get("apiStartTime")
	if err == nil {
		startTime, _ := startTimeIntf.(time.Time)
		execTime := time.Since(startTime) / time.Millisecond
		dataDogAgent.Histogram(key, float64(execTime), tags, float64(execTime))
	}
}
