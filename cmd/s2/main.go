package main

import (
	"encoding/json"
	"flag"
	"github.com/cocoagaurav/devops/database"
	"github.com/cocoagaurav/devops/kafka"
	"github.com/cocoagaurav/devops/models"
	"github.com/golang/glog"
	"net/http"
)

func init() {
	flag.Parse()
}
func main() {
	database.Opendatabase()
	db := database.GetDB()
	kafka.InitKafkaConsumer()
	mssg := make(chan []byte)
	go kafka.KafkaConsumer(mssg)
	go func() {
		for {
			select {
			case twt := <-mssg:
				tw := &models.Tweet{}
				err := json.Unmarshal(twt, tw)
				if err != nil {
					glog.Errorf("error while unmarshalling tweet err :%v", err)
				}
				db.Create(tw)


			}
		}
	}()
	defer db.Close()
	defer kafka.CloseConsumner()
	glog.Info("listening at 8007")
	http.ListenAndServe(":8007", nil)
	glog.Info("shutting down s2")
}
