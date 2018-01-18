package profiler

import (
	"github.com/jabong/floRest/src/common/monitor"
)

var rate float64 // rate to be used for monitoring

// InitProfiler inits the profiler
func InitProfiler(r float64) {
	rate = r
}

//StartProfile starts a profiling using profiler instance p for key. Key should have the
//following name '<package-file-method>'. For example a method 'Initialise'
//in the file 'search_dao.go' file inside package accessor should have a profile
//key as 'accessor-search_dao-Initialise'
func StartProfile(p *Profiler, key string) {
	if p == nil {
		return
	}
	p.StartProfile(key)
}

//EndProfile ends the profiling starting using profiler instance p for key k
func EndProfile(p *Profiler, k string, a ...interface{}) int64 {
	if p == nil {
		return 0
	}
	t := p.EndProfile(k)
	if t != NilTime {
		monitor.GetInstance().Histogram(k, float64(t), nil, rate)
	}
	return t
}

//EndProfileCustomMetric ends the profiling starting using profiler instance p for key k
//And sends the mettic with the name n
func EndProfileCustomMetric(p *Profiler, k string, n string, a ...interface{}) int64 {
	if p == nil {
		return 0
	}
	t := p.EndProfile(k)
	if t != NilTime {
		monitor.GetInstance().Histogram(n, float64(t), nil, rate)
	}
	return t
}

//EndProfileCustomRate ends the profiling starting using profiler instance p for key k
func EndProfileCustomRate(p *Profiler, k string, r float64, a ...interface{}) int64 {
	if p == nil {
		return 0
	}
	t := p.EndProfile(k)
	if t != NilTime {
		monitor.GetInstance().Histogram(k, float64(t), nil, r)
	}
	return t
}

//EndProfileCustomMetricCustomRate ends the profiling starting using profiler instance p for key k
//And sends the mettic with the name n
func EndProfileCustomMetricCustomRate(p *Profiler, k string, n string, r float64, a ...interface{}) int64 {
	if p == nil {
		return 0
	}
	t := p.EndProfile(k)
	if t != NilTime {
		monitor.GetInstance().Histogram(n, float64(t), nil, r)
	}
	return t
}
