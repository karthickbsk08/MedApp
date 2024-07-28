package adduserpage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	common "medapp/FMPPackage/Common"
	"net/http"
)

type AddUserReqStruct struct {
	User_id         string `json:"user_id"`
	Password        string `json:"password"`
	Role            string `json:"role"`
	Creater_User_id string `json:"cuser_id"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("AAAddUser(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "PUT" {

		var lNewUserRec AddUserReqStruct

		var lfinalrespRec common.CommonRespStruct

		//Reading and Unmarshal

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			lfinalrespRec.ErrMsg = "AA01" + err.Error()
			lfinalrespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lNewUserRec)
			fmt.Println(lNewUserRec)
			if err != nil {
				lfinalrespRec.ErrMsg = "AA02" + err.Error()
				lfinalrespRec.Status = "E"
			} else {

				//Login Validation
				lDetails, err := AddingUser(lNewUserRec)
				if err != nil {
					lfinalrespRec.ErrMsg = "AA03" + err.Error()
					lfinalrespRec.Status = "E"

				} else {

					if lDetails != 0 {
						lfinalrespRec.Status = "S"
						lfinalrespRec.Msg = "User Added Successful"
					} else {
						lfinalrespRec.Msg = "User Already Exists"
						lfinalrespRec.Status = "E"
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
		fmt.Println("AAAddUser(-)")

	}
}
