package main

import (
	"fmt"
	"os"
	"time"
)

func cmdAdd(title string) {
	todos, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	todo := Todo{ID: nextID(todos), Title: title, CreatedAt: time.Now()}
	todos = append(todos, todo)
	if err := save(todos); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Printf("追加しました [%d] %s\n", todo.ID, todo.Title)
}

func cmdList() {
	todos, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if len(todos) == 0 {
		fmt.Println("TODOはありません")
		return
	}
	for _, t := range todos {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%s %d: %s\n", status, t.ID, t.Title)
	}
}

func cmdDone(id int) {
	todos, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	for i, t := range todos {
		if t.ID == id {
			todos[i].Done = true
			if err := save(todos); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
			fmt.Printf("完了しました [%d] %s\n", t.ID, t.Title)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "ID %d が見つかりません\n", id)
	os.Exit(1)
}

func cmdDelete(id int) {
	todos, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			if err := save(todos); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
			fmt.Printf("削除しました [%d] %s\n", t.ID, t.Title)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "ID %d が見つかりません\n", id)
	os.Exit(1)
}
