package main

import (
	"io/ioutil"

	"github.com/Goryudyuma/anaconda"
	"github.com/Goryudyuma/tlc/tlc"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

func loadconfig() []byte {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	return data
}

func loadyaml() tlc.MyTwitterKey {
	key := tlc.MyTwitterKey{}

	err := yaml.Unmarshal(loadconfig(), &key)
	if err != nil {
		panic(err)
	}
	return key
}

func main() {
	key := loadyaml()
	spew.Dump(key)
	//	fmt.Println(key.AccessToken)

	/*
		operator := byte('*')
		list1 := tlc.List{Listname: "aaa", Owner_screen_name: "Goryudyuma", Owner_id: 0}
		list2 := tlc.List{Listname: "bbb", Owner_screen_name: "Goryudyuma", Owner_id: 0}
		resultliststring := "ccc"
	*/

	anaconda.SetConsumerKey(key.ConsumerKey)
	anaconda.SetConsumerSecret(key.ConsumerSecret)

	spew.Dump(anaconda.AuthorizationURL("https://tlc.063.jp/"))

	//api := anaconda.NewTwitterApi(key.AccessToken, key.AccessTokenSecret)

	//tlc.Tlc(*api, operator, list1, list2, resultliststring)
}
