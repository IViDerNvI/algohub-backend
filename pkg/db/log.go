package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type MyDatabaseLog struct{}

func (c *MyDatabaseLog) LogMode(level logger.LogLevel) logger.Interface {
	return c
}

func (c *MyDatabaseLog) Info(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"message": msg,
		"data":    data,
		"realm":   "pkg/db",
	}).Infof("[Database]")
}

func (c *MyDatabaseLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"message": msg,
		"data":    data,
		"realm":   "pkg/db",
	}).Warnf("[Database]")
}

func (c *MyDatabaseLog) Error(ctx context.Context, msg string, data ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"message": msg,
		"data":    data,
		"realm":   "pkg/db",
	}).Errorf("[Database]")
}

func (c *MyDatabaseLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"sql":   sql,
			"rows":  rows,
			"realm": "pkg/db",
		}).Errorf("[Database] %s", err)

	} else {
		logrus.WithFields(logrus.Fields{
			"sql":     sql,
			"rows":    rows,
			"elapsed": elapsed,
			"realm":   "pkg/db",
		}).Infof("[Database] SQL executed.")
	}
}
