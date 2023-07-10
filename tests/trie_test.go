package tests

import (
	"testing"

	"com.vincentcodes/bpe/trie"
)

func setup() *trie.Trie{
    vocab := trie.New()
    vocab.Add("asdsa")
    vocab.Add("asdsa")
    vocab.Add("asdsa")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asd")
    vocab.Add("asb")
    return vocab
}

func TestTrieCounter(t *testing.T){
    vocab := setup()
    if !vocab.Has("asdsa"){
        t.Fail()
    }
    if !vocab.Has("asd"){
        t.Fail()
    }
    if !vocab.Has("asb"){
        t.Fail()
    }
}