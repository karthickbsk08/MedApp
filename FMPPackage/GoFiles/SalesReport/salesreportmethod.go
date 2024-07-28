package salesreport

import (
	"fmt"
	"medapp/FMPPackage/DBCONNECT"
)

func FetchSalesReport(pdaterangeRec SalesReportReqStruct) ([]SalesReportStruct, error) {
	fmt.Println(pdaterangeRec)

	var lSalesArrRec []SalesReportStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SRFS01")
		return lSalesArrRec, err
	} else {
		defer db.Close()

		sqlstring := ` Select mbd.bill_no ,mbm.bill_date ,mmm.medicine_name ,mbd.quantity,mbd.amount 
		From medapp_bill_details mbd ,medapp_medicine_master mmm ,medapp_bill_master mbm 
		Where mbd.bill_no =mbm.bill_no  and mbd.medicine_master_id =mmm.medicine_master_id and mbm.bill_date between  ? and  ? ;`

		rows, err := db.Query(sqlstring, pdaterangeRec.From_date, pdaterangeRec.To_date)
		if err != nil {
			fmt.Println("SRFS02", err.Error())
			return lSalesArrRec, err
		} else {

			for rows.Next() {
				var lsalesDetailsRec SalesReportStruct

				err := rows.Scan(&lsalesDetailsRec.Bill_No, &lsalesDetailsRec.Bill_date, &lsalesDetailsRec.Medicine_Name, &lsalesDetailsRec.Quantity, &lsalesDetailsRec.Amount)
				if err != nil {
					fmt.Println("SRFS03")
					return lSalesArrRec, err
				} else {

					lSalesArrRec = append(lSalesArrRec, lsalesDetailsRec)
				}
			}
		}
	}

	return lSalesArrRec, nil
}
