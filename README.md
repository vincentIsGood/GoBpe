# Byte-pair Encoding
Very loose (bad) implementation of bpe. Instead of doing "byte"-pair, this does 
"char"-pair tokenization, forgive me.

This serves as an example implementation to bpe. The best way to learn how bpe 
works is by looking at the code. I left some comments as well.

## Test
```sh
go test -v com.vincentcodes/bpe/tests
```

## Caution
I did not read any documentation for bpe while implementing the algorithm.
I read some articles about it. Got the gist of bpe and implement it in my own way.

If you are looking for real bpe, this repo shouldn't suit your needs.