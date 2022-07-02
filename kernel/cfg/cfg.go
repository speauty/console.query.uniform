package cfg

import (
	"console.query.uniform/kernel/constants"
	"encoding/json"
	"io/ioutil"
	"sync"
)

var (
	cfgService *Cfg
	cfgOnce    sync.Once
)

func NewCfgService() *Cfg {
	cfgOnce.Do(func() {
		cfgService = &Cfg{}
		err := cfgService.LoadCfg()
		if err != nil {
			cfgService.loadDefaultCfg()
			_ = cfgService.Flush()
		}
	})
	return cfgService
}

// Sys 配置-系统
type Sys struct {
	Mode           string `json:"mode"`             // 模式: debug test release
	CmdHistoryFile string `json:"cmd_history_file"` // 命令历史文件
	CmdLinePrompt  string `json:"cmd_line_prompt"`  // 命令行提示符
	EnableLog      int    `json:"enable_log"`       // 是否启用日志
	EnableDbLog    int    `json:"enable_db_log"`    // 是否启用数据库查询及结果日志
}

func (s Sys) Default() *Sys {
	return &Sys{
		Mode:           constants.SysModeDebug,
		CmdHistoryFile: constants.DefaultCmdHistoryFile,
		CmdLinePrompt:  constants.DefaultCmdLinePrompt,
		EnableLog:      1,
		EnableDbLog:    1,
	}
}

// App 配置-应用
type App struct {
	Name        string `json:"name"`        // 应用名称
	Version     string `json:"version"`     // 应用版本
	Description string `json:"description"` // 应用描述
}

func (a App) Default() *App {
	return &App{
		Name:        constants.AppName,
		Version:     constants.AppVersion,
		Description: constants.AppDescription,
	}
}

// Log 配置-日志
type Log struct {
	DbLogFile        string `json:"db_log_file"`        // 数据库查询及结果日志
	LogFile          string `json:"log_file"`           // 日志文件
	LogLevel         uint32 `json:"log_level"`          // 日志级别
	LogRotationTime  int    `json:"log_rotation_time"`  // 日志分割时间
	LogRotationCount uint   `json:"log_rotation_count"` // 日志文件最大保存数量(和LogMaxAge只能设置一个, 优先采用LogRotationCount)
	LogMaxAge        int    `json:"log_max_age"`        // 日志文件清理前最长保存时间
}

func (l Log) Default() *Log {
	return &Log{
		DbLogFile:        constants.DefaultDbLogFile,
		LogFile:          constants.DefaultLogFile,
		LogLevel:         constants.DefaultLogLevel,
		LogRotationTime:  constants.DefaultLogRotationTime,
		LogRotationCount: constants.DefaultLogRotationCount,
		LogMaxAge:        constants.DefaultLogMaxAge,
	}
}

// Db 配置-数据库
type Db struct {
	Name   string `json:"name"`   // 名称
	Driver string `json:"driver"` // 数据库驱动
	Dsn    string `json:"dsn"`    // 连接
}

func (d Db) Default() []*Db {
	return []*Db{{
		Name:   constants.DbDriverMysql,
		Driver: constants.DbDriverMysql,
		Dsn:    constants.DefaultDbDsn,
	}}
}

type Cfg struct {
	Sys *Sys  `json:"sys"`
	App *App  `json:"app"`
	Log *Log  `json:"log"`
	Db  []*Db `json:"db"`
}

// LoadCfg 从执行文件加载配置
func (cfg *Cfg) LoadCfg() error {
	data, err := ioutil.ReadFile(cfg.getCfgFile())
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, cfg); err != nil {
		return err
	}
	return nil
}

// Flush 将当前配置刷入指定文件, 然后重新载入配置
func (cfg *Cfg) Flush() error {
	data, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(cfg.getCfgFile(), data, 0666); err != nil {
		return err
	}

	if err = cfg.LoadCfg(); err != nil {
		return err
	}

	return nil
}

// 获取配置文件, 这里可能要做调整，如果配置文件位置变了，这里就要适当调整一下，这也是为什么单独使用getter获取的原因
func (cfg *Cfg) getCfgFile() string { return constants.DefaultCfgFile }

// 载入默认配置，分别调用配置结构体自带的Default生成相应的默认配置
func (cfg *Cfg) loadDefaultCfg() {
	cfg.Sys = Sys{}.Default()
	cfg.App = App{}.Default()
	cfg.Log = Log{}.Default()
	cfg.Db = Db{}.Default()
}
