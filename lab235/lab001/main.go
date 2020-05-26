package main

import (
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"

	"github.com/slack-go/slack"
)

var (
	token string
)

func init() {
	pflag.StringVarP(&token, "token", "t", "", "token")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if token == "" {
		log.Fatalf("need token")
	}
}
func main() {
	log.Println(token)
	api := slack.New(token, slack.OptionDebug(true))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))

	chans, err := api.GetChannels(false)
	if err != nil {
		log.Fatalf("api.GetChannels error:%v", err)
	}
	for _, v := range chans {
		log.Println(v)
	}

	//groups, err := api.GetGroups(false)
	//if err != nil {
	//	log.Fatalf("api.GetGroups error:%v", err)
	//}
	//for _, group := range groups {
	//	log.Printf("ID: %s, Name: %s", group.ID, group.Name)
	//}

}
