package pclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWithAuthToken(t *testing.T) {
	token := "token"

	validate := func(t *testing.T, c *http.Client, token string) {
		t.Helper()
		want := token
		gotReq := false
		headerVal := ""
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotReq = true
			headerVal = r.Header.Get("Authorization")
		}))
		_, err := c.Get(srv.URL)
		assertNilError(t, err)
		if !gotReq {
			t.Error("request not sent")
		}
		if headerVal != want {
			t.Errorf("Authorization header is %v, want %v", headerVal, want)
		}
	}

	t.Run("zero-value Client", func(t *testing.T) {
		c := new(Client)
		WithAuthToken(token)(c)
		validate(t, c.httpClient, token)
	})

	t.Run("NewClient", func(t *testing.T) {
		// httpClient := &http.Client{}
		client := NewClient(WithAuthToken(token))
		validate(t, client.httpClient, token)
		// // make sure the original client isn't setting auth headers now
		// validate(t, httpClient, "")
	})

	t.Run("NewTokenClient", func(t *testing.T) {
		validate(t, NewTokenClient(token).httpClient, token)
	})
}
