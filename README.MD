# Timing Request

## Introduce

Timing Request is used for send http request timing

## Quick Start

- get code

    ```bash
    go get -v github.com/VinkDong/TimingRequest
    ```

- build execute file

    ```bash
    go build
    ```

- edit config file

  config file is under the yaml format you can use follow keys:

  - url: address you want requested
  - method: request method you need
  - body: request body, if you have 2 or more bodies it will send with difference body
  - range: specail time you want to send request, you can use only `hour` from now,then you can use `gt`,`lt`,`gte`,`lte` and `eq` as condition,they're <span style="color:red">AND</span> relationship
  - run_every: interval time for wait next request, you can use `seconds`,`minutes`,`hours` and `days`,if there have 2 or more, they're <span style="color:red">AND</span> relationship too.

- start program

    ```bash
    ./TimingRequest --conf path/your/config
    ```

BTW, your can use `--help` to display help information

## Monitoring with prometheus

you can monitoring this process via prometheus, this is a experimental

- start with prometheus monitoring

    ```bash
    .TimingRequest --conf path/your/config -metrics --port 9800
    ```

- then your can scape metrics by prometheus

    ```yaml
    scrape_configs:
    - job_name: 'TimingRequest'
    static_configs:
    - targets:
        - '127.0.0.1:9800'
    ```