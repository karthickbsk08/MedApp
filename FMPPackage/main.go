package main

import (
	"log"
	adduserpage "medapp/FMPPackage/GoFiles/AddUserPage"
	billentry "medapp/FMPPackage/GoFiles/BillEntry"
	dashboard "medapp/FMPPackage/GoFiles/Dashboard"
	gofiles "medapp/FMPPackage/GoFiles/LoginPage"
	loginhistory "medapp/FMPPackage/GoFiles/Loginhistory"
	salesreport "medapp/FMPPackage/GoFiles/SalesReport"
	stockentry "medapp/FMPPackage/GoFiles/StockEntry"
	stockview "medapp/FMPPackage/GoFiles/StockView"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.Println("Server Started...")
	//setting cookies part
	http.HandleFunc("/set", gofiles.SetCookieHandler)
	// http.HandleFunc("/get", gofiles.GetCookieHandler)

	//login page
	http.HandleFunc("/login", gofiles.Login)

	//AddUser
	http.HandleFunc("/adduser", adduserpage.AddUser)

	//login history
	http.HandleFunc("/loginhistory", loginhistory.LoginHistory)

	//stock view
	http.HandleFunc("/stockview", stockview.Stocksview)

	//salesreport
	http.HandleFunc("/salesreport", salesreport.GetSalesReport)

	//stockentry
	http.HandleFunc("/addmedicine", stockentry.AddMedicine)
	http.HandleFunc("/getstock", stockentry.GetStocks)
	http.HandleFunc("/updatestock", stockentry.UpdateStock)

	//bill entry
	http.HandleFunc("/billstock", billentry.BillStock)
	http.HandleFunc("/addbill", billentry.AddBill)
	http.HandleFunc("/savebill", billentry.SaveBill)

	//logout
	http.HandleFunc("/logout", gofiles.Logout)

	//dashboard
	http.HandleFunc("/billerdash", dashboard.BillerDash)
	http.HandleFunc("/managerdashboard", dashboard.ManagerDash)

	http.ListenAndServe(":22881", nil)
}
