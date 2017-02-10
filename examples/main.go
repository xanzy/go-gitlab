package main

import "github.com/sindrepm/go-gitlab"
import "fmt"

func main() {
	// See the separate files in this directory for the examples. This file is only
	// here to provide a main() function for the `example` package, keeping Travis happy.

	c := gitlab.NewClient(nil, "LJsoVMYMZxpYCU9zY7_5")
	c.SetBaseURL("https://git.appstrada.net/api/v3/")
	s, _, err := c.NotificationSettings.UpdateGlobalSettings(gitlab.NotificationSettings{
		Level: gitlab.CustomNotificationLevel,
		Email: "admin@example.com",
		NotificationEventOptions: &gitlab.NotificationEventOptions{
			FailedPipelineEvent: true,
		},
	}, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
