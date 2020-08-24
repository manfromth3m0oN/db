package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

const (
	key       = "test"
	value     = "hello world"
	charset   = "abcdefghijklmnopqrstuvwxyz"
	chunkSize = 10
)

var dataBuf []byte
var errValueNotFound error = errors.New("Value not found in chunk")

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func write(key, value string) {
	kvp := []byte(fmt.Sprintf("%v:%v;", key, value))
	dataBuf = append(dataBuf, kvp...)
	err := ioutil.WriteFile("data.db", dataBuf, 0644)
	checkErr(err)
}

func populateRandomData() {
	write(key, value)
	for i := 0; i < 1008; i++ {
		var randomKey []byte
		for j := 0; j < 4; j++ {
			randomKey = append(randomKey, charset[rand.Intn(26)])
		}
		var randomValue []byte
		for k := 0; k < 4; k++ {
			randomValue = append(randomValue, charset[rand.Intn(26)])
		}
		write(string(randomKey), string(randomValue))
	}
}

func valueInChunk(pairs []string, key string, response chan string) {
	for _, pairString := range pairs {
		pair := strings.Split(pairString, ":")
		if pair[0] == key {
			response <- pair[1]
			//println(pair[1])
		}
	}
}

// Find a value by key
func findK(key string) (string, error) {
	data, err := ioutil.ReadFile("data.db")
	checkErr(err)
	pairs := strings.Split(string(data), ";")
	numPairs := len(pairs)
	response := make(chan string)
	if numPairs%chunkSize == 0 {
		for i := 0; i < numPairs/chunkSize; i++ {
			fmt.Printf("Current i: %v, current bottom: %v, current top: %v\n", i, i*10, (i+1)*10)
			go valueInChunk(pairs[i*10:(i+1)*10], key, response)
		}
		val := <-response
		fmt.Println(val)
		return val, nil
	}
	return "", errValueNotFound
}

func main() {
	//populateRandomData()
	if val, err := findK("not a key"); err != nil {
		fmt.Println(val)
	} else {
		fmt.Println("value is not in the store")
	}
}
