package cache

import (
	"Hertz_refactored/biz/config"
	"Hertz_refactored/biz/pkg/logging"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"time"

	redsyncs "github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
)

var (
	redisClient *redis.Pool
	rs          *redsyncs.Redsync
	//锁过期时间
	lockExpiry = 2 * time.Second
	//获取锁失败重试时间间隔
	retryDelay = 500 * time.Millisecond
	//值过期时间
	valueExpire  = 3600 * 24 * 30
	ErrMissCache = errors.New("miss Cache")
	//锁设置
	option = []redsyncs.Option{
		redsyncs.WithExpiry(lockExpiry),
		redsyncs.WithRetryDelay(retryDelay),
	}
)

func Init() {
	network := "tcp"
	auth := ""
	redisClient = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, config.ConfigInfo.Redis.Addr,
				redis.DialPassword(auth),
				redis.DialDatabase(1),
			)
			if err != nil {
				logrus.Error("conn redis failed,", err)
				return nil, err
			}

			return c, err
		},
	}
	//这条指令每次会把redis的缓存清空
	//redisClient.Get().Do("flushdb")
	sync := redigo.NewPool(redisClient)
	rs = redsyncs.New(sync)
	logrus.Info("redis conn success")

}

func GetRedis() redis.Conn {
	return redisClient.Get()
}
func CloseRedis() {
	redisClient.Close()
}

func Exists(key string) bool {
	conn := redisClient.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func GetLock(key string) (*redsyncs.Mutex, error) {
	mutex := rs.NewMutex(key+"_lock", option...)
	if err := mutex.Lock(); err != nil {
		return mutex, err
	}
	return mutex, nil
}

func UnLock(mutex *redsyncs.Mutex) error {
	if _, err := mutex.Unlock(); err != nil {
		return err
	}
	return nil
}

/////////////////////////String类型接口////////////////////////////////////////

func CacheSet(key string, data interface{}) error {
	//这个函数里面有两个参数,一个是key作为缓存项的键,另一个为dat,空接口,表示为它可以接受任何类型的数据
	conn := redisClient.Get()
	//redisClient是一个已经初始化的Redis客户端对象，它有一个Get()方法用于获取一个Redis连接
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value, "EX", valueExpire)
	/*
		使用Redis的Set命令将序列化的数据存入Redis中
		key：用作缓存键。
		value：序列化后的数据。
		"EX"：设置过期时间。
		valueExpire：应该是一个在函数外部定义的变量，表示过期时间（秒）
	*/
	if err != nil {
		logrus.Info("第二步出错")
		return err
	}
	return nil
}

func CacheGet(key string) ([]byte, error) { //用于获取键
	conn := redisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	//这是一个辅助函数（可能是go-redis库提供的），用于将Redis命令的响应转换为字节切片（[]byte）
	if err != nil {
		return nil, err
	}

	//这表示为字节切片，所以使用len表示其长度，以判读这个切片是否为空
	if len(reply) == 0 {
		return nil, ErrMissCache
	}

	return reply, nil
}

// ToDO 为实现排行功能完成
func RangeAdd(value, id int64) error {
	conn := redisClient.Get()
	defer conn.Close()
	_, err := conn.Do("ZADD", "Rank", value, id)
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

func RangeList(key string) ([]string, error) {
	conn := redisClient.Get()
	defer conn.Close()
	res, err := redis.Strings(conn.Do("ZRevRange", key, 0, -1))
	if err != nil {
		logging.Error(err)
	}
	return res, nil
}

///////////////////////////List类型接口////////////////////////////////////////

func CacheLPush(key string, value ...interface{}) error {
	return listPush("LPUSH", key, value)
}

func CacheRPush(key string, value ...interface{}) error {
	return listPush("RPUSH", key, value)
}

func CacheLPop(key string) ([]byte, error) {
	return listPop("LPOP", key)
}

func CacheRPop(key string) ([]byte, error) {
	return listPop("RPOP", key)
}

func CacheLGetAll(key string) ([][]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.ByteSlices(conn.Do("LRANGE", key, "0", "-1"))
	if err != nil {
		return [][]byte{}, err
	}
	return data, nil
}

func listPush(op, key string, data ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()
	for _, d := range data {
		value, err := json.Marshal(d)
		if err != nil {
			return err
		}
		_, err = conn.Do(op, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func listPop(op, key string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do(op, key))
	if err != nil {
		return reply, err
	}

	return reply, nil
}

/////////////////////////Hash类型接口///////////////////////////////////////////

func CacheHSet(key, mkey string, value ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()

	for _, d := range value {
		data, err := json.Marshal(d)
		if err != nil {
			return nil
		}
		//对该redis缓存的解释为:key是哈希表的名字，而mkey则是哈希表的键名
		_, err = conn.Do("HSET", key, mkey, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func CacheHGet(key, mkey string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("HGET", key, mkey))
	//fmt.Printf("data:%v", data)
	if err != nil {
		return []byte{}, err
	}
	if len(data) == 0 {
		return []byte{}, ErrMissCache
	}
	return data, nil
}
func CacheHGet2(key, mkey string) (string, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("HGET", key, mkey))
	//fmt.Printf("data:%v", data)
	if err != nil {
		logrus.Info(err)
	}
	return data, nil
}
func CacheDelHash(key, mkey string) error {

	conn := redisClient.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, mkey)
	if err != nil {
		return err
	}
	return nil
}

func CacheDelHash2(key, mkey, comment_id string) error {

	conn := redisClient.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", key, mkey, comment_id)
	if err != nil {
		return err
	}
	return nil
}
