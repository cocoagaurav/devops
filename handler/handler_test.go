package handler

import (
	"bytes"
	"context"
	"flag"
	"github.com/cocoagaurav/devops/database"
	"github.com/cocoagaurav/devops/es"
	"github.com/cocoagaurav/devops/kafka"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	flag.Parse()
	database.Opendatabase()
	database.InitDB()
	kafka.NewKafkaProducer()
	es.ElacticConn()

	test := m.Run()
	database.CloseDb()
	kafka.CloseProducer()
	os.Exit(test)
}
func TestLogin(t *testing.T) {
	body := `{
	"email":"gaurav@hotcocoasoftware.com",
	"password":"password"
}`

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(body)))
	if err != nil{
		t.Logf("error file creating request err:%v ", err)
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(res, req)

	if res.Code != 200{
		t.Fatalf("expected status code 200 got %v", res.Code)
	}
	es := es.GetEs()
	re, _, err := es.Ping("http://localhost:9200").Do(context.Background())
	if err!= nil{
		t.Fatalf("es is not working err: %v", err )
	}
	t.Logf("es res is :%v", re)

}


func TestTweets(t *testing.T) {
	body := `{
	"title":"tweet test",
	"discription":"desc test"
}`
	req, err := http.NewRequest("POST", "/tweet", bytes.NewBuffer([]byte(body)))

	if err != nil{
		t.Logf("error file posting request err:%v ", err)
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Tweet)
	handler.ServeHTTP(res, req)

	if res.Code != 200{
		t.Fatalf("expected status code 200 got %v", res.Code)
	}

}
