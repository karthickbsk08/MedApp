package stockentry

import (
	"fmt"
	common "medapp/FMPPackage/Common"
	"medapp/FMPPackage/DBCONNECT"
)

func InsertNewMedicine(pnewmed AddMedicineReqStruct) (common.StockStruct, error) {
	var lLastMed common.StockStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SEIM01")
		return lLastMed, err

	} else {
		defer db.Close()
		sqlstring := `insert into medapp_medicine_master 
		(medicine_name,brand,created_by,created_date, updated_by,updated_date)
		select ?,?,?,now(),?,now()
		where not exists (select medicine_master_id
		from medapp_medicine_master mmm
		where mmm.medicine_name=? and mmm.brand=?);`

		record, err := db.Exec(sqlstring, pnewmed.Medicine_Name, pnewmed.Brand, pnewmed.User_id, pnewmed.User_id, pnewmed.Medicine_Name, pnewmed.Brand)

		if err != nil {
			fmt.Println("SEIM02")
			return lLastMed, err
		} else {
			// lNum, err := record.RowsAffected()

			// if err != nil {
			// 	fmt.Println("SEIM03")
			// 	return lverifyvalue, err
			// } else {
			// 	lverifyvalue = int(lNum)
			// }

			newNum, err := record.LastInsertId()
			var Medicine_Master_id int = int(newNum)
			if err != nil {
				fmt.Println("SEIM04")
				return lLastMed, err
			} else {
				lDetails, err := TakeLastMed(Medicine_Master_id)
				if err != nil {
					fmt.Println("SEIM05")
					return lLastMed, err
				} else {
					lLastMed = lDetails
				}
			}
		}
	}
	return lLastMed, nil

}

func TakeLastMed(pLastMed_id int) (common.StockStruct, error) {
	var lLastMed common.StockStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SETLM01")
		return lLastMed, err

	} else {
		defer db.Close()
		sqlstring := `select mmm.medicine_name ,mmm.brand 
		from medapp_medicine_master mmm 
		where mmm.medicine_master_id =?`

		rows, err := db.Query(sqlstring, pLastMed_id)
		if err != nil {
			fmt.Println("SETLM02", err.Error())
			return lLastMed, err

		} else {
			if rows.Next() {
				err := rows.Scan(&lLastMed.Medicine_Name, &lLastMed.Brand)
				if err != nil {
					fmt.Println("SETLM03")
					return lLastMed, err

				}
			}
		}
	}
	return lLastMed, nil

}

func FetchStocks() ([]common.StockStruct, error) {
	var lStocksArr []common.StockStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SEFS01")
		return lStocksArr, err
	} else {
		defer db.Close()

		sqlstring := `select medicine_name ,brand
		from medapp_medicine_master mmm ;`

		rows, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("SEFS02", err.Error())
			return lStocksArr, err
		} else {

			for rows.Next() {
				var lStockDetailsRec common.StockStruct

				err := rows.Scan(&lStockDetailsRec.Medicine_Name, &lStockDetailsRec.Brand)
				if err != nil {
					fmt.Println("SEFS03")
					return lStocksArr, err
				} else {
					lStocksArr = append(lStocksArr, lStockDetailsRec)
				}
			}
		}
	}

	return lStocksArr, nil
}

func StockUpdate(pUpdateMedRec UpdateStockReqStruct) (int, error) {

	var lvalue int
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SESU01")
		return lvalue, err

	} else {
		defer db.Close()
		sqlstring := `update medapp_stock 
		set quantity=quantity +?,unit_price =?,updated_by =?,updated_date =now()
		where medicine_master_id in (select medicine_master_id
		from medapp_medicine_master mmm
		where mmm.medicine_name=? and mmm.brand=?);`

		record, err := db.Exec(sqlstring, pUpdateMedRec.Quantity, pUpdateMedRec.Unit_price, pUpdateMedRec.User_id, pUpdateMedRec.Medicine_Name, pUpdateMedRec.Brand)

		if err != nil {
			fmt.Println("SESU02")
			return lvalue, err
		} else {
			lhide, err := record.RowsAffected()
			if err != nil {
				fmt.Println("SESU03")
				return lvalue, err
			} else {
				lvalue = int(lhide)

			}
		}
	}
	return lvalue, nil

}

func StockInsert(pInsertStockRec UpdateStockReqStruct) (int, error) {
	var lvalue int
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SESI01")
		return lvalue, err

	} else {
		defer db.Close()
		sqlstring := `insert into medapp_stock (medicine_master_id,quantity,unit_price,created_by,created_date,
			updated_by,updated_date)
			select mmm.medicine_master_id,?,?,?,now(),?,now()
			from medapp_medicine_master mmm 
			where mmm.medicine_name =? and mmm.brand =?`

		record, err := db.Exec(sqlstring, pInsertStockRec.Quantity, pInsertStockRec.Unit_price, pInsertStockRec.User_id, pInsertStockRec.User_id, pInsertStockRec.Medicine_Name, pInsertStockRec.Brand)
		if err != nil {
			fmt.Println("SESI02")
			return lvalue, err
		} else {
			lhide, err := record.RowsAffected()
			if err != nil {
				fmt.Println("SESI03")
				return lvalue, err
			} else {
				lvalue = int(lhide)

			}
		}
	}
	return lvalue, nil
}
