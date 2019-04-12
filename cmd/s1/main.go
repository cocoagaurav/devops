package main

import (
	"flag"
	"github.com/cocoagaurav/devops/database"
	"github.com/cocoagaurav/devops/es"
	"github.com/cocoagaurav/devops/handler"
	"github.com/cocoagaurav/devops/kafka"
	"github.com/go-chi/chi"
	"github.com/golang/glog"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/profile"
	"net/http"
)

func init(){
	flag.Parse()
}
func main() {
	defer profile.Start(profile.MemProfile).Stop()
	database.Opendatabase()
	database.InitDB()
	es.ElacticConn()
	kafka.NewKafkaProducer()

	defer kafka.CloseProducer()
	defer database.CloseDb()
	r := chi.NewRouter()

	r.Get("/health", handler.Health)

	r.Post("/login", handler.Login)
	r.Post("/tweet", handler.Tweet)


	// Register pprof handlers
	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//
	//r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	//r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	//r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	//r.Handle("/debug/pprof/block", pprof.Handler("block"))


	glog.Infof("server running at 8008")
	http.ListenAndServe(":8008", r)

}
