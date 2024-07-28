package loginhistory

import (
	"fmt"
	"medapp/FMPPackage/DBCONNECT"
)

func FetchLoginHistory() ([]LoginHistoryStruct, error) {

	// var lHistoryDetailsRec LoginHistoryStruct
	var lHistoryArrRec []LoginHistoryStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("LHFL01")
		return lHistoryArrRec, err
	} else {
		defer db.Close()

		sqlstring := `select concat(upper(substr(created_by,1,1)) ,substr(created_by,2)),mlh.login_date ,mlh.login_time ,nvl(mlh.logout_date,"Working"),nvl(mlh.logout_time,"Working") 
		from medapp_login_history mlh ;`

		rows, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("LHFL02", err.Error())
			return lHistoryArrRec, err
		} else {

			for rows.Next() {
				var lHistoryDetailsRec LoginHistoryStruct

				err := rows.Scan(&lHistoryDetailsRec.Created_by, &lHistoryDetailsRec.Login_date, &lHistoryDetailsRec.Login_time, &lHistoryDetailsRec.Logout_date, &lHistoryDetailsRec.Logout_time)
				if err != nil {
					fmt.Println("LHFL03")
					return lHistoryArrRec, err
				} else {
					lHistoryArrRec = append(lHistoryArrRec, lHistoryDetailsRec)
				}
			}
		}
	}

	return lHistoryArrRec, nil
}
