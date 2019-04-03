package cmd

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	fmt.Println("BuildTree=============")
	root := BuildTree()

	level1woda := &MultiTree{11, 0, "woda", groupType, "http://git.woda.ink/groups/woda", []*MultiTree{}}

	level1business := &MultiTree{179, 0, "business", groupType, "http://git.woda.ink/groups/business", []*MultiTree{}}

	level1dw := &MultiTree{140, 0, "dw", groupType, "http://git.woda.ink/groups/dw", []*MultiTree{}}

	level1h5 := &MultiTree{29, 0, "h5", groupType, "http://git.woda.ink/groups/h5", []*MultiTree{}}

	level1jifanfei := &MultiTree{97, 0, "jifanfei", groupType, "http://git.woda.ink/groups/jifanfei", []*MultiTree{}}

	level1mini := &MultiTree{164, 0, "mini", groupType, "http://git.woda.ink/groups/mini", []*MultiTree{}}

	root.Children = append(root.Children, level1woda, level1business, level1dw, level1h5, level1jifanfei, level1mini)

	level2minidata := &MultiTree{164, 11, "woda/mini_data", groupType, "http://git.woda.ink/groups/woda/mini_data", []*MultiTree{}}

	level2product := &MultiTree{159, 11, "woda/product", groupType, "http://git.woda.ink/groups/woda/product", []*MultiTree{}}
	level1woda.Children = append(level1woda.Children, level2minidata, level2product)

	fmt.Println("DFS=============")
	root.DFS(true)

	fmt.Println("BFS=============")
	root.BFS(true)

	fmt.Println("FindByID=============")
	temp := root.FindByID(11)
	temp.DFS(true)

	fmt.Println("FindByName=============")
	temp = root.FindByName("woda/product")
	temp.DFS(true)

	fmt.Println("AddNode=============")
	root.AddNode(&MultiTree{17, 11, "woda/services", groupType, "http://git.woda.ink/groups/woda/services", []*MultiTree{}})
	temp = root.FindByID(11)
	temp.DFS(false)

	fmt.Println("DelNode=============")
	root.DelNode("mini")
	root.DFS(true)

	fmt.Println("DelNode=============")
	root.DelNode("woda/mini_data")
	root.DFS(true)
}
