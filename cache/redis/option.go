/**
 * @Author: kwens
 * @Date: 2022-11-15 16:53:02
 * @Description:
 */
package redis

type options struct {
	redisConfig Config
}

type Option interface {
	apply(opt *options)
}

type redisConfigOption struct {
	redisConfig Config
}

func (r redisConfigOption) apply(opt *options) {
	opt.redisConfig = r.redisConfig
}

func WithRedisConfig(redisConf Config) Option {
	return &redisConfigOption{redisConf}
}
