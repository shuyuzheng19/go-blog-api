package response

type Token struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	Type        string `json:"type"`
	Role        string `json:"role"`
	Ip          string `json:"ip"`
	CreateAt    string `json:"createAt"`
	ExpireAt    string `json:"expireAt"`
}
