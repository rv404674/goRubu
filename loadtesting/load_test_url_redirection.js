import { check } from "k6";
import http from "k6/http";
import { Rate } from 'k6/metrics';


export let errorRate = new Rate('errors');

export default function() {
    var url = "http://localhost:8080/all/redirect";
    var params = {
        headers: {
            'Content-Type': 'application/json'
          }
    };

    var data = JSON.stringify({
        "Url": "https://goRubu/Nzg0ODA="
    });

    check(http.post(url, data, params), {
        'status is 20': r => r.status == 200
      }) || errorRate.add(1);

//   check(res, {
//     "is status 200": (r) => r.status === 200
//   });
}