package main

import (
	"time"

	"github.com/Makrorof/mongodbTutorial/internal/app"
	"github.com/Makrorof/mongodbTutorial/tools"
	"go.uber.org/zap"
)

//
//https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
//

func init() {
	tools.SetupLog(zap.DebugLevel, "./", time.Hour*24, time.Hour, true)
}

func main() {
	//LOAD CONFIG

	app.RunCLI()
}
