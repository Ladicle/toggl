package toggl

import (
	"errors"
	"strconv"
)

type Project struct {
	Active        bool   `json:"active"`
	ActualHours   int    `json:"actual_hours"`
	At            string `json:"at"`
	AutoEstimates bool   `json:"auto_estimates"`
	Billable      bool   `json:"billable"`
	Color         string `json:"color"`
	CreatedAt     string `json:"created_at"`
	HexColor      string `json:"hex_color"`
	ID            int    `json:"id"`
	IsPrivate     bool   `json:"is_private"`
	Name          string `json:"name"`
	Template      bool   `json:"template"`
	Wid           int    `json:"wid"`
}

type Data struct {
	Project Project `json:"data"`
}

type Projects []Project

func (repository Projects) FindByID(id int) (Project, error) {
	for _, item := range repository {
		if item.ID == id {
			return item, nil
		}
	}
	return Project{}, errors.New("Find Failed")
}

func FetchWorkspaceProjects(token string, wid int) (Projects, error) {
	var workspaces Projects
	res, err := Request("GET", "/workspaces/"+strconv.Itoa(wid)+"/projects", nil, token)
	if err != nil {
		return Projects{}, err
	}
	err = res.Body.FromJsonTo(&workspaces)
	if err != nil {
		return Projects{}, err
	}
	return workspaces, nil
}

func CreateWorkspaceProject(token string, wid int, name string) (Project, error) {
	param := make(map[string]map[string]interface{})
	param["project"] = make(map[string]interface{})
	param["project"]["wid"] = wid
	param["project"]["name"] = name

	var data Data
	res, err := Request("POST", "/projects", param, token)
	if err != nil {
		return Project{}, err
	}
	err = res.Body.FromJsonTo(&data)
	if err != nil {
		return Project{}, err
	}
	return data.Project, nil
}

func DeleteProject(token string, id int) error {
	_, err := Request("DELETE", "/projects/"+strconv.Itoa(id), nil, token)
	return err
}
