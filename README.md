# VTB API 2024

## Table of Contents

- [VTB API 2024](#vtb-api-2024)
  - [Table of Contents](#table-of-contents)
  - [Architecture](#architecture)
    - [Generalized architecture](#generalized-architecture)
    - [Client - Server](#client---server)
  - [Project Structure](#project-structure)
  - [Testing](#testing)
  - [Docker Image](#docker-image)
  - [Implementation](#implementation)

## Architecture

### Generalized architecture

![PlantUML Architecture Schema](https://www.plantuml.com/plantuml/svg/L8rDJiGm34RtFeKrvnYDW9aTFohq1COrQTGqhXmNWH1t9mr5OhBCU_xUJN4oN-S4KAQO5pAQSOWcx54p9dbpT4MBONG0ObV3SugI7M2JA8SaU0WltNd_Fo0L3BRmdvPrgCP5UH9hy36oQH5xZopu1LWVwxnehle-rLUwsSVZavYwV5mTdKwjKKNu1XZ8wPLARNHQF9Zqpm-jSxc3Vxltyib9QMheTlNnjm7mqoWPt2eC0q7qxYhSG1huXo-AWIirG11SuCuDA669Y1859RJmjYFMH8fM1YHEAyjlTiOZZpbZKVoxsF4lSaHuw8f1ry7T9D2dCVILcSRnmHD9uYnDm3bnLDGsEoGYL5eosB5DYtHdimrEzC6M2pMDRhZhmSIYvRSzRt6idEo9aIvoHMYe8pU4dbxtkWs6_cRk6q1oqxpViRJ1jAc6CgzuuZXDI9WjbdFxh4y0)

### Client - Server

![Client Server Schema](https://www.plantuml.com/plantuml/svg/RPBVIiCm5CRlynI7tky5nXYJWTCBeUYybj35BTWb8us2YA3pmWkzyW7q5RfXcVXdliB96tcIK5YBlMdFuVj-tvUIGM6viPVpd1KNKYsuuNEOI2CoAxM2N9nRi1gCdCuiOpopsW2-uHQ_t3DkwBt6qYsnyZFm0auBOHXJUY8WIQ_jZ1X7CZEQrbSo4mdSCA0dq_E5La9PIFPvOVVImmHwlMZezpre4RxFl8-8RTFqE1t2C9Sqj8rPsBBRkiL8vrnDMoruqLYqMjK19o7S1unqXCSw7kuF2frElraKkd0m7gU06opnrztXrjspfe8iiI9fySKznRy88H2_iN9Bcj-nP54LCcfuK3NXDNKN6rLR8gqzhoRzKnPVSKMW4bT1_edbZW_BTdKYX7Dtwplmu1wnmX_ymAffP0EX9dIh_W00)

## Project Structure

```tree
├── api                  # proto files
│   └── v1
├── cmd
│   └── grpc_server      # starting point of the project
├── go.mod
├── go.sum
├── internal             # main folder of project
│   ├── api
│   ├── app              # provides server app
│   ├── config           # provides configuration of app
│   ├── converter        # provides converters between different models
│   ├── model            # data transfer objects (entities)
│   ├── repository       # implementation of business logic 
│   └── service          # business logic of this application
├── Makefile             # scripts for developers
├── pkg                  # generated with proto go-files
│   └── v1
├── README.md            # readme-file
└── vendor               # folder for external dependencies
```

## Testing

Before start autotests, you should run `make local`. This command starts
clear database container for autotesting.

Start autotests with command: `

After autotests you can use `make local-down` to stop container.

## Docker Image

You can download Docker Image from [Docker Hub](https://hub.docker.com/r/andytakker/vtb-api-2024-grpc).

To build new image locally use `make docker-build`.

A new image is built every time a new commit gets into master via
[Github Actions](https://github.com/NEROTEX-Team/vtb-api-2024-grpc/blob/master/.github/workflows/ci.yaml).

## Implementation
