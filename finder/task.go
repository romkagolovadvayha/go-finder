package finder

import (
    "net/http"
    "io/ioutil"
    "bytes"
    "log"
    "fmt"
)

type Task struct {
    url string
    searchWord string
    countWord int
}

func NewTask(url, searchWord string) *Task {
    t := new(Task)
    t.url = url
    t.searchWord = searchWord
    return t
}

/*
* Старт задачи - получить кол-во найденых слов на сайте
*/
func (t *Task) GetCountWordsFoundOnSite() int {
    body := t.getBody()
    t.countWord = t.CountSubStr(body)
    return t.countWord
}

/*
* Получаем тело
*/
func (t Task) getBody() []byte {
    resp, err := http.Get(t.url)
    if err != nil {
        log.Println(err.Error())
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err.Error())
    }
    return body
}

/*
* Количество найденной строки в байтовом массиве
*/
func (t Task) CountSubStr(body []byte) int {
    return bytes.Count(body, []byte(t.searchWord));
}

/*
* Вывод в консоль
*/
func (t Task) Render() {
    fmt.Println("Count for " + t.url + ":", t.countWord)
}