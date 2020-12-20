package enum

var (
	// Special index returned for non-member enums. Notice that the member
	// indexes always start from 1, hence 0 is never used.
	NoRecord uint = 0
)

// New creates an Enum with the given values. The items in the Enum will be indexed sequentially in the order
// of values, starting at 1. For instance:
//
//	// one is indexed at 1
//	// two is indexed at 2
//	// three is indexed at 3
//	New("one", "two", "three")
//
// This constructor fits the most use cases where a value is assigned a single enumeration. For instance:
//
//	var someType = New("one", "two", "three")
//	var some = "one" // or "two", or "three"
//	var indexOfOne = someType.Index(some)
//	var _some = someType.Value(indexOfOne)
//
// This constructor does NOT cover cases where a value may correspond to several values of enumerations, such as
// in a "Set" data structure. See NewComposite for this case.
func New(values ...string) *Enum {
	return newEnum(func(i int) uint {
		return uint(i + 1)
	}, values...)
}

// NewComposite creates an Enum with the given values. The items in the Enum will be indexed sequentially in the
// order of 2. For instance:
//
//	// one is indexed at 1
//	// two is indexed at 2
//	// three is indexed at 4
//	NewComposite("one", "two", "three")
//
// This constructor fits the most use cases where a value is assigned one or more enumeration. For instance:
//
//	var someType = NewComposite("one", "two", "three")
//	var some = []string{"one", "three"}
//	var bitmap = someType.BitMap(some...)	// bitmap == 5
//	var values = someType.Hydrate(bitmap)	// values == ["one", "three"]
func NewComposite(values ...string) *Enum {
	return newEnum(func(i int) uint {
		return 1 << i
	}, values...)
}

// NewCustom creates an Enum with the given values and the indexFunc will be used to index each value at position i.
func NewCustom(values []string, indexFunc func(int) uint) *Enum {
	return newEnum(indexFunc, values...)
}

func newEnum(indexFunc func(int) uint, values ...string) *Enum {
	if len(values) == 0 {
		panic("at least one value is expected")
	}

	var records []*record
	for i, value := range values {
		records = append(records, &record{
			value: value,
			index: indexFunc(i),
		})
	}

	enum := &Enum{
		p: map[string]*record{},
		q: map[uint]*record{},
	}

	for _, r := range records {
		if _, ok := enum.p[r.value]; !ok {
			enum.p[r.value] = r
			enum.q[r.index] = r
		}
	}

	return enum
}

// Enum is the enumeration object, which holds one or more string values. Each string value is assigned a unique
// uint index and allows for quick lookup. An Enum is read-only after creation.
type Enum struct {
	p map[string]*record
	q map[uint]*record
}

// Index returns the index of the given value. If the value does not exist in the Enum, a special NoRecord index
// is returned.
func (e *Enum) Index(value string) uint {
	if r, ok := e.p[value]; ok {
		return r.index
	}
	return NoRecord
}

// Value returns the value that corresponds to the given index. If the provided index does not map to a value, empty
// string is returned.
func (e *Enum) Value(index uint) string {
	v, ok := e.ValueOK(index)
	if !ok {
		return ""
	}
	return v
}

// ValueOK is an alternative to Value in that it explicitly returns a boolean value indicating whether is value
// is found. This is useful when the member values include the empty string
func (e *Enum) ValueOK(index uint) (string, bool) {
	r, ok := e.q[index]
	if !ok {
		return "", ok
	}
	return r.value, true
}

// BitMap returns a bit map of given values in the Enum. Caution that only Enum created with NewComposite should use
// this method, as other Enum could have applied overlapping indexed which result in a irreversible bitmap.
func (e *Enum) BitMap(values ...string) uint {
	var bits bitmap

	for _, value := range values {
		if i := e.Index(value); i != NoRecord {
			bits.Set(bitmap(i))
		}
	}

	return bits.UInt()
}

// Hydrate reverses the computed bitmap and returns the corresponding values. Caution that only Enum created with
// NewComposite should use this method.
func (e *Enum) Hydrate(bits uint) []string {
	var values []string

	var i uint
	for i = 1; i <= bits; i <<= 1 {
		if bitmap(bits).Has(bitmap(i)) {
			values = append(values, e.q[i].value)
		}
	}

	return values
}

// Contains returns true if all the given values are registered with Enum.
func (e *Enum) Contains(values ...string) bool {
	for _, value := range values {
		if _, ok := e.p[value]; !ok {
			return false
		}
	}
	return true
}

type record struct {
	value string
	index uint
}
