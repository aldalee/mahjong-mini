package mahjong

// 麻将包
import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("牌库启动")
}

const ALLCARDNUMBER int = 136

// CARD
// [ 1- 9]壹万～玖万的个数
// [11-19]壹筒～玖筒的个数
// [21-29]壹条～玖条的个数
// [31-37]东南西北中发白的个数
var CARD = []int{
	1, 1, 1, 1,
	2, 2, 2, 2,
	3, 3, 3, 3,
	4, 4, 4, 4,
	5, 5, 5, 5,
	6, 6, 6, 6,
	7, 7, 7, 7,
	8, 8, 8, 8,
	9, 9, 9, 9,
	11, 11, 11, 11,
	12, 12, 12, 12,
	13, 13, 13, 13,
	14, 14, 14, 14,
	15, 15, 15, 15,
	16, 16, 16, 16,
	17, 17, 17, 17,
	18, 18, 18, 18,
	19, 19, 19, 19,
	21, 21, 21, 21,
	22, 22, 22, 22,
	23, 23, 23, 23,
	24, 24, 24, 24,
	25, 25, 25, 25,
	26, 26, 26, 26,
	27, 27, 27, 27,
	28, 28, 28, 28,
	29, 29, 29, 29,
	31, 31, 31, 31,
	32, 32, 32, 32,
	33, 33, 33, 33,
	34, 34, 34, 34,
	35, 35, 35, 35,
	36, 36, 36, 36,
	37, 37, 37, 37,
}

var instance = &Mahjong{}

func GetSingleton() *Mahjong {
	instance.Sum_Number = ALLCARDNUMBER
	instance.Card = make([]int, instance.Sum_Number)
	instance.Out_Number = 0
	instance.OutCardArr = make([]int, instance.Out_Number)
	copy(instance.Card, CARD)
	return instance
}

type Mahjong struct {
	Sum_Number int
	Card       []int
	Out_Number int
	OutCardArr []int
}

func (m *Mahjong) RandomCard() int {
	card_rand := rand.Intn(m.Sum_Number)
	m.Sum_Number--
	fmt.Println("抽到的牌为:", m.Card[card_rand], ", 剩余总数量:", m.Sum_Number)
	ret := m.Card[card_rand]
	m.Card = append(m.Card[:card_rand], m.Card[card_rand+1:]...)

	return ret
}
