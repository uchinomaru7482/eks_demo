package utils

import "time"

func InitTimeLocationAsiaTokyo() error {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	time.Local = loc
	return nil
}
