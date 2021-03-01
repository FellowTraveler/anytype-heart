module github.com/anytypeio/go-anytype-middleware

go 1.13

require (
	github.com/JohannesKaufmann/html-to-markdown v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/anytypeio/go-slip10 v0.0.0-20200330112030-a352ca8495e4
	github.com/anytypeio/go-slip21 v0.0.0-20200218204727-e2e51e20ab51
	github.com/blevesearch/bleve v1.0.14
	github.com/cheggaaa/mb v1.0.2
	github.com/dave/jennifer v1.4.1
	github.com/dgtony/collections v0.1.5
	github.com/disintegration/imaging v1.6.2
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gobwas/glob v0.2.3
	github.com/goccy/go-graphviz v0.0.9
	github.com/gogo/protobuf v1.3.1
	github.com/gogo/status v1.1.0
	github.com/golang/mock v1.4.4
	github.com/google/martian v2.1.0+incompatible
	github.com/google/uuid v1.1.4
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/h2non/filetype v1.1.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/hsanjuan/ipfs-lite v1.1.17
	github.com/improbable-eng/grpc-web v0.13.0
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.5
	github.com/ipfs/go-ds-badger v0.2.6
	github.com/ipfs/go-ipfs-blockstore v1.0.3
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipld-cbor v0.0.5
	github.com/ipfs/go-ipld-format v0.2.0
	github.com/ipfs/go-log/v2 v2.1.2-0.20200810212702-264b187bb04f
	github.com/ipfs/go-merkledag v0.3.2
	github.com/ipfs/go-path v0.0.7
	github.com/ipfs/go-unixfs v0.2.4
	github.com/ipfs/interface-go-ipfs-core v0.4.0
	github.com/jsummers/gobmp v0.0.0-20151104160322-e2ba15ffa76e
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/libp2p/go-libp2p v0.12.0
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.7.0
	github.com/libp2p/go-libp2p-kad-dht v0.11.1
	github.com/libp2p/go-libp2p-peerstore v0.2.6
	github.com/libp2p/go-libp2p-tls v0.1.3
	github.com/magiconair/properties v1.8.4
	github.com/mauidude/go-readability v0.0.0-20141216012317-2f30b1a346f1
	github.com/mb0/diff v0.0.0-20131118162322-d8d9a906c24d
	github.com/microcosm-cc/bluemonday v1.0.4
	github.com/miolini/datacounter v1.0.2
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-base32 v0.0.3
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multihash v0.0.14
	github.com/otiai10/opengraph v1.1.3
	github.com/prometheus/client_golang v1.9.0
	github.com/rwcarlsen/goexif v0.0.0-20190401172101-9e8deecbddbd
	github.com/stretchr/testify v1.7.0
	github.com/textileio/go-threads v1.0.2-0.20201217154614-3a79caa4def7
	github.com/tyler-smith/go-bip39 v1.0.1-0.20190808214741-c55f737395bc
	github.com/yuin/goldmark v1.3.1
	go.uber.org/zap v1.16.0
	golang.org/x/text v0.3.5
	google.golang.org/grpc v1.34.0
	gopkg.in/Graylog2/go-gelf.v2 v2.0.0-20180125164251-1832d8546a9f
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

replace github.com/JohannesKaufmann/html-to-markdown => github.com/anytypeio/html-to-markdown v0.0.0-20200617145221-2afd2a14bae1

replace github.com/textileio/go-threads => github.com/anytypeio/go-threads v1.0.2-0.20201217154614-3a79caa4def7

replace github.com/ipfs/go-log/v2 => github.com/anytypeio/go-log/v2 v2.1.2-0.20200810212702-264b187bb04f

replace gopkg.in/Graylog2/go-gelf.v2 => github.com/anytypeio/go-gelf v0.0.0-20200813115635-198b2af80f88
