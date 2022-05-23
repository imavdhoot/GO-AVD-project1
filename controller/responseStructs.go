package controller

type AddMemberResp struct {
	Email      string `json:"email"`
	MemberID   int    `json:"memberId"`
	MerchantID string `json:"merchantId"`
	Message    string `json:"message"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type MemberListByMerchantResp struct {
	Count   int `json:"count"`
	Members []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		MerchantID string `json:"merchantId"`
	} `json:"members"`
	Message  string `json:"message"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Status   int    `json:"status"`
}
