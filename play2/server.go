package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

var SerialNumber = map[string]int{}
var Check = map[string]int{}
var Repeat = map[string]int{}

func savedata() {
	f, _ := os.Create("SerialNumber")
	t, _ := json.Marshal(SerialNumber)
	f.Write(t)
	f.Close()
	f, _ = os.Create("Check")
	t, _ = json.Marshal(Check)
	f.Write(t)
	f.Close()
	f, _ = os.Create("Repeat")
	t, _ = json.Marshal(Repeat)
	f.Write(t)
	f.Close()
}
func authentication(w http.ResponseWriter, r *http.Request) {

	engineid, _ := ioutil.ReadAll(r.Body)
	rep := map[string]string{}
	rep["engineID"] = string(engineid)
	if SerialNumber[string(engineid)] == 0 {
		rep["status"] = "0"
		rep["count"] = "0"
	} else if SerialNumber[string(engineid)] == 1 {
		rep["status"] = "1"
		rep["count"] = "0"
	} else if SerialNumber[string(engineid)] == 2 {
		SerialNumber[string(engineid)] = 1
		h := sha1.New()
		Check[string(engineid)] = 0
		rep["status"] = "2"
		h.Write([]byte(strconv.Itoa(0)))
		rep["count"] = hex.EncodeToString(h.Sum(nil))
	}
	fmt.Println(rep)
	savedata()
	t, _ := json.Marshal(rep)
	w.Write(t)
}
func checkauthority(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(data)
	resp := map[string]string{}
	req := map[string]string{}
	json.Unmarshal(data, &req)
	fmt.Println(req)

	//if Repeat[req["engineID"]] == 1 {
	//	resp["engineID"] = "-1"
	//	resp["Count"] = "-1"
	count, exist := Check[req["engineID"]]
	if !exist || count >= 50 {
		resp["engineID"] = "-1"
		resp["Count"] = "-1"
	} else {
		resp["engineID"] = req["engineID"]
		resp["Count"] = req["Count"]
		/*
			h := sha1.New()
			h.Write([]byte(strconv.Itoa(count)))
			fmt.Println("req Count:" + req["Count"])
			fmt.Println("count:" + hex.EncodeToString(h.Sum(nil)))
			if hex.EncodeToString(h.Sum(nil)) == req["Count"] {
				count = Check[resp["engineID"]]
				h = sha1.New()
				h.Write([]byte(strconv.Itoa(count)))
				resp["engineID"] = req["engineID"]
				resp["Count"] = hex.EncodeToString(h.Sum(nil))
			} else {
				Repeat[req["engineID"]] = 1
				resp["engineID"] = "-1"
				resp["Count"] = "-1"
			}
		*/
	}
	fmt.Println(resp)
	savedata()
	data, _ = json.Marshal(resp)
	w.Write(data)

}
func complete(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(data)
	resp := map[string]string{}
	req := map[string]string{}
	json.Unmarshal(data, &req)
	fmt.Println(req)
	count := Check[req["engineID"]]
	h := sha1.New()
	h.Write([]byte(strconv.Itoa(count)))
	fmt.Println("req Count:" + req["Count"])
	fmt.Println("count:" + hex.EncodeToString(h.Sum(nil)))

	fi, err := os.Stat("config")
	if err != nil {
		os.MkdirAll("config", 0777)
	} else if !fi.IsDir() {
		os.MkdirAll("config", 0777)
	}
	fi, err = os.Stat(path.Join("config", req["engineID"]))
	if err != nil {
		os.MkdirAll(path.Join("config", req["engineID"]), 0777)
	} else if !fi.IsDir() {
		os.MkdirAll(path.Join("config", req["engineID"]), 0777)
	}
	err = ioutil.WriteFile(path.Join("config", req["engineID"], time.Now().String()), []byte(req["Config"]), 0x664)
	if err != nil {
		fmt.Println(err.Error())
	}
	Check[req["engineID"]]++
	count = Check[resp["engineID"]]
	h = sha1.New()
	h.Write([]byte(strconv.Itoa(count)))
	resp["engineID"] = req["engineID"]
	resp["Count"] = hex.EncodeToString(h.Sum(nil))

	/*
		if Repeat[req["engineID"]] == 1 {
			resp["engineID"] = "-1"
			resp["Count"] = "-1"
		} else {
			count := Check[req["engineID"]]
			h := sha1.New()
			h.Write([]byte(strconv.Itoa(count)))
			fmt.Println("req Count:" + req["Count"])
			fmt.Println("count:" + hex.EncodeToString(h.Sum(nil)))
			if hex.EncodeToString(h.Sum(nil)) == req["Count"] {
				Check[req["engineID"]]++
				count = Check[resp["engineID"]]
				h = sha1.New()
				h.Write([]byte(strconv.Itoa(count)))
				resp["engineID"] = req["engineID"]
				resp["Count"] = hex.EncodeToString(h.Sum(nil))
			} else {
				Repeat[req["engineID"]] = 1
				resp["engineID"] = "-1"
				resp["Count"] = "-1"
			}
		}
	*/
	fmt.Println(resp)
	savedata()
	data, _ = json.Marshal(resp)
	w.Write(data)
}
func main() {
	b, err := ioutil.ReadFile("SerialNumber")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(b, &SerialNumber)
	}
	b, err = ioutil.ReadFile("Check")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(b, &Check)
	}
	b, err = ioutil.ReadFile("Repeat")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(b, &Repeat)
	}
	fmt.Println(SerialNumber)
	//SerialNumber["aaaaa"] = 2
	http.HandleFunc("/authentication", authentication)
	http.HandleFunc("/check", checkauthority)
	http.HandleFunc("/complete", complete)
	err = http.ListenAndServe(":50000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
