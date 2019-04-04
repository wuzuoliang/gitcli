package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"runtime"
	"strings"
)

// TreeHandler serves for tree command handler
var TreeHandler = func(c *cli.Context) error {
	dictionary, err := LoadBoltDB("")
	if err != nil {
		return err
	}
	tree := MakeOrderTree(dictionary)
	temp := tree.FindByName(c.Args().First())
	if temp != nil {
		temp.DFS(false)
	} else {
		fmt.Println("TreeHandler::FindByName", c.Args().First(), "not found")
	}
	return nil
}

func Path2StrSlice(path string) []string {
	if runtime.GOOS == "windows" {
		path = strings.TrimSuffix(path, "\\")
		if len(path) > 0 {
			tmp := strings.Split(path, "\\")
			length := len(tmp)
			ret := make([]string, 0, length)
			for cur := 0; cur < length; cur++ {
				t := ""
				for i := 0; i <= cur; i++ {
					if i == 0 {
						t += tmp[i]
					} else {
						t += "\\" + tmp[i]
					}
				}
				ret = append(ret, t)
			}
			return ret
		}
		return []string{"\\"}
	} else {
		path = strings.TrimSuffix(path, "/")
		if len(path) > 0 {
			tmp := strings.Split(path, "/")
			length := len(tmp)
			ret := make([]string, 0, length)
			for cur := 0; cur < length; cur++ {
				t := ""
				for i := 0; i <= cur; i++ {
					if i == 0 {
						t += tmp[i]
					} else {
						t += "/" + tmp[i]
					}
				}
				ret = append(ret, t)
			}
			return ret
		}
		return []string{"/"}
	}
}
