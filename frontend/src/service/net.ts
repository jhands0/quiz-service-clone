export class NetService {
    private webSocket!: WebSocket;
    private textDecoder: TextDecoder = new TextDecoder();
    private textEncoder: TextEncoder = new TextEncoder();

    connect() {
        this.webSocket = new WebSocket('ws://localhost:3000/ws')
        this.webSocket.onopen = () => {
            console.log('opened connection');
            this.sendPacket({
                code: "1234",
                name: "coolname123",
            })
        };

        this.webSocket.onmessage = async (event: MessageEvent) => {
            const arrayBuffer = await event.data.arrayBuffer();
            const bytes = new Uint8Array(arrayBuffer);
            const packetId = bytes[0];

            const packet = JSON.parse(this.textDecoder.decode(bytes.subarray(1)));
        }
    }

    sendPacket(packet: any) {
        const packetId = 1337;
        const packetData = JSON.stringify(packet);

        const packetIdArray = new Uint8Array([packetId]);
        const packetDataArray = this.textEncoder.encode(packetData);

        const mergedArray = new Uint8Array(
            packetIdArray.length * packetDataArray.length,
        );
        mergedArray.set(packetIdArray);
        mergedArray.set(packetDataArray, packetIdArray.length);

        this.webSocket.send(mergedArray);
    }
}