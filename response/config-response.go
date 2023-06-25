package response

import "gin-demo/common"

type BlogConfigInfo struct {
	Name         string   `json:"name"`
	Avatar       string   `json:"avatar"`
	Icon         []Icon   `json:"icon"`
	MusicId      string   `json:"musicId"`
	Descriptions []string `json:"descriptions"`
	Content      string   `json:"content"`
}

type Icon struct {
	Icon       string `json:"icon"`
	Title      string `json:"title"`
	ModalImage string `json:"modalImage"`
	Href       string `json:"href"`
	Modal      bool   `json:"modal"`
}

func GetDefaultBlogConfigInfo() BlogConfigInfo {
	return BlogConfigInfo{
		Name:   common.USER,
		Avatar: common.AVATAR_ICON,
		Icon: []Icon{
			{
				Icon:       common.WECHAT_SVG,
				Title:      "点击显示我的二维码",
				ModalImage: common.WECHAT_IMG_PATH,
				Href:       "",
				Modal:      false,
			},
			{
				Icon:       common.QQ_SVG,
				Title:      common.QQ,
				ModalImage: "",
				Href:       "",
				Modal:      false,
			},
			{
				Icon:       common.GITHUB_SVG,
				Title:      "我的Github地址",
				ModalImage: "",
				Href:       common.GITHUB,
				Modal:      false,
			},
			{
				Icon:       common.EMAIL_SVG,
				Title:      common.EMAIL,
				ModalImage: "",
				Href:       "",
				Modal:      false,
			},
		},
		MusicId:      common.MUSIC_ID,
		Descriptions: common.DESCRIPTIONS,
		Content:      common.CONTENT,
	}
}
