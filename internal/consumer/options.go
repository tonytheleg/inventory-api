package consumer

import (
	"fmt"
	"github.com/spf13/pflag"
)

type Options struct {
	BootstrapServers  string `mapstructure:"bootstrap-servers"`
	ConsumerGroupID   string `mapstructure:"consumer-group-id"`
	Topic             string `mapstructure:"topic"`
	SessionTimeout    string `mapstructure:"session-timeout"`
	HeartbeatInterval string `mapstructure:"heartbeat-interval"`
	MaxPollInterval   string `mapstructure:"max-poll-interval"`
	EnableAutoCommit  string `mapstructure:"enable-auto-commit"`
	AutoOffsetReset   string `mapstructure:"auto-offset-reset"`
	Debug             string `mapstructure:"debug"`
}

func NewOptions() *Options {
	return &Options{
		ConsumerGroupID:   "inventory-consumer",
		Topic:             "outbox.event.kessel.tuples",
		SessionTimeout:    "45000",
		HeartbeatInterval: "3000",
		MaxPollInterval:   "300000",
		EnableAutoCommit:  "false",
		AutoOffsetReset:   "earliest",
		Debug:             "",
	}
}

func (o *Options) AddFlags(fs *pflag.FlagSet, prefix string) {
	if prefix != "" {
		prefix = prefix + "."
	}
	fs.StringVar(&o.BootstrapServers, prefix+"bootstrap-servers", o.BootstrapServers, "sets the bootstrap server address and port for Kafka")
	fs.StringVar(&o.ConsumerGroupID, prefix+"consumer-groupd-id", o.ConsumerGroupID, "sets the Kafka consumer group name (default: inventory-consumer)")
	fs.StringVar(&o.Topic, prefix+"topic", o.Topic, "Kafka topic to monitor for events")
	fs.StringVar(&o.SessionTimeout, prefix+"session-timeout", o.SessionTimeout, "time a consumer can live without sending heartbeat (default: 45000)")
	fs.StringVar(&o.HeartbeatInterval, prefix+"heartbeat-interval", o.HeartbeatInterval, "interval between heartbeats sent to Kafka (default: 3000, must be lower then session-timeout)")
	fs.StringVar(&o.MaxPollInterval, prefix+"max-poll", o.MaxPollInterval, "length of time consumer can go without polling before considered dead (default: 300000)")
}

func (o *Options) Validate() []error {
	var errs []error

	if len(o.BootstrapServers) == 0 {
		errs = append(errs, fmt.Errorf("bootstrap servers can not be empty"))
	}

	return errs
}

func (o *Options) Complete() []error {
	var errs []error

	return errs
}
