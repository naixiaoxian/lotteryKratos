package algorithm

import (
	"fmt"
	"lotteryKratos/internal/data"
	"lotteryKratos/internal/data/strategy/vo"
)

//抽奖
type Impl interface {
	// 初始化策略概率
	/**
	* 程序启动时初始化概率元祖，在初始化完成后使用过程中不允许修改元祖数据
	* <p>
	* 元祖数据作用在于讲百分比内(0.2、0.3、0.5)的数据，转换为一整条数组上分区数据，如下；
	* 0.2 = 0 ~ 0.2
	* 0.3 = 0 + 0.2 ~ 0.2 + 0.3 = 0.2 ~ 0.5
	* 0.5 = 0.5 ~ 1 （计算方式同上）
	* <p>
	* 通过数据拆分为整条后，再根据0-100中各个区间的奖品信息，使用斐波那契散列计算出索引位置，把奖品数据存放到元祖中。比如：
	* <p>
	* 1. 把 0.2 转换为 20
	* 2. 20 对应的斐波那契值哈希值：（20 * HASH_INCREMENT + HASH_INCREMENT）= -1549107828 HASH_INCREMENT = 0x61c88647
	* 3. 再通过哈希值计算索引位置：hashCode & (rateTuple.length - 1) = 12
	* 4. 那么tup[14] = 0.2 中奖概率对应的奖品
	* 5. 当后续通过随机数获取到1-100的值后，可以直接定位到对应的奖品信息，通过这样的方式把轮训算奖的时间复杂度从O(n) 降低到 0(1)
		这个地方需要学习一下看看。通过算法进行了什么样子的优化
	*/

	initRateTuple(strategyId int64, strategyMode int, awardRateInfoList []vo.AwardRateVo)
	// 判断对应策略数据是否存在
	isExist(strategyId int64) bool
	// 生成 随机数据 返回对应奖品返回结果
	// 这里用组合的办法去做对应的实现
	// RandomDraw(strategyId int64, excludeAwardIds []string) string
}

type DrawImpl interface {
	RandomDraw(strategyId int64, excludeAwardIds []string) string
	IsExist(int2 int64) bool
	InitRateTuple(strategyId int64, strategyMode int, awardRateInfoList []vo.AwardRateVo)
}

const (
	HASH_INCREMENT    = 0x61c88647
	RATE_TUPLE_LENGTH = 128
)

type BaseAlgorithm struct {
	//数据可优化至redis
	RateTupleMap     map[int64][]string
	AwardRateInfoMap map[int64][]vo.AwardRateVo
}

func (ba *BaseAlgorithm) initRateTuple(strategyId int64, strategyMode int, awardRateInfoList []vo.AwardRateVo) {
	if ba.isExist(strategyId) {
		return
	}
	//保存概率信息
	ba.AwardRateInfoMap[strategyId] = awardRateInfoList

	//非单向概率，不必存入缓存，因为这部分抽奖算法需要实时处理中奖概率
	if strategyMode == data.SINGLE {
		return
	}

	if _, err := ba.RateTupleMap[strategyId]; err == false {
		ba.RateTupleMap[strategyId] = make([]string, RATE_TUPLE_LENGTH)
	}
	rateTuple, _ := ba.RateTupleMap[strategyId]

	cusorVal := 0
	for _, awardRateVo := range awardRateInfoList {
		rateVal := int(awardRateVo.GetAwardRate() * 100)
		//循环填充概率值
		for i := cusorVal + 1; i <= (rateVal + cusorVal); i++ {
			rateTuple[ba.hashIdx(i)] = awardRateVo.AwardId
		}
		cusorVal += rateVal
	}

	fmt.Println(ba.RateTupleMap[strategyId])
	return
}

func (ba *BaseAlgorithm) hashIdx(val int) int {
	hashCode := val*HASH_INCREMENT + HASH_INCREMENT
	return hashCode & (RATE_TUPLE_LENGTH - 1)
}

func (ba *BaseAlgorithm) isExist(strategyId int64) (ret bool) {
	_, ret = ba.AwardRateInfoMap[strategyId]
	return
}

func IsExistInArray(value string, arrays []string) (ret bool) {
	ret = false
	for _, v := range arrays {
		if v == value {
			ret = true
			return
		}
	}
	return
}
