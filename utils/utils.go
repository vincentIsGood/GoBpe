package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func IsSpace[T fmt.Stringer](obj T) bool{
    return strings.TrimSpace(obj.String()) == ""
}

func Min[T int | float32](a, b T) T{
    if a < b{
        return a
    }
    return b
}

func InitArray[T any](length int, defaultVal T) *[]T{
    array := make([]T, length)
    for i := range array{
        array[i] = defaultVal
    }
    return &array
}

func PrintObject(obj any){
    jsonRaw, _ := json.Marshal(obj)
    fmt.Println(string(jsonRaw))
}