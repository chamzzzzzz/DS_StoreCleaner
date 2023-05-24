package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

func main() {
	app := cli.App{}
	app.Name = "DS_StoreCleaner"
	app.Usage = "递归删除指定目录下的DS_Store文件"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "dir",
			Usage: "指定要递归删除DS_Store文件的目录路径",
		},
	}

	app.Action = func(c *cli.Context) error {
		dirPath := c.String("dir")
		if dirPath == "" {
			dirPath, _ = os.Getwd()
		}

		files := listDSStoreFiles(dirPath)
		if len(files) == 0 {
			fmt.Println("没有找到任何DS_Store文件。")
			return nil
		}

		for _, file := range files {
			fmt.Println(file)
		}
		fmt.Printf("是否确认删除以上共计%d个文件？(y/n): ", len(files))
		reader := bufio.NewReader(os.Stdin)
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)

		if confirm == "y" || confirm == "Y" {
			deleteDSStoreFiles(files)
			fmt.Println("删除操作完成。")
		} else {
			fmt.Println("取消删除操作。")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func listDSStoreFiles(dirPath string) []string {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == ".DS_Store" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录发生错误: %v\n", err)
	}
	return files
}

func deleteDSStoreFiles(files []string) {
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("删除文件 %s 发生错误: %v\n", file, err)
		}
	}
}
