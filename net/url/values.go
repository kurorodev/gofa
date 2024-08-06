package url

import (
	"strconv"
	"unsafe"

	"github.com/anton2920/gofa/alloc"
	"github.com/anton2920/gofa/database"
	"github.com/anton2920/gofa/errors"
	"github.com/anton2920/gofa/slices"
	"github.com/anton2920/gofa/strings"
	"github.com/anton2920/gofa/trace"
)

type Values struct {
	Arena alloc.Arena

	Keys   []string
	Values [][]string
}

func ParseQuery(vs *Values, query string) error {
	t := trace.Begin("")

	var err error

	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")
		if strings.FindChar(key, ';') != -1 {
			err = errors.New("invalid semicolon separator in query")
			continue
		}
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")

		keyBuffer := vs.Arena.NewSlice(len(key))
		n, ok := QueryDecode(keyBuffer, key)
		if !ok {
			if err == nil {
				err = errors.New("invalid key")
			}
			continue
		}
		key = unsafe.String(unsafe.SliceData(keyBuffer), n)

		valueBuffer := vs.Arena.NewSlice(len(value))
		n, ok = QueryDecode(valueBuffer, value)
		if !ok {
			if err == nil {
				err = errors.New("invalid value")
			}
			continue
		}
		value = unsafe.String(unsafe.SliceData(valueBuffer), n)

		vs.Add(key, value)
	}

	trace.End(t)
	return err
}

func (vs *Values) Add(key string, value string) {
	t := trace.Begin("")

	for i := 0; i < len(vs.Keys); i++ {
		if key == vs.Keys[i] {
			vs.Values[i] = append(vs.Values[i], value)

			trace.End(t)
			return
		}
	}
	vs.Keys = append(vs.Keys, key)

	if len(vs.Values) == cap(vs.Values) {
		vs.Values = append(vs.Values, []string{value})

		trace.End(t)
		return
	}
	n := len(vs.Values)
	vs.Values = vs.Values[:n+1]
	vs.Values[n] = vs.Values[n][:0]
	vs.Values[n] = append(vs.Values[n], value)

	trace.End(t)
}

func (vs *Values) Get(key string) string {
	t := trace.Begin("")

	for i := 0; i < len(vs.Keys); i++ {
		if key == vs.Keys[i] {
			trace.End(t)
			return vs.Values[i][0]
		}
	}

	trace.End(t)
	return ""
}

func (vs Values) GetInt(key string) (int, error) {
	t := trace.Begin("")

	n, err := strconv.Atoi(vs.Get(key))

	trace.End(t)
	return n, err
}

func (vs Values) GetID(key string) (database.ID, error) {
	t := trace.Begin("")

	id, err := strconv.Atoi(vs.Get(key))
	if err != nil {
		trace.End(t)
		return -1, err
	}
	if (id < database.MinValidID) || (id > database.MaxValidID) {
		trace.End(t)
		return -1, errors.New("ID out of range")
	}

	trace.End(t)
	return database.ID(id), nil
}

func (vs *Values) GetMany(key string) []string {
	t := trace.Begin("")

	for i := 0; i < len(vs.Keys); i++ {
		if key == vs.Keys[i] {
			trace.End(t)
			return vs.Values[i]
		}
	}

	trace.End(t)
	return nil
}

func (vs *Values) Has(key string) bool {
	t := trace.Begin("")

	for i := 0; i < len(vs.Keys); i++ {
		if key == vs.Keys[i] {
			trace.End(t)
			return true
		}
	}

	trace.End(t)
	return false
}

func (vs *Values) Reset() {
	vs.Keys = vs.Keys[:0]
	vs.Values = vs.Values[:0]
	vs.Arena.Reset()
}

func (vs *Values) Set(key string, value string) {
	t := trace.Begin("")

	for i := 0; i < len(vs.Keys); i++ {
		if key == vs.Keys[i] {
			vs.Values[i] = vs.Values[i][:0]
			vs.Values[i] = append(vs.Values[i], value)

			trace.End(t)
			return
		}
	}
	vs.Keys = append(vs.Keys, key)

	if len(vs.Values) == cap(vs.Values) {
		vs.Values = append(vs.Values, []string{value})

		trace.End(t)
		return
	}
	n := len(vs.Values)
	vs.Values = vs.Values[:n+1]
	vs.Values[n] = vs.Values[n][:0]
	vs.Values[n] = append(vs.Values[n], value)

	trace.End(t)
}

func (vs *Values) SetInt(key string, value int) {
	t := trace.Begin("")

	buffer := vs.Arena.NewSlice(20)
	n := slices.PutInt(buffer, value)
	vs.Set(key, unsafe.String(&buffer[0], n))

	trace.End(t)
}
