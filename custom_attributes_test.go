package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListCustomAttributes(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"key":"testkey1", "value":"testvalue1"}, {"key":"testkey2", "value":"testvalue2"}]`)
	})

	customAttributes, _, err := client.CustomAttributes.ListCustomAttributes(CustomAttributeOptions{
		CustomeAttributeResourceID: "2",
		CustomeAttributeResource:   "groups",
	})
	if err != nil {
		t.Errorf("CustomAttributes.ListCustomAttributes returned error: %v", err)
	}

	want := []*CustomAttributes{{Key: "testkey1", Value: "testvalue1"}, {Key: "testkey2", Value: "testvalue2"}}
	if !reflect.DeepEqual(want, customAttributes) {
		t.Errorf("CustomAttributes.ListCustomAttributes returned %+v, want %+v", customAttributes, want)
	}
}

func TestGetCustomAttribute(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttributes.GetCustomAttribute(CustomAttributeOptions{
		CustomeAttributeResourceID: "2",
		CustomeAttributeResource:   "groups",
		CA: CustomAttributes{
			Key: "testkey1",
		},
	})

	if err != nil {
		t.Errorf("CustomAttributes.GetCustomAttribute returned error: %v", err)
	}

	want := &CustomAttributes{Key: "testkey1", Value: "testvalue1"}
	if !reflect.DeepEqual(want, customAttribute) {
		t.Errorf("CustomAttributes.GetCustomAttribute returned %+v, want %+v", customAttribute, want)
	}
}

func TestSetCustomAttribute(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttributes.SetCustomAttribute(CustomAttributeOptions{
		CustomeAttributeResourceID: "2",
		CustomeAttributeResource:   "groups",
		CA: CustomAttributes{
			Key:   "testkey1",
			Value: "testvalue1",
		},
	})

	if err != nil {
		t.Errorf("CustomAttributes.SetCustomAttributes returned error: %v", err)
	}

	want := &CustomAttributes{Key: "testkey1", Value: "testvalue1"}
	if !reflect.DeepEqual(want, customAttribute) {
		t.Errorf("CustomAttributes.SetCustomAttributes returned %+v, want %+v", customAttribute, want)
	}

}

func TestDeleteCustomAttribute(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.CustomAttributes.DeleteCustomAttribute(CustomAttributeOptions{
		CustomeAttributeResourceID: "2",
		CustomeAttributeResource:   "groups",
		CA: CustomAttributes{
			Key: "testkey1",
		},
	})
	if err != nil {
		t.Errorf("CustomAttributes.DeleteCustomAttributes returned error: %v", err)
	}

	want := http.StatusAccepted
	got := resp.StatusCode
	if got != want {
		t.Errorf("CustomAttributes.DeleteCustomAttributes returned %d, want %d", got, want)
	}
}
