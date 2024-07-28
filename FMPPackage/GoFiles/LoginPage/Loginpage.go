package gofiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LoginReqStruct struct {
	User_Id  string `json:"user_id"`
	Password string `json:"password"`
}

type UserStruct struct {
	Login_id         int    `json:"login_id"`
	Role             string `json:"role"`
	User_id          string `json:"user_id"`
	Login_History_id int    `json:"lhistory_id"`
}

type LoginRespStruct struct {
	User_Details UserStruct `json:"userdetails"`
	ErrMsg       string     `json:"errmsg"`
	Status       string     `json:"status"`
	Msg          string     `json:"msg"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GFLogin(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {

		var lLoginDetailRec LoginReqStruct
		var lfinalrespRec LoginRespStruct

		//Reading and Unmarshal

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			lfinalrespRec.ErrMsg = "GL01" + err.Error()
			lfinalrespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lLoginDetailRec)
			fmt.Println(lLoginDetailRec)
			if err != nil {
				lfinalrespRec.ErrMsg = "GL02" + err.Error()
				lfinalrespRec.Status = "E"
			} else {

				//Login Validation
				lUserDetailRec, err := LoginValidation(lLoginDetailRec)
				if err != nil {
					lfinalrespRec.ErrMsg = "GL03" + err.Error()
					lfinalrespRec.Status = "E"

				} else {
					fmt.Println(lUserDetailRec, "API")
					if lUserDetailRec.Login_id != 0 {

						lfinalrespRec.User_Details = lUserDetailRec
						lfinalrespRec.Status = "S"
						lfinalrespRec.Msg = "Login Successfull"
					} else {

						lfinalrespRec.Status = "E"
						lfinalrespRec.Msg = "Invalid Credentials"

					}
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

		fmt.Println("GFLogin(-)")
	}
}
