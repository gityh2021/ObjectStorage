package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Handler w http.ResponseWriter用于写入HTTP的响应。
// *http.Request 用于获取Request请求。
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodGet {
		Get(w, r)
		return
	}
	if m == http.MethodPut {
		Put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func Get(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	strSplits := strings.Split(url, "/")
	if len(strSplits) == 3 {
		objectName := strSplits[2]
		f, err := os.Open(DataPath + "/" + objectName)

		// 文件关闭失败
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println("defer关闭文件失败:", err)
			} else {
				fmt.Println("defer 关闭文件成功")
			}
		}()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if _, err = io.Copy(w, f); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func Put(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	strSplits := strings.Split(url, "/")
	if len(strSplits) == 3 {
		objectName := strSplits[2]
		f, err := os.Create(DataPath + "/" + objectName)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 文件关闭
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println("defer关闭文件失败:", err)
			} else {
				fmt.Println("defer关闭文件成功")
			}
		}()

		// 文件拷贝
		if _, err = io.Copy(f, r.Body); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
