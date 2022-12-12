package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIssueBoardsService_CreateIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
			  {
				"id": 1,
				"project": {
				  "id": 5,
				  "name": "Diaspora Project Site",
				  "name_with_namespace": "Diaspora / Diaspora Project Site",
				  "path": "diaspora-project-site",
				  "path_with_namespace": "diaspora/diaspora-project-site",
				  "http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
				  "web_url": "http://example.com/diaspora/diaspora-project-site"
				},
				"name": "newboard",
				"lists" : [],
				"group": null,
				"milestone": null,
				"assignee" : null,
				"labels" : [],
				"weight" : null
			  }
		`)
	})

	want := &IssueBoard{
		ID:   1,
		Name: "newboard",
		Project: &Project{
			ID:                5,
			HTTPURLToRepo:     "http://example.com/diaspora/diaspora-project-site.git",
			WebURL:            "http://example.com/diaspora/diaspora-project-site",
			Name:              "Diaspora Project Site",
			NameWithNamespace: "Diaspora / Diaspora Project Site",
			Path:              "diaspora-project-site",
			PathWithNamespace: "diaspora/diaspora-project-site",
		},
		Lists:  []*BoardList{},
		Labels: []*LabelDetails{},
	}

	ib, resp, err := client.Boards.CreateIssueBoard(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ib)

	ib, resp, err = client.Boards.CreateIssueBoard(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ib)

	ib, resp, err = client.Boards.CreateIssueBoard(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ib)

	ib, resp, err = client.Boards.CreateIssueBoard(7, nil, nil)
	require.Error(t, err)
	require.Nil(t, ib)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_UpdateIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
			  {
				"id": 1,
				"project": {
				  "id": 5,
				  "name": "Diaspora Project Site",
				  "name_with_namespace": "Diaspora / Diaspora Project Site",
				  "path": "diaspora-project-site",
				  "path_with_namespace": "diaspora/diaspora-project-site",
				  "http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
				  "web_url": "http://example.com/diaspora/diaspora-project-site"
				},
				"name": "new_name",
				"lists" : [],
				"group": null,
				"milestone": null,
				"assignee" : null,
				"labels" : [],
				"weight" : null
			  }
		`)
	})

	want := &IssueBoard{
		ID:   1,
		Name: "new_name",
		Project: &Project{
			ID:                5,
			HTTPURLToRepo:     "http://example.com/diaspora/diaspora-project-site.git",
			WebURL:            "http://example.com/diaspora/diaspora-project-site",
			Name:              "Diaspora Project Site",
			NameWithNamespace: "Diaspora / Diaspora Project Site",
			Path:              "diaspora-project-site",
			PathWithNamespace: "diaspora/diaspora-project-site",
		},
		Lists:  []*BoardList{},
		Labels: []*LabelDetails{},
	}

	ib, resp, err := client.Boards.UpdateIssueBoard(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ib)

	ib, resp, err = client.Boards.UpdateIssueBoard(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ib)

	ib, resp, err = client.Boards.UpdateIssueBoard(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ib)

	ib, resp, err = client.Boards.UpdateIssueBoard(7, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, ib)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_DeleteIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Boards.DeleteIssueBoard(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Boards.DeleteIssueBoard(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Boards.DeleteIssueBoard(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Boards.DeleteIssueBoard(7, 1, nil, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_ListIssueBoards(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
			[
			  {
				"id" : 1,
				"name": "board1",
				"project": {
				  "id": 5,
				  "name": "Diaspora Project Site",
				  "name_with_namespace": "Diaspora / Diaspora Project Site",
				  "path": "diaspora-project-site",
				  "path_with_namespace": "diaspora/diaspora-project-site",
				  "http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
				  "web_url": "http://example.com/diaspora/diaspora-project-site"
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
					"position" : 1,
					"max_issue_count": 0,
					"max_issue_weight": 0,
					"limit_metric": null
				  },
				  {
					"id" : 2,
					"label" : {
					  "name" : "Ready",
					  "color" : "#FF0000",
					  "description" : null
					},
					"position" : 2,
					"max_issue_count": 0,
					"max_issue_weight": 0,
					"limit_metric":  null
				  },
				  {
					"id" : 3,
					"label" : {
					  "name" : "Production",
					  "color" : "#FF5F00",
					  "description" : null
					},
					"position" : 3,
					"max_issue_count": 0,
					"max_issue_weight": 0,
					"limit_metric":  null
				  }
				]
			  }
			]
		`)
	})

	want := []*IssueBoard{{
		ID:   1,
		Name: "board1",
		Project: &Project{
			ID:                5,
			HTTPURLToRepo:     "http://example.com/diaspora/diaspora-project-site.git",
			WebURL:            "http://example.com/diaspora/diaspora-project-site",
			Name:              "Diaspora Project Site",
			NameWithNamespace: "Diaspora / Diaspora Project Site",
			Path:              "diaspora-project-site",
			PathWithNamespace: "diaspora/diaspora-project-site",
		},
		Milestone: &Milestone{
			ID:    12,
			Title: "10.0",
		},
		Lists: []*BoardList{
			{
				ID: 1,
				Label: &Label{
					Name:  "Testing",
					Color: "#F0AD4E",
				}, Position: 1,
			},
			{
				ID: 2,
				Label: &Label{
					Name:  "Ready",
					Color: "#FF0000",
				},
				Position: 2,
			},
			{
				ID: 3,
				Label: &Label{
					Name:  "Production",
					Color: "#FF5F00",
				},
				Position: 3,
			},
		},
	}}

	ibs, resp, err := client.Boards.ListIssueBoards(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ibs)

	ibs, resp, err = client.Boards.ListIssueBoards(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ibs)

	ibs, resp, err = client.Boards.ListIssueBoards(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ibs)

	ibs, resp, err = client.Boards.ListIssueBoards(7, nil, nil)
	require.Error(t, err)
	require.Nil(t, ibs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_GetIssueBoard(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id" : 1,
			"name": "board1",
			"project": {
			  "id": 5,
			  "name": "Diaspora Project Site",
			  "name_with_namespace": "Diaspora / Diaspora Project Site",
			  "path": "diaspora-project-site",
			  "path_with_namespace": "diaspora/diaspora-project-site",
			  "http_url_to_repo": "http://example.com/diaspora/diaspora-project-site.git",
			  "web_url": "http://example.com/diaspora/diaspora-project-site"
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
				"position" : 1,
				"max_issue_count": 0,
				"max_issue_weight": 0,
				"limit_metric": null
			  },
			  {
				"id" : 2,
				"label" : {
				  "name" : "Ready",
				  "color" : "#FF0000",
				  "description" : null
				},
				"position" : 2,
				"max_issue_count": 0,
				"max_issue_weight": 0,
				"limit_metric":  null
			  },
			  {
				"id" : 3,
				"label" : {
				  "name" : "Production",
				  "color" : "#FF5F00",
				  "description" : null
				},
				"position" : 3,
				"max_issue_count": 0,
				"max_issue_weight": 0,
				"limit_metric":  null
			  }
			]
		  }
		`)
	})

	want := &IssueBoard{
		ID:   1,
		Name: "board1",
		Project: &Project{
			ID:                5,
			HTTPURLToRepo:     "http://example.com/diaspora/diaspora-project-site.git",
			WebURL:            "http://example.com/diaspora/diaspora-project-site",
			Name:              "Diaspora Project Site",
			NameWithNamespace: "Diaspora / Diaspora Project Site",
			Path:              "diaspora-project-site",
			PathWithNamespace: "diaspora/diaspora-project-site",
		},
		Milestone: &Milestone{
			ID:    12,
			Title: "10.0",
		},
		Lists: []*BoardList{
			{
				ID: 1,
				Label: &Label{
					Name:  "Testing",
					Color: "#F0AD4E",
				},
				Position: 1,
			},
			{
				ID: 2,
				Label: &Label{
					Name:  "Ready",
					Color: "#FF0000",
				},
				Position: 2,
			},
			{
				ID: 3,
				Label: &Label{
					Name:  "Production",
					Color: "#FF5F00",
				},
				Position: 3,
			},
		},
	}

	ib, resp, err := client.Boards.GetIssueBoard(5, 1, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, ib)

	ib, resp, err = client.Boards.GetIssueBoard(5.01, 1, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, ib)

	ib, resp, err = client.Boards.GetIssueBoard(5, 1, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, ib)

	ib, resp, err = client.Boards.GetIssueBoard(7, 1, nil, nil)
	require.Error(t, err)
	require.Nil(t, ib)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_GetIssueBoardLists(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1/lists", func(w http.ResponseWriter, r *http.Request) {
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
				"position" : 1,
				"max_issue_count": 0,
				"max_issue_weight": 0,
				"limit_metric":  null
			  },
			  {
				"id" : 2,
				"label" : {
				  "name" : "Ready",
				  "color" : "#FF0000",
				  "description" : null
				},
				"position" : 2,
				"max_issue_count": 0,
				"max_issue_weight": 0,
				"limit_metric":  null
			  },
			  {
				"id" : 3,
				"label" : {
				  "name" : "Production",
				  "color" : "#FF5F00",
				  "description" : null
				},
				"position" : 3,
				"max_issue_count": 0,
				"max_issue_weight": 0,
				"limit_metric":  null
			  }
			]
		`)
	})

	want := []*BoardList{
		{
			ID: 1,
			Label: &Label{
				Name:  "Testing",
				Color: "#F0AD4E",
			},
			Position: 1,
		},
		{
			ID: 2,
			Label: &Label{
				Name:  "Ready",
				Color: "#FF0000",
			},
			Position: 2,
		},
		{
			ID: 3,
			Label: &Label{
				Name:  "Production",
				Color: "#FF5F00",
			},
			Position: 3,
		},
	}

	bls, resp, err := client.Boards.GetIssueBoardLists(5, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bls)

	bls, resp, err = client.Boards.GetIssueBoardLists(5.01, 1, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bls)

	bls, resp, err = client.Boards.GetIssueBoardLists(5, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bls)

	bls, resp, err = client.Boards.GetIssueBoardLists(3, 1, nil)
	require.Error(t, err)
	require.Nil(t, bls)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_GetIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		  {
			"id" : 1,
			"label" : {
			  "name" : "Testing",
			  "color" : "#F0AD4E",
			  "description" : null
			},
			"position" : 1,
			"max_issue_count": 0,
			"max_issue_weight": 0,
			"limit_metric":  null
		  }
		`)
	})

	want := &BoardList{
		ID: 1,
		Label: &Label{
			Name:  "Testing",
			Color: "#F0AD4E",
		},
		Position: 1,
	}

	bl, resp, err := client.Boards.GetIssueBoardList(5, 1, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bl)

	bl, resp, err = client.Boards.GetIssueBoardList(5.01, 1, 1, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.Boards.GetIssueBoardList(5, 1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.Boards.GetIssueBoardList(3, 1, 1, nil)
	require.Error(t, err)
	require.Nil(t, bl)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_CreateIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
		  {
			"id" : 1,
			"label" : {
			  "name" : "Testing",
			  "color" : "#F0AD4E",
			  "description" : null
			},
			"position" : 1,
			"max_issue_count": 0,
			"max_issue_weight": 0,
			"limit_metric":  null
		  }
		`)
	})

	want := &BoardList{
		ID: 1,
		Label: &Label{
			Name:  "Testing",
			Color: "#F0AD4E",
		},
		Position: 1,
	}

	bl, resp, err := client.Boards.CreateIssueBoardList(5, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bl)

	bl, resp, err = client.Boards.CreateIssueBoardList(5.01, 1, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.Boards.CreateIssueBoardList(5, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.Boards.CreateIssueBoardList(3, 1, nil)
	require.Error(t, err)
	require.Nil(t, bl)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_UpdateIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprintf(w, `
		  {
			"id" : 1,
			"label" : {
			  "name" : "Testing",
			  "color" : "#F0AD4E",
			  "description" : null
			},
			"position" : 1,
			"max_issue_count": 0,
			"max_issue_weight": 0,
			"limit_metric":  null
		  }
		`)
	})

	want := &BoardList{
		ID: 1,
		Label: &Label{
			Name:  "Testing",
			Color: "#F0AD4E",
		},
		Position: 1,
	}

	bl, resp, err := client.Boards.UpdateIssueBoardList(5, 1, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, bl)

	bl, resp, err = client.Boards.UpdateIssueBoardList(5.01, 1, 1, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.Boards.UpdateIssueBoardList(5, 1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, bl)

	bl, resp, err = client.Boards.UpdateIssueBoardList(3, 1, 1, nil)
	require.Error(t, err)
	require.Nil(t, bl)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestIssueBoardsService_DeleteIssueBoardList(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/5/boards/1/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Boards.DeleteIssueBoardList(5, 1, 1, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp, err = client.Boards.DeleteIssueBoardList(5.01, 1, 1, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)

	resp, err = client.Boards.DeleteIssueBoardList(5, 1, 1, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)

	resp, err = client.Boards.DeleteIssueBoardList(3, 1, 1, nil)
	require.Error(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
