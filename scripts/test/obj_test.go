package obj

import (
	"testing"
)

// go test ./... -run TestUser_Success
func TestUser_Success(t *testing.T) {
	t.Logf("TestUser: %v", t.Name())

	args := "taro"
	want := "taro"

	u := NewUser("1")
	u.SetName(args)

	got := u.GetName()

	if got != want {
		t.Errorf("GetName() == %s, want %s", got, want)
	}

	t.Run("Sub", func(t *testing.T) {
		t.Log("Sub test")
	})
}

func TestUser_Fail(t *testing.T) {
	t.Fail()
}

