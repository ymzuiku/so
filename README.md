# so

Tiny assert in go, No nonsense log.

## Install

```bash
go get github.com/ymzuiku/so
```

## APIs

- `so.True`
- `so.False`
- `so.Nil`
- `so.NotNil`
- `so.Error`
- `so.Equal`
- `so.NotEqual`
- `so.Empty`
- `so.NotEmpty`

## Use

```go
package sotest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ymzuiku/so"
)

func dog() error {
	err1 := errors.New("err1")
	err2 := fmt.Errorf("%w, error2", err1)
	return err2
}

func cat() error {
	return nil
}

func TestAssetx(t *testing.T) {
	t.Run("assetx error", func(t *testing.T) {
		err := dog()
		so.Error(t, err)
	})

	t.Run("assetx nil", func(t *testing.T) {
		err := cat()
		so.Nil(t, err)
	})

	t.Run("assetx true", func(t *testing.T) {
		err := dog()
		so.True(t, err != nil, err)
	})

	t.Run("assetx false", func(t *testing.T) {
		err := dog()
		so.False(t, err == nil, err)
	})
}

```
