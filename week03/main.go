package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/webfortune", fortunehandler)

	http.ListenAndServe(":8080", nil)
}

func fortunehandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	items := []string{"大吉", "中吉", "吉", "凶"}
	result := items[rand.New(rand.NewSource(seed)).Intn(len(items))]

	fmt.Println("今日の運勢は", result, "です")
}
