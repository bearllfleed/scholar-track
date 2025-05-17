package snowflake

import (
	"fmt"
	"sync"
	"time"
)

type Snowflake struct {
	mu            sync.Mutex // 互斥锁，确保并发安全
	epoch         int64      // 起始时间戳
	workerID      int64      // 工作节点ID
	sequence      int64      // 当前毫秒内的序列号
	lastTimestamp int64      // 上次生成ID的时间戳
}

const (
	timestampBits = 41 // 时间戳占用41位
	workerIDBits  = 10 // 工作节点ID占用10位
	sequenceBits  = 12 // 序列号占用12位

	maxWorkerID   = -1 ^ (-1 << workerIDBits) // 最大工作节点ID: 1023
	maxSequence   = -1 ^ (-1 << sequenceBits) // 最大序列号: 4095
	timeShift     = workerIDBits + sequenceBits
	workerIDShift = sequenceBits                // 工作节点ID左移位数
	timestampMask = sequenceBits + workerIDBits // 时间戳左移位数
)

func NewSnowflake(epoch, workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", maxWorkerID)
	}
	return &Snowflake{
		epoch:         epoch,
		workerID:      workerID,
		lastTimestamp: -1,
		sequence:      0,
	}, nil
}

func (s *Snowflake) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixNano() / 1e6

	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards, refusing to generate ID")
	}
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = timestamp
	id := ((timestamp - s.epoch) << timeShift) | (s.workerID << workerIDShift) | s.sequence
	return id, nil
}

// func GenerateID(workerID int64) (int64, error) {
// 	epoch := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
// 	sf, err := NewSnowflake(epoch, workerID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	id, err := sf.NextID()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }
