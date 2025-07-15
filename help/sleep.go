package help

import "time"

func Sleep(duration int) {
	if duration == 0 {
		return
	}
	time.Sleep(time.Duration(duration) * time.Second)
}
func MSleep(duration int) {
	if duration == 0 {
		return
	}
	time.Sleep(time.Duration(duration) * time.Millisecond)
}

func USleep(duration int) {
	if duration == 0 {
		return
	}
	time.Sleep(time.Duration(duration) * time.Microsecond)
}
