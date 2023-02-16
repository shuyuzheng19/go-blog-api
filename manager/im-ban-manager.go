package manager

//
//import (
//	"encoding/json"
//	"time"
//	"vs-blog-api/common"
//	"vs-blog-api/config"
//	"vs-blog-api/utils"
//)
//
//type IpBanManager struct {
//
//}
//
//func NewIpBanManager() *IpBanManager {
//	return &IpBanManager{}
//}
//
//type IpBan struct {
//	//IP封禁原因
//	Description string `json:"description"`
//	//封禁时间
//	BanTime time.Time `json:"banTime"`
//	//解封时间
//	RestoreTime time.Time `json:"restoreTime"`
//	//封禁的接口
//	Url string `json:"url"`
//}
//
//func (*IpBanManager) GetIpBan (url string,ip string)*IpBan {
//
//	var KEY = ip + "->" + url
//
//	ipBanStr := config.Redis.HGet(common.IpBanMap, KEY).Val()
//
//	if ipBanStr == "" {
//		return nil
//	}else{
//		var ipBan *IpBan
//		if err:=json.Unmarshal([]byte(ipBanStr),&ipBan);err!=nil{
//			return nil
//		}
//		now := time.Now()
//
//		if now.After(ipBan.RestoreTime){
//			config.Redis.HDel(common.IpBanMap,KEY)
//			return nil
//		}
//
//		return ipBan
//	}
//}
//
//func (ban IpBanManager) IsBanGetIpBan(url string,ip string)*IpBan{
//
//	var md5Ip = utils.ToMd5String([]byte(ip))
//
//	var ipBan = ban.GetIpBan(url,md5Ip)
//
//	if ipBan!=nil {return ipBan}
//
//	var KEY = common.IpBan +":"+md5Ip+":"+url
//
//	if config.Redis.Exists(KEY).Val()==0 {
//
//		config.Redis.Set(KEY, 1, common.IpTime).Val()
//
//	}else{
//
//		count:=utils.ToInt(config.Redis.Get(KEY).Val())
//
//		if count>common.IpCount{
//			now:=time.Now()
//			var ipBan=IpBan{
//				Description: "访问频率过快",
//				BanTime:     now,
//				RestoreTime: time.Now().Add(common.IpBanExpire),
//				Url:         url,
//			}
//			return &ipBan
//		}
//
//		config.Redis.Set(KEY,count+1,0)
//
//	}
//
//
//	return nil
//}
