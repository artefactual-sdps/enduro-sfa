module github.com/artefactual-labs/enduro

go 1.13

require (
	github.com/GeertJohan/go.rice v1.0.1-0.20190430230923-c880e3cd4dd8
	github.com/anmitsu/go-shlex v0.0.0-20161002113705-648efa622239 // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/atrox/go-migrate-rice v1.0.1
	github.com/aws/aws-sdk-go v1.25.30
	github.com/bmizerany/perks v0.0.0-20141205001514-d9a9656a3a4b // indirect
	github.com/cenkalti/backoff/v3 v3.0.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fatih/structtag v1.1.0 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/frankban/quicktest v1.5.0 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.1
	github.com/go-redis/redis/v7 v7.0.0-beta.4.0.20190923123950-4b6ad6a95349
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/googleapis v1.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gogo/status v1.1.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.7.0
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/schema v1.1.0
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/oklog/run v1.0.0
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pierrec/lz4 v2.3.0+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/prashantv/protectmem v0.0.0-20171002184600-e20412882b3a // indirect
	github.com/prometheus/client_golang v1.2.1
	github.com/radovskyb/watcher v1.0.7
	github.com/robfig/cron v1.2.0 // indirect
	github.com/samuel/go-thrift v0.0.0-20190219015601-e8b6b52668fe // indirect
	github.com/spf13/afero v1.2.2
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.5.0
	github.com/streadway/quantile v0.0.0-20150917103942-b0c588724d25 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber-go/mapdecode v1.0.0 // indirect
	github.com/uber-go/tally v3.3.13+incompatible
	github.com/uber/jaeger-client-go v2.17.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	github.com/uber/tchannel-go v1.16.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.opencensus.io v0.22.1 // indirect
	go.uber.org/cadence v0.9.3
	go.uber.org/dig v1.7.0 // indirect
	go.uber.org/fx v1.9.0 // indirect
	go.uber.org/goleak v0.10.0 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/net/metrics v1.2.0 // indirect
	go.uber.org/thriftrw v1.20.2 // indirect
	go.uber.org/yarpc v1.42.0
	go.uber.org/zap v1.12.0
	goa.design/goa v2.0.7+incompatible
	goa.design/goa/v3 v3.0.7
	goa.design/plugins/v3 v3.0.7
	gocloud.dev v0.17.0
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0 // indirect
	golang.org/x/sys v0.0.0-20191105231009-c1f44814a5cd // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20191107185733-c07e1c6ef61c // indirect
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6 // indirect
	google.golang.org/grpc v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

// "go.uber.org/cadence" requires it but "go mod" selects "v0.12.0".
// I suspect the problem is in that Thrift tags are not using the "v" prefix.
replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7

// v1.5.0 not released yet!
replace github.com/go-sql-driver/mysql => github.com/go-sql-driver/mysql v1.4.1-0.20191001060945-14bb9c0fc20f
