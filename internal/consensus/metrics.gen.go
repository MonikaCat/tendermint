// Code generated by metricsgen. DO NOT EDIT.

package consensus

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Height: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "height",
			Help:      "Height of the chain.",
		}, labels).With(labelsAndValues...),
		ValidatorLastSignedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_last_signed_height",
			Help:      "Last height signed by this validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		Rounds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rounds",
			Help:      "Number of rounds.",
		}, labels).With(labelsAndValues...),
		RoundDuration: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "round_duration",
			Help:      "Histogram of round duration.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, labels).With(labelsAndValues...),
		Validators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validators",
			Help:      "Number of validators.",
		}, labels).With(labelsAndValues...),
		ValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validators_power",
			Help:      "Total power of all validators.",
		}, labels).With(labelsAndValues...),
		ValidatorPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_power",
			Help:      "Power of a validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		ValidatorMissedBlocks: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_missed_blocks",
			Help:      "Amount of blocks missed per validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		MissingValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_validators",
			Help:      "Number of validators who did not sign.",
		}, labels).With(labelsAndValues...),
		MissingValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_validators_power",
			Help:      "Total power of the missing validators.",
		}, labels).With(labelsAndValues...),
		ByzantineValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "byzantine_validators",
			Help:      "Number of validators who tried to double sign.",
		}, labels).With(labelsAndValues...),
		ByzantineValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "byzantine_validators_power",
			Help:      "Total power of the byzantine validators.",
		}, labels).With(labelsAndValues...),
		BlockIntervalSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_interval_seconds",
			Help:      "Time between this and the last block.",
		}, labels).With(labelsAndValues...),
		NumTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "num_txs",
			Help:      "Number of transactions.",
		}, labels).With(labelsAndValues...),
		BlockSizeBytes: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_size_bytes",
			Help:      "Size of the block.",
		}, labels).With(labelsAndValues...),
		TotalTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "total_txs",
			Help:      "Total number of transactions.",
		}, labels).With(labelsAndValues...),
		CommittedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "latest_block_height",
			Help:      "The latest block height.",
		}, labels).With(labelsAndValues...),
		BlockSyncing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_syncing",
			Help:      "Whether or not a node is block syncing. 1 if yes, 0 if no.",
		}, labels).With(labelsAndValues...),
		StateSyncing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "state_syncing",
			Help:      "Whether or not a node is state syncing. 1 if yes, 0 if no.",
		}, labels).With(labelsAndValues...),
		BlockParts: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_parts",
			Help:      "Number of blockparts transmitted by peer.",
		}, append(labels, "peer_id")).With(labelsAndValues...),
		StepDuration: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "step_duration",
			Help:      "Histogram of step duration.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, append(labels, "step")).With(labelsAndValues...),
		BlockGossipReceiveLatency: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_gossip_receive_latency",
			Help:      "Histogram of time taken to receive a block in seconds, measured between when a new block is first discovered to when the block is completed.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, labels).With(labelsAndValues...),
		BlockGossipPartsReceived: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_gossip_parts_received",
			Help:      "Number of block parts received by the node, separated by whether the part was relevant to the block the node is trying to gather or not.",
		}, append(labels, "matches_current")).With(labelsAndValues...),
		QuorumPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "quorum_prevote_delay",
			Help:      "QuroumPrevoteMessageDelay is the interval in seconds between the proposal timestamp and the timestamp of the earliest prevote that achieved a quorum during the prevote step. // To compute it, sum the voting power over each prevote received, in increasing order of timestamp. The timestamp of the first prevote to increase the sum to be above 2/3 of the total voting power of the network defines the endpoint the endpoint of the interval. Subtract the proposal timestamp from this endpoint to obtain the quorum delay.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		FullPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "full_prevote_delay",
			Help:      "FullPrevoteDelay is the interval in seconds between the proposal timestamp and the timestamp of the latest prevote in a round where 100% of the voting power on the network issued prevotes.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		ProposalTimestampDifference: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_timestamp_difference",
			Help:      "ProposalTimestampDifference is the difference between the timestamp in the proposal message and the local time of the validator at the time that the validator received the message.",

			Buckets: []float64{-10, -.5, -.025, 0, .1, .5, 1, 1.5, 2, 10},
		}, append(labels, "is_timely")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Height:                      discard.NewGauge(),
		ValidatorLastSignedHeight:   discard.NewGauge(),
		Rounds:                      discard.NewGauge(),
		RoundDuration:               discard.NewHistogram(),
		Validators:                  discard.NewGauge(),
		ValidatorsPower:             discard.NewGauge(),
		ValidatorPower:              discard.NewGauge(),
		ValidatorMissedBlocks:       discard.NewGauge(),
		MissingValidators:           discard.NewGauge(),
		MissingValidatorsPower:      discard.NewGauge(),
		ByzantineValidators:         discard.NewGauge(),
		ByzantineValidatorsPower:    discard.NewGauge(),
		BlockIntervalSeconds:        discard.NewHistogram(),
		NumTxs:                      discard.NewGauge(),
		BlockSizeBytes:              discard.NewHistogram(),
		TotalTxs:                    discard.NewGauge(),
		CommittedHeight:             discard.NewGauge(),
		BlockSyncing:                discard.NewGauge(),
		StateSyncing:                discard.NewGauge(),
		BlockParts:                  discard.NewCounter(),
		StepDuration:                discard.NewHistogram(),
		BlockGossipReceiveLatency:   discard.NewHistogram(),
		BlockGossipPartsReceived:    discard.NewCounter(),
		QuorumPrevoteDelay:          discard.NewGauge(),
		FullPrevoteDelay:            discard.NewGauge(),
		ProposalTimestampDifference: discard.NewHistogram(),
	}
}
