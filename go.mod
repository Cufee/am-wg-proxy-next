module aftermath.link/repo/am-wg-proxy

replace aftermath.link/repo/logs => ../logs

go 1.17

require (
	aftermath.link/repo/logs v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.23.0
	github.com/google/uuid v1.3.0
)

require (
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
)
