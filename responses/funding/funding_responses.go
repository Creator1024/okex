package funding

import (
	models "github.com/Creator1024/okex/models/funding"
)

type Data struct {
	TotalBal string  `json:"totalBal"`
	Ts       string  `json:"ts"`
	Details  Details `json:"details"`
}
type Details struct {
	Funding string `json:"funding"`
	Trading string `json:"trading"`
	Classic string `json:"classic"`
	Earn    string `json:"earn"`
}

type (
	GetCurrencies struct {
		Code       string             `json:"code"`
		Msg        string             `json:"msg,omitempty"`
		Currencies []*models.Currency `json:"data"`
	}
	GetAssetValuation struct {
		Code           string                  `json:"code"`
		Msg            string                  `json:"msg,omitempty"`
		AssetValuation []models.AssetValuation `json:"data"`
	}
	GetBalance struct {
		Code     string            `json:"code"`
		Msg      string            `json:"msg,omitempty"`
		Balances []*models.Balance `json:"data"`
	}
	FundsTransfer struct {
		Code      string             `json:"code"`
		Msg       string             `json:"msg,omitempty"`
		Transfers []*models.Transfer `json:"data"`
	}
	AssetBillsDetails struct {
		Code  string         `json:"code"`
		Msg   string         `json:"msg,omitempty"`
		Bills []*models.Bill `json:"data"`
	}
	GetDepositAddress struct {
		Code             string                   `json:"code"`
		Msg              string                   `json:"msg,omitempty"`
		DepositAddresses []*models.DepositAddress `json:"data"`
	}
	GetDepositHistory struct {
		Code             string                   `json:"code"`
		Msg              string                   `json:"msg,omitempty"`
		DepositHistories []*models.DepositHistory `json:"data"`
	}
	Withdrawal struct {
		Code        string               `json:"code"`
		Msg         string               `json:"msg,omitempty"`
		Withdrawals []*models.Withdrawal `json:"data"`
	}
	GetWithdrawalHistory struct {
		Code                string                      `json:"code"`
		Msg                 string                      `json:"msg,omitempty"`
		WithdrawalHistories []*models.WithdrawalHistory `json:"data"`
	}
	PiggyBankPurchaseRedemption struct {
		Code       string              `json:"code"`
		Msg        string              `json:"msg,omitempty"`
		PiggyBanks []*models.PiggyBank `json:"data"`
	}
	GetPiggyBankBalance struct {
		Code     string                     `json:"code"`
		Msg      string                     `json:"msg,omitempty"`
		Balances []*models.PiggyBankBalance `json:"data"`
	}
)
