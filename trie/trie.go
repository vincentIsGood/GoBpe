package trie

import (
	"encoding/json"
	"fmt"
)

type Trie struct {
    rootNode  *TrieNode
    longestWordLen int
}

type TrieNode struct {
    value     rune
    counter   int
    nextNodes map[rune]*TrieNode
    endOfWord bool
}

type TrieEntry struct {
    Word string

    // indicates how many times it has been inserted into the Trie
    Counter int
}

func New() *Trie {
    return &Trie{
        rootNode: newTrieNode(0),
        longestWordLen: 0,
    }
}

func newTrieNode(value rune) *TrieNode{
    return &TrieNode{
        value:     value,
        counter:   1,
        nextNodes: make(map[rune]*TrieNode),
        endOfWord: false,
    }
}

func (trie *Trie) Add(word string){
    currentNode := trie.rootNode
    var lastNode *TrieNode = currentNode
    if len(word) > trie.longestWordLen{
        trie.longestWordLen = len(word)
    }
    for _, char := range word{
        node := currentNode.nextNodes[char]
        if node != nil{
            lastNode = node
            node.counter++
            currentNode = node
            continue
        }

        // create new entry for this char
        lastNode = newTrieNode(char)
        currentNode.nextNodes[char] = lastNode
        currentNode = currentNode.nextNodes[char]
    }
    lastNode.endOfWord = true
}

func (trie *Trie) Has(word string) bool{
    currentNode := trie.rootNode
    var lastNode *TrieNode = currentNode
    for _, char := range word{
        node := currentNode.nextNodes[char]
        if node == nil{
            return false
        }
        currentNode = node
        lastNode = node
    }
    // full word
    if lastNode.endOfWord{
        return true
    }
    return false
}

func (trie *Trie) LongestWordLength() int{
    return trie.longestWordLen
}

func (trie *Trie) GetWords() *[]*TrieEntry{
    foundWords := &[]*TrieEntry{}
    for _, nextNode := range trie.rootNode.nextNodes{
        trie.walkNode(nextNode, foundWords, []rune{})
    }
    return foundWords
}

func (trie *Trie) walkNode(node *TrieNode, foundWords *[]*TrieEntry, currentString []rune){
    currentString = append(currentString, node.value)

    if node.endOfWord{
        *foundWords = append(*foundWords, &TrieEntry{
            Word: string(currentString),
            Counter: node.counter,
        })
    }

    for _, nextNode := range node.nextNodes{
        trie.walkNode(nextNode, foundWords, currentString)
    }
}

// TODO: remove

// Helpers
func (trie *Trie) String() string{
    return fmt.Sprintf("{Trie 'rootNode': %s}", trie.rootNode)
}
func (node *TrieNode) String() string{
    jsonRaw, _ := json.Marshal(node)
    return fmt.Sprintf("{TrieNode 'value': '%s', 'nextNodes': %s}", string(node.value), string(jsonRaw))
}
func (node *TrieNode) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "value": string(node.value),
        "counter": node.counter,
        "nextNodes": node.nextNodes,
        "endOfWord": node.endOfWord,
    })
}