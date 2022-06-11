package constants

const (
	SysOsWindows string = "windows"

	SysModeDebug   string = "debug"
	SysModeTest    string = "test"
	SysModeRelease string = "release"

	AppName        string = "统一查询终端"
	AppVersion     string = "v0.0.0"
	AppDescription string = "this is a description for the application"

	DefaultTimestampFormat string = "2006-01-02 15:04:05"

	DefaultCmdHistoryFile string = "./runtime/log/history.log"
	DefaultCmdLinePrompt  string = "query>"

	DefaultLogFile          string = "./runtime/log/app.log"
	DefaultDbLogFile        string = "./runtime/log/db.log"
	DefaultLogLevel         uint32 = 6
	DefaultLogRotationTime  int    = 86400
	DefaultLogRotationCount uint   = 15
	DefaultLogMaxAge        int    = 0

	DbDriverMysql string = "mysql"
	DbDriverPgsql string = "pgsql"
	DefaultDbDsn  string = "root:root@tcp(127.0.0.1:3306)/console.query.uniform?charset=utf8mb4&multiStatements=true"

	DefaultCfgFile string = "./console.json"
)
