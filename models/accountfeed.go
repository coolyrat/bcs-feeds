package models

type AccountFeed struct {
	UserId         int64   `json:"userId"`
	Balance        float64 `json:"balance"`
	Budget         float64 `json:"budget"`
	UserStat       int     `json:"userStat"`
	BalancePackage int     `json:"balancePackage"`
	UaStatus       int     `json:"uaStatus"`
	ValidFlows     []int   `json:"validFlows"`
}
