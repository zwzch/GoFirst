package tomcat

import (
	"encoding/xml"
	"io/ioutil"
	//"fmt"
	"log"

	//"strings"
	"strconv"
	"os"
	"bufio"
	//"fmt"
	"io"
)

type parseResult struct {
	ip string
	port int
	staticpath string
	sessionSize int
	}

type StringResources struct {
	XMLName        xml.Name         `xml:"resources"`
	ResourceString []ResourceString `xml:"property"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"property"`
	PropertyName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}
func XmlParse(path string ) *parseResult{


	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	ip := result.ResourceString[0].InnerText
	port := result.ResourceString[1].InnerText
	parse :=new(parseResult)
	parse.ip = ip
	parse.port,_ = strconv.Atoi(port)
	staticpath :=result.ResourceString[2]
	parse.staticpath = staticpath.InnerText
	sessionSize :=result.ResourceString[3]
	parse.sessionSize,_ = strconv.Atoi(sessionSize.InnerText)
	return parse
	//fmt.Println(ip,port)
	//log.Println(result)
	//log.Println(result.ResourceString)
	//for _, o := range result.ResourceString {
	//	log.Println(o.PropertyName + "===" + o.InnerText)
	//}
	//content,err :=ioutil.ReadFile(path)
	//checkErr(err)
	//var xmlc xmlConfig
	//err = xml.Unmarshal(content,&xmlc)
	//checkErr(err)
	//fmt.Println(xmlc)
}

func HtmlBody(path string,pathName string)string{
	pathAll := path+pathName
	//ioutil.ReadFile(pathAll)
	f, err := os.Open(pathAll)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	body:=""
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		line=line+"\n"
		body+=line
	}
	//fmt.Println(body)
	return body
}

