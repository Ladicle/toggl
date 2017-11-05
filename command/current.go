package command

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Ladicle/toggl/cache"
	"github.com/Ladicle/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func formatTimeDuration(duration time.Duration) string {
	hours := duration / time.Hour
	minutes := duration / time.Minute % 60
	seconds := duration / time.Second % 60
	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}

func calcDuration(duration int64) time.Duration {
	second := duration + time.Now().Unix()
	return time.Duration(second * int64(time.Second))
}

func CmdCurrent(c *cli.Context) error {
	var project toggl.Project
	var timeEntry toggl.TimeEntry
	var workspace toggl.Workspace

	timeEntry = cache.GetContent().CurrentTimeEntry

	if !c.GlobalBool("cache") {
		current, err := toggl.GetCurrentTimeEntry(viper.GetString("token"))
		timeEntry = current.Data
		if err != nil {
			return err
		}
		cache.SetCurrentTimeEntry(timeEntry)
		cache.Write()

		workspaces, err := GetWorkspaces(c)
		if err != nil {
			return err
		}
		workspace, err = workspaces.FindByID(timeEntry.WID)

		if timeEntry.PID != 0 {
			projects, err := GetProjects(c)
			if err != nil {
				return err
			}
			project, err = projects.FindByID(timeEntry.PID)
			if err != nil {
				return err
			}
		}
	}

	if timeEntry.ID == 0 {
		fmt.Println("No time entry")
		return nil
	}

	writer := NewWriter(c)

	defer writer.Flush()

	records := [][]string{
		[]string{"ID", strconv.Itoa(timeEntry.ID)},
		[]string{"Description", timeEntry.Description},
		[]string{"Project", project.Name},
		[]string{"Workspace", workspace.Name},
		[]string{"Duration", formatTimeDuration(calcDuration(timeEntry.Duration))},
	}

	for _, record := range records {
		writer.Write(record)
	}

	return nil
}
