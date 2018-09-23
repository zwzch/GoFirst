package mobula

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
	"time"
	"log"
)

func PbulishTransform()  {
	//fin,_ := os.Open("test3.txt")
	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
	//input := bufio.NewScanner(fin)//初始化一个扫表对象
	for input.Scan() { //扫描输入内容
		lines := input.Text() //把输入内容转换为字符串
		lineArr := strings.Split(lines,"\n")
		for i, n := 0, len(lineArr); i < n; i++ {
			if (lineArr[i]!=""&&lineArr[i]!="NULL") {
				var Fields []string
				if (strings.ContainsAny(lineArr[i], "/t")==true){
					Fields = strings.Split(lineArr[i], "\t")
					if (len(Fields)!=30){
						continue
					} else {
						//publish_time, api_name, apppkg, sdk_vn, app_vc, token, op_code, license, child, lc, resolution, \
						//logid, publish_json, sid, stype, income_type, server_v, app_vn, publish_ext, publish_forms, \
						//dllv, svn, anid, goid, hostname, sidtags, qs_request, publish_ext_user, \
						//publish_ext_md5, adpkg_request
						//fmt.Println(Fields)
						parsePublishTransform(Fields)
						//parseTest(Fields)
					}
				} else {
					continue
				}
				//fmt.Println(len(Fields))
			}
		}

	}
}

func parseTest(Fields []string) {


	publish_json_raw :=Fields[12]
	fmt.Println(Fields[12])

	strings.Replace(publish_json_raw,"\\","",-1)
	var publish_json []byte = []byte(publish_json_raw)
	groups := jsoniter.Get(publish_json, "mgs").ToString()
	fmt.Println(groups)

	na_adid := jsoniter.Get(publish_json, "adIds").ToString()
	if (na_adid == "") {
		na_adid = jsoniter.Get(publish_json, "adids").ToString()
	}
	if (na_adid == "") {
		na_adid = "0"
	}
	na_adidArr := getArr(na_adid)
	fmt.Println(na_adidArr)
}

func parsePublishTransform(Fields []string){
	logFile, err := os.Create("./"+ time.Now().Format("goerror") + ".txt");
	if err != nil {
		fmt.Println(err);
	}
	loger := log.New(logFile, "test_", log.Ldate|log.Ltime|log.Lshortfile);

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("Exception has been caught.")
			//fmt.Println(err)
			loger.Println(err)
		}
	}()
	opcode := "NULL"
	op_code := Fields[6]
	//fmt.Println(op_code)
	op_arr := strings.Split(string(op_code), ",")
	for e := range op_arr {
		if (op_arr[e] != "null" && op_arr[e] != "NULL") {
			opcode = op_arr[e]
		}
	}

	publish_json_raw := Fields[12]
	if (publish_json_raw != "" && publish_json_raw != "NULL" && publish_json_raw != "null") {
		strings.Replace(publish_json_raw,"\\","",-1)
		var publish_json []byte = []byte(publish_json_raw)
		pid := jsoniter.Get(publish_json, "pId").ToString()
		if (pid == "") {pid = jsoniter.Get(publish_json, "pid").ToString()}
		if (pid == "") {pid="0"}
		tid := jsoniter.Get(publish_json, "tId").ToString()
		if (tid == "") {tid = jsoniter.Get(publish_json, "tid").ToString()}
		if (tid == "") {tid="0"}
		pn := jsoniter.Get(publish_json, "pn").ToString()
		if(pn == ""){pn="0"}
		ps := jsoniter.Get(publish_json, "ps").ToString()
		if(ps == ""){ps="0"}
		groups := jsoniter.Get(publish_json, "mgs").ToString()
		pc := jsoniter.Get(publish_json, "pc").ToString()
		if(pc == ""){pc="-1"}
		isNew := jsoniter.Get(publish_json, "isNew").ToString()
		if(isNew == ""){isNew="-1"}
		rcp := jsoniter.Get(publish_json, "rcp").ToString()
		if(rcp == ""){rcp="-1"}
		bigImg := jsoniter.Get(publish_json, "bigImg").ToString()
		if(bigImg == ""){bigImg="-1"}
		tcppType := jsoniter.Get(publish_json, "tcppType").ToString()
		if(tcppType == ""){tcppType="-1"}
		rlogid := jsoniter.Get(publish_json, "rlogid").ToString()
		if(rlogid == ""){rlogid="NULL"}
		dlHighIsNew := jsoniter.Get(publish_json, "dlHighIsNew").ToString()
		if(dlHighIsNew == ""){dlHighIsNew="-1"}
		materialList := jsoniter.Get(publish_json, "materialList").ToString()
		if(materialList == ""){materialList="-1"}
		pkgBigimgList := jsoniter.Get(publish_json, "pkgBigimgList").ToString()
		if(pkgBigimgList == ""){pkgBigimgList="-1"}
		pkgLanguageList := jsoniter.Get(publish_json, "pkgLanguageList").ToString()
		if(pkgLanguageList == ""){pkgLanguageList="-1"}
		rlogidArr := getArr(rlogid)
		if (len(rlogidArr) == 0) {
			rlogid = "NULL"
		}
		groupArr := GetGroups(groups);
		if (groupArr!=nil && len(groupArr) != 0 ) {
			for e := range groupArr {

				var group_json []byte = []byte(groupArr[e])
				bid := "0"
				gid := jsoniter.Get(group_json, "mgId").ToString()
				if (gid == "") {gid = jsoniter.Get(group_json, "mgid").ToString()}
				if (gid == "") {gid="0"}

				gtype := jsoniter.Get(group_json, "mgType").ToString()
				if (gtype == "") {gtype = jsoniter.Get(group_json, "mgtype").ToString()}
				if (gtype == "") {gtype="0"}

				preclick := jsoniter.Get(group_json, "pc").ToString()
				if(preclick == ""){preclick = "-1"}
				preclickArr := getArr(preclick)

				if (gtype == "100" || gtype == "102" ) {
					ids := jsoniter.Get(group_json, "adIds").ToString()
					if (ids == "") {ids = jsoniter.Get(group_json, "adids").ToString()}
					if (ids == "") {ids="[]"}
					adindex := 0

					pn := jsoniter.Get(group_json, "pn").ToString()
					if(pn == ""){pn="0"}

					ps := jsoniter.Get(group_json, "ps").ToString()
					if(ps == ""){ps="0"}


					isNew := jsoniter.Get(group_json, "isNew").ToString()
					if(isNew == ""){isNew="-1"}
					isNewArr := getArr(isNew)

					rcp := jsoniter.Get(group_json, "rcp").ToString()
					if(rcp == ""){rcp="-1"}
					rcpArr :=getArr(rcp)

					bigImg := jsoniter.Get(group_json, "bigImg").ToString()
					if(bigImg == ""){bigImg="-1"}
					bigImgArr := getArr(bigImg)

					tcppType := jsoniter.Get(group_json, "tcppType").ToString()
					if(tcppType == ""){tcppType="-1"}
					tcppTypeArr := getArr(tcppType)

					dlHighIsNew := jsoniter.Get(group_json, "dlHighIsNew").ToString()
					if(dlHighIsNew == ""){dlHighIsNew="-1"}
					dlHighIsNewArr := getArr(dlHighIsNew)

					materialList := jsoniter.Get(group_json, "materialList").ToString()
					if(materialList == ""){materialList="-1"}
					materialListArr := getArr(materialList)

					pkgBigimgList := jsoniter.Get(publish_json, "pkgBigimgList").ToString()
					if(pkgBigimgList == ""){pkgBigimgList="-1"}
					pkgBigimgListArr := getArr(pkgBigimgList)

					pkgLanguageList := jsoniter.Get(publish_json, "pkgLanguageList").ToString()
					if(pkgLanguageList == ""){pkgLanguageList="-1"}
					pkgLanguageListArr := getArr(pkgLanguageList)

					rlogid := jsoniter.Get(group_json, "rlogid").ToString()
					rlogidArr := getArr(rlogid)

					idsArr := getArr(ids)
					for e := range idsArr {
						if (rlogidArr != nil) {
							if (len(rlogidArr) == 0&&rlogidArr[0] != "-1") {
								rlogid="NULL"
							} else {
								rlogid=rlogidArr[adindex]

								}
						}


						if (preclickArr != nil) {
							if (len(preclickArr) == 0&&preclickArr[0] != "-1") {
								preclick="NULL"
							} else {
								preclick=preclickArr[adindex]
							}
						}

						if (isNewArr != nil&&isNewArr[0] != "-1") {
							if (len(isNewArr) == 0) {
								isNew="NULL"
							} else {
								isNew=isNewArr[adindex]
							}
						}

						if (rcpArr != nil&&rcpArr[0] != "-1") {
							if (len(rcpArr) == 0) {
								rcp="NULL"
							} else {
								rcp=rcpArr[adindex]
							}
						}

						if (bigImgArr != nil&&bigImgArr[0] != "-1") {
							if (len(bigImgArr) == 0) {
								bigImg="NULL"
							} else {
								bigImg=bigImgArr[adindex]
							}
						}

						if (tcppTypeArr != nil&&tcppTypeArr[0] != "-1") {
							if (len(tcppTypeArr) == 0) {
								tcppType="NULL"
							} else {
								tcppType=tcppTypeArr[adindex]
							}
						}

						if (dlHighIsNewArr != nil&&dlHighIsNewArr[0] != "-1") {
							if (len(dlHighIsNewArr) == 0) {
								dlHighIsNew="NULL"
							} else {
								dlHighIsNew=dlHighIsNewArr[adindex]
							}
						}

						if (materialListArr != nil&&materialListArr[0] != "-1") {
							if (len(materialListArr) == 0) {
								materialList="NULL"
							} else {
								materialList=materialListArr[adindex]
							}
						}

						if (pkgBigimgListArr != nil&& pkgBigimgListArr[0] != "-1") {
							if (len(pkgBigimgListArr) == 0) {
								pkgBigimgList="NULL"
							} else {
								pkgBigimgList=pkgBigimgListArr[adindex]
							}
						}

						if (pkgLanguageListArr != nil&&pkgLanguageListArr[0] != "-1") {
							if (len(pkgLanguageListArr) == 0) {
								pkgLanguageList="NULL"
							} else {
								pkgLanguageList=pkgLanguageListArr[adindex]
							}
						}

						adindex += 1
						adindexStr :=strconv.Itoa(adindex)
						//publish_time, api_name, apppkg, sdk_vn, app_vc, token, op_code, license, child, lc, resolution, \
						//        logid, publish_json, sid, stype, income_type, server_v, app_vn, publish_ext, publish_forms, \
						//        dllv, svn, anid, goid, hostname, sidtags, qs_request, publish_ext_user, \
						//        publish_ext_md5, adpkg_request
						s := []string{string(Fields[0]), Fields[1], Fields[11], Fields[2], string(Fields[3]), string(Fields[4]),
							string(Fields[17]), Fields[8], string(opcode), Fields[5], Fields[9], Fields[7], string(Fields[13]), string(pid), string(tid),
						    string(gid), string(gtype), string(idsArr[e]), string(bid), string(pn),
							string(ps), string(adindexStr), Fields[15], string(Fields[16]), Fields[14], string(preclick), Fields[18], string(isNew),
							string(rcp), rlogid, Fields[19], string(bigImg), string(tcppType), string(Fields[20]), string(Fields[21]), string(dlHighIsNew),
							Fields[22], Fields[23], Fields[24], Fields[25], string(materialList), string(idsArr[e]), string(Fields[26]), Fields[27],
							Fields[28], string(pkgBigimgList), string(pkgLanguageList), string(Fields[29])}
						fmt.Println(strings.Join(s, "\t"))
					}
				} else if (gtype == "101") {
					banners := jsoniter.Get(group_json, "bnrs").ToString()
					if(banners == ""){banners="[]"}
					bannersArr := getArr(banners)
					b_adidIndex := 0
					if (bannersArr != nil && bannersArr[0] != "-1") {
						if (len(bannersArr) != 0) {
							for e := range bannersArr {
								b_adidIndex += 1
								b_adindexStr :=strconv.Itoa(b_adidIndex)
								var bannerJson []byte = []byte(bannersArr[e])
								b_bid := jsoniter.Get(bannerJson, "bId").ToString()
								if (b_bid == "") {b_bid = jsoniter.Get(bannerJson, "bid").ToString()}
								if (b_bid == "") {b_bid="0"}

								b_adid := jsoniter.Get(bannerJson, "adId").ToString()
								if (b_adid == "") {b_adid = jsoniter.Get(bannerJson, "adid").ToString()}
								if (b_adid == "") {b_adid="NULL"}

								s := []string{string(Fields[0]), Fields[1], Fields[11], Fields[2], string(Fields[3]), string(Fields[4]),
									string(Fields[17]), Fields[8], string(opcode), Fields[5], Fields[9], Fields[7], string(Fields[13]), string(pid), string(tid),
									string(gid), string(gtype), string(b_adid), string(b_bid), string(pn),
									string(ps), string(b_adindexStr), Fields[15], string(Fields[16]), Fields[14], string(preclick), Fields[18], string(isNew),
									string(rcp), rlogid, Fields[19], string(bigImg), string(tcppType), string(Fields[20]), string(Fields[21]), string(dlHighIsNew),
									Fields[22], Fields[23], Fields[24], Fields[25], string(materialList), string(b_adid), string(Fields[26]), Fields[27],
									Fields[28], string(pkgBigimgList), string(pkgLanguageList), string(Fields[29])}
								fmt.Println(strings.Join(s, "\t"))


								}
						}
					}

				}
 			}
		} else {
			na_adid := jsoniter.Get(publish_json, "adIds").ToString()
			if (na_adid == "") {
				na_adid = jsoniter.Get(publish_json, "adids").ToString()
			}
			if (na_adid == "") {
				na_adid = "0"
			}
			na_adidArr := getArr(na_adid)
			//fmt.Println(na_adidArr)
			//n_adidIndex := 0
			if (na_adid == "0" || na_adid == "") {
				s := []string{string(Fields[0]), Fields[1], Fields[11], Fields[2], string(Fields[3]), string(Fields[4]),
					string(Fields[17]), Fields[8],
					string(opcode), Fields[5], Fields[9], Fields[7], string(Fields[13]), "0", "0", "0", "0", string("-1"), "0", string(pn),
					string(ps), "-1", Fields[15], string(Fields[16]), Fields[14], "-1", Fields[18], string(isNew),
					string(rcp), rlogid, Fields[19], string(bigImg), string(tcppType), string(Fields[20]), string(Fields[21]), string(dlHighIsNew),
					Fields[22], Fields[23], Fields[24], Fields[25], string(materialList), "-1", string(Fields[26]), Fields[27],
					Fields[28], string(pkgBigimgList), string(pkgLanguageList), string(Fields[29])}
				fmt.Println(strings.Join(s, "\t"))
			} else {
				preclick := jsoniter.Get(publish_json, "pc").ToString()
				if(preclick == ""){preclick = "-1"}
				preclickArr := getArr(preclick)

				isNew := jsoniter.Get(publish_json, "isNew").ToString()
				if(isNew == ""){isNew="-1"}
				isNewArr := getArr(isNew)

				rcp := jsoniter.Get(publish_json, "rcp").ToString()
				if(rcp == ""){rcp="-1"}
				rcpArr :=getArr(rcp)

				bigImg := jsoniter.Get(publish_json, "bigImg").ToString()
				if(bigImg == ""){bigImg="-1"}
				bigImgArr := getArr(bigImg)

				tcppType := jsoniter.Get(publish_json, "tcppType").ToString()
				if(tcppType == ""){tcppType="-1"}
				tcppTypeArr := getArr(tcppType)

				dlHighIsNew := jsoniter.Get(publish_json, "dlHighIsNew").ToString()
				if(dlHighIsNew == ""){dlHighIsNew="-1"}
				dlHighIsNewArr := getArr(dlHighIsNew)

				materialList := jsoniter.Get(publish_json, "materialList").ToString()
				if(materialList == ""){materialList="-1"}
				materialListArr := getArr(materialList)

				pkgBigimgList := jsoniter.Get(publish_json, "pkgBigimgList").ToString()
				if(pkgBigimgList == ""){pkgBigimgList="-1"}
				pkgBigimgListArr := getArr(pkgBigimgList)

				pkgLanguageList := jsoniter.Get(publish_json, "pkgLanguageList").ToString()
				if(pkgLanguageList == ""){pkgLanguageList="-1"}
				pkgLanguageListArr := getArr(pkgLanguageList)

				rlogid := jsoniter.Get(publish_json, "rlogid").ToString()
				rlogidArr := getArr(rlogid)
				adindex := 0
				for e := range na_adidArr {
					if (rlogidArr != nil) {
						if (len(rlogidArr) == 0&&rlogidArr[0] != "-1") {
							rlogid="NULL"
						} else {
							rlogid=rlogidArr[adindex]

						}
					}
					if (preclickArr != nil) {
						if (len(preclickArr) == 0&&preclickArr[0] != "-1") {
							preclick="NULL"
						} else {
							preclick=preclickArr[adindex]
						}
					}

					if (isNewArr != nil&&isNewArr[0] != "-1") {
						if (len(isNewArr) == 0) {
							isNew="NULL"
						} else {
							isNew=isNewArr[adindex]
						}
					}

					if (rcpArr != nil&&rcpArr[0] != "-1") {
						if (len(rcpArr) == 0) {
							rcp="NULL"
						} else {
							rcp=rcpArr[adindex]
						}
					}

					if (bigImgArr != nil&&bigImgArr[0] != "-1") {
						if (len(bigImgArr) == 0) {
							bigImg="NULL"
						} else {
							bigImg=bigImgArr[adindex]
						}
					}

					if (tcppTypeArr != nil&&tcppTypeArr[0] != "-1") {
						if (len(tcppTypeArr) == 0) {
							tcppType="NULL"
						} else {
							tcppType=tcppTypeArr[adindex]
						}
					}

					if (dlHighIsNewArr != nil&&dlHighIsNewArr[0] != "-1") {
						if (len(dlHighIsNewArr) == 0) {
							dlHighIsNew="NULL"
						} else {
							dlHighIsNew=dlHighIsNewArr[adindex]
						}
					}

					if (materialListArr != nil&&materialListArr[0] != "-1") {
						if (len(materialListArr) == 0) {
							materialList="NULL"
						} else {
							materialList=materialListArr[adindex]
						}
					}

					if (pkgBigimgListArr != nil&& pkgBigimgListArr[0] != "-1") {
						if (len(pkgBigimgListArr) == 0) {
							pkgBigimgList="NULL"
						} else {
							pkgBigimgList=pkgBigimgListArr[adindex]
						}
					}

					if (pkgLanguageListArr != nil&&pkgLanguageListArr[0] != "-1") {
						if (len(pkgLanguageListArr) == 0) {
							pkgLanguageList="NULL"
						} else {
							pkgLanguageList=pkgLanguageListArr[adindex]
						}
					}
					adindex += 1
					adindexStr :=strconv.Itoa(adindex)
					s := []string{string(Fields[0]), Fields[1], Fields[11], Fields[2], string(Fields[3]), string(Fields[4]),
						string(Fields[17]), Fields[8], string(opcode), Fields[5], Fields[9], Fields[7], string(Fields[13]), "0", "0",
						"0", "0", string(na_adidArr[e]), "0", string(pn),
						string(ps), string(adindexStr), Fields[15], string(Fields[16]), Fields[14], string(preclick), Fields[18], string(isNew),
						string(rcp), rlogid, Fields[19], string(bigImg), string(tcppType), string(Fields[20]), string(Fields[21]), string(dlHighIsNew),
						Fields[22], Fields[23], Fields[24], Fields[25], string(materialList), string(na_adidArr[e]), string(Fields[26]), Fields[27],
						Fields[28], string(pkgBigimgList), string(pkgLanguageList), string(Fields[29])}
					fmt.Println(strings.Join(s, "\t"))


					}
			}
		}
		//pid := jsoniter.Get(publish_json, "tId").ToString()
		//if (pid=="") {pid="0"}
		}
	}

func getArr(strRaw string) []string  {

	if ( strRaw == "" ) {
		return nil
	} else {
		strArr := strings.Split(strings.TrimRight(strings.TrimLeft(strRaw, "["), "]"), ",")
		return strArr
		}
	//return strArr
}

func GetGroups(str string) []string {
	//var str_json []byte = []byte(str)

	if ( str == "" ) {
		return nil
	}else {
	strArr := strings.Split(strings.TrimRight(strings.TrimLeft(str, "["), "]"), "},{")
	//fmt.Println(strArr)
	return strArr }
	}

func PublishTest()  {

	fmt.Println("error\t" + "xxx")
}