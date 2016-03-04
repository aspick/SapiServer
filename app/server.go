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

    "crypto/md5"
    "encoding/hex"
)

func currentSourcePath() string {
    _, filename, _, _ := runtime.Caller(1)
    return path.Dir(filename)
}

func storagePath() string {
    return currentSourcePath() + "\\..\\audio_data";
}

func storageFilePath(sapiId string) string {
    return currentSourcePath() + "\\..\\audio_data\\" + sapiId
}

func getVoicesInfo() string {
    cmd       := exec.Command("CScript", "//nologo", currentSourcePath() + "\\sapi_voices.js")
    output, _ := cmd.Output()
    return string(output)
}

func md5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func createWaveWithSapi(message string, voiceIndex string, sapiId string) bool {
    // create message txt
    err := ioutil.WriteFile(storageFilePath(sapiId) + ".txt", []byte(message), os.ModePerm)
    if err != nil {
        // err handle
        log.Print("file output error, is there permission?")
        return false
    }

    // create wave
    cmd := exec.Command("CScript", "//nologo", currentSourcePath() +"\\create_sapi.js", storageFilePath(sapiId), voiceIndex);
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
    return true
}

func isExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func sapiHandle(w http.ResponseWriter, r *http.Request) {
    message    := r.FormValue("message")
    voiceIndex := r.FormValue("voice_index")

    if message == "" || voiceIndex == "" {
        w.WriteHeader(400)
        fmt.Fprintf(w, "{\"error\":\"required params missing\"}")
        return
    }

    sapiId := md5Hash(voiceIndex + message)

    if !isExist(storageFilePath(sapiId) + ".wav") {
        success := createWaveWithSapi(message, voiceIndex, sapiId)

        if !success {
            w.WriteHeader(500)
            fmt.Fprintf(w, "{\"error\":\"unable to create wave file\"}")
            return
        }
    }
    http.ServeFile(w, r, storageFilePath(sapiId) + ".wav")
}

func voicesHandle(w http.ResponseWriter, r *http.Request) {
    data := getVoicesInfo()

    fmt.Fprintf(w, data)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, currentSourcePath() + "\\index.html")
}

func initialize() {
    fmt.Println("start init")
    currentVoices := getVoicesInfo()

    // load prevoius voices
    fmt.Println("load previous voices")
    fi, fierr := os.Open("voices.json")
    defer fi.Close()
    previousVoices := ""
    if fierr != nil {
        fmt.Println("previous voices not detected")
    }else{
        fmt.Println("previous voices loaded")
        buf := make([]byte, 1024)
        fi.Read(buf)
        previousVoices = string(buf)
    }

    // compare voices
    if previousVoices != currentVoices {
        fmt.Println("current voices chagned, clear caches")
        // clear caches if voices updated
        err := os.RemoveAll(storagePath())
        if err != nil {
            panic(err)
        }

        err = os.Mkdir(storagePath(), 077)
        if err != nil {
            panic(err)
        }
    }

    // update latest voices
    fmt.Println("update voices")
    fo, err := os.Create("voices.json")
    if err != nil {
        panic(err)
    }
    defer fo.Close()

    fo.WriteString(getVoicesInfo())
}


func main(){
    initialize()

    http.HandleFunc("/sapi/create", sapiHandle)
    http.HandleFunc("/sapi/voices", voicesHandle)
    http.HandleFunc("/", indexHandle)

    if err := http.ListenAndServe(":9081", nil); err != nil {
        log.Fatal("ListenAndServe ", err)
    }
}
