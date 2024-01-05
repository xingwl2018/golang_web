package handlers

import (
	"fmt"
	. "golang_web/structs"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
	"unicode"
)

var (
	indexHtmlPath    = "/simpleweb/templates/index.html"
	homeHtmlPath     = "/simpleweb/templates/home.html"
	passageHtmlPath  = "/simpleweb/templates/passage.html"
	loginHtmlPath    = "/simpleweb/templates/login.html"
	songHtmlPath     = "/simpleweb/templates/song.html"
	progressHtmlPath = "/simpleweb/templates/progress.html"
	currentPath, _   = os.Getwd()
)

// 构建 首页 handler函数
func home(writer http.ResponseWriter, req *http.Request) {
	var c Contents
	c = []Content{
		{
			Tag:           "Go",
			Title:         "How to learn Golang",
			Time:          time.Now(),
			Content:       "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
			CommentInt:    1,
			CollectionInt: 12,
			ClickInt:      100,
		},
		{
			Tag:           "Python",
			Title:         "How to learn Python",
			Time:          time.Now().Add(-24 * time.Hour),
			Content:       "Python is a programming language that lets you work quickly and integrate systems more effectively.",
			CommentInt:    2,
			CollectionInt: 34,
			ClickInt:      1000,
		},
		{
			Tag:           "Java",
			Title:         "How to learn Java",
			Time:          time.Now().Add(-24 * 2 * time.Hour),
			Content:       "Java is a general-purpose programming language that is class-based, object-oriented, and designed to have as few implementation dependencies as possible.",
			CommentInt:    3,
			CollectionInt: 124,
			ClickInt:      900,
		},
		{
			Tag:           "JavaScript",
			Title:         "How to learn JavaScript",
			Time:          time.Now().Add(-24 * 2 * 2 * time.Hour),
			Content:       "JavaScript often abbreviated as JS, is a high-level, interpreted programming language that conforms to the ECMAScript specification. JavaScript has curly-bracket syntax, dynamic typing, prototype-based object-orientation, and first-class functions.",
			CommentInt:    212,
			CollectionInt: 1224,
			ClickInt:      9030,
		},
	}
	// 获取：gopath
	//currentPath, _ := os.Getwd()
	// 读取模版
	temp := template.New("index.html")
	// 添加模版处理函数
	t := temp.Funcs(template.FuncMap{"timeHandle": timeHandle})
	// 添加解析文件
	t, err := t.ParseFiles(
		path.Join(currentPath, indexHtmlPath),
		path.Join(currentPath, homeHtmlPath),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 将数据填充到模版中
	err = t.Execute(writer, c)
	if err != nil {
		panic(err)
	}
}

// 文章详情
func passage(writer http.ResponseWriter, request *http.Request) {
	var content = struct {
		PassageContent
		Side
	}{
		PassageContent: PassageContent{
			Title:     "How to learn golang",
			CreatedAt: time.Now(),
			Author:    "Go Team",
			Detail: `The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.`,
		},
		Side: Side{
			Tag: "状态",
			Items: []string{
				"用户数: 62",
				"分享数: 27",
				"评论数: 19",
				"收藏数: 12",
			},
		},
	}

	temp := template.New("index.html")
	t := temp.Funcs(template.FuncMap{"time": timeFormat})
	t, err := t.ParseFiles(
		path.Join(currentPath, passageHtmlPath),
		path.Join(currentPath, indexHtmlPath),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(writer, content)
	if err != nil {
		panic(err)
	}
}

// 登陆逻辑处理
func login(writer http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles(
		path.Join(currentPath, indexHtmlPath),
		path.Join(currentPath, loginHtmlPath),
	)
	var lgInfo LoginInfo
	// get请求，则执行登陆页处理
	if req.Method == http.MethodGet {
		temp.Execute(writer, lgInfo)
		return
	}

	if err := req.ParseForm(); err != nil {
		return
	}
	// 获取请求中的用户名密码
	fmt.Printf("请求信息：%+v \n", req.PostForm)
	UserName := req.PostFormValue("username")
	Password := req.PostFormValue("password")

	lgInfo.UserName = UserName
	lgInfo.Password = Password

	// 匿名方法
	errFunc := func(values string) error {
		v := values
		if len(v) < 8 || len(v) == 0 {
			return fmt.Errorf("the length should be larger 8")
		}
		if unicode.IsNumber(rune(v[0])) {
			return fmt.Errorf("should not start number")
		}
		return nil
	}

	if errFunc(lgInfo.UserName) == nil && errFunc(lgInfo.Password) == nil {
		// 重定向到 home页
		http.Redirect(writer, req, "/home", http.StatusSeeOther)
		return
	} else {
		if errFunc(lgInfo.UserName) != nil {
			lgInfo.Error = append(lgInfo.Error, fmt.Errorf("username is not suitable"))
		}
		if errFunc(lgInfo.Password) != nil {
			lgInfo.Error = append(lgInfo.Error, fmt.Errorf("password is not suitable"))
		}
		log.Println(lgInfo)
		temp.Execute(writer, lgInfo)
	}
}

// 退出登陆逻辑
func logout(writer http.ResponseWriter, req *http.Request) {
	req.Header.Del("Authorization")
	// 重定向到登陆页
	http.Redirect(writer, req, "/login", http.StatusSeeOther)
}

func song(writer http.ResponseWriter, req *http.Request) {
	var ss Songs
	ss = []SingleSong{
		{
			ID:     1,
			Name:   "全部都是你",
			Author: "DP龙猪",
			Time:   "03:23",
			Album:  "全部都是你",
		},
		{
			ID:     2,
			Name:   "对你的感觉",
			Author: "DP龙猪",
			Time:   "04:23",
			Album:  "对你的感觉",
		},
		{
			ID:     3,
			Name:   "我可不可以",
			Author: "DP龙猪",
			Time:   "05:23",
			Album:  "我可不可以",
		},
		{
			ID:     4,
			Name:   "围绕",
			Author: "DP龙猪",
			Time:   "03:23",
			Album:  "围绕",
		},
	}
	var aps Apis
	aps = []Api{
		{
			Title:   "获取人员信息",
			Content: "通过ID获取人员信息",
			Method:  http.MethodGet,
			Path:    fmt.Sprintf("/person/get?id=ID"),
			Comment: fmt.Sprintf("ID 选择 1 或者 2"),
		},
		{
			Title:   "获取所有人员信息",
			Content: "获取内置所有人员的信息",
			Method:  http.MethodGet,
			Path:    fmt.Sprintf("/persons"),
			Comment: fmt.Sprintf("无须传入请求参数"),
		},
		{
			Title:   "创建人员信息",
			Content: "传入参数 id 和 telephone 创建新人",
			Method:  http.MethodPost,
			Path:    fmt.Sprintf("/person/post"),
			Comment: fmt.Sprintf("传入参数 id 或者 telephone"),
		},
		{
			Title:   "更新人员信息",
			Content: "传入参数 id 更新人员 telephone 信息",
			Method:  http.MethodPatch,
			Path:    fmt.Sprintf("/person/patch?id=ID"),
			Comment: fmt.Sprintf("传入路径参数 id 和请求参数 telephone"),
		},
	}
	var all = struct {
		Songs Songs
		APis  Apis
	}{
		Songs: ss,
		APis:  aps,
	}
	temp, err := template.ParseFiles(
		path.Join(currentPath, indexHtmlPath),
		path.Join(currentPath, songHtmlPath),
	)
	if err != nil {
		log.Println(err)
		return
	}
	err2 := temp.Execute(writer, all)
	if err2 != nil {
		log.Println(err2)
		return
	}
}

func progress(writer http.ResponseWriter, req *http.Request) {
	var proStatus ProgressStatus
	monthDays := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	now := time.Now()
	y, m, d := now.Date()
	ok := 0
	sum := monthDays[time.Month(m)-1] + d
	if (y%400 == 0) || ((y%4 == 0) && (y%100 != 0)) {
		ok = 1
	}
	if (ok == 1) && (time.Month(m) > 2) {
		sum += 1
	}
	proStatus.Now, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(sum)/float64(365)*100), 64)
	proStatus.Year = y
	temp, _ := template.ParseFiles(
		path.Join(currentPath, indexHtmlPath),
		path.Join(currentPath, progressHtmlPath),
	)
	temp.Execute(writer, proStatus)
}
