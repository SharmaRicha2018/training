package nosqldb

import (
	"github.com/jabong/floRest/src/common/logger"
	"github.com/jabong/floRest/src/common/utils/misc"
	"gopkg.in/redis.v3"
	"strings"
	"time"
)

type redisAdapter struct {
	client *redis.ClusterClient
	hashes []string
}

type mgetResult struct {
	Keys           []string
	SliceCmdOutput *redis.SliceCmd
}

type mdelResult struct {
	Keys         []string
	IntCmdOutput *redis.IntCmd
}

func (ra *redisAdapter) getHashKey(key string) string {
	hash := ra.getHash(key)
	return ra.getHashKeyFromHash(key, hash)
}

func (ra *redisAdapter) getHashKeyFromHash(key string, hash string) string {
	return "{" + hash + "}" + key
}

func (ra *redisAdapter) getHash(key string) string {
	hash := misc.GetHash(key, len(ra.hashes))
	return ra.hashes[hash]
}

func (ra *redisAdapter) Init(conf *Config) *NoSqlDbError {
	ra.client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(conf.ConnStr, ","),
		Password: conf.Password,
	})
	ra.hashes = conf.BucketHashes
	return nil
}

func (ra *redisAdapter) Get(key string, serialize bool, compress bool) (item *Item, err *NoSqlDbError) {
	hashKey := ra.getHashKey(key)
	val, getErr := ra.client.Get(hashKey).Result()
	if getErr != nil {
		return nil, getErrObj(ERR_GET_FAILURE, "Getting key failed with error : "+getErr.Error())
	}
	item = new(Item)
	item.Key = key
	item.Value = val
	return item, nil
}

func (ra *redisAdapter) Set(item Item, serialize bool, compress bool) *NoSqlDbError {
	hashKey := ra.getHashKey(item.Key)
	err := ra.client.Set(hashKey, item.Value, 0).Err()
	if err != nil {
		return getErrObj(ERR_SET_FAILURE, "Setting key failed with error : "+err.Error())
	}
	return nil
}

func (ra *redisAdapter) SetWithTimeout(item Item, serialize bool, compress bool, ttl int32) *NoSqlDbError {
	hashKey := ra.getHashKey(item.Key)
	err := ra.client.Set(hashKey, item.Value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return getErrObj(ERR_SET_FAILURE, "Setting key with ttl failed with error : "+err.Error())
	}
	return nil
}

func (ra *redisAdapter) Delete(key string) *NoSqlDbError {
	hashKey := ra.getHashKey(key)
	val, err := ra.client.Del(hashKey).Result()

	if err != nil {
		return getErrObj(ERR_DELETE_FAILURE, "Deleting key failed with error : "+err.Error())
	}

	if val != 1 {
		logger.Info("Delete failed because the key does not exist on the server")
	}

	return nil
}

func (ra *redisAdapter) DeleteBatch(keys []string) *NoSqlDbError {
	hashCountTemp := 0
	hashKeysMap := make(map[string][]string)
	keysMap := make(map[string][]string)
	valuesChannel := make(chan *mdelResult)
	defer close(valuesChannel)

	for _, key := range keys {
		hash := ra.getHash(key)
		if hashKeysMap[hash] == nil {
			hashCountTemp++
			hashKeysMap[hash] = make([]string, 0)
			keysMap[hash] = make([]string, 0)
		}
		hashKeysMap[hash] = append(hashKeysMap[hash], ra.getHashKeyFromHash(key, hash))
		keysMap[hash] = append(keysMap[hash], key)
	}

	for hash, hashKeys := range hashKeysMap {
		go func(keys []string, hashKeys []string) {
			result := new(mdelResult)
			result.Keys = keys
			result.IntCmdOutput = ra.client.Del(hashKeys...)
			valuesChannel <- result
		}(keysMap[hash], hashKeys)
	}

	for i := 0; i < hashCountTemp; i++ {
		result := <-valuesChannel
		val, err := result.IntCmdOutput.Result()

		if err != nil {
			return getErrObj(ERR_DELETE_BATCH_FAILURE, "Deleting bulk keys failed with error : "+err.Error())
		}

		if val != int64(len(result.Keys)) {
			logger.Info("Delete failed because the keys does not exists in the server")
		}
	}

	return nil
}

func (ra *redisAdapter) GetBatch(keys []string, serialize bool, compress bool) (items map[string]*Item, err *NoSqlDbError) {
	resMap := make(map[string]*Item, len(keys))
	hashCountTemp := 0
	hashKeysMap := make(map[string][]string)
	keysMap := make(map[string][]string)
	valuesChannel := make(chan *mgetResult)
	defer close(valuesChannel)

	for _, key := range keys {
		hash := ra.getHash(key)
		if hashKeysMap[hash] == nil {
			hashCountTemp++
			hashKeysMap[hash] = make([]string, 0)
			keysMap[hash] = make([]string, 0)
		}
		hashKeysMap[hash] = append(hashKeysMap[hash], ra.getHashKeyFromHash(key, hash))
		keysMap[hash] = append(keysMap[hash], key)
	}

	for hash, hashKeys := range hashKeysMap {
		go func(keys []string, hashKeys []string) {
			result := new(mgetResult)
			result.Keys = keys
			result.SliceCmdOutput = ra.client.MGet(hashKeys...)
			valuesChannel <- result
		}(keysMap[hash], hashKeys)
	}

	for i := 0; i < hashCountTemp; i++ {
		result := <-valuesChannel
		vals, err := result.SliceCmdOutput.Result()
		if err != nil {
			return nil, getErrObj(ERR_GET_BATCH_FAILURE, "Getting bulk keys failed with error : "+err.Error())
		}

		for index, val := range vals {
			item := new(Item)
			item.Key = result.Keys[index]
			if val != nil {
				item.Value = val
			} else {
				item.Error = "The key does not exist on the server"
			}
			resMap[item.Key] = item
		}
	}
	return resMap, nil
}
