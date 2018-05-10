package beater

import (
	"fmt"
	"time"
	"net/http"
	"os"
	"log"
	"io"
	"bytes"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/andreiburuntia/cbeat/config"

	"github.com/rjeczalik/notify"
)

const (
    MAX_CONCURRENT_WRITERS = 5
)

var (
    pipePath string
    filePath string
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

func readPipe(){
	var p *os.File
    var err error
    var e notify.EventInfo

    pipePath = "/tmp/cupsbeat"

    // The usual stuff: checking wether the named pipe exists etc
    if p, err = os.Open(pipePath); os.IsNotExist(err) {
        log.Fatalf("Named pipe '%s' does not exist", pipePath)
    } else if os.IsPermission(err) {
        log.Fatalf("Insufficient permissions to read named pipe '%s': %s", pipePath, err)
    } else if err != nil {
        log.Fatalf("Error while opening named pipe '%s': %s", pipePath, err)
    }
    // Yep, there and readable. Close the file handle on exit
    defer p.Close()

    c := make(chan notify.EventInfo, MAX_CONCURRENT_WRITERS)

    notify.Watch(pipePath, c, notify.Write|notify.Remove)


	var buf bytes.Buffer
    // We start an infinite loop...
    for {
        // ...waiting for an event to be passed.
        e = <-c

        switch e.Event() {

        case notify.Write:
			io.Copy(&buf, p)
			newStr := buf.String()
			logp.Info("HIT!")
			logp.Info(newStr)
			messageQueue <- newStr

        case notify.Remove:
            log.Fatalf("Named pipe '%s' was removed. Quitting", pipePath)
        }
    }
}

func (bt *Cbeat) Run(b *beat.Beat) error {
	logp.Info("cbeat is running! Hit CTRL-C to stop it.")

	go func(){
		readPipe()
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
