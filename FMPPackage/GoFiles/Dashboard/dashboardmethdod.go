package dashboard

import (
	"fmt"
	"log"
	"medapp/FMPPackage/DBCONNECT"
)

func BillerTodaySales(plogin_id int) (float64, error) {
	var ltodaysales float64

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBBT01")
		return ltodaysales, err
	} else {
		defer db.Close()

		sqlstring := `select nvl(round(sum(mbm.net_price),2),0) as net_price
					  from medapp_bill_master mbm 
					  where login_id =? and 
					        bill_date =current_date();`

		rows, err := db.Query(sqlstring, plogin_id)
		if err != nil {
			fmt.Println("DBBT02")
			return ltodaysales, err
		} else {
			if rows.Next() {

				err := rows.Scan(&ltodaysales)
				if err != nil {
					fmt.Println("DBBT03")
					return ltodaysales, err
				} else {

				}

			}
		}
	}
	fmt.Println(ltodaysales)
	return ltodaysales, nil
}

func BillerYesterdaySales(plogin_id int) (float64, error) {
	var yesterdaysales float64

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBBY01")
		return yesterdaysales, err
	} else {
		defer db.Close()

		sqlstring := `select nvl(round(sum(mbm.net_price),2),0) from medapp_bill_master mbm
		where login_id =? and bill_date = (select date_add(current_date(),interval -1 day));`

		rows, err := db.Query(sqlstring, plogin_id)
		if err != nil {
			fmt.Println("DBBY02")
			return yesterdaysales, err
		} else {
			if rows.Next() {

				err := rows.Scan(&yesterdaysales)
				if err != nil {
					fmt.Println("DBBY03")
					return yesterdaysales, err
				}

			}
		}
	}
	fmt.Println(yesterdaysales)
	return yesterdaysales, nil
}

func SalesValue() (float64, float64, ErrorStruct) {
	fmt.Println("DBSalesValue(+)")
	var lTodaysales float64
	var lInventoryValue float64
	var Errorstr ErrorStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBSV01")
		Errorstr.ErrorCode = "DBSV01"
		Errorstr.ErrMsg = err.Error()
		return lTodaysales, lInventoryValue, Errorstr
	} else {
		defer db.Close()

		sqlstring := `select nvl(round(sum(mbm.net_price),2),0)
		from medapp_bill_master mbm where bill_date =current_date();`

		rows, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("DBSV02")
			Errorstr.ErrorCode = "DBSV02"
			Errorstr.ErrMsg = err.Error()
			return lTodaysales, lInventoryValue, Errorstr
		} else {
			if rows.Next() {

				err := rows.Scan(&lTodaysales)
				fmt.Println("Today sales", lTodaysales)

				if err != nil {
					fmt.Println("DBSV03")
					Errorstr.ErrorCode = "DBSV03"
					Errorstr.ErrMsg = err.Error()
					return lTodaysales, lInventoryValue, Errorstr

				} else {
					sqlstring1 := `select nvl(round(sum(ms.quantity *ms.unit_price),2),0)
					from medapp_stock ms;`

					rows1, err := db.Query(sqlstring1)

					if err != nil {
						fmt.Println("DBSV04")
						Errorstr.ErrorCode = "DBSV04"
						Errorstr.ErrMsg = err.Error()
						return lTodaysales, lInventoryValue, Errorstr

					} else {
						if rows1.Next() {
							err := rows1.Scan(&lInventoryValue)

							fmt.Println("Inven", lInventoryValue)

							if err != nil {
								fmt.Println("DBSV05")
								Errorstr.ErrorCode = "DBSV05"
								Errorstr.ErrMsg = err.Error()
								return lTodaysales, lInventoryValue, Errorstr
							}
						}
					}

				}

			}
		}
	}
	fmt.Println("DBSalesValue(-)")
	return lTodaysales, lInventoryValue, Errorstr

}

func DailySales() ([]DailySalesStruct, ErrorStruct) {
	fmt.Println("DBDailySales(+)")

	var lDailySalesArr []DailySalesStruct
	var lDailyDetailsRec DailySalesStruct
	var Errorstr ErrorStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBDS01")
		Errorstr.ErrorCode = "DBDS01"
		Errorstr.ErrMsg = err.Error()
		return lDailySalesArr, Errorstr
	} else {
		defer db.Close()

		sqlstring := `select nvl(round(sum(mbm.net_price),2),0) ,dayname(bill_date)
		from medapp_bill_master mbm
		where bill_date >= (current_date() - interval 6 day)
		group by bill_date ;`

		// sqlstring := `select SalesValue, nvl(round(mbm.net_price,2),0) ,dayname(SalesTable.SalesValue)
		// from (select current_date()-interval  6 day as SalesValue
		// 	union all
		// 	select current_date()-interval  5 day
		// 	union all
		// 	select current_date()-interval  4 day
		// 	union all
		// 	select current_date()-interval  3 day
		// 	union all
		// 	select current_date()-interval  2 day
		// 	union all
		// 	select current_date()-interval  1 day
		// 	union all
		// 	select current_date()
		// ) as SalesTable left join medapp_bill_master mbm on Date(mbm.bill_date) =SalesTable.SalesValue
		// group by SalesTable.SalesValue
		// order by SalesTable.SalesValue;`

		rows, err := db.Query(sqlstring)

		if err != nil {
			fmt.Println("DBDS02")
			Errorstr.ErrorCode = "DBDS02"
			Errorstr.ErrMsg = err.Error()
			return lDailySalesArr, Errorstr
		} else {
			var count = 0
			for rows.Next() {
				count++

				err := rows.Scan(&lDailyDetailsRec.Total_amount, &lDailyDetailsRec.Day)
				if err != nil {
					fmt.Println("DBDS03")
					Errorstr.ErrorCode = "DBDS03"
					Errorstr.ErrMsg = err.Error()
					return lDailySalesArr, Errorstr
				} else {
					lDailySalesArr = append(lDailySalesArr, lDailyDetailsRec)
				}

			}
			if count == 0 {
				lDailyDetailsRec.Total_amount = 0
				lDailyDetailsRec.Day = "No Data1"
				lDailySalesArr = append(lDailySalesArr, lDailyDetailsRec)
				fmt.Println(lDailySalesArr)
			}
		}
	}
	fmt.Println(lDailySalesArr)
	fmt.Println("DBDailySales(-)")
	return lDailySalesArr, Errorstr
}

func MonthlySales() ([]MonthlySalesStruct, ErrorStruct) {
	fmt.Println("DBMonthlySales(+)")

	var lMonthlySalesArr []MonthlySalesStruct
	var lMonthlyDetailsRec MonthlySalesStruct
	var Errorstr ErrorStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBMS01")
		Errorstr.ErrorCode = "DBMS01"
		Errorstr.ErrMsg = err.Error()
		return lMonthlySalesArr, Errorstr
	} else {
		defer db.Close()

		// sqlstring := `select bill_date,nvl(round(sum(mbm.net_price),2),0),
		// to_char(bill_date,'MON')
		// from medapp_bill_master mbm
		// where year(mbm.bill_date) = year(current_date())
		// group by month(bill_date);`

		sqlstring := `select nvl(round(sum(mbm.net_price),2),0),
		to_char(bill_date,'MON')
		from medapp_bill_master mbm 
		where mbm.bill_date between (mbm.bill_date -interval 12 month ) and current_date() 
		group by month(bill_date);`

		rows, err := db.Query(sqlstring)

		if err != nil {
			fmt.Println("DBMS02")
			Errorstr.ErrorCode = "DBMS02"
			Errorstr.ErrMsg = err.Error()
			return lMonthlySalesArr, Errorstr
		} else {
			var count = 0
			for rows.Next() {
				count++

				err := rows.Scan(&lMonthlyDetailsRec.Total_amount, &lMonthlyDetailsRec.Month)
				if err != nil {
					fmt.Println("DBMS03")
					Errorstr.ErrorCode = "DBMS03"
					Errorstr.ErrMsg = err.Error()
					return lMonthlySalesArr, Errorstr
				} else {
					lMonthlySalesArr = append(lMonthlySalesArr, lMonthlyDetailsRec)
				}

			}
			if count == 0 {
				lMonthlyDetailsRec.Total_amount = 0
				lMonthlyDetailsRec.Month = "No Data2"
				lMonthlySalesArr = append(lMonthlySalesArr, lMonthlyDetailsRec)
				fmt.Println(lMonthlySalesArr)
			}
		}
	}
	fmt.Println(lMonthlySalesArr)
	fmt.Println("DBMonthlySales(-)")
	return lMonthlySalesArr, Errorstr
}

func CurrentMonthPef() ([]CurrentMonthPefStruct, ErrorStruct) {
	fmt.Println("DBCurrentMonthPef(+)")

	var lCurrentMonArr []CurrentMonthPefStruct
	var lCurrentMonDetailsRec CurrentMonthPefStruct
	var Errorstr ErrorStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBCMP01")
		Errorstr.ErrorCode = "DBCMP01"
		Errorstr.ErrMsg = err.Error()
		return lCurrentMonArr, Errorstr
	} else {
		defer db.Close()

		sqlstring := `select nvl(round(sum(net_price),2),0),(select ml.user_id  from medapp_login ml where ml.login_id = mbm.login_id)
		from medapp_bill_master mbm
		where monthname(mbm.bill_date)=monthname(current_date())  and date_format(mbm.bill_date,'%m-%Y')= date_format(current_date() ,'%m-%Y')
		group by mbm.login_id ;`

		// sqlstring := `select nvl(Bill_Master.Net_price,0),Login.user_id
		// from (select ml.user_id,ml.login_id
		// from medapp_login ml
		// where ml.role ='Biller' ) as Login left join (select mbm.login_id ,round(sum(mbm.net_price),2) as Net_price
		// from medapp_bill_master mbm
		// where date_format(mbm.bill_date,'%m-%Y')= date_format(current_date() ,'%m-%Y')
		// group by mbm.login_id
		// order by mbm.bill_date ) as Bill_Master on Login.login_id=Bill_Master.login_id`

		rows, err := db.Query(sqlstring)

		if err != nil {
			fmt.Println("DBCMP02")
			Errorstr.ErrorCode = "DBCMP02"
			Errorstr.ErrMsg = err.Error()
			fmt.Println(Errorstr.ErrMsg)
			return lCurrentMonArr, Errorstr
		} else {
			var count = 0
			for rows.Next() {
				count++
				err := rows.Scan(&lCurrentMonDetailsRec.Total_amount, &lCurrentMonDetailsRec.User_id)
				if err != nil {
					fmt.Println("DBCMP03")
					Errorstr.ErrorCode = "DBCMP03"
					Errorstr.ErrMsg = err.Error()
					return lCurrentMonArr, Errorstr
				} else {
					lCurrentMonArr = append(lCurrentMonArr, lCurrentMonDetailsRec)
				}

			}
			if count == 0 {
				lCurrentMonDetailsRec.Total_amount = 0
				lCurrentMonDetailsRec.User_id = "No Data3"
				lCurrentMonArr = append(lCurrentMonArr, lCurrentMonDetailsRec)
				fmt.Println(lCurrentMonArr)
			}
		}
	}
	fmt.Println(lCurrentMonArr)
	fmt.Println("DBCurrentMonthPef(-)")
	return lCurrentMonArr, Errorstr
}

func TodaySalesPef() ([]TodaySalesPefStruct, ErrorStruct) {
	fmt.Println("DBTodaySalesPef(+)")

	var lTodaySalesArr []TodaySalesPefStruct
	var lTodaySalesPefDetailsRec TodaySalesPefStruct
	var Errorstr ErrorStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("DBTSP01")
		Errorstr.ErrorCode = "DBTSP01"
		Errorstr.ErrMsg = err.Error()
		return lTodaySalesArr, Errorstr
	} else {
		defer db.Close()

		sqlstring := `select nvl(round(sum(mbm.net_price),2),0) ,
		(select ml.user_id 
		 from medapp_login ml 
		where ml.login_id =mbm.login_id)
		from medapp_bill_master mbm 
		where bill_date = current_date() 
		group by mbm.login_id;`

		rows, err := db.Query(sqlstring)

		if err != nil {
			fmt.Println("DBTSP02")
			Errorstr.ErrorCode = "DBTSP02"
			Errorstr.ErrMsg = err.Error()
			return lTodaySalesArr, Errorstr
		} else {
			// log.Println("------------else----------------")
			var count = 0
			for rows.Next() {
				count++
				err := rows.Scan(&lTodaySalesPefDetailsRec.Total_amount, &lTodaySalesPefDetailsRec.User_id)
				log.Println("------------for----------------")

				if err != nil {
					fmt.Println("DBTSP03")
					Errorstr.ErrorCode = "DBTSP03"
					Errorstr.ErrMsg = err.Error()
					return lTodaySalesArr, Errorstr
				} else {
					log.Println("------------else2----------------")

					lTodaySalesArr = append(lTodaySalesArr, lTodaySalesPefDetailsRec)
				}
			}
			log.Println(count, "copyparthi---///////////////////")
			if count == 0 {
				lTodaySalesPefDetailsRec.Total_amount = 0
				lTodaySalesPefDetailsRec.User_id = "No Data4"
				lTodaySalesArr = append(lTodaySalesArr, lTodaySalesPefDetailsRec)
				fmt.Println(lTodaySalesArr)
			}
		}
	}
	fmt.Println(lTodaySalesArr)
	fmt.Println("DBTodaySalesPef(-)")
	return lTodaySalesArr, Errorstr
}

func FinalMGRDashboard() ManagerDashRespStruct {

	fmt.Println("DBFinalMGRDashboard(+)")
	var lManagerDashrespRec ManagerDashRespStruct

	lTodaysales, lInventory, Err := SalesValue()
	if Err.ErrMsg != "" {
		lManagerDashrespRec.Error = append(lManagerDashrespRec.Error, Err)
		lManagerDashrespRec.Status = "E"
	} else {

		lManagerDashrespRec.OAToday_sales = lTodaysales
		lManagerDashrespRec.Inventory_value = lInventory
		lManagerDashrespRec.Status = "S"
		lManagerDashrespRec.Msg = "Sales Value Received"
	}

	var lDailySales = make([]DailySalesStruct, 0)
	lDailySales, Err = DailySales()
	if Err.ErrMsg != "" {
		lManagerDashrespRec.Error = append(lManagerDashrespRec.Error, Err)
		lManagerDashrespRec.Status = "E"
	} else {
		if lDailySales != nil {
			lManagerDashrespRec.DailySalesArr = lDailySales
			lManagerDashrespRec.Status = "S"
			lManagerDashrespRec.Msg = "Daily sales Received"
		}
		// lManagerDashrespRec.DailySalesArr = lDailySales
		// lManagerDashrespRec.Status = "S"
		// lManagerDashrespRec.Msg = "Daily sales Received"
	}

	lMonthlySales, Err := MonthlySales()
	if Err.ErrMsg != "" {
		lManagerDashrespRec.Error = append(lManagerDashrespRec.Error, Err)
		lManagerDashrespRec.Status = "E"
	} else {

		lManagerDashrespRec.MonthlySalesArr = lMonthlySales
		lManagerDashrespRec.Status = "S"
		lManagerDashrespRec.Msg = "Monthly sales Received"
	}

	lCurrentMon, Err := CurrentMonthPef()
	if Err.ErrMsg != "" {
		lManagerDashrespRec.Error = append(lManagerDashrespRec.Error, Err)
		lManagerDashrespRec.Status = "E"
	} else {
		lManagerDashrespRec.CurrentMonthArr = lCurrentMon
		lManagerDashrespRec.Status = "S"
		lManagerDashrespRec.Msg = "Current Mon sales Received"
	}

	lTodaySalesPef, Err := TodaySalesPef()
	if Err.ErrMsg != "" {
		lManagerDashrespRec.Error = append(lManagerDashrespRec.Error, Err)
		lManagerDashrespRec.Status = "E"
	} else {
		lManagerDashrespRec.TodaySalesArr = lTodaySalesPef
		lManagerDashrespRec.Status = "S"
		lManagerDashrespRec.Msg = "Todaysalespef sales Received"
	}

	fmt.Println("DBFinalMGRDashboard(-)")
	// log.Println(lManagerDashrespRec)
	return lManagerDashrespRec
}
