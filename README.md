# Dashboard API

## Env

>recommend

```text
go 1.10+
```

>needed

```text
influxdb 1.5.1+
```

## Project structure

```text
┌── api
├── cmd
├── common
├── db
├── log
└── vendor
```

## Build

```shell
cmd dashboard-api
# generate the executable file
make
# execute the file
build/dashboard-api start -c <configfile>
```

## Tip

The format column `time` in result series can be changed by parameter `precision`.

Precision format like `rfc3339|h|m|s|ms|u|ns`
