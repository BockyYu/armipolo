package dto

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Commodity struct {
	ID    uint64          `json:"id"`    // ID(與紅龍網一致)
	Name  string          `json:"name"`  // 物品名稱
	Price decimal.Decimal `json:"price"` // 價格
	Count uint64          `json:"count"` // 數量
}

func (c Commodity) Validate() error {
	if c.ID == 0 {
		return fmt.Errorf("ID不能為空")
	}
	if c.Name == "" {
		return fmt.Errorf("商品名稱不能為空")
	}
	if c.Price.IsNegative() {
		return fmt.Errorf("價格不能為負")
	}
	if c.Count == 0 {
		return fmt.Errorf("數量必須大於0")
	}
	return nil
}
