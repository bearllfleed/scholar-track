package initialize

import (
	"time"

	"github.com/bearllfleed/scholar-track/pkg/snowflake"
)

func MustInitSnowflake(workerID int64) *snowflake.Snowflake {
	epoch := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
	snowflake, err := snowflake.NewSnowflake(epoch, workerID)
	if err != nil {
		panic(err)
	}
	return snowflake
}
