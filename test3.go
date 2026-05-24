package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "os/exec"
    "strings"
)

func main() {
    http.HandleFunc("/exec", dangerousExec)
    http.HandleFunc("/db", unsafeQuery)
    http.ListenAndServe(":8080", nil)
}

func dangerousExec(w http.ResponseWriter, r *http.Request) {
    cmd := r.URL.Query().Get("cmd")
    // 危险：直接执行用户输入的命令
    out, _ := exec.Command("sh", "-c", cmd).Output()
    w.Write(out)
}

func unsafeQuery(w http.ResponseWriter, r *http.Request) {
    db, _ := sql.Open("mysql", "user:pass@/dbname")
    defer db.Close()

    id := r.URL.Query().Get("id")
    // 危险：SQL注入漏洞
    query := "SELECT name FROM users WHERE id = " + id
    rows, _ := db.Query(query)
    
    var name string
    rows.Next()
    rows.Scan(&name)
    w.Write([]byte(name))
}

func ignoredError() {
    file := "test.txt"
    // 错误：未处理返回值
    _ = strings.Contains(file, "test")
    // 错误：检查了错误但未处理
    cmd := exec.Command("ls", "-l")
    _, _ = cmd.Output()
}
