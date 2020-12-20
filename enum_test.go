package enum_test

import (
	"fmt"
	"github.com/imulab/go-enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnum_Index(t *testing.T) {
	e := enum.New("one", "two", "three")

	assert.Equal(t, uint(1), e.Index("one"))
	assert.Equal(t, uint(2), e.Index("two"))
	assert.Equal(t, uint(3), e.Index("three"))
	assert.Equal(t, enum.NoRecord, e.Index("invalid"))
}

func TestEnum_Value(t *testing.T) {
	e := enum.New("one", "two", "three")

	assert.Equal(t, "one", e.Value(1))
	assert.Equal(t, "two", e.Value(2))
	assert.Equal(t, "three", e.Value(3))
	assert.Empty(t, e.Value(4))
}

func TestEnum_Hydrate(t *testing.T) {
	e := enum.NewComposite("one", "two", "three")

	cases := [][]string{
		{"one"},
		{"two"},
		{"three"},
		{"one", "two"},
		{"one", "three"},
		{"two", "three"},
		{"one", "two", "three"},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%s/%d", t.Name(), i+1), func(t *testing.T) {
			bits := e.BitMap(c...)
			hydrated := e.Hydrate(bits)
			assert.True(t, assert.ObjectsAreEqual(c, hydrated))
		})
	}
}

func TestEnum_BitMap(t *testing.T) {
	cases := []struct {
		enum   *enum.Enum
		values []string
		expect uint
	}{
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"one"},
			expect: 1,
		},
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"two"},
			expect: 2,
		},
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"three"},
			expect: 4,
		},
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"one", "two"},
			expect: 3,
		},
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"one", "three"},
			expect: 5,
		},
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"two", "three"},
			expect: 6,
		},
		{
			enum:   enum.NewComposite("one", "two", "three"),
			values: []string{"one", "two", "three"},
			expect: 7,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%s/%d", t.Name(), i+1), func(t *testing.T) {
			bits := c.enum.BitMap(c.values...)
			assert.Equal(t, c.expect, bits)
		})
	}
}

func TestEnum_Contains(t *testing.T) {
	cases := []struct {
		enum   *enum.Enum
		values []string
		expect bool
	}{
		{
			enum:   enum.New("one", "two"),
			values: []string{"one"},
			expect: true,
		},
		{
			enum:   enum.New("one", "two"),
			values: []string{"one", "two"},
			expect: true,
		},
		{
			enum:   enum.New("one", "two"),
			values: []string{"three"},
			expect: false,
		},
		{
			enum:   enum.NewComposite("one", "two"),
			values: []string{"one"},
			expect: true,
		},
		{
			enum:   enum.NewComposite("one", "two"),
			values: []string{"one", "two"},
			expect: true,
		},
		{
			enum:   enum.NewComposite("one", "two"),
			values: []string{"three"},
			expect: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%s/%d", t.Name(), i+1), func(t *testing.T) {
			contains := c.enum.Contains(c.values...)
			assert.Equal(t, c.expect, contains)
		})
	}
}
