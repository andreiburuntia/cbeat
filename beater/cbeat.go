package beater

import (
	"fmt"
	"time"
	"net/http"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/andreiburuntia/cbeat/config"
)

type Cbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Cbeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

var messageQueue = make(chan string)

func test(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for key, _ := range req.Form {
		logp.Info("received JSON: " + key)
		messageQueue <- key
	}
}

func (bt *Cbeat) Run(b *beat.Beat) error {
	logp.Info("cbeat is running! Hit CTRL-C to stop it.")

	go func(){
		http.HandleFunc("/test", test)
		http.ListenAndServe(":9000", nil)
	}()
	logp.Info("here")
	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		var msg = <-messageQueue
		logp.Info(msg)
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
			"msg": msg,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Cbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
