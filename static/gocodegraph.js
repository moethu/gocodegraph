components = {}

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
        var widget = $("#example").data("flowchart-flowchart");
        result = {}
        for (var field in widget.data) {
            if (widget.data.hasOwnProperty(field)) {
                var name = widget.data[field];
                result[field] = name;
            }
        }
        let djson = JSON.stringify(result);
        $.post("./solve", djson, function (data, status) {
            console.log("Data: " + data + "\nStatus: " + status);
        });
    });
}