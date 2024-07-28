package adduserpage

import (
	"fmt"
	"medapp/FMPPackage/DBCONNECT"
)

func AddingUser(pnewuser AddUserReqStruct) (int, error) {
	var lvalue int

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("AAU01")
		return lvalue, err
	} else {
		defer db.Close()

		sqlstring := `insert into medapp_login 
		(user_id,password,role,created_by,created_date,updated_by,updated_date) 
		select ?,?,?,?,now(),?,now()
		where not exists (select login_id 
		from medapp_login ml 
		where ml.user_id=?);`

		record, err := db.Exec(sqlstring, pnewuser.User_id, pnewuser.Password, pnewuser.Role, pnewuser.Creater_User_id, pnewuser.Creater_User_id, pnewuser.User_id)
		if err != nil {
			fmt.Println("AAU02")
			return lvalue, err
		} else {
			Num, err := record.RowsAffected()
			if err != nil {
				fmt.Println("AAU03")
				return lvalue, err
			} else {
				lvalue = int(Num)
				fmt.Println(lvalue, "Records inserted")
			}
		}

	}
	return lvalue, nil
}
