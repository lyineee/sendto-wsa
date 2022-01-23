package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"unicode"
)

func main() {
	log.Println("start transfer")
	filePath := os.Args[1]
	_, fileName := filepath.Split(filePath)
	ext := filepath.Ext(fileName)
	isChinese := false
	outFileName := randomString(10)
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
	if isChinese {
		log.Println("Change filename to Chinese")
		shellCmd := fmt.Sprintf("mv /sdcard/Documents/%s /sdcard/Documents/%s", outFileName+ext, fileName+ext)
		log.Println("Shell command is: " + shellCmd)
		output, err = exec.Command("adb", "shell", shellCmd).Output()
		handleErr(err)
		log.Println(output)
	}

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

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
