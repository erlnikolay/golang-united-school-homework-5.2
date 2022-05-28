package cache

//package main

import (
	"time"
)

type cSubCache struct {
	cElement       string
	cDeadLineExist bool
	cDeadLine      time.Time
}

type Cache struct {
	cBuffer map[string]cSubCache
}

func NewCache() Cache {
	var c Cache

	c = Cache{cBuffer: make(map[string]cSubCache)}
	// Cache{
	//		cBuffer {
	//			"2022-05-28 11:00:00": {cElement: "value1",cDeadLineExist: false, cDeadLine: 2022-05-28 12:00:00},
	//			"2022-05-28 11:56:00": {cElement: "valueNew2",cDeadLineExist: false, cDeadLine: 2022-05-28 14:00:00},
	//  		...
	//		}
	//}

	return c
}

func (c Cache) Get(key string) (string, bool) {
	var kValue cSubCache
	var ok bool

	// get cSubCache data
	kValue, ok = c.cBuffer[key]
	if ok {
		// check dead line
		if kValue.cDeadLineExist {
			// if dead time hasn't achived
			if kValue.cDeadLine.After(time.Now()) {
				return kValue.cElement, true
			} else {
				// dead time is achived
				return "", false
			}
		} else {
			// dead line is not set
			return kValue.cElement, true
		}
	} else {
		// element is not presented
		return "", false
	}
}

func (c Cache) Put(key, value string) {
	var newSubCache cSubCache

	newSubCache = cSubCache{
		cElement:       value,
		cDeadLineExist: false,
		cDeadLine:      time.Now(),
	}
	c.cBuffer[key] = newSubCache
}

func (c Cache) Keys() []string {
	var resKeys []string

	for keyOfElement, valueOfElement := range c.cBuffer {
		// there is not deadline
		if !valueOfElement.cDeadLineExist {
			resKeys = append(resKeys, keyOfElement)
		} else {
			// there is deadline
			if valueOfElement.cDeadLine.After(time.Now()) {
				// deadline has been not expired yet
				resKeys = append(resKeys, keyOfElement)
			}
		}
	}
	return resKeys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	var newSubCache cSubCache

	newSubCache = cSubCache{
		cElement:       value,
		cDeadLineExist: true,
		cDeadLine:      deadline,
	}

	c.cBuffer[key] = newSubCache
}

//func main() {
//	var nBuffer Cache
//	var time1 string
//	var time2 string
//	var time3 string
//	var time4 string
//	var resVal string
//	var ok bool
//
//	nBuffer = NewCache()
//
//	time1 = time.Now().Format(time.UnixDate)
//	time.Sleep(1 * time.Second)
//	time2 = time.Now().Format(time.UnixDate)
//	time.Sleep(1 * time.Second)
//	time3 = time.Now().Format(time.UnixDate)
//	time.Sleep(1 * time.Second)
//	fmt.Printf("time1: %s\n", time1)
//	fmt.Printf("time2: %s\n", time2)
//	fmt.Printf("time3: %s\n", time3)
//
//	// save into cache three new value
//	nBuffer.Put(time1, "one")
//	// gets one by one
//	if resVal, ok = nBuffer.Get(time1); ok {
//		fmt.Printf("Cache key: %s, cache value: %s\n", time1, resVal)
//	} else {
//		fmt.Printf("Cache key: %s is not found\n", time1)
//	}
//	nBuffer.Put(time2, "two")
//	if resVal, ok = nBuffer.Get(time2); ok {
//		fmt.Printf("Cache key: %s, cache value: %s\n", time2, resVal)
//	} else {
//		fmt.Printf("Cache key: %s is not found\n", time2)
//	}
//	nBuffer.Put(time3, "three")
//	if resVal, ok = nBuffer.Get(time3); ok {
//		fmt.Printf("Cache key: %s, cache value: %s\n", time3, resVal)
//	} else {
//		fmt.Printf("Cache key: %s is not found\n", time3)
//	}
//	fmt.Printf("All keys: \n")
//	for _, oneKey := range nBuffer.Keys() {
//		fmt.Println(oneKey)
//		//fmt.Printf(nBuffer.Get(oneKey))
//	}
//	fmt.Printf("Add key with deadline\n")
//	tmpTime := time.Now()
//	time4 = tmpTime.Format(time.UnixDate)
//	nBuffer.PutTill(time4, "four", tmpTime.Add(5*time.Second))
//	fmt.Printf("All keys: \n")
//	for _, oneKey := range nBuffer.Keys() {
//		fmt.Println(oneKey)
//	}
//	fmt.Printf("value of dead line Element:\n")
//	if resVal, ok = nBuffer.Get(time4); ok {
//		fmt.Printf("Cache key: %s, cache value: %s\n", time4, resVal)
//	} else {
//		fmt.Printf("Cache key: %s is not found\n", time4)
//	}
//	time.Sleep(10 * time.Second)
//	fmt.Printf("value of dead line Element:\n")
//	if resVal, ok = nBuffer.Get(time4); ok {
//		fmt.Printf("Cache key: %s, cache value: %s\n", time4, resVal)
//	} else {
//		fmt.Printf("Cache key: %s is not found\n", time4)
//	}
//}
