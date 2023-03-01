package godemo

import (
	"fmt"
	"sync"
)

func testSyncMap() {
	mp := sync.Map{}
	mp.Store("test", "10")
	v, ok := mp.Load("test")
	fmt.Println(ok, v.(string))

	mp.Range(func(key, value any) bool {
		k := key.(string)
		v := value.(string)
		fmt.Println(k, v)
		return true
	})
}
