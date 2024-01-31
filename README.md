# Kuma-push

![GitHub Repo stars](https://img.shields.io/github/stars/Nigh/gpt-kook?style=flat&color=ffaaaa)
[![Software License](https://img.shields.io/github/license/Nigh/gpt-kook)](LICENSE)
[![uptime-kuma](https://img.shields.io/badge/Work_with-Uptime--kuma-a8e7bf)](https://github.com/louislam/uptime-kuma)

`Kuma push` is a package for auto sending heartbeat to uptime-kuma service.

## Usage

```go
import (
	"time"
	kuma "github.com/Nigh/kuma-push"
)

func startKuma() {
	k := kuma.New("https://kuma.test.cc/api/push/testToken")
	k.SetInterval(120 * time.Second)
	k.Start()
}
```

You can set three push parameters by calling these funcs

```go
k.SetStatus("statu")
k.SetMsg("msg")
k.SetPing("123")
```

You can set the retry count for calling the push url

```go
k.SetRetry(3)
```

And you can call the stop func to stop the auto heartbeat

```go
k.Stop()
```

The default parameters of new kuma push instance is below:

```go
{
	Status: "up",
	Msg: "ok",
	Interval: 1 * time.Minute,
	Retry: 1
}
```
