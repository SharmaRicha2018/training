package examples

import (
	"fmt"
	"github.com/jabong/floRest/src/common/cache"
	workflow "github.com/jabong/floRest/src/common/orchestrator"
)

type HelloWorld struct {
	id string
}

func (n *HelloWorld) SetID(id string) {
	n.id = id
}

func (n HelloWorld) GetID() (id string, err error) {
	return n.id, nil
}

func (a HelloWorld) Name() string {
	return "HelloWorld"
}

func (a HelloWorld) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	// fill cache config
	conf := new(cache.Config)
	conf.Platform = "centralCache"
	conf.Host = "http://localhost:8080/cache/api/v1/buckets" // Blitz should be running in localhost:8080
	conf.KeyPrefix = "default"
	conf.ExpirySec = 60

	// get cache object
	cacheAdapter, err := cache.Get(*conf) // It should be called only once and can be shared across go routines

	// Put some items with TTL
	item1 := cache.Item{"somekey1", "somevalue1"}
	item2 := cache.Item{"somekey2", "somevalue2"}
	item3 := cache.Item{"somekey3", "somevalue3"}

	err = cacheAdapter.SetWithTimeout(item1, false, false, 60)
	if err != nil {
		fmt.Println("Error in setting keys in blitz. Error - " + err.Error())
		return io, nil
	}

	cacheAdapter.SetWithTimeout(item2, false, false, 60)
	cacheAdapter.SetWithTimeout(item3, false, false, 60)

	fmt.Println("Setting items are successful")

	// Get an item
	item, err := cacheAdapter.Get("somekey1", false, false)
	if err != nil {
		fmt.Println("Getting item from blitz failed. Error - " + err.Error())
		return io, nil
	}

	fmt.Println("Got the item - key : " + item.Key + ", value : " + item.Value.(string))

	// Get bulk items
	keys := []string{"somekey1", "somekey2", "somekey4"}

	items, err := cacheAdapter.GetBatch(keys, false, false)
	if err != nil {
		fmt.Println("Getting bulk items from blitz failed. Error - " + err.Error())
		return io, nil
	}

	fmt.Println("Got bulk items " + items["somekey1"].Value.(string) + ", " + items["somekey2"].Value.(string) + ", " + items["somekey4"].Error + ", " + items["somekey4"].Value.(string))

	// Delete an item
	err = cacheAdapter.Delete("somekey1")
	if err != nil {
		fmt.Println("Error in deleting item from blitz. Error - " + err.Error())
		return io, nil
	}
	item, err = cacheAdapter.Get("somekey1", false, false)
	if err == nil {
		fmt.Println("Item deleted.. But still present in cache. Value : " + item.Value.(string))
		return io, nil
	}
	fmt.Println("Item deleted successfully..")

	// Delete bulk items
	keysToDelete := []string{"somekey2", "somekey3"}
	err = cacheAdapter.DeleteBatch(keysToDelete)
	if err != nil {
		fmt.Println("Error in deleting bulk items from blitz. Error - " + err.Error())
		return io, nil
	}
	item, err = cacheAdapter.Get("somekey2", false, false)
	if err == nil {
		fmt.Println("Item deleted.. But still present in cache. Value : " + item.Value.(string))
		return io, nil
	}

	fmt.Println("Bulk items deleted successfully..")

	//Business Logic
	return io, nil
}
