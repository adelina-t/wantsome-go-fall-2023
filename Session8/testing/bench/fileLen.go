package bench

import "os"

func GetFileLen(name string, bufferSize int) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	size := 0
	for {
		buf := make([]byte, bufferSize)
		c, err := file.Read(buf)
		size += c
		if err != nil {
			break
		}
	}
	return size, nil
}
