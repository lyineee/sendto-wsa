package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path"
	"unicode"
)

func main() {
	log.Println("start transfer")
	filePath := os.Args[1]
	_, fileName := path.Split(filePath)
	ext := path.Ext(fileName)
	isChinese := false
	outFileName := "output"
	for _, char := range fileName {
		if unicode.Is(unicode.Han, char) {
			isChinese = true
			break
		}
	}
	if !isChinese {
		outFileName = fileName
	}
	output, err := exec.Command("adb", "push", filePath, "/sdcard/Documents/"+outFileName+ext).Output()
	handleErr(err)
	log.Println(string(output))

}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
		read := bufio.NewReader(os.Stdin)
		_, _, err := read.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
	}
}
