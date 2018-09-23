package mobula

import (
	"os"
	"fmt"
	"bufio"
	"encoding/json"
	"github.com/json-iterator/go"
	"strings"
	)

func checkErr(e error){
	if e != nil{
		panic(e)
	}
}


func Transform2() {

	//fin,_ := os.Open("test.txt")
	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
	//input := bufio.NewScanner(fin)//初始化一个扫表对象
	for input.Scan() { //扫描输入内容
		lines := input.Text() //把输入内容转换为字符串
		//fmt.Println(line)//输出到标准输出
		lineArr := strings.Split(lines, "\n")
		for i, n := 0, len(lineArr); i < n; i++ {
			if (lineArr[i]!=""&&lineArr[i]!="NULL") {
				var Fields []string
				if (strings.ContainsAny(lineArr[i], "/t")==true){
					Fields = strings.Split(lineArr[i], "\t")
							if (len(Fields)!=20){
								continue
							} else {
								//fmt.Println(Fields)
								parseSplit(Fields)
					}
					} else {
					continue
				}

				//fmt.Println(len(Fields))
			}
		}
	}
}
func Transform()  {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(err)
		}
	}()
	fin,_ := os.Open("test.txt")
	//input := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
	input := bufio.NewScanner(fin)//初始化一个扫表对象
	for input.Scan() {//扫描输入内容

		lines := input.Text()//把输入内容转换为字符串
		//fmt.Println(line)//输出到标准输出
		lineArr := strings.Split(lines,"\n")
		for i,n := 0, len(lineArr); i<n; i++ {

			Fields :=strings.Split(lineArr[i],"\t")
			//fmt.Println(Fields[12])
			parseSplit(Fields)



			//fmt.Println(coin, logid, level, pid, gid, key, ac, fbid)
			//printSlice(adpkgsArr)
			//printSlice(adidsArrRaw)
			//printSlice(adidsArr)
			//fmt.Println(key)
			//fmt.Println(ac)
			}
	}
}

func parseSplit(Fields []string)  {

	json_raw := Fields[12]
	//fmt.Println(len(Fields))
	strings.Replace(json_raw,"\\","",-1)
	var report_json []byte = []byte(json_raw)

	key := jsoniter.Get(report_json, "key").ToString()

	ac := jsoniter.Get(report_json, "ac").ToString()
	if (ac=="") {ac="NULL"}

	coins := jsoniter.Get(report_json, "coins").ToString()
	if(coins==""){coins="0"}

	logid := jsoniter.Get(report_json, "logid").ToString()
	if(logid==""){logid="NULL"}

	level := jsoniter.Get(report_json, "level").ToString()
	if(level==""){level="1"}

	pid := jsoniter.Get(report_json, "pid").ToString()
	if(pid==""){pid="0"}

	tid := jsoniter.Get(report_json, "tid").ToString()
	if(tid==""){tid="0"}

	gid := jsoniter.Get(report_json, "gid").ToString()
	if(gid==""){gid="0"}

	adid := jsoniter.Get(report_json, "id").ToString()
	if(adid==""){adid="0"}

	gtype := jsoniter.Get(report_json, "gtype").ToString()
	if(gtype==""){adid="0"}

	bid := jsoniter.Get(report_json, "bid").ToString()
	if(bid==""){bid="0"}

	entry := jsoniter.Get(report_json, "entry").ToString()
	if(entry==""){bid="NULL"}

	adpkg := jsoniter.Get(report_json, "adpkg").ToString()
	if(adpkg==""){adpkg="NULL"}

	adpkgs := jsoniter.Get(report_json, "adpkgs").ToString()
	if(adpkgs==""){adpkgs="NULL"}

	adids := jsoniter.Get(report_json, "ids").ToString()
	if(adids==""){adpkgs="NULL"}

	sid := jsoniter.Get(report_json, "sid").ToString()
	if(sid==""){sid="0"}

	ts := jsoniter.Get(report_json, "ts").ToString()
	if(ts==""){ts="NULL"}

	sc := jsoniter.Get(report_json, "sc").ToString()
	if(sc==""){sc="NULL"}

	st := jsoniter.Get(report_json, "st").ToString()
	if(st==""){st="NULL"}

	directgp := jsoniter.Get(report_json, "directgp").ToString()
	if(directgp==""){directgp="True"}
	fbid_all :=Fields[len(Fields)-1]
	var fbid string=""
	if (strings.Contains(fbid_all,"_") && fbid_all != ""){
		fbid = strings.Split(fbid_all,"_")[1]
	}else {
		fbid = "NULL"
	}

	entry2 := jsoniter.Get(report_json, "entry2").ToString()
	if(entry2==""){entry2="NULL"}

	vs := jsoniter.Get(report_json, "vs").ToString()
	if(vs==""){vs="0"}
	var adpkgsArr [] string

	if (adpkg!="" && adpkg !="NULL") {
		var adpkgsArrRaw [] string
		if (adpkgs!="NULL") {
			adpkgsArrRaw = strings.Split(strings.TrimRight(strings.TrimLeft(adpkgs, "["), "]"), ",")
		}
		adpkgsArr = append(adpkgsArrRaw, adpkg)
	}
	adidsArrRaw := strings.Split(strings.TrimRight(strings.TrimLeft(adids, "["), "]"), ",")
	if (adidsArrRaw == nil){
		adidsArrRaw = append(adidsArrRaw, adid)
	}
	if (len(adidsArrRaw) == 1) {
		adid = adidsArrRaw[0]
	}
	if (len(adpkgsArr)!=0) {
		//report_time, apppkg, child, token, lc, opcode, sdk_vn, app_vc, app_vn, time_stamp, licence
		for e := range adpkgsArr {
			s:=[]string{string(Fields[0]), Fields[1], Fields[3], Fields[4], Fields[5], string(Fields[6]),
					Fields[7], string(Fields[8]), string(Fields[14]), Fields[9], Fields[10], logid,
					entry, string(level), key, string(pid), string(tid), string(gid), string(adid),
					string(bid), string(gtype), adpkgsArr[e], string(sid), Fields[13],
					string(ac), string(coins), Fields[15], Fields[16], Fields[17],
					Fields[2], string(ts), string(Fields[18]), string(directgp), string(adid),
					string(Fields[19]), string(fbid), string(sc), string(st), entry2, string(vs)}
			fmt.Println(strings.Join(s,"\t"))
			//fmt.Println(string(Fields[0]), Fields[1], Fields[3], Fields[4], Fields[5], string(Fields[6]),
			//	Fields[7], string(Fields[8]), string(Fields[14]), Fields[9], Fields[10], logid,
			//	entry, string(level), key, string(pid), string(tid), string(gid), string(adid),
			//	string(bid), string(gtype), adpkgsArr[e], string(sid), Fields[13],
			//	string(ac), string(coins), Fields[15], Fields[16], Fields[17],
			//	Fields[2], string(ts), string(Fields[18]), string(directgp), string(adid),
			//	string(Fields[19]), string(fbid), string(sc), string(st), entry2, string(vs))
		}
	}else {
		for e := range adidsArrRaw {
			s:=[]string{string(Fields[0]), Fields[1], Fields[3], Fields[4], Fields[5], string(Fields[6]),
					Fields[7], string(Fields[8]), string(Fields[14]), Fields[9], Fields[10], logid,
					entry, string(level), key, string(pid), string(tid), string(gid), string(adid),
					string(bid), string(gtype), adidsArrRaw[e], string(sid), Fields[13],
					string(ac), string(coins), Fields[15], Fields[16], Fields[17],
					Fields[2], string(ts), string(Fields[18]), string(directgp), string(adid),
					string(Fields[19]), string(fbid), string(sc), string(st), entry2, string(vs)}
			fmt.Println(strings.Join(s,"\t"))

			//fmt.Println(string(Fields[0]), Fields[1], Fields[3], Fields[4], Fields[5], string(Fields[6]),
			//	Fields[7], string(Fields[8]), string(Fields[14]), Fields[9], Fields[10], logid,
			//	entry, string(level), key, string(pid), string(tid), string(gid), string(adid),
			//	string(bid), string(gtype), adidsArrRaw[e], string(sid), Fields[13],
			//	string(ac), string(coins), Fields[15], Fields[16], Fields[17],
			//	Fields[2], string(ts), string(Fields[18]), string(directgp), string(adid),
			//	string(Fields[19]), string(fbid), string(sc), string(st), entry2, string(vs))
		}
	}
}
func printSlice(x []string){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func ParseJson()  {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json_iterator.Marshal(group)
	os.Stdout.Write(b)
}

func Unmarshal()  {
	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

	var animals2 []Animal
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	json_iterator.Unmarshal(jsonBlob, &animals2)
	fmt.Printf("%+v", animals2)


}

func Test()  {

	//str:="[55350588,53011003,56277357,56253589]"
	//ids:=strings.Split(strings.TrimRight(strings.TrimLeft(str,"["),"]"),",")
	//
	//for e := range ids {
	//	fmt.Println(ids[e])
	//}
	//s:=[]string{"1","2","3"}
	//
	//fmt.Println(strings.Join(s,","))
	//a:=123

}