package utils

import "fmt"

func FormatBytes(n int64) string {
	units := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB"}

	var i int
	size := float64(n)

	for i = 0; size >= 1024 && i < len(units)-1; i++ {
		size /= 1024
	}

	return fmt.Sprintf("%.2f %s", size, units[i])
}
