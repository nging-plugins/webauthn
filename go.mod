module github.com/nging-plugins/webauthn

go 1.19

replace github.com/admpub/nging/v5 => ../../admpub/nging

//replace github.com/admpub/webx => ../../admpub/webx

replace github.com/coscms/webauthn => ../../coscms/webauthn

require (
	github.com/admpub/log v1.3.3
	github.com/admpub/nging/v5 v5.0.0
	github.com/admpub/useragent v0.0.1
	github.com/coscms/webauthn v0.1.3
	github.com/go-webauthn/webauthn v0.7.0
	github.com/webx-top/com v0.9.1
	github.com/webx-top/db v1.23.14
	github.com/webx-top/echo v2.33.1+incompatible
)

require (
	github.com/abh/errorutil v1.0.0 // indirect
	github.com/admpub/ccs-gm v0.0.3 // indirect
	github.com/admpub/checksum v1.1.0 // indirect
	github.com/admpub/color v1.8.1 // indirect
	github.com/admpub/confl v0.2.2 // indirect
	github.com/admpub/cron v0.0.1 // indirect
	github.com/admpub/decimal v1.3.1 // indirect
	github.com/admpub/email v2.4.1+incompatible // indirect
	github.com/admpub/errors v0.8.2 // indirect
	github.com/admpub/events v1.3.5 // indirect
	github.com/admpub/fasthttp v0.0.5 // indirect
	github.com/admpub/fsnotify v1.5.0 // indirect
	github.com/admpub/gifresize v1.0.2 // indirect
	github.com/admpub/go-bindata-assetfs v0.0.0-20170428090253-36eaa4c19588 // indirect
	github.com/admpub/go-download/v2 v2.1.12 // indirect
	github.com/admpub/go-isatty v0.0.10 // indirect
	github.com/admpub/go-password v0.1.3 // indirect
	github.com/admpub/go-pretty/v6 v6.0.3 // indirect
	github.com/admpub/go-reuseport v0.0.4 // indirect
	github.com/admpub/go-utility v0.0.1 // indirect
	github.com/admpub/godotenv v1.4.2 // indirect
	github.com/admpub/humanize v0.0.0-20190501023926-5f826e92c8ca // indirect
	github.com/admpub/i18n v0.2.1 // indirect
	github.com/admpub/identicon v1.0.2 // indirect
	github.com/admpub/imageproxy v0.9.3 // indirect
	github.com/admpub/imaging v1.5.0 // indirect
	github.com/admpub/ip2region/v2 v2.0.1 // indirect
	github.com/admpub/json5 v0.0.1 // indirect
	github.com/admpub/license_gen v0.1.0 // indirect
	github.com/admpub/mahonia v0.0.0-20151019004008-c528b747d92d // indirect
	github.com/admpub/mail v0.0.0-20170408110349-d63147b0317b // indirect
	github.com/admpub/marmot v0.0.0-20200702042226-2170d9ff59f5 // indirect
	github.com/admpub/mysql-schema-sync v0.2.4 // indirect
	github.com/admpub/null v8.0.4+incompatible // indirect
	github.com/admpub/once v0.0.1 // indirect
	github.com/admpub/pester v0.0.0-20200411024648-005672a2bd48 // indirect
	github.com/admpub/randomize v0.0.2 // indirect
	github.com/admpub/realip v0.0.0-20210421084339-374cf5df122d // indirect
	github.com/admpub/redistore v1.2.1 // indirect
	github.com/admpub/resty/v2 v2.7.0 // indirect
	github.com/admpub/securecookie v1.1.2 // indirect
	github.com/admpub/service v0.0.5 // indirect
	github.com/admpub/sessions v0.1.3 // indirect
	github.com/admpub/sonyflake v0.0.1 // indirect
	github.com/admpub/tail v1.1.0 // indirect
	github.com/admpub/timeago v1.2.1 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coscms/forms v1.12.1 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/fcjr/aia-transport-go v1.2.2 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/garyburd/redigo v1.6.4 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/go-webauthn/revoke v0.1.8 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-tpm v0.3.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/h2non/filetype v1.1.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20230110061619-bbe2e5e100de // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.22 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/muesli/smartcrop v0.3.0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.40.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rwcarlsen/goexif v0.0.0-20190401172101-9e8deecbddbd // indirect
	github.com/segmentio/fasthash v1.0.3 // indirect
	github.com/shirou/gopsutil/v3 v3.23.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/strmangle v0.0.4 // indirect
	github.com/webx-top/captcha v0.1.0 // indirect
	github.com/webx-top/chardet v0.0.1 // indirect
	github.com/webx-top/client v0.8.8 // indirect
	github.com/webx-top/codec v0.2.1 // indirect
	github.com/webx-top/image v0.0.9 // indirect
	github.com/webx-top/pagination v0.2.1 // indirect
	github.com/webx-top/poolx v0.0.0-20210912044716-5cfa2d58e380 // indirect
	github.com/webx-top/restyclient v0.0.3 // indirect
	github.com/webx-top/tagfast v0.0.0-20161020041435-9a2065ce3dd2 // indirect
	github.com/webx-top/validation v0.0.3 // indirect
	github.com/webx-top/validator v0.2.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/image v0.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)
