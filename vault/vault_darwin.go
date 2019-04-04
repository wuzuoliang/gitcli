package vault

import (
	"github.com/keybase/go-keychain"
	"log"
)

type Vault interface {
	Store(key string, secret []byte) error
	Get(key string) ([]byte, error)
	Erase(key string) error
	ClientType() string
}

func NewClient() (Vault, error) {
	return &macCredClient{}, nil
}

type macCredClient struct {
}

func (w *macCredClient) Store(key string, secret []byte) error {
	// Create generic password item with service, account, label, password, access group
	item := keychain.NewGenericPassword("gitlab", key, "label", secret, "accessGroup")
	item.SetSynchronizable(keychain.SynchronizableDefault)
	item.SetAccessible(keychain.AccessibleAlways)
	err := keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		// Duplicate
		log.Println(keychain.ErrorDuplicateItem)
		err = keychain.DeleteGenericPasswordItem("gitlab", key)
		if err != nil {
			log.Println("DeleteGenericPasswordItem error", err)
		}
	}
	return nil
}

func (w *macCredClient) Get(key string) ([]byte, error) {
	return keychain.GetGenericPassword("gitlab", key, "label", "accessGroup")
}

func (w *macCredClient) Erase(key string) error {
	return keychain.DeleteGenericPasswordItem("gitlab", key)
}
func (w *macCredClient) ClientType() string {
	return "Mac"
}
