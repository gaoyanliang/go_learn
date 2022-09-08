/* 使用 buffer channel 实现对象池 */
package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T)  {
	numOfObj := 10
	pool := NewObjPool(numOfObj)

	/* 尝试放置超出池大小的对象 */
	//if err := pool.ReleaseObj(&ReuseableObj{}); err != nil {
	//	t.Error(err)
	//}

	for i := 0; i < numOfObj; i ++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("get obj %T \n", v)

			/* 放回去 */
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}

	fmt.Println("Done")
}

