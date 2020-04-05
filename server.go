package main

import (
    "bufio"
    "log"
    "net/http"
    "os"
    "github.com/emirpasic/gods/sets/hashset"
    "encoding/json"
)

// Pass around the 100K set
type HundredK struct {
    set *hashset.Set
}

type Check100kRequest struct {
    Value string
}

type Check100kResponse struct {
    Common bool
}

// Process request to check 100k list
func (hundredK *HundredK) Check100kServer(w http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)
    var body Check100kRequest
    decodeErr := decoder.Decode(&body)
    check(decodeErr)

    result := Check100kResponse { Common: hundredK.set.Contains(body.Value) }

    w.Header().Set("Content-Type", "application/json")

    encoder := json.NewEncoder(w)

    encodeErr := encoder.Encode(&result)
    check(encodeErr)
}

// Error handling for file IO
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// https://github.com/denji/golang-tls
func main() {


    log.Print("Loading resources...")

    hundredKFile, err := os.Open("ncsc-common-100k.txt")
    check(err)

    fileScanner := bufio.NewScanner(hundredKFile)
    fileScanner.Split(bufio.ScanLines)
    hundredKSet := hashset.New()

    for fileScanner.Scan() {
        hundredKSet.Add(fileScanner.Text())
    }

    hundredKFile.Close()

    hundredK := &HundredK{set: hundredKSet}

    log.Print("Resources Loaded.")

    // TODO: Check file existence and output nice error message
    log.Print("Starting Riddler Server...")

    http.HandleFunc("/check-100k", hundredK.Check100kServer)
    listenErr := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
    if listenErr != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}