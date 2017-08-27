import { Game } from './game'
import { Player } from './player'
import { Map, ViewportHeight, ViewportWidth } from './map'

interface GreetingData {
    gameMap: number[][]
    x: number
    y: number
    mapWidth: number
    mapHeight: number
}

let mainCanvas: HTMLCanvasElement = <HTMLCanvasElement> document.getElementById('main-canvas')
let ctx: CanvasRenderingContext2D = <CanvasRenderingContext2D> mainCanvas.getContext('2d')
let websocket = new WebSocket(`ws://${window.location.host}/wss`)

websocket.onopen = function(ev: Event) {
    websocket.send("First message")
}

let map: Map
let player: Player

function drawCell(ctx: CanvasRenderingContext2D, x: number, y: number, color: string) {
    let xx: number = x * 50
    let yy: number = y * 50
    ctx.fillStyle = color
    ctx.fillRect(xx, yy, 50, 50)
    ctx.stroke()
}

function mapFieldColor(fieldContent: number): string {
    switch (fieldContent){
        case 1: {
            return "#333"
        }

        case 2: {
            return "#f3f"
        }

        default: {
            return "#fff"
        }
    }
}

function render(canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D, map: Map) {
    ctx.save()

    map.loopOverPartOfMap(player.x, player.y, (x: number, y: number, content: number) => {
        drawCell(ctx, x, y, mapFieldColor(content))
    })
    
    ctx.restore()
}

function step() {
    render(mainCanvas, ctx, map)
    window.requestAnimationFrame(step)
}

websocket.onmessage = function(ev: MessageEvent) {
    let greetingData: GreetingData = JSON.parse(ev.data)
    map = new Map(greetingData.mapWidth, greetingData.mapHeight, greetingData.gameMap)
    player = new Player(greetingData.x, greetingData.y)
    window.requestAnimationFrame(step)
}

document.addEventListener("keydown", function(this: Document, ev: KeyboardEvent){
    ev.preventDefault()

    switch(ev.keyCode) {
        case 37: {
            player.x--
            //websocket.send("move:left")
            break
        }

        case 38: {
            player.y--
            //websocket.send("move:up")
            break
        }

        case 39: {
            player.x++
            //websocket.send("move:right")
            break
        }

        case 40: {
            player.y++
            //websocket.send("move:down")
            break
        }
    }  
})