package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/wuzuoliang/gitcli/vault"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/howeyc/gopass"
	"github.com/urfave/cli"
)

const (
	key = "git.example.com"
)

type credential struct {
	Token string
	Host  string
}

// Configure serves as a sub command
var Configure = func(c *cli.Context) error {
	cred, err := getCredential()
	if err != nil {
		fmt.Println("Configure::GetCredential error", err)
		return err
	}

	buff, err := json.Marshal(cred)
	if err != nil {
		fmt.Println("Configure::Marshal error", err)
		return err
	}

	client, err := vault.NewClient()
	if err != nil {
		fmt.Println("Configure::NewClient error", err)
		return err
	}

	err = client.Store(key, buff)
	if err != nil {
		fmt.Println("Configure::store error", err)
		return err
	}
	return nil
}

// RemoveConfig remove current config of app
var RemoveConfig = func(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	return client.Erase(key)
}

func getCredential() (*credential, error) {
	c := &credential{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Configure now, press Enter to use default.")
	fmt.Printf("Enter the base url of your project [default:http://git.example.com]: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("getCredential::ReadString error", err)
		return nil, err
	}

	host = strings.TrimSpace(host)
	if host == "" {
		host = "http://git.example.com"
	}
	c.Host = host

	fmt.Printf("Enter Private Token (visit %s/profile/personal_access_tokens if you do not have): ", c.Host)
	token, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println("getCredential::ReadString error", err)
		return nil, err
	}

	c.Token = strings.TrimSpace(string(token))
	if c.Token == "" {
		return nil, errors.New("invalid token")
	}
	return c, nil
}

func loadCredential() (*credential, error) {
	cred := credential{}
	client, err := vault.NewClient()
	if err != nil {
		fmt.Println("loadCredential::NewClient error", err)
		return nil, err
	}
	secret, err := client.Get(key)
	if err != nil {
		return nil, err
	}
	if len(secret) == 0 {
		fmt.Println("loadCredential::configuration not set")
		return nil, errors.New("configuration not set")
	}
	err = json.Unmarshal(secret, &cred)
	if err != nil {
		fmt.Println("loadCredential::Unmarshal error", err, "secret=", secret)
		return nil, err
	}
	return &cred, nil
}
