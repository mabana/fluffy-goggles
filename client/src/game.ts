import { Object } from './object';

export class Game {
    ws: WebSocket
    gameMap: number[][]
    objects: Object[][]

    constructor(ws: WebSocket, gameMap: number[][], objects: Object[][]) {
        this.ws = ws
        this.gameMap = gameMap
        this.objects = objects
    }

    initWsLoop() {
    }
}