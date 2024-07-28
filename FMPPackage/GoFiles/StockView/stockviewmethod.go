package stockview

import (
	"fmt"
	"medapp/FMPPackage/DBCONNECT"
)

func FetchStockviewRecords() ([]StockviewStruct, error) {

	var lStocksArr []StockviewStruct

	db, err := DBCONNECT.LocalDBConnect()

	if err != nil {
		fmt.Println("SVFR01", err.Error())
		return lStocksArr, err
	} else {
		defer db.Close()

		sqlstring := `select mmm.medicine_name , mmm.brand , ms.quantity , ms.unit_price 
		from medapp_stock ms ,medapp_medicine_master mmm 
		where  mmm.medicine_master_id = ms.medicine_master_id ;`

		rows, err := db.Query(sqlstring)
		if err != nil {
			fmt.Println("SVFR02", err.Error())
			return lStocksArr, err
		} else {

			for rows.Next() {
				var lStockDetailsRec StockviewStruct

				err := rows.Scan(&lStockDetailsRec.Medicine_Name, &lStockDetailsRec.Brand, &lStockDetailsRec.Quantity, &lStockDetailsRec.Unit_price)
				if err != nil {
					fmt.Println("SVFR03", err.Error())
					return lStocksArr, err
				} else {
					lStocksArr = append(lStocksArr, lStockDetailsRec)
				}
			}
		}
	}

	return lStocksArr, nil
}
