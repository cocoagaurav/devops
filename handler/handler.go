package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/devops/database"
	"github.com/cocoagaurav/devops/es"
	"github.com/cocoagaurav/devops/kafka"
	"github.com/cocoagaurav/devops/models"
	"github.com/golang/glog"
	"net/http"

)

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	reg := &models.Registration{}
	err := json.NewDecoder(r.Body).Decode(reg)
	if err != nil {
		glog.Errorf("error while parsing request err:%v", err)
	}
	db.Create(reg)
	json.NewEncoder(w).Encode("you are ready to tweet")

}

func Health(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "congrats server is up")

}

func Tweet(w http.ResponseWriter, r *http.Request) {
	tweet := &models.Tweet{}
	err := json.NewDecoder(r.Body).Decode(tweet)
	if err != nil {
		glog.Errorf("error while parsing request err:%v", err)
	}
	esClient := es.GetEs()
	_, err = esClient.Index().Index("tweets").Type("tweet").BodyJson(tweet).Do(context.Background())
	if err != nil{
		glog.Errorf("error while writing tweets into es err:%v",err)
	}
	b, err := json.Marshal(tweet)
	if err != nil{
		glog.Errorf("error while marshalling tweets  err: %v",err)
		return
	}
	kafka.SentMessg(b)
	json.NewEncoder(w).Encode("you are ready to tweet")

}
