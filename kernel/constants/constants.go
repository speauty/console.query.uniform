package constants

const (
	AppName        string = "console.query.uniform"
	AppVersion     string = "v0.0.0"
	AppAuthor      string = "speauty"
	AppEmail       string = "speauty@163.com"
	AppUsage       string = "-"
	AppDescription string = "this is a description for the application"

	SysModeDebug   string = "debug"
	SysModeTest    string = "test"
	SysModeRelease string = "release"

	DefaultTimestampFormat string = "2006-01-02 15:04:05"

	DefaultCmdHistoryFile string = "./runtime/log/cliApp.log"
	DefaultCmdLinePrompt  string = "query>"

	DefaultLogFile          string = "./runtime/log/app.log"
	DefaultDbLogFile        string = "./runtime/log/db.log"
	DefaultLogLevel         uint32 = 6
	DefaultLogRotationTime  int    = 86400
	DefaultLogRotationCount uint   = 15
	DefaultLogMaxAge        int    = 0

	DbDriverMysql string = "mysql"
	DbDriverPgsql string = "pgsql"
	DefaultDbDsn  string = "root:root@tcp(127.0.0.1:3306)/console.query.uniform?charset=utf8mb4"

	DefaultCfgFile string = "./console.json"
)
