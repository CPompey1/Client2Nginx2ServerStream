const net = require('net');

// create a TCP socket
const client = new net.Socket();

// connect to the server
client.connect(8080, 'localhost', () => {
    console.log('Connected to server!');
});

// send hello message every second
setInterval(() => {
    client.write('hello\n');
}, 1000);

// handle incoming data (if needed)
client.on('data', (data) => {
    console.log('Received: ' + data);
});

// handle when connection is closed
client.on('close', () => {
    console.log('Connection closed');
});