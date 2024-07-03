import http from "k6/http";
import { check, sleep } from "k6";

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
