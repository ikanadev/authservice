import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
  // vus: 1000,
  // iterations: 100000,
  // duration: '5s',
  stages: [
    {duration: '10s', target: 100},
    {duration: '10s', target: 1400},
    {duration: '1m', target: 1400},
    {duration: '10s', target: 100},
    {duration: '10s', target: 0},
  ],
};

export default function () {
  const url = 'http://127.0.0.1:8000/authenticate';
  const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJBcHAiLCJleHAiOjE2MTUxOTgwMDcsImlhdCI6MTYxNDU5MzIwNywiSUQiOjEsIlJvbGUiOjJ9.eymqVG7bzuw4cg1CsEXD7rnrM_HcI2Z33ITYybkP9Gw';

  const params = {
    headers: {
      'Authorization': token,
    },
  };
  http.get(url, params);
  // sleep(1);
}