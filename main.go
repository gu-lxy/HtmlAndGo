package main

import (
	"LolHeros/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("欢迎来到Lol...")

	//1、确定采集数据源：https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js
	url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"
	//2、执行网络请求，获取数据源
	client := http.Client{}//相当于：浏览器软件

	request, err := http.NewRequest("GET",url,nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response,err :=client.Do(request)//开始发起一个请求
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dataBytes, err :=ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(dataBytes))


	//2、从数据源中解析得到所有的目标数据
	// 从string类型数据转换成为Go语言结构体对象类型
	//马歇尔
	//json.Marshal()
	//安马歇尔
	var herolist entity.HeroList
	err =json.Unmarshal(dataBytes, &herolist)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(herolist)
	//list := entity.HeroList{Version:"1.0",FileName:"hero_list.js",FileTime:"2020-06-16 15:22"}
	//fmt.Println(list)

}
