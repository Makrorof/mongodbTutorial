package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Makrorof/mongodbTutorial/internal/repository/task/taskrepo"
	"github.com/Makrorof/mongodbTutorial/internal/usecase/itasks/tasksusecase"
	"github.com/Makrorof/mongodbTutorial/tools"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func init() {
	tools.SetupLog(zap.DebugLevel, "./", time.Hour*24, time.Hour, true)
}

func main() {
	collection, err := tools.SetupMongo("localhost", "27017")

	if err != nil {
		zap.L().Panic("Couldn't connect to the database.", zap.Error(err))
		return
	}

	taskRepo := taskrepo.New(collection)
	tasksUsecase := tasksusecase.New(taskRepo)

	app := &cli.App{
		Name:  "tasker",
		Usage: "manage tasks",
		Commands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("cannot add an empty task")
					}

					data, err := tasksUsecase.Create(context.TODO(), str)
					fmt.Println("Data: ", data)

					return err
				},
			},
		},
	}

	if err = app.Run(os.Args); err != nil {
		zap.L().Panic("The CLI app couldn't be started.", zap.Error(err))
		return
	}
}
