package kernel

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

	DefaultCmdHistoryFile string = "./cmd.log"
	DefaultCmdLinePrompt  string = " cqu > "

	DefaultLogFile   string = "./log/app.log"
	DefaultDbLogFile string = "./log/db.log"

	DbDriverMysql string = "mysql"
	DbDriverPgsql string = "pgsql"
	DefaultDbDsn  string = "root:root@tcp(127.0.0.1:3306)/console.query.uniform?charset=utf8mb4"

	DefaultCfgFile string = "./console.json"
)
