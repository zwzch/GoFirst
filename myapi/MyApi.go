package myapi

import (
	//"encoding/base64"
	//"fmt"
	//"encoding/json"
	//"os"

	//"bufio"
	//"os"
	//"strings"

	"fmt"
	//"time"
	//"crypto/sha1"
	//"strconv"
	//"strconv"
	"strings"
	//"time"
	"net"
	"net/url"
)

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}
func MyApi(){
	/*base64*/
	//data := "abc123!?$*&()'-=@~"
	//sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	//fmt.Println(sEnc)
	////对byte数据编码
	//sDec,_ := base64.StdEncoding.DecodeString(sEnc)
	////解码
	//fmt.Println(string(sDec))

	//URLEncoding
	//uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	//fmt.Println(uEnc)
	//uDec,_ :=base64.URLEncoding.DecodeString(uEnc)
	//fmt.Println(string(uDec))

	/*json*/
	//bolB,_ := json.Marshal(true)
	//fmt.Println(string(bolB)
	//intB, _ := json.Marshal(1)
	//fmt.Println(string(intB))
	//
	//fltB, _ := json.Marshal(2.34)
	//fmt.Println(string(fltB))
	//
	//strB, _ := json.Marshal("gopher")
	//fmt.Println(string(strB))

	//slcD := []string{"apple", "peach", "pear"}
	//slcB, _ := json.Marshal(slcD)
	//fmt.Println(string(slcB))

	//mapD := map[string]int{"apple": 5, "lettuce": 7}
	//mapB, _ := json.Marshal(mapD)
	//fmt.Println(string(mapB))

	//res1D := &Response1{
	//	Page:   1,
	//	Fruits: []string{"apple", "peach", "pear"}}
	//res1B, _ := json.Marshal(res1D)
	//fmt.Println(string(res1B))
	//json序列化
	//res2D := &Response2{
	//	Page:   1,
	//	Fruits: []string{"apple", "peach", "pear"}}
	//res2B, _ := json.Marshal(res2D)
	//fmt.Println(string(res2B))


	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//
	//var dat map[string]interface{}
	//
	//if err := json.Unmarshal(byt, &dat); err != nil {
	//	panic(err)
	//}
	//fmt.Println(dat)

	//fmt.Println(num)

	//strs := dat["strs"].([]interface{})
	//str1 := strs[0].(string)
	//fmt.Println(str1)
	//
	//str := `{"page": 1, "fruits": ["apple", "peach"]}`
	//res := Response2{}
	//json.Unmarshal([]byte(str), &res)
	//fmt.Println(res)
	//fmt.Println(res.Fruits[0])

	//enc := json.NewEncoder(os.Stdout)
	//d := map[string]int{"apple": 5, "lettuce": 7}
	//enc.Encode(d)

	/*filter*/
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan(){
	//	ucl :=strings.ToUpper(scanner.Text())
	//	fmt.Println(ucl)
	//	}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "error:", err)
	//	os.Exit(1)
	//}
	//监听输入流只要有值就变成大写


	/*Random*/
	//fmt.Print(rand.Intn(100), ",")
	//fmt.Print(rand.Intn(100))
	//fmt.Println()

	//fmt.Println(rand.Float64())
	//fmt.Print((rand.Float64()*5)+5, ",")
	//fmt.Print((rand.Float64() * 5) + 5)
	//fmt.Println()
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//fmt.Println(r1.Intn(100))
	//fmt.Println(r1.Intn(100))
	//种子是随机的
	// If you seed a source with the same number, it
	// produces the same sequence of random numbers.

	//s2 := rand.NewSource(42)
	//r2 := rand.New(s2)
	//fmt.Print(r2.Intn(100), ",")
	//fmt.Print(r2.Intn(100))
	//fmt.Println()
	//s3 := rand.NewSource(42)
	//r3 := rand.New(s3)
	//fmt.Print(r3.Intn(100), ",")
	//fmt.Print(r3.Intn(100))


	/*正则表达式*/
	/*sha1*/
	//s := "sha1 this string"
	//h := sha1.New()
	//h.Write([]byte(s))
	//bs := h.Sum(nil)
	//fmt.Println(s)
	//fmt.Println(bs)
	//fmt.Printf("%x\n", bs)


	/*str*/
	//str_sample()
	//string_function_sample()
	//s:= "string"
	//var p  = fmt.Println
	//p("Contains:  ", s.Contains("test", "es"))
	//p("Count:     ", s.Count("test", "t"))
	//p("HasPrefix: ", s.HasPrefix("test", "te"))
	//p("HasSuffix: ", s.HasSuffix("test", "st"))
	//p("Index:     ", s.Index("test", "e"))
	//p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	//p("Repeat:    ", s.Repeat("a", 5))
	//p("Replace:   ", s.Replace("foo", "o", "0", -1))
	//p("Replace:   ", s.Replace("foo", "o", "0", 1))
	//p("Split:     ", s.Split("a-b-c-d-e", "-"))
	//p("ToLower:   ", s.ToLower("TEST"))
	//p("ToUpper:   ", s.ToUpper("test"))
	//p()
	//
	//p("Len: ", len("hello"))
	//p("Char:", "hello"[1])



	/*time*/
	//now := time.Now()
	//secs := now.Unix()
	//nanos := now.UnixNano()
	//fmt.Println(now)
	//millis := nanos / 1000000
	//fmt.Println(secs)
	//fmt.Println(millis)
	//fmt.Println(nanos)
	//fmt.Println(time.Unix(secs, 0))
	//fmt.Println(time.Unix(0, nanos))


	//tNow := time.Now()
	////时间转化为string，layout必须为 "2006-01-02 15:04:05"
	//timeNow := tNow.Format("2006-01-02 15:04:05")
	//fmt.Println("tNow(time format):", tNow)
	//fmt.Println("tNow(string format):", timeNow)
	//
	////string转化为时间，layout必须为 "2006-01-02 15:04:05"
	//t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	//fmt.Println("t(time format)", t)
	//
	////某个时间点 前后判断
	//trueOrFalse := t.After(tNow)
	//if trueOrFalse == true {
	//	fmt.Println("t（2014-06-15 08:37:18）在tNow之后!")
	//} else {
	//	fmt.Println("t（2014-06-15 08:37:18）在tNow之前!")
	//}
	//MyUrlParser()
}
func MyUrlParser() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

/*
* strings包函数使用。
*/
func string_function_sample() {
	//该函数代码收集自互联网

	fmt.Println("查找子串是否在指定的字符串中")
	fmt.Println(" Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true

	fmt.Println("")
	fmt.Println(" ContainsAny 函数的用法")
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false

	fmt.Println("")
	fmt.Println(" ContainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符

	fmt.Println("")
	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5 , 源码中有实现

	fmt.Println("")
	fmt.Println(" EqualFold 函数的用法")
	fmt.Println(strings.EqualFold("Go", "go")) //大小写忽略

	fmt.Println("")
	fmt.Println(" Fields 函数的用法")
	fmt.Println("Fields are: %q", strings.Fields(" foo bar baz ")) //["foo" "bar" "baz"] 返回一个列表

	//相当于用函数做为参数，支持匿名函数
	for _, record := range []string{" aaa*1892*122", "aaataat", "124|939|22"} {
		fmt.Println(strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case ch > '5':
				return true
			}
			return false
		}))
	}

	fmt.Println("")
	fmt.Println(" HasPrefix 函数的用法")
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) //前缀是以NLT开头的

	fmt.Println("")
	fmt.Println(" HasSuffix 函数的用法")
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) //后缀是以NLT开头的

	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 在存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中"))     // 在存在返回 6

	fmt.Println("")
	fmt.Println(" IndexAny 函数的用法")
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1

	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.IndexRune("NLT_abc", 'b')) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 's')) // 在存在返回 -1
	fmt.Println(strings.IndexRune("我是中国人", '中'))   // 在存在返回 6

	fmt.Println("")
	fmt.Println(" Join 函数的用法")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 返回字符串：foo, bar, baz

	fmt.Println("")
	fmt.Println(" LastIndex 函数的用法")
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3
	//最后一个索引
	fmt.Println("")
	fmt.Println(" LastIndexAny 函数的用法")
	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("我是中国人", "中"))      // 6

	fmt.Println("")
	fmt.Println(" Map 函数的用法")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	//map函数
	fmt.Println("")
	fmt.Println(" Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana

	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Println("")
	fmt.Println(" Split 函数的用法")
	fmt.Printf("%qn", strings.Split("a,b,c", ","))
	fmt.Printf("%qn", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%qn", strings.Split(" xyz ", ""))
	fmt.Printf("%qn", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println("")
	fmt.Println(" SplitAfter 函数的用法")
	fmt.Printf("%qn", strings.SplitAfter("/home/m_ta/src", "/")) //["/" "home/" "m_ta/" "src"]

	fmt.Println("")
	fmt.Println(" SplitAfterN 函数的用法")
	fmt.Printf("%qn", strings.SplitAfterN("/home/m_ta/src", "/", 2))  //["/" "home/m_ta/src"]
	fmt.Printf("%qn", strings.SplitAfterN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]

	fmt.Println("")
	fmt.Println(" SplitN 函数的用法")
	fmt.Printf("%qn", strings.SplitN("/home/m_ta/src", "/", 1))

	fmt.Printf("%qn", strings.SplitN("/home/m_ta/src", "/", 2))  //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%qn", strings.SplitN("/home/m_ta/src", "/", -1)) //["" "home" "m_ta" "src"]
	fmt.Printf("%qn", strings.SplitN("home,m_ta,src", ",", 2))   //["/" "home/" "m_ta/" "src"]

	fmt.Printf("%qn", strings.SplitN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]

	fmt.Println("")
	fmt.Println(" Title 函数的用法") //这个函数，还真不知道有什么用
	fmt.Println(strings.Title("her royal highness"))

	fmt.Println("")
	fmt.Println(" ToLower 函数的用法")
	fmt.Println(strings.ToLower("Gopher")) //gopher

	fmt.Println("")
	fmt.Println(" ToLowerSpecial 函数的用法")

	fmt.Println("")
	fmt.Println(" ToTitle 函数的用法")
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("loud 中国"))

	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", 2)) // aBaACEDF
	//第四个参数小于0，表示所有的都替换， 可以看下golang的文档
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", -1)) // aBaaCEDF

	fmt.Println("")
	fmt.Println(" ToUpper 函数的用法")
	fmt.Println(strings.ToUpper("Gopher")) //GOPHER

	fmt.Println("")
	fmt.Println(" Trim 函数的用法")
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]

	fmt.Println("")
	fmt.Println(" TrimLeft 函数的用法")
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung !!! ", "! ")) // ["Achtung !!! "]

	fmt.Println("")
	fmt.Println(" TrimSpace 函数的用法")
	fmt.Println(strings.TrimSpace(" tn a lone gopher ntrn")) // a lone gopher

}

func str_sample(){
		//str := "Hello"
		//str += "World"
		//fmt.Println(str)
		//fmt.Println(len(str))
		//fmt.Printf("获取对应下标字符： %c  %d \n", str[0], str[1])
		//r := []rune(str+"中国")
		//for i := 0; i < len(r); i++ {
		//	fmt.Println(string(r[i]), r[i])
		//}
		//n,err:=strconv.Atoi("12")
		//if err == nil{
		//	fmt.Printf("字符串转整数 : %d \n", n)
		//}
		//fmt.Println(strconv.Itoa(666))

		//fmt.Println([]byte("helloworld"))
		//fmt.Println("[]byte 转 字符串 : ", string([]byte{51, 52, 53, 54, 55, 56})) //可能存在unicode问题
		//fmt.Println("十六进制 : ", strconv.FormatInt(int64(28), 16))
		//fmt.Printf("二进制 : ")
		//fmt.Println(strconv.FormatInt(123, 2))
		////十六进制字符串转整数
		//n2, _ := strconv.ParseInt("be", 16, 32)
		//fmt.Printf("十六进制字符串转整数: %d  \n", int(n2))
		//fmt.Println(fmt.Sprintf("格式化字符串 0x%X", 62))
	}