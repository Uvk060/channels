package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Producer struct {
	OutChan chan int
}

func (p *Producer) getOutChan() <-chan int { //вернет канал из которого можно только читать
	return p.OutChan //поэтомузаписать сюда ничего не можем

}
func (p *Producer) produce() {
	for {
		time.Sleep(1 * time.Second)      //пусть пишет каждую сек
		rand.Seed(time.Now().UnixNano()) //чтобы ранд всегда давал новые! числа https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
		p.OutChan <- rand.Intn(10)       //рандомное чосло от 0 до 10
	}
}
func main() {
	p := Producer{
		OutChan: make(chan int, 10),
	}
	go p.produce()
	for i := range p.getOutChan() { //getOutChan - канал только для считывания
		fmt.Println("Got massage from chan ", i)

		if i == 5 { //когда рандомно выйдет 5 - останови цикл
			close(p.OutChan)
			fmt.Println("Done")
		}
	}
}
