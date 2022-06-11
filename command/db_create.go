package command

//@todo 等待实现
// 结构体:DbCreate 指令:db.create
// 注意: 需要保证配置是最新的, 因此注册Before和After函数句柄, 主要是为了重载配置
// NewCfgService().Db.append() 后调用Cfg.Flush写入文件
