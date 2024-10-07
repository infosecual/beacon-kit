package phuslu

import (
	"io"
	"testing"

	"github.com/phuslu/log"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_Formatter_AddKeyColor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var color Color
		fill_err = tp.Fill(&color)
		if fill_err != nil {
			return
		}

		f1 := NewFormatter()
		f1.AddKeyColor(key, color)
	})
}

func Fuzz_Nosy_Formatter_AddKeyValColor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var val string
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var color Color
		fill_err = tp.Fill(&color)
		if fill_err != nil {
			return
		}

		f1 := NewFormatter()
		f1.AddKeyValColor(key, val, color)
	})
}

func Fuzz_Nosy_Formatter_Format__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var args *log.FormatterArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		f1 := NewFormatter()
		f1.Format(out, args)
	})
}

func Fuzz_Nosy_Formatter_ensureLineBreak__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *byteBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		f1 := NewFormatter()
		f1.ensureLineBreak(b)
	})
}

func Fuzz_Nosy_Formatter_formatHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *log.FormatterArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var b *byteBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var color string
		fill_err = tp.Fill(&color)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		if args == nil || b == nil {
			return
		}

		f1 := NewFormatter()
		f1.formatHeader(args, b, color, label)
	})
}

func Fuzz_Nosy_Formatter_printWithColor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *log.FormatterArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var b *byteBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var color string
		fill_err = tp.Fill(&color)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		if args == nil || b == nil {
			return
		}

		f1 := NewFormatter()
		f1.printWithColor(args, b, color, label)
	})
}

// skipping Fuzz_Nosy_Logger_AddKeyColor__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Logger_AddKeyValColor__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Logger_Debug__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_Logger_Error__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_Logger_Impl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.Impl()
	})
}

// skipping Fuzz_Nosy_Logger_Info__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_Logger_Warn__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_Logger_WithConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var c2 *Config
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		var c3 *Config
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		if c2 == nil || c3 == nil {
			return
		}

		l := NewLogger(out, c2)
		l.WithConfig(c3)
	})
}

func Fuzz_Nosy_Logger_Writer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.Writer()
	})
}

// skipping Fuzz_Nosy_Logger_msgWithContext__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_Logger_setWriter__ because parameters include func, chan, or unsupported interface: github.com/phuslu/log.Writer

func Fuzz_Nosy_Logger_useConsoleWriter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.useConsoleWriter()
	})
}

func Fuzz_Nosy_Logger_useJSONWriter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.useJSONWriter()
	})
}

func Fuzz_Nosy_Logger_withLogLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var level string
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.withLogLevel(level)
	})
}

func Fuzz_Nosy_Logger_withStyle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var style string
		fill_err = tp.Fill(&style)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.withStyle(style)
	})
}

func Fuzz_Nosy_Logger_withTimeFormat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var formatStr string
		fill_err = tp.Fill(&formatStr)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewLogger(out, cfg)
		l.withTimeFormat(formatStr)
	})
}

func Fuzz_Nosy_byteBuffer_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *byteBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Reset()
	})
}

// skipping Fuzz_Nosy_Logger_With__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_parseFormat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, formatStr string) {
		parseFormat(formatStr)
	})
}
