package obj_pool_test

import (
	"ch32/obj_pool"
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := obj_pool.NewObjPool(10)

	for i := 0; i < 12; i++ {
		if v,err := pool.GetObj(time.Second * 1);err != nil{
			t.Error(err)
		}else {
			fmt.Printf("%T\n",v)
			if err := pool.ReleaseObj(v);err != nil{
				t.Error(err)
			}
		}
	}

}
