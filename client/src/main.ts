let websocket = new WebSocket("ws://127.0.0.1:3000/wss")

websocket.onopen = function(ev: Event) {
    websocket.send("First message")
}

websocket.onmessage = function(ev: MessageEvent) {
    map = JSON.parse(ev.data)
}

let mainCanvas: HTMLCanvasElement = <HTMLCanvasElement> document.getElementById('main-canvas')
let ctx: CanvasRenderingContext2D = <CanvasRenderingContext2D> mainCanvas.getContext('2d')

function drawCell(ctx: CanvasRenderingContext2D, x:number, y:number, color: string) {
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

function render(canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D, map: number[][]) {
    ctx.save()

    for(let y: number = 0; y<12; y++) {
        for(let x: number = 0; x<16; x++) {
            let cell: number = map[y][x]
            drawCell(ctx, x, y, mapFieldColor(cell))
        }
    }

    ctx.restore()
}

let map: number[][] = [
    [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
    [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
]

render(mainCanvas, ctx, map)

document.addEventListener("keydown", function(this: Document, ev: KeyboardEvent){
    ev.preventDefault()

    switch(ev.keyCode) {
        case 37: {
            websocket.send("move:left")
            break
        }

        case 38: {
            websocket.send("move:up")
            break
        }

        case 39: {
            websocket.send("move:right")
            break
        }

        case 40: {
            websocket.send("move:down")
            break
        }
    }  
})

function step() {
    render(mainCanvas, ctx, map)
    window.requestAnimationFrame(step)
}

window.requestAnimationFrame(step)