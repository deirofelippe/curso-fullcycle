const express = require("express");
const promMid = require("express-prometheus-middleware");

const app = express();

const port = 3000;

app.use(
  promMid({
    metricsPath: "/metrics",
    collectDefaultMetrics: true,
    collectGCMetrics: true,
    requestDurationBuckets: [0.1, 0.5, 1, 1.5],
    requestLengthBuckets: [512, 1024, 5120, 10240, 51200, 102400],
    responseLengthBuckets: [512, 1024, 5120, 10240, 51200, 102400],
  })
);

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
