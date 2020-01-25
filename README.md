# gocodegraph

gocodegraph is an engine for a visual coding environment in go.
The current implementation first initializes all nodes which do require inputs in individual go routines. It uses channels to represent edges of the graph ```Channel chan interface{}```, so each of the routines is awaiting inputs on channels. It then initializes all nodes which do not require inputs, which are ready to be solved. Once solved, they are propagating their results down the line through channels. If all channels of a node received data it can be solved and it will send its result down the line, and so on.

![graph](https://github.com/moethu/gocodegraph/raw/master/images/graph.png)

Once results are ready, they can be streamed to the WebUI using a websocket.

![socket](https://github.com/moethu/gocodegraph/raw/master/images/arc.png)

## Web UI

![webui](https://github.com/moethu/gocodegraph/raw/master/images/screenshot.png)
