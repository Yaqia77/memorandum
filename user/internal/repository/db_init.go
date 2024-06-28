package repository

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset + "&parseTime=True&loc=Local"}, "")
	err := Database(dsn)
	if err != nil {
		panic(err)
	}
}

func Database(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == gin.DebugMode {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true, //禁用datetime精度，MySQL5.6不支持
		DontSupportRenameIndex:   true, //重命名索引的时候采用删除重建的方式
		DontSupportRenameColumn:  true, //重命名列时不支持
	}), &gorm.Config{
		Logger: ormLogger,
		// 禁用默认表名复数形式，启用单数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加复数
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)                 //最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) //连接最大存活时间
	DB = db
	return nil
}
