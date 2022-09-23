package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
	vientiane "vientiane/pub/idl/grpc"
)

func TestGetArticle(t *testing.T) {
	wg := WorkerGroup{
		Workers: make(chan chan *http.Request, 100),
	}
	resCh := make(chan string, 100000)
	for i := 0; i < 100; i++ {
		worker := make(chan *http.Request, 0)
		wg.Workers <- worker
		go Worker(worker, resCh)
	}
	for j := 0; j < 20000; j++ {
		go func() {
			worker := <-wg.Workers
			defer func() {
				wg.Workers <- worker
			}()
			request, _ := http.NewRequest("GET", "http://localhost:8088/article/get/1", nil)
			request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
			request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
			request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
			request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
			request.Header.Set("Cache-Control", "max-age=0")
			request.Header.Set("Connection", "keep-alive")
			worker <- request
		}()
	}
	time.Sleep(time.Second * 1)
	log.Println("res-len: ", len(resCh))
}

func TestAddArticle(t *testing.T) {
	wg := WorkerGroup{
		Workers: make(chan chan *http.Request, 50),
	}
	resCh := make(chan string, 100000)
	for i := 0; i < 50; i++ {
		worker := make(chan *http.Request, 0)
		wg.Workers <- worker
		go Worker(worker, resCh)
	}
	param := &vientiane.Article{
		Title:   "title",
		Content: "content",
		Author:  "http",
	}
	bytes, _ := json.Marshal(param)
	for j := 0; j < 10000; j++ {
		go func() {
			worker := <-wg.Workers
			defer func() {
				wg.Workers <- worker
			}()
			request, _ := http.NewRequest("POST", "http://localhost:8088/article/add", strings.NewReader(string(bytes)))
			request.Header.Set("Cache-Control", "max-age=0")
			request.Header.Set("Connection", "keep-alive")
			request.Header.Set("Content-Type", "application/json")
			worker <- request
		}()
	}
	time.Sleep(time.Second * 1)
	log.Println("res-len: ", len(resCh))
}

type WorkerGroup struct {
	Workers chan chan *http.Request
}

func Worker(worker chan *http.Request, ch chan string) {
	client := &http.Client{}
	for {
		select {
		case request := <-worker:
			response, err := client.Do(request)
			if err != nil {
				log.Println(err)
			}

			if request.Method == http.MethodGet && response.StatusCode == 200 {
				body, _ := ioutil.ReadAll(response.Body)
				ch <- string(body)
				response.Body.Close()
			}
		}
	}
}
