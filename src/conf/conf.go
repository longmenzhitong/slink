package conf

import (
	"fmt"
	"log"
	"os"
	"slink/src/files"
	"strconv"
	"strings"
)

var Port string
var Domain string
var RdsAddr string
var RdsPswd string
var RdsDb int

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("get user home dir error: %v", err)
	}
	path := homedir + "\\slink.conf"

	if _, err = os.Stat(path); err != nil {
		// 如果配置文件不存在就创建一个
		_, err = os.Create(path)
		if err != nil {
			log.Fatalf("create %s error: %v", path, err)
		}
		return
	}

	// 如果配置文件存在则读取配置
	lines, err := files.ReadLinesFromPath(path)
	if err != nil {
		log.Fatalf("read lines from %s error: %v", path, err)
	}

	for _, line := range lines {
		i := strings.Index(line, ":")
		key := line[:i]
		val := line[i+1:]
		switch key {
		case "port":
			Port = val
		case "domain":
			Domain = val
		case "rdsaddr":
			RdsAddr = val
		case "rdspswd":
			RdsPswd = val
		case "rdsdb":
			RdsDb, err = strconv.Atoi(val)
			if err != nil {
				fmt.Printf("invalid redis db: %s", val)
			}
		}
	}

	check()
}

func check() {
	if Port == "" {
		Port = "8080"
	}
	fmt.Printf("%-25s: %s\n", "port", Port)

	if Domain == "" {
		Domain = "localhost:8080"
	}
	fmt.Printf("%-25s: %s\n", "domain", Domain)

	if RdsAddr == "" {
		log.Fatal("redis address not found")
	} else {
		fmt.Printf("%-25s: %s\n", "redis address", RdsAddr)
	}

	fmt.Printf("%-25s: %d\n", "redis db", RdsDb)
}
