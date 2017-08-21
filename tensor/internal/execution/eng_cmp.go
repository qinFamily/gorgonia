package execution

import (
	"reflect"

	"github.com/chewxy/gorgonia/tensor/internal/storage"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func (e E) Gt(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is a scalar. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()

		switch {
		case as && bs:
			GtI(at, bt, rt)
		case as && !bs:
			GtSVI(at[0], bt, rt)
		case !as && bs:
			GtVSI(at, bt[0], rt)
		default:
			GtI(at, bt, rt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()

		switch {
		case as && bs:
			GtI8(at, bt, rt)
		case as && !bs:
			GtSVI8(at[0], bt, rt)
		case !as && bs:
			GtVSI8(at, bt[0], rt)
		default:
			GtI8(at, bt, rt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()

		switch {
		case as && bs:
			GtI16(at, bt, rt)
		case as && !bs:
			GtSVI16(at[0], bt, rt)
		case !as && bs:
			GtVSI16(at, bt[0], rt)
		default:
			GtI16(at, bt, rt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()

		switch {
		case as && bs:
			GtI32(at, bt, rt)
		case as && !bs:
			GtSVI32(at[0], bt, rt)
		case !as && bs:
			GtVSI32(at, bt[0], rt)
		default:
			GtI32(at, bt, rt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()

		switch {
		case as && bs:
			GtI64(at, bt, rt)
		case as && !bs:
			GtSVI64(at[0], bt, rt)
		case !as && bs:
			GtVSI64(at, bt[0], rt)
		default:
			GtI64(at, bt, rt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()

		switch {
		case as && bs:
			GtU(at, bt, rt)
		case as && !bs:
			GtSVU(at[0], bt, rt)
		case !as && bs:
			GtVSU(at, bt[0], rt)
		default:
			GtU(at, bt, rt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()

		switch {
		case as && bs:
			GtU8(at, bt, rt)
		case as && !bs:
			GtSVU8(at[0], bt, rt)
		case !as && bs:
			GtVSU8(at, bt[0], rt)
		default:
			GtU8(at, bt, rt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()

		switch {
		case as && bs:
			GtU16(at, bt, rt)
		case as && !bs:
			GtSVU16(at[0], bt, rt)
		case !as && bs:
			GtVSU16(at, bt[0], rt)
		default:
			GtU16(at, bt, rt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()

		switch {
		case as && bs:
			GtU32(at, bt, rt)
		case as && !bs:
			GtSVU32(at[0], bt, rt)
		case !as && bs:
			GtVSU32(at, bt[0], rt)
		default:
			GtU32(at, bt, rt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()

		switch {
		case as && bs:
			GtU64(at, bt, rt)
		case as && !bs:
			GtSVU64(at[0], bt, rt)
		case !as && bs:
			GtVSU64(at, bt[0], rt)
		default:
			GtU64(at, bt, rt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()

		switch {
		case as && bs:
			GtF32(at, bt, rt)
		case as && !bs:
			GtSVF32(at[0], bt, rt)
		case !as && bs:
			GtVSF32(at, bt[0], rt)
		default:
			GtF32(at, bt, rt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()

		switch {
		case as && bs:
			GtF64(at, bt, rt)
		case as && !bs:
			GtSVF64(at[0], bt, rt)
		case !as && bs:
			GtVSF64(at, bt[0], rt)
		default:
			GtF64(at, bt, rt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()

		switch {
		case as && bs:
			GtStr(at, bt, rt)
		case as && !bs:
			GtSVStr(at[0], bt, rt)
		case !as && bs:
			GtVSStr(at, bt[0], rt)
		default:
			GtStr(at, bt, rt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Gt", t)
	}
}

func (e E) Gte(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is a scalar. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()

		switch {
		case as && bs:
			GteI(at, bt, rt)
		case as && !bs:
			GteSVI(at[0], bt, rt)
		case !as && bs:
			GteVSI(at, bt[0], rt)
		default:
			GteI(at, bt, rt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()

		switch {
		case as && bs:
			GteI8(at, bt, rt)
		case as && !bs:
			GteSVI8(at[0], bt, rt)
		case !as && bs:
			GteVSI8(at, bt[0], rt)
		default:
			GteI8(at, bt, rt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()

		switch {
		case as && bs:
			GteI16(at, bt, rt)
		case as && !bs:
			GteSVI16(at[0], bt, rt)
		case !as && bs:
			GteVSI16(at, bt[0], rt)
		default:
			GteI16(at, bt, rt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()

		switch {
		case as && bs:
			GteI32(at, bt, rt)
		case as && !bs:
			GteSVI32(at[0], bt, rt)
		case !as && bs:
			GteVSI32(at, bt[0], rt)
		default:
			GteI32(at, bt, rt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()

		switch {
		case as && bs:
			GteI64(at, bt, rt)
		case as && !bs:
			GteSVI64(at[0], bt, rt)
		case !as && bs:
			GteVSI64(at, bt[0], rt)
		default:
			GteI64(at, bt, rt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()

		switch {
		case as && bs:
			GteU(at, bt, rt)
		case as && !bs:
			GteSVU(at[0], bt, rt)
		case !as && bs:
			GteVSU(at, bt[0], rt)
		default:
			GteU(at, bt, rt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()

		switch {
		case as && bs:
			GteU8(at, bt, rt)
		case as && !bs:
			GteSVU8(at[0], bt, rt)
		case !as && bs:
			GteVSU8(at, bt[0], rt)
		default:
			GteU8(at, bt, rt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()

		switch {
		case as && bs:
			GteU16(at, bt, rt)
		case as && !bs:
			GteSVU16(at[0], bt, rt)
		case !as && bs:
			GteVSU16(at, bt[0], rt)
		default:
			GteU16(at, bt, rt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()

		switch {
		case as && bs:
			GteU32(at, bt, rt)
		case as && !bs:
			GteSVU32(at[0], bt, rt)
		case !as && bs:
			GteVSU32(at, bt[0], rt)
		default:
			GteU32(at, bt, rt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()

		switch {
		case as && bs:
			GteU64(at, bt, rt)
		case as && !bs:
			GteSVU64(at[0], bt, rt)
		case !as && bs:
			GteVSU64(at, bt[0], rt)
		default:
			GteU64(at, bt, rt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()

		switch {
		case as && bs:
			GteF32(at, bt, rt)
		case as && !bs:
			GteSVF32(at[0], bt, rt)
		case !as && bs:
			GteVSF32(at, bt[0], rt)
		default:
			GteF32(at, bt, rt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()

		switch {
		case as && bs:
			GteF64(at, bt, rt)
		case as && !bs:
			GteSVF64(at[0], bt, rt)
		case !as && bs:
			GteVSF64(at, bt[0], rt)
		default:
			GteF64(at, bt, rt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()

		switch {
		case as && bs:
			GteStr(at, bt, rt)
		case as && !bs:
			GteSVStr(at[0], bt, rt)
		case !as && bs:
			GteVSStr(at, bt[0], rt)
		default:
			GteStr(at, bt, rt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Gte", t)
	}
}

func (e E) Lt(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is a scalar. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()

		switch {
		case as && bs:
			LtI(at, bt, rt)
		case as && !bs:
			LtSVI(at[0], bt, rt)
		case !as && bs:
			LtVSI(at, bt[0], rt)
		default:
			LtI(at, bt, rt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()

		switch {
		case as && bs:
			LtI8(at, bt, rt)
		case as && !bs:
			LtSVI8(at[0], bt, rt)
		case !as && bs:
			LtVSI8(at, bt[0], rt)
		default:
			LtI8(at, bt, rt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()

		switch {
		case as && bs:
			LtI16(at, bt, rt)
		case as && !bs:
			LtSVI16(at[0], bt, rt)
		case !as && bs:
			LtVSI16(at, bt[0], rt)
		default:
			LtI16(at, bt, rt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()

		switch {
		case as && bs:
			LtI32(at, bt, rt)
		case as && !bs:
			LtSVI32(at[0], bt, rt)
		case !as && bs:
			LtVSI32(at, bt[0], rt)
		default:
			LtI32(at, bt, rt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()

		switch {
		case as && bs:
			LtI64(at, bt, rt)
		case as && !bs:
			LtSVI64(at[0], bt, rt)
		case !as && bs:
			LtVSI64(at, bt[0], rt)
		default:
			LtI64(at, bt, rt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()

		switch {
		case as && bs:
			LtU(at, bt, rt)
		case as && !bs:
			LtSVU(at[0], bt, rt)
		case !as && bs:
			LtVSU(at, bt[0], rt)
		default:
			LtU(at, bt, rt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()

		switch {
		case as && bs:
			LtU8(at, bt, rt)
		case as && !bs:
			LtSVU8(at[0], bt, rt)
		case !as && bs:
			LtVSU8(at, bt[0], rt)
		default:
			LtU8(at, bt, rt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()

		switch {
		case as && bs:
			LtU16(at, bt, rt)
		case as && !bs:
			LtSVU16(at[0], bt, rt)
		case !as && bs:
			LtVSU16(at, bt[0], rt)
		default:
			LtU16(at, bt, rt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()

		switch {
		case as && bs:
			LtU32(at, bt, rt)
		case as && !bs:
			LtSVU32(at[0], bt, rt)
		case !as && bs:
			LtVSU32(at, bt[0], rt)
		default:
			LtU32(at, bt, rt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()

		switch {
		case as && bs:
			LtU64(at, bt, rt)
		case as && !bs:
			LtSVU64(at[0], bt, rt)
		case !as && bs:
			LtVSU64(at, bt[0], rt)
		default:
			LtU64(at, bt, rt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()

		switch {
		case as && bs:
			LtF32(at, bt, rt)
		case as && !bs:
			LtSVF32(at[0], bt, rt)
		case !as && bs:
			LtVSF32(at, bt[0], rt)
		default:
			LtF32(at, bt, rt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()

		switch {
		case as && bs:
			LtF64(at, bt, rt)
		case as && !bs:
			LtSVF64(at[0], bt, rt)
		case !as && bs:
			LtVSF64(at, bt[0], rt)
		default:
			LtF64(at, bt, rt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()

		switch {
		case as && bs:
			LtStr(at, bt, rt)
		case as && !bs:
			LtSVStr(at[0], bt, rt)
		case !as && bs:
			LtVSStr(at, bt[0], rt)
		default:
			LtStr(at, bt, rt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Lt", t)
	}
}

func (e E) Lte(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is a scalar. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()

		switch {
		case as && bs:
			LteI(at, bt, rt)
		case as && !bs:
			LteSVI(at[0], bt, rt)
		case !as && bs:
			LteVSI(at, bt[0], rt)
		default:
			LteI(at, bt, rt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()

		switch {
		case as && bs:
			LteI8(at, bt, rt)
		case as && !bs:
			LteSVI8(at[0], bt, rt)
		case !as && bs:
			LteVSI8(at, bt[0], rt)
		default:
			LteI8(at, bt, rt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()

		switch {
		case as && bs:
			LteI16(at, bt, rt)
		case as && !bs:
			LteSVI16(at[0], bt, rt)
		case !as && bs:
			LteVSI16(at, bt[0], rt)
		default:
			LteI16(at, bt, rt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()

		switch {
		case as && bs:
			LteI32(at, bt, rt)
		case as && !bs:
			LteSVI32(at[0], bt, rt)
		case !as && bs:
			LteVSI32(at, bt[0], rt)
		default:
			LteI32(at, bt, rt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()

		switch {
		case as && bs:
			LteI64(at, bt, rt)
		case as && !bs:
			LteSVI64(at[0], bt, rt)
		case !as && bs:
			LteVSI64(at, bt[0], rt)
		default:
			LteI64(at, bt, rt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()

		switch {
		case as && bs:
			LteU(at, bt, rt)
		case as && !bs:
			LteSVU(at[0], bt, rt)
		case !as && bs:
			LteVSU(at, bt[0], rt)
		default:
			LteU(at, bt, rt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()

		switch {
		case as && bs:
			LteU8(at, bt, rt)
		case as && !bs:
			LteSVU8(at[0], bt, rt)
		case !as && bs:
			LteVSU8(at, bt[0], rt)
		default:
			LteU8(at, bt, rt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()

		switch {
		case as && bs:
			LteU16(at, bt, rt)
		case as && !bs:
			LteSVU16(at[0], bt, rt)
		case !as && bs:
			LteVSU16(at, bt[0], rt)
		default:
			LteU16(at, bt, rt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()

		switch {
		case as && bs:
			LteU32(at, bt, rt)
		case as && !bs:
			LteSVU32(at[0], bt, rt)
		case !as && bs:
			LteVSU32(at, bt[0], rt)
		default:
			LteU32(at, bt, rt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()

		switch {
		case as && bs:
			LteU64(at, bt, rt)
		case as && !bs:
			LteSVU64(at[0], bt, rt)
		case !as && bs:
			LteVSU64(at, bt[0], rt)
		default:
			LteU64(at, bt, rt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()

		switch {
		case as && bs:
			LteF32(at, bt, rt)
		case as && !bs:
			LteSVF32(at[0], bt, rt)
		case !as && bs:
			LteVSF32(at, bt[0], rt)
		default:
			LteF32(at, bt, rt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()

		switch {
		case as && bs:
			LteF64(at, bt, rt)
		case as && !bs:
			LteSVF64(at[0], bt, rt)
		case !as && bs:
			LteVSF64(at, bt[0], rt)
		default:
			LteF64(at, bt, rt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()

		switch {
		case as && bs:
			LteStr(at, bt, rt)
		case as && !bs:
			LteSVStr(at[0], bt, rt)
		case !as && bs:
			LteVSStr(at, bt[0], rt)
		default:
			LteStr(at, bt, rt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Lte", t)
	}
}

func (e E) Eq(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is a scalar. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()

		switch {
		case as && bs:
			EqB(at, bt, rt)
		case as && !bs:
			EqSVB(at[0], bt, rt)
		case !as && bs:
			EqVSB(at, bt[0], rt)
		default:
			EqB(at, bt, rt)
		}
		return
	case Int:
		at := a.Ints()
		bt := b.Ints()

		switch {
		case as && bs:
			EqI(at, bt, rt)
		case as && !bs:
			EqSVI(at[0], bt, rt)
		case !as && bs:
			EqVSI(at, bt[0], rt)
		default:
			EqI(at, bt, rt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()

		switch {
		case as && bs:
			EqI8(at, bt, rt)
		case as && !bs:
			EqSVI8(at[0], bt, rt)
		case !as && bs:
			EqVSI8(at, bt[0], rt)
		default:
			EqI8(at, bt, rt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()

		switch {
		case as && bs:
			EqI16(at, bt, rt)
		case as && !bs:
			EqSVI16(at[0], bt, rt)
		case !as && bs:
			EqVSI16(at, bt[0], rt)
		default:
			EqI16(at, bt, rt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()

		switch {
		case as && bs:
			EqI32(at, bt, rt)
		case as && !bs:
			EqSVI32(at[0], bt, rt)
		case !as && bs:
			EqVSI32(at, bt[0], rt)
		default:
			EqI32(at, bt, rt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()

		switch {
		case as && bs:
			EqI64(at, bt, rt)
		case as && !bs:
			EqSVI64(at[0], bt, rt)
		case !as && bs:
			EqVSI64(at, bt[0], rt)
		default:
			EqI64(at, bt, rt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()

		switch {
		case as && bs:
			EqU(at, bt, rt)
		case as && !bs:
			EqSVU(at[0], bt, rt)
		case !as && bs:
			EqVSU(at, bt[0], rt)
		default:
			EqU(at, bt, rt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()

		switch {
		case as && bs:
			EqU8(at, bt, rt)
		case as && !bs:
			EqSVU8(at[0], bt, rt)
		case !as && bs:
			EqVSU8(at, bt[0], rt)
		default:
			EqU8(at, bt, rt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()

		switch {
		case as && bs:
			EqU16(at, bt, rt)
		case as && !bs:
			EqSVU16(at[0], bt, rt)
		case !as && bs:
			EqVSU16(at, bt[0], rt)
		default:
			EqU16(at, bt, rt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()

		switch {
		case as && bs:
			EqU32(at, bt, rt)
		case as && !bs:
			EqSVU32(at[0], bt, rt)
		case !as && bs:
			EqVSU32(at, bt[0], rt)
		default:
			EqU32(at, bt, rt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()

		switch {
		case as && bs:
			EqU64(at, bt, rt)
		case as && !bs:
			EqSVU64(at[0], bt, rt)
		case !as && bs:
			EqVSU64(at, bt[0], rt)
		default:
			EqU64(at, bt, rt)
		}
		return
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()

		switch {
		case as && bs:
			EqUintptr(at, bt, rt)
		case as && !bs:
			EqSVUintptr(at[0], bt, rt)
		case !as && bs:
			EqVSUintptr(at, bt[0], rt)
		default:
			EqUintptr(at, bt, rt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()

		switch {
		case as && bs:
			EqF32(at, bt, rt)
		case as && !bs:
			EqSVF32(at[0], bt, rt)
		case !as && bs:
			EqVSF32(at, bt[0], rt)
		default:
			EqF32(at, bt, rt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()

		switch {
		case as && bs:
			EqF64(at, bt, rt)
		case as && !bs:
			EqSVF64(at[0], bt, rt)
		case !as && bs:
			EqVSF64(at, bt[0], rt)
		default:
			EqF64(at, bt, rt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()

		switch {
		case as && bs:
			EqC64(at, bt, rt)
		case as && !bs:
			EqSVC64(at[0], bt, rt)
		case !as && bs:
			EqVSC64(at, bt[0], rt)
		default:
			EqC64(at, bt, rt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()

		switch {
		case as && bs:
			EqC128(at, bt, rt)
		case as && !bs:
			EqSVC128(at[0], bt, rt)
		case !as && bs:
			EqVSC128(at, bt[0], rt)
		default:
			EqC128(at, bt, rt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()

		switch {
		case as && bs:
			EqStr(at, bt, rt)
		case as && !bs:
			EqSVStr(at[0], bt, rt)
		case !as && bs:
			EqVSStr(at, bt[0], rt)
		default:
			EqStr(at, bt, rt)
		}
		return
	case UnsafePointer:
		at := a.UnsafePointers()
		bt := b.UnsafePointers()

		switch {
		case as && bs:
			EqUnsafePointer(at, bt, rt)
		case as && !bs:
			EqSVUnsafePointer(at[0], bt, rt)
		case !as && bs:
			EqVSUnsafePointer(at, bt[0], rt)
		default:
			EqUnsafePointer(at, bt, rt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Eq", t)
	}
}

func (e E) Ne(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is a scalar. a: %d, b %d", a.Len(), b.Len())
	}

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()

		switch {
		case as && bs:
			NeB(at, bt, rt)
		case as && !bs:
			NeSVB(at[0], bt, rt)
		case !as && bs:
			NeVSB(at, bt[0], rt)
		default:
			NeB(at, bt, rt)
		}
		return
	case Int:
		at := a.Ints()
		bt := b.Ints()

		switch {
		case as && bs:
			NeI(at, bt, rt)
		case as && !bs:
			NeSVI(at[0], bt, rt)
		case !as && bs:
			NeVSI(at, bt[0], rt)
		default:
			NeI(at, bt, rt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()

		switch {
		case as && bs:
			NeI8(at, bt, rt)
		case as && !bs:
			NeSVI8(at[0], bt, rt)
		case !as && bs:
			NeVSI8(at, bt[0], rt)
		default:
			NeI8(at, bt, rt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()

		switch {
		case as && bs:
			NeI16(at, bt, rt)
		case as && !bs:
			NeSVI16(at[0], bt, rt)
		case !as && bs:
			NeVSI16(at, bt[0], rt)
		default:
			NeI16(at, bt, rt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()

		switch {
		case as && bs:
			NeI32(at, bt, rt)
		case as && !bs:
			NeSVI32(at[0], bt, rt)
		case !as && bs:
			NeVSI32(at, bt[0], rt)
		default:
			NeI32(at, bt, rt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()

		switch {
		case as && bs:
			NeI64(at, bt, rt)
		case as && !bs:
			NeSVI64(at[0], bt, rt)
		case !as && bs:
			NeVSI64(at, bt[0], rt)
		default:
			NeI64(at, bt, rt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()

		switch {
		case as && bs:
			NeU(at, bt, rt)
		case as && !bs:
			NeSVU(at[0], bt, rt)
		case !as && bs:
			NeVSU(at, bt[0], rt)
		default:
			NeU(at, bt, rt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()

		switch {
		case as && bs:
			NeU8(at, bt, rt)
		case as && !bs:
			NeSVU8(at[0], bt, rt)
		case !as && bs:
			NeVSU8(at, bt[0], rt)
		default:
			NeU8(at, bt, rt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()

		switch {
		case as && bs:
			NeU16(at, bt, rt)
		case as && !bs:
			NeSVU16(at[0], bt, rt)
		case !as && bs:
			NeVSU16(at, bt[0], rt)
		default:
			NeU16(at, bt, rt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()

		switch {
		case as && bs:
			NeU32(at, bt, rt)
		case as && !bs:
			NeSVU32(at[0], bt, rt)
		case !as && bs:
			NeVSU32(at, bt[0], rt)
		default:
			NeU32(at, bt, rt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()

		switch {
		case as && bs:
			NeU64(at, bt, rt)
		case as && !bs:
			NeSVU64(at[0], bt, rt)
		case !as && bs:
			NeVSU64(at, bt[0], rt)
		default:
			NeU64(at, bt, rt)
		}
		return
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()

		switch {
		case as && bs:
			NeUintptr(at, bt, rt)
		case as && !bs:
			NeSVUintptr(at[0], bt, rt)
		case !as && bs:
			NeVSUintptr(at, bt[0], rt)
		default:
			NeUintptr(at, bt, rt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()

		switch {
		case as && bs:
			NeF32(at, bt, rt)
		case as && !bs:
			NeSVF32(at[0], bt, rt)
		case !as && bs:
			NeVSF32(at, bt[0], rt)
		default:
			NeF32(at, bt, rt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()

		switch {
		case as && bs:
			NeF64(at, bt, rt)
		case as && !bs:
			NeSVF64(at[0], bt, rt)
		case !as && bs:
			NeVSF64(at, bt[0], rt)
		default:
			NeF64(at, bt, rt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()

		switch {
		case as && bs:
			NeC64(at, bt, rt)
		case as && !bs:
			NeSVC64(at[0], bt, rt)
		case !as && bs:
			NeVSC64(at, bt[0], rt)
		default:
			NeC64(at, bt, rt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()

		switch {
		case as && bs:
			NeC128(at, bt, rt)
		case as && !bs:
			NeSVC128(at[0], bt, rt)
		case !as && bs:
			NeVSC128(at, bt[0], rt)
		default:
			NeC128(at, bt, rt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()

		switch {
		case as && bs:
			NeStr(at, bt, rt)
		case as && !bs:
			NeSVStr(at[0], bt, rt)
		case !as && bs:
			NeVSStr(at, bt[0], rt)
		default:
			NeStr(at, bt, rt)
		}
		return
	case UnsafePointer:
		at := a.UnsafePointers()
		bt := b.UnsafePointers()

		switch {
		case as && bs:
			NeUnsafePointer(at, bt, rt)
		case as && !bs:
			NeSVUnsafePointer(at[0], bt, rt)
		case !as && bs:
			NeVSUnsafePointer(at, bt[0], rt)
		default:
			NeUnsafePointer(at, bt, rt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Ne", t)
	}
}

func (e E) GtSame(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			GtSameI(at, bt)
		case as && !bs:
			GtSameSVI(at[0], bt)
		case !as && bs:
			GtSameVSI(at, bt[0])
		default:
			GtSameI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			GtSameI8(at, bt)
		case as && !bs:
			GtSameSVI8(at[0], bt)
		case !as && bs:
			GtSameVSI8(at, bt[0])
		default:
			GtSameI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			GtSameI16(at, bt)
		case as && !bs:
			GtSameSVI16(at[0], bt)
		case !as && bs:
			GtSameVSI16(at, bt[0])
		default:
			GtSameI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			GtSameI32(at, bt)
		case as && !bs:
			GtSameSVI32(at[0], bt)
		case !as && bs:
			GtSameVSI32(at, bt[0])
		default:
			GtSameI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			GtSameI64(at, bt)
		case as && !bs:
			GtSameSVI64(at[0], bt)
		case !as && bs:
			GtSameVSI64(at, bt[0])
		default:
			GtSameI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			GtSameU(at, bt)
		case as && !bs:
			GtSameSVU(at[0], bt)
		case !as && bs:
			GtSameVSU(at, bt[0])
		default:
			GtSameU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			GtSameU8(at, bt)
		case as && !bs:
			GtSameSVU8(at[0], bt)
		case !as && bs:
			GtSameVSU8(at, bt[0])
		default:
			GtSameU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			GtSameU16(at, bt)
		case as && !bs:
			GtSameSVU16(at[0], bt)
		case !as && bs:
			GtSameVSU16(at, bt[0])
		default:
			GtSameU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			GtSameU32(at, bt)
		case as && !bs:
			GtSameSVU32(at[0], bt)
		case !as && bs:
			GtSameVSU32(at, bt[0])
		default:
			GtSameU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			GtSameU64(at, bt)
		case as && !bs:
			GtSameSVU64(at[0], bt)
		case !as && bs:
			GtSameVSU64(at, bt[0])
		default:
			GtSameU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			GtSameF32(at, bt)
		case as && !bs:
			GtSameSVF32(at[0], bt)
		case !as && bs:
			GtSameVSF32(at, bt[0])
		default:
			GtSameF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			GtSameF64(at, bt)
		case as && !bs:
			GtSameSVF64(at[0], bt)
		case !as && bs:
			GtSameVSF64(at, bt[0])
		default:
			GtSameF64(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			GtSameStr(at, bt)
		case as && !bs:
			GtSameSVStr(at[0], bt)
		case !as && bs:
			GtSameVSStr(at, bt[0])
		default:
			GtSameStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Gt", t)
	}
}

func (e E) GteSame(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			GteSameI(at, bt)
		case as && !bs:
			GteSameSVI(at[0], bt)
		case !as && bs:
			GteSameVSI(at, bt[0])
		default:
			GteSameI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			GteSameI8(at, bt)
		case as && !bs:
			GteSameSVI8(at[0], bt)
		case !as && bs:
			GteSameVSI8(at, bt[0])
		default:
			GteSameI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			GteSameI16(at, bt)
		case as && !bs:
			GteSameSVI16(at[0], bt)
		case !as && bs:
			GteSameVSI16(at, bt[0])
		default:
			GteSameI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			GteSameI32(at, bt)
		case as && !bs:
			GteSameSVI32(at[0], bt)
		case !as && bs:
			GteSameVSI32(at, bt[0])
		default:
			GteSameI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			GteSameI64(at, bt)
		case as && !bs:
			GteSameSVI64(at[0], bt)
		case !as && bs:
			GteSameVSI64(at, bt[0])
		default:
			GteSameI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			GteSameU(at, bt)
		case as && !bs:
			GteSameSVU(at[0], bt)
		case !as && bs:
			GteSameVSU(at, bt[0])
		default:
			GteSameU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			GteSameU8(at, bt)
		case as && !bs:
			GteSameSVU8(at[0], bt)
		case !as && bs:
			GteSameVSU8(at, bt[0])
		default:
			GteSameU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			GteSameU16(at, bt)
		case as && !bs:
			GteSameSVU16(at[0], bt)
		case !as && bs:
			GteSameVSU16(at, bt[0])
		default:
			GteSameU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			GteSameU32(at, bt)
		case as && !bs:
			GteSameSVU32(at[0], bt)
		case !as && bs:
			GteSameVSU32(at, bt[0])
		default:
			GteSameU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			GteSameU64(at, bt)
		case as && !bs:
			GteSameSVU64(at[0], bt)
		case !as && bs:
			GteSameVSU64(at, bt[0])
		default:
			GteSameU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			GteSameF32(at, bt)
		case as && !bs:
			GteSameSVF32(at[0], bt)
		case !as && bs:
			GteSameVSF32(at, bt[0])
		default:
			GteSameF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			GteSameF64(at, bt)
		case as && !bs:
			GteSameSVF64(at[0], bt)
		case !as && bs:
			GteSameVSF64(at, bt[0])
		default:
			GteSameF64(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			GteSameStr(at, bt)
		case as && !bs:
			GteSameSVStr(at[0], bt)
		case !as && bs:
			GteSameVSStr(at, bt[0])
		default:
			GteSameStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Gte", t)
	}
}

func (e E) LtSame(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			LtSameI(at, bt)
		case as && !bs:
			LtSameSVI(at[0], bt)
		case !as && bs:
			LtSameVSI(at, bt[0])
		default:
			LtSameI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			LtSameI8(at, bt)
		case as && !bs:
			LtSameSVI8(at[0], bt)
		case !as && bs:
			LtSameVSI8(at, bt[0])
		default:
			LtSameI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			LtSameI16(at, bt)
		case as && !bs:
			LtSameSVI16(at[0], bt)
		case !as && bs:
			LtSameVSI16(at, bt[0])
		default:
			LtSameI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			LtSameI32(at, bt)
		case as && !bs:
			LtSameSVI32(at[0], bt)
		case !as && bs:
			LtSameVSI32(at, bt[0])
		default:
			LtSameI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			LtSameI64(at, bt)
		case as && !bs:
			LtSameSVI64(at[0], bt)
		case !as && bs:
			LtSameVSI64(at, bt[0])
		default:
			LtSameI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			LtSameU(at, bt)
		case as && !bs:
			LtSameSVU(at[0], bt)
		case !as && bs:
			LtSameVSU(at, bt[0])
		default:
			LtSameU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			LtSameU8(at, bt)
		case as && !bs:
			LtSameSVU8(at[0], bt)
		case !as && bs:
			LtSameVSU8(at, bt[0])
		default:
			LtSameU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			LtSameU16(at, bt)
		case as && !bs:
			LtSameSVU16(at[0], bt)
		case !as && bs:
			LtSameVSU16(at, bt[0])
		default:
			LtSameU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			LtSameU32(at, bt)
		case as && !bs:
			LtSameSVU32(at[0], bt)
		case !as && bs:
			LtSameVSU32(at, bt[0])
		default:
			LtSameU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			LtSameU64(at, bt)
		case as && !bs:
			LtSameSVU64(at[0], bt)
		case !as && bs:
			LtSameVSU64(at, bt[0])
		default:
			LtSameU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			LtSameF32(at, bt)
		case as && !bs:
			LtSameSVF32(at[0], bt)
		case !as && bs:
			LtSameVSF32(at, bt[0])
		default:
			LtSameF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			LtSameF64(at, bt)
		case as && !bs:
			LtSameSVF64(at[0], bt)
		case !as && bs:
			LtSameVSF64(at, bt[0])
		default:
			LtSameF64(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			LtSameStr(at, bt)
		case as && !bs:
			LtSameSVStr(at[0], bt)
		case !as && bs:
			LtSameVSStr(at, bt[0])
		default:
			LtSameStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Lt", t)
	}
}

func (e E) LteSame(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			LteSameI(at, bt)
		case as && !bs:
			LteSameSVI(at[0], bt)
		case !as && bs:
			LteSameVSI(at, bt[0])
		default:
			LteSameI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			LteSameI8(at, bt)
		case as && !bs:
			LteSameSVI8(at[0], bt)
		case !as && bs:
			LteSameVSI8(at, bt[0])
		default:
			LteSameI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			LteSameI16(at, bt)
		case as && !bs:
			LteSameSVI16(at[0], bt)
		case !as && bs:
			LteSameVSI16(at, bt[0])
		default:
			LteSameI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			LteSameI32(at, bt)
		case as && !bs:
			LteSameSVI32(at[0], bt)
		case !as && bs:
			LteSameVSI32(at, bt[0])
		default:
			LteSameI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			LteSameI64(at, bt)
		case as && !bs:
			LteSameSVI64(at[0], bt)
		case !as && bs:
			LteSameVSI64(at, bt[0])
		default:
			LteSameI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			LteSameU(at, bt)
		case as && !bs:
			LteSameSVU(at[0], bt)
		case !as && bs:
			LteSameVSU(at, bt[0])
		default:
			LteSameU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			LteSameU8(at, bt)
		case as && !bs:
			LteSameSVU8(at[0], bt)
		case !as && bs:
			LteSameVSU8(at, bt[0])
		default:
			LteSameU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			LteSameU16(at, bt)
		case as && !bs:
			LteSameSVU16(at[0], bt)
		case !as && bs:
			LteSameVSU16(at, bt[0])
		default:
			LteSameU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			LteSameU32(at, bt)
		case as && !bs:
			LteSameSVU32(at[0], bt)
		case !as && bs:
			LteSameVSU32(at, bt[0])
		default:
			LteSameU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			LteSameU64(at, bt)
		case as && !bs:
			LteSameSVU64(at[0], bt)
		case !as && bs:
			LteSameVSU64(at, bt[0])
		default:
			LteSameU64(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			LteSameF32(at, bt)
		case as && !bs:
			LteSameSVF32(at[0], bt)
		case !as && bs:
			LteSameVSF32(at, bt[0])
		default:
			LteSameF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			LteSameF64(at, bt)
		case as && !bs:
			LteSameSVF64(at[0], bt)
		case !as && bs:
			LteSameVSF64(at, bt[0])
		default:
			LteSameF64(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			LteSameStr(at, bt)
		case as && !bs:
			LteSameSVStr(at[0], bt)
		case !as && bs:
			LteSameVSStr(at, bt[0])
		default:
			LteSameStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Lte", t)
	}
}

func (e E) EqSame(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()
		switch {
		case as && bs:
			EqSameB(at, bt)
		case as && !bs:
			EqSameSVB(at[0], bt)
		case !as && bs:
			EqSameVSB(at, bt[0])
		default:
			EqSameB(at, bt)
		}
		return
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			EqSameI(at, bt)
		case as && !bs:
			EqSameSVI(at[0], bt)
		case !as && bs:
			EqSameVSI(at, bt[0])
		default:
			EqSameI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			EqSameI8(at, bt)
		case as && !bs:
			EqSameSVI8(at[0], bt)
		case !as && bs:
			EqSameVSI8(at, bt[0])
		default:
			EqSameI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			EqSameI16(at, bt)
		case as && !bs:
			EqSameSVI16(at[0], bt)
		case !as && bs:
			EqSameVSI16(at, bt[0])
		default:
			EqSameI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			EqSameI32(at, bt)
		case as && !bs:
			EqSameSVI32(at[0], bt)
		case !as && bs:
			EqSameVSI32(at, bt[0])
		default:
			EqSameI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			EqSameI64(at, bt)
		case as && !bs:
			EqSameSVI64(at[0], bt)
		case !as && bs:
			EqSameVSI64(at, bt[0])
		default:
			EqSameI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			EqSameU(at, bt)
		case as && !bs:
			EqSameSVU(at[0], bt)
		case !as && bs:
			EqSameVSU(at, bt[0])
		default:
			EqSameU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			EqSameU8(at, bt)
		case as && !bs:
			EqSameSVU8(at[0], bt)
		case !as && bs:
			EqSameVSU8(at, bt[0])
		default:
			EqSameU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			EqSameU16(at, bt)
		case as && !bs:
			EqSameSVU16(at[0], bt)
		case !as && bs:
			EqSameVSU16(at, bt[0])
		default:
			EqSameU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			EqSameU32(at, bt)
		case as && !bs:
			EqSameSVU32(at[0], bt)
		case !as && bs:
			EqSameVSU32(at, bt[0])
		default:
			EqSameU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			EqSameU64(at, bt)
		case as && !bs:
			EqSameSVU64(at[0], bt)
		case !as && bs:
			EqSameVSU64(at, bt[0])
		default:
			EqSameU64(at, bt)
		}
		return
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()
		switch {
		case as && bs:
			EqSameUintptr(at, bt)
		case as && !bs:
			EqSameSVUintptr(at[0], bt)
		case !as && bs:
			EqSameVSUintptr(at, bt[0])
		default:
			EqSameUintptr(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			EqSameF32(at, bt)
		case as && !bs:
			EqSameSVF32(at[0], bt)
		case !as && bs:
			EqSameVSF32(at, bt[0])
		default:
			EqSameF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			EqSameF64(at, bt)
		case as && !bs:
			EqSameSVF64(at[0], bt)
		case !as && bs:
			EqSameVSF64(at, bt[0])
		default:
			EqSameF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			EqSameC64(at, bt)
		case as && !bs:
			EqSameSVC64(at[0], bt)
		case !as && bs:
			EqSameVSC64(at, bt[0])
		default:
			EqSameC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			EqSameC128(at, bt)
		case as && !bs:
			EqSameSVC128(at[0], bt)
		case !as && bs:
			EqSameVSC128(at, bt[0])
		default:
			EqSameC128(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			EqSameStr(at, bt)
		case as && !bs:
			EqSameSVStr(at[0], bt)
		case !as && bs:
			EqSameVSStr(at, bt[0])
		default:
			EqSameStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Eq", t)
	}
}

func (e E) NeSame(t reflect.Type, a *storage.Header, b *storage.Header) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()
		switch {
		case as && bs:
			NeSameB(at, bt)
		case as && !bs:
			NeSameSVB(at[0], bt)
		case !as && bs:
			NeSameVSB(at, bt[0])
		default:
			NeSameB(at, bt)
		}
		return
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			NeSameI(at, bt)
		case as && !bs:
			NeSameSVI(at[0], bt)
		case !as && bs:
			NeSameVSI(at, bt[0])
		default:
			NeSameI(at, bt)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			NeSameI8(at, bt)
		case as && !bs:
			NeSameSVI8(at[0], bt)
		case !as && bs:
			NeSameVSI8(at, bt[0])
		default:
			NeSameI8(at, bt)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			NeSameI16(at, bt)
		case as && !bs:
			NeSameSVI16(at[0], bt)
		case !as && bs:
			NeSameVSI16(at, bt[0])
		default:
			NeSameI16(at, bt)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			NeSameI32(at, bt)
		case as && !bs:
			NeSameSVI32(at[0], bt)
		case !as && bs:
			NeSameVSI32(at, bt[0])
		default:
			NeSameI32(at, bt)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			NeSameI64(at, bt)
		case as && !bs:
			NeSameSVI64(at[0], bt)
		case !as && bs:
			NeSameVSI64(at, bt[0])
		default:
			NeSameI64(at, bt)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			NeSameU(at, bt)
		case as && !bs:
			NeSameSVU(at[0], bt)
		case !as && bs:
			NeSameVSU(at, bt[0])
		default:
			NeSameU(at, bt)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			NeSameU8(at, bt)
		case as && !bs:
			NeSameSVU8(at[0], bt)
		case !as && bs:
			NeSameVSU8(at, bt[0])
		default:
			NeSameU8(at, bt)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			NeSameU16(at, bt)
		case as && !bs:
			NeSameSVU16(at[0], bt)
		case !as && bs:
			NeSameVSU16(at, bt[0])
		default:
			NeSameU16(at, bt)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			NeSameU32(at, bt)
		case as && !bs:
			NeSameSVU32(at[0], bt)
		case !as && bs:
			NeSameVSU32(at, bt[0])
		default:
			NeSameU32(at, bt)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			NeSameU64(at, bt)
		case as && !bs:
			NeSameSVU64(at[0], bt)
		case !as && bs:
			NeSameVSU64(at, bt[0])
		default:
			NeSameU64(at, bt)
		}
		return
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()
		switch {
		case as && bs:
			NeSameUintptr(at, bt)
		case as && !bs:
			NeSameSVUintptr(at[0], bt)
		case !as && bs:
			NeSameVSUintptr(at, bt[0])
		default:
			NeSameUintptr(at, bt)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			NeSameF32(at, bt)
		case as && !bs:
			NeSameSVF32(at[0], bt)
		case !as && bs:
			NeSameVSF32(at, bt[0])
		default:
			NeSameF32(at, bt)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			NeSameF64(at, bt)
		case as && !bs:
			NeSameSVF64(at[0], bt)
		case !as && bs:
			NeSameVSF64(at, bt[0])
		default:
			NeSameF64(at, bt)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			NeSameC64(at, bt)
		case as && !bs:
			NeSameSVC64(at[0], bt)
		case !as && bs:
			NeSameVSC64(at, bt[0])
		default:
			NeSameC64(at, bt)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			NeSameC128(at, bt)
		case as && !bs:
			NeSameSVC128(at[0], bt)
		case !as && bs:
			NeSameVSC128(at, bt[0])
		default:
			NeSameC128(at, bt)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			NeSameStr(at, bt)
		case as && !bs:
			NeSameSVStr(at[0], bt)
		case !as && bs:
			NeSameVSStr(at, bt[0])
		default:
			NeSameStr(at, bt)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Ne", t)
	}
}

func (e E) GtIter(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header, ait Iterator, bit Iterator, rit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is scalar while len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			GtI(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVI(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSI(at, bt[0], rt, ait, rit)
		default:
			return GtIterI(at, bt, rt, ait, bit, rit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			GtI8(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVI8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSI8(at, bt[0], rt, ait, rit)
		default:
			return GtIterI8(at, bt, rt, ait, bit, rit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			GtI16(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVI16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSI16(at, bt[0], rt, ait, rit)
		default:
			return GtIterI16(at, bt, rt, ait, bit, rit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			GtI32(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVI32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSI32(at, bt[0], rt, ait, rit)
		default:
			return GtIterI32(at, bt, rt, ait, bit, rit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			GtI64(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVI64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSI64(at, bt[0], rt, ait, rit)
		default:
			return GtIterI64(at, bt, rt, ait, bit, rit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			GtU(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVU(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSU(at, bt[0], rt, ait, rit)
		default:
			return GtIterU(at, bt, rt, ait, bit, rit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			GtU8(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVU8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSU8(at, bt[0], rt, ait, rit)
		default:
			return GtIterU8(at, bt, rt, ait, bit, rit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			GtU16(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVU16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSU16(at, bt[0], rt, ait, rit)
		default:
			return GtIterU16(at, bt, rt, ait, bit, rit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			GtU32(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVU32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSU32(at, bt[0], rt, ait, rit)
		default:
			return GtIterU32(at, bt, rt, ait, bit, rit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			GtU64(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVU64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSU64(at, bt[0], rt, ait, rit)
		default:
			return GtIterU64(at, bt, rt, ait, bit, rit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			GtF32(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVF32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSF32(at, bt[0], rt, ait, rit)
		default:
			return GtIterF32(at, bt, rt, ait, bit, rit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			GtF64(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVF64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSF64(at, bt[0], rt, ait, rit)
		default:
			return GtIterF64(at, bt, rt, ait, bit, rit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			GtStr(at, bt, rt)
			return
		case as && !bs:
			return GtIterSVStr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GtIterVSStr(at, bt[0], rt, ait, rit)
		default:
			return GtIterStr(at, bt, rt, ait, bit, rit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Gt", t)
	}
}

func (e E) GteIter(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header, ait Iterator, bit Iterator, rit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is scalar while len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			GteI(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVI(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSI(at, bt[0], rt, ait, rit)
		default:
			return GteIterI(at, bt, rt, ait, bit, rit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			GteI8(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVI8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSI8(at, bt[0], rt, ait, rit)
		default:
			return GteIterI8(at, bt, rt, ait, bit, rit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			GteI16(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVI16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSI16(at, bt[0], rt, ait, rit)
		default:
			return GteIterI16(at, bt, rt, ait, bit, rit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			GteI32(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVI32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSI32(at, bt[0], rt, ait, rit)
		default:
			return GteIterI32(at, bt, rt, ait, bit, rit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			GteI64(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVI64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSI64(at, bt[0], rt, ait, rit)
		default:
			return GteIterI64(at, bt, rt, ait, bit, rit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			GteU(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVU(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSU(at, bt[0], rt, ait, rit)
		default:
			return GteIterU(at, bt, rt, ait, bit, rit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			GteU8(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVU8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSU8(at, bt[0], rt, ait, rit)
		default:
			return GteIterU8(at, bt, rt, ait, bit, rit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			GteU16(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVU16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSU16(at, bt[0], rt, ait, rit)
		default:
			return GteIterU16(at, bt, rt, ait, bit, rit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			GteU32(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVU32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSU32(at, bt[0], rt, ait, rit)
		default:
			return GteIterU32(at, bt, rt, ait, bit, rit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			GteU64(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVU64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSU64(at, bt[0], rt, ait, rit)
		default:
			return GteIterU64(at, bt, rt, ait, bit, rit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			GteF32(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVF32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSF32(at, bt[0], rt, ait, rit)
		default:
			return GteIterF32(at, bt, rt, ait, bit, rit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			GteF64(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVF64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSF64(at, bt[0], rt, ait, rit)
		default:
			return GteIterF64(at, bt, rt, ait, bit, rit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			GteStr(at, bt, rt)
			return
		case as && !bs:
			return GteIterSVStr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return GteIterVSStr(at, bt[0], rt, ait, rit)
		default:
			return GteIterStr(at, bt, rt, ait, bit, rit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Gte", t)
	}
}

func (e E) LtIter(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header, ait Iterator, bit Iterator, rit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is scalar while len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			LtI(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVI(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSI(at, bt[0], rt, ait, rit)
		default:
			return LtIterI(at, bt, rt, ait, bit, rit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			LtI8(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVI8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSI8(at, bt[0], rt, ait, rit)
		default:
			return LtIterI8(at, bt, rt, ait, bit, rit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			LtI16(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVI16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSI16(at, bt[0], rt, ait, rit)
		default:
			return LtIterI16(at, bt, rt, ait, bit, rit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			LtI32(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVI32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSI32(at, bt[0], rt, ait, rit)
		default:
			return LtIterI32(at, bt, rt, ait, bit, rit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			LtI64(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVI64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSI64(at, bt[0], rt, ait, rit)
		default:
			return LtIterI64(at, bt, rt, ait, bit, rit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			LtU(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVU(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSU(at, bt[0], rt, ait, rit)
		default:
			return LtIterU(at, bt, rt, ait, bit, rit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			LtU8(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVU8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSU8(at, bt[0], rt, ait, rit)
		default:
			return LtIterU8(at, bt, rt, ait, bit, rit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			LtU16(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVU16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSU16(at, bt[0], rt, ait, rit)
		default:
			return LtIterU16(at, bt, rt, ait, bit, rit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			LtU32(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVU32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSU32(at, bt[0], rt, ait, rit)
		default:
			return LtIterU32(at, bt, rt, ait, bit, rit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			LtU64(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVU64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSU64(at, bt[0], rt, ait, rit)
		default:
			return LtIterU64(at, bt, rt, ait, bit, rit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			LtF32(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVF32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSF32(at, bt[0], rt, ait, rit)
		default:
			return LtIterF32(at, bt, rt, ait, bit, rit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			LtF64(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVF64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSF64(at, bt[0], rt, ait, rit)
		default:
			return LtIterF64(at, bt, rt, ait, bit, rit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			LtStr(at, bt, rt)
			return
		case as && !bs:
			return LtIterSVStr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LtIterVSStr(at, bt[0], rt, ait, rit)
		default:
			return LtIterStr(at, bt, rt, ait, bit, rit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Lt", t)
	}
}

func (e E) LteIter(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header, ait Iterator, bit Iterator, rit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is scalar while len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			LteI(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVI(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSI(at, bt[0], rt, ait, rit)
		default:
			return LteIterI(at, bt, rt, ait, bit, rit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			LteI8(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVI8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSI8(at, bt[0], rt, ait, rit)
		default:
			return LteIterI8(at, bt, rt, ait, bit, rit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			LteI16(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVI16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSI16(at, bt[0], rt, ait, rit)
		default:
			return LteIterI16(at, bt, rt, ait, bit, rit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			LteI32(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVI32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSI32(at, bt[0], rt, ait, rit)
		default:
			return LteIterI32(at, bt, rt, ait, bit, rit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			LteI64(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVI64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSI64(at, bt[0], rt, ait, rit)
		default:
			return LteIterI64(at, bt, rt, ait, bit, rit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			LteU(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVU(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSU(at, bt[0], rt, ait, rit)
		default:
			return LteIterU(at, bt, rt, ait, bit, rit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			LteU8(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVU8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSU8(at, bt[0], rt, ait, rit)
		default:
			return LteIterU8(at, bt, rt, ait, bit, rit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			LteU16(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVU16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSU16(at, bt[0], rt, ait, rit)
		default:
			return LteIterU16(at, bt, rt, ait, bit, rit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			LteU32(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVU32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSU32(at, bt[0], rt, ait, rit)
		default:
			return LteIterU32(at, bt, rt, ait, bit, rit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			LteU64(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVU64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSU64(at, bt[0], rt, ait, rit)
		default:
			return LteIterU64(at, bt, rt, ait, bit, rit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			LteF32(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVF32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSF32(at, bt[0], rt, ait, rit)
		default:
			return LteIterF32(at, bt, rt, ait, bit, rit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			LteF64(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVF64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSF64(at, bt[0], rt, ait, rit)
		default:
			return LteIterF64(at, bt, rt, ait, bit, rit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			LteStr(at, bt, rt)
			return
		case as && !bs:
			return LteIterSVStr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return LteIterVSStr(at, bt[0], rt, ait, rit)
		default:
			return LteIterStr(at, bt, rt, ait, bit, rit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Lte", t)
	}
}

func (e E) EqIter(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header, ait Iterator, bit Iterator, rit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is scalar while len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()
		switch {
		case as && bs:
			EqB(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVB(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSB(at, bt[0], rt, ait, rit)
		default:
			return EqIterB(at, bt, rt, ait, bit, rit)
		}
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			EqI(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVI(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSI(at, bt[0], rt, ait, rit)
		default:
			return EqIterI(at, bt, rt, ait, bit, rit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			EqI8(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVI8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSI8(at, bt[0], rt, ait, rit)
		default:
			return EqIterI8(at, bt, rt, ait, bit, rit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			EqI16(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVI16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSI16(at, bt[0], rt, ait, rit)
		default:
			return EqIterI16(at, bt, rt, ait, bit, rit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			EqI32(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVI32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSI32(at, bt[0], rt, ait, rit)
		default:
			return EqIterI32(at, bt, rt, ait, bit, rit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			EqI64(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVI64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSI64(at, bt[0], rt, ait, rit)
		default:
			return EqIterI64(at, bt, rt, ait, bit, rit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			EqU(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVU(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSU(at, bt[0], rt, ait, rit)
		default:
			return EqIterU(at, bt, rt, ait, bit, rit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			EqU8(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVU8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSU8(at, bt[0], rt, ait, rit)
		default:
			return EqIterU8(at, bt, rt, ait, bit, rit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			EqU16(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVU16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSU16(at, bt[0], rt, ait, rit)
		default:
			return EqIterU16(at, bt, rt, ait, bit, rit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			EqU32(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVU32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSU32(at, bt[0], rt, ait, rit)
		default:
			return EqIterU32(at, bt, rt, ait, bit, rit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			EqU64(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVU64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSU64(at, bt[0], rt, ait, rit)
		default:
			return EqIterU64(at, bt, rt, ait, bit, rit)
		}
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()
		switch {
		case as && bs:
			EqUintptr(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVUintptr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSUintptr(at, bt[0], rt, ait, rit)
		default:
			return EqIterUintptr(at, bt, rt, ait, bit, rit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			EqF32(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVF32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSF32(at, bt[0], rt, ait, rit)
		default:
			return EqIterF32(at, bt, rt, ait, bit, rit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			EqF64(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVF64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSF64(at, bt[0], rt, ait, rit)
		default:
			return EqIterF64(at, bt, rt, ait, bit, rit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			EqC64(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVC64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSC64(at, bt[0], rt, ait, rit)
		default:
			return EqIterC64(at, bt, rt, ait, bit, rit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			EqC128(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVC128(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSC128(at, bt[0], rt, ait, rit)
		default:
			return EqIterC128(at, bt, rt, ait, bit, rit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			EqStr(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVStr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSStr(at, bt[0], rt, ait, rit)
		default:
			return EqIterStr(at, bt, rt, ait, bit, rit)
		}
	case UnsafePointer:
		at := a.UnsafePointers()
		bt := b.UnsafePointers()
		switch {
		case as && bs:
			EqUnsafePointer(at, bt, rt)
			return
		case as && !bs:
			return EqIterSVUnsafePointer(at[0], bt, rt, bit, rit)
		case !as && bs:
			return EqIterVSUnsafePointer(at, bt[0], rt, ait, rit)
		default:
			return EqIterUnsafePointer(at, bt, rt, ait, bit, rit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Eq", t)
	}
}

func (e E) NeIter(t reflect.Type, a *storage.Header, b *storage.Header, retVal *storage.Header, ait Iterator, bit Iterator, rit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)
	rs := isScalar(retVal)
	rt := retVal.Bools()

	if (as && !bs) || (bs && !as) && rs {
		return errors.Errorf("retVal is scalar while len(a): %d, len(b) %d", a.Len(), b.Len())
	}

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()
		switch {
		case as && bs:
			NeB(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVB(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSB(at, bt[0], rt, ait, rit)
		default:
			return NeIterB(at, bt, rt, ait, bit, rit)
		}
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			NeI(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVI(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSI(at, bt[0], rt, ait, rit)
		default:
			return NeIterI(at, bt, rt, ait, bit, rit)
		}
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			NeI8(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVI8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSI8(at, bt[0], rt, ait, rit)
		default:
			return NeIterI8(at, bt, rt, ait, bit, rit)
		}
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			NeI16(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVI16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSI16(at, bt[0], rt, ait, rit)
		default:
			return NeIterI16(at, bt, rt, ait, bit, rit)
		}
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			NeI32(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVI32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSI32(at, bt[0], rt, ait, rit)
		default:
			return NeIterI32(at, bt, rt, ait, bit, rit)
		}
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			NeI64(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVI64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSI64(at, bt[0], rt, ait, rit)
		default:
			return NeIterI64(at, bt, rt, ait, bit, rit)
		}
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			NeU(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVU(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSU(at, bt[0], rt, ait, rit)
		default:
			return NeIterU(at, bt, rt, ait, bit, rit)
		}
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			NeU8(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVU8(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSU8(at, bt[0], rt, ait, rit)
		default:
			return NeIterU8(at, bt, rt, ait, bit, rit)
		}
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			NeU16(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVU16(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSU16(at, bt[0], rt, ait, rit)
		default:
			return NeIterU16(at, bt, rt, ait, bit, rit)
		}
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			NeU32(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVU32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSU32(at, bt[0], rt, ait, rit)
		default:
			return NeIterU32(at, bt, rt, ait, bit, rit)
		}
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			NeU64(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVU64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSU64(at, bt[0], rt, ait, rit)
		default:
			return NeIterU64(at, bt, rt, ait, bit, rit)
		}
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()
		switch {
		case as && bs:
			NeUintptr(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVUintptr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSUintptr(at, bt[0], rt, ait, rit)
		default:
			return NeIterUintptr(at, bt, rt, ait, bit, rit)
		}
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			NeF32(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVF32(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSF32(at, bt[0], rt, ait, rit)
		default:
			return NeIterF32(at, bt, rt, ait, bit, rit)
		}
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			NeF64(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVF64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSF64(at, bt[0], rt, ait, rit)
		default:
			return NeIterF64(at, bt, rt, ait, bit, rit)
		}
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			NeC64(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVC64(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSC64(at, bt[0], rt, ait, rit)
		default:
			return NeIterC64(at, bt, rt, ait, bit, rit)
		}
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			NeC128(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVC128(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSC128(at, bt[0], rt, ait, rit)
		default:
			return NeIterC128(at, bt, rt, ait, bit, rit)
		}
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			NeStr(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVStr(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSStr(at, bt[0], rt, ait, rit)
		default:
			return NeIterStr(at, bt, rt, ait, bit, rit)
		}
	case UnsafePointer:
		at := a.UnsafePointers()
		bt := b.UnsafePointers()
		switch {
		case as && bs:
			NeUnsafePointer(at, bt, rt)
			return
		case as && !bs:
			return NeIterSVUnsafePointer(at[0], bt, rt, bit, rit)
		case !as && bs:
			return NeIterVSUnsafePointer(at, bt[0], rt, ait, rit)
		default:
			return NeIterUnsafePointer(at, bt, rt, ait, bit, rit)
		}
	default:
		return errors.Errorf("Unsupported type %v for Ne", t)
	}
}

func (e E) GtSameIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			GtSameI(at, bt)
		case as && !bs:
			GtSameIterSVI(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSI(at, bt[0], ait)
		default:
			GtSameIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			GtSameI8(at, bt)
		case as && !bs:
			GtSameIterSVI8(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSI8(at, bt[0], ait)
		default:
			GtSameIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			GtSameI16(at, bt)
		case as && !bs:
			GtSameIterSVI16(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSI16(at, bt[0], ait)
		default:
			GtSameIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			GtSameI32(at, bt)
		case as && !bs:
			GtSameIterSVI32(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSI32(at, bt[0], ait)
		default:
			GtSameIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			GtSameI64(at, bt)
		case as && !bs:
			GtSameIterSVI64(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSI64(at, bt[0], ait)
		default:
			GtSameIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			GtSameU(at, bt)
		case as && !bs:
			GtSameIterSVU(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSU(at, bt[0], ait)
		default:
			GtSameIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			GtSameU8(at, bt)
		case as && !bs:
			GtSameIterSVU8(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSU8(at, bt[0], ait)
		default:
			GtSameIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			GtSameU16(at, bt)
		case as && !bs:
			GtSameIterSVU16(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSU16(at, bt[0], ait)
		default:
			GtSameIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			GtSameU32(at, bt)
		case as && !bs:
			GtSameIterSVU32(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSU32(at, bt[0], ait)
		default:
			GtSameIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			GtSameU64(at, bt)
		case as && !bs:
			GtSameIterSVU64(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSU64(at, bt[0], ait)
		default:
			GtSameIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			GtSameF32(at, bt)
		case as && !bs:
			GtSameIterSVF32(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSF32(at, bt[0], ait)
		default:
			GtSameIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			GtSameF64(at, bt)
		case as && !bs:
			GtSameIterSVF64(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSF64(at, bt[0], ait)
		default:
			GtSameIterF64(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			GtSameStr(at, bt)
		case as && !bs:
			GtSameIterSVStr(at[0], bt, bit)
		case !as && bs:
			GtSameIterVSStr(at, bt[0], ait)
		default:
			GtSameIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Gt", t)
	}
}

func (e E) GteSameIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			GteSameI(at, bt)
		case as && !bs:
			GteSameIterSVI(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSI(at, bt[0], ait)
		default:
			GteSameIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			GteSameI8(at, bt)
		case as && !bs:
			GteSameIterSVI8(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSI8(at, bt[0], ait)
		default:
			GteSameIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			GteSameI16(at, bt)
		case as && !bs:
			GteSameIterSVI16(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSI16(at, bt[0], ait)
		default:
			GteSameIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			GteSameI32(at, bt)
		case as && !bs:
			GteSameIterSVI32(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSI32(at, bt[0], ait)
		default:
			GteSameIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			GteSameI64(at, bt)
		case as && !bs:
			GteSameIterSVI64(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSI64(at, bt[0], ait)
		default:
			GteSameIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			GteSameU(at, bt)
		case as && !bs:
			GteSameIterSVU(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSU(at, bt[0], ait)
		default:
			GteSameIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			GteSameU8(at, bt)
		case as && !bs:
			GteSameIterSVU8(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSU8(at, bt[0], ait)
		default:
			GteSameIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			GteSameU16(at, bt)
		case as && !bs:
			GteSameIterSVU16(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSU16(at, bt[0], ait)
		default:
			GteSameIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			GteSameU32(at, bt)
		case as && !bs:
			GteSameIterSVU32(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSU32(at, bt[0], ait)
		default:
			GteSameIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			GteSameU64(at, bt)
		case as && !bs:
			GteSameIterSVU64(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSU64(at, bt[0], ait)
		default:
			GteSameIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			GteSameF32(at, bt)
		case as && !bs:
			GteSameIterSVF32(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSF32(at, bt[0], ait)
		default:
			GteSameIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			GteSameF64(at, bt)
		case as && !bs:
			GteSameIterSVF64(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSF64(at, bt[0], ait)
		default:
			GteSameIterF64(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			GteSameStr(at, bt)
		case as && !bs:
			GteSameIterSVStr(at[0], bt, bit)
		case !as && bs:
			GteSameIterVSStr(at, bt[0], ait)
		default:
			GteSameIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Gte", t)
	}
}

func (e E) LtSameIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			LtSameI(at, bt)
		case as && !bs:
			LtSameIterSVI(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSI(at, bt[0], ait)
		default:
			LtSameIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			LtSameI8(at, bt)
		case as && !bs:
			LtSameIterSVI8(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSI8(at, bt[0], ait)
		default:
			LtSameIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			LtSameI16(at, bt)
		case as && !bs:
			LtSameIterSVI16(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSI16(at, bt[0], ait)
		default:
			LtSameIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			LtSameI32(at, bt)
		case as && !bs:
			LtSameIterSVI32(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSI32(at, bt[0], ait)
		default:
			LtSameIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			LtSameI64(at, bt)
		case as && !bs:
			LtSameIterSVI64(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSI64(at, bt[0], ait)
		default:
			LtSameIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			LtSameU(at, bt)
		case as && !bs:
			LtSameIterSVU(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSU(at, bt[0], ait)
		default:
			LtSameIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			LtSameU8(at, bt)
		case as && !bs:
			LtSameIterSVU8(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSU8(at, bt[0], ait)
		default:
			LtSameIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			LtSameU16(at, bt)
		case as && !bs:
			LtSameIterSVU16(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSU16(at, bt[0], ait)
		default:
			LtSameIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			LtSameU32(at, bt)
		case as && !bs:
			LtSameIterSVU32(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSU32(at, bt[0], ait)
		default:
			LtSameIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			LtSameU64(at, bt)
		case as && !bs:
			LtSameIterSVU64(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSU64(at, bt[0], ait)
		default:
			LtSameIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			LtSameF32(at, bt)
		case as && !bs:
			LtSameIterSVF32(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSF32(at, bt[0], ait)
		default:
			LtSameIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			LtSameF64(at, bt)
		case as && !bs:
			LtSameIterSVF64(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSF64(at, bt[0], ait)
		default:
			LtSameIterF64(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			LtSameStr(at, bt)
		case as && !bs:
			LtSameIterSVStr(at[0], bt, bit)
		case !as && bs:
			LtSameIterVSStr(at, bt[0], ait)
		default:
			LtSameIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Lt", t)
	}
}

func (e E) LteSameIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			LteSameI(at, bt)
		case as && !bs:
			LteSameIterSVI(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSI(at, bt[0], ait)
		default:
			LteSameIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			LteSameI8(at, bt)
		case as && !bs:
			LteSameIterSVI8(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSI8(at, bt[0], ait)
		default:
			LteSameIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			LteSameI16(at, bt)
		case as && !bs:
			LteSameIterSVI16(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSI16(at, bt[0], ait)
		default:
			LteSameIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			LteSameI32(at, bt)
		case as && !bs:
			LteSameIterSVI32(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSI32(at, bt[0], ait)
		default:
			LteSameIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			LteSameI64(at, bt)
		case as && !bs:
			LteSameIterSVI64(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSI64(at, bt[0], ait)
		default:
			LteSameIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			LteSameU(at, bt)
		case as && !bs:
			LteSameIterSVU(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSU(at, bt[0], ait)
		default:
			LteSameIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			LteSameU8(at, bt)
		case as && !bs:
			LteSameIterSVU8(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSU8(at, bt[0], ait)
		default:
			LteSameIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			LteSameU16(at, bt)
		case as && !bs:
			LteSameIterSVU16(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSU16(at, bt[0], ait)
		default:
			LteSameIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			LteSameU32(at, bt)
		case as && !bs:
			LteSameIterSVU32(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSU32(at, bt[0], ait)
		default:
			LteSameIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			LteSameU64(at, bt)
		case as && !bs:
			LteSameIterSVU64(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSU64(at, bt[0], ait)
		default:
			LteSameIterU64(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			LteSameF32(at, bt)
		case as && !bs:
			LteSameIterSVF32(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSF32(at, bt[0], ait)
		default:
			LteSameIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			LteSameF64(at, bt)
		case as && !bs:
			LteSameIterSVF64(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSF64(at, bt[0], ait)
		default:
			LteSameIterF64(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			LteSameStr(at, bt)
		case as && !bs:
			LteSameIterSVStr(at[0], bt, bit)
		case !as && bs:
			LteSameIterVSStr(at, bt[0], ait)
		default:
			LteSameIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Lte", t)
	}
}

func (e E) EqSameIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()
		switch {
		case as && bs:
			EqSameB(at, bt)
		case as && !bs:
			EqSameIterSVB(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSB(at, bt[0], ait)
		default:
			EqSameIterB(at, bt, ait, bit)
		}
		return
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			EqSameI(at, bt)
		case as && !bs:
			EqSameIterSVI(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSI(at, bt[0], ait)
		default:
			EqSameIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			EqSameI8(at, bt)
		case as && !bs:
			EqSameIterSVI8(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSI8(at, bt[0], ait)
		default:
			EqSameIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			EqSameI16(at, bt)
		case as && !bs:
			EqSameIterSVI16(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSI16(at, bt[0], ait)
		default:
			EqSameIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			EqSameI32(at, bt)
		case as && !bs:
			EqSameIterSVI32(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSI32(at, bt[0], ait)
		default:
			EqSameIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			EqSameI64(at, bt)
		case as && !bs:
			EqSameIterSVI64(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSI64(at, bt[0], ait)
		default:
			EqSameIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			EqSameU(at, bt)
		case as && !bs:
			EqSameIterSVU(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSU(at, bt[0], ait)
		default:
			EqSameIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			EqSameU8(at, bt)
		case as && !bs:
			EqSameIterSVU8(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSU8(at, bt[0], ait)
		default:
			EqSameIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			EqSameU16(at, bt)
		case as && !bs:
			EqSameIterSVU16(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSU16(at, bt[0], ait)
		default:
			EqSameIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			EqSameU32(at, bt)
		case as && !bs:
			EqSameIterSVU32(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSU32(at, bt[0], ait)
		default:
			EqSameIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			EqSameU64(at, bt)
		case as && !bs:
			EqSameIterSVU64(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSU64(at, bt[0], ait)
		default:
			EqSameIterU64(at, bt, ait, bit)
		}
		return
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()
		switch {
		case as && bs:
			EqSameUintptr(at, bt)
		case as && !bs:
			EqSameIterSVUintptr(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSUintptr(at, bt[0], ait)
		default:
			EqSameIterUintptr(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			EqSameF32(at, bt)
		case as && !bs:
			EqSameIterSVF32(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSF32(at, bt[0], ait)
		default:
			EqSameIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			EqSameF64(at, bt)
		case as && !bs:
			EqSameIterSVF64(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSF64(at, bt[0], ait)
		default:
			EqSameIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			EqSameC64(at, bt)
		case as && !bs:
			EqSameIterSVC64(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSC64(at, bt[0], ait)
		default:
			EqSameIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			EqSameC128(at, bt)
		case as && !bs:
			EqSameIterSVC128(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSC128(at, bt[0], ait)
		default:
			EqSameIterC128(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			EqSameStr(at, bt)
		case as && !bs:
			EqSameIterSVStr(at[0], bt, bit)
		case !as && bs:
			EqSameIterVSStr(at, bt[0], ait)
		default:
			EqSameIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Eq", t)
	}
}

func (e E) NeSameIter(t reflect.Type, a *storage.Header, b *storage.Header, ait Iterator, bit Iterator) (err error) {
	as := isScalar(a)
	bs := isScalar(b)

	switch t {
	case Bool:
		at := a.Bools()
		bt := b.Bools()
		switch {
		case as && bs:
			NeSameB(at, bt)
		case as && !bs:
			NeSameIterSVB(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSB(at, bt[0], ait)
		default:
			NeSameIterB(at, bt, ait, bit)
		}
		return
	case Int:
		at := a.Ints()
		bt := b.Ints()
		switch {
		case as && bs:
			NeSameI(at, bt)
		case as && !bs:
			NeSameIterSVI(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSI(at, bt[0], ait)
		default:
			NeSameIterI(at, bt, ait, bit)
		}
		return
	case Int8:
		at := a.Int8s()
		bt := b.Int8s()
		switch {
		case as && bs:
			NeSameI8(at, bt)
		case as && !bs:
			NeSameIterSVI8(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSI8(at, bt[0], ait)
		default:
			NeSameIterI8(at, bt, ait, bit)
		}
		return
	case Int16:
		at := a.Int16s()
		bt := b.Int16s()
		switch {
		case as && bs:
			NeSameI16(at, bt)
		case as && !bs:
			NeSameIterSVI16(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSI16(at, bt[0], ait)
		default:
			NeSameIterI16(at, bt, ait, bit)
		}
		return
	case Int32:
		at := a.Int32s()
		bt := b.Int32s()
		switch {
		case as && bs:
			NeSameI32(at, bt)
		case as && !bs:
			NeSameIterSVI32(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSI32(at, bt[0], ait)
		default:
			NeSameIterI32(at, bt, ait, bit)
		}
		return
	case Int64:
		at := a.Int64s()
		bt := b.Int64s()
		switch {
		case as && bs:
			NeSameI64(at, bt)
		case as && !bs:
			NeSameIterSVI64(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSI64(at, bt[0], ait)
		default:
			NeSameIterI64(at, bt, ait, bit)
		}
		return
	case Uint:
		at := a.Uints()
		bt := b.Uints()
		switch {
		case as && bs:
			NeSameU(at, bt)
		case as && !bs:
			NeSameIterSVU(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSU(at, bt[0], ait)
		default:
			NeSameIterU(at, bt, ait, bit)
		}
		return
	case Uint8:
		at := a.Uint8s()
		bt := b.Uint8s()
		switch {
		case as && bs:
			NeSameU8(at, bt)
		case as && !bs:
			NeSameIterSVU8(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSU8(at, bt[0], ait)
		default:
			NeSameIterU8(at, bt, ait, bit)
		}
		return
	case Uint16:
		at := a.Uint16s()
		bt := b.Uint16s()
		switch {
		case as && bs:
			NeSameU16(at, bt)
		case as && !bs:
			NeSameIterSVU16(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSU16(at, bt[0], ait)
		default:
			NeSameIterU16(at, bt, ait, bit)
		}
		return
	case Uint32:
		at := a.Uint32s()
		bt := b.Uint32s()
		switch {
		case as && bs:
			NeSameU32(at, bt)
		case as && !bs:
			NeSameIterSVU32(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSU32(at, bt[0], ait)
		default:
			NeSameIterU32(at, bt, ait, bit)
		}
		return
	case Uint64:
		at := a.Uint64s()
		bt := b.Uint64s()
		switch {
		case as && bs:
			NeSameU64(at, bt)
		case as && !bs:
			NeSameIterSVU64(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSU64(at, bt[0], ait)
		default:
			NeSameIterU64(at, bt, ait, bit)
		}
		return
	case Uintptr:
		at := a.Uintptrs()
		bt := b.Uintptrs()
		switch {
		case as && bs:
			NeSameUintptr(at, bt)
		case as && !bs:
			NeSameIterSVUintptr(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSUintptr(at, bt[0], ait)
		default:
			NeSameIterUintptr(at, bt, ait, bit)
		}
		return
	case Float32:
		at := a.Float32s()
		bt := b.Float32s()
		switch {
		case as && bs:
			NeSameF32(at, bt)
		case as && !bs:
			NeSameIterSVF32(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSF32(at, bt[0], ait)
		default:
			NeSameIterF32(at, bt, ait, bit)
		}
		return
	case Float64:
		at := a.Float64s()
		bt := b.Float64s()
		switch {
		case as && bs:
			NeSameF64(at, bt)
		case as && !bs:
			NeSameIterSVF64(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSF64(at, bt[0], ait)
		default:
			NeSameIterF64(at, bt, ait, bit)
		}
		return
	case Complex64:
		at := a.Complex64s()
		bt := b.Complex64s()
		switch {
		case as && bs:
			NeSameC64(at, bt)
		case as && !bs:
			NeSameIterSVC64(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSC64(at, bt[0], ait)
		default:
			NeSameIterC64(at, bt, ait, bit)
		}
		return
	case Complex128:
		at := a.Complex128s()
		bt := b.Complex128s()
		switch {
		case as && bs:
			NeSameC128(at, bt)
		case as && !bs:
			NeSameIterSVC128(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSC128(at, bt[0], ait)
		default:
			NeSameIterC128(at, bt, ait, bit)
		}
		return
	case String:
		at := a.Strings()
		bt := b.Strings()
		switch {
		case as && bs:
			NeSameStr(at, bt)
		case as && !bs:
			NeSameIterSVStr(at[0], bt, bit)
		case !as && bs:
			NeSameIterVSStr(at, bt[0], ait)
		default:
			NeSameIterStr(at, bt, ait, bit)
		}
		return
	default:
		return errors.Errorf("Unsupported type %v for Ne", t)
	}
}