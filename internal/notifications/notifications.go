package notifications

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Opts struct {
	ServiceName  string
	DeploymentId string
	Reason       string
	Message      string
}

func Info(opts Opts) Notification {
	return Notification{
		Level:         "info",
		Source:        opts.ServiceName,
		Time:          time.Now().Unix(),
		Message:       opts.Message,
		Reason:        opts.Reason,
		TransactionId: opts.DeploymentId,
	}
}

func Error(opts Opts) Notification {
	return Notification{
		Level:         "error",
		Source:        opts.ServiceName,
		Time:          time.Now().Unix(),
		Message:       opts.Message,
		Reason:        opts.Reason,
		TransactionId: opts.DeploymentId,
	}
}

type Notification struct {
	Level         string `json:"level"`
	Time          int64  `json:"time"`
	Message       string `json:"message"`
	Source        string `json:"source"`
	Reason        string `json:"reason"`
	TransactionId string `json:"deploymentId"`
}

type Notifier struct {
	logServiceUrl string
	httpClient    *http.Client
}

func NewNotifier(logServiceUrl string) *Notifier {
	return &Notifier{logServiceUrl: logServiceUrl}
}

func (n *Notifier) Send(evt Notification) error {
	dat, err := json.Marshal(&evt)
	if err != nil {
		return fmt.Errorf("sending notification (deploymentId:%s): %w", evt.TransactionId, err)
	}

	ctx, cncl := context.WithTimeout(context.Background(), time.Second*40)
	defer cncl()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, n.logServiceUrl, bytes.NewBuffer(dat))
	if err != nil {
		return fmt.Errorf("sending notification (deploymentId:%s): %w", evt.TransactionId, err)
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending notification (deploymentId:%s): %w", evt.TransactionId, err)
	}

	return nil
}
