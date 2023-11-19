package trie

import (
	"encoding/json"
	"fmt"
)

type Trie struct {
    RootNode  *TrieNode  `json:"root"`
    LongestWordLen int   `json:"maxlen"`
}

type TrieNode struct {
    Value     rune                `json:"v"`
    Counter   int                 `json:"c"`
    NextNodes map[rune]*TrieNode  `json:"next"`
    EndOfWord bool                `json:"ends"`
}

type TrieEntry struct {
    Word string

    // indicates how many times it has been inserted into the Trie
    Counter int
}

func New() *Trie {
    return &Trie{
        RootNode: newTrieNode(0),
        LongestWordLen: 0,
    }
}

func newTrieNode(value rune) *TrieNode{
    return &TrieNode{
        Value:     value,
        Counter:   1,
        NextNodes: make(map[rune]*TrieNode),
        EndOfWord: false,
    }
}

func (trie *Trie) Add(word string){
    currentNode := trie.RootNode
    var lastNode *TrieNode = currentNode
    if len(word) > trie.LongestWordLen{
        trie.LongestWordLen = len(word)
    }
    for _, char := range word{
        node := currentNode.NextNodes[char]
        if node != nil{
            lastNode = node
            node.Counter++
            currentNode = node
            continue
        }

        // create new entry for this char
        lastNode = newTrieNode(char)
        currentNode.NextNodes[char] = lastNode
        currentNode = currentNode.NextNodes[char]
    }
    lastNode.EndOfWord = true
}

func (trie *Trie) Has(word string) bool{
    currentNode := trie.RootNode
    var lastNode *TrieNode = currentNode
    for _, char := range word{
        node := currentNode.NextNodes[char]
        if node == nil{
            return false
        }
        currentNode = node
        lastNode = node
    }
    // full word
    if lastNode.EndOfWord{
        return true
    }
    return false
}

func (trie *Trie) LongestWordLength() int{
    return trie.LongestWordLen
}

func (trie *Trie) GetWords() *[]*TrieEntry{
    foundWords := &[]*TrieEntry{}
    for _, nextNode := range trie.RootNode.NextNodes{
        trie.walkNode(nextNode, foundWords, []rune{})
    }
    return foundWords
}

// walk through the Trie from a `node` and add words into `foundWords`
func (trie *Trie) walkNode(node *TrieNode, foundWords *[]*TrieEntry, currentString []rune){
    currentString = append(currentString, node.Value)

    if node.EndOfWord{
        *foundWords = append(*foundWords, &TrieEntry{
            Word: string(currentString),
            Counter: node.Counter,
        })
    }

    for _, nextNode := range node.NextNodes{
        trie.walkNode(nextNode, foundWords, currentString)
    }
}

// TODO: remove

// Helpers
func (trie *Trie) String() string{
    return fmt.Sprintf("{Trie 'rootNode': %s}", trie.RootNode)
}
func (node *TrieNode) String() string{
    jsonRaw, _ := json.Marshal(node)
    return fmt.Sprintf("{TrieNode 'value': '%s', 'nextNodes': %s}", string(node.Value), string(jsonRaw))
}
// func (node *TrieNode) MarshalJSON() ([]byte, error) {
//     return json.Marshal(map[string]interface{}{
//         "value": string(node.value),
//         "counter": node.counter,
//         "nextNodes": node.nextNodes,
//         "endOfWord": node.endOfWord,
//     })
// }