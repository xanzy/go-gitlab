package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBlockUser(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/block", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.BlockUser(1)
	if err != nil {
		t.Errorf("Users.BlockUser returned error: %v", err)
	}
}

func TestBlockUser_UserNotFound(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/block", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.BlockUser(1)
	if err != UserNotFoundError {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\n     Got: %+v", UserNotFoundError, err)
	}
}

func TestBlockUser_BlockPrevented(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/block", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.BlockUser(1)
	if err != UserBlockPreventedError {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\n     Got: %+v", UserBlockPreventedError, err)
	}
}

func TestBlockUser_UnknownError(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/block", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Errorf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.BlockUser(1)
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\n     Got: %+v", want, err)
	}
}

func TestBlockUser_BadResponseFromNet(t *testing.T) {
	client := NewClient(nil, "")
	client.SetBaseURL("")

	want := "Post /users/1/block: unsupported protocol scheme \"\""

	err := client.Users.BlockUser(1)
	if err.Error() != want {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\n     Got: %+v", want, err)
	}
}

//  ------------------------  Unblock user -------------------------
func TestUnblockUser(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/unblock", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.UnblockUser(1)
	if err != nil {
		t.Errorf("Users.UnblockUser returned error: %v", err)
	}
}

func TestUnblockUser_UserNotFound(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/unblock", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.UnblockUser(1)
	if err != UserNotFoundError {
		t.Errorf("Users.UnblockUser error.\nExpected: %+v\n     Got: %+v", UserNotFoundError, err)
	}
}

func TestUnblockUser_UnblockPrevented(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/unblock", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.UnblockUser(1)
	if err != UserUnblockPreventedError {
		t.Errorf("Users.UnblockUser error.\nExpected: %+v\n     Got: %+v", UserUnblockPreventedError, err)
	}
}

func TestUnblockUser_UnknownError(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/users/1/unblock", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Errorf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.UnblockUser(1)
	if !reflect.DeepEqual(err, want) {
		t.Errorf("Users.UnblockUser error.\nExpected: %+v\n     Got: %+v", want, err)
	}
}

func TestUnblockUser_BadResponseFromNet(t *testing.T) {
	client := NewClient(nil, "")
	client.SetBaseURL("")

	want := "Post /users/1/unblock: unsupported protocol scheme \"\""

	err := client.Users.UnblockUser(1)
	if err.Error() != want {
		t.Errorf("Users.UnblockUser error.\nExpected: %+v\n     Got: %+v", want, err)
	}
}
