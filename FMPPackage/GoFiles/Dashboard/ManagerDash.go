package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ManagerDash(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DBManagerDash(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "LOGINID,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {

		var lManagerDashrespRec ManagerDashRespStruct

		//Reading and Unmarshal
		lDetails := FinalMGRDashboard()
		lManagerDashrespRec = lDetails

		//UnMarshal
		datas, err := json.Marshal(lManagerDashrespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}
		fmt.Println("Final", lManagerDashrespRec)
		fmt.Println("DBManagerDash(-)")
	}
}
