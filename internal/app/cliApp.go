package app

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Makrorof/mongodbTutorial/internal/model"
	"github.com/Makrorof/mongodbTutorial/internal/repository/task/taskrepo"
	"github.com/Makrorof/mongodbTutorial/internal/usecase/itasks/tasksusecase"
	"github.com/Makrorof/mongodbTutorial/tools"
	"github.com/gookit/color"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func RunCLI() {
	collection, err := tools.SetupMongo("localhost", "27017")

	if err != nil {
		zap.L().Panic("Couldn't connect to the database.", zap.Error(err))
		return
	}

	taskRepo, err := taskrepo.New(collection)
	if err != nil {
		zap.L().Panic("Couldn't connect to the database.", zap.Error(err))
		return
	}

	tasksUsecase := tasksusecase.New(taskRepo)

	app := &cli.App{
		Name:  "tasker",
		Usage: "manage tasks",
		Action: func(c *cli.Context) error {
			tasks, err := tasksUsecase.GetPending(context.TODO())
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
					return nil
				}
				return err
			}

			printTasks(tasks)
			return nil
		},
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

					data, created, err := tasksUsecase.Create(context.TODO(), str)

					if err == nil && created {
						fmt.Println("CREATED OK. ID: ", data.ID)
					} else if err == nil {
						fmt.Println("UPDATED OK. ID: ", data.ID)
					}

					return err
				},
			},
			{
				Name:    "all",
				Aliases: []string{"l"},
				Usage:   "list all tasks",
				Action: func(c *cli.Context) error {
					tasks, err := tasksUsecase.GetAll(context.TODO())
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
							return nil
						}

						return err
					}

					printTasks(tasks)
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					text := c.Args().First()

					return tasksUsecase.Update(context.TODO(), &model.UpdateTask{
						FilterText: text,
						Completed:  true,
					})
				},
			},
			{
				Name:    "finished",
				Aliases: []string{"f"},
				Usage:   "list completed tasks",
				Action: func(c *cli.Context) error {
					tasks, err := tasksUsecase.GetFinished(context.TODO())
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `done 'task'` to complete a task")
							return nil
						}

						return err
					}

					printTasks(tasks)
					return nil
				},
			},
		},
	}

	if err = app.Run(os.Args); err != nil {
		zap.L().Panic("The CLI app couldn't be started.", zap.Error(err))
		return
	}
}

func printTasks(tasks []*model.Task) {
	for i, v := range tasks {
		if v.Completed {
			color.Green.Printf("%d: %s\n", i+1, v.Text)
		} else {
			color.Yellow.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}
