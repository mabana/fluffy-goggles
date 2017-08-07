let mainCanvas: HTMLCanvasElement = <HTMLCanvasElement> document.getElementById('main-canvas')
let ctx: CanvasRenderingContext2D = <CanvasRenderingContext2D> mainCanvas.getContext('2d')

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

ctx.save()

for(let y: number = 0; y<12; y++) {
    for(let x: number = 0; x<16; x++) {
        let cell: number = map[y][x]
        let xx: number = x * 50
        let yy: number = y * 50

        ctx.fillStyle = (cell == 1) ? "#333" : "#fff"
        ctx.fillRect(xx, yy, xx+50, yy+50)
        ctx.stroke()
    }
}

ctx.restore()

document.addEventListener("keydown", function(this: Document, ev: KeyboardEvent){
    switch(ev.keyCode) {
        case 37: {
            console.log("LeftArrow")
            break
        }

        case 38: {
            console.log("UpArrow")
            break
        }

        case 39: {
            console.log("RightArrow")
            break
        }

        case 40: {
            console.log("BottomArrow")
        }
    }
})