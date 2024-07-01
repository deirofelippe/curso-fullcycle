import http from "k6/http";
import { check } from "k6";
import { group, sleep } from "k6";

/*
smoke test
1vu - 1m

load test
100vu - 1m
100vu - 2m
0vu - 1m

stress test
100vu - 1m
100vu - 1m
200vu - 1m
200vu - 1m
300vu - 1m
300vu - 1m
400vu - 1m
400vu - 1m
0vu - 2m

spike test
100vu - 0.5m
100vu - 1m
1500vu - 5m
100vu - 1m
0vu - 0.5s
- excelente: sistema nao quebrou, no tempo esperado, depois voltou ao normal
- bom: sistema ficou mais lento, depois voltou ao normal, nao deu erro
- ruim: deu erro no pico, mas depois seguiu em frente
- muito ruim: no pico deu erro, na baixa continuou com erro

soak test
400vu - 1m
400vu - 1d
0vu - 1m
- verifica capacidade do banco de dados
- verifica se nao vai ter vazamento de memoria ou bugs
- todos os sistemas nao vao parar de funcionar

- vale a pena implementar
- use ambiente sandbox, use o ambiente de producao em horario de baixo acesso


const opt1 = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "shared-iterations",
      vus: 10,
      iterations: 200,
      maxDuration: "30s",
    },
  },
};

// numero de vu complete o msm numero de iterations
const opt2 = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "per-vu-iterations",
      vus: 10,
      iterations: 20,
      maxDuration: "30s",
    },
  },
};
const opt3 = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "constant-vus",
      vus: 10,
      maxDuration: "30s",
    },
  },
};

// vu seja rampado
const opt4 = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "ramping-vus",
      startVUs: 0,
      stages: [
        { duration: "20s", target: 10 },
        { duration: "10s", target: 0 },
      ],
      gracefulRampDown: "0s",
    },
  },
};

// bom para usar em teste de stress e pico
// vai fazer as iteracoes, independente do tempo de resposta e erro
const opt5 = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "constant-arrival-rate",
      duration: "30s",
      rate: 30,
      timeUnit: "1s",
      preAllocatedVUs: 2,
      maxVUs: 50,
    },
  },
};

// bom para usar em teste de stress e pico
// vai fazer as iteracoes, independente do tempo de resposta e erro, terá a rampa
const opt6 = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: "ramping-arrival-rate",
      startRate: 300,
      timeUnit: "1m",
      preAllocatedVUs: 2,
      maxVUs: 50,
      stages: [
        { duration: "1m", target: 300 },
        { duration: "2m", target: 600 },
        { duration: "4m", target: 600 },
        { duration: "2m", target: 60 },
      ],
    },
  },
};
*/

export let options = {
  stages: [
    { duration: "5s", target: 1000 },
    { duration: "5m", target: 1000 },
    { duration: "5s", target: 0 },
  ],
  thresholds: {
    http_req_failed: ["rate < 0.01"], // http errors should be less than 1%
    http_req_duration: ["p(98) < 500"], // 95% of requests should be below 200ms
  },
};

function GetHelloWorld() {
  let res = http.get("http://nodeapp:3000/");
  check(res, { "status is 200": (r) => r.status === 200 });
  sleep(1);
}

export default function () {
  group("Endpoint GET - Hello World", () => {
    GetHelloWorld();
  });
}
