package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	// 如果是debug模式
	if gin.Mode() == "debug" {
		// 使用info级别的日志记录
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 否则默认
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, //路径
		DefaultStringSize:         256,      //string类型字段默认长度
		DisableDatetimePrecision:  true,     // mysql 5.6之前的数据库不支持datetime精度
		DontSupportRenameIndex:    true,     //不允许重命名索引，得先删掉再重建索引
		DontSupportRenameColumn:   true,     //用change重命名列，mysql 8之前不支持重命名列
		SkipInitializeWithVersion: false,    //如果为true则根据当前mysql版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //创建表格不用加s即不用复数化
		},
	})
	if err != nil {
		fmt.Printf("Database connection failed,error:%v\n", err)
		return
	}
	mydb, _ := db.DB()
	mydb.SetMaxIdleConns(20)                  //设置连接池，最大连接为20
	mydb.SetMaxOpenConns(100)                 //打开连接数
	mydb.SetConnMaxLifetime(time.Second * 30) // 连接的生存时间
	_db = db

	// 主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      //写操作
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, //读操作
		Policy:   dbresolver.RandomPolicy{},                                    //读写平衡策略
	}))
	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
