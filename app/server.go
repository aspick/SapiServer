package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "os"
    "io/ioutil"
    "runtime"
    "path"

    "crypto/rand"
    "math/big"
    "strconv"
)

func currentSourcePath() string {
    _, filename, _, _ := runtime.Caller(1)
    return path.Dir(filename)
}

func storagePath(sapiId string) string {
    return currentSourcePath() + "\\..\\audio_data\\" + sapiId
}

func createWaveWithSapi(message string, voiceIndex string, sapiId string) bool {
    // create message txt
    err := ioutil.WriteFile(storagePath(sapiId) + ".txt", []byte(message), os.ModePerm)
    if err != nil {
        // err handle
        log.Print("file output error, is there permission?")
        return false
    }

    // create wave
    cmd := exec.Command("CScript", "//nologo", currentSourcePath() +"\\create_sapi.js", storagePath(sapiId), voiceIndex);
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
    return true
}

func randomStr(length int) string {
    const base = 36
    size := big.NewInt(base)
    n := make([]byte, length)
    for i, _ := range n {
        c, _ := rand.Int(rand.Reader, size)
        n[i] = strconv.FormatInt(c.Int64(), base)[0]
    }
    return string(n)
}


func sapiHandle(w http.ResponseWriter, r *http.Request) {
    sapiId     := r.FormValue("sapi_id")
    message    := r.FormValue("message")
    voiceIndex := r.FormValue("voice_index")

    if message == "" || voiceIndex == "" {
        w.WriteHeader(400)
        fmt.Fprintf(w, "{\"error\":\"required params missing\"}")
        return
    }

    if sapiId == "" {
        sapiId = randomStr(10)
    }

    success := createWaveWithSapi(message, voiceIndex, sapiId)

    if success {
        http.ServeFile(w, r, storagePath(sapiId) + ".wav")
    }else{
        w.WriteHeader(500)
        fmt.Fprintf(w, "{\"error\":\"unable to create wave file\"}")
        return
    }

}

func voicesHandle(w http.ResponseWriter, r *http.Request) {
    cmd       := exec.Command("CScript", "//nologo", currentSourcePath() + "\\sapi_voices.js")
    output, _ := cmd.Output()

    fmt.Fprintf(w, string(output))
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, currentSourcePath() + "\\index.html")
}

func main(){
    http.HandleFunc("/sapi/create", sapiHandle)
    http.HandleFunc("/sapi/voices", voicesHandle)
    http.HandleFunc("/", indexHandle)

    if err := http.ListenAndServe(":9081", nil); err != nil {
        log.Fatal("ListenAndServe ", err)
    }
}
