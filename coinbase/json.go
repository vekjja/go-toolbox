package coinbase

import "time"

type PriceData struct {
	TradeID int       `json:"trade_id"`
	Price   string    `json:"price"`
	Size    string    `json:"size"`
	Time    time.Time `json:"time"`
	Bid     string    `json:"bid"`
	Ask     string    `json:"ask"`
	Volume  string    `json:"volume"`
}

type AccountData []struct {
	ID             string `json:"id"`
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Hold           string `json:"hold"`
	Available      string `json:"available"`
	ProfileID      string `json:"profile_id"`
	TradingEnabled bool   `json:"trading_enabled"`
}

type Currency struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	MinSize       string        `json:"min_size"`
	Status        string        `json:"status"`
	Message       string        `json:"message"`
	MaxPrecision  string        `json:"max_precision"`
	ConvertibleTo []interface{} `json:"convertible_to"`
	Details       struct {
		Type                  string      `json:"type"`
		Symbol                string      `json:"symbol"`
		NetworkConfirmations  int         `json:"network_confirmations"`
		SortOrder             int         `json:"sort_order"`
		CryptoAddressLink     string      `json:"crypto_address_link"`
		CryptoTransactionLink string      `json:"crypto_transaction_link"`
		PushPaymentMethods    []string    `json:"push_payment_methods"`
		GroupTypes            []string    `json:"group_types"`
		DisplayName           interface{} `json:"display_name"`
		ProcessingTimeSeconds interface{} `json:"processing_time_seconds"`
		MinWithdrawalAmount   float64     `json:"min_withdrawal_amount"`
		MaxWithdrawalAmount   int         `json:"max_withdrawal_amount"`
	} `json:"details"`
}

type Currencies []struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	MinSize       string        `json:"min_size"`
	Status        string        `json:"status"`
	Message       string        `json:"message"`
	MaxPrecision  string        `json:"max_precision"`
	ConvertibleTo []interface{} `json:"convertible_to"`
	Details       struct {
		Type                  string        `json:"type"`
		Symbol                string        `json:"symbol"`
		NetworkConfirmations  int           `json:"network_confirmations"`
		SortOrder             int           `json:"sort_order"`
		CryptoAddressLink     string        `json:"crypto_address_link"`
		CryptoTransactionLink string        `json:"crypto_transaction_link"`
		PushPaymentMethods    []string      `json:"push_payment_methods"`
		GroupTypes            []interface{} `json:"group_types"`
		DisplayName           interface{}   `json:"display_name"`
		ProcessingTimeSeconds interface{}   `json:"processing_time_seconds"`
		MinWithdrawalAmount   float64       `json:"min_withdrawal_amount"`
		MaxWithdrawalAmount   int           `json:"max_withdrawal_amount"`
	} `json:"details"`
}

type OrderData []struct {
	ID            string `json:"id"`
	ClientOid     string `json:"client_oid"`
	Price         string `json:"price"`
	Size          string `json:"size"`
	ProductID     string `json:"product_id"`
	ProfileID     string `json:"profile_id"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	TimeInForce   string `json:"time_in_force"`
	PostOnly      bool   `json:"post_only"`
	CreatedAt     string `json:"created_at"`
	FillFees      string `json:"fill_fees"`
	FilledSize    string `json:"filled_size"`
	ExecutedValue string `json:"executed_value"`
	Status        string `json:"status"`
	Settled       bool   `json:"settled"`
}
