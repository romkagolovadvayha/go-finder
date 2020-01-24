package finder

import (
	"log"
	"sync"
)

type Finder struct {
    MaxCountWorkers int
    countWorkers int
    SearchWord string  // Поисковое слово
    totalCountWord int
    ProcessedPutWG sync.WaitGroup
    ProcessedRenderWG sync.WaitGroup
}


/*
* Запуск процесса
*/
func (f *Finder) Start(str string) {
    if f.countWorkers < f.MaxCountWorkers {
		f.countWorkers++
        f.ProcessedPutWG.Add(1)
        go func() {
            defer f.ProcessedPutWG.Done()
            for task := range f.taskCn {
                f.renderCn <- task
            }
        }()
    }
	f.taskCn <- str
}

/*
* Пишем результат
*/
func (f *Finder) Render() {
    f.ProcessedRenderWG.Add(1)
	go func() {
        defer f.ProcessedRenderWG.Done()
		for t := range f.renderCn {
            log.Println("Count for ", t.Url, ":", t.CountWord)
		}
        log.Println("Total:", f.totalCountWord)
	}()
}

/*
* Стопаем все потоки
*/
func (f *Finder) StopWait() {
	close(f.taskCn)
	f.ProcessedPutWG.Wait()

	close(f.renderCn)
	f.ProcessedRenderWG.Wait()
}