package command

import (
	"fmt"
	"strconv"

	"errors"

	"github.com/Ladicle/toggl/cache"
	"github.com/Ladicle/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func GetProjects(c *cli.Context) (projects toggl.Projects, err error) {
	projects = cache.GetContent().Projects
	if len(projects) == 0 || !c.GlobalBool("cache") {
		projects, err = toggl.FetchWorkspaceProjects(viper.GetString("token"), viper.GetInt("wid"))
		cache.SetProjects(projects)
		cache.Write()
	}
	return
}

func CmdShowProjects(c *cli.Context) error {
	projects, err := GetProjects(c)
	if err != nil {
		return err
	}

	writer := NewWriter(c)

	defer writer.Flush()

	for _, project := range projects {
		writer.Write([]string{strconv.Itoa(project.ID), project.Name})
	}

	return nil
}

func CmdCreateProject(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("You need to specify project name.")
	}

	projects := cache.GetContent().Projects
	project, err := toggl.CreateWorkspaceProject(
		viper.GetString("token"), viper.GetInt("wid"), c.Args().First())
	if err != nil {
		return err
	}
	projects = append(projects, project)
	cache.SetProjects(projects)
	cache.Write()

	writer := NewWriter(c)
	defer writer.Flush()
	writer.Write([]string{strconv.Itoa(project.ID), project.Name})
	return nil
}

func CmdDeleteProject(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("You need to specify project ID.")
	}

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}

	err = toggl.DeleteProject(viper.GetString("token"), id)
	if err != nil {
		return err
	}

	projects := cache.GetContent().Projects
	result := toggl.Projects{}
	for _, project := range projects {
		if project.ID != id {
			result = append(result, project)
		}
	}
	cache.SetProjects(result)
	cache.Write()

	writer := NewWriter(c)
	defer writer.Flush()
	writer.Write([]string{fmt.Sprintf("Delete %d completed.", id)})
	return nil
}
