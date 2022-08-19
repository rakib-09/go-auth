package types

type BrandReq struct {
	Title     string `json:"title,omitempty"`
	CompanyId uint   `json:"companyId,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
}

type BrandResp struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	CompanyId uint        `json:"companyId"`
	Phone     string      `json:"phone"`
	Address   string      `json:"address"`
	Company   CompanyResp `json:"company,omitempty"`
}
