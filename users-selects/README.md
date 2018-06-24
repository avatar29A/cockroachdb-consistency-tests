# users-selects

# Usage

```
$ time ./users-selects -servername="10.8.162.205:26257" -times=1000
```

# Minibench SELECTS

Query:

```
SELECT * FROM users WHERE username = 'username_%d' LIMIT 1
```
DataSet: 1 million rows

VC1S:

2 X86 64bit Cores
2GM RAM

CocroachDB: 1Gb Cache

Results: (1000 random queries)

real	0m10.746s
user	0m0.756s
sys	0m0.984s

Latency: ~10ms per query





