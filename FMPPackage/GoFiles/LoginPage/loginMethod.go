package gofiles

import (
	"fmt"
	"medapp/FMPPackage/DBCONNECT"
)

func LoginValidation(ploginRec LoginReqStruct) (UserStruct, error) {

	var lUserDetailRec UserStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("GLV01")
		return lUserDetailRec, err
	} else {
		defer db.Close()

		sqlstring := `Select login_id,role,user_id
					from medapp_login
					where user_id =? and password =?;`

		rows, err := db.Query(sqlstring, ploginRec.User_Id, ploginRec.Password)
		if err != nil {
			fmt.Println("GLV02")
			return lUserDetailRec, err
		} else {
			if rows.Next() {

				err := rows.Scan(&lUserDetailRec.Login_id, &lUserDetailRec.Role, &lUserDetailRec.User_id)
				if err != nil {
					fmt.Println("GLV03")
					return lUserDetailRec, err
				} else {

					Login_History_id, err := LoginHistoryInsert(lUserDetailRec)
					if err != nil {
						fmt.Println("GLV03")
						return lUserDetailRec, err
					} else {
						lUserDetailRec.Login_History_id = Login_History_id
						fmt.Println("Login History is Inserted")
					}

				}

			}
		}
	}
	fmt.Println(lUserDetailRec)
	return lUserDetailRec, nil

}

func LoginHistoryInsert(puserdetail UserStruct) (int, error) {
	var lLogin_History_id int
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("GLHI01")
		return lLogin_History_id, err

	} else {
		defer db.Close()
		sqlstring := `insert into medapp_login_history (login_id,login_date,login_time,created_by,created_date,updated_by,updated_date) 
		values (?,now(),now(),?,now(),?,now());`

		record, err := db.Exec(sqlstring, puserdetail.Login_id, puserdetail.User_id, puserdetail.User_id)

		if err != nil {
			fmt.Println("GLHI02")
			return lLogin_History_id, err
		} else {
			lhide, err := record.LastInsertId()
			if err != nil {
				fmt.Println("GLHI03")
				return lLogin_History_id, err
			} else {
				lLogin_History_id = int(lhide)

			}
		}
	}
	return lLogin_History_id, nil

}

func DoLogout(PlogoutDetails LogoutReqStruct) (int, error) {
	var lvalue int
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("GFDL01")
		return lvalue, err

	} else {
		defer db.Close()
		sqlstring := `update medapp_login_history 
		set logout_date =now(),logout_time =now(),updated_by=?,updated_date=now()
		where login_history_id = ?;`

		record, err := db.Exec(sqlstring, PlogoutDetails.User_id, PlogoutDetails.Login_History_Id)
		if err != nil {
			fmt.Println("GFDL02")
			return lvalue, err
		} else {
			lNum, err := record.RowsAffected()
			if err != nil {
				fmt.Println("GFDL03")
				return lvalue, err
			} else {
				lvalue = int(lNum)
			}
		}
	}
	return lvalue, nil
}
