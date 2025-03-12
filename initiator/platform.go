package initiator

import (
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/platform"
	"github.com/BoruTamena/jobboard/platform/rediscache"
)

type Platform struct {
	redis platform.RedisCache
}

func InitPlatform(cfg dto.Config) Platform {

	return Platform{
		redis: rediscache.InitRedisCache(cfg),
	}

}
