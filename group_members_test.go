//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListBillableGroupMembers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/groups/1/billable_members",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[{"id":1,"username":"ray","name":"Raymond","state":"active","avatar_url":"https://foo.bar/mypic","web_url":"http://192.168.1.8:3000/root","last_activity_on":"2021-01-27"}]`)
		})

	billableMembers, _, err := client.Groups.ListBillableGroupMembers(1, &ListBillableGroupMembersOptions{})
	if err != nil {
		t.Errorf("Groups.ListBillableGroupMembers returned error: %v", err)
	}

	testTime := ISOTime{}
	err = testTime.UnmarshalJSON([]byte(`"2021-01-27"`))
	if err != nil {
		t.Errorf("Could not unmarshal date string to ISOTime: %v", err)
	}
	want := []*BillableGroupMember{{ID: 1, Username: "ray", Name: "Raymond", State: "active", AvatarURL: "https://foo.bar/mypic", WebURL: "http://192.168.1.8:3000/root", LastActivityOn: testTime}}
	if !reflect.DeepEqual(want, billableMembers) {
		t.Errorf("Groups.ListBillableGroupMembers returned %+v, want %+v", billableMembers, want)
	}
}
