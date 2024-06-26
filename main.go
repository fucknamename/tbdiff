package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"tbdiff/handle"
	"tbdiff/utils"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
	_ = utils.MakeDir("./sql")
}

func main() {
	// 全量数据库（要参考的数据库）
	db_s := flag.String("ds", "", "short of target mysql database connection string")
	// // 全量数据表（要参考的数据表）
	// table_s := flag.String("ts", "", "shor of target table")

	// 本地数据库（要更新的数据库）
	db_l := flag.String("dl", "", "short of local mysql database connection string")
	// // 本地数据表（要更新的数据表）
	// table_l := flag.String("tl", "", "shor of local table")
	flag.Parse()

	if err := handle.InitDB(*db_s, *db_l); err != nil {
		fmt.Println(err.Error())
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if strings.Contains(path, "favicon.ico") {
			http.NotFound(w, r)
		} else {
			w.Write([]byte("ok"))
		}
	})
	mux.HandleFunc("/diff/*", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")
		table := parts[len(parts)-1] // 取最后一个参数作为要对比的表名

		handle.HandleCompared(w, r, table)
	})
	mux.HandleFunc("/backup/*", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")
		table := parts[len(parts)-1] // 取最后一个参数作为要备份的表名

		handle.HandleBackUp(w, r, table)
	})
	mux.HandleFunc("/table/*", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")
		count := len(parts)

		db := ""
		if count > 1 {
			db = parts[1] // 取最第二个参数作为要查看的库
		}

		handle.HandleShowTable(w, r, db, count > 2)
	})

	srv := &http.Server{
		Addr:    ":3844",
		Handler: mux,
	}

	fmt.Println("db table diff server run at 3844 port")
	fmt.Println("telegram: @echoty")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefulExitWeb(srv)
}

// 优雅退出
func gracefulExitWeb(server *http.Server) {
	quit := make(chan os.Signal, 4)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-quit

	fmt.Println("got a signal\ndbdiff server stoped", sig)

	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(cxt); err != nil {
		fmt.Println("err", err)
	}

	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
	os.Exit(0)
}
