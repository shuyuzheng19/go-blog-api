package manager

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type MusicPlayListManager struct {
}

func NewMusicPlayListManager() *MusicPlayListManager {
	return &MusicPlayListManager{}
}

func (*MusicPlayListManager) SetMusicCloudPlayList(mid string) {
	var url = "https://api.injahow.cn/meting/?type=playlist&id=" + mid

	fmt.Println(url)

	parse := utils.GetColly(url)

	var musicInfo []response.MusicVo

	parse.OnResponse(func(response2 *colly.Response) {
		array := gjson.ParseBytes(response2.Body).Array()
		for _, result := range array {
			music := response.MusicVo{
				Artist: result.Get("artist").Str,
				Name:   result.Get("name").Str,
				Url:    result.Get("url").Str,
				Pic:    result.Get("pic").Str,
				Lrc:    result.Get("lrc").Str,
			}
			musicInfo = append(musicInfo, music)
		}
	})

	parse.Wait()

	config.Redis.Set(common.MusicPlayList, utils.ObjectToJson(musicInfo), -1)
}

func (*MusicPlayListManager) SetMusicArrayPlayList(musics []response.MusicVo) {

	var musicInfo []string

	for _, music := range musics {
		musicInfo = append(musicInfo, utils.ObjectToJson(music))
	}

	config.Redis.RPush(common.MusicPlayList, musicInfo)
}

func (*MusicPlayListManager) GetMusicPlayList() []response.MusicVo {

	const KEY = common.MusicPlayList

	val := config.Redis.Exists(KEY).Val()

	if val == 0 {
		return make([]response.MusicVo, 0)
	}

	var playList []response.MusicVo

	musics := config.Redis.Get(KEY).Val()

	if err := json.Unmarshal([]byte(musics), &playList); err != nil {
		return make([]response.MusicVo, 0)
	}

	return playList
}
