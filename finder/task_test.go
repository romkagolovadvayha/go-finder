package finder

import (
    "testing"
)

func TestCountSubStr(t *testing.T) {
    cases := []struct {
        str string
        subStr string
        cntFindWord int
    }{
        {
            "Текст с содержанием одного слова Go", 
            "Go",
            1,
        },
        {
            "Go Текст с Go содержанием 5 слов Go Go Go", 
            "Go",
            5,
        },
        {
            "Go Текст с Go содержанием 9 слов Go Go Go Go Go Go Go", 
            "Go",
            9,
        },
    }
    for _, c := range cases {
        task := Task {}
        cnt := task.CountSubStr([]byte(c.str), c.subStr)
        if cnt != c.cntFindWord {
            t.Errorf("task.CountSubStr(%q, %q) == %d, cntFindWord %d", c.str, c.subStr, cnt, c.cntFindWord)
        }
    }
}

func TestGetCountWordsFoundOnSite(t *testing.T) {
    cases := []struct {
        url string
        searchWord string
        countWord int
    }{
        {
            "https://golang.org", 
            "Go",
            20,
        },
        {
            "https://golang.org", 
            "code",
            5,
        },
    }
    for _, c := range cases {
        task := NewTask(c.url, c.searchWord)
        countFindWord := task.Run()
        if c.countWord != countFindWord {
            t.Errorf("task.Run(%q), url: %q, countWord: %d != task.CountWord: %d", c.searchWord, c.url, c.countWord, countFindWord)
        }
    }
}
