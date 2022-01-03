package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IsExistsFile(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err) && err == nil
}

func ReadInt64FromFile(path string) (int64, error) {
	line, err := readFirstLine(path)
	if err != nil {
		return 0, fmt.Errorf("read first line")
	}

	v, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int64: %w", err)
	}

	return v, nil
}

func ReadUInt64FromFile(path string) (uint64, error) {
	line, err := readFirstLine(path)
	if err != nil {
		return 0, fmt.Errorf("read first line")
	}

	v, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse uint64: %w", err)
	}

	return v, nil
}

func readFirstLine(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return "", fmt.Errorf("scan first line: %w", err)
		}

		return "", nil
	}

	return scanner.Text(), nil
}
