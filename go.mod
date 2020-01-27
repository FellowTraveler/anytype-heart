module github.com/anytypeio/go-anytype-middleware

go 1.13

require (
	github.com/anytypeio/go-anytype-library v0.0.0-20200123103712-e18be5ef2d32
	github.com/anytypeio/goldmark v1.1.13-0.20200127143441-60bce0023a7d

	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.3.1
	github.com/ipfs/go-log v0.0.1
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/stretchr/testify v1.4.0
	github.com/textileio/go-textile v0.7.8-0.20200102164400-98b263e32c0c
	gotest.tools v2.1.0+incompatible
)

replace github.com/textileio/go-textile => github.com/anytypeio/go-textile v0.6.10-0.20200113110756-d18f10d572cb

replace github.com/libp2p/go-eventbus => github.com/libp2p/go-eventbus v0.1.0
