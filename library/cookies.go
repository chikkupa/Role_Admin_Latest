package library

import(
	"time"
	"net/http"
)

var expiration  = time.Now().Add(24 * time.Hour)

func PutCookie(w http.ResponseWriter, key string, value string){
	cookie := http.Cookie{Name: key, Value: value, Expires: expiration}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, key string) string {
	cookie, err := r.Cookie(key)
	if err != nil{
		return ""
	}
		
	return cookie.Value
}