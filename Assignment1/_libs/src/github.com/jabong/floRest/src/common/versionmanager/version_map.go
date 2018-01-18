package versionmanager

import (
	"github.com/jabong/floRest/src/common/ratelimiter"
)

/*
Data structure to store the API that is versioned
*/
type Version struct {
	Resource string
	Version  string
	Action   string
	BucketId string
}

type FullVersion struct {
	versionable Versionable
	rateLimiter *ratelimiter.RateLimiter
}

type VersionMap map[Version]FullVersion

func NewFullVersion(ver Versionable, rateLimit *ratelimiter.RateLimiter) FullVersion {
	p := FullVersion{
		versionable: ver,
		rateLimiter: rateLimit,
	}

	return p
}
