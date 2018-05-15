package beater

import (
	"fmt"
	"time"
	"os"
	"log"
	"io"
	"bytes"
	"encoding/json"
	"strconv"
	"io/ioutil"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/andreiburuntia/cbeat/config"

	"github.com/rjeczalik/notify"

	"github.com/andreiburuntia/cbeat/cups_itf"
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

var messageQueue = make(chan Message_s_t)


/*

{"CutMedia": 0, "Duplex": 0, "HWResolution": {"hr1": 600, "hr2": 600}, 
"ImagingBoundingBox": {"ibb1": 0, "ibb2": 0, "ibb3": 595, "ibb4": 842}, "InsertSheet": 0,     
"Jog": 0, "LeadingEdge": 0, "ManualFeed": 0, "MediaPosition": 0, "MediaWeight": 0, "NumCopies": 1, 
"Orientation": 0, "PageSize": {"ps1": 595, "ps2": 842}}, "Tumble": 0,      "cupsWidth": 4958, "cupsHeight": 7017, "cupsBitsPerColor": 8, "cupsBitsPerPixel": 24, 
"cupsColorOrder": 14874, "cupsColorSpace": 0, "cupsNumColors": 1}

*/

type HWResolution_t struct {
	Hr1 int
	Hr2 int
}

type ImaginBoundingBox_t struct {
	Ibb1 int
	Ibb2 int
	Ibb3 int
	Ibb4 int
}

type PageSize_t struct {
	Ps1 int
	Ps2 int
}

type Message_t struct{
	CutMedia int
	Duplex int
	HWResolution HWResolution_t
	ImagingBoundingBox ImaginBoundingBox_t
	InsertSheet int
	Orientation int
	NumCopies int
	PageSize PageSize_t
	Tumble int
	CupsWidth int
	CupsHeight int
	CupsBitsPerColor int
	CupsBitsPerPixel int
	CupsColorOrder int
	CupsColorSpace int
	CupsNumColors int
}

type Message_s_t struct{
	CutMedia string
	Duplex int
	HWResolution HWResolution_t
	ImagingBoundingBox ImaginBoundingBox_t
	InsertSheet int
	Orientation string
	NumCopies int
	PageSize PageSize_t
	Tumble int
	CupsWidth int
	CupsHeight int
	CupsBitsPerColor int
	CupsBitsPerPixel int
	CupsColorOrder string
	CupsColorSpace string
	CupsNumColors int
}

func isJSON(s string) bool {
    var js map[string]interface{}
    return json.Unmarshal([]byte(s), &js) == nil

}

func processMsg(msg string) Message_s_t{
	var maps = cups_itf.Maps
	fmt.Println(maps["Duplex"]["0"])
	logp.Info("PROCESSING STRING:")
	logp.Info(msg)
	var message Message_t
	var message_s Message_s_t
	if isJSON(msg){
		logp.Info("VALID JSON")
		json.Unmarshal([]byte(msg), &message)
		message_s.CutMedia = maps["CutMedia"][strconv.Itoa(message.CutMedia)]
		message_s.Duplex = message_s.Duplex
		message_s.HWResolution = message.HWResolution
		message_s.ImagingBoundingBox = message.ImagingBoundingBox
		message_s.InsertSheet = message.InsertSheet
		message_s.Orientation = maps["Orientation"][strconv.Itoa(message.Orientation)]
		message_s.NumCopies = message.NumCopies
		message_s.PageSize = message.PageSize
		message_s.Tumble = message.Tumble
		message_s.CupsWidth = message.CupsWidth
		message_s.CupsHeight = message.CupsHeight
		message_s.CupsBitsPerColor = message.CupsBitsPerColor
		message_s.CupsBitsPerPixel = message.CupsBitsPerPixel
		message_s.CupsColorOrder = maps["cupsColorOrder"][strconv.Itoa(message.CupsColorOrder)]
		message_s.CupsColorSpace = maps["cupsColorSpace"][strconv.Itoa(message.CupsColorSpace)]
		message_s.CupsNumColors = message.CupsNumColors
		logp.Info("CREATED NEW STRUCT")
	}
	return message_s
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
			//newStr := buf.String()
			logp.Info("HIT!")
			//logp.Info(newStr)
			dat, err := ioutil.ReadFile("./test.json")
			if err==nil{}
			var res = processMsg(string(dat))
			messageQueue <- res

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
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"CutMedia": msg.CutMedia,
			"Duplex": msg.Duplex,
			"HWResolutionX": msg.HWResolution.Hr1,
			"HWResolutionY": msg.HWResolution.Hr2,
			"ImagingBoundingBox1": msg.ImagingBoundingBox.Ibb1,
			"ImagingBoundingBox2": msg.ImagingBoundingBox.Ibb2,
			"ImagingBoundingBox3": msg.ImagingBoundingBox.Ibb3,
			"ImagingBoundingBox4": msg.ImagingBoundingBox.Ibb4,
			"InsertSheet": msg.InsertSheet,
			"Orientation": msg.Orientation,
			"Tumble": msg.Tumble,
			"PageSizeX": msg.PageSize.Ps1,
			"PageSizeY": msg.PageSize.Ps2,
			"NumCopies": msg.NumCopies,
			"ColorSpace": msg.CupsColorSpace,
			"ColorOrder": msg.CupsColorOrder,
			"BitsPerPixel": msg.CupsBitsPerPixel,
			"BitsPerColor": msg.CupsBitsPerColor,
			"NumColors": msg.CupsNumColors,
			"cupsWidth": msg.CupsWidth,
			"cupsHeight": msg.CupsHeight,
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
