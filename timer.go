package main

import (
	"fmt"
	"os"
	"time"
)

func Maintance_7200() {
	Timer := 0
	fmt.Print("Enter date (YYYY-MM-DD): ")
	var dateStr string
	fmt.Scanln(&dateStr)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	finish_time := date.Add(7200 * time.Hour)
	fmt.Println("Tahmini 7200 saat  bakım tarihi:", finish_time.Format("2006-01-02"))

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(finish_time)

		Timer++
		if date == finish_time {
			fmt.Println("Bakım zamanı geldi.")
			break
		} else {

			fmt.Printf("Days: %02d \n", timeRemaining.d)
			if Timer == 10 {
				break
			}

		}

	}
}

func main() {

	var Bakim string
	fmt.Println("giriş yapılan zaman:", time.Now())

	fmt.Println("a : 3600 saat bakımı \nb : 7200 saat bakımı")
	fmt.Printf("bakım türünü seçiniz: ")
	fmt.Scan(&Bakim)

	switch Bakim {

	case "b", "B":
		Maintance_7200()

		break

	}

}

type countdown struct {
	t int
	d int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))

	return countdown{
		t: total,
		d: days,
	}
}
