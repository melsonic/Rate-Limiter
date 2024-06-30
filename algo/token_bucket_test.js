import http from "k6/http";
import { check, sleep } from "k6";
// import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
  vus: 10,
  iterations: 60,
};

export default () => {
  sleep(Math.random() * 2)
  const res = http.get("http://localhost:3045/limited");
  check(res, {
    "is status 200": (r) => r.status == 200,
  });
};
