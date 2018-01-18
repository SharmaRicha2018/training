package monitor

import (
	"errors"
)

type myMonitor struct {
	agent     MonitorInterface
	isEnabled bool
}

var monitorObj *myMonitor // singleton object

// method to ensure singleton behaviour
func GetInstance() *myMonitor {
	if monitorObj == nil {
		monitorObj = new(myMonitor)
	}
	// return
	return monitorObj
}

// initialize the object
func (obj *myMonitor) Initialize(cnfg *MonitorConf) (err error) {
	// set the config
	obj.isEnabled = cnfg.Enabled
	// set the Agent now
	obj.agent, err = Get(cnfg)
	if err != nil {
		obj.agent = nil
	} else {
		// start sending server stats, e.g mem usage, disk usage, etc.
		obj.agent.SendAppMetrics(cnfg.MetricsServer)
	}
	return err
}

// histogram
func (obj *myMonitor) Histogram(name string, val float64, tags []string, rate float64) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Histogram(name, val, tags, rate)
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// error
func (obj *myMonitor) Error(prefix string, msg string, tags map[string]string) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Error(getMonitorDataObj(prefix, msg, tags))
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// warning
func (obj *myMonitor) Warning(prefix string, msg string, tags map[string]string) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Warning(getMonitorDataObj(prefix, msg, tags))
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// info
func (obj *myMonitor) Info(prefix string, msg string, tags map[string]string) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Info(getMonitorDataObj(prefix, msg, tags))
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// success
func (obj *myMonitor) Success(prefix string, msg string, tags map[string]string) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Success(getMonitorDataObj(prefix, msg, tags))
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// count
func (obj *myMonitor) Count(name string, value int64, tags []string, rate float64) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Count(name, value, tags, rate)
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// guage
func (obj *myMonitor) Guage(name string, value float64, tags []string, rate float64) (err error) {
	if !monitorObj.isEnabled {
		return nil
	}
	if obj.AgentInitialized() {
		err = obj.agent.Gauge(name, value, tags, rate)
	} else {
		err = errors.New("Monitor agent nil")
	}
	return err
}

// is agent initialized
func (obj *myMonitor) AgentInitialized() bool {

	agent := false
	if obj.agent != nil {
		agent = true
	}
	// return bool
	return agent
}

// getMonitorDataObj: get monitor data object
func getMonitorDataObj(prefix string, msg string, tags map[string]string) (ret *MonitorData) {
	ret = new(MonitorData)
	ret.Body = msg
	ret.Tags = tags
	ret.Title = prefix
	ret.ToSendEvent = true
	// return
	return ret
}
