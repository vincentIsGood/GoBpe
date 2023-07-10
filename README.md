# Byte-pair Encoding
Very loose (bad) implementation of bpe. Instead of doing "byte"-pair, this does 
"char"-pair tokenization.

This serves as an example implementation to bpe. 

## Test
```sh
go test -v com.vincentcodes/bpe/tests
```