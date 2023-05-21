package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var scoreMap = make(map[string]int)
	scoreMap["王子"] = 100
	scoreMap["王后"] = 200
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["王子"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var stuMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		stuMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range stuMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, stuMap[key])
	}
}
