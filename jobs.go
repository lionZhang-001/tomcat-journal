package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v6"
	"os"
	"strings"
)

func job(){

	var tom tomcat
	ctx := context.Background()
	var servers  string
	var tomcatFile *os.File
	var  name , defaultsValue ,info * string


	//获取tomcat配置文件
	var names string = "fpath"
	name = &names
	var defaultsValues string = "/home/inspur/icp-collect/config/collect-tomcat.cfg"
	defaultsValue = &defaultsValues
	var information string = "reading tomcat config"
	info = &information
	tomcatFile = getParamFromCommand(name , defaultsValue ,info)
	defer tomcatFile.Close()

	s := bufio.NewScanner(tomcatFile)
	for s.Scan() {
		line := s.Text()
		lines := strings.Split(line , "|")
		tom.id = lines[0]
		tom.ip = lines[1]
		tom.port = lines[2]
		tom.identification = lines[3]

		fmt.Println(tom)

	}



	client ,err := elastic.NewClient(elastic.SetURL(servers) , elastic.SetSniff(false))
	if err != nil {
		panic(err)
		return
	}


	exists , err := client.IndexExists("tomcat-0002-accesslog-*").Do(ctx)
	if err != nil {
		fmt.Println("searching index went wrong : " , err)
		return
	}
	if exists {
		fmt.Println("the index tomcat-0002-accesslog exists .")
	}else {
		fmt.Println("the index tomcat-0002-accesslog doesnot exists .")
	}
}

func getParamFromCommand(name , defaultValues , info *string ) (f *os.File) {
	str := flag.String(*name , *defaultValues , *info )
	flag.Parse()
	f , err := os.Open(*str)
	if err != nil {
		fmt.Println("reading file collect-tomcat.cfg went wrong : ", err)
		return
	}

	return f

}
