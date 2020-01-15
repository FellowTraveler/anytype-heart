module github.com/anytypeio/go-anytype-middleware

go 1.12

require (
	github.com/anytypeio/go-anytype-library v0.0.0-20200115140827-a14699895401

	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.3.1
	github.com/ipfs/go-log v0.0.1
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/stretchr/testify v1.4.0
	github.com/textileio/go-textile v0.7.2-0.20190907000013-95a885123536
	gotest.tools v2.1.0+incompatible
)

replace github.com/textileio/go-textile => github.com/anytypeio/go-textile v0.6.10-0.20191224183538-ba056fbef614

replace github.com/libp2p/go-eventbus => github.com/libp2p/go-eventbus v0.1.0
