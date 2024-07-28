import axios from "axios"

const BaseApiClient=axios.create({
    baseURL:"http://localhost:22881",
    withCredentials:false,
    headers:{
        "Accept":"application/json",
        "Content-Type":"application/json"
    }
})
const BaseApiClient1=axios.create({
    baseURL:"http://localhost:22881",
    withCredentials:true,
    headers:{
        "Accept":"application/json",
        "Content-Type":"application/json"
    }
})

export default {
    LoginValidationCall(body){
        return BaseApiClient.put('/login',body)
    },
    AddNewUser(body){
        return BaseApiClient.put('/adduser',body)
    },
    LoginHistory(){
        return BaseApiClient.get('/loginhistory')
    },
    StockView(){
        return BaseApiClient.get('/stockview')
    },
    SalesReport(body){
        return BaseApiClient.put('/salesreport',body)
    },
    AddMedicine(body){
        return BaseApiClient.put('/addmedicine',body)
    },
    StockEntryGetStocks(){
        return BaseApiClient.get('/getstock')
    },
    UpdateStock(body){
        return BaseApiClient.put('/updatestock',body)
    },
    BillStock(){
        return BaseApiClient.get('/billstock')
    },
    AddBill(body){
        return BaseApiClient.put('/addbill',body)
    },
    SaveBill(body){
        return BaseApiClient.put('/savebill',body)
    },
    Logout(input){
        const hdr={
            headers:{
                "USER":input,                 
               
            }
        }
        return BaseApiClient.get('/logout',hdr)
    },
    BillerDash(input){
        const hdr={
            headers:{
                "LOGINID":input
            }
        }
        return BaseApiClient.get('/billerdash',hdr)
    },
    ManagerDash(){
        return BaseApiClient.get('managerdashboard')
    },
    SetCookie(body){
        return BaseApiClient1.put('/set',body)
    },
    GetCookie(){
        return BaseApiClient1.get('/get')
    }
}