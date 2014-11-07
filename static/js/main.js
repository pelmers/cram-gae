(function() {
    var textArea = document.getElementById("codeInput"),
        codeMirror = CodeMirror.fromTextArea(textArea, {
            mode: "javascript",
        }),
        submitButton = document.getElementById("submitCode");

    submitButton.addEventListener('click', function() {
        var code = codeMirror.getValue(),
            ratio = document.getElementById("ratioInput").value,
            shape = document.getElementById("shapeSelect").value,
            xhr = new XMLHttpRequest();
        if (shape == "") {
            // just default to rectangle
            shape = "rect";
        }
        xhr.open("POST","/post", true);
        xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4) {
                // replace body of codemirror with the response
                codeMirror.setValue(xhr.response);
            }
        };
        xhr.send(JSON.stringify({
            Shape: shape,
            Ratio: parseFloat(ratio),
            Contents: code,
        }));
    });
})();
