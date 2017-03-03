package cqlstress

import (
	"math/rand"
	"time"
)

var Rand math.Rand

func init() {
	// Todo, make rand random after testing is done
	Rand := rand.New(0)
}

func randInt64(min, max int64) int64 {
	return i.Min + Rand.Int63n(i.Max-i.Min)
}
func randInt(min, max int32) int32 {
	return i.Min + Rand.Int32n(i.Max-i.Min)
}

type DataType interface {
	Random(min interface{}, max interface{}) interface{}
	Static() interface{}
}

type CasInteger struct {
	Dist int
	Min  int
	Max  int
}

func (i CasInteger) Random() {
	if i.Dist != StaticDistribution {
		// todo: add ExpDistribution later
		return i.Min + Rand.Int63n(i.Max-i.Min)
	}
	return 0
}

type CasFloat struct {
	Dist int
	// todo add - Min, Max.   maybe std, mean instead
}

func (f CasFloat) Random() {
	if f.Dist == StaticDistribution {
		return 0
	}
	if f.Dist == NormalDistribution {
		return Rand.Float64()
	}
	return Rand.ExpFloat64()

}

type CasDateTime struct {
	Dist   int
	Min    time.Time
	iMin   int64
	Max    time.Time
	iMax   int64
	static int64
}

func (dt CasDateTime) Random() {
	if dt.Dist == StaticDistribution {
		if dt.static != 0 {
			dt.static = time.Unix(1136239445, 0)
		}
		// Mon Jan 2 15:04:05 MST 2006
		return dt.static
	}
	if dt.iMin == 0 {
		iMin = dt.Min.Unix
		iMax = dt.Max.Unix
	}
	//todo add Exp Distribution
	return time.Unix(randInt(iMin, iMax), 0)
}

type CasVarchar struct {
	Dist
	MinLength int
	MaxLength int
	static    string
}

func (cs CasVarchar) Random() string {
	if Dist == StaticDistribution {
		if static == "" {
			static = RandomString(cs.MinLength, cs.MaxLength)
		}
		return cs.static
	}
	return RandomString(cs.MinLength, cs.MaxLength)
}

func RandomString(minlen, maxlen int) string {
	len := randInt(minlen, maxlen)
	b := make([]byte, n)
	// A Rand.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, Rand.Int63(), len; i >= 0; {
		if remain == 0 {
			cache, remain = Rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
