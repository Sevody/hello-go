package greeting

import "time"

func IsAM() bool {
	localTime := time.Now()
	return localTime.Hour() <= 12
}

func IsAfternoon() bool {
	localTime := time.Now()
	return localTime.Hour() <=16
}

func IsEvening() bool{
	localTime := time.Now() 
	return localTime.Hour() <= 22
}