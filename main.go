package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 執行 `git remote remove origin` 來移除原有的 Origin 遠端。
	exec.Command("git", "remote", "remove", "origin").Run()

	// 打開目前目錄底下的 `.gitremotes` 檔案。
	file, err := os.Open("./.gitremotes")
	if err != nil {
		log.Fatalf("gitremotes: %v", err)
	}
	defer file.Close()

	// 顯示目前的目錄。
	wd, _ := os.Getwd()
	fmt.Printf(">> %s\n", wd)

	first := true
	// 掃描 `.gitremotes` 的每一行文字。
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 取得該行的遠端路徑，並移除多餘的空白和換行。
		path := strings.TrimSpace(scanner.Text())

		if first {
			// 如果是第一次則執行 `git remote add origin <path>`。
			err := exec.Command("git", "remote", "add", "origin", path).Run()
			if err != nil {
				log.Fatalf("gitremotes: %v", err)
			}
		} else {
			// 如果是不是第一次則執行 `git remote set-url origin --add <path>`。
			err := exec.Command("git", "remote", "set-url", "origin", "--add", path).Run()
			if err != nil {
				log.Fatalf("gitremotes: %v", err)
			}
		}
		first = false
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("gitremotes: %v", err)
	}

	// 接著執行 `git remote -v` 來顯示最終的遠端結果。
	cmd := exec.Command("git", "remote", "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("gitremotes: %v\n%v", output, err)
	}
	fmt.Println(string(output))
}
