package sqltocsv

import (
	"bufio"
	"os"
	"path"
	"regexp"
	"strings"
)

func TransToCsv(sqlFile string) (string, error) {
	fRead, err := os.Open(sqlFile)
	if err != nil {
		return "", err
	}
	defer fRead.Close()

	baseName := path.Base(sqlFile)
	ext := path.Ext(sqlFile)
	filename := strings.TrimSuffix(baseName, ext)

	csvFile := filename + ".csv"
	fWrite, err := os.Create(csvFile)
	if err != nil {
		return "", err
	}
	defer fWrite.Close()

	fs := bufio.NewScanner(fRead)

	reg, err := regexp.Compile(`INSERT INTO .* \(`)
	if err != nil {
		return "", err
	}

	first := true
	for fs.Scan() {
		strLine := fs.Text()
		strs := strings.Split(strLine, ") VALUES (")
		if first {
			header := reg.ReplaceAllString(strs[0], "")
			fWrite.WriteString(header + "\n")
			first = false
		}
		value := strings.Trim(strs[1], ");")
		fWrite.WriteString(value + "\n")
	}
	return csvFile, nil
}
