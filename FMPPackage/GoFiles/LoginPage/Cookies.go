package gofiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// func main() {
// 	// Start a web server with the two endpoints.
// 	log.Println("Server Started")
// http.HandleFunc("/set", setCookieHandler)
// 	http.HandleFunc("/get", getCookieHandler)
// 	http.ListenAndServe(":29081", nil)
// }

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	// SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

type CookieResp struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SetCookies(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	// Initialize a new cookie containing the string "kavin" and some
	// non-default attributes.

	// body,err:=ioutil.ReadAll(r.Body)
	if r.Method == "PUT" {
		var lfinalrespRec CookieResp
		var lLoginDetailRec LoginReqStruct

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))

		// fmt.Println(body)
		if err != nil {
			lfinalrespRec.ErrMsg = "GSCH01" + err.Error()
			lfinalrespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lLoginDetailRec)
			fmt.Println()
			if err != nil {
				lfinalrespRec.ErrMsg = "GSCH02" + err.Error()
				lfinalrespRec.Status = "E"
			} else {
				cookie := http.Cookie{
					Name:  lLoginDetailRec.User_Id,
					Value: lLoginDetailRec.Password,
					// Domain:   "http://192.168.2.188/",
					Path:     "/",
					MaxAge:   3600,
					HttpOnly: true,
					Secure:   false,
					// SameSite: http.SameSiteLaxMode,
				}
				fmt.Println(cookie)
				http.SetCookie(w, &cookie)

				log.Println("************************")
				log.Println(cookie)
				log.Println("Cookie set Successfully")
				lfinalrespRec.Status = "S"
				lfinalrespRec.Msg = "Cookie set Successfully"
			}

		}
		datas, err := json.Marshal(lfinalrespRec)
		log.Println(lfinalrespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

	}

	//Login Validation

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.

	// Write a HTTP response as normal.
	// w.Write([]byte("cookie set!"))
	fmt.Println("SetCookies(-)")
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	fmt.Println("GetCookies(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "http://192.168.2.162:8080")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	// cookie, err := r.Cookie("karthick")

	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, http.ErrNoCookie):
	// 		http.Error(w, "cookie not found", http.StatusBadRequest)
	// 	default:
	// 		log.Println(err)
	// 		http.Error(w, "server error", http.StatusInternalServerError)
	// 	}
	// 	return
	// } else {
	// 	fmt.Println("Cookiew", cookie)
	// 	fmt.Println("Cookie", cookie.Name, cookie.Value)
	// }

	cookies := r.Cookies() // Retrieve all cookies from the request

	for _, cookie := range cookies {
		fmt.Println("Cookie Name:", cookie.Name)
		fmt.Println("Cookie Value:", cookie.Value)
		fmt.Println("Cookie Domain:", cookie.MaxAge)
		fmt.Println("Cookie Path:", cookie.Path)
		// ... access other cookie attributes
	}

	// Echo out the cookie value in the response body.
	// w.Write([]byte(cookie.Value))
	fmt.Println("GetCookies(-)")

}
