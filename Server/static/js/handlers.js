let ep;
function init(endpoint) {
    ep = endpoint;
}
window.addEventListener("load", function(evt) {

    const output = document.getElementById("output");
    const input = document.getElementById("input");
    let ws;

    const print = function (message) {
        const d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket(ep);
        ws.onopen = function(evt) {
            print("OPEN");
        };
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        };
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        };
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        };
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});