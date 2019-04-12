package es

import (
	"context"
	"github.com/golang/glog"
	"github.com/olivere/elastic"
	"time"
)
var client *elastic.Client

func ElacticConn(){
	var err error
	client, err = elastic.NewSimpleClient(elastic.SetURL("http://elasticsearch-cluster:9200"))
	if err != nil {
		glog.Errorf("error connecting ES err:%v", err)
		time.Sleep(5 * time.Second)
		ElacticConn()
	}
	glog.Info("elasticsearch connected")
}

func GetEs()*elastic.Client{
	if client == nil{
		ElacticConn()
	}
	exist ,err := client.IndexExists("tweets").Do(context.Background())
	if err!=nil{
		glog.Errorf("error while checking index exist err:%v", err)
	}

	if !exist {
		_, err = client.CreateIndex("tweets").Do(context.Background())
		if err != nil {
			glog.Errorf("error while creating index err:%v", err)
		}
	}

	return client
}