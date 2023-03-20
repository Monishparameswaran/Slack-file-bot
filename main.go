package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	// os package provide us functionalities to interact with os
	//Environment variable are the KEY VALUE PAIR that is used to store important info such useername passsword
	// to set the environment variable we use os.setenv function
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5002117372720-4971960476470-dQjEp1k9u2u1iE6qPbr09kJm")
	os.Setenv("CHANNEL_ID", "C04UUTWVDA8")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"gio.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s is the error", err)
			return
		}
		fmt.Printf("name :%s,URL:%s \n", file.Name, file.URL)
	}
}
