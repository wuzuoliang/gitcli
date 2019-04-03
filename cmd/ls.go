package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

// LsHandler serves for tree command handler
var LsHandler = func(c *cli.Context) error {
	dictionary, err := LoadBoltDB("")
	if err != nil {
		return err
	}
	tree := MakeOrderTree(dictionary)

	temp := tree.FindByName(c.Args().First())
	if temp != nil {
		temp.LS(false)
	} else {
		fmt.Println("LsHandler::FindByName", c.Args().First(), "not found")
	}
	return nil
}
