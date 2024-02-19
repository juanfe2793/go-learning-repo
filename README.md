# go-learning-repo

This is a Repo to learn the dev language Go with some sample apps.

## Weather Command

Weather CLI Tool: sample how to get the weather for your location in the terminal.

You will need an API key from <https://www.weatherapi.com/>

### install

```shell
mv weather /usr/local/bin
```

### usage

```shell
weather -location Cali -key myAwesomeKeyGoesHere
he weather in Cali, Colombia is 28C and Partly cloudy
Hour - Temp, Rain%, Condition
19:00 - 19C, 100%, Patchy rain nearby 
20:00 - 18C, 100%, Patchy rain nearby 
21:00 - 17C, 100%, Patchy rain nearby 
22:00 - 17C, 0%, Partly Cloudy  
23:00 - 18C, 0%, Partly Cloudy
```

## General useful commands

```shell
go mod init foo/var
go build .
go run main.go

```
