let ep;
function init(endpoint) {
    ep = endpoint;
}
window.addEventListener("load", function(evt) {

    const output = document.getElementById("output");
    const input = document.getElementById("input");
    let ws;

    const print = function (message) {
        const d = document.createElement("li");
        const s = message.split('::');
        if (s.length > 1) {
            switch (s[0]) {
                case 'd': //description
                    d.classList.add('description');
                    break;
                case 's': //say
                    d.classList.add('say');
                    break;
                case 'sys': //system message
                    d.classList.add('system');
                    break;
                case 'w':
                    d.classList.add('whisper');
                    break;
            }
            d.textContent = s[1];
        } else {
            console.log(`ERROR: malformed message :${message}`)
        }
        if (output.childElementCount >= 50) {
            output.removeChild(output.firstChild)
        }
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket(ep);
        ws.onopen = function(evt) {
            //print("OPEN");
        };
        ws.onclose = function(evt) {
            //print("CLOSE");
            ws = null;
        };
        ws.onmessage = function(evt) {
            print(evt.data);
            output.scrollTop = output.scrollHeight - output.clientHeight;
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

        ws.send(logic(input.value));
        input.value = '';
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

    function logic(inputString) {
        const s = inputString.toLowerCase().split(' ');
        if (s.length > 0) {
            if (s.length > 1) {
                switch (s[0]) {
                    case 'l':
                    case 'look':
                        return `l::${s[1]}`;
                    case 'm':
                    case 'move':
                        return `${s[1]}::`;
                    case 'p':
                    case 'pick up':
                        return `p::${s[1]}`;
                    case 'a':
                    case 'attack':
                        return `a::${s[1]}`;
                    case 'i':
                    case 'inventory':
                        return 'i::';
                    case 'e':
                    case 'equip':
                        return `eq::${s[1]}`;
                    case 'u':
                    case 'unequip':
                        return `uq::${s[1]}`;
                    case 'w':
                    case 'whisper':
                        return `wh::${s[1]}::${inputString}`;
                    default:
                        return inputString;
                }

            } else {
                if (s[0] === 'l' || s[0] === 'look') {
                    return 'l::';
                } else if (s[0] === 'm' || s[0] === 'move') {
                    return 'm::';
                }
            }
        }
        return inputString;
    }
});