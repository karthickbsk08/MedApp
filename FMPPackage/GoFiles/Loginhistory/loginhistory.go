package loginhistory

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginHistoryStruct struct {
	Created_by  string `json:"created_by"`
	Login_date  string `json:"login_date"`
	Login_time  string `json:"login_time"`
	Logout_date string `json:"logout_date"`
	Logout_time string `json:"logout_time"`
}

type LoginHistoryRespStruct struct {
	LoginHistoryArr []LoginHistoryStruct `json:"login_historyarr"`
	ErrMsg          string               `json:"errmsg"`
	Status          string               `json:"status"`
	Msg             string               `json:"msg"`
}

func LoginHistory(w http.ResponseWriter, r *http.Request) {

	fmt.Println("LLoginHistory(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "GET" {

		var lhistoryRespRec LoginHistoryRespStruct

		//LoginHistory Fetch
		lDetails, err := FetchLoginHistory()
		if err != nil {
			lhistoryRespRec.ErrMsg = "LHL01" + err.Error()
			lhistoryRespRec.Status = "E"

		} else {

			if lDetails[0].Created_by != "" {
				lhistoryRespRec.LoginHistoryArr = lDetails
				lhistoryRespRec.Status = "S"
				lhistoryRespRec.Msg = "Logih History record received"
			} else {
				lhistoryRespRec.Msg = "Record not Found"
				lhistoryRespRec.Status = "E"
			}
		}

		//UnMarshal
		datas, err := json.Marshal(lhistoryRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}
		fmt.Println("LLoginHistory(-)")

	}
}
