package utils

import (
	"fmt"
	"log"
)

type printer struct {
	normalPrint bool
	logPrint    bool
}

var (
	Printer = &printer{}
)

func SetNormal() {
	Printer.normalPrint = true
}

func SetLog() {
	Printer.logPrint = true
}

func (p *printer) Print(msg string) {
	if p.normalPrint {
		fmt.Println(msg)
	}
}

func (p *printer) Log(msg string) {
	if p.logPrint {
		log.Println(msg)
	}
}
