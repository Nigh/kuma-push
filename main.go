package kumapush

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	TimeoutMS int64 = 99999
)

type KumaPush struct {
	BaseURL          string
	Status           string
	Msg              string
	Ping             string
	Interval         time.Duration
	Retry            int
	lastResponseTime int64
	stop             chan int
}

func New(url string) (k *KumaPush) {
	k = &KumaPush{BaseURL: url, Status: "up", Msg: "ok", lastResponseTime: 0, Interval: 1 * time.Minute, Retry: 1}
	return
}

func (k *KumaPush) Stop() {
	k.stop <- 0
}
func (k *KumaPush) Start() error {
	k.stop = make(chan int, 1)
	if k.BaseURL == "" {
		return errors.New("push URL is empty")
	}
	go func() {
		interval := time.NewTicker(k.Interval)
		defer interval.Stop()
		for {
			k.push()
			select {
			case <-k.stop:
				return
			case <-interval.C:
				continue
			}
		}
	}()
	return nil
}

func (k KumaPush) url() string {
	params := url.Values{}
	params.Add("status", k.Status)
	params.Add("msg", k.Msg)
	if k.Ping == "" {
		params.Add("ping", strconv.FormatInt(k.lastResponseTime, 10))
	} else {
		params.Add("ping", k.Ping)
	}
	return k.BaseURL + "?" + params.Encode()
}

func (k *KumaPush) push() {
	retried := 0
	for retried <= k.Retry {
		retried++
		start := time.Now()
		client := http.Client{Timeout: 10 * time.Second}
		_, err := client.Get(k.url())
		if err == nil {
			k.lastResponseTime = time.Since(start).Milliseconds()
			retried = k.Retry + 1
		} else {
			k.lastResponseTime = TimeoutMS
		}
	}
}

func (k *KumaPush) SetStatus(status string) {
	k.Status = status
}
func (k *KumaPush) SetMsg(msg string) {
	k.Msg = msg
}
func (k *KumaPush) SetPing(ping string) {
	k.Ping = ping
}
func (k *KumaPush) SetInterval(interval time.Duration) {
	k.Interval = interval
}
func (k *KumaPush) SetRetry(retry int) {
	k.Retry = retry
}
