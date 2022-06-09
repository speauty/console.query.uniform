package log

import (
	cfg2 "console.query.uniform/kernel/cfg"
	"console.query.uniform/kernel/constants"
	"console.query.uniform/util"
	"fmt"
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

// Log 日志采用logrus实现，参考链接 https://github.com/sirupsen/logrus
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
	// 报告调用者文件及行数，似乎没啥用，没有trace
	//l.logger.SetReportCaller(true)
	l.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: constants.DefaultTimestampFormat, DisableColors: false,
		ForceColors: true, FullTimestamp: true})
	if l.cfg.Log.LogFile != "" { // 如果日志文件非空, 将日志打到对应文件
		var fd *rotatelogs.RotateLogs
		var err error

		optLog := l.cfg.Log.LogFile + ".%Y%m%d"
		optLink := rotatelogs.WithLinkName(l.cfg.Log.LogFile)
		optMaxAge := rotatelogs.WithMaxAge(time.Duration(l.cfg.Log.LogMaxAge) * time.Second)
		optMaxRotationCount := rotatelogs.WithRotationCount(l.cfg.Log.LogRotationCount)
		if l.cfg.Log.LogRotationTime == 0 {
			l.cfg.Log.LogRotationTime = 60 * 60 * 24
		}
		optRotationTime := rotatelogs.WithRotationTime(time.Duration(l.cfg.Log.LogRotationTime) * time.Second)
		// windows默认环境，没有创建软连接的权限，所以这里需要区分处理
		// 参考链接 https://github.com/golang/go/issues/22874
		if util.GetOS() == constants.SysOsWindows {
			if l.cfg.Log.LogRotationCount > 0 {
				fd, err = rotatelogs.New(optLog, optMaxAge, optMaxRotationCount)
			} else {
				fd, err = rotatelogs.New(optLog, optMaxAge, optRotationTime)
			}
		} else {
			if l.cfg.Log.LogRotationCount > 0 {
				fd, err = rotatelogs.New(optLog, optLink, optMaxAge, optMaxRotationCount)
			} else {
				fd, err = rotatelogs.New(optLog, optLink, optMaxAge, optRotationTime)
			}
		}

		if err != nil {
			fmt.Println(err)
		}

		l.logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: constants.DefaultTimestampFormat})
		l.logger.SetOutput(fd)
	}
}
