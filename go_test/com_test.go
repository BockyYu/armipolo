package go_test

import (
	"armipolo/dto"
	"armipolo/service"
	"github.com/shopspring/decimal" // 需要導入 decimal 包
	"testing"
)

func TestDo(t *testing.T) {
	commodityService := service.NewCommodityService()
	testCases := []dto.Commodity{
		{
			ID:    1,
			Name:  "商品A",
			Price: decimal.RequireFromString("100"),
			Count: 2,
		},
		{
			ID:    2,
			Name:  "商品B",
			Price: decimal.RequireFromString("200"),
			Count: 1,
		},
	}

	err := commodityService.PrintDetail(testCases...)
	if err != nil {
		t.Errorf("PrintDetail() error = %v", err)
	}
}
