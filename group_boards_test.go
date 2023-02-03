package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupIssueBoardsService_ListGroupIssueBoards(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id": 1,
				"name": "group issue board",
				"group": {
				  "id": 5,
				  "name": "Documentcloud",
				  "web_url": "http://example.com/groups/documentcloud"
				},
				"milestone":   {
				  "id": 12,
				  "title": "10.0"
				},
				"lists" : [
				  {
					"id" : 1,
					"label" : {
					  "name" : "Testing",
					  "color" : "#F0AD4E",
					  "description" : null
					},
					"position" : 1
				  },
				  {
					"id" : 2,
					"label" : {
					  "name" : "Ready",
					  "color" : "#FF0000",
					  "description" : null
					},
					"position" : 2
				  },
				  {
					"id" : 3,
					"label" : {
					  "name" : "Production",
					  "color" : "#FF5F00",
					  "description" : null
					},
					"position" : 3
				  }
				]
			  }
			]
		`)
	})

	want := []*GroupIssueBoard{{
		ID:   1,
		Name: "group issue board",
		Group: &Group{
			ID:     5,
			Name:   "Documentcloud",
			WebURL: "http://example.com/groups/documentcloud",
		},
		Milestone: &Milestone{
			ID:          12,
			IID:         0,
			ProjectID:   0,
			Title:       "10.0",
			Description: "",
		},
		Lists: []*BoardList{
			{
				ID: 1,
				Label: &Label{
					Name:        "Testing",
					Color:       "#F0AD4E",
					Description: "",
				},
				Position: 1,
			},
			{
				ID: 2,
				Label: &Label{
					Name:        "Ready",
					Color:       "#FF0000",
					Description: "",
				},
				Position: 2,
			},
			{
				ID: 3,
				Label: &Label{
					Name:        "Production",
					Color:       "#FF5F00",
					Description: "",
				},
				Position: 3,
			},
		},
	}}

	gibs, resp, err := client.GroupIssueBoards.ListGroupIssueBoards(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gibs)

	gibs, resp, err = client.GroupIssueBoards.ListGroupIssueBoards(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gibs)

	gibs, resp, err = client.GroupIssueBoards.ListGroupIssueBoards(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gibs)

	gibs, resp, err = client.GroupIssueBoards.ListGroupIssueBoards(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, gibs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_CreateGroupIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			  {
				"id": 1,
				"name": "newboard",
				"project": null,
				"lists" : [],
				"group": {
				  "id": 5,
				  "name": "Documentcloud",
				  "web_url": "http://example.com/groups/documentcloud"
				},
				"milestone": null,
				"assignee" : null,
				"labels" : [],
				"weight" : null
			  }
		`)
	})

	want := &GroupIssueBoard{
		ID:   1,
		Name: "newboard",
		Group: &Group{
			ID:     5,
			Name:   "Documentcloud",
			WebURL: "http://example.com/groups/documentcloud",
		},
		Milestone: nil,
		Lists:     []*BoardList{},
	}

	gib, resp, err := client.GroupIssueBoards.CreateGroupIssueBoard(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gib)

	gib, resp, err = client.GroupIssueBoards.CreateGroupIssueBoard(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gib)

	gib, resp, err = client.GroupIssueBoards.CreateGroupIssueBoard(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gib)

	gib, resp, err = client.GroupIssueBoards.CreateGroupIssueBoard(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, gib)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_GetGroupIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id": 1,
			"name": "group issue board",
			"group": {
			  "id": 5,
			  "name": "Documentcloud",
			  "web_url": "http://example.com/groups/documentcloud"
			},
			"milestone":   {
			  "id": 12,
			  "title": "10.0"
			},
			"lists" : [
			  {
				"id" : 1,
				"label" : {
				  "name" : "Testing",
				  "color" : "#F0AD4E",
				  "description" : null
				},
				"position" : 1
			  },
			  {
				"id" : 2,
				"label" : {
				  "name" : "Ready",
				  "color" : "#FF0000",
				  "description" : null
				},
				"position" : 2
			  },
			  {
				"id" : 3,
				"label" : {
				  "name" : "Production",
				  "color" : "#FF5F00",
				  "description" : null
				},
				"position" : 3
			  }
			]
		  }
		`)
	})

	want := &GroupIssueBoard{
		ID:   1,
		Name: "group issue board",
		Group: &Group{
			ID:     5,
			Name:   "Documentcloud",
			WebURL: "http://example.com/groups/documentcloud",
		},
		Milestone: &Milestone{
			ID:          12,
			IID:         0,
			ProjectID:   0,
			Title:       "10.0",
			Description: "",
		},
		Lists: []*BoardList{
			{
				ID: 1,
				Label: &Label{
					Name:        "Testing",
					Color:       "#F0AD4E",
					Description: "",
				},
				Position: 1,
			},
			{
				ID: 2,
				Label: &Label{
					Name:        "Ready",
					Color:       "#FF0000",
					Description: "",
				},
				Position: 2,
			},
			{
				ID: 3,
				Label: &Label{
					Name:        "Production",
					Color:       "#FF5F00",
					Description: "",
				},
				Position: 3,
			},
		},
	}

	gib, resp, err := client.GroupIssueBoards.GetGroupIssueBoard(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gib)

	gib, resp, err = client.GroupIssueBoards.GetGroupIssueBoard(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gib)

	gib, resp, err = client.GroupIssueBoards.GetGroupIssueBoard(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gib)

	gib, resp, err = client.GroupIssueBoards.GetGroupIssueBoard(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, gib)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_UpdateIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			  {
				"id": 1,
				"project": null,
				"lists": [],
				"name": "new_name",
				"group": {
				  "id": 5,
				  "name": "Documentcloud",
				  "web_url": "http://example.com/groups/documentcloud"
				},
				"milestone": {
				  "id": 44,
				  "iid": 1,
				  "group_id": 5,
				  "title": "Group Milestone",
				  "description": "Group Milestone Desc",
				  "state": "active",
				  "web_url": "http://example.com/groups/documentcloud/-/milestones/1"
				},
				"assignee": {
				  "id": 1,
				  "name": "Administrator",
				  "username": "root",
				  "state": "active",
				  "avatar_url": "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon",
				  "web_url": "http://example.com/root"
				},
				"labels": [{
				  "id": 11,
				  "name": "GroupLabel",
				  "color": "#428BCA",
				  "description": ""
				}],
				"weight": 4
			  }
		`)
	})

	want := &GroupIssueBoard{
		ID:   1,
		Name: "new_name",
		Group: &Group{
			ID:     5,
			Name:   "Documentcloud",
			WebURL: "http://example.com/groups/documentcloud",
		},
		Milestone: &Milestone{
			ID:          44,
			IID:         1,
			ProjectID:   0,
			Title:       "Group Milestone",
			Description: "Group Milestone Desc",
			State:       "active",
			WebURL:      "http://example.com/groups/documentcloud/-/milestones/1",
		},
		Lists: []*BoardList{},
	}

	gib, resp, err := client.GroupIssueBoards.UpdateIssueBoard(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gib)

	gib, resp, err = client.GroupIssueBoards.UpdateIssueBoard(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gib)

	gib, resp, err = client.GroupIssueBoards.UpdateIssueBoard(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gib)

	gib, resp, err = client.GroupIssueBoards.UpdateIssueBoard(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, gib)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_DeleteIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.GroupIssueBoards.DeleteIssueBoard(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.GroupIssueBoards.DeleteIssueBoard(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.GroupIssueBoards.DeleteIssueBoard(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.GroupIssueBoards.DeleteIssueBoard(3, 1, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_ListGroupIssueBoardLists(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id" : 1,
				"label" : {
				  "name" : "Testing",
				  "color" : "#F0AD4E",
				  "description" : null
				},
				"position" : 1
			  },
			  {
				"id" : 2,
				"label" : {
				  "name" : "Ready",
				  "color" : "#FF0000",
				  "description" : null
				},
				"position" : 2
			  },
			  {
				"id" : 3,
				"label" : {
				  "name" : "Production",
				  "color" : "#FF5F00",
				  "description" : null
				},
				"position" : 3
			  }
			]
		`)
	})

	want := []*BoardList{
		{
			ID: 1,
			Label: &Label{
				Name:        "Testing",
				Color:       "#F0AD4E",
				Description: "",
			},
			Position: 1,
		},
		{
			ID: 2,
			Label: &Label{
				Name:        "Ready",
				Color:       "#FF0000",
				Description: "",
			},
			Position: 2,
		},
		{
			ID: 3,
			Label: &Label{
				Name:        "Production",
				Color:       "#FF5F00",
				Description: "",
			},
			Position: 3,
		},
	}

	bls, resp, err := client.GroupIssueBoards.ListGroupIssueBoardLists(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bls)

	bls, resp, err = client.GroupIssueBoards.ListGroupIssueBoardLists(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bls)

	bls, resp, err = client.GroupIssueBoards.ListGroupIssueBoardLists(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bls)

	bls, resp, err = client.GroupIssueBoards.ListGroupIssueBoardLists(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, bls)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_GetGroupIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id" : 1,
			"label" : {
			  "name" : "Testing",
			  "color" : "#F0AD4E",
			  "description" : null
			},
			"position" : 1
		  }
		`)
	})

	want := &BoardList{
		ID: 1,
		Label: &Label{
			Name:        "Testing",
			Color:       "#F0AD4E",
			Description: "",
		},
		Position: 1,
	}

	bl, resp, err := client.GroupIssueBoards.GetGroupIssueBoardList(5, 1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bl)

	bl, resp, err = client.GroupIssueBoards.GetGroupIssueBoardList(5.01, 1, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.GroupIssueBoards.GetGroupIssueBoardList(5, 1, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.GroupIssueBoards.GetGroupIssueBoardList(3, 1, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, bl)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_CreateGroupIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			{
			  "id": 9,
			  "label": null,
			  "position": 0,
			  "milestone": {
				"id": 7,
				"iid": 3,
				"group_id": 12,
				"title": "Milestone with due date",
				"description": "",
				"state": "active",
				"web_url": "https://gitlab.example.com/groups/issue-reproduce/-/milestones/3"
			  }
			}
		`)
	})

	want := &BoardList{
		ID:       9,
		Label:    nil,
		Position: 0,
		Milestone: &Milestone{
			ID:          7,
			IID:         3,
			Title:       "Milestone with due date",
			Description: "",
			State:       "active",
			WebURL:      "https://gitlab.example.com/groups/issue-reproduce/-/milestones/3",
		},
	}

	bl, resp, err := client.GroupIssueBoards.CreateGroupIssueBoardList(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bl)

	bl, resp, err = client.GroupIssueBoards.CreateGroupIssueBoardList(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.GroupIssueBoards.CreateGroupIssueBoardList(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.GroupIssueBoards.CreateGroupIssueBoardList(3, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, bl)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_UpdateIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			[
			  {
				"id" : 1,
				"label" : {
				  "name" : "Testing",
				  "color" : "#F0AD4E",
				  "description" : null
				},
				"position" : 1
			  },
			  {
				"id" : 2,
				"label" : {
				  "name" : "Ready",
				  "color" : "#FF0000",
				  "description" : null
				},
				"position" : 2
			  },
			  {
				"id" : 3,
				"label" : {
				  "name" : "Production",
				  "color" : "#FF5F00",
				  "description" : null
				},
				"position" : 3
			  }
			]
		`)
	})

	want := []*BoardList{
		{
			ID: 1,
			Label: &Label{
				Name:        "Testing",
				Color:       "#F0AD4E",
				Description: "",
			},
			Position: 1,
		},
		{
			ID: 2,
			Label: &Label{
				Name:        "Ready",
				Color:       "#FF0000",
				Description: "",
			},
			Position: 2,
		},
		{
			ID: 3,
			Label: &Label{
				Name:        "Production",
				Color:       "#FF5F00",
				Description: "",
			},
			Position: 3,
		},
	}

	bl, resp, err := client.GroupIssueBoards.UpdateIssueBoardList(5, 1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bl)

	bl, resp, err = client.GroupIssueBoards.UpdateIssueBoardList(5.01, 1, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.GroupIssueBoards.UpdateIssueBoardList(5, 1, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.GroupIssueBoards.UpdateIssueBoardList(3, 1, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, bl)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGroupIssueBoardsService_DeleteGroupIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/boards/1/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.GroupIssueBoards.DeleteGroupIssueBoardList(5, 1, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.GroupIssueBoards.DeleteGroupIssueBoardList(5.01, 1, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.GroupIssueBoards.DeleteGroupIssueBoardList(5, 1, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.GroupIssueBoards.DeleteGroupIssueBoardList(3, 1, 1, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
