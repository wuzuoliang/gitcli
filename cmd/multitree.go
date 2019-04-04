package cmd

import (
	"fmt"
	"runtime"
	"strings"
)

type MultiTree struct {
	ID       int
	PID      int
	Path     string
	Type     int
	SSH      string
	Children []*MultiTree
}

func BuildTree() *MultiTree {
	if runtime.GOOS == "windows" {
		return &MultiTree{0, 0, "\\", groupType, "", make([]*MultiTree, 0)}
	}
	return &MultiTree{0, 0, "/", groupType, "", make([]*MultiTree, 0)}
}

func (m MultiTree) Print(detailPrint bool) {
	if detailPrint {
		fmt.Println("ID=", m.ID, "PID=", m.PID, "Path=", m.Path, "Type=", m.Type, "Children Num=", len(m.Children), "SSH", m.SSH)
	} else {
		if m.Type == groupType && m.ID != 0 {
			if runtime.GOOS == "windows" {
				fmt.Println(m.Path + "\\")
			} else {
				fmt.Println(m.Path + "/")
			}
		} else {
			fmt.Println(m.Path)
		}
	}
}

func (m MultiTree) DFS(detailPrint bool) {
	m.Print(detailPrint)
	for i := range m.Children {
		if m.Children[i] != nil {
			m.Children[i].DFS(detailPrint)
		}
	}
}

func (m MultiTree) BFS(detailPrint bool) {
	queue := make([]MultiTree, 0)
	queue = append(queue, m)
	for len(queue) > 0 {
		temp := queue[0]
		temp.Print(detailPrint)
		for i := range temp.Children {
			if temp.Children[i] != nil {
				queue = append(queue, *temp.Children[i])
			}
		}
		queue = queue[1:]
	}
}
func (m MultiTree) LS(detailPrint bool) {
	for _, v := range m.Children {
		v.Print(detailPrint)
	}
}
func (m *MultiTree) FindByName(path string) *MultiTree {
	if len(path) == 0 {
		if runtime.GOOS == "windows" {
			path = "\\"
		} else {
			path = "/"
		}
	}
	if m.Path == path {
		return m
	}
	for i := range m.Children {
		if m.Children[i] != nil {
			ret := m.Children[i].FindByName(path)
			if ret != nil {
				return ret
			}
		}
	}
	return nil
}

func (m *MultiTree) FindByID(ID int) *MultiTree {
	if m.ID == ID {
		return m
	}
	for i := range m.Children {
		if m.Children[i] != nil {
			ret := m.Children[i].FindByID(ID)
			if ret != nil {
				return ret
			}
		}
	}
	return nil
}

func (m *MultiTree) AddNode(treeNode *MultiTree) {
	t := m.FindByID(treeNode.PID)
	if t == nil {
		if runtime.GOOS == "windows" {
			fmt.Println("MultiTree::AddNode not found parent path", treeNode.Path[:strings.LastIndex(treeNode.Path, "\\")])
		} else {
			fmt.Println("MultiTree::AddNode not found parent path", treeNode.Path[:strings.LastIndex(treeNode.Path, "/")])
		}
		return
	}
	found := false
	for _, v := range t.Children {
		if v.ID == treeNode.ID {
			found = true
		}
	}
	if !found {
		t.Children = append(t.Children, treeNode)
	}
}

func (m *MultiTree) DelNode(name string) {
	if name == "/" || name == "\\" || len(name) == 0 {
		return
	}
	t := m.FindByName(name)
	if t != nil {
		pt := m.FindByID(t.PID)
		if pt != nil {
			for i := range pt.Children {
				if pt.Children[i] != nil {
					if pt.Children[i].Path == name {
						pt.Children[i] = nil
					}
				}
			}
		}
	}
}
func MakeOrderTree(dictionary []GitLabInfo) *MultiTree {
	nameIDRel := make(map[string]int, len(dictionary))
	tree := BuildTree()
	index := 0
	if runtime.GOOS == "windows" {
		nameIDRel["\\"] = index
	} else {
		nameIDRel["/"] = index
	}
	for _, v := range dictionary {
		pathSlice := Path2StrSlice(v.Path)
		pathSliceLen := len(pathSlice)
		for i := 0; i < pathSliceLen; i++ {
			if _, ok := nameIDRel[pathSlice[i]]; !ok {
				index++
				nameIDRel[pathSlice[i]] = index
				tPid := 0
				if i > 0 {
					if runtime.GOOS == "windows" {
						prePath := strings.TrimSuffix(pathSlice[i-1], "\\")
						if pid, ok := nameIDRel[prePath]; ok {
							tPid = pid
						}
					} else {
						prePath := strings.TrimSuffix(pathSlice[i-1], "/")
						if pid, ok := nameIDRel[prePath]; ok {
							tPid = pid
						}
					}
				} else {
					tree.AddNode(&MultiTree{index, 0, pathSlice[0], groupType, "", make([]*MultiTree, 0)})
				}
				if pathSlice[i] != v.Path {
					tree.AddNode(&MultiTree{index, tPid, pathSlice[i], groupType, "", make([]*MultiTree, 0)})
				} else {
					tree.AddNode(&MultiTree{index, tPid, pathSlice[i], v.Type, v.SSH, make([]*MultiTree, 0)})
				}
			}
		}
	}
	return tree
}
