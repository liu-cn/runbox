package dto

type ViewReq struct {
	FilePath string `json:"file_path"`
}

type SplitJoin struct {
	Data  string `json:"data"`
	Stp1  string `json:"stp1"`
	Stp2  string `json:"stp2"`
	Index string `json:"index"`
}
