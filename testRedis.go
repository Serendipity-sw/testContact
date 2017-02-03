package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/guotie/deferinit"
	"time"
	"fmt"
	"strconv"
)

func main() {

	timeNow:=time.Now()
	offTime:=timeNow.AddDate(0,0,1)
	baseTime:=time.Date(offTime.Year(),offTime.Month(),offTime.Day(),0,0,0,0,time.Local)
	timeTickf:=baseTime.Sub(time.Now())
	scenepotal:=timeTickf/time.Hour
	fmt.Println(int(scenepotal))



	deferinit.InitAll()
	defer deferinit.FiniAll()

	err:=setRedisCachePs(22,30)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jTmr:=time.NewTimer(1*time.Second)
	<-jTmr.C
	for {
		redisCode,err:=getRedisCachePs()
		if err != nil {
			fmt.Println(err.Error())
		}else{
			fmt.Println(redisCode)
		}
		jTmr.Reset(1*time.Second)
		<-jTmr.C
	}

}

var (
	rpool           *redis.Pool
	cacheBssSeconds int = 86400
)

func init() {
	deferinit.AddInit(connectRedis, disconnectRedis, 500) // 优先级最高
	deferinit.AddInit(func() {
		cacheBssSeconds = 86400
	}, nil, 50)
}

func disconnectRedis() {
	closeRedis()
}

func connectRedis() {
	proto := "tcp"
	addr := "10.10.188.10:6379"
	dbindex := 5
	openRedis(proto, addr, dbindex)

	fmt.Println("open redis successfully.")
}

func openRedis(proto, addr string, idx int) {
	rpool = &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 600 * time.Second,
		Dial: func() (redis.Conn, error) {
			do := redis.DialDatabase(idx)
			c, err := redis.Dial(proto, addr, do)
			if err != nil {
				panic(err)
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	conn := rpool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		panic(err.Error())
	}

	return
}

func closeRedis() {
	rpool.Close()
}

func setRedisCachePs(ps int,number int) error {

	c := rpool.Get()
	defer c.Close()
	_, err := c.Do("SETEX", "smsMob",number,ps)
	if err != nil {
		fmt.Sprintf("setRedisCachePs set redis cache failed: string %s error=%v\n",
			ps, err)
	}
	return err
}

func getRedisCachePs() (int, error) {

	c := rpool.Get()
	r, err := c.Do("GET", "smsMob")
	c.Close()
	if err != nil {
		return 0, err
	}
	if r == nil {
		err = fmt.Errorf("getRedisCache: key %s is nil", "smsControl")
		return 0, err
	}
	ps := string(r.([]byte))
	indexNumber,err:=strconv.Atoi(ps)
	if err != nil {
		return 0,err
	}
	return indexNumber, nil
}