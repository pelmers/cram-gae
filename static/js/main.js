(function() {
    "use strict;"
    var textArea = document.getElementById("codeInput"),
        codeMirror = CodeMirror.fromTextArea(textArea, {
            mode: "javascript"
        }),
        selectedShape = null,
        shapeDivs = document.getElementById("shapeSelect").children;

    // Post the code string to /post/language with specified shape and ratio.
    // Once we have the response, call callback(response)
    function postCode(language, code, ratio, shape, callback) {
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/post/" +language, true);
        xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4) {
                callback(xhr.response);
            }
        };
        xhr.send(JSON.stringify({
            Shape: shape,
            Ratio: ratio,
            Contents: code,
        }));
    }

    function getSelectedShape() {
        return selectedShape;
    }

    function setSelectedShape(shapeName) {
        document.getElementById(shapeName).classList.add('selected');
        if (selectedShape !== null) {
            document.getElementById(selectedShape).classList.remove('selected');
        }
        selectedShape = shapeName;
    }
    setSelectedShape("rect");

    function getRatio() {
        return parseFloat(document.getElementById("ratioText").innerHTML);
    }

    // bind event listeners
    for (var i = 0; i < shapeDivs.length; i++) {
        (function(id) {
            document.getElementById(id).addEventListener('click', function() {
                setSelectedShape(id);
            });
        })(shapeDivs[i].getAttribute("id"));
    }

    document.getElementById("ratioInput").addEventListener('wheel', function(e) {
        function sign(val) {
            return (val >= 0)?1:-1;
        }
        document.getElementById("ratioText").innerHTML = "" + Math.max(
            0.05, getRatio() + -0.05*sign(e.deltaY)).toFixed(2);
    });

    document.getElementById("submitCode").addEventListener('click', function() {
        var code = codeMirror.getValue(),
            ratio = getRatio(),
            shape = getSelectedShape();
        if (shape == "") {
            // just default to rectangle
            shape = "rect";
        }
        // eventually this will pick a specific endpoint based on langauge
        // but right now only js is supported
        postCode("js", code, ratio, shape, function(response) {
            codeMirror.setValue(response);
        });
    });
})();
