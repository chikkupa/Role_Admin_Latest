package middleware

import(
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "strconv"
    "net/http"
    "time"
    "io"

    "../library"
)

var key = "key@login#dfghgfkjkldfgd2asdwwds"
var redirect_location = "/"

func SetLoggedIn(w http.ResponseWriter) {
    timestamp := time.Now().Unix()
    sTimestamp := strconv.FormatInt(timestamp, 10)

    // authToken, _ := encrypt([]byte(sTimestamp), []byte(key))

    library.PutCookie(w, "auth_token", string(sTimestamp))
}

func IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
    sAuthTimestamp := library.GetCookie(r, "auth_token")

    if sAuthTimestamp == ""{
        return false
    }

    // sTimestamp, _ := decrypt([]byte(authToken), []byte(key))

    auth_timestamp, _ := strconv.Atoi(sAuthTimestamp)

    timestamp := time.Now().Unix()

    if(timestamp - int64(auth_timestamp) > 30 * 60){
        return false;
    }

    SetLoggedIn(w)

    return true;
}

func Logout(w http.ResponseWriter, r *http.Request){
    library.PutCookie(w, "auth_token", "")

    http.Redirect(w, r, redirect_location, http.StatusSeeOther)
}

func RedirectIfNotLoggedIn(w http.ResponseWriter, r *http.Request) {
    if !IsLoggedIn(w, r) {
        http.Redirect(w, r, redirect_location, http.StatusSeeOther)
    }
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}
