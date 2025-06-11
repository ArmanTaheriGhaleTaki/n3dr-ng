package count

import (
	"encoding/csv"
	"os"

	"n3dr/internal/app/n3dr/connection"
)

type csvWriter struct {
	file   *os.File
	writer *csv.Writer
}

type Nexus3 struct {
	*connection.Nexus3
	CsvFile string
	Sort    bool
	writer  *csv.Writer
}
