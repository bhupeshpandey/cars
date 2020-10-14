package cache

import (
	model "github.com/bhupeshpandey/cars/model"
)

func New(cacheConfig *model.CacheConfig) model.Cache {
	return NewLRU(cacheConfig.Size)
}


//func New()
