package p2p

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"

	"github.com/go-kit/kit/metrics"
)

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "p2p"
)

var (
	// valueToLabelRegexp is used to find the golang package name and type name
	// so that the name can be turned into a prometheus label where the characters
	// in the label do not include prometheus special characters such as '*' and '.'.
	valueToLabelRegexp = regexp.MustCompile(`\*?(\w+)\.(.*)`)
)

//go:generate go run github.com/tendermint/tendermint/scripts/metricsgen -struct=Metrics

// Metrics contains metrics exposed by this package.
type Metrics struct {
	// Number of peers.
	Peers metrics.Gauge
	// Number of bytes received from a given peer.
	PeerReceiveBytesTotal metrics.Counter `metrics_labels:"peer_id, chID, message_type"`
	// Number of bytes sent to a given peer.
	PeerSendBytesTotal metrics.Counter `metrics_labels:"peer_id, chID, message_type"`
	// Pending bytes to be sent to a given peer.
	PeerPendingSendBytes metrics.Gauge `metrics_labels:"peer_id"`

	// RouterPeerQueueRecv defines the time taken to read off of a peer's queue
	// before sending on the connection.
	RouterPeerQueueRecv metrics.Histogram

	// RouterPeerQueueSend defines the time taken to send on a peer's queue which
	// will later be read and sent on the connection (see RouterPeerQueueRecv).
	RouterPeerQueueSend metrics.Histogram

	// RouterChannelQueueSend defines the time taken to send on a p2p channel's
	// queue which will later be consued by the corresponding reactor/service.
	RouterChannelQueueSend metrics.Histogram

	// PeerQueueDroppedMsgs defines the number of messages dropped from a peer's
	// queue for a specific flow (i.e. Channel).
	PeerQueueDroppedMsgs metrics.Counter `metrics_labels:"ch_id" metrics_name:"router_channel_queue_dropped_msgs"`

	// PeerQueueMsgSize defines the average size of messages sent over a peer's
	// queue for a specific flow (i.e. Channel).
	PeerQueueMsgSize metrics.Gauge `metrics_labels:"ch_id" metric_name:"router_channel_queue_msg_size"`
}

type MetricsLabelCache struct {
	mtx               *sync.RWMutex
	messageLabelNames map[reflect.Type]string
}

// ValueToMetricLabel is a method that is used to produce a prometheus label value of the golang
// type that is passed in.
// This method uses a map on the Metrics struct so that each label name only needs
// to be produced once to prevent expensive string operations.
func (m *MetricsLabelCache) ValueToMetricLabel(i interface{}) string {
	t := reflect.TypeOf(i)
	m.mtx.RLock()

	if s, ok := m.messageLabelNames[t]; ok {
		m.mtx.RUnlock()
		return s
	}
	m.mtx.RUnlock()

	s := t.String()
	ss := valueToLabelRegexp.FindStringSubmatch(s)
	l := fmt.Sprintf("%s_%s", ss[1], ss[2])
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.messageLabelNames[t] = l
	return l
}

func NewMetricsLabelCache() *MetricsLabelCache {
	return &MetricsLabelCache{
		mtx:               &sync.RWMutex{},
		messageLabelNames: map[reflect.Type]string{},
	}
}
