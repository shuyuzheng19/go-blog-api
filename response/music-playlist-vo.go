package response

type MusicVo struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
	Pic    string `json:"pic"`
	Lrc    string `json:"lrc"`
}
