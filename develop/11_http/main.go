package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)
var calendar map[string]string

func main() {
	http.HandleFunc("/create_event", eventCreation)
	http.HandleFunc("/update_event", eventUpdate)
	http.HandleFunc("/delete_event", eventDelete)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)
	port := ":8080"
	calendar = make(map[string]string)
	err := http.ListenAndServe(port, nil)
	check(err)
}

type day struct{
	Day int `json:"day"`
	Month int `json:"month"`
	Year int `json:"year"`
	Event string `json:"event"`
}

type doc struct {
	Result string `json:"result"`
}

type jsError struct {
	Error string `json:"error"`
}

func eventCreation(w http.ResponseWriter,r *http.Request) {
	date, stringIndex := serveRequest(r)
	_, e := calendar[stringIndex]
	var out interface{}
	if !e {
		out = doc{"Success"}
		calendar[stringIndex] = date.Event
	} else {
		out = jsError{"Can't put on this date"}
	}
	js, err := json.MarshalIndent(out, "", "\n")
	if err != nil {
		log.Fatal("Json ", err.Error())
	}
	log.Print("Got query create event on", date)
	log.Print("Event calendar is now ", calendar)
	w.Write(js)
}

func eventUpdate(w http.ResponseWriter,r *http.Request) {
	date, stringIndex := serveRequest(r)
	var out interface{}
	if len(calendar[stringIndex]) > 0 {
		out = doc{"Success"}
		calendar[stringIndex] = date.Event
	} else {
		out = jsError{"Nothing to update"}
	}
	js, err := json.MarshalIndent(out, "", "\n")
	if err != nil {
		log.Fatal("Json ", err.Error())
	}
	log.Print("Got query update event on ", date)
	log.Print("Event calendar is now ", calendar)
	w.Write(js)
}

func eventDelete(w http.ResponseWriter,r *http.Request) {
	date, stringIndex := serveRequest(r)
	var out interface{}
	if len(calendar[stringIndex]) > 0 {
		out = doc{"Success"}
		delete(calendar,stringIndex)
	} else {
		out = jsError{"Nothing to delete"}
	}
	js, err := json.MarshalIndent(out, "", "\n")
	if err != nil {
		log.Fatal("Json ", err.Error())
	}
	log.Print("Got query delete event on", date)
	log.Print("Event calendar is now ", calendar)
	w.Write(js)
}

func serveRequest(r *http.Request) (day, string){
	event := r.FormValue("event")
	today, errDay := strconv.Atoi(r.FormValue("day"))
	if errDay != nil {
		log.Fatal("day ", errDay.Error())
	}
	month, errMonth := strconv.Atoi(r.FormValue("month"))
	if errMonth != nil {
		log.Fatal("Month ", errMonth.Error())
	}
	year, errYear := strconv.Atoi(r.FormValue("year"))
	if errYear != nil {
		log.Fatal("Year ", errYear.Error())
	}
	date := day{today, month, year, event}
	stringIndex := r.FormValue("day") + "-" + r.FormValue("month") + "-" + r.FormValue("year")
	return date, stringIndex
}

func eventsForDay(w http.ResponseWriter,r *http.Request) {
	date, stringIndex := serveRequest(r)
	date.Event = calendar[stringIndex]
	var out interface{}
	if len(calendar[stringIndex]) > 0 {
		out = doc{"Success"}
	} else {
		out = jsError{"There is no events on"}
	}
	js, err := json.MarshalIndent(out, "", "\n")
	if err != nil {
		log.Fatal("Json ", err.Error())
	}
	log.Print("Got query show events on", date)
	w.Write(js)
	js, err = json.MarshalIndent(date, "", "\n")
	if err != nil {
		log.Fatal("Json ", err.Error())
	}
	w.Write(js)
}

func eventsForWeek(w http.ResponseWriter,r *http.Request) {
	date, _ := serveRequest(r)
	tmp := date
	newStringIndex := ""
	for tmp.Day = date.Day; tmp.Day < date.Day + 7; tmp.Day++ {
		if tmp.Day > 31 {
			break
		}
		newStringIndex = strconv.Itoa(tmp.Day) + "-" + strconv.Itoa(tmp.Month) + "-" + strconv.Itoa(tmp.Year)
		tmp.Event = calendar[newStringIndex]
		var out interface{}
		if len(calendar[newStringIndex]) > 0 {
			out = doc{"Success"}
		} else {
			out = jsError{"There is no events on"}
		}
		js, err := json.MarshalIndent(out, "", "\n")
		if err != nil {
			log.Fatal("Json ", err.Error())
		}
		w.Write(js)
		js, err = json.Marshal(tmp)
		if err != nil {
			log.Fatal("Json ", err.Error())
		}
		w.Write(js)
	}
	log.Print("Got query show events in a week since", date)
}

func eventsForMonth(w http.ResponseWriter,r *http.Request) {
	date, _ := serveRequest(r)
	tmp := date
	newStringIndex := ""
	for tmp.Day = date.Day; tmp.Day < date.Day + 31; tmp.Day++ {
		if tmp.Day > 31 {
			break
		}
		newStringIndex = strconv.Itoa(tmp.Day) + "-" + strconv.Itoa(tmp.Month) + "-" + strconv.Itoa(tmp.Year)
		tmp.Event = calendar[newStringIndex]
		var out interface{}
		if len(calendar[newStringIndex]) > 0 {
			out = doc{"Success"}
		} else {
			out = jsError{"There is no events on"}
		}
		js, err := json.MarshalIndent(out, "", "\n")
		if err != nil {
			log.Fatal("Json ", err.Error())
		}
		w.Write(js)
		js, err = json.Marshal(tmp)
		if err != nil {
			log.Fatal("Json ", err.Error())
		}
		w.Write(js)
	}
	log.Print("Got query show events in a month since", date)
}

func check(e error) {
    if e != nil {
        log.Fatal(e.Error())
    }
}
