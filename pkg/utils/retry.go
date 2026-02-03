package utils

import "time"

// Retry 简单重试封装
func Retry(times int, fn func() error) error {
	var err error
	for i := 0; i < times; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(time.Millisecond * 100)
	}
	return err
}
