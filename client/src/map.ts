export const ViewportHeight = 13
export const ViewportWidth = 17

class Point {
    x: number
    y: number

    setX(x: number) {
        this.x = x
    }

    setY(y: number) {
        this.y = y
    }

    getX(): number {
        return this.x
    }

    getY(): number {
        return this.y
    }
}

export class Map {
    width: number
    height: number
    content: number[][]

    constructor(width: number, height: number, content: number[][]) {
        this.width = width
        this.height = height
        this.content = content
    }

    loopOverPartOfMap(playerX: number, playerY: number, lambda: (x: number, y: number, content: number)=> void) {
        let startPoint: Point = this.calculateStartPoint(playerX, playerY)

        for(let y: number = 0; y<ViewportHeight; y++) {
            for(let x: number = 0; x<ViewportWidth; x++) {
                lambda(x, y, this.content[startPoint.getY() + y][startPoint.getX()+x])
            }
        }

        if(startPoint.x > 0) {
            playerX -= startPoint.x
        }

        if(startPoint.y > 0) {
            playerY -= startPoint.y
        }

        lambda(playerX, playerY, 2)
    }

    calculateStartPoint(x: number, y: number): Point {
        let rx: number = Math.floor(ViewportWidth / 2)
        let ry: number = Math.floor(ViewportHeight / 2)

        let startPoint: Point = new Point

        if(x < rx) {
            startPoint.setX(0)
        } else if(x+rx >= this.width) {
            startPoint.setX(this.width - ViewportWidth)
        } else {
            startPoint.setX(x - rx)
        }

        if (y < ry) {
            startPoint.setY(0)
        } else if (y + ry >= this.height) {
            startPoint.setY(this.height - ViewportHeight)
        } else {
            startPoint.setY(y - ry)
        }

        return startPoint
    }
}