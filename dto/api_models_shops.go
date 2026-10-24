// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Models.Shops
package dto

import (
	"time"
)

type ExchangeLimit struct {
	Id                         int64            `json:"id" msgpack:"id"`                                                 // Key: 0
	ExchangeShopThingId        int64            `json:"exchangeShopThingId" msgpack:"exchangeShopThingId"`               // Key: 1
	ReplaceType                ShopReplaceTypes `json:"replaceType" msgpack:"replaceType"`                               // Key: 2
	SpecifiedNumberOfDaysLimit *int             `json:"specifiedNumberOfDaysLimit" msgpack:"specifiedNumberOfDaysLimit"` // Key: 3
	ExchangedCount             int              `json:"exchangedCount" msgpack:"exchangedCount"`                         // Key: 4
	Until                      *time.Time       `json:"until" msgpack:"until"`                                           // Key: 5
}

type JewelShop struct {
	Id                    int64      `json:"id" msgpack:"id"`                                       // Key: 0
	JewelShopItemMasterId int64      `json:"jewelShopItemMasterId" msgpack:"jewelShopItemMasterId"` // Key: 1
	PurchaseCount         int        `json:"purchaseCount" msgpack:"purchaseCount"`                 // Key: 2
	TotalPurchaseCount    int        `json:"totalPurchaseCount" msgpack:"totalPurchaseCount"`       // Key: 3
	RePurchaseDate        *time.Time `json:"rePurchaseDate" msgpack:"rePurchaseDate"`               // Key: 4
}

type LoginPassStatus struct {
	Id         int64     `json:"id" msgpack:"id"`                 // Key: 0
	ValidUntil time.Time `json:"validUntil" msgpack:"validUntil"` // Key: 1
}

type PermanentMarketThing struct {
	PermanentMarketThingMasterId int64 `json:"permanentMarketThingMasterId" msgpack:"permanentMarketThingMasterId"` // Key: 0
	PurchaseCount                int   `json:"purchaseCount" msgpack:"purchaseCount"`                               // Key: 1
}

type StarPassStatus struct {
	Id                  int64         `json:"id" msgpack:"id"`                                   // Key: 0
	Type                StarPassTypes `json:"type" msgpack:"type"`                               // Key: 1
	TotalPurchasedCount int           `json:"totalPurchasedCount" msgpack:"totalPurchasedCount"` // Key: 2
	ValidUntil          time.Time     `json:"validUntil" msgpack:"validUntil"`                   // Key: 3
}

type ViewedShop struct {
	Id                   int64                   `json:"id" msgpack:"id"`                                     // Key: 0
	ExchangeShopMasterId *int64                  `json:"exchangeShopMasterId" msgpack:"exchangeShopMasterId"` // Key: 1
	LastViewedAt         time.Time               `json:"lastViewedAt" msgpack:"lastViewedAt"`                 // Key: 2
	ViewedShopCategory   ViewedShopCategoryTypes `json:"viewedShopCategory" msgpack:"viewedShopCategory"`     // Key: 3
}
