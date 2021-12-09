package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alexsuslov/godotenv"
	"github.com/alexsuslov/gokdesk"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var version string
var help string
var debugger bool

var config string
var get string
var update string
var status string


func init() {
	// Config
	flag.StringVar(&config, "config", ".env", "path to  env file")
	//GET
	flag.StringVar(&get, "get", "", "get issue by id")
	// SetStatus
	flag.StringVar(&status, "status", "", "set issue status by id")

	flag.Parse()
}

func main() {
	if err := godotenv.Load(config); err != nil {
		logrus.Warningf("no %s file", config)
	}
	// Getter
	if get != "" {
		body, err := gokdesk.Getter(context.Background(), get)
		Done(body, err)
		return
	}

	if status!=""{
		body, err := gokdesk.SetStatusRAW(context.Background(), status, os.Stdin, nil)
		Done(body, err)
		return
	}

	// help
	fmt.Printf("Gokdesk is a Golang wrapper for accessing OKDESK using the REST API. Version %v \n", version)
}

func Done(body io.ReadCloser, err error) {
	if err != nil {
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err != nil {
		panic(err)
	}
}
