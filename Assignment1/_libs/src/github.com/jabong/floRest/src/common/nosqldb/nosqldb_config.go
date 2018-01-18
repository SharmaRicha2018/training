package nosqldb

import ()

type Config struct {
	Platform     string
	ConnStr      string
	Password     string
	BucketHashes []string
	Cluster      bool
}
