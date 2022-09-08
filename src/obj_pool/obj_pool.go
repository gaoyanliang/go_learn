package obj_pool

import (
	"errors"
	"time"
)

/* 可复用对象 */
type ReuseableObj struct {}

/* 对象池，用于缓存可复用对象 */
type ObjPool struct {
	bufChan chan *ReuseableObj
}

/* 创建对象池 */
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReuseableObj, numOfObj)

	// 向 buf channel 中放入可复用对象
	for i := 0; i < numOfObj; i ++ {
		objPool.bufChan <- &ReuseableObj{}
	}

	return &objPool
}

/* 获取对象 */
func (p *ObjPool) GetObj(timeout time.Duration) (*ReuseableObj, error)  {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <- time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

/* 释放对象 */
func (p *ObjPool) ReleaseObj(obj *ReuseableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")

	}
}