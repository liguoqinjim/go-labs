package main

import (
	"flag"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	dsn string
)

func init() {
	pflag.StringVarP(&dsn, "dsn", "d", "", "dsn")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if dsn == "" {
		log.Fatalf("dsn is empty")
	}
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	sentry.CaptureMessage("It works!")

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.

	defer func() {
		log.Println("sentry.Flush:", sentry.Flush(time.Second*20))
	}()

	resp, err := http.Get("http://example.com")
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("reported to Sentry: %s", err)
		return
	}
	defer resp.Body.Close()
}
