# gocodegraph

gocodegraph is an engine for a visual coding environment in go.
The current implementation first initializes all nodes which do require inputs. It uses channels to represent edges of the graph ```Channel chan interface{}```, so each of the nodes is awaiting inputs with channels.
It then initializes all nodes which do not require inputs, which are ready to be solved. Once solved, they are propagating their results down the line through channels. If all channels of a node received data it can be solved and it will send its result down the line, and so on.

![alt text](https://github.com/moethu/gocodegraph/raw/master/images/screenshot.png)
