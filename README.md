# TimeStream Simple CLI

- This is a sample to operate TimeStream with AWS SDK for Go v2

- You can use both AWS and LocalStack

- References
  - TimeStream Official
    - [Amazon Timestream](https://docs.aws.amazon.com/timestream/latest/developerguide/what-is-timestream.html)
    - [LocalStack TimeStream](https://docs.localstack.cloud/user-guide/aws/timestream/)
  - GitHub
    - [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)
    - [sample_apps/goV2](https://github.com/awslabs/amazon-timestream-tools/tree/mainline/sample_apps/goV2)

## Preparation for AWS

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

## Preparation for LocalStack

- Install LocalStack

```bash
pip install localstack
```

## Build

- build

  ```bash
  make build
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
ts database describe -n sampleDB
```

## Describe Table

```bash
ts table describe -d sampleDB -t sampleTable
```

## Generate Sample Data

```bash
ts preset data -t home > ./sample/home.json
ts preset data -t building > ./sample/building.json
```
