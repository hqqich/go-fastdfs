package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	levelDB()
}
func levelDB() {
	db, _ := leveldb.OpenFile("./fileserver.db", nil)

	db.Put([]byte("hello"), []byte("world"), nil)

	data, _ := db.Get([]byte("hello"), nil)
	fmt.Println(string(data))

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Print(string(iter.Key()))
		fmt.Print(":")
		fmt.Print(string(iter.Value()))
		fmt.Print("\n")
		fmt.Println("===============================================================================")
	}

}
