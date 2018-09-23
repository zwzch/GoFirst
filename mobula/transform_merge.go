package mobula

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
)

func MergeTransform()  {
	//fin,_ := os.Open("test3.txt")
	input := bufio.NewScanner(os.Stdin) //初始化一个扫表对象
	//input := bufio.NewScanner(fin)//初始化一个扫表对象
	for input.Scan() { //扫描输入内容
		lines := input.Text() //把输入内容转换为字符串
		lineArr := strings.Split(lines,"\n")
		for i, n := 0, len(lineArr); i < n; i++ {
			if (lineArr[i]!=""&&lineArr[i]!="NULL") {
				var Fields []string
				if (strings.ContainsAny(lineArr[i], "\t")==true){
					Fields = strings.Split(lineArr[i], "\t")
					//fmt.Println(len(Fields))
					if (len(Fields)!=45){
						continue
					} else {
						//logid, publish_time, op_code, publish_json, sid, stype, server_v, \
						//income_type, report_time, time_stamp, app_pkg, child, token, lc, \
						//sdk_vn, app_vc, app_vn, license, key, entry, level, pid, tid, gid, gtype, \
						//bid, adid, adpkg, ac, coins, goid, anid,publish_ext, report_ext, ip,mdu, \
						//publish_forms, dllv, svn, hostname, sidtags, directgp, ps_request,\
						//api_name, vs
						parseMergeTransform(Fields)
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

func parseMergeTransform(Fields []string)  {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	//fmt.Println(Fields)
	logid, publish_time, op_code, publish_json_raw, sid, stype, server_v,
	income_type, report_time, time_stamp, app_pkg, child, token, lc,
	sdk_vn, app_vc, app_vn, license, key, entry, level, pid, tid, gid, gtype,
	bid, adid, adpkg, ac, coins, goid, anid,publish_ext, report_ext, ip,mdu,
	publish_forms, dllv, svn, hostname, sidtags, directgp, ps_request,
	api_name, vs := getMergeTuple(Fields)
	//fmt.Println(logid, publish_time, op_code, publish_json_raw, sid, stype, server_v,
	//	income_type, report_time, time_stamp, app_pkg, child, token, lc,
	//	sdk_vn, app_vc, app_vn, license, key, entry, level, pid, tid, gid, gtype,
	//	bid, adid, adpkg, ac, coins, goid, anid,publish_ext, report_ext, ip,mdu,
	//	publish_forms, dllv, svn, hostname, sidtags, directgp, ps_request,
	//	api_name, vs)
	op_code_sub := "NULL"
	//fmt.Println(op_code)
	op_arr := strings.Split(string(op_code), ",")
	for e := range op_arr {
		if (op_arr[e] != "null" && op_arr[e] != "NULL") {
			op_code_sub = op_arr[e]
		}
	}
	strings.Replace(publish_json_raw,"\\","",-1)
	var publish_json []byte = []byte(publish_json_raw)
	pn := jsoniter.Get(publish_json, "pn").ToString()
	if(pn == ""){pn="0"}
	ps := jsoniter.Get(publish_json, "ps").ToString()
	if(ps == ""){ps="0"}
	rlogid := jsoniter.Get(publish_json, "rlogid").ToString()
	if(rlogid == ""){rlogid="NULL"}
	if (sid == "null" || sid == "NULL") {
		sid = "0"
	}
	if (key == "show" && adid == "0") {

		p_pid := jsoniter.Get(publish_json, "pId").ToString()
		if (p_pid == "") {p_pid = jsoniter.Get(publish_json, "pid").ToString()}
		if (p_pid == "") {p_pid="0"}

		p_tid := jsoniter.Get(publish_json, "tId").ToString()
		if (p_tid == "") {p_tid = jsoniter.Get(publish_json, "tid").ToString()}
		if (p_tid == "") {tid="0"}

		p_sid := jsoniter.Get(publish_json, "sid").ToString()
		if (p_sid == "") {p_sid = "0"}

		groups := jsoniter.Get(publish_json, "mgs").ToString()

		preclick := jsoniter.Get(publish_json, "pc").ToString()
		if(preclick == ""){preclick = "-1"}
		isNew := jsoniter.Get(publish_json, "isNew").ToString()
		if(isNew == ""){isNew="-1"}
		dlHighIsNew := jsoniter.Get(publish_json, "dlHighIsNew").ToString()
		if(dlHighIsNew == ""){dlHighIsNew="-1"}
		materialList := jsoniter.Get(publish_json, "materialList").ToString()
		if(materialList == ""){materialList="-1"}
		pkgBigimgList := jsoniter.Get(publish_json, "pkgBigimgList").ToString()
		if(pkgBigimgList == ""){pkgBigimgList="-1"}
		pkgLanguageList := jsoniter.Get(publish_json, "pkgLanguageList").ToString()
		if(pkgLanguageList == ""){pkgLanguageList="-1"}
		groupArr := GetGroups(groups);

		if (groupArr!=nil && len(groupArr) != 0 ) {
			for e := range groupArr {
				var group_json []byte = []byte(groupArr[e])
				p_gid := jsoniter.Get(group_json, "mgId").ToString()
				if (p_gid == "") {p_gid = jsoniter.Get(group_json, "mgid").ToString()}
				if (p_gid == "") {p_gid="0"}

				p_gtype := jsoniter.Get(group_json, "mgType").ToString()
				if (p_gtype == "") {p_gtype = jsoniter.Get(group_json, "mgtype").ToString()}
				if (p_gtype == "") {p_gtype="0"}

				if (gtype == "100" || gtype == "102" ) {
					adindex :=0
					adids := jsoniter.Get(group_json, "adIds").ToString()
					if (adids == "") {adids = jsoniter.Get(group_json, "adids").ToString()}
					if (adids == "") {adids="[]"}
					pn := jsoniter.Get(group_json, "pn").ToString()
					if(pn == ""){pn="0"}

					ps := jsoniter.Get(group_json, "ps").ToString()
					if(ps == ""){ps="0"}

					preclick := jsoniter.Get(group_json, "pc").ToString()
					if(preclick == ""){preclick = "-1"}
					preclickArr := getArr(preclick)

					isNew := jsoniter.Get(group_json, "isNew").ToString()
					if(isNew == ""){isNew="-1"}
					isNewArr := getArr(isNew)

					rcp := jsoniter.Get(group_json, "rcp").ToString()
					if(rcp == ""){rcp="-1"}
					rcpArr :=getArr(rcp)

					rlogid := jsoniter.Get(group_json, "rlogid").ToString()
					rlogidArr := getArr(rlogid)

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

					adidsArr := getArr(adids)
					for e := range adidsArr {
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
						s := []string{
							publish_time, report_time, time_stamp, income_type,
							string(server_v), child, token, lc, op_code_sub, sdk_vn, app_pkg,
							app_vn, app_vc, license, entry, string(level), key, string(p_pid), string(p_tid),
							string(p_gid), string(adidsArr[e]), adpkg, "0", string(p_gtype), string(sid), stype, string(pn), string(ps),
							logid, string(ac), string(coins), goid, anid, string(preclick), publish_ext, report_ext, ip,
							string(adindexStr), string(isNew), string(rcp), mdu, rlogid, publish_forms, string(bigImg),
							string(tcppType), string(dllv), string(svn), string(dlHighIsNew), hostname, sidtags, string(materialList),
							directgp, string(adidsArr[e]), string(ps_request), api_name, string(pkgBigimgList), string(pkgLanguageList),
							string(vs)}
						fmt.Println(strings.Join(s, "\t"))
					}

				} else if (gtype == "101") {
					banners := jsoniter.Get(group_json, "bnrs").ToString()
					if(banners == ""){banners="[]"}
					bannersArr := getArr(banners)
					b_adidIndex := 0
					if (bannersArr != nil && bannersArr[0] != "-1") {
						b_adidIndex += 1
						b_adindexStr :=strconv.Itoa(b_adidIndex)
						var bannerJson []byte = []byte(bannersArr[e])
						if (len(bannersArr) != 0) {
							p_bid := jsoniter.Get(bannerJson, "bId").ToString()
							if (p_bid == "") {p_bid = jsoniter.Get(bannerJson, "bid").ToString()}
							if (p_bid == "") {p_bid="0"}

							tmp_adid := jsoniter.Get(bannerJson, "adId").ToString()
							if (tmp_adid == "") {tmp_adid = jsoniter.Get(bannerJson, "adid").ToString()}
							if (tmp_adid == "") {tmp_adid="0"}

							s := []string{publish_time, report_time, time_stamp,
								income_type, string(server_v), child, token, lc, op_code_sub, sdk_vn, app_pkg,
								app_vn, app_vc, license, entry, string(level), key, string(p_pid), string(p_tid),
								string(p_gid), string(tmp_adid), adpkg, string(p_bid), string(p_gtype), string(sid),
								stype, string(pn), string(ps), logid,
								string(ac), string(coins), goid, anid, string(preclick), publish_ext, report_ext, ip,
								string(b_adindexStr), string(isNew), "NULL", mdu, rlogid, publish_forms, "NULL",
								"NULL", string(dllv), string(svn), string(dlHighIsNew), hostname, sidtags, string(materialList),
								directgp, string(tmp_adid), string(ps_request), api_name, string(pkgBigimgList), string(pkgLanguageList),
								string(vs)}
							fmt.Println(strings.Join(s, "\t"))
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
			if (na_adid == "0" || na_adid == "" || len(na_adidArr) == 0) {
				s := []string{publish_time, report_time, time_stamp,
					income_type, string(server_v), child, token, lc, op_code_sub, sdk_vn, app_pkg,
					app_vn, app_vc, license, entry, level, key, "0", "0",
					"0", "-1", adpkg, "0", string(gtype), string(sid), stype, "0", "0", logid,
					string(ac), string(coins), goid, anid, "-1", publish_ext, report_ext, ip,
					string("-1"), string(isNew), string(""), mdu, rlogid, publish_forms, string(""),
					string(""), string(dllv), string(svn), string(dlHighIsNew), hostname, sidtags, string(materialList),
					directgp, "-1", string(ps_request), api_name, string(pkgBigimgList), string(pkgLanguageList), string(vs)}
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
					s := []string{publish_time, report_time, time_stamp,
						income_type, string(server_v), child, token, lc, op_code_sub, sdk_vn, app_pkg,
						app_vn, app_vc, license, entry, level, key, "0", "0",
						"0", string(na_adidArr[e]), adpkg, "0", string(gtype), string(sid), stype, "0", "0", logid,
						string(ac), string(coins), goid, anid, string(preclick), publish_ext, report_ext, ip,
						string(adindexStr), string(isNew), string(rcp), mdu, rlogid, publish_forms, string(bigImg),
						string(tcppType), string(dllv), string(svn), string(dlHighIsNew), hostname, sidtags,
						string(materialList), directgp, string(na_adidArr[e]), string(ps_request), api_name, string(pkgBigimgList),
						string(pkgLanguageList), string(vs)}
					fmt.Println(strings.Join(s, "\t"))
				}
				}
		}
	} else {
		//fmt.Println(publish_json_raw)
		strings.Replace(publish_json_raw,"\\","",-1)
		var publish_json []byte = []byte(publish_json_raw)
		na_adid := jsoniter.Get(publish_json, "adIds").ToString()
		if (na_adid == "") {
			na_adid = "-1"
		}
		na_adidArr := getArr(na_adid)
		//fmt.Println(na_adidArr)
		adindex := 0
		if (len(na_adidArr) != 0 && na_adidArr != nil) {
			for e := range na_adidArr {
				if (na_adidArr[e] == adid) {
					break;
				}
				adindex += 1
			}
		}

		preclick := jsoniter.Get(publish_json, "pc").ToString()
		if(preclick == ""){preclick = "-1"}
		preclickArr := getArr(preclick)

		isNew := jsoniter.Get(publish_json, "isNew").ToString()
		if(isNew == ""){isNew="-1"}
		isNewArr := getArr(isNew)

		rcp := jsoniter.Get(publish_json, "rcp").ToString()
		if(rcp == ""){rcp="-1"}
		rcpArr :=getArr(rcp)

		rlogid := jsoniter.Get(publish_json, "rlogid").ToString()
		rlogidArr := getArr(rlogid)

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


		if (preclickArr != nil) {
			if (len(preclickArr) == 0&&preclickArr[0] != "-1") {
				preclick="-1"
			} else {
				preclick=preclickArr[adindex]
			}
		}

		if (isNewArr != nil&&isNewArr[0] != "-1") {
			if (len(isNewArr) == 0) {
				isNew="-1"
			} else {
				isNew=isNewArr[adindex]
			}
		}

		if (rcpArr != nil&&rcpArr[0] != "-1") {
			if (len(rcpArr) == 0) {
				rcp="-1"
			} else {
				rcp=rcpArr[adindex]
			}
		}

		if (bigImgArr != nil&&bigImgArr[0] != "-1") {
			if (len(bigImgArr) == 0) {
				bigImg="-1"
			} else {
				bigImg=bigImgArr[adindex]
			}
		}

		if (tcppTypeArr != nil&&tcppTypeArr[0] != "-1") {
			if (len(tcppTypeArr) == 0) {
				tcppType="-1"
			} else {
				tcppType=tcppTypeArr[adindex]
			}
		}

		if (dlHighIsNewArr != nil&&dlHighIsNewArr[0] != "-1") {
			if (len(dlHighIsNewArr) == 0) {
				dlHighIsNew="-1"
			} else {
				dlHighIsNew=dlHighIsNewArr[adindex]
			}
		}

		if (materialListArr != nil&&materialListArr[0] != "-1") {
			if (len(materialListArr) == 0) {
				materialList="-1"
			} else {
				materialList=materialListArr[adindex]
			}
		}

		if (pkgBigimgListArr != nil&& pkgBigimgListArr[0] != "-1") {
			if (len(pkgBigimgListArr) == 0) {
				pkgBigimgList="-1"
			} else {
				pkgBigimgList=pkgBigimgListArr[adindex]
			}
		}

		if (pkgLanguageListArr != nil&&pkgLanguageListArr[0] != "-1") {
			if (len(pkgLanguageListArr) == 0) {
				pkgLanguageList="-1"
			} else {
				pkgLanguageList=pkgLanguageListArr[adindex]
			}
		}

		if (rlogidArr != nil) {
			if (len(rlogidArr) == 0&&rlogidArr[0] != "-1") {
				rlogid="NULL"
			} else {
				rlogid=rlogidArr[adindex]

			}
		}else {
			rlogid = "NULL"
		}


		adindexStr :=strconv.Itoa(adindex+1)
		s := []string{publish_time, report_time, time_stamp, income_type,
			string(server_v), child, token, lc, op_code_sub, sdk_vn, app_pkg,
			app_vn, app_vc, license, entry, level, key, string(pid), string(tid),
			string(gid), string(adid), adpkg, string(bid), string(gtype), string(sid), stype, string(pn), string(ps), logid,
			string(ac), string(coins), goid, anid, string(preclick), publish_ext, report_ext, ip, string(adindexStr),
			string(isNew), string(rcp), mdu, rlogid, publish_forms, string(bigImg), string(tcppType),
			string(dllv), string(svn), string(dlHighIsNew), hostname, sidtags, string(materialList), directgp, string(adid),
			string(ps_request), api_name, string(pkgBigimgList), string(pkgLanguageList), string(vs)}

		fmt.Println(strings.Join(s, "\t"))
		}
}
func getMergeTuple(Fields []string) (logid string, publish_time string, op_code string, publish_json string, sid string, stype string, server_v string, income_type string, report_time string, time_stamp string, app_pkg string, child string, token string, lc string, sdk_vn string, app_vc string, app_vn string, license string, key string, entry string, level string, pid string, tid string, gid string, gtype string, bid string, adid string, adpkg string, ac string, coins string, goid string, anid string, publish_ext string, report_ext string, ip string, mdu string, publish_forms string, dllv string, svn string, hostname string, sidtags string, directgp string, ps_request string, api_name string, vs string) {
	return Fields[0],  Fields[1],  Fields[2],  Fields[3],  Fields[4],  Fields[5],  Fields[6],  Fields[7],  Fields[8],  Fields[9],  Fields[10],  Fields[11],  Fields[12],  Fields[13],  Fields[14],  Fields[15],  Fields[16],  Fields[17],  Fields[18],  Fields[19],  Fields[20],  Fields[21],  Fields[22],  Fields[23],  Fields[24],  Fields[25],  Fields[26],  Fields[27],  Fields[28],  Fields[29],  Fields[30],  Fields[31],  Fields[32],  Fields[33],  Fields[34],  Fields[35],  Fields[36],  Fields[37],  Fields[38],  Fields[39],  Fields[40],  Fields[41],  Fields[42],  Fields[43],  Fields[44]
}