package main

import (
	"fmt"
	"sync"
)

type MyMap struct{
	m map[interface{}]interface{}
	mu sync.RWMutex
}

func NewMyMap() *MyMap {
	return &MyMap{
		m:make(map[interface{}]interface{}),
	}
}

func (mMap *MyMap) Delete(key interface{}) {
	mMap.mu.Lock()
	defer mMap.mu.Unlock()
	delete(mMap.m,key)
}

func (mMap *MyMap) Load(key interface{}) (value interface{},ok bool) {
	mMap.mu.RLock()
	defer mMap.mu.RUnlock()
	value,ok =mMap.m[key]
	return
}

func (mMap *MyMap) LoadOrStore(key interface{},value interface{}) (actual interface{},loaded bool) {
	mMap.mu.Lock()
	defer mMap.mu.Unlock()
	actual,loaded=mMap.m[key]
	if loaded {
		return
	}
	mMap.m[key]=value
	actual=value
	return
}

func (mMap *MyMap) Range(f func(key,value interface{}) bool) {
	mMap.mu.RLock()
	defer mMap.mu.RUnlock()
	for k,v := range mMap.m {
		if !f(k,v) {
			break
		}
	}
}

func (mMap *MyMap) Store(key interface{},value interface{}) {
	mMap.mu.Lock()
	defer mMap.mu.Unlock()
	mMap.m[key]=value
}

func main(){
	m:=NewMyMap()
	m.Store("15916648982","zhenzhen")
	m.Store("15566648982","tangyuan")
	m.Store("15916643482","xinhua")
	f:= func(key interface{},value interface{}) bool {
		fmt.Printf("key is %v,value is %v.\n",key,value)
		return true
	}

	v1,ok:=m.Load("15916648982")
	act1,ok:=m.LoadOrStore("15916648982",v1)
	if ok {
		fmt.Printf("act1 is %v.it exist.",act1)
	}
	m.Delete("15916648982")
	act1,ok=m.LoadOrStore("15916648982",v1)
	if ok {
		fmt.Printf("act1 is %v.it exist.",act1)
	}
	fmt.Println(v1)
	if ok {
		fmt.Println("v1 delete unsuccessfully.")
	}
	m.Range(f)
}