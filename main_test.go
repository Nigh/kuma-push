package kumapush

import (
	"testing"
	"time"
)

func TestPushAndStop(t *testing.T) {
	k := New("https://kuma.test.cc/api/push/testToken")
	k.SetInterval(20 * time.Second)
	k.Start()
	<-time.After(2 * time.Minute)
	k.Stop()
	<-time.After(40 * time.Second)
}
