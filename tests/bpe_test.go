package tests

import (
	"testing"

	"com.vincentcodes/bpe/bpetrainer"
	"com.vincentcodes/bpe/utils"
)

func TestBpeEnglish(t *testing.T) {
    trainer := bpetrainer.NewDefault(true)
    trainer.Train(
        "I wanna test the code. Before that, I want to introduce you to the code. This is a trainer. The thingy used to train a model. " +
        "If you have the right training data, the test should go as planned, I think. I do really think.")
    utils.PrintObject(trainer.GetLearnedVocabs())

    // expect     "code", "test", "think", "train" -- occurance >= 2
    // not expect "data"                           -- occurance is 1
}

func TestBpeJapanese(t *testing.T) {
    trainer := bpetrainer.NewDefault(true)
    trainer.Train(
        "コードをテストしたいんですが、まずはこのコードを紹介してあげる。トレーナーなんです。モデルのトレーニングを行うための装置です。" +
        "適切なトレーニングデータさえあれば、テストも上手く行くはずですと思います。本当にそう思いますよ。")
    utils.PrintObject(trainer.GetLearnedVocabs())

    // expect     "コード", "トレーニ" (16 bytes), "テスト", "思います"
    // not expect "データ"
}