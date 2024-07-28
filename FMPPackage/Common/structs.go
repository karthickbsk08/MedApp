package common

type CommonRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type StockStruct struct {
	Medicine_Name string `json:"medicine_name"`
	Brand         string `json:"brand"`
}
