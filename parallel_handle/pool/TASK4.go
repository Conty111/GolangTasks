/*
Напишите собственный worker pool (пакет pool). Он будет обрабатывать задачи, реализующие интерфейс PoolTask:

type PoolTask interface {
	// Execute запускает выполнение задачи и возвращает nil,
	// либо возникшую ошибку.
	Execute() error
	// OnFailure будет обрабатывать ошибки, возникшие в Execute(), то есть
	// пул должен вызывать OnFailure в случае, если Execute возвращает ошибку.
	OnFailure(error)
}

Для этого создайте структуру MyPool, которая удовлетворяет следующему интерфейсу:

type WorkerPool interface {
	// Start подготавливает пул для обработки задач. Должен вызываться один раз
	// перед использованием пула. Очередные вызовы должны игнорироваться.
	Start()
	// Stop останавливает обработку в пуле. Должен вызываться один раз.
	// Очередные вызовы должны игнорироваться.
	Stop()
	// AddWork добавляет задачу для обработки пулом. Добавлять задачи
	// можно после вызова Start() и до вызова Stop().
	// Если на момент добавления в пуле нет
	// свободных ресурсов (очередь заполнена) -
	// эту функция ожидает их освобождения (либо вызова Stop).
	AddWork(PoolTask)
}

Код должен содержать конструктор для MyPool:

// NewWorkerPool возвращает новый пул
// numWorkers - количество воркеров
// channelSize - размер очереди ожидания
// В случае ошибок верните nil и описание ошибки

func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error){
// ваша реализация
}
*/

package pool

import (
	"fmt"
	"sync"
)

type PoolTask interface {
	// Execute запускает выполнение задачи и возвращает nil,
	// либо возникшую ошибку.
	Execute() error
	// OnFailure будет обрабатывать ошибки, возникшие в Execute(), то есть
	// пул должен вызывать OnFailure в случае, если Execute возвращает ошибку.
	OnFailure(error)
}

type MyPool struct {
	poolTasks   chan PoolTask
	wg          sync.WaitGroup
	channelSize int
	numWorkers  int
	isWorking   bool
	mu          sync.Mutex
}

func NewWorkerPool(numWorkers int, channelSize int) (*MyPool, error) {
	// ваша реализация
	if numWorkers < 1 || channelSize < 0 {
		return nil, fmt.Errorf("invalid argument")
	}
	return &MyPool{
		numWorkers:  numWorkers,
		mu:          sync.Mutex{},
		channelSize: channelSize,
	}, nil
}

func (p *MyPool) Start() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.isWorking {
		p.isWorking = true
		p.poolTasks = make(chan PoolTask, p.channelSize)
		p.wg = sync.WaitGroup{}
		p.wg.Add(p.numWorkers)
		for i := 0; i < p.numWorkers; i++ {
			go func() {
				defer p.wg.Done()
				for t := range p.poolTasks {
					p.wg.Add(1)
					err := t.Execute()
					if err != nil {
						t.OnFailure(err)
					}
					p.wg.Done()
				}
			}()
		}
	}
}

func (p *MyPool) Stop() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.isWorking {
		close(p.poolTasks)
		p.wg.Wait()
		p.isWorking = false
	}

}

func (p *MyPool) AddWork(task PoolTask) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.isWorking {
		p.poolTasks <- task
	}
}
