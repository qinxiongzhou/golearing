package obj_pool

import (
	"errors"
	"time"
)

type ResuableObj struct {
}

type ObjPool struct {
	bufChan chan *ResuableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ResuableObj, numOfObj)

	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ResuableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ResuableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret,nil
	case <-time.After(timeout):
		return nil,errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ResuableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")

	}
}

