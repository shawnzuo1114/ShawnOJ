package database

import (
	"ShawnOJ/utils/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Db *sqlx.DB

func InitMysql(cfg *setting.MysqlConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return err
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	Db = db // 显式赋值给全局变量
	zap.L().Info("MySQL connection established successfully")
	return nil
}

func CloseMysql() {
	_ = Db.Close()
}
