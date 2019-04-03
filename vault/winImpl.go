package vault

import (
	"github.com/danieljoos/wincred"
)

type winCredClient struct {
}

func (w *winCredClient) Store(key string, secret []byte) error {
	cred := wincred.NewGenericCredential(key)
	cred.CredentialBlob = secret
	return cred.Write()
}

func (w *winCredClient) Get(key string) ([]byte, error) {
	cred, err := wincred.GetGenericCredential(key)
	if err != nil {
		return nil, err
	}
	return cred.CredentialBlob, nil
}

func (w *winCredClient) Erase(key string) error {
	cred, err := wincred.GetGenericCredential(key)
	if err != nil {
		return err
	}
	return cred.Delete()
}

func (w *winCredClient) ClientType() string {
	return "Windows"
}
