package models

//文档 https://github.com/redis/go-redis
import (
	"context"
	"time"
)

type RedisStore struct{}

// RedisStore 实现Store接口
var ctx = context.Background()

// CAPTCHA 验证码前缀
var CAPTCHA = "captcha:"

// Set 实现设置captcha的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := RedisDb.Set(ctx, key, value, time.Minute*2).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get 实现获取captcha的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	// 如果clear为true，删除验证码
	if clear {
		err := RedisDb.Del(ctx, key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
