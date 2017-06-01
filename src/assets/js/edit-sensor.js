function drawCircle(context, x, y, radius, color) {
  context.beginPath();
  context.arc(x, y, radius, 0, 2 * Math.PI);
  context.fillStyle = color;
  context.fill();
  context.stroke();
}

var Circle = function(x, y, radius) {
    this.left = x - radius;
    this.top = y - radius;
    this.right = parseInt(x) + radius;
    this.bottom = parseInt(y) + radius;
};

function loadFloorPlan(){
  var naturalWidth = 0;
  var naturalHeight = 0;
  var imgWidth = 0;
  var imgHeight = 0;
  var img_buffer = document.createElement('img');
  img_buffer.src = '/assets/img/CHIBB_0.png';
  var canvas = document.getElementById("myCanvas");
  var context = canvas.getContext('2d');
  img_buffer.onload = function() {
    imgWidth = img_buffer.width;
    imgHeight = img_buffer.height;
    naturalWidth = img_buffer.naturalWidth;
    naturalHeight = img_buffer.naturalHeight;
    canvas.width = imgWidth;
    canvas.height = imgHeight;
    context.drawImage(img_buffer, 0, 0, imgWidth, imgHeight);
  }
  $('select[name="floor"]').on('change', function() {
    img_buffer.src = '/assets/img/CHIBB_' + this.value + ".png";
    context.drawImage(img_buffer, 0, 0, imgWidth, imgHeight);
  });
  $("#myCanvas").on("click", function(event) {
    context.drawImage(img_buffer, 0, 0, imgWidth, imgHeight);
    var currentClickPosX = event.pageX - this.offsetLeft;
    var currentClickPosY = event.pageY - this.offsetTop;
    var naturalClickPosX = (naturalWidth / canvas.scrollWidth) * currentClickPosX;
    var naturalClickPosY = (naturalHeight / canvas.scrollHeight) * currentClickPosY;
    $("#field_x").attr("value", parseInt(naturalClickPosX));
    $("#field_y").attr("value", parseInt(naturalClickPosY));
    drawCircle(context, naturalClickPosX, naturalClickPosY, 8, "black");
  });
}
