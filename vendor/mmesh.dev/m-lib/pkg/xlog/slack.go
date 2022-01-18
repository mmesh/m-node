package xlog

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/slack-go/slack"
	"mmesh.dev/m-lib/pkg/errors"
)

type SlackOptions struct {
	Level        LogLevel
	Webhook      string
	User         string
	Icon         string
	TraceChannel string
	DebugChannel string
	InfoChannel  string
	WarnChannel  string
	ErrorChannel string
	AlertChannel string
}

type slackLoggerCfg struct {
	webhook  string
	user     string
	icon     string
	logLevel LogLevel
	channels map[LogLevel]string
	colors   map[LogLevel]string
}

func (l *LoggerSpec) SetSlackLogger(opt *SlackOptions) *LoggerSpec {
	l.slackLogger = &slackLoggerCfg{
		webhook:  opt.Webhook,
		user:     opt.User,
		icon:     opt.Icon,
		logLevel: opt.Level,
		channels: map[LogLevel]string{
			TRACE: opt.TraceChannel,
			DEBUG: opt.DebugChannel,
			INFO:  opt.InfoChannel,
			WARN:  opt.WarnChannel,
			ERROR: opt.ErrorChannel,
			ALERT: opt.AlertChannel,
		},
		colors: map[LogLevel]string{
			TRACE: "#ff77ff",
			DEBUG: "#444999",
			INFO:  "#009999",
			WARN:  "#fff000",
			ERROR: "#ff4444",
			ALERT: "#990000",
		},
	}
	return l
}

func (l *LoggerSpec) slackMsgTitle(level LogLevel, timestamp time.Time) string {
	return "[" + l.severity(level) + "] " + timestamp.Format(TIME_FORMAT) + " @" + l.hostID
}

func (l *LoggerSpec) slackLog(level LogLevel, timestamp time.Time, msg string) error {
	if len(l.slackLogger.channels[level]) == 0 {
		return nil
	}

	attachment := slack.Attachment{
		Title:      l.slackMsgTitle(level, timestamp),
		Text:       "```" + msg + "```",
		Color:      l.slackLogger.colors[level],
		AuthorName: l.slackLogger.user,
		AuthorIcon: l.slackLogger.icon,
		Ts:         json.Number(strconv.Itoa(int(timestamp.Unix()))),
		Fields: []slack.AttachmentField{
			{
				Title: "Priority",
				Value: string(l.priority(level)),
				Short: true,
			},
			{
				Title: "Severity",
				Value: l.severity(level),
				Short: true,
			},
			{
				Title: "Timestamp",
				Value: timestamp.Format(time.RFC3339),
				Short: false,
			},
		},
	}

	m := slack.WebhookMessage{
		Username: l.slackLogger.user,
		IconURL:  l.slackLogger.icon,
		Channel:  l.slackLogger.channels[level],
		// Text: msg,
		Attachments: []slack.Attachment{attachment},
		Parse:       "full",
	}

	if err := slack.PostWebhook(l.slackLogger.webhook, &m); err != nil {
		return errors.Wrapf(err, "[%v] function slack.PostWebhook()", errors.Trace())
	}

	return nil
}
