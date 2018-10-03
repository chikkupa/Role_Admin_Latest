package library

import(
    "net/http"
    "strconv"
    "log"
    "io/ioutil"
    "time"
    "bytes"
)

func ForceDownload(w http.ResponseWriter, r *http.Request, file string, download_name string) {
    downloadBytes, err := ioutil.ReadFile(file)

    if err != nil {
    	log.Print(err)
    }

    // set the default MIME type to send
    mime := http.DetectContentType(downloadBytes)

    fileSize := len(string(downloadBytes))

    // Generate the server headers
    w.Header().Set("Content-Type", mime)
    w.Header().Set("Content-Disposition", "attachment; filename=" + download_name + "")
    w.Header().Set("Expires", "0")
    w.Header().Set("Content-Transfer-Encoding", "binary")
    w.Header().Set("Content-Length", strconv.Itoa(fileSize))
    w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

    http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))
 }