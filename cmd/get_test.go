package cmd

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_Get(t *testing.T) {

	dictionary, err := LoadBoltDB("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink/gitcli/")
	if err != nil {
		t.Fatal(err)
	}
	tree := MakeOrderTree(dictionary)

	root := tree.FindByName("business")
	if root != nil {
		root.DFS(true)
	} else {
		fmt.Println("not found")
	}

	syncDir("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink", root)
}
func Test_FilePathJoin(t *testing.T) {
	t.Log(filepath.Join("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink", "mini"))
	t.Log(filepath.Join("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink/", "mini"))
	t.Log(filepath.Join("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink/", "mini/"))
	t.Log(filepath.Join("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink/", "mini/woda"))
}
