package bpetrainer

import (
	"fmt"

	"com.vincentcodes/bpe/utils"
)

func (trainer *Trainer) log(format string, args ...any){
    if !trainer.DoLog{
        return
    }

    utils.Info(format, args...)
}

func (trainer *Trainer) println(args ...any){
    if !trainer.DoLog{
        return
    }

    fmt.Println(args...)
}