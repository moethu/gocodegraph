components = {}

var ws;

function uuidv4() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0,
            v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

function createComponent(name) {
    component = {
        top: 10,
        left: 10,
        properties: {
            title: name,
            inputs: {},
            outputs: {}
        }
    }

    item = components[name]

    if (item.I) {
        ctr = 1
        item.I.forEach(i => {
            component.properties.inputs[`input_${ctr}`] = {
                label: i.Name
            }
            ctr++
        })
    }

    if (item.O) {
        ctr = 1
        item.O.forEach(i => {
            component.properties.outputs[`output_${ctr}`] = {
                label: i.Name
            }
            ctr++
        })
    }

    return component
}


function placeComponent(name) {
    id = uuidv4()
    var widget = $("#example").data("flowchart-flowchart");
    widget.data.operators[id] = createComponent(name)
    widget.createOperator(id, widget.data.operators[id]);
}

function load() {
    console.log("x")
    $.get("./components", function (data, status) {
        console.log("Data: " + data + "\nStatus: " + status);
        d = JSON.parse(data)
        d.forEach(item => {
            name = item.Name.replace("components.", "")
            components[name] = item
            $('#components').append(
                `<li class="nav-item"><a class="nav-link" href="#" onclick="placeComponent('${name}')"><span data-feather="file"></span>${name}</a></li>`)
        })
    });

    $('#example').flowchart({
        data: {}
    });

    $("#solve").click(function () {
        if (ws) {
            return false;
        }

        ws = new WebSocket(`ws://localhost:8000/solver`);
        var widget = $("#example").data("flowchart-flowchart");

        ws.onopen = function(evt) {
            console.log("Connected to Server");
            
            result = {}
            for (var field in widget.data) {
                if (widget.data.hasOwnProperty(field)) {
                    var name = widget.data[field];
                    result[field] = name;
                }
            }
            let djson = JSON.stringify(result);
            console.log(djson)
            ws.send(djson)
        }

        ws.onclose = function(evt) {
			console.log("Closed Connection");
            ws = null;
        }
        
        ws.onmessage = function(evt) {
            console.log(evt.data)
            data = JSON.parse(evt.data);
            widget.data.operators[data.Id].internal.properties.outputs[data.Port].label = data.Value.toString()
            widget.setOperatorData(data.Id,widget.data.operators[data.Id]);
        }

        ws.onerror = function(evt) {
            console.log("Error: " + evt.data);
        }
        return false;

    });
}