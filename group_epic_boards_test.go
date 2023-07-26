package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupEpicBoardsService_ListGroupEpicBoards(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/5/epic_boards", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `
		[
			{
			  "id": 1,
			  "name": "group epic board",
			  "group": {
				"id": 5,
				"name": "Documentcloud",
				"web_url": "http://example.com/groups/documentcloud"
			  },
			  "hide_backlog_list": false,
			  "hide_closed_list": false,
			  "labels": [
				{
				  "id": 1,
				  "name": "Board Label",
				  "color": "#c21e56",
				  "group_id": 5,
				  "description": "label applied to the epic board"
				}
			  ],
			  "lists": [
				{
				  "id": 1,
				  "label": {
					"id": 69,
					"name": "Testing",
					"color": "#F0AD4E",
					"description": null
				  },
				  "position": 1,
				  "list_type": "label"
				},
				{
				  "id": 2,
				  "label": {
					"id": 70,
					"name": "Ready",
					"color": "#FF0000",
					"description": null
				  },
				  "position": 2,
				  "list_type": "label"
				},
				{
				  "id": 3,
				  "label": {
					"id": 71,
					"name": "Production",
					"color": "#FF5F00",
					"description": null
				  },
				  "position": 3,
				  "list_type": "label"
				}
			  ]
			}
		  ]		  
		`)
	})

	want := []*GroupEpicBoard{{
		ID:   1,
		Name: "group epic board",
		Group: &Group{
			ID:     5,
			Name:   "Documentcloud",
			WebURL: "http://example.com/groups/documentcloud",
		},
		Labels: []*LabelDetails{
			{
				ID:          1,
				Name:        "Board Label",
				Color:       "#c21e56",
				Description: "label applied to the epic board",
			},
		},
		Lists: []*BoardList{
			{
				ID: 1,
				Label: &Label{
					ID:          69,
					Name:        "Testing",
					Color:       "#F0AD4E",
					Description: "",
				},
				Position: 1,
			},
			{
				ID: 2,
				Label: &Label{
					ID:          70,
					Name:        "Ready",
					Color:       "#FF0000",
					Description: "",
				},
				Position: 2,
			},
			{
				ID: 3,
				Label: &Label{
					ID:          71,
					Name:        "Production",
					Color:       "#FF5F00",
					Description: "",
				},
				Position: 3,
			},
		},
	}}

	gibs, resp, err := client.GroupEpicBoards.ListGroupEpicBoards(5, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, gibs)

	gibs, resp, err = client.GroupEpicBoards.ListGroupEpicBoards(5.01, nil, nil)
	require.EqualError(t, err, "invalid ID type 5.01, the ID must be an int or a string")
	require.Nil(t, resp)
	require.Nil(t, gibs)

	gibs, resp, err = client.GroupEpicBoards.ListGroupEpicBoards(5, nil, nil, errorOption)
	require.EqualError(t, err, "RequestOptionFunc returns an error")
	require.Nil(t, resp)
	require.Nil(t, gibs)

	gibs, resp, err = client.GroupEpicBoards.ListGroupEpicBoards(3, nil, nil)
	require.Error(t, err)
	require.Nil(t, gibs)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
