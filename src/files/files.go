package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadLinesFromPath(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return ReadLinesFromFile(file)
}

func ReadLinesFromFile(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func WriteLinesToFile(file *os.File, lines *[]string) error {
	w := bufio.NewWriter(file)
	for _, line := range *lines {
		_, err := fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}

	return w.Flush()
}

func WriteLinesToPath(path string, lines *[]string) error {
	var file *os.File
	_, err := os.Stat(path)
	if err != nil {
		// 如果文件不存在则创建
		file, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		// 如果文件存在则打开
		file, err = os.Open(path)
		if err != nil {
			return err
		}
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return WriteLinesToFile(file, lines)
}
