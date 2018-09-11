package humanize

import (
	"fmt"
)

const (
	cUnit = int64(1024)
)

// Bytes convert bytes to human readable format
//          Bytes     	 Humanized
//              0            "0 B"
//             27           "27 B"
//            999          "999 B"
//           1000         "1000 B"
//           1023         "1023 B"
//           1024        "1.0 KiB"
//           1728        "1.7 KiB"
//  1855425871872        "1.7 TiB"
//  math.MaxInt64        "8.0 EiB
func Bytes(b int64) string {
	if b < cUnit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := cUnit, 0
	for n := b / cUnit; n >= cUnit; n /= cUnit {
		div *= cUnit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
