const canvas = document.getElementById("myCanvas");
const ctx = canvas.getContext("2d");

ctx.beginPath();                
ctx.arc(150, 100, 50, 0, Math.PI * 2); 
ctx.fillStyle = "red";           
ctx.fill();                      
ctx.stroke(); 