package log

import (
	cfg2 "console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/constants"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var logService *Log
var logOnce sync.Once

func NewLogService() *Log {
	logOnce.Do(func() {
		logService = &Log{}
		logService.init()
	})
	return logService
}

type Log struct {
	logger *logrus.Logger
	cfg    *cfg2.Cfg
}

func (l Log) Print(args ...interface{}) {
	l.logger.Println(args)
}

func (l Log) Trace(args ...interface{}) {
	l.logger.Traceln(args)
}

func (l Log) Debug(args ...interface{}) {
	l.logger.Debugln(args)
}

func (l Log) Info(args ...interface{}) {
	l.logger.Infoln(args)
}

func (l Log) Warn(args ...interface{}) {
	l.logger.Warnln(args)
}

func (l Log) Error(args ...interface{}) {
	l.logger.Errorln(args)
}

func (l Log) Fatal(args ...interface{}) {
	l.logger.Fatalln(args)
}

func (l Log) Panic(args ...interface{}) {
	logrus.Panicln(args)
}

func (l *Log) init() {
	l.logger = logrus.New()
	l.cfg = cfg2.NewCfgService()

	l.initLogrus()
}

func (l *Log) initLogrus() {
	if l.cfg.Log.LogLevel < 7 {
		l.logger.SetLevel(logrus.Level(l.cfg.Log.LogLevel))
	}
	l.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: constants.DefaultTimestampFormat, DisableColors: false,
		ForceColors: true, FullTimestamp: true})
	if l.cfg.Log.LogFile != "" { // 如果日志文件非空, 将日志打到对应文件
		var fd *rotatelogs.RotateLogs
		if l.cfg.Log.LogRotationCount > 0 {
			fd, _ = rotatelogs.New(
				l.cfg.Log.LogFile+".%Y%m%d",
				rotatelogs.WithLinkName(l.cfg.Log.LogFile),
				rotatelogs.WithMaxAge(time.Duration(l.cfg.Log.LogMaxAge)*time.Second),
				rotatelogs.WithRotationCount(l.cfg.Log.LogRotationCount),
			)
		} else {
			if l.cfg.Log.LogRotationTime == 0 {
				l.cfg.Log.LogRotationTime = 60 * 60 * 24
			}
			fd, _ = rotatelogs.New(
				l.cfg.Log.LogFile+".%Y%m%d",
				rotatelogs.WithLinkName(l.cfg.Log.LogFile),
				rotatelogs.WithMaxAge(time.Duration(l.cfg.Log.LogMaxAge)*time.Second),
				rotatelogs.WithRotationTime(time.Duration(l.cfg.Log.LogRotationTime)*time.Second),
			)
		}

		l.logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: constants.DefaultTimestampFormat})
		l.logger.SetOutput(fd)
	}
}
