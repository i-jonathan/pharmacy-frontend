package inventory

import (
	"PharmUI/account"
	"time"

	"github.com/shopspring/decimal"
)

type inventoryResponse struct {
	Previous bool      `json:"previous"`
	Next     bool      `json:"next"`
	Page     int       `json:"page"`
	Count    int64     `json:"count"`
	Data     []product `json:"data"`
}

type product struct {
	ID             uint            `json:"id"`
	ItemName       string          `json:"name"`
	BarCode        string          `json:"bar_code"`
	Description    string          `json:"description"`
	Category       category        `json:"category"`
	CategoryID     int             `json:"category_id"`
	ExpiryDate     time.Time       `json:"expiry_date"`
	PurchaseDate   time.Time       `json:"purchase_date"`
	ProductionDate time.Time       `json:"production_date"`
	Quantity       uint            `json:"quantity"`
	PurchasePrice  decimal.Decimal `json:"purchase_price"`
	SellingPrice   decimal.Decimal `json:"selling_price"`
	User           account.User    `json:"user,-"`
	UserID         int             `json:"user_id,-"`
	ReorderLevel   int             `json:"reorder_level"`
	SKU            string          `json:"sku"`
	QuantitySold   int             `json:"quantity_sold"`
}

type category struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DateCreated time.Time    `json:"date_created"`
	Creator     account.User `json:"created_by,-"`
	CreatorID   int          `json:"creator_id,-"`
}
