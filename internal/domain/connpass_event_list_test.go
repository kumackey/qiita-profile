package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConpassEventList_ToProfileMarkdown(t *testing.T) {
	tests := map[string]struct {
		events []ConnpassEvent
		limit  int
		output string
	}{
		"マークダウンに変換できること": {
			events: []ConnpassEvent{
				connpass("イベントの例1", "https://example.com/1", "2022-02-01T14:59:00+00:00"),
				connpass("イベントの例2", "https://example.com/2", "2022-02-01T15:00:00+00:00"),
			},
			limit:  5,
			output: "\n- Feb 2 [イベントの例2](https://example.com/2)\n- Feb 1 [イベントの例1](https://example.com/1)\n",
		},
		"イベント数を制限できること": {
			events: []ConnpassEvent{
				connpass("イベントの例1", "https://example.com/1", "2022-02-01T14:59:00+00:00"),
				connpass("イベントの例2", "https://example.com/2", "2022-02-01T15:00:00+00:00"),
			},
			limit:  1,
			output: "\n- Feb 2 [イベントの例2](https://example.com/2)\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			markdown := ToMarkdown(test.events, test.limit)
			assert.Equal(t, test.output, markdown)
		})
	}
}

func connpass(title string, link string, startedAt string) ConnpassEvent {
	sa, err := time.Parse(time.RFC3339, startedAt)
	if err != nil {
		panic("connpass, time.Parse failed")
	}
	return ConnpassEvent{title: title, link: link, startedAt: sa}
}