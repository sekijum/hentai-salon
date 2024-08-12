package util

import (
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

type SonyflakeIDGenerator struct {
	flake *sonyflake.Sonyflake
	mu    sync.Mutex
}

func NewSonyflakeIDGenerator() *SonyflakeIDGenerator {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(), // Snowflake IDの生成を開始する時間を設定
	})
	if flake == nil {
		panic("Failed to initialize Sonyflake ID generator")
	}

	return &SonyflakeIDGenerator{
		flake: flake,
	}
}

func (gen *SonyflakeIDGenerator) GenerateID() (uint64, error) {
	gen.mu.Lock()
	defer gen.mu.Unlock()

	id, err := gen.flake.NextID()
	if err != nil {
		return 0, err
	}
	return id, nil
}