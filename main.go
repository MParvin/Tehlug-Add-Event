package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Event struct {
	EVENT_TITLE   string
	CURRENT_DATE  string
	EVENT_NUMBER  int
	EVENT_DATE    string
	START_TIME    string
	END_TIME      string
	ADDRESS       string
	REGISTER_LINK string
	SPEAKERS      []string
	TOPICS        []string
	LAT           string
	LNG           string
	LOCATION      string
	DESCRIPTION   string
}

func main() {
	if _, err := os.Stat("./content/events"); os.IsNotExist(err) {
		log.Fatal("Directory ./content/events does not exist")
	}

	files, err := ioutil.ReadDir("./content/events")
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, file := range files {
		if file.IsDir() {
			num, err := strconv.Atoi(file.Name())
			if err == nil && num > max {
				max = num
			}
		}
	}

	eventNumber := max + 1
	fmt.Println("Event number:", eventNumber)

	now := time.Now()
	nowStr := now.Format("2006-01-02T15:04:05-07:00")
	fmt.Println("Current date:", nowStr)

	// Start getting data from user
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Event title( e.g. جلسه ۲۷۳): ")
	eventTitle, _ := reader.ReadString('\n')

	fmt.Print("Enter Event date (e.g. 1402-11-26): ")
	eventDate, _ := reader.ReadString('\n')

	fmt.Print("Enter start time (e.g. 19:00): ")
	startTime, _ := reader.ReadString('\n')

	fmt.Print("Enter end time (e.g. 21:00): ")
	endTime, _ := reader.ReadString('\n')

	fmt.Println("Enter speakers(comma-separated): ")
	speakersStr, _ := reader.ReadString('\n')
	speakersStr = speakersStr[:len(speakersStr)-1]
	speakers := strings.Split(speakersStr, ",")

	fmt.Print("Enter registration link(Evand link): ")
	registerLink, _ := reader.ReadString('\n')

	fmt.Println("Enter topics(comma-separated): ")
	topicsStr, _ := reader.ReadString('\n')
	topicsStr = topicsStr[:len(topicsStr)-1]
	topics := strings.Split(topicsStr, ",")

	fmt.Println("Enter event address: ")
	eventAddress, _ := reader.ReadString('\n')

	fmt.Print("Enter Latitude: ")
	lat, _ := reader.ReadString('\n')

	fmt.Print("Enter Longitude: ")
	lng, _ := reader.ReadString('\n')

	fmt.Println("Enter Event decription: ")
	eventDescription, _ := reader.ReadString('\n')

	fmt.Print("Enter Event location(e.g کدانو): ")
	eventLocation, _ := reader.ReadString('\n')

	event := Event{
		EVENT_TITLE:   eventTitle,
		CURRENT_DATE:  nowStr,
		EVENT_NUMBER:  eventNumber,
		EVENT_DATE:    eventDate,
		START_TIME:    startTime,
		END_TIME:      endTime,
		SPEAKERS:      speakers,
		REGISTER_LINK: registerLink,
		TOPICS:        topics,
		ADDRESS:       eventAddress,
		LAT:           lat,
		LNG:           lng,
		LOCATION:      eventLocation,
		DESCRIPTION:   eventDescription,
	}

	fmt.Println(event)

	eventDirectory := fmt.Sprintf("./content/events/%d", eventNumber)
	err = os.Mkdir(eventDirectory, 0755)
	if err != nil {
		panic(err)
	}

	eventFile := eventDirectory + "/info.md"

	tmpl, err := template.ParseFiles("templates/event.md")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(eventFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, event)
	if err != nil {
		panic(err)
	}

	fmt.Println("Event Created in ", eventFile)
	fmt.Println("Now, you can add this file using the following command: ")
	fmt.Println("git add ", eventFile)
	fmt.Println("And then commit it using the following command: ")
	fmt.Println("git commit -m \"Added event " + eventTitle + "\"")
	fmt.Println("And then push it using the following command: ")
	fmt.Println("git push")
}
