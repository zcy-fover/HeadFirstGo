package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	os "os"
	"text/template"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/addView", addViewHandler)
	http.HandleFunc("/guestbook/commit", commitHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	checkError(err)
	fmt.Println("服务启动。。。")
}

//func viewHandler(writer http.ResponseWriter, request *http.Request) {
//	placeholder := []byte("这瓜保熟吗")
//	_, err := writer.Write(placeholder)
//	checkError(err)
//}
//
//func viewHandler(writer http.ResponseWriter, request *http.Request) {
//	html, err := template.ParseFiles("src/webGo/guestbook/guestH5.html")
//	checkError(err)
//	err = html.Execute(writer, nil)
//	checkError(err)
//}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	strings := getStrings("src/webGo/guestbook/signature.txt")
	html, err := template.ParseFiles("src/webGo/guestbook/guestH5.html")
	checkError(err)
	guestbook := Guestbook{
		SignatureCount: len(strings),
		Signatures:     strings,
	}
	err = html.Execute(writer, guestbook)
	checkError(err)
}

func addViewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("src/webGo/guestbook/addView.html")
	checkError(err)
	err = html.Execute(writer, nil)
	checkError(err)
}

func commitHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	//打开文件进行读写操作，如果文件存在进行追加，不存在就创建文件
	file, err := os.OpenFile("src/webGo/guestbook/signature.txt", options, 0600)
	checkError(err)
	_, err = fmt.Fprintln(file, signature)
	err = file.Close()
	checkError(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func getStrings(fileName string) []string {
	var lines []string //保存文件的每一行
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
