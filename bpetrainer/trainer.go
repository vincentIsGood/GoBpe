package bpetrainer

import (
	"fmt"
	"unsafe"

	"com.vincentcodes/bpe/trie"
	"com.vincentcodes/bpe/utils"
)

type Token string

func (t Token) String() string{
    return string(t)
}

type Trainer struct{
    vocab *trie.Trie
    
    // number of occurances that makes a pair eligible to become a vocab
    threshold int

    // in bytes
    ngramMax int
    endOfToken string

    ignoreSpaces bool
}

func New() *Trainer{
    return &Trainer{
        vocab: trie.New(),
        threshold: 2,
        ngramMax: 15,
        ignoreSpaces: true,
        endOfToken: "</eot>",
    }
}

func (trainer *Trainer) Train(sentence string){
    trainer.learnLetters(sentence)

    trainer.findPairToken(sentence)
    for true{
        pairedToken := trainer.findPairToken(sentence)
        if pairedToken == nil{
            break
        }
        
        newVocab := string(*pairedToken)
        if len(newVocab) >= trainer.ngramMax{
            // TODO: fix the long repeating word problem
            break
        }
        trainer.vocab.Add(newVocab)
        // fmt.Printf("'%s'\n", newVocab)
    }
}

func (trainer *Trainer) GetLearnedVocabs() []string{
    entries := trainer.vocab.GetWords()
    result := []string{}
    for _, entry := range *entries{
        result = append(result, entry.Word)
    }
    return result
}

func (trainer *Trainer) learnLetters(sentence string){
    for _, intChar := range sentence{
        trainer.vocab.Add(string(intChar))
    }
}

func (trainer *Trainer) findPairToken(sentence string) (*Token){
    // eg. farmer, driver, mother, eat -> (? -> e)    most frequent 
    //                                 -> (e,? -> er) paired 
    //                                 -> Done
    // er -> 3; ea -> 1
    tokens := trainer.TokenizeSubwords(sentence)
    token, indexList := trainer.findMostFrequentPairToken(tokens)
    frequency := len(*indexList)
    // utils.PrintObject(tokens)
    // printIndexLocation(sentence, tokens, *indexList)
    if frequency < trainer.threshold{
        return nil
    }
    return &token
}

func (trainer *Trainer) findMostFrequentPairToken(tokens *[]Token) (Token, *[]int) {
    tokenToIndexListMap := make(map[Token](*[]int))
    var mostFrequentToken Token = ""
    maxFrequency := 0
    for i, token := range *tokens{
        if i >= len(*tokens)-1{
            break
        }
        nextToken := (*tokens)[i+1]
        if trainer.ignoreSpaces && (utils.IsSpace[Token](token) || utils.IsSpace[Token](nextToken)){
            continue
        }
        if combinedTokensLength(token, nextToken) >= trainer.ngramMax{
            continue
        }
        indexList := tokenToIndexListMap[token + nextToken]
        if indexList == nil{
            tokenToIndexListMap[token + nextToken] = &[]int{}
            indexList = tokenToIndexListMap[token + nextToken]
        }
        *indexList = append(*indexList, i)
        length := len(*indexList)
        if length > maxFrequency{
            maxFrequency = length
            mostFrequentToken = token + nextToken
        }
    }
    return mostFrequentToken, tokenToIndexListMap[mostFrequentToken]
}

func (trainer *Trainer) TokenizeSubwords(sentence string) *[]Token{
    // start from ngram size, shrink to smaller ngram (if not found)
    startI := 0
    endI := utils.Min(len(sentence), trainer.ngramMax)
    tokens := []string{}
    for startI < len(sentence){
        if startI == endI{
            panic(fmt.Sprintf("IllegalState: Encountered an unknown word starting from index: %d\n", startI))
        }
        subword := sentence[startI:endI]
        if trainer.vocab.Has(subword){
            tokens = append(tokens, subword)
            startI = endI
            endI = utils.Min(len(sentence), startI + trainer.ngramMax)
        }else{
            // shrink ngram size
            endI--
        }
    }
    return (*[]Token)(unsafe.Pointer(&tokens))
}

// utils
func combinedTokensLength(a Token, b Token) int{
    return len(a) + len(b)
}

func printIndexLocation(str string, tokens *[]Token, location *[]int){
    fmt.Println(str)
    finalStr := *utils.InitArray[byte](len(str), ' ')
    for _, loc := range *location{
        finalStr[accumulatedTokensLength(tokens, loc)-1] = byte('^')
    }
    fmt.Println(string(finalStr))
}

func accumulatedTokensLength(tokens *[]Token, index int) int{
    finalLength := 0
    for i, token := range *tokens{
        if i > index{
            break
        }
        finalLength += len(token)
    }
    return finalLength
}