package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// go func() {
	f, _ := os.Open("US.txt")
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = '\t'

	count := 0
	total := 0

	for {
		record, err := r.Read()
		total++
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if len(record) >= 5 {
			lat, err := strconv.ParseFloat(record[4], 64)
			lon, err := strconv.ParseFloat(record[5], 64)
			if err != nil {
				fmt.Println(err)
				break
			}
			location := redis.GeoLocation{Name: record[1], Latitude: lat, Longitude: lon}

			err = client.GeoAdd("places", &location).Err()
			count++
			if err != nil {
				fmt.Println(err)
				// panic(err)
			}
		}
	}
	// }()

	// err := client.Set("key", "Chris Test", 0).Err()
	fmt.Println(count)
	fmt.Println(total)
	val, err := client.GeoHash("races", "Brighton").Result()
	if err != nil {
		panic(err)
	}
	fmt.Print(val)

}
