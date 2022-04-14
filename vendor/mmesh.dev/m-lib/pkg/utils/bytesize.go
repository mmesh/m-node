package utils

import "fmt"

type ByteSize float64

const (
	KB ByteSize = 1 << (10 * (iota + 1))
	MB
	GB
)

func (b ByteSize) String() string {
	switch {
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

/*
func test() {
	fmt.Println(1001*KB, 2.5*MB, 3.5*GB)
	fmt.Println(ByteSize(121000000))
}

Output:

1001.00KB 2.50MB 3.50GB
115.39MB
*/
