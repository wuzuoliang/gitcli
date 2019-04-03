package vault

import (
	"github.com/pkg/errors"
	"runtime"
)

type Vault interface {
	Store(key string, secret []byte) error
	Get(key string) ([]byte, error)
	Erase(key string) error
	ClientType() string
}

func NewClient() (Vault, error) {
	if runtime.GOOS == "windows" {
		return &winCredClient{}, nil
	} else if runtime.GOOS == "darwin" {
		return &macCredClient{}, nil
	}
	return nil, errors.New("unsupported os")
}
