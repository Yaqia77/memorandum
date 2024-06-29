package dao

import (
	"os"

	"github.com/yaqia77/memorandum/apps/user/internal/repository/db/model"
	"github.com/yaqia77/memorandum/pkg/util/logger"
)

func migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
		)
	if err != nil {
		logger.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}
	logger.LogrusObj.Infoln("register table success")
}
