package global

import (
	"github.com/bearllflee/scholar-track/pkg/snowflake"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Snowflake *snowflake.Snowflake
)
