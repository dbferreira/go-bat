package main

import (
	"fmt"
	"time"
)

import as "github.com/aerospike/aerospike-client-go"

func main() {

	client, err := as.NewClient("127.0.0.1", 3000)

	// key, err := as.NewKey("namespace", "set",
	// "key value goes here and can be any supported primitive")

	policy := as.NewWritePolicy(0, 0)
	policy.Timeout = 50 * time.Millisecond

	key2, _ := as.NewKey("test", "myset", "mykey")
	bin := as.NewBin("mybin", "myvalue")

	fmt.Println("bin is ", bin)

	// client.PutBins(policy, key2, bin)

	// bin1 := as.NewBin("bin1", "value1")
	// bin2 := as.NewBin("bin2", "value2")

	// Write a record
	// err = client.PutBins(nil, key, bin1, bin2)

	// Read a record
	record, err := client.Get(nil, key2)

	fmt.Println("record and error: ", record, err)
	client.Close()

}
