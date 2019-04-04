package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateBoltDB(dbPath string) (*bolt.DB, error) {
	if len(dbPath) != 0 {
		if runtime.GOOS=="windows"{
			if !strings.HasSuffix(dbPath, "\\") {
				dbPath += "\\"
			}
		}else{
			if !strings.HasSuffix(dbPath, "/") {
				dbPath += "/"
			}
		}

	}

	if b, _ := pathExists(dbPath + boltDB); b {
		err := os.Remove(dbPath + boltDB)
		if err != nil {
			fmt.Println("CreateBoltDB::remove bolt.db error", err)
			return nil, err
		}
	}

	// Open the boltDB data file in your current directory.
	// It will be created if it doesn't exist.
	// Options 1 second timeout is purpose to avoid race.
	var err error
	db, err := bolt.Open(dbPath+boltDB, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("CreateBoltDB::bolt.Open error", err)
		return nil, err
	}
	return db, nil
}

func LoadBoltDB(dbPath string) ([]GitLabInfo, error) {
	if len(dbPath) != 0 {
		if runtime.GOOS=="windows"{
			if !strings.HasSuffix(dbPath, "\\") {
				dbPath += "\\"
			}
		}else{
			if !strings.HasSuffix(dbPath, "/") {
				dbPath += "/"
			}
		}
	}
	dictionary := make([]GitLabInfo, 0)

	// Open the boltDB data file in your current directory.
	// It will be created if it doesn't exist.
	// Options 1 second timeout is purpose to avoid race.
	var err error
	db, err := bolt.Open(dbPath+boltDB, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("LoadBoltDB::bolt.Open error", err)
		return dictionary, err
	}
	defer db.Close()

	var dict GitLabInfo
	err = db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(buckets))
		err = b.ForEach(func(k, v []byte) error {
			err = json.Unmarshal(v, &dict)
			if err != nil {
				fmt.Println("LoadBoltDB::json.Unmarshal error", err)
				return err
			}
			dictionary = append(dictionary, dict)
			return nil
		})
		if err != nil {
			fmt.Println("LoadBoltDB::b.ForEach error", err)
		}
		return err
	})
	if err != nil {
		fmt.Println("LoadBoltDB::db.View error", err)
	}

	sort.Sort(GitLabInfoList(dictionary))
	return dictionary, nil
}
