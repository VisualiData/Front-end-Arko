function drawCircle(context, x, y, radius, color) {
  context.beginPath();
  context.arc(x, y, radius, 0, 2 * Math.PI);
  context.fillStyle = color;
  context.fill();
  context.stroke();
}

var Circle = function(x, y, radius, sensor_info) {
  this.left = x - radius;
  this.top = y - radius;
  this.right = parseInt(x) + radius;
  this.bottom = parseInt(y) + radius;
  this.sensor_info = sensor_info;
};

function loadFloorPlan() {
  var currentFloot = $('select[name="floor"]').val();
  var naturalWidth = 0;
  var naturalHeight = 0;
  var imgWidth = 0;
  var imgHeight = 0;
  // load floorplan image
  var image = document.createElement('img');
  image.src = '/assets/img/CHIBB_' + currentFloot + '.png';
  var canvas = document.getElementById("myCanvas");
  var context = canvas.getContext('2d');
  image.onload = function() {
    imgWidth = image.width;
    imgHeight = image.height;
    naturalWidth = image.naturalWidth;
    naturalHeight = image.naturalHeight;
    canvas.width = imgWidth;
    canvas.height = imgHeight;
    context.drawImage(image, 0, 0, imgWidth, imgHeight);
    if ($('#field_x').val() != ""){
      drawCircle(context, $('#field_x').val(), $('#field_y').val(), 8, "black");
    }
  }
  // change floorplan image
  $('select[name="floor"]').on('change', function() {
    image.src = '/assets/img/CHIBB_' + this.value + ".png";
    context.drawImage(image, 0, 0, imgWidth, imgHeight);
  });
  // draw circle on clicked position
  $("#myCanvas").on("click", function(event) {
    context.drawImage(image, 0, 0, imgWidth, imgHeight);
    var currentClickPosX = event.pageX - this.offsetLeft;
    var currentClickPosY = event.pageY - this.offsetTop;
    var naturalClickPosX = (naturalWidth / canvas.scrollWidth) * currentClickPosX;
    var naturalClickPosY = (naturalHeight / canvas.scrollHeight) * currentClickPosY;
    $("#field_x").attr("value", parseInt(naturalClickPosX));
    $("#field_y").attr("value", parseInt(naturalClickPosY));
    drawCircle(context, naturalClickPosX, naturalClickPosY, 8, "black");
  });
}
