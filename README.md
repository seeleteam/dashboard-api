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
