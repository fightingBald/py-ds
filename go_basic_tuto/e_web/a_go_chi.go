package e_web

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func testChi(t *testing.T) {
	// 新建一个 chi 路由器
	r := chi.NewRouter()

	// 定义路由和 handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 从 query string 里取出参数 q
		q := r.URL.Query().Get("q")

		if q == "hello" {
			w.Write([]byte("world"))
		} else {
			w.Write([]byte("unknown"))
		}
	})

	// 启动 http server，监听 8080 端口
	http.ListenAndServe(":8080", r)
}
