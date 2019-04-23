package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v7"
)

//uses info
type User struct {
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	HeadPhoto interface{} `json:"head_photo"`
	Age       int         `json:"age"`
	Gender    string      `json:"gender"`
	School    string      `json:"school"`
}

type Elasti struct {
	client *elastic.Client
	index  string
	etype  string
}

func NewElastic(url string, index string, etype string) *Elasti {
	//创建链接 elastic  默认 localhost9200  设置URL elastic.SetURL() docker 开的 elasticsearch elastic.SetSniff(false)
	cli, e := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))

	if e != nil {
		panic(e)
	}
	fmt.Println(cli)
	ela := &Elasti{
		client: cli,
		index:  index,
		etype:  etype,
	}
	return ela
}

//添加
func (e *Elasti) Create(user *User) error {

	// index（库名）/type（表名）/id（ID字段，不传自动生成）post可以不用带ID
	//这个index方法 写入修改 elastic 接口 设计 用的是 RESTful
	//response, e := client.Index().
	//	Index("test_elastic").BodyJson(user).
	//	Do(context.Background())
	//
	//if e!=nil {
	//	panic(e)
	//}
	//
	//fmt.Println(response)

	_, i := e.client.Index().Index(e.index).Type(e.etype).BodyJson(user).Do(context.Background())
	if i != nil {
		return i
	}
	return nil
}

//通过ID获取
func (e *Elasti) Get(id string) (user *User, err error) {

	result, i := e.client.Get().Index(e.index).Type(e.etype).Id(id).Do(context.Background())

	if i != nil {
		return nil, i
	}

	var u User

	err = json.Unmarshal(result.Source, &u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func elasticClient(url string) (*elastic.Client, error) {

	//创建链接 elastic  默认 localhost9200  设置URL elastic.SetURL() docker 开的 elasticsearch elastic.SetSniff(false)
	//client, e := elastic.NewClient(elastic.SetSniff(false))
	//
	//if e != nil {
	//	panic("NewClient fail:" + e.Error())
	//}

	//创建链接 elastic  默认 localhost9200  设置URL elastic.SetURL() docker 开的 elasticsearch elastic.SetSniff(false)
	client, e := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))

	if e != nil {
		panic("NewClient fail:" + e.Error())
	}

	return client, nil
}

func main() {
	user := &User{
		Name:      "elastic",
		Email:     "github@gamil.com",
		Phone:     "13888888888",
		HeadPhoto: `{"image":"/uploads/2019/01/23/yuiejswwwfcs.jgp","tumb_image":"/uploads/2019/01/23/tumb_yuiejswwwfcs.jgp"}`,
		Age:       28,
		Gender:    "0", //0 男 1女
		School:    "北京航空航天大学",
	}

	_ = user

	e := NewElastic("http://127.0.0.1:9200", "test_elastic", "_doc")

	user, err := e.Get("vMOmSmoB6i_QyiY4e3vw")

	if err != nil {
		panic(e)
	}

	fmt.Printf("%+v", user)
}
