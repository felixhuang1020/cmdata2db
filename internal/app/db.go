package app

import (
	"cmdata2db/config"
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Engine *gorm.DB

// InitializeCk 数据库初始化
func InitializeCk() error {
	var err error
	Engine, err = gorm.Open(clickhouse.New(clickhouse.Config{
		DSN: config.Conf.Database.Source,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "t_", // 表名前缀（所有表名会自动添加 t_）
			SingularTable: true, // 使用单数表名（默认是复数，设置为 true 后 User → user）
		},
	})
	if err != nil {
		log.Error("数据库连接失败: %v", err)
		panic("数据库连接失败")
		return err
	}
	if err = pingDB(Engine); err != nil {
		log.Error("数据库连接失败: %v", err)
		return err
	}
	sqlDB, err := Engine.DB()
	if err != nil {
		log.Error("数据库连接失败: %v", err)
		return err
	}
	sqlDB.SetMaxIdleConns(10)                 // 增加空闲连接数
	sqlDB.SetMaxOpenConns(10)                 // 最大打开连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // 减少连接生命周期
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return nil
}

// 检查数据库连接是否有效
func pingDB(db *gorm.DB) error {
	// 获取底层sql.DB对象
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 执行ping操作
	return sqlDB.PingContext(ctx)
}
