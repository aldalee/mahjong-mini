package main

import (
	"fmt"
	"mahjong-mini/mahjong"
	"mahjong-mini/mycard"
	"mahjong-mini/rule"
)

type Start struct {
	Card_Lib *mahjong.Mahjong
	My_Card  *mycard.MyCard
}

// InitCard 抽取13张基础卡牌
func (s *Start) InitCard() {
	for i := 0; i < 13; i++ {
		s.SetMyCard()
		fmt.Println()
	}
	fmt.Println()

}

// SetMyCard 摸牌
func (s *Start) SetMyCard() {
	s.My_Card.SetCard(s.Card_Lib.RandomCard())
}

// OutMyCard 打出
func (s *Start) OutMyCard() {
	s.My_Card.CatAllCard()
	var outCard int
	fmt.Printf("您要打出的牌：")
	for {
		fmt.Scan(&outCard)
		if outCard <= s.My_Card.Card_Number && outCard >= 1 {
			break
		} else {
			fmt.Printf("输入错误，请重新输入:")
		}
	}

	out := s.My_Card.OutCard(outCard - 1)
	s.Card_Lib.OutCardArr = append(s.Card_Lib.OutCardArr, out)
	s.Card_Lib.Out_Number++
}

func (s *Start) IsHu() bool {

	sorted := make([]int, s.My_Card.Card_Number)
	sorted = append(s.My_Card.Card_Arr[:s.My_Card.Card_Number])

	// fmt.Println(ss.My_Card.Card_Arr)
	// fmt.Println(sorted)
	// 没有一个对子，没有胡牌
	if rule.FindPairPos(sorted) == nil {
		return false
	}

	if rule.IsShiSanYao(sorted) { // 十三幺
		return true
	} else if rule.IsQiXiaoDui(sorted) { // 小七对
		if rule.IsQingYiSe(sorted) {
			fmt.Println("恭喜你胡牌：", rule.HuTypeQingYiSe, rule.HuTypeQiXiaoDui)
			return true
		}
		fmt.Println("恭喜你胡牌：", rule.HuTypeQiXiaoDui)
		return true
	} else if rule.IsNormal(sorted) { // 普通胡牌
		if rule.IsQingYiSe(sorted) {
			fmt.Println("恭喜你胡牌：", rule.HuTypeQingYiSe, rule.HuTypeNormal)
			return true
		}
		fmt.Println("恭喜你胡牌：", rule.HuTypeNormal)
		return true
	} else if rule.IsDuiDui(sorted) {
		if rule.IsQingYiSe(sorted) {
			fmt.Println("恭喜你胡牌：", rule.HuTypeQingYiSe, rule.HuTypeDuiDui)
			return true
		}
		fmt.Println("恭喜你胡牌：", rule.HuTypeDuiDui)
		return true
	}

	return false
}

func main() {

	start := &Start{}
	start.Card_Lib = mahjong.GetSingleton()
	start.My_Card = mycard.GetSingleton()
	// 初始摸13张
	start.InitCard()

	for i := 0; ; i++ {
		start.SetMyCard()
		fmt.Println()
		if start.My_Card.RemoveQuaternion() {
			continue
		}
		if start.IsHu() {
			start.My_Card.CatAllCard()
			fmt.Println("对局结束")
			break
		}
		start.OutMyCard()
	}
}
