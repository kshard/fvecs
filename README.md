# fvecs - vector file formats

The library implement codec for .bvecs, .fvecs and .ivecs vector file formats 
developed by [INRIA LEAR and TEXMEX groups](http://corpus-texmex.irisa.fr).

[![Version](https://img.shields.io/github/v/tag/kshard/fvecs?label=version)](https://github.com/kshard/fvecs/releases)
[![Documentation](https://pkg.go.dev/badge/github.com/kshard/fvecs)](https://pkg.go.dev/github.com/kshard/fvecs)
[![Build Status](https://github.com/kshard/fvecs/workflows/test/badge.svg)](https://github.com/kshard/fvecs/actions/)
[![Git Hub](https://img.shields.io/github/last-commit/kshard/fvecs.svg)](https://github.com/kshard/fvecs)
[![Coverage Status](https://coveralls.io/repos/github/kshard/fvecs/badge.svg?branch=main)](https://coveralls.io/github/kshard/fvecs?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kshard/fvecs)](https://goreportcard.com/report/github.com/kshard/fvecs)

## Inspiration

Vector file formats (.bvecs, .fvecs and .ivecs) is a serialized sequence of vectors `[]byte`, `[]float32` or `[]int32`.

Each vector is

```
      32bit          ⟨𝒅⟩ * n-bits 
|--------------|-------/ ... /-------|
 ⟨𝒅⟩ dimension
```

- .fvecs stores 32-bit float (n = 32bit)
- .ivecs stores 32-bit signed or unsigned integers (n = 32bits)
- .bvecs stores bytes (n = 8bit)

## Getting started

The latest version of the library is available at `main` branch of this repository. All development, including new features and bug fixes, take place on the `main` branch using forking and pull requests as described in contribution guidelines. The stable version is available via Golang modules.

```go
import "github.com/kshard/fvecs"

//
// Reading vectors
r, err := os.Open("siftsmall_base.fvecs")
d := fvecs.NewDecoder[float32](r)

for {
  v, err := d.Read()
  if err != nil {
    break
  }
}

//
// Writing vectors
w, err := os.Create("siftsmall_base.fvecs")
e := fvecs.NewEncoder[float32](w)

for _, v := range /* source of []float32 vectors */ {
  err := e.Write(v)
  if err != nil {
    break
  }
}
```

## How To Contribute

The library is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

The build and testing process requires [Go](https://golang.org) version 1.13 or later.

**build** and **test** library.

```bash
git clone https://github.com/kshard/fvecs
cd fvecs
go test
```

### commit message

The commit message helps us to write a good release note, speed-up review process. The message should address two question what changed and why. The project follows the template defined by chapter [Contributing to a Project](http://git-scm.com/book/ch5-2.html) of Git book.

### bugs

If you experience any issues with the library, please let us know via [GitHub issues](https://github.com/kshard/fvecs/issue). We appreciate detailed and accurate reports that help us to identity and replicate the issue. 


## License

[![See LICENSE](https://img.shields.io/github/license/kshard/fvecs.svg?style=for-the-badge)](LICENSE)