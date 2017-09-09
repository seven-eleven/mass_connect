package common

import (
	"errors"
	"sync"
)

var kernelMap = make(map[string]string) // key-value存储内核

var rwlock = new(sync.RWMutex) // 内核操作读写锁

// 使用key-value更新内核数据
func KernelUpdate(key string, value string) {
	if "" == key {
		Log("key is null")
		return
	}

	rwlock.Lock()
	kernelMap[key] = value
	//Log("kernel write: ", key, ",", value)
	rwlock.Unlock()
}

// 使用KEY查询数据
func KernelQueryByKey(key string) (string, error) {
	if "" == key {
		Log("key is null")
		return "", errors.New("key is null")
	}

	rwlock.RLock()
	value, ok := kernelMap[key]
	//Log("kernel read by ", key)
	rwlock.RUnlock()

	if ok {
		return value, nil
	} else {
		return "", errors.New("key not found")
	}
}

// 查询内核所有数据
func KernelQueryAll() map[string]string {
	return kernelMap
}
