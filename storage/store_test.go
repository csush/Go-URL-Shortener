package storage

import "testing"

func TestStoreSave(t *testing.T) {
	t.Run("save URL mapping", func(t *testing.T) {
		store := &Store{urls: make(map[string]string)}
		store.Save("http://localhost:5001/abcd", "http://csu.sh")

		got, err := store.Get("http://localhost:5001/abcd")
		want := "http://csu.sh"

		if err != nil {
			t.Error("shouldn't get an error but got one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
