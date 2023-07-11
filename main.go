package main

import (
	"flag"

	"com.vincentcodes/bpe/bpetrainer"
	"com.vincentcodes/bpe/utils"
)

func main(){
    textPtr := flag.String("text", "", "the text to be used by the trainer")
    trainPtr := flag.Bool("train", false, "train it")
    loadFilePtr := flag.String("load", "", "load trainer from json file (eg. './trainer.json')")
    saveFilePtr := flag.String("save", "./trainer.json", "save trainer to json file")

    verbosePtr := flag.Bool("v", false, "verbose output")
    thresholdPtr := flag.Int("threshold", 2, "num of occurance needed in order to be added to vocab (per train basis)")
    ngramMaxPtr := flag.Int("ngrammax", 15, "max vocab length in bytes")
    flag.Parse()

    if utils.IsStringSpace(*textPtr){
        flag.Usage()
        return
    }

    trainer := bpetrainer.New(*verbosePtr, *thresholdPtr, *ngramMaxPtr)
    if !utils.IsStringSpace(*loadFilePtr){
        utils.Info("Loading trainer from '%s'", *loadFilePtr)
        trainer = bpetrainer.LoadFromFile(*loadFilePtr)
    }

    if *trainPtr{
        utils.Info("Start training")
        trainer.Train(*textPtr)
        trainer.SaveToFile(*saveFilePtr)
        utils.Info("Saved trainer to '%s'", *saveFilePtr)
        utils.Info("Learned vocabs are shown here:")
        utils.PrintObject(trainer.GetLearnedVocabs())
        return
    }

    utils.Info("Do segmentation only")
    utils.PrintObject(*trainer.TokenizeSubwords(*textPtr))
}