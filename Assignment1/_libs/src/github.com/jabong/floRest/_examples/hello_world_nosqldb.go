package examples

import (
	"fmt"
	"github.com/jabong/floRest/src/common/nosqldb"
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
	// fill nosqldb config
	conf := new(nosqldb.Config)
	conf.Platform = "redis"
	conf.BucketHashes = []string{"dog", "cat"}
	conf.Cluster = true
	conf.Password = ""
	conf.ConnStr = ":30001,:30002,:30003,:30004,:30005,:30006"

	// get redis object
	nosqlAdapter, err := nosqldb.Get(conf) // It should be called only once and can be shared across go routines

	// Put some items with TTL
	item1 := nosqldb.Item{"somekey1", "somevalue1", ""}
	item2 := nosqldb.Item{"somekey2", "somevalue2", ""}
	item3 := nosqldb.Item{"somekey3", "somevalue3", ""}

	err = nosqlAdapter.SetWithTimeout(item1, false, false, 1000)
	if err != nil {
		fmt.Println("Error in setting keys in nosqldb. Error - " + err.DeveloperMessage)
		return io, nil
	}

	nosqlAdapter.SetWithTimeout(item2, false, false, 1000)
	nosqlAdapter.SetWithTimeout(item3, false, false, 1000)

	fmt.Println("Setting items are successful")

	// Get an item
	item, err := nosqlAdapter.Get("somekey1", false, false)
	if err != nil {
		fmt.Println("Getting item from nosqldb failed. Error - " + err.DeveloperMessage)
		return io, nil
	}

	fmt.Println("Got the item - key : " + item.Key + ", value : " + item.Value.(string))

	// Get bulk items
	keys := []string{"somekey1", "somekey2", "somekey4"}

	items, err := nosqlAdapter.GetBatch(keys, false, false)
	if err != nil {
		fmt.Println("Getting bulk items from nosqldb failed. Error - " + err.DeveloperMessage)
		return io, nil
	}
	fmt.Println("Got bulk items " + items["somekey1"].Value.(string) + ", " + items["somekey2"].Value.(string) + ", " + items["somekey4"].Error)

	// Delete an item
	err = nosqlAdapter.Delete("somekey1")
	if err != nil {
		fmt.Println("Error in deleting item from nosqldb. Error - " + err.DeveloperMessage)
		return io, nil
	}
	item, err = nosqlAdapter.Get("somekey1", false, false)
	if err == nil {
		fmt.Println("Item deleted.. But still present in cache. Value : " + item.Value.(string))
		return io, nil
	}
	fmt.Println("Item deleted successfully..")

	// Delete bulk items
	keysToDelete := []string{"somekey2", "somekey3"}
	err = nosqlAdapter.DeleteBatch(keysToDelete)
	if err != nil {
		fmt.Println("Error in deleting bulk items from nosqldb. Error - " + err.DeveloperMessage)
		return io, nil
	}
	item, err = nosqlAdapter.Get("somekey2", false, false)
	if err == nil {
		fmt.Println("Item deleted.. But still present in nosqldb. Value : " + item.Value.(string))
		return io, nil
	}

	fmt.Println("Bulk items deleted successfully..")

	//Business Logic
	return io, nil
}
