package log

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/inconshreveable/log15"
)

const (
	errorKey       = "LOG15_ERROR"
	timeFormat     = "2006-01-02T15:04:05-0700"
	termTimeFormat = "01-02|15:04:05"
	floatFormat    = 'f'
	termMsgJust    = 40
)

func TerminalFormatEx() log15.Format {
	return log15.FormatFunc(func(r *Record) []byte {
		var color = 0
		switch r.Lvl {
		case LvlCrit:
			color = 35
		case LvlError:
			color = 31
		case LvlWarn:
			color = 33
		case LvlInfo:
			color = 32
		case LvlDebug:
			color = 36
		}

		b := &bytes.Buffer{}
		lvl := strings.ToUpper(r.Lvl.String())
		if color > 0 {
			fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s] %s ", color, lvl, r.Time.Format(termTimeFormat), r.Msg)
		} else {
			fmt.Fprintf(b, "[%s] [%s] %s ", lvl, r.Time.Format(termTimeFormat), r.Msg)
		}

		// try to justify the log output for short messages
		if len(r.Ctx) > 0 && len(r.Msg) < termMsgJust {
			b.Write(bytes.Repeat([]byte{' '}, termMsgJust-len(r.Msg)))
		}

		// print the keys logfmt style
		logfmt(b, r.Ctx, color)
		return b.Bytes()
	})
}

func logfmt(buf *bytes.Buffer, ctx []interface{}, color int) {
	for i := 0; i < len(ctx); i += 2 {
		if i != 0 {
			buf.WriteByte(' ')
		}

		k, ok := ctx[i].(string)
		v := formatLogfmtValue(ctx[i+1])
		if !ok {
			k, v = errorKey, formatLogfmtValue(k)
		}

		// XXX: we should probably check that all of your key bytes aren't invalid
		if color > 0 {
			fmt.Fprintf(buf, "\x1b[%dm%s\x1b[0m=%s", color, k, v)
		} else {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}

	buf.WriteByte('\n')
}

// formatValue formats a value for serialization
func formatLogfmtValue(value interface{}) string {
	if value == nil {
		return "nil"
	}

	if t, ok := value.(time.Time); ok {
		// Performance optimization: No need for escaping since the provided
		// timeFormat doesn't have any escape characters, and escaping is
		// expensive.
		return t.Format(timeFormat)
	}
	value = formatShared(value)
	switch v := value.(type) {
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), floatFormat, 3, 64)
	case float64:
		return strconv.FormatFloat(v, floatFormat, 3, 64)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", value)
	case string:
		return v
	default:
		return fmt.Sprintf("%+v", value)
	}
}

func formatShared(value interface{}) (result interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if v := reflect.ValueOf(value); v.Kind() == reflect.Ptr && v.IsNil() {
				result = "nil"
			} else {
				panic(err)
			}
		}
	}()

	switch v := value.(type) {
	case time.Time:
		return v.Format(timeFormat)

	case error:
		return v.Error()

	case fmt.Stringer:
		return v.String()

	case []string:
		return fmt.Sprintf("%+q", v)

	case []byte:
		return hex.EncodeToString(v)

	case [][]byte:
		str := "[ "
		for _, item := range v {
			str += hex.EncodeToString(item) + " "
		}
		str += "]"
		return str

	default:
		t := reflect.TypeOf(v)
		if t.Kind() != reflect.Slice {
			return v
		}
		s := reflect.ValueOf(v)
		if s.Len() == 0 || !s.Index(0).CanInterface() {
			return v
		}
		switch s.Index(0).Interface().(type) {
		case fmt.Stringer:
			str := "[ "
			for i := 0; i < s.Len(); i++ {
				elem := s.Index(i).Interface()
				str += elem.(fmt.Stringer).String() + " "
			}
			str += "]"
			return str
		}
		return v
	}
}
