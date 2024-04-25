package cache

import (
	"encoding/json"
	"errors"
	"time"

	redsyncs "github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"

	"Hertz_refactored/biz/config"
	"Hertz_refactored/biz/pkg/logging"
)

var (
	redisClient *redis.Pool
	rs          *redsyncs.Redsync
	//锁过期时间
	lockExpiry = 2 * time.Second
	//获取锁失败重试时间间隔
	retryDelay = 500 * time.Millisecond
	//值过期时间  过期时间需要使用int整形
	valueExpire  = int(3600 * time.Second) 
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
	err := redisClient.Close()
	if err != nil {
		return
	}
}

func Exists(key string) bool {
	conn := redisClient.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

//GenerateID

func GenerateID(key string) int64 {
	conn := redisClient.Get()
	defer conn.Close()

	val, err := redis.Int64(conn.Do("INCR", key))
	if err != nil {
		logrus.Info(err)
		return -1
	}
	return val
}

/*
当多个进程或线程需要访问共享资源时，为了避免并发问题，我们通常会使用锁来保证在同一时刻只有一个进程或线程能够访问共享资源。
假设我们有一个电商网站，用户可以在网站上购买商品。当一个商品的库存只剩下最后一件时，可能会有多个用户同时尝试购买这个商品。
当一个用户尝试购买一个商品时，我们首先获取一个分布式锁。只有成功获取到锁的用户才能继续执行购买操作。这样可以确保在同一时刻，只有一个用户能够购买最后一件商品，从而避免了超卖的问题。
*/

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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)
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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)
	// ZADD racer_scores 10 "Norem"
	//其中ZADD命令中 第一个为这个事务的名称 第二个是值(分数) 第三个是键(分数对应的对象)
	_, err := conn.Do("ZADD", "Rank", value, id)
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

// To Get the RankList
func RangeList(key string) ([]string, error) {
	conn := redisClient.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)
	//ZRANGE命令的排序顺序是从低到高的，而ZREVRANGE命令的顺序是从高到低的。	0和-1表示从元素索引为0到最后一个元素
	//ZRANGE racer_scores 0 -1 withscores 这个命令代表的是不但会返回排序后的键 还会返回这个键也就是对象对应的值也就是分数
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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

	reply, err := redis.Bytes(conn.Do(op, key))
	if err != nil {
		return reply, err
	}

	return reply, nil
}

/////////////////////////Hash类型接口///////////////////////////////////////////

func CacheHSet(key, mkey string, value ...interface{}) error {
	conn := redisClient.Get()
	//这是一个延迟执行的匿名函数。它接受一个 redis.Conn 类型的参数 conn，这个参数是 Redis 连接
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)
	//value ...interface{} 表示 value 可以是任意数量、任意类型的参数。在函数内部，你可以像处理切片一样处理 value
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
		_, err = conn.Do("EXPIRE", key, valueExpire)
		if err != nil {
			logrus.Info(err)
			return err
		}
	}
	return nil
}

func CacheHGet(key, mkey string) ([]byte, error) {
	conn := redisClient.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

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
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

	data, err := redis.String(conn.Do("HGET", key, mkey))
	//fmt.Printf("data:%v", data)
	if err != nil {
		logrus.Info(err)
	}
	return data, nil
}
func CacheDelHash(key, mkey string) error {

	conn := redisClient.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

	_, err := conn.Do("HDEL", key, mkey)
	if err != nil {
		return err
	}
	return nil
}

func CacheDelHash2(key, mkey, comment_id string) error {

	conn := redisClient.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(conn)

	_, err := conn.Do("HDEL", key, mkey, comment_id)
	if err != nil {
		return err
	}
	return nil
}
