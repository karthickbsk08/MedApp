package gofiles

import (
	"encoding/json"
	"fmt"
	common "medapp/FMPPackage/Common"
	"net/http"
)

type LogoutReqStruct struct {
	Login_History_Id int    `json:"login_history_id"`
	User_id          string `json:"user_id"`
}

func Logout(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GFLogout(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {

		var lLogoutReqRec LogoutReqStruct
		var lfinalrespRec common.CommonRespStruct

		//Reading and Unmarshal

		value := r.Header.Get("USER")

		err := json.Unmarshal([]byte(value), &lLogoutReqRec)
		if err != nil {
			lfinalrespRec.ErrMsg = "GFLG01" + err.Error()
			lfinalrespRec.Status = "E"
		} else {
			fmt.Println(lLogoutReqRec)
			lDetails, err := DoLogout(lLogoutReqRec)
			if err != nil {
				lfinalrespRec.ErrMsg = "GFLG01" + err.Error()
				lfinalrespRec.Status = "E"
			} else {
				if lDetails != 0 {
					lfinalrespRec.Status = "S"
					lfinalrespRec.Msg = "Logout Successfull"
				} else {
					lfinalrespRec.Status = "E"
					lfinalrespRec.Msg = "Logout Fail"
				}
			}
		}

		//UnMarshal
		datas, err := json.Marshal(lfinalrespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("GFLogout(-)")
	}
}
