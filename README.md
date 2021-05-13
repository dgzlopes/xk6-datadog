# xk6-datadog

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  $ go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/dgzlopes/xk6-datadog@latest
  ```

## Example

```javascript
import datadog from 'k6/x/datadog';

const client = datadog.newClient("<your-api-key>", "<your-app-key>");

export default function() {
    var ingestAvgPerZone = client.metricQuery("service_ingest.kb.avg{*}by{availability-zone}", 3600);

    ingestAvgPerZone.forEach(t => {
        console.log(`${t.expression} with ${t.points.length} data points`);
        console.log(`unit=${t.unit} avg=${t.avg()} min=${t.min()} max=${t.max()}`);
    });

    var dockerThreadCount = client.metricQuery("docker.thread.count{*}");

    dockerThreadCount.forEach(t => {
        console.log(`${t.expression} with ${t.points.length} data points`);
        console.log(`unit=${t.unit} avg=${t.avg()} min=${t.min()} max=${t.max()}`);
    });
}
```

Result output:

```
$ ./k6 run script.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: ../example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0001] service_ingest.kb.avg{availability-zone:us-east-1b} with 16 data points  source=console
INFO[0001] unit=KiB avg=358.32016603090034 min=0 max=497.97467041015625  source=console
INFO[0001] service_ingest.kb.avg{availability-zone:us-east-1a} with 16 data points  source=console
INFO[0001] unit=KiB avg=372.6228407504512 min=0 max=519.2804870605469  source=console
INFO[0001] service_ingest.kb.avg{availability-zone:us-east-1c} with 16 data points  source=console
INFO[0001] unit=KiB avg=258.3870323896408 min=0 max=548.156005859375  source=console
INFO[0001] docker.thread.count{*} with 179 data points   source=console
INFO[0001] unit=unknown avg=11.463663193997492 min=11.303571428571429 max=11.71875  source=console

running (00m00.9s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.9s/10m0s  1/1 iters, 1 per VU

    data_received........: 0 B 0 B/s
    data_sent............: 0 B 0 B/s
    iteration_duration...: avg=905.32ms min=905.32ms med=905.32ms max=905.32ms p(90)=905.32ms p(95)=905.32ms
    iterations...........: 1   1.080766/s
```
