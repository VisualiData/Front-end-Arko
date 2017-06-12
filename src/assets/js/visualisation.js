function updateGraph(house, floor) {
  var checkedSensors = [];
  var checkedTypes = [];
  $("input:checkbox[name=sensor]:checked").each(function() {
    checkedSensors.push($(this).val());
  });
  $("input:checkbox[name=sensor_type]:checked").each(function() {
    checkedTypes.push($(this).val());
  });
  var to = $("#to").val();
  to = to.substr(6, 4) + "-" + to.substr(3, 2) + "-" + to.substr(0, 2) + "T00:00:00Z";
  var from = $("#from").val();
  from = from.substr(6, 4) + "-" + from.substr(3, 2) + "-" + from.substr(0, 2) + "T00:00:00Z";
  var url = "/visualisation/" + house + "/" + floor + "/" + checkedSensors.join() + "/" + checkedTypes.join() + "/" + from + "/" + to;
  window.location.href = url;
}

function displayData(input, sensors) {
  var sensors = sensors.split(",");
  var svg = d3.select("svg"),
    margin = {
      top: 20,
      right: 80,
      bottom: 30,
      left: 50
    },
    width = svg.attr("width") - margin.left - margin.right,
    height = svg.attr("height") - margin.top - margin.bottom,
    g = svg.append("g").attr("transform", "translate(" + margin.left + "," + margin.top + ")");
  var parseTime = d3.timeParse("%Y-%m-%dT%H:%M:%S");

  var x = d3.scaleTime().rangeRound([0, width]);
  var y = d3.scaleLinear().rangeRound([height, 0]);
  // loop data from every sensor
  for (var i = 0, len = input.length; i < len; i++) {
    var color = "#" + ((1 << 24) * Math.random() | 0).toString(16);
    var data = input[i];
    var line = d3.line()
      .x(function(d) {
        return x(parseTime(d.timestamp));
      })
      .y(function(d) {
        return y(d.value);
      });

    x.domain(d3.extent(data, function(d) {
      return parseTime(d.timestamp);
    }));
    y.domain(d3.extent(data, function(d) {
      return d.value;
    }));
    // add data to graph
    g.append("path")
      .datum(data)
      .attr("fill", "none")
      .attr("stroke", color)
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("stroke-width", 1.5)
      .attr("d", line);
  }
  // add axis labels
  g.append("g")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(x))
    .select(".domain")
    .remove();
  g.append("g")
    .call(d3.axisLeft(y))
    .append("text")
    .attr("fill", "#000")
    .attr("transform", "rotate(-90)")
    .attr("y", 6)
    .attr("dy", "0.71em")
    .attr("text-anchor", "end")
    .text("Temperature (Celcius)");

}
