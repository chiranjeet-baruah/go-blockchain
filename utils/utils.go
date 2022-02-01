package utils

import "time"

// GetEpochTime returns the current epoch time
func GetEpochTime() int64 {
	return time.Now().Unix()
}
