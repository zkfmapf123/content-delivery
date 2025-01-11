# Content-Delivery

- [x] API
- [x] Kafka
- [x] Consumer Server
- [x] Database

## Docker-Compose

- [docker-compose.yml](./docker-compose.yml) api 서버
- [docker-compose.kafka.yml](./docker-compose.kafka.yml) kafka + ui
- [docker-compose.log.yml](./docker-compose.log.yml) prometheus + loki + grafana + open telemetry

# Huge Traffic Architecture

![asis](./public/asis.png)

- 대용량에 트래픽이 발생하게 될때 -> <b>DB 모든 부하가 발생함</b>
- DB 자체에 장애가 발생한다면 모든시스템에 장애가 발생할 수 있음 (SPOF)

![tobe](./public/tobe.png)
