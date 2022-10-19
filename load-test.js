import http from "k6/http";

// import { sleep } from "k6";
import { check } from "k6";

export const options = {
  scenarios: {
    constant_request_rate: {
      executor: "constant-arrival-rate",
      rate: 100,
      timeUnit: "1s", // 1000 iterations per second, i.e. 1000 RPS
      duration: "10s",
      preAllocatedVUs: 100, // how large the initial pool of VUs would be
      maxVUs: 1000, // if the preAllocatedVUs are not enough, we can initialize more
    },
  },
};

export default function () {
  const res = http.get(
    "http://localhost:8082/v1/query/NA/accounts/1016034603/vehicles"
  );
  check(res, {
    "status was 200": (r) => r.status == 200,
  });
  check(res, {
    "payload is valid": (r) => {
      const json = JSON.parse(r.body);
      return Array.isArray(json.data) && json.data.length > 0;
    },
  });
}
