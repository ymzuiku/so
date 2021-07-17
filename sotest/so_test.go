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

func dog2() error {
	err1 := errors.New("err2")
	err2 := fmt.Errorf("%w, error2", err1)
	return err2
}

func dogAgain() error {
	err1 := errors.New("err1")
	err2 := fmt.Errorf("%w, error2", err1)
	return err2
}

func cat() error {
	return nil
}

func TestAssetx(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		err := dog()
		so.Error(t, err)
	})

	t.Run("nil", func(t *testing.T) {
		err := cat()
		so.Nil(t, err)
	})

	t.Run("true", func(t *testing.T) {
		err := dog()
		so.True(t, err != nil, err)
	})

	t.Run("false", func(t *testing.T) {
		err := dog()
		so.False(t, err == nil, err)
	})

	t.Run("equal", func(t *testing.T) {
		err := dog()
		err2 := dog()
		so.Equal(t, err, err2)

		err3 := dog()
		err4 := dogAgain()
		so.Equal(t, err3, err4)

		type dog struct {
			name string
			age  int
		}

		var dog1 dog
		dog2 := dog{}
		dog2b := dog{name: "", age: 0}
		so.Equal(t, dog1, dog2)
		so.Equal(t, dog1, dog2b)

		dog3 := dog{name: "a", age: 2}
		dog4 := dog{name: "a", age: 2}
		so.Equal(t, dog3, dog4)

	})

	t.Run("not equal", func(t *testing.T) {
		err := dog()
		err2 := dog2()
		so.NotEqual(t, err, err2)
		type dog struct {
			name string
			age  int
		}

		type cat struct {
			name string
			age  int
		}

		var dog1 dog
		dog2 := dog{name: "1", age: 0}
		so.NotEqual(t, dog1, dog2)

		dog3 := dog{name: "a", age: 2}
		dog4 := dog{name: "ab", age: 2}
		so.NotEqual(t, dog3, dog4)

		var cat1 cat
		cat2 := cat{name: "", age: 0}
		so.NotEqual(t, dog1, cat1)
		so.NotEqual(t, dog1, cat2)
	})

	t.Run("empty", func(t *testing.T) {
		a := 0
		b := ""
		c := false
		so.Empty(t, a)
		so.Empty(t, b)
		so.Empty(t, c)

		type cat struct {
			name string
			age  int
		}
		type dog struct {
			name string
			age  int
			c    cat
			cs   []cat
		}
		var dog1 dog
		so.Empty(t, dog1)

		dog2 := dog{name: "", age: 0, c: cat{name: "", age: 0}}
		so.Empty(t, dog2)
		so.Empty(t, dog2.cs)

		dog3 := dog{name: "", age: 0, c: cat{name: "", age: 0}, cs: []cat{}}
		so.Empty(t, dog3.cs)

		var c1 []cat
		so.Empty(t, c1)
		c2 := []cat{}
		so.Empty(t, c2)

	})

	t.Run("not empty", func(t *testing.T) {
		a := 1
		b := " "
		c := true
		so.NotEmpty(t, a)
		so.NotEmpty(t, b)
		so.NotEmpty(t, c)

		type cat struct {
			name string
			age  int
		}
		type dog struct {
			name string
			age  int
			c    cat
		}

		dog1 := dog{name: "1", age: 0, c: cat{name: "", age: 0}}
		so.NotEmpty(t, dog1)
		dog2 := dog{name: "0", age: 0, c: cat{name: "1", age: 0}}
		so.NotEmpty(t, dog2.c)

		c3 := []cat{{name: "string", age: 20}, {name: "bbb", age: 21}}
		so.NotEmpty(t, c3)
	})

}
