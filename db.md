# Making a flatfile nosql database

* data is stored in a big ole blob
* the start of the blob has all of the tables, fields and uuids for quick indexing
* each uuid links to a data address in the file



goroutines make this type of work much less intensive

with goroutines:

```
./db  0.00s user 0.01s system 21% cpu 0.049 total
```

without:

```
./db  0.00s user 0.00s system 86% cpu 0.009 total
```

