package tools

import (
	"context"
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupMongo(host string, port string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", host, port))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database("tasker").Collection("tasks"), nil
}

func SetupLog(level zapcore.Level, path string, maxAge time.Duration, rotationTime time.Duration, devMode bool) {
	// initialize the rotator
	logFile := path + "/app-%Y-%m-%d-%H-%M-%S.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	if err != nil {
		panic(err)
	}

	// initialize the JSON encoding config
	encCfg := zap.NewProductionEncoderConfig()

	if devMode {
		encCfg = zap.NewDevelopmentEncoderConfig()
	}

	// add the encoder config and rotator to create a new zap logger
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encCfg),
		w,
		level)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)
}
