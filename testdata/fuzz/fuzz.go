package fuzzing

import "github.com/gibbz00/glamour"

func Fuzz(data []byte) int {
	_, err := glamour.RenderBytes(data, "dark")
	if err != nil {
		return 0
	}
	return 1
}
