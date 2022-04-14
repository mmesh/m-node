package xlog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"mmesh.dev/m-lib/pkg/errors"
)

const mmeshUserAgent string = "mmesh-xlog/1.0"

type SumologicOptions struct {
	Level    LogLevel
	URL      string
	Name     string // source name (mmesh.namespace)
	Host     string // source host name (mmID)
	Category string // source category (mm.app)
}

type sumologicLogMsg struct {
	level     LogLevel
	timestamp time.Time
	msg       string
}

type sumologicMessages struct {
	lowPriority    []*sumologicLogMsg
	mediumPriority []*sumologicLogMsg
	highPriority   []*sumologicLogMsg
}

type sumologicLogger struct {
	logLevel LogLevel
	url      string
	headers  *http.Header
	client   *http.Client
	logQueue    chan *sumologicLogMsg
	logMessages *sumologicMessages
	endCh       chan struct{}
}

func (l *LoggerSpec) SetSumologicLogger(opts *SumologicOptions) *LoggerSpec {
	l.sumologicLogger = &sumologicLogger{
		logLevel: opts.Level,
		url:      opts.URL,
		headers:  &http.Header{},
		client:   &http.Client{},
		logQueue: make(chan *sumologicLogMsg, 128),
		logMessages: &sumologicMessages{
			lowPriority:    make([]*sumologicLogMsg, 0),
			mediumPriority: make([]*sumologicLogMsg, 0),
			highPriority:   make([]*sumologicLogMsg, 0),
		},
		endCh: make(chan struct{}),
	}

	// set headers
	if len(opts.Name) > 0 {
		l.sumologicLogger.headers.Add("X-Sumo-Name", opts.Name)
	}
	if len(opts.Host) > 0 {
		l.sumologicLogger.headers.Add("X-Sumo-Host", opts.Host)
	}
	if len(opts.Category) > 0 {
		l.sumologicLogger.headers.Add("X-Sumo-Category", opts.Category)
	}

	go l.sumologicLogger.processor()

	return l
}

func (l *LoggerSpec) sumologicLog(level LogLevel, timestamp time.Time, msg string) error {
	l.sumologicLogger.logQueue <- &sumologicLogMsg{
		level:     level,
		timestamp: timestamp,
		msg:       msg,
	}

	return nil
}

func (sml *sumologicLogger) processor() {
	for {
		select {
		case m := <-sml.logQueue:
			var logMessages []*sumologicLogMsg
			prio := logPriorities[m.level]
			switch prio {
			case LOW:
				sml.logMessages.lowPriority = append(sml.logMessages.lowPriority, m)
				logMessages = sml.logMessages.lowPriority
			case MEDIUM:
				sml.logMessages.mediumPriority = append(sml.logMessages.mediumPriority, m)
				logMessages = sml.logMessages.mediumPriority
			case HIGH:
				sml.logMessages.highPriority = append(sml.logMessages.highPriority, m)
				logMessages = sml.logMessages.highPriority
			default:
				prio = LOW
				sml.logMessages.lowPriority = append(sml.logMessages.lowPriority, m)
				logMessages = sml.logMessages.lowPriority
			}
			if err := sml.send(logMessages, prio, false); err != nil {
				tm := time.Now().Format(TIME_FORMAT)
				fmt.Printf("[ERROR] %s %v\n", tm, err)
			}
		case <-time.After(300 * time.Second):
			sml.flushAll()
		case <-sml.endCh:
			sml.flushAll()
			return
		}
	}
}

func (sml *sumologicLogger) flushAll() {
	if err := sml.send(sml.logMessages.highPriority, HIGH, true); err != nil {
		tm := time.Now().Format(TIME_FORMAT)
		fmt.Printf("[ERROR] %s %v\n", tm, err)
	}
	if err := sml.send(sml.logMessages.mediumPriority, MEDIUM, true); err != nil {
		tm := time.Now().Format(TIME_FORMAT)
		fmt.Printf("[ERROR] %s %v\n", tm, err)
	}
	if err := sml.send(sml.logMessages.lowPriority, LOW, true); err != nil {
		tm := time.Now().Format(TIME_FORMAT)
		fmt.Printf("[ERROR] %s %v\n", tm, err)
	}
}

func (sml *sumologicLogger) send(logMessages []*sumologicLogMsg, prio Priority, now bool) error {
	if len(logMessages) == 0 {
		return nil
	}

	if len(logMessages) < 100 && !now {
		return nil
	}

	var payload string
	for _, m := range logMessages {
		prefix := "[" + logPrefixes[m.level] + "]"
		msg := fmt.Sprintf("%s %s %s", prefix, m.timestamp.Format(TIME_FORMAT), m.msg)
		payload = fmt.Sprintf("%s%s\n", payload, msg)
	}

	if len(payload) > 0 {
		headers := &http.Header{}
		p := strings.ToLower(string(prio))
		fields := fmt.Sprintf("priority=%s", p)
		headers.Add("X-Sumo-Fields", fields)

		if err := sml.upload(strings.NewReader(payload), headers); err != nil {
			return errors.Wrapf(err, "[%v] function sml.upload()", errors.Trace())
		}

		switch prio {
		case LOW:
			sml.logMessages.lowPriority = make([]*sumologicLogMsg, 0)
		case MEDIUM:
			sml.logMessages.mediumPriority = make([]*sumologicLogMsg, 0)
		case HIGH:
			sml.logMessages.highPriority = make([]*sumologicLogMsg, 0)
		}
	}

	return nil
}

func (sml *sumologicLogger) upload(payload io.Reader, headers *http.Header) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	u, err := url.ParseRequestURI(sml.url)
	if err != nil {
		return errors.Wrapf(err, "[%v] function url.ParseRequestURI()", errors.Trace())
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), payload)
	if err != nil {
		return errors.Wrapf(err, "[%v] function http.NewRequest()", errors.Trace())
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", mmeshUserAgent)

	if sml.headers != nil {
		for k, v := range *sml.headers {
			req.Header.Add(k, v[0])
		}
	}
	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v[0])
		}
	}

	resp, err := sml.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "[%v] function sml.client.Do()", errors.Trace())
	}
	defer resp.Body.Close()

	if !validResponseStatus(resp.StatusCode) {
		return fmt.Errorf("unexpected response code from Sumologic: %v", resp.StatusCode)
	}

	return nil
}

func validResponseStatus(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}
