package bpetrainer

import (
	"encoding/json"
	"fmt"
	"os"
	"unsafe"

	"com.vincentcodes/bpe/trie"
	"com.vincentcodes/bpe/utils"
)

type Token string

func (t Token) String() string{
    return string(t)
}

type Trainer struct{
    Vocab *trie.Trie `json:"vocab"`
    
    // number of occurances that makes a pair eligible to become a vocab (per train basis)
    Threshold int `json:"threshold"`

    // in bytes
    NgramMax int `json:"ngramMax"`

    // not used
    EndOfToken string `json:"eot"`

    IgnoreSpaces bool `json:"ignoreSpaces"`

    DoLog bool  `json:"dolog"`
}

func New(doLog bool, threshold int, ngramMax int) *Trainer{
    return &Trainer{
        Vocab: trie.New(),

        Threshold: threshold,
        NgramMax: ngramMax,
        EndOfToken: "</eot>",

        IgnoreSpaces: true,
        DoLog: doLog,
    }
}
func NewDefault(doLog bool) *Trainer{
    return &Trainer{
        Vocab: trie.New(),

        Threshold: 2,
        NgramMax: 15,
        EndOfToken: "</eot>",

        IgnoreSpaces: true,
        DoLog: doLog,
    }
}

func LoadFromFile(fullFilePath string) *Trainer{
    data, err := os.ReadFile(fullFilePath)
    utils.PanicOnError(err)

    var trainer *Trainer = &Trainer{}
    err = json.Unmarshal(data, trainer)
    utils.PanicOnError(err)
    return trainer
}

func (trainer *Trainer) SaveToFile(fullFilePath string){
    jsonRaw, _ := json.Marshal(trainer)
    err := os.WriteFile(fullFilePath, jsonRaw, 0755)
    utils.PanicOnError(err)
}

func (trainer *Trainer) Train(sentence string){
    trainer.learnLetters(sentence)

    for true{
        pairedToken := trainer.findPairToken(sentence)
        if pairedToken == nil{
            break
        }
        
        newVocab := string(*pairedToken)
        if len(newVocab) >= trainer.NgramMax{
            break
        }
        trainer.Vocab.Add(newVocab)
        trainer.log("Learned '%s'", newVocab)
    }
}

func (trainer *Trainer) GetLearnedVocabs() []string{
    entries := trainer.Vocab.GetWords()
    result := []string{}
    for _, entry := range *entries{
        result = append(result, entry.Word)
    }
    return result
}

func (trainer *Trainer) learnLetters(sentence string){
    for _, intChar := range sentence{
        trainer.Vocab.Add(string(intChar))
    }
}

// returns The most frequent token pair (combined already)
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
    if frequency < trainer.Threshold{
        return nil
    }
    return &token
}

// returns The most frequent token pair (combined already) and 
//         its index list which contains the location from which 
//         the func. has found it
func (trainer *Trainer) findMostFrequentPairToken(tokens *[]Token) (Token, *[]int) {
    tokenToIndexListMap := make(map[Token](*[]int))
    var mostFrequentToken Token = ""
    maxFrequency := 0
    for i, token := range *tokens{
        if i >= len(*tokens)-1{
            break
        }
        nextToken := (*tokens)[i+1]
        if trainer.IgnoreSpaces && (utils.IsSpace[Token](token) || utils.IsSpace[Token](nextToken)){
            continue
        }
        if combinedTokensLength(token, nextToken) >= trainer.NgramMax{
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
    endI := utils.Min(len(sentence), trainer.NgramMax)
    tokens := []string{}
    for startI < len(sentence){
        if startI == endI{
            panic(fmt.Sprintf("IllegalState: Encountered an unknown word starting from index: %d\n", startI))
        }
        subword := sentence[startI:endI]
        if trainer.Vocab.Has(subword){
            tokens = append(tokens, subword)
            startI = endI
            endI = utils.Min(len(sentence), startI + trainer.NgramMax)
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