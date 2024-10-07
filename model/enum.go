package model

var RunnerTypeMap = map[string]string{
	"可执行程序(linux)":   "linux",
	"可执行程序(windows)": "windows",
}

var PcMapToLabel = map[string]string{
	"linux":   "可执行程序(linux)",
	"windows": "可执行程序(windows)",
}
