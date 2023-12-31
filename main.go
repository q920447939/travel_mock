package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/io/flutter_app/json/travel_page.json", func(w http.ResponseWriter, r *http.Request) {
		// 读取 travel_page.json 文件
		b, err := ioutil.ReadFile("travel_page.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		// 将 JSON 数据写入响应
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	http.HandleFunc("/restapi/soa2/16189/json/searchTripShootListForHomePageV2", func(w http.ResponseWriter, r *http.Request) {

		// 获取请求参数
		pageIndex := GetUrlArg(r, "pageIndex")

		// 判断 pageIndex
		if pageIndex == "1" {
			// 读取 1.json 文件
			b, err := ioutil.ReadFile("1.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			// 将 JSON 数据写入响应
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
		} else {
			// 返回错误响应
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("pageIndex 必须为 1"))
		}
	})

	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {

		// 获取请求参数
		name := GetUrlArg(r, "name")

		// 读取 1.json 文件
		b, err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 将 JSON 数据写入响应
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(b)
	})

	http.ListenAndServe("0.0.0.0:18088", nil)
}

// 获取URL的GET参数
func GetUrlArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg = values.Get(name)
	return arg
}
