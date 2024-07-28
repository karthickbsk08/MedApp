package billentry

import (
	"fmt"
	"medapp/FMPPackage/DBCONNECT"
)

func GetBillStock1() ([]BillStockStruct1, error) {
	var lbillStockArr []BillStockStruct1

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEGB01", err.Error())
		return lbillStockArr, err
	} else {
		defer db.Close()

		sqlstring := `select mmm.medicine_master_id ,mmm.medicine_name , mmm.brand , ms.quantity , ms.unit_price from medapp_stock ms ,medapp_medicine_master mmm 
		where  mmm.medicine_master_id = ms.medicine_master_id;`

		rows, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("BEGB02", err.Error())
			return lbillStockArr, err
		} else {

			for rows.Next() {
				var lbillStockRec BillStockStruct1

				err := rows.Scan(&lbillStockRec.Medicine_id, &lbillStockRec.Medicine_Name, &lbillStockRec.Brand, &lbillStockRec.Quantity, &lbillStockRec.Unit_price)
				if err != nil {
					fmt.Println("BEGB03", err.Error())
					return lbillStockArr, err
				} else {
					lbillStockArr = append(lbillStockArr, lbillStockRec)
				}
			}
		}
	}

	return lbillStockArr, nil

}

func GetBillStock() ([]BillStockStruct, int, error) {
	var lbillStockArr []BillStockStruct
	var lvalue int

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEGB01", err.Error())
		return lbillStockArr, lvalue, err
	} else {
		defer db.Close()

		sqlstring := `select mmm.medicine_name, ms.quantity from medapp_stock ms ,medapp_medicine_master mmm 
		where  mmm.medicine_master_id = ms.medicine_master_id;`

		rows, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("BEGB02", err.Error())
			return lbillStockArr, lvalue, err
		} else {

			for rows.Next() {
				var lbillStockRec BillStockStruct

				err := rows.Scan(&lbillStockRec.Medicine_Name, &lbillStockRec.Quantity)
				if err != nil {
					fmt.Println("BEGB03", err.Error())
					return lbillStockArr, lvalue, err
				} else {
					lbillStockArr = append(lbillStockArr, lbillStockRec)
					lNum, err := FetchBillNo()
					if err != nil {
						fmt.Println("BEGB04", err.Error())
						return lbillStockArr, lvalue, err
					} else {
						lvalue = lNum
					}
				}
			}
		}
	}
	return lbillStockArr, lvalue, nil

}

func AddBillMedicine(pnewMedRec BillStockStruct) ([]MedStruct, error) {
	var lAddMedArr []MedStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEGB01", err.Error())
		return lAddMedArr, err
	} else {
		defer db.Close()

		sqlstring := `Select mmm.medicine_master_id, mmm.medicine_name
		,mmm.brand ,? ,(? * ms.unit_price) 
		from medapp_medicine_master mmm, medapp_stock ms
		where mmm.medicine_master_id =ms.medicine_master_id and
		mmm.medicine_name =? ;`

		rows, err := db.Query(sqlstring, pnewMedRec.Quantity, pnewMedRec.Quantity, pnewMedRec.Medicine_Name)
		if err != nil {
			fmt.Println("BEGB02", err.Error())
			return lAddMedArr, err
		} else {

			for rows.Next() {
				var lAddmedDetailsRec MedStruct

				err := rows.Scan(&lAddmedDetailsRec.Medicine_id, &lAddmedDetailsRec.Medicine_Name, &lAddmedDetailsRec.Brand, &lAddmedDetailsRec.Quantity, &lAddmedDetailsRec.Amount)
				if err != nil {
					fmt.Println("BEGB03", err.Error())
					return lAddMedArr, err
				} else {
					lAddMedArr = append(lAddMedArr, lAddmedDetailsRec)
				}
			}
		}
	}
	return lAddMedArr, nil

}
func InsertBillDetails(psaveDetail BillSaveReqStruct) (BillPriceStruct, error) {

	// var lvalue int
	var lBillPrice BillPriceStruct
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEIBD01")
		return lBillPrice, err

	} else {
		defer db.Close()
		for i := 0; i < len(psaveDetail.BillsArr); i++ {

			sqlstring := `insert into medapp_bill_details (bill_no,medicine_master_id,quantity,unit_price,amount,created_by,created_date,updated_by,updated_date) values (?,?,?,round(?,2),?,?,now(),?,now());`

			Unit_Price := psaveDetail.BillsArr[i].Amount / float64(psaveDetail.BillsArr[i].Quantity)
			fmt.Println("Unit_price :", Unit_Price)

			lBillPrice.Total_Price = lBillPrice.Total_Price + psaveDetail.BillsArr[i].Amount

			_, err := db.Exec(sqlstring, psaveDetail.Bill_No, psaveDetail.BillsArr[i].Medicine_id, psaveDetail.BillsArr[i].Quantity, Unit_Price, psaveDetail.BillsArr[i].Amount, psaveDetail.User_id, psaveDetail.User_id)

			if err != nil {
				fmt.Println("BEIBD02")
				return lBillPrice, err
			}
		}

		lBillPrice.Gst = (lBillPrice.Total_Price * 18) / 100
		// lBillPrice.Gst = 123
		lBillPrice.Net_Price = lBillPrice.Gst + lBillPrice.Total_Price
	}
	fmt.Println("Bill Price", lBillPrice)
	return lBillPrice, nil
}

func InsertBillMaster(psaveDetail BillSaveReqStruct, pbillprice BillPriceStruct) (int, error) {

	var lvalue int
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEIBM01")
		return lvalue, err

	} else {
		defer db.Close()

		sqlstring := `insert into medapp_bill_master (bill_no,bill_date,bill_amount,bill_gst,net_price,
				login_id,created_by,created_date,updated_by,updated_date) values (?,?,?,?,?,?,?,now(),?,now());`

		record, err := db.Exec(sqlstring, psaveDetail.Bill_No, psaveDetail.Bill_date, pbillprice.Total_Price,
			pbillprice.Gst, pbillprice.Net_Price, psaveDetail.Login_id, psaveDetail.User_id, psaveDetail.User_id)

		if err != nil {
			fmt.Println("BEIBM02")
			return lvalue, err
		} else {
			rNum, err := record.RowsAffected()
			if err != nil {
				fmt.Println("BEIBM03")
				return lvalue, err
			} else {
				// fmt.Println("Bill Master ID")
				// fmt.Println(int(rNum))
				if rNum != 0 {
					lNum, err := FetchBillNo()
					if err != nil {
						fmt.Println("BEIBM04")
						return lvalue, err
					} else {
						fmt.Println("Bill No Save", lNum)
						lDetails, err := UpdateBillStocks(psaveDetail)
						if err != nil {
							fmt.Println("BEIBM05")
							return lvalue, err
						} else {
							if lDetails != 0 {
								lvalue = lNum + 1
								fmt.Println("Updated Successfully")
							}
						}
					}
				}
			}

		}
	}
	return lvalue, nil
}

func FetchBillNo() (int, error) {
	var lvalue int

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEFBN01")
		return lvalue, err

	} else {
		defer db.Close()

		sqlstring := `select max(bill_no)
		from medapp_bill_master mbm;`

		record, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("BEFBN01")
			return lvalue, err

		} else {
			if record.Next() {
				err = record.Scan(&lvalue)
				if err != nil {
					fmt.Println("BEFBN01")
					return lvalue, err
				} else {
					fmt.Println("Bill Fetched Successfully")
				}
			}
		}

	}
	return lvalue, nil
}

func UpdateBillStocks(psaveDetail BillSaveReqStruct) (int, error) {
	var lvalue int
	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("BEUB01")
		return lvalue, err

	} else {
		defer db.Close()
		for i := 0; i < len(psaveDetail.BillsArr); i++ {

			sqlstring := `update medapp_stock 
			set quantity =quantity -?,updated_by =?,updated_date =now()
			where medicine_master_id = (select medicine_master_id 
			from medapp_medicine_master mmm 
			where medicine_name =? and brand=?);`

			record, err := db.Exec(sqlstring, psaveDetail.BillsArr[i].Quantity, psaveDetail.User_id, psaveDetail.BillsArr[i].Medicine_Name, psaveDetail.BillsArr[i].Brand)

			if err != nil {
				fmt.Println("BEUB02")
				return lvalue, err
			} else {
				lNum, err := record.RowsAffected()
				if err != nil {
					fmt.Println("BEUB03")
					return lvalue, err
				} else {
					lvalue = int(lNum)
				}

			}
		}

	}
	return lvalue, nil
}
