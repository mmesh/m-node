// Copyright (C) 2019 x6a
//
// pkg is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// pkg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with pkg. If not, see <http://www.gnu.org/licenses/>.

package xlog

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/nlopes/slack"
	"x6a.dev/pkg/errors"
)

func (l *logger) slackMsgTitle(level LogLevel, timestamp time.Time) string {
	return "[" + l.severity(level) + "] " + timestamp.Format(TIME_FORMAT) + " @" + l.hostID
}

func (l *logger) slackLog(level LogLevel, timestamp time.Time, msg string) error {
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
