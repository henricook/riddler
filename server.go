package main

import (
    "bufio"
    "encoding/json"
    "github.com/emirpasic/gods/sets/hashset"
    "log"
    "net/http"
    "os"
)

// Library functions

// Error handling for file IO
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Endpoints
type PingResponse struct {
    Response string  `json:"response"`
}

func Ping(w http.ResponseWriter, req *http.Request) {
    result := PingResponse{ Response: "PONG" }

    w.Header().Set("Content-Type", "application/json")

    encoder := json.NewEncoder(w)

    encodeErr := encoder.Encode(&result)
    check(encodeErr)
}

// Pass around the 100K set
type HundredK struct {
    set *hashset.Set
}

type Check100kRequest struct {
    Value string  `json:"value"`
}

type Check100kResponse struct {
    Common bool  `json:"common"`
}

// Process request to check 100k list
func (hundredK *HundredK) Check100kServer(w http.ResponseWriter, req *http.Request) {

    decoder := json.NewDecoder(req.Body)
    var body Check100kRequest
    decodeErr := decoder.Decode(&body)
    check(decodeErr)

    result := Check100kResponse{ Common: hundredK.set.Contains(body.Value) }

    w.Header().Set("Content-Type", "application/json")

    encoder := json.NewEncoder(w)

    encodeErr := encoder.Encode(&result)
    check(encodeErr)
}

// Server
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

    http.HandleFunc("/ping", Ping)
    http.HandleFunc("/check-100k", hundredK.Check100kServer)
    listenErr := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
    if listenErr != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
