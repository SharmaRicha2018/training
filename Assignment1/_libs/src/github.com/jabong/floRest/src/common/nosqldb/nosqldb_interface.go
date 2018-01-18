package nosqldb

//Item repesents the structure of an item to be stored in some nosqldb data store
type Item struct {
	Key   string
	Value interface{}
	Error string
}

// nosqldb interface
type NoSqlDbInterface interface {
	// Init initialises a connection to a nosql db server where all keys are to be namespaced by keyPrefix
	//	and the dump of all key-val pair in the cache to be stored in dumpFilePath (in case Dump is called)
	Init(conf *Config) *NoSqlDbError

	// Get gets an item from a nosqldb store indexed with key. serialize and compress indicates if the nosqldb implementation
	// has to undergo some serialization or compression before returning the item
	Get(key string, serialize bool, compress bool) (item *Item, err *NoSqlDbError)

	// Set sets an item into a nosqldb store. serialise and compress indicates if the nosqldb implementation
	// has to undergo some serialization or compression before setting the item in nosqldb
	Set(item Item, serialize bool, compress bool) *NoSqlDbError

	// SetWithTimeout sets the item into nosqldb, same as Set but this function takes an extra add_argument
	// which sets the timeout for the particular item, does not take expirySec from config
	SetWithTimeout(item Item, serialize bool, compress bool, ttl int32) *NoSqlDbError

	// Delete deletes a key from nosqldb
	Delete(key string) *NoSqlDbError

	// DeleteBatch deletes an array of keys from nosqldb
	DeleteBatch(keys []string) *NoSqlDbError

	// GetBatch gets a list of all items indexed with keys. serialize and compress indicates if the
	// cache implementation has to undergo some serialization or compression before returning the items
	GetBatch(keys []string, serialize bool, compress bool) (items map[string]*Item, err *NoSqlDbError)
}
