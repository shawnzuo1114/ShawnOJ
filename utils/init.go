package utils

import (
	"ShawnOJ/utils/database"
	"ShawnOJ/utils/logger"
	"ShawnOJ/utils/setting"
	"ShawnOJ/utils/snowflake"
	"fmt"

	"go.uber.org/zap"
)

func UtilsInit() (err error) {
	//加载配置
	if err = setting.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	//初始化日志
	if err = logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")

	//初始化snowflake
	if err = snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	//初始化MySQL连接
	if err = database.InitMysql(setting.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	//初始化redis连接
	if err = database.InitRedis(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	return nil
}
