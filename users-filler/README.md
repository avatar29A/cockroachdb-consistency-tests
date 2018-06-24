# users-filler

This project makes a lot of insert queries to cockroach cluster. All queries gathering to batches by 100 items.

Template of query:

<username>_N, where N is a monotonic incrementing number, and username is a predefined prefix.

## Usage

```
$ users-filler -username=testuser -from=0 -to=500000 -servername="localhost:26572"
```

Args:

* *from* is an initial value of N. (default 0)
* *to* is an upper bound of N. (default 0)

## Minibench INSERTS (with ON CONFLICT condition)

50000 rows
100 rows per batch

VC1S:

2 X86 64bit Cores
2GM RAM

Cluster: 5 nodes

Results:


```
$ time ./users-filler -servername="10.8.162.205:26257" -from=0 -to=50000

real	0m24.742s
user	0m0.184s
sys	0m0.180s
```

~2000 rows per second.