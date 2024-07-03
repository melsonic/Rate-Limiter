import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  vus: 10,
  iterations: 150,
};

export default () => {
  sleep(Math.random() * 1)
  const res = http.get("http://localhost:3045/limited");
  check(res, {
    "is status 200": (r) => r.status == 200,
  });
};
