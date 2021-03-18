# logspam
You know how sometimes when you go to read the logs of a random service, it spews a torrent of log messages at you - at such a high speed you can't make anything out (and it also clutters up your shell)? Well, logspam is a cli to avoid cluttering up you shell and also get a nice little measurement of the spam speed.

# install

````bash
$ go get github.com/karlpokus/logspam
````

# usage
```bash
$ ugly-process | logspam [-v]
> last sample: 500 lines/s
> last sample: 800 lines/s
```

# development
```bash
$ go test
$ ./testdata/spam.sh 5 | go run .
# run test and build the binary
$ ./build.sh
```

On a side note: this project became an interesting challenge to get an accurate speed reading. In order to not drop any messages we need to run at least 4 threads where (1) is listening on stdin, (2) is continuously notifying when sample time is up, (3) is counting log messages and (4) dumps calculated log speed to stdout.

# todos
- [x] use io.Reader
- [x] tests
- [x] build script
- [x] verbose output and flag
- [ ] sample rate cli opt
- [ ] use bufio.Reader if log entry size becomes an issue
- [ ] start timer as close to sample start as possible
- [ ] limit allocations if possible
- [ ] keyboard opt to release logspam valve on demand
- [x] make go getable

# license
MIT
