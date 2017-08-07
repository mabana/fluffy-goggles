export class Player {
    x: number
    y: number

    constructor(x: number, y: number) {
        this.x = x
        this.y = y
    }

    moveLeft() {
        this.x--
    }

    moveRight() {
        this.x++
    }

    moveUp() {
        this.y--
    }

    moveBottom() {
        this.y++
    }
}