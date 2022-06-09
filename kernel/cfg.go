package kernel

import (
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

// CfgSys 配置-系统
type CfgSys struct {
	Mode           string `json:"mode"`             // 模式: debug test release
	CmdHistoryFile string `json:"cmd_history_file"` // 命令历史文件
	CmdLinePrompt  string `json:"cmd_line_prompt"`  // 命令行提示符
	EnableLog      int    `json:"enable_log"`       // 是否启用日志
	EnableDbLog    int    `json:"enable_db_log"`    // 是否启用数据库查询及结果日志
}

// CfgApp 配置-应用
type CfgApp struct {
	Name        string `json:"name"`        // 应用名称
	Version     string `json:"version"`     // 应用版本
	Author      string `json:"author"`      // 应用作者
	Email       string `json:"email"`       // 作者邮箱
	Usage       string `json:"usage"`       // 应用描述
	Description string `json:"description"` // 应用描述
}

// CfgLog 配置-日志
type CfgLog struct {
	LogFile   string `json:"log_file"`    // 日志文件
	DbLogFile string `json:"db_log_file"` // 数据库查询及结果日志
}

// CfgDB 配置-数据库
type CfgDB struct {
	Name   string `json:"name"`   // 名称
	Driver string `json:"driver"` // 数据库驱动
	Dsn    string `json:"dsn"`    // 连接
}

type Cfg struct {
	Sys CfgSys  `json:"sys"`
	App CfgApp  `json:"app"`
	Log CfgLog  `json:"log"`
	DB  []CfgDB `json:"db"`
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

func (cfg Cfg) getCfgFile() string {
	return DefaultCfgFile
}

func (cfg *Cfg) loadDefaultCfg() {
	cfg.Sys = CfgSys{
		Mode:           SysModeDebug,
		CmdHistoryFile: DefaultCmdHistoryFile,
		CmdLinePrompt:  DefaultCmdLinePrompt,
		EnableLog:      1,
		EnableDbLog:    1,
	}
	cfg.App = CfgApp{
		Name:        AppName,
		Version:     AppVersion,
		Author:      AppAuthor,
		Email:       AppEmail,
		Usage:       AppUsage,
		Description: AppDescription,
	}
	cfg.Log = CfgLog{
		LogFile:   DefaultLogFile,
		DbLogFile: DefaultDbLogFile,
	}
	cfg.DB = []CfgDB{{
		Name:   DbDriverMysql,
		Driver: DbDriverMysql,
		Dsn:    DefaultDbDsn,
	}}
}
