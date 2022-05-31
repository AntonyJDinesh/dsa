package queue

import "testing"

func TestMovingAverage_Next(t *testing.T) {

}

func Test_openLock(t *testing.T) {
	testData := []struct {
		deadends []string
		target   string
	}{
		{
			[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888",
		},
	}

	for _, td := range testData {
		openLock(td.deadends, td.target)
	}
}
