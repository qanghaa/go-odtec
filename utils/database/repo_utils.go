package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
)

// GeneratePlaceholders returns a string of "$1, $2, ..., $n".
func GeneratePlaceholders(n int) string {
	if n <= 0 {
		return ""
	}

	var builder strings.Builder
	sep := ", "
	for i := 1; i <= n; i++ {
		if i == n {
			sep = ""
		}
		builder.WriteString("$" + strconv.Itoa(i) + sep)
	}

	return builder.String()
}

func ScanAll[T Entity](rows pgx.Rows) ([]T, error) {
	var es = []T{}
	for rows.Next() {
		var e T
		_, values := e.FieldMap()
		err := rows.Scan(values)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		es = append(es, e)
	}
	return es, nil
}
