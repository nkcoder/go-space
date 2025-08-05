# Description

A CLI script for processing CSV data and sending emails using:
- [Cobra](https://github.com/spf13/cobra): to create CLI application
- [Task](https://taskfile.dev/): for task runner
- [golangci-lint](https://github.com/golangci/golangci-lint): for linting and formatting
- [pgx](https://github.com/jackc/pgx): communicating with PostgreSQL
- [testify](https://github.com/stretchr/testify): for testing


## Configuration

You need to update the variables in `Taskfile.yml` to run for different environments and testing purpose:

```
vars:
  ENV: dev
  AWS_PROFILE: dev
  TEST_EMAIL: "daniel.guo@vivalabs.com.au"
  SENDER: no-reply@the-hub.ai
  PIF_INPUT: data/pif_club_transfer.csv
  DD_INPUT: data/dd_club_transfer.csv
```

## Running with Task locally

```sh
export AWS_PROFILE=xxx

# Build and run PIF transfers
task send-email-pif

# Build and run DD transfers
task send-email-dd

# Run tests
task test

# Run tests with coverage
task test:coverage
```

## Running with Docker

### Run with Docker

```sh
task docker:send-email-pif
task docker:send-email-dd
```

### Run tests in Docker

```sh
task docker:test
```

## Troubleshooting

Error:
```
EBUG: 2025/08/05 10:14:21 Debug logging enabled
INFO:  2025/08/05 10:14:21 Transfer type: PIF, filename: data/pif_club_transfer.csv, env: prod
INFO:  2025/08/05 10:14:21 Loading database configuration from secret: hub-insights-rds-cluster-readonly-prod
INFO:  2025/08/05 10:14:21 Getting secret: hub-insights-rds-cluster-readonly-prod
INFO:  2025/08/05 10:14:22 Successfully connected to database at hub-insights-cluster-prod.cluster-ro-cnfhlchb9rtt.ap-southeast-2.rds.amazonaws.com:5432/hub_insights
INFO:  2025/08/05 10:14:22 Starting club transfer process for type: PIF
ERROR: 2025/08/05 10:14:22 Failed to process club transfers: failed to read club transfer data: error reading club transfer data: column Member Id not found
task: Failed to run task "send-email-pif": exit status 1
```

The project is: before the column `Member Id`, there is leading space.
