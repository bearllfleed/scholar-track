package initialize

import (
	"time"

	"github.com/bearllflee/scholar-track/pkg/snowflake"
	"github.com/bearllflee/scholar-track/rpc/achieve/internal/global"
)

func MustInitSnowflake(workerID int64) *snowflake.Snowflake {
	epoch := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
	snowflake, err := snowflake.NewSnowflake(epoch, workerID)
	if err != nil {
		panic(err)
	}
	global.Snowflake = snowflake
	return snowflake
}
