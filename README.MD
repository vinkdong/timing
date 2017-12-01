# Timing Request

[![Circle CI](https://circleci.com/gh/VinkDong/TimingRequest/tree/master.svg?style=shield)](https://circleci.com/gh/VinkDong/TimingRequest/tree/master)

## Introduction

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
  
  you can find a example conf file at [`./config.yaml`](https://github.com/VinkDong/TimingRequest/blob/master/config.yaml), config file is under the yaml format you can use follow keys:

  - url: address you want requested
  - method: request method you need
  - body: request body, if you have 2 or more bodies it will send with difference body
  - range: special time period you want to send request, you can use `hour`,`minute`,`second`,`month` and `weekday`, from now,then you can use `gt`,`lt`,`gte`,`lte` and `eq` as condition,they're <span style="color:red">AND</span> relationship,Time is <span style="color:red">UTC</span> time, and Sunday in weekday is `0`, Monday is `1` by parity of reasoning.
  - run_every: interval time for wait next request, you can use `seconds`,`minutes`,`hours` and `days`,if there have 2 or more, they're <span style="color:red">AND</span> relationship too.
  - log_response: whether log response body

  in additional, you can use multi yaml split `---` as multi request entities

- start program

    ```bash
    ./TimingRequest --conf path/your/config
    ```

BTW, your can use `--help` to display help information

## Monitoring with prometheus

you can monitoring this process via prometheus, this is a experimental

- start with prometheus monitoring

    ```bash
    .TimingRequest --conf path/your/config -enable_metrics --addr :9800
    ```

- then your can scape metrics by prometheus

    ```yaml
    scrape_configs:
    - job_name: 'TimingRequest'
    static_configs:
    - targets:
        - '127.0.0.1:9800'
    ```