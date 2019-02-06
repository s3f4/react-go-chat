import { EventEmitter } from 'events';
class Socket {
    constructor(ws = new WebSocket(), ee = new EventEmitter) {
        this.ee = ee;
        this.ws = ws;
        ws.onmessage = this.message.bind(this)
        ws.onopen = this.open.bind(this)
        ws.onclose = this.close.bind(this)
    }

    on(name, fn) {
        this.ee.on(name, fn);
    }

    of(name, fn) {
        this.ee.removeListener(name, fn);
    }

    emit(data, fn) {
        const message = JSON.stringify({ name, data });
        this.ws.send(message);
    }
    open() {
        this.ee.emit('connect');
    }

    message(e) {
        try {
            const message = JSON.parse(e.data);
            this.ee.emit(message.name, message.data);
        } catch (error) {
            this.ee.emit('error', error)
        }
    }

    close() {
        this.ee.emit('disconnect');
    }
}

export default Socket