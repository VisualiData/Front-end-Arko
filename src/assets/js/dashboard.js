function dashboard(data) {
  var circles = [];
  var active = 0;
  var warning = 0;
  var inactive = 0;
  var plans = ["CHIBB_0", "CHIBB_1", "CHIBB_2"];
  var floorplans = {};
  fillSensorList(data);

  plans.forEach(function(value) {
    var newFloor = {
      "floor_plan": value
    };

    var $div = $("<div>", {
      id: value,
      "class": "columns large-12"
    });
    $("#floor_plan_container").append($div);
    $('<p></p>', {
      id: "p_" + value,
      'style': 'text-align: center;',
      'text': value
    }).appendTo("#" + value);
    var image = document.createElement('img');
    image.src = '/assets/img/' + value + '.png';
    $('<canvas></canvas>', {
      id: "canvas_" + value,
      class: "floor_plan",
      'style': 'width: 100%'
    }).appendTo("#" + value);
    var canvas = document.getElementById("canvas_" + value);
    newFloor["canvas"] = "canvas_" + value;
    var context = canvas.getContext('2d');
    image.onload = function() {
      var imgWidth = image.width;
      var imgHeight = image.height;
      var naturalWidth = image.naturalWidth;
      var naturalHeight = image.naturalHeight;
      newFloor['naturalWidth'] = naturalWidth;
      newFloor['naturalHeight'] = naturalHeight;
      canvas.width = imgWidth;
      canvas.height = imgHeight;
      context.drawImage(image, 0, 0, imgWidth, imgHeight);
      var vars = value.split("_");
      var wanted = data.filter(function(item) {
        return (item.position.floor == vars[1] && item.position.house == vars[0]);
      });
      wanted.forEach(function(sensor) {
        var color = "red";
        if (sensor.status == "active") {
          color = 'green';
          active++;
        } else if (sensor.status == "intermittent failures") {
          color = 'orange';
          warning++;
        } else {
          inactive++;
        }
        var circle = new Circle(sensor.position.x, sensor.position.y, 10, sensor);
        circles.push(circle);
        drawCircle(context, sensor.position.x, sensor.position.y, 10, color);
      });
      $("#active").text(active);
      $("#warning").text(warning);
      $("#inactive").text(inactive);
    }
    floorplans[value] = newFloor;
  });

  $('.floor_plan').click(function(e) {
    var clickedCanvas = e.target.id.replace("canvas_", "");
    var clickedCoordinates = getClickedCoordinates(e, this, floorplans[clickedCanvas]);
    checkClickedPosition(clickedCoordinates, circles);
  });
}

function getClickedCoordinates(event, _self, floor_plan) {
  var canvas = document.getElementById(floor_plan["canvas"]);
  var clickedX = event.pageX - _self.offsetLeft;
  var clickedY = event.pageY - _self.offsetTop;
  var naturalClickPosX = (floor_plan.naturalWidth / canvas.scrollWidth) * clickedX;
  var naturalClickPosY = (floor_plan.naturalHeight / canvas.scrollHeight) * clickedY;
  return {
    "x": naturalClickPosX,
    "y": naturalClickPosY
  };
}

function checkClickedPosition(coordinates, circles) {
  var x = coordinates.x;
  var y = coordinates.y;
  for (var i = 0; i < circles.length; i++) {
    if (x < circles[i].right && x > circles[i].left && y > circles[i].top && y < circles[i].bottom) {
      $('#exampleModal1').html(sensorInfo(circles[i].sensor_info)).foundation('open');
    }
  }
}

function sensorInfo(data) {
  var html = "<h1>" + data.sensor_id + "</h1>";
  html += "<p>House: " + data.position.house + "<br>";
  html += "Floor: " + data.position.floor + "</p>";
  html += "<p>Status: " + data.status + "<br>";
  html += "Type: " + data.type + "<br>";
  html += "Node: " + data.nodeName + "<br>";
  html += "Node type: " + data.nodeType + "</p>";
  html += '<button class="close-button" data-close aria-label="Close reveal" type="button"><span aria-hidden="true">&times;</span></button>';
  return html;
}

function fillSensorList(data) {
  for (var i = 0; i < data.length; i++) {
    if (!$('#floor' + data[i].position.floor).length) {
      $('#sensors').append('<li><a href="#">Floor ' + data[i].position.floor + '</a><ul class="menu vertical sublevel-1" id="floor' + data[i].position.floor + '"></ul></li>');
    }
    $('#floor' + data[i].position.floor).append('<li><a class="subitem" href="/sensor/edit/' + data[i].sensor_id + '">' + data[i].sensor_id + '</a></li>');
  }
}
