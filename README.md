# `watch` with a graph

[`watch`](https://www.freebsd.org/cgi/man.cgi?query=watch&manpath=SuSE+Linux/i386+11.3) +
[`asciigraph`](https://github.com/guptarohit/asciigraph)

![demo.gif](demo.gif)

## Usage

`watchgraph --interval=500ms <some command that simply returns an integer>`

I needed this for testing a search API that returned some JSON which contained the time taken to conduct the search. The
command you run is passed to Bash so you can use pipes and stuff. In the demo, I used `jq` to extract the time taken
value as an integer which is passed to the graph and rendered.
