
package vault

import (
	"github.com/keybase/go-keychain"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	t.Log(client.ClientType())
}
func TestStore(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	t.Log(client.ClientType())
	t.Log(client.Erase("git.woda.ink"))
	err = client.Store("git.woda.ink", []byte(`{
	"Token": "WzgYPXu2z6ynFs1XLWmH",
	"Host": "http://git.woda.ink"
}`))
	if err != nil {
		t.Error(err)
	}
	t.Log("Store succ")

}
func TestGet(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}

	secret, err := client.Get("wuzuoliang@woda.ink")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(secret))
}
func TestDel(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	_, err = client.Get("wuzuoliang@woda.ink")
	if err != nil {
		t.Error(err)
	}
	t.Log(client.Erase("wuzuoliang@woda.ink"))
}

func TestComSeq(t *testing.T) {
	// Create generic password item with service, account, label, password, access group
	item := keychain.NewGenericPassword("MyService", "gabriel", "A label", []byte("toomanysecrets"), "A123456789.group.com.mycorp")
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	err := keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		// Duplicate
		t.Log(keychain.ErrorDuplicateItem)
	}

	accounts, err := keychain.GetGenericPasswordAccounts("MyService")
	t.Log(accounts)
	// Should have 1 account == "gabriel"
	passwd, e := keychain.GetGenericPassword("MyService", "gabriel", "A label", "A123456789.group.com.mycorp")
	t.Log(string(passwd))
	t.Log(e)

	err = keychain.DeleteGenericPasswordItem("MyService", "gabriel")
	t.Log(err)

	accounts, err = keychain.GetGenericPasswordAccounts("MyService")
	t.Log(accounts)
}
