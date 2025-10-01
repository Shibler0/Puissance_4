const container = document.getElementById("container");
const buttons = document.getElementById("buttons")

const rows = 6;
const cols = 7;

let grid = []; // tableau 2D [ligne][colonne]

for (let row = 0; row < rows; row++) {
    grid[row] = [];
    for (let col = 0; col < cols; col++) {
        const canvas = document.createElement("canvas");
        canvas.width = 60;
        canvas.height = 60;
        container.appendChild(canvas);

        const ctx = canvas.getContext("2d");
        ctx.beginPath();
        ctx.arc(30, 30, 25, 0, Math.PI * 2);
        ctx.fillStyle = "white";
        ctx.fill();

        grid[row][col] = canvas;
    }
}

for(let i = 0; i < 7; i++) {
    const button = document.createElement("button");
    button.textContent = i +1;
    if (i == 0) {
        button.addEventListener('click', () => colorCell(5, 0, "red"))
    }
    buttons.appendChild(button);
}

function colorCell(row, col, color) {
    const canvas = grid[row][col];
    const ctx = canvas.getContext("2d");
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.beginPath();
    ctx.arc(30, 30, 25, 0, Math.PI * 2);
    ctx.fillStyle = color;
    ctx.fill();
}

