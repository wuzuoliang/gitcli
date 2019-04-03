package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/urfave/cli"
)

// LoadHandler serves for load command handler
var LoadHandler = func(c *cli.Context) error {

	db, err := CreateBoltDB("")
	if err != nil {
		return err
	}
	defer db.Close()

	// load cred client
	cred, err := loadCredential()
	if err != nil {
		fmt.Println("LoadHandler::loadCredential err", err)
		return err
	}
	client := newClient(cred.Host, cred.Token)

	projects, err := client.GetProjects()
	if err != nil {
		fmt.Println("LoadHandler::GetProjects error", err)
		return err
	}
	groups, err := client.GetGroups()
	if err != nil {
		fmt.Println("LoadHandler::GetGroups error", err)
		return err
	}

	infoList := make([]GitLabInfo, 0, len(projects)+len(groups))
	for _, v := range projects {
		infoList = append(infoList, GitLabInfo{v.ID, v.PID, v.Name, v.Path, v.SSH, projectType})
	}
	for _, v := range groups {
		infoList = append(infoList, GitLabInfo{v.ID, v.PID, v.Name, v.Path, "", groupType})
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(buckets))
		if err != nil {
			fmt.Print("LoadHandler::tx.CreateBucketIfNotExists error", err)
			return err
		}
		for _, v := range infoList {
			err = b.Put([]byte(v.Path), v.Byte())
			if err != nil {
				fmt.Println("LoadHandler::b.Put error", err)
			}
		}
		return nil
	})
	return err
}
