package cmd

import (
	"fmt"
	"strings"
	"testing"
)

// Before run `*_test.go` please use load cmd to init `bolt.db` file.
func Test_Tree(t *testing.T) {
	dictionary, err := LoadBoltDB("/Users/wuzuoliang/Documents/workspace/go/src/git.woda.ink/gitcli/")
	if err != nil {
		t.Fatal(err)
	}
	tree := MakeOrderTree(dictionary)

	root := tree.FindByName("")
	if root != nil {
		root.DFS(true)
	} else {
		fmt.Println("not found")
	}

	woda := tree.FindByName("woda")
	if woda != nil {
		fmt.Println("DFS=======================")
		woda.DFS(false)
		fmt.Println("BFS=======================")
		woda.BFS(false)
	} else {
		fmt.Println("not found")
	}
}
func Test_Spilt(t *testing.T) {
	fmt.Println(strings.Split("s", "/"))
	fmt.Println(strings.Split("s/", "/"))
	fmt.Println(strings.Split("s/b", "/"))
	fmt.Println(strings.Split("s/b/a", "/"))
	fmt.Println(strings.Split("s/b/a/", "/"))
	fmt.Println(Path2StrSlice(""))
	fmt.Println(Path2StrSlice("a"))
	fmt.Println(Path2StrSlice("a/"))
	fmt.Println(Path2StrSlice("a/b"))
	fmt.Println(Path2StrSlice("a/b/"))
	fmt.Println(Path2StrSlice("a/b/c"))
	fmt.Println(Path2StrSlice("a/b/c/"))

}
