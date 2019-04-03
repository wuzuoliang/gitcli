package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"go/build"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	// GetParams is parameters of get sub command
	GetParams = struct {
		OutputDirectory string
		ExcludeList     cli.StringSlice
	}{}
)

// GetHandler handle the get command
var GetHandler = func(c *cli.Context) error {
	cred, err := loadCredential()
	if err != nil {
		fmt.Println("GetHandler::loadCredential error", err)
		return err
	}

	u, err := url.Parse(cred.Host)
	if err != nil {
		fmt.Println("GetHandler::url.Parse error", err)
		return err
	}

	dics, err := LoadBoltDB("")
	if err != nil {
		fmt.Println("GetHandler::LoadBoltDB error", err)
		return err
	}

	tree := MakeOrderTree(dics)

	destTree := tree.FindByName(c.Args().First())
	if destTree == nil {
		fmt.Println("GetHandler::not Find path", c.Args().First())
		return errors.New("Not Find Path")
	}

	fmt.Println("GetParams.OutputDirectory", GetParams.OutputDirectory)
	// check output dir
	outDir := GetParams.OutputDirectory
	if outDir == "" {
		outDir = os.Getenv("GOPATH")
		if outDir == "" {
			outDir = build.Default.GOPATH
		}
		outDir = filepath.Join(outDir, "src/"+u.Host)
	}

	fmt.Println("Output Directory", outDir)
	fmt.Println("Exclude Directory", GetParams.ExcludeList)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		fmt.Println("Path", outDir, "not exist")
		return err
	}

	// exclude list
	for _, v := range GetParams.ExcludeList {
		tree.DelNode(v)
	}

	syncDir(outDir, destTree)
	return nil
}

func syncDir(outDir string, destTree *MultiTree) {
	if destTree == nil {
		return
	}
	paths := Path2StrSlice(destTree.Path)
	for _, v := range paths {
		tPath := filepath.Join(outDir, v)
		if v != destTree.Path {
			if _, err := os.Stat(tPath); os.IsNotExist(err) {
				err = os.Mkdir(tPath, 0750)
				if err != nil {
					fmt.Println("Mkdir", tPath, "error", err)
					return
				}
			}
		} else {
			if destTree.Type == projectType {
				tParPath := tPath[:strings.LastIndex(tPath, "/")]
				err := os.Chdir(tParPath)
				if err != nil {
					fmt.Println("Chdir", outDir+v, "error", err)
					return
				}
				repo := destTree.SSH
				cloneOrUpdate(repo, tPath)
			} else {
				for i := range destTree.Children {
					syncDir(outDir, destTree.Children[i])
				}
			}
		}
	}
}

func exeCmd(cmd string) error {
	fmt.Println("now executing: ", cmd)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(out) != 0 {
		fmt.Println(string(out))
	}
	return nil
}

func cloneOrUpdate(repo, dir string) {
	dst := filepath.Clean(dir + "/.git")
	parentDir := filepath.Clean(dir + "/../")
	// check to see if git repo exist
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err := os.Chdir(parentDir)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("now object:", dir)
		cmd := "git clone " + repo
		err = exeCmd(cmd)
		if err != nil {
			fmt.Println("[] exeCmd error", err)
		}
	} else {
		err := os.Chdir(dir)
		if err != nil {
			fmt.Println("chdir failed, err", err, dir)
			return
		}
		fmt.Println("now object:", dir)
		cmd := "git pull"
		err = exeCmd(cmd)
		if err != nil {
			fmt.Println("failed for ", repo)
		}
	}
}
