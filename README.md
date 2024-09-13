# TimeStream Simple CLI

- This is a sample to operate TimeStream with AWS SDK for Go v2

- References
  - TimeStream Official
    - [Amazon Timestream](https://docs.aws.amazon.com/timestream/latest/developerguide/what-is-timestream.html)
  - GitHub
    - [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)
    - [sample_apps/goV2](https://github.com/awslabs/amazon-timestream-tools/tree/mainline/sample_apps/goV2)

## Preparation

- Add SDK

```bash
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/timestreamquery
go get github.com/aws/aws-sdk-go-v2/service/timestreamwrite
```

- Set AWS credentials

```bash
aws configure
```

- Edit `.envrc` for your environment after copying `.envrc.sample`

```bash
cp .envrc.sample .envrc
```

- Set ENV

```bash
direnv allow
```

[direnv](https://github.com/direnv/direnv)

## Build

- build

  ```bash
  go build -o ./.bin/tt ./
  ```

- set PATH

  ```bash
  export PATH=$PATH:./.bin
  ```

## Create Database

```bash

```

## Describe Database

```bash
tt database describe -n "sampleDB"
```

## List Databases

```bash

```

## Update Database

```bash

```

## Delete Database

```bash

```
