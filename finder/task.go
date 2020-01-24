package finder

import (
    "net/http"
    "io/ioutil"
    "bytes"
    "log"
)

type Task struct {
    Url string
    CountWord int
}

/*
* Старт задачи
*/
func (t *Task) Run(searchWord string) {
    body := t.getBody()
    t.CountWord = t.CountSubStr(body, searchWord)
}

/*
* Получаем тело
*/
func (t Task) getBody() []byte {
    resp, err := http.Get(t.Url)
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
func (t Task) CountSubStr(body []byte, text string) int {
    return bytes.Count(body, []byte(text));
}