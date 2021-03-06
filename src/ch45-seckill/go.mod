module ch45-seckill

go 1.16

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.4.0
	github.com/gohouse/gorose/v2 v2.1.10
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/consul/api v1.3.0
	github.com/juju/ratelimit v1.0.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/prometheus/client_golang v1.3.0
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	github.com/unknwon/com v1.0.1
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/grpc v1.38.0
)

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.29.1
