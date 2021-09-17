module github.com/openziti/foundation

go 1.16

// replace github.com/openziti/dilithium => ../dilithium

// replace github.com/kataras/go-events => ../go-events

// replace github.com/michaelquigley/pfxlog => ../pfxlog

require (
	github.com/antlr/antlr4 v0.0.0-20210114010855-d34d2e1c271a
	github.com/biogo/store v0.0.0-20190426020002-884f370e325d
	github.com/emirpasic/gods v1.12.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/influxdata/influxdb1-client v0.0.0-20191209144304-8bf82d3c094d
	github.com/kataras/go-events v0.0.3-0.20201007151548-c411dc70c0a6
	github.com/lucas-clemente/quic-go v0.23.0
	github.com/michaelquigley/pfxlog v0.6.1
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/go-ps v1.0.0
	github.com/openziti/dilithium v0.3.3
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/sirupsen/logrus v1.8.1
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	go.etcd.io/bbolt v1.3.5-0.20200615073812-232d8fc87f50
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)
