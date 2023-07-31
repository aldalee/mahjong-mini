package rule

import "fmt"

const (
	// 胡牌类型
	HuTypeNormal     = "普通胡牌"
	HuTypeDuiDui     = "对对胡"
	HuTypeQiXiaoDui  = "7小对"
	HuTypeYiTiaoLong = "一条龙"
	HuTypeShiSanYao  = "十三么"
	HuTypeQingYiSe   = "清一色"
)

func init() {
	fmt.Println("")
}

// FindPairPos 找出所有对牌的位置
func FindPairPos(sortedCards []int) []int {
	var pos = []int{}
	length := len(sortedCards) - 1
	for i := 0; i < length; i++ {
		if sortedCards[i] == sortedCards[i+1] {
			pos = append(pos, i)
			i++
		}
	}
	return pos
}

// RemovePair 从已排序的牌中，移除一对
func RemovePair(sortedCards []int, pos int) []int {
	remainCards := make([]int, 0, len(sortedCards)-2)
	remainCards = append(remainCards, sortedCards[:pos]...)
	remainCards = append(remainCards, sortedCards[pos+2:]...)
	return remainCards
}

// RemoveTriplet 剔除所有111刻子
func RemoveTriplet(arr []int) []int {
	length := len(arr)
	var retArr []int
	copy(retArr, arr)
	for i := 0; i < length-2; i++ {
		if arr[i] == arr[i+1] && arr[i] == arr[i+2] {
			retArr = append(retArr, i)
			i += 2
		}
	}

	for i := len(retArr) - 1; i >= 0; i-- {
		arr = append(arr[:retArr[i]], arr[retArr[i]+3:]...)
	}
	return arr
}

// RemoveSequence 判断刻子
func RemoveSequence(arr []int) (int, int) {
	for i := 1; i < len(arr); i++ {
		if arr[0] == arr[i]-1 {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] == arr[j]-1 {
					return i, j
				}
			}
		}
	}
	return 0, 0
}

// IsSequence 判断剩余是否全为顺子
func IsSequence(sortedCards []int) bool {
	length := len(sortedCards)
	arr := make([]int, length)
	copy(arr, sortedCards)
	if length%3 != 0 && arr[length-1] < 30 {
		return false
	}
	for {
		if len(arr) == 0 {
			return true
		}
		i, j := RemoveSequence(arr)
		if i == 0 {
			return false
		}
		arr = append(arr[:j], arr[j+1:]...)
		arr = append(arr[:i], arr[i+1:]...)
		arr = append(arr[1:])
	}
}

// IsNormal
// 普通胡牌规则,只有一个 对牌
// for循环剔除一个对牌
// 检测是否为 123、111、三张牌型
func IsNormal(sortedCards []int) bool {
	length := len(sortedCards)
	pos := FindPairPos(sortedCards)

	for _, value := range pos {
		arr := make([]int, length)
		copy(arr, sortedCards)
		arr = RemovePair(arr, value)
		arr = RemoveTriplet(arr)
		if IsSequence(arr) {
			return true
		}
	}
	return false
}

// IsDuiDui 对对胡
func IsDuiDui(sortedCards []int) bool {
	length := len(sortedCards)
	arr := make([]int, length)
	copy(arr, sortedCards)
	// 处理掉单个,13幺只有一个对
	pos := FindPairPos(arr)

	for _, value := range pos {
		arr := make([]int, length)
		copy(arr, sortedCards)
		// 剔除一个对子
		arr = RemovePair(arr, value)
		arr = RemoveTriplet(arr)
		if len(arr) == 0 {
			return true
		}
	}
	return false
}

// IsShiSanYao 13幺
func IsShiSanYao(sortedCards []int) bool {
	length := len(sortedCards)
	arr := make([]int, length)
	copy(arr, sortedCards)
	// 处理掉单个,13幺只有一个对
	pos := FindPairPos(arr)
	if len(pos) == 0 || length != 14 {
		return false
	}
	arr = append(arr[:pos[0]], arr[pos[0]+1:]...)
	// 验证 万 条 筒
	if !(arr[0] == 1 && arr[1] == 9) {
		return false
	}
	if !(arr[2] == 11 && arr[3] == 19) {
		return false
	}
	if !(arr[4] == 21 && arr[5] == 29) {
		return false
	}
	if arr[6] != 31 {
		return false
	}

	for i := 6; i < 12; i++ {
		if arr[i]+1 != arr[i+1] {
			return false
		}
	}
	fmt.Println("恭喜你胡牌：", HuTypeShiSanYao)
	return true
}

// IsQiXiaoDui 七小对
func IsQiXiaoDui(sortedCards []int) bool {
	length := len(sortedCards)
	if length != 14 {
		return false
	}
	arr := make([]int, length)
	copy(arr, sortedCards)

	pos := FindPairPos(arr)
	if len(pos) == 7 && length == 14 {
		return true
	}
	return false
}

// IsQingYiSe 清一色
func IsQingYiSe(sortedCards []int) bool {
	maxLength := sortedCards[len(sortedCards)-1]
	minLength := sortedCards[0]
	if (maxLength-minLength) < 10 && maxLength < 30 {
		return true
	}
	return false
}
