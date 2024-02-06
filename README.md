# VZX Creative - Network API Client in Go

This is an implementation of the API defined in the [VZX Creative - Network API](https://github.com/vovoid/vzx_udp_client/tree/master) repository, reimplemented in Go.

## Usage

```go get github.com/PixPMusic/goVoxVoid```

Them, you can use the package in your code:

```
import (
    vzxapi "github.com/PixPMusic/goVoxVoid"
)

func main() {
    vzxapi.PlayQueueCurrentItemFxLevelSet(0.0)
}
```

Currently, this is very beta and subject to major changes. Use at your own risk.

## Roadmap

* [x] Implement all API methods
* [ ] Add examples - current example is very bare bones
* [ ] Add error handling
* [ ] Add logging
* [ ] Add configuration - will require refactor, as the current implementation is static
* [ ] Add tests
* [ ] Add documentation
