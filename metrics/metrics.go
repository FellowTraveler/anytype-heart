package metrics

import (
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/textileio/go-threads/metrics"
	"net/http"
	"os"
	"sync"
	"time"
)

var log = logging.Logger("anytype-logger")

type ThreadsMetrics struct {}

func (t *ThreadsMetrics) AcceptRecord(tp metrics.RecordType, isNAT bool) {
	if tp == metrics.RecordTypePush {
		if isNAT {
			PushRecordAcceptedNAT.Inc()
		} else {
			PushRecordAccepted.Inc()
		}
	} else {
		if isNAT {
			GetRecordAcceptedNAT.Inc()
		} else {
			GetRecordAccepted.Inc()
		}
	}
}

var (
	Enabled       bool
	once          sync.Once
	ServedThreads = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "threads_total",
		Help:      "Number of served threads",
	})

	PushRecordAcceptedNAT = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "go_threads",
		Name:      "push_record_accepted_nat",
		Help:      "Number of push records accepted when under NAT",
	})

	GetRecordAcceptedNAT = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "go_threads",
		Name:      "get_record_accepted_nat",
		Help:      "Number of get records accepted when under NAT (updateRecordsFromPeer)",
	})

	PushRecordAccepted = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "go_threads",
		Name:      "push_record_accepted",
		Help:      "Number of push records accepted",
	})

	GetRecordAccepted = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "go_threads",
		Name:      "get_record_accepted",
		Help:      "Number of get records accepted (updateRecordsFromPeer)",
	})

	ChangeCreatedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "change_created",
		Help:      "Number of changes created",
	})

	ChangeReceivedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "change_received",
		Help:      "Number of changes received",
	})

	ExternalThreadReceivedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "external_thread_received",
		Help:      "New thread received",
	})

	ExternalThreadHandlingAttempts = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "external_thread_handling_attempts",
		Help:      "New thread handling attempts",
	})

	ExternalThreadHandlingDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "external_thread_handling_seconds",
		Help:      "New thread successfully handling duration",
		Buckets: MetricTimeBuckets([]time.Duration{
			256 * time.Millisecond,
			512 * time.Millisecond,
			1024 * time.Millisecond,
			2 * time.Second,
			4 * time.Second,
			8 * time.Second,
			16 * time.Second,
			30 * time.Second,
			45 * time.Second,
			60 * time.Second,
			90 * time.Second,
			120 * time.Second,
			180 * time.Second,
			240 * time.Second,
		}),
	})

	ThreadAdded = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "thread_added",
		Help:      "New thread added",
	})

	ThreadAddReplicatorAttempts = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "thread_add_replicator_attempts",
		Help:      "New thread handling attempts",
	})

	ThreadAddReplicatorDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "thread_add_replicator_seconds",
		Help:      "New thread successfully handling duration",
		Buckets: MetricTimeBuckets([]time.Duration{
			256 * time.Millisecond,
			512 * time.Millisecond,
			1024 * time.Millisecond,
			2 * time.Second,
			4 * time.Second,
			8 * time.Second,
			16 * time.Second,
			30 * time.Second,
			45 * time.Second,
			60 * time.Second,
			90 * time.Second,
			120 * time.Second,
			180 * time.Second,
			240 * time.Second,
		}),
	})

	ObjectFTUpdatedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "fulltext_index_updated",
		Help:      "Fulltext updated for an object",
	})
	ObjectDetailsUpdatedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "details_index_updated",
		Help:      "Details updated for an object",
	})
	ObjectRelationsUpdatedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "anytype",
		Subsystem: "mw",
		Name:      "relations_index_updated",
		Help:      "Relations updated for an object",
	})
)

func registerPrometheusExpvars() {
	expvarCollector := prometheus.NewExpvarCollector(map[string]*prometheus.Desc{
		"badger_blocked_puts_total": prometheus.NewDesc(
			"badger_blocked_puts_total",
			"badger_blocked_puts_total",
			nil, nil,
		),
		"badger_lsm_size_bytes": prometheus.NewDesc(
			"badger_lsm_size_bytes",
			"badger_lsm_size_bytes",
			[]string{"dir"}, nil,
		),
		"badger_vlog_size_bytes": prometheus.NewDesc(
			"badger_vlog_size_bytes",
			"badger_vlog_size_bytes",
			[]string{"dir"}, nil,
		),
		"badger_pending_writes_total": prometheus.NewDesc(
			"badger_pending_writes_total",
			"badger_pending_writes_total",
			[]string{"dir"}, nil,
		),
		"badger_disk_reads_total": prometheus.NewDesc(
			"badger_disk_reads_total",
			"badger_disk_reads_total",
			nil, nil,
		),
		"badger_disk_writes_total": prometheus.NewDesc(
			"badger_disk_writes_total",
			"badger_disk_writes_total",
			nil, nil,
		),
		"badger_read_bytes": prometheus.NewDesc(
			"badger_read_bytes",
			"badger_read_bytes",
			nil, nil,
		),
		"badger_written_bytes": prometheus.NewDesc(
			"badger_written_bytes",
			"badger_written_bytes",
			nil, nil,
		),
		"badger_lsm_level_gets_total": prometheus.NewDesc(
			"badger_lsm_level_gets_total",
			"badger_lsm_level_gets_total",
			[]string{"level"}, nil,
		),
		"badger_lsm_bloom_hits_total": prometheus.NewDesc(
			"badger_lsm_bloom_hits_total",
			"badger_lsm_bloom_hits_total",
			[]string{"level"}, nil,
		),
		"badger_gets_total": prometheus.NewDesc(
			"badger_gets_total",
			"badger_gets_total",
			nil, nil,
		),
		"badger_puts_total": prometheus.NewDesc(
			"badger_puts_total",
			"badger_puts_total",
			nil, nil,
		),
		"badger_memtable_gets_total": prometheus.NewDesc(
			"badger_memtable_gets_total",
			"badger_memtable_gets_total",
			nil, nil,
		),
	})

	prometheus.MustRegister(expvarCollector)
}

func runPrometheusHttp(addr string) {
	once.Do(func() {
		registerPrometheusExpvars()
		// Create a HTTP server for prometheus.
		httpServer := &http.Server{Handler: promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}), Addr: addr}
		Enabled = true

		// Start your http server for prometheus.
		go func() {
			if err := httpServer.ListenAndServe(); err != nil {
				Enabled = false
				log.Errorf("Unable to start a prometheus http server.")
			}
		}()
	})
}

func MetricTimeBuckets(scale []time.Duration) []float64 {
	buckets := make([]float64, len(scale))
	for i, b := range scale {
		buckets[i] = b.Seconds()
	}
	return buckets
}

func init() {
	if addr := os.Getenv("ANYTYPE_PROM"); addr != "" {
		runPrometheusHttp(addr)
	}
}
