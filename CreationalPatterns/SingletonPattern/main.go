package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type database struct {
}

var dbInstance *database

func getInstance() *database {
	//lock operation is expensive, so its better to check if instance is nii before applying lock
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			fmt.Println("Db instance created Now...")
			dbInstance = &database{}
		} else {
			fmt.Println("DB instance already created...")
		}
	} else {
		fmt.Println("DB instance already created...")
	}

	return dbInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
