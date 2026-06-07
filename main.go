package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println(`使い方:
  todo add <タイトル>   TODOを追加
  todo list            一覧を表示
  todo done <ID>       完了にする
  todo delete <ID>     削除する`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "タイトルを指定してください")
			os.Exit(1)
		}
		cmdAdd(os.Args[2])
	case "list":
		cmdList()
	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "IDを指定してください")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "IDは数値で指定してください")
			os.Exit(1)
		}
		cmdDone(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "IDを指定してください")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "IDは数値で指定してください")
			os.Exit(1)
		}
		cmdDelete(id)
	default:
		usage()
		os.Exit(1)
	}
}
