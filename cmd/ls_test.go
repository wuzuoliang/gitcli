package cmd

import (
	"fmt"
	"testing"
)

// Before run `*_test.go` please use load cmd to init `bolt.db` file.
func Test_LS(t *testing.T) {

	dictionary, err := LoadBoltDB("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink/gitcli/")
	if err != nil {
		t.Fatal(err)
	}
	tree := MakeOrderTree(dictionary)

	root := tree.FindByName("")
	if root != nil {
		root.LS(true)
	} else {
		fmt.Println("not found")
	}

	woda := tree.FindByName("woda")
	if woda != nil {
		fmt.Println("LS=======================")
		woda.LS(false)
	} else {
		fmt.Println("not found")
	}

	tree.DelNode("business/proto")
	tree.FindByName("business").DFS(false)
}
