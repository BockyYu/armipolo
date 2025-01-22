package service

import (
	"armipolo/dto"
	"fmt"
)

type Commodities interface {
	PrintDetail(com ...dto.Commodity) error
}

type CommodityService struct {
}

// PrintDetail 商品清單
func (s *CommodityService) PrintDetail(commodities ...dto.Commodity) error {

	for index, com := range commodities {
		fmt.Printf("序號: %d\n", index+1)
		fmt.Printf("商品ID: %d\n", com.ID)
		fmt.Printf("名稱: %s\n", com.Name)
		fmt.Printf("價格: %s\n", com.Price)
		fmt.Printf("數量: %d\n", com.Count)
		fmt.Println("-------------------")
	}
	return nil
}

func NewCommodityService() Commodities {
	return &CommodityService{}
}
