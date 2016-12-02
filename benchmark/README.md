# go-raml generated code benchmark

We benchmarked our generated code to test how fast it is.
benchmark condition:
- use [hey](https://github.com/rakyll/hey) and [boom](https://github.com/tarekziade/boom) as benchmark tools
- both benchmark tools and server code in one machine
- generated code needs some modification
- benchmark tools and server are in different VMs, each VM has 2 GB of RAM and 2GB of CPU.


Results summary:
- Nim is generally faster than Go if both use 1 CPU core
  Sometimes Go faster.
- Go faster if allowed to use all cores
- Python is the slowest
- `hey` give more request per seconds, it seems `hey` is better tool `boom`
  in term of concurency
- Nim doesn't support chunked transfer encoding which is used by `hey` and `ab`

result details:
- [hey result](./results.hey.md)
- [boom result](./results.boom.md)
