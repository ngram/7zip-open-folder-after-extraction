package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("begin extract-open")
	fmt.Println(os.Args)

	targetDir := filepath.Base(os.Args[1])
	targetDirArray := strings.Split(targetDir, ".")
	fmt.Println(targetDirArray[0])
	path := filepath.Clean(filepath.Dir(os.Args[1]))
	fmt.Println(path)

	openPath := filepath.Join(path, targetDirArray[0])
	fmt.Println(openPath)

	err := os.Mkdir(openPath, 0755)
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command("C:\\Program Files\\7-Zip\\7zG.exe", "x", os.Args[1], strings.Join([]string{"-o", openPath}, ""))
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(output))

	cmd = exec.Command("explorer", strings.Join([]string{"/root", ",", openPath}, ""))
	output, err = cmd.Output()
	if err != nil {
		log.Println("err :", err)
	}
	log.Println(string(output))
	fmt.Println("end extract-open")
}
