# Byte-pair Encoding
Very loose (bad) implementation of bpe. Instead of doing "byte"-pair, this does 
"char"-pair tokenization, forgive me.

This serves as an example implementation to bpe. The best way to learn how bpe 
works is by looking at the code. I left some comments as well.

## Command line
```sh
$ go run main.go -h
Usage of /tmp/go-build2260239964/b001/exe/main:
  -load string
        load trainer from json file (eg. './trainer.json')
  -save string
        save trainer to json file (default "./trainer.json")
  -text string
        the text to be used by the trainer
  -train
        train it
  -v    verbose output
```

## Sample run
```sh
$ go run main.go -text "Hello, World. Hello" -train
[+] Start training
[+] Saved trainer to './trainer.json'
[+] Learned vocabs are shown here:
["W","r",".","H","He","Hel","Hell","Hello"," ","o",",","d","e","l"]

$ go run main.go -text "test the program. Quick! Do the
 tests now!" -train -load trainer.json
[+] Loading trainer from 'trainer.json'
[+] Start training
[+] Saved trainer to './trainer.json'
[+] Learned vocabs are shown here:
[",","s","h","p","o","r","H","He","Hel","Hell","Hello","W","u","k","w","l"," ","g","Q","!","n","i","c","d","e",".","t","te","tes","test","th","the","a","m","D"]
```

## Test
```sh
go test -v com.vincentcodes/bpe/tests
```

## Caution
I did not read any documentation for bpe while implementing the algorithm.
I read some articles about it. Got the gist of bpe and implement it in my own way.

If you are looking for real bpe, this repo shouldn't suit your needs.