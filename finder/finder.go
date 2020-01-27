package finder

import (
    "sync"
    "fmt"
)

type Finder struct {
    maxCountWorkers int
    countWorkers int
    searchWord string  // Поисковое слово
    totalCountWord int
    processedPutWG sync.WaitGroup
    processedRenderWG sync.WaitGroup
    taskCn chan string
    renderCn chan *Task
    mu sync.Mutex
}

func NewFinder(searchWord string) *Finder {
    f := new(Finder)
    f.maxCountWorkers = 10
    f.searchWord = searchWord
    f.taskCn =  make(chan string)
    f.renderCn = make(chan *Task)
    return f
}

/*
* Максимальное кол-во потоков
*/
func (f *Finder) SetMaxCountWorkers(countWorkers int) {
    f.maxCountWorkers = countWorkers
}

/*
* Запуск процесса
*/
func (f *Finder) Start(str string) {
    if f.countWorkers < f.maxCountWorkers {
        f.countWorkers++
        f.processedPutWG.Add(1)
        go func() {
            defer f.processedPutWG.Done()
            for url := range f.taskCn {
                t := NewTask(url, f.searchWord)
                countWordsFoundOnSite := t.GetCountWordsFoundOnSite();
                f.mu.Lock();
                f.totalCountWord += countWordsFoundOnSite
                f.mu.Unlock();
                f.renderCn <- t
            }
        }()
    }
    f.taskCn <- str
}

/*
* Пишем результат
*/
func (f *Finder) Render() {
    f.processedRenderWG.Add(1)
    go func() {
        defer f.processedRenderWG.Done()
        for t := range f.renderCn {
            t.Render()
        }
        fmt.Println("Total:", f.totalCountWord)
    }()
}

/*
* Стопаем все потоки
*/
func (f *Finder) StopWait() {
    close(f.taskCn)
    f.processedPutWG.Wait()

    close(f.renderCn)
    f.processedRenderWG.Wait()
}