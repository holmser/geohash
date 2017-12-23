package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

func lineCounter() {
	file, _ := os.Open("US.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
		// result := strings.Split(fileScanner.Text(), "\t")
		// fmt.Println(result)
	}
	fmt.Println("number of lines:", lineCount)
}

// parseRecord takes a line of text and returns a pointer to a redis GeoLocation
// for insertion into the database
func parseRecord(record string) *redis.GeoLocation {
	res := strings.Split(record, "\t")
	lat, err := strconv.ParseFloat(res[4], 64)
	lon, err := strconv.ParseFloat(res[5], 64)
	if err != nil {
		fmt.Println(err)
	}
	location := redis.GeoLocation{
		Name:      res[1],
		Latitude:  lat,
		Longitude: lon}
	return &location
}

func main() {

	go lineCounter()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	file, _ := os.Open("US.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
		// result := strings.Split(fileScanner.Text(), "\t")
		loc := parseRecord(fileScanner.Text())
		// fmt.Println(loc)
		err := client.GeoAdd("places", loc).Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	val, err := client.GeoHash("races", "Brighton").Result()
	if err != nil {
		panic(err)
	}
	fmt.Print(val)

}
