package mycard

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("我的卡牌")
}

const (
	MAXCARD     int = 38
	MAXHANDCARD int = 18
)

var instance = &MyCard{}

func GetSingleton() *MyCard {
	instance.Card_Number = 0
	instance.Card_Arr = make([]int, MAXHANDCARD)
	return instance
}

// MyCard
// [ 1- 9]壹万～玖万的个数
// [11-19]壹筒～玖筒的个数
// [21-29]壹条～玖条的个数
// [31-37]东南西北中发白的个数
type MyCard struct {
	Card_Number     int   // 注意***所有边界都由 card_Number 来控制
	Card_Arr        []int // 手牌
	Card_Quaternion []int // 杠牌
	Universal_Card  int   // 万能牌
}

// SetCard 插入麻将
func (this *MyCard) SetCard(card int) {
	this.Card_Arr[this.Card_Number] = card
	this.Card_Number++
	fmt.Print("摸到的麻将：")
	this.CatCard(card)
	// 使用浅拷贝，将赋值过的数组进行排序
	s := this.Card_Arr[:this.Card_Number]
	sort.Ints(s)
}

func (this *MyCard) CatCard(value int) {
	switch value {
	case 31:
		fmt.Printf("东")
	case 32:
		fmt.Printf("南")
	case 33:
		fmt.Printf("西")
	case 34:
		fmt.Printf("北")
	case 35:
		fmt.Printf("中")
	case 36:
		fmt.Printf("发")
	case 37:
		fmt.Printf("白")
	default:
		if value < 10 {
			fmt.Printf(" %d万", value)
		} else if value < 20 {
			fmt.Printf(" %d筒", value%10)
		} else if value < 30 {
			fmt.Printf(" %d条", value%20)
		}
	}
}

func (this *MyCard) CatAllCard() {
	fmt.Printf("我的牌组：")
	for i := 0; i < this.Card_Number; i++ {
		value := this.Card_Arr[i]
		this.CatCard(value)
	}
	fmt.Println()
	fmt.Print("我的牌号：")
	for i := 1; i <= this.Card_Number; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Print("\n我的杠牌: ")
	for _, value := range this.Card_Quaternion {
		this.CatCard(value)
	}
	fmt.Println()
}

// RemoveQuaternion 找出杠
func (this *MyCard) RemoveQuaternion() bool {
	for i := 0; i < this.Card_Number-3; i++ {
		if this.Card_Arr[i] == this.Card_Arr[i+3] {
			fmt.Print("你有一张杠:")
			this.CatCard(this.Card_Arr[i])
			fmt.Print("\n是否杠?\n输入0不杠,输入1杠;\n请输入:")
			var outCard int
			for {
				fmt.Scan(&outCard)
				if outCard == 0 || outCard == 1 {
					break
				}
				fmt.Print("输入错误，请重新输入:")
			}
			if outCard == 0 {
				i += 3
				continue
			} else if outCard == 1 {
				this.Card_Quaternion = append(this.Card_Quaternion, this.Card_Arr[i])
				this.Card_Arr[i] = 38
				this.Card_Arr[i+1] = 38
				this.Card_Arr[i+2] = 38
				this.Card_Arr[i+3] = 38
				// 将4对排序后移，用card_Number 作为边界
				s := this.Card_Arr[:this.Card_Number]
				sort.Ints(s)
				this.Card_Number = this.Card_Number - 4
				return true
			}
		}
	}
	return false
}

func (this *MyCard) OutCard(set int) int {
	var out int

	fmt.Printf("打出的牌为：")
	this.CatCard(this.Card_Arr[set])
	fmt.Println()
	out = this.Card_Arr[set]

	// 将打出的牌改为最大，再将手牌数量减少
	this.Card_Arr[set] = this.Card_Arr[this.Card_Number-1]
	this.Card_Number--
	// 使用浅拷贝，将赋值过的数组进行排序
	s := this.Card_Arr[:this.Card_Number]
	sort.Ints(s)
	fmt.Println()
	return out
}
