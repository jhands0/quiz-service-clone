import type { Player } from '../model/quiz';

export enum PacketTypes {
    Connect,
    HostGame,
    QuestionShow,
    ChangeGameState,
    PlayerJoin,
    StartGame
}

export enum GameState {
    Lobby,
    Play,
    Reveal,
    End
}

export interface ChangeGameStatePacket {
    state: GameState;
}

export interface PlayerJoinPacket {
    player: Player;
}

export class NetService {
    private webSocket!: WebSocket;
    private textDecoder: TextDecoder = new TextDecoder();
    private textEncoder: TextEncoder = new TextEncoder();

    private onPacketCallback?: (packet: any) => void;

    connect() {
        this.webSocket = new WebSocket('ws://localhost:3000/ws')
        this.webSocket.onopen = () => {
            console.log('opened connection');
        };

        this.webSocket.onmessage = async (event: MessageEvent) => {
            const arrayBuffer = await event.data.arrayBuffer();
            const bytes = new Uint8Array(arrayBuffer);
            const packetId = bytes[0];

            const packet = JSON.parse(this.textDecoder.decode(bytes.subarray(1)));

            packet.id = packetId;

            console.log(packetId);
            console.log(packet);

            if (this.onPacketCallback)
                this.onPacketCallback(packet);
        }
    }

    onPacket(callback: (packet: any) => void) {
        this.onPacketCallback = callback
    }

    sendPacket(packet: any) {
        const packetId = packet.id;
        const packetData = JSON.stringify(packet, (key, value) => 
            key == "id" ? undefined : value
        );

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