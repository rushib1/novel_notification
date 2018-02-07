package main

import (
	"github.com/boltdb/bolt"
	"fmt"
	"encoding/json"
)

func (novel *Novel) save(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(novel)
		if err != nil {
			return err
		}
		return b.Put([]byte(novel.NovelName), []byte(encoded))
	})
	return err
}

func Novelcheck(db *bolt.DB, key, value []byte) {
	fmt.Printf("gelta %s",key)
}

func mental() {
	fmt.Println("hello roka")
}

func iter(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		c := b.Cursor()

		for k,v := c.First(); k!=nil; k, v = c.Next() {
			go Novelcheck(db, k, v)
		}
		return nil
	})
}

type Novel struct {
	NovelName string
	Url string
	Chapter int
}

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err!=nil {
		fmt.Println("hello")
	}
	iter(db)
	defer db.Close()
	var input string
	fmt.Scanln(&input)
}