package response

type Token struct {
	Type string `json:"type"`

	Token string `json:"token"`

	Create string `json:"create"`

	Expire string `json:"expire"`
}
