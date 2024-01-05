package structs

import "time"

// Content home页填充结构体
type Content struct {
	Tag   string `json:"tag"`
	Title string `json:"title"`
	// 时间字段
	Time          time.Time `json:"time"`
	Content       string    `json:"content"`
	CommentInt    int       `json:"comment_int"`
	CollectionInt int       `json:"collection_int"`
	ClickInt      int       `json:"click_int"`
}

// Contents 定义数组类型
type Contents []Content

// PassageContent 文章详情结构体
type PassageContent struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
	Detail    string    `json:"detail"`
}

type Side struct {
	Tag   string   `json:"tag"`
	Items []string `json:"items"`
}

// LoginInfo 登陆信息结构体
type LoginInfo struct {
	UserName string  `json:"user_name"`
	Password string  `json:"password"`
	Error    []error `json:"error"`
}

type SingleSong struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Time   string `json:"time"`
	Album  string `json:"album"`
}

type Songs []SingleSong

type Api struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Comment string `json:"comment"`
}

type Apis []Api

// ProgressStatus 进度条
type ProgressStatus struct {
	Now  float64 `json:"now"`
	Year int     `json:"year"`
}
