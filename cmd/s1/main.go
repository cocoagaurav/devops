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
	"net/http"
)

func init(){
	flag.Parse()
}
func main() {
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
	glog.Infof("server running at 8008")
	http.ListenAndServe(":8008", r)

}
