language: go
go:
  - 1.8.x
env:
  - "PATH=/home/travis/gopath/bin:$PATH"
before_install:
  - make restore
script:
  - make build
// should be added later to publish
// after_success:
//   - make release
