package main

type Topic struct {
	Name  string
	Words []string
}

var TOPICS = []Topic{
	{"Android", []string{"Android", "安卓"}},
	{"大数据", []string{"大数据", "数据挖掘", "数据分析"}},
	{"人工智能", []string{"人工智能", "机器学习", "自然语言处理"}},
	{"设计", []string{"设计", "交互", "UI", "UE", "美工"}},
	{"云计算", []string{"云计算", "分布式"}},
	{"实习/兼职", []string{"实习", "兼职"}},
	{"Web前端", []string{"前端", "h5", "html", "js", "javascript"}},
	{"创业", []string{"创业"}},
	{"产品", []string{"产品"}},
	{"PHP", []string{"PHP"}},
	{"iOS", []string{"iOS"}},
	{"C++", []string{"C++", "cpp"}},
	{"Java", []string{"Java"}},
	{"Python", []string{"Python"}},
	{"运营/市场", []string{"运营", "市场"}},
}
