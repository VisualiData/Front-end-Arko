$(document).ready(function() {
  var img_buffer = document.createElement('img');
  img_buffer.src = '/assets/img/CHIBB_0.png';
  var naturalWidth = 0;
  var imgWidth = 0;
  var imgHeight = 0;

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
  $('select').on('change', function() {
    img_buffer.src = '/assets/img/CHIBB_' + this.value + ".png";
    context.drawImage(img_buffer, 0, 0, imgWidth, imgHeight);
  })

  $("#myCanvas").on("click", function(event) {
    context.drawImage(img_buffer, 0, 0, imgWidth, imgHeight);
    var currentClickPosX = event.pageX - this.offsetLeft;
    var currentClickPosY = event.pageY - this.offsetTop;
    var naturalClickPosX = (naturalWidth / canvas.scrollWidth) * currentClickPosX;
    var naturalClickPosY = (naturalHeight / canvas.scrollHeight) * currentClickPosY;
    $("#field_x").val(parseInt(naturalClickPosX));
    $("#field_y").val(parseInt(naturalClickPosY));
    context.beginPath();
    context.arc(naturalClickPosX, naturalClickPosY, 5, 0, 2 * Math.PI);
    context.stroke();
  });
});
