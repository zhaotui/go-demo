package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

//全局变量
var (
	Body [] string //总容器
	spa [] string = make([]string, 13) //黑桃
	rpea [] string = make([]string, 13) //红桃
	squ [] string = make([]string, 13) //方块
	plum [] string = make([]string, 13) //梅花
)

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
//分配花色牌
func flower(c string) {
	r := generateRandomNumber(1,14,13)
	switch (c) {
	case "spa":
		for k, v := range r {
		    spa[k] = "黑桃" + strconv.Itoa(v)
		}
	break
	case "rpea":
		for k, v := range r {
		    rpea[k] = "红桃" + strconv.Itoa(v)
		}
	break
	case "squ":
		for k, v := range r {
		    squ[k] = "方块" + strconv.Itoa(v)
		}
	break
	case "plum":
		for k, v := range r {
		    plum[k] = "梅花" + strconv.Itoa(v)
		}
	break
	}
}

//洗牌
func Shuffle () {
	r := generateRandomNumber(0,52,52)
	copyb := Body
	copy(Body,copyb)
	for k, v := range r {
		Body[k] = copyb[v]
	}
}

func main () {
	flower("spa")
	flower("rpea")
	flower("squ")
	flower("plum")
	Body = append(Body, spa...)
	Body = append(Body, rpea...)
	Body = append(Body, squ...)
	Body = append(Body, plum...)
	Shuffle()
	fmt.Println(Body)
}