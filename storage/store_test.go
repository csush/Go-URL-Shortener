package storage

import "testing"

func TestStore(t *testing.T) {
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
	t.Run("get URL mapping", func(t *testing.T) {
		urls := map[string]string{
			"http://localhost:5001/a123": "https://google.com",
			"http://localhost:5001/321a": "http://csu.sh",
		}
		store := &Store{urls: urls}

		got, err := store.Get("http://localhost:5001/321a")
		want := "https://google.com"

		if err != nil {
			t.Error("shouldn't get an error but got one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("URL mapping doesn't exist", func(t *testing.T) {
		urls := make(map[string]string)
		urls["http://localhost:5001/a123"] = "http://csu.sh"
		urls["http://localhost:5001/321a"] = "https://google.com"
		store := &Store{urls: urls}

		_, err := store.Get("http://localhost:5001/doesntexist")

		if err == nil {
			t.Error("should get an error but didnt get one")
		}
		if err.Error() != ErrNotFoundMsg {
			t.Errorf("should get error %q but got %q", ErrNotFoundMsg, err.Error())
		}
	})
}
