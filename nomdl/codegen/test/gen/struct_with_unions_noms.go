// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_with_unions_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("StructWithUnions",
			[]types.Field{
				types.Field{"a", types.MakeTypeRef(ref.Ref{}, 1), false},
				types.Field{"d", types.MakeTypeRef(ref.Ref{}, 2), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("",
			[]types.Field{},
			types.Choices{
				types.Field{"b", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"c", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
		),
		types.MakeStructTypeRef("",
			[]types.Field{},
			types.Choices{
				types.Field{"e", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"f", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_with_unions_CachedRef = types.RegisterPackage(&p)
}

// StructWithUnions

type StructWithUnions struct {
	m   types.Map
	ref *ref.Ref
}

func NewStructWithUnions() StructWithUnions {
	return StructWithUnions{types.NewMap(
		types.NewString("a"), New__unionOfBOfFloat64AndCOfString(),
		types.NewString("d"), New__unionOfEOfFloat64AndFOfString(),
	), &ref.Ref{}}
}

type StructWithUnionsDef struct {
	A __unionOfBOfFloat64AndCOfStringDef
	D __unionOfEOfFloat64AndFOfStringDef
}

func (def StructWithUnionsDef) New() StructWithUnions {
	return StructWithUnions{
		types.NewMap(
			types.NewString("a"), def.A.New(),
			types.NewString("d"), def.D.New(),
		), &ref.Ref{}}
}

func (s StructWithUnions) Def() (d StructWithUnionsDef) {
	d.A = s.m.Get(types.NewString("a")).(__unionOfBOfFloat64AndCOfString).Def()
	d.D = s.m.Get(types.NewString("d")).(__unionOfEOfFloat64AndFOfString).Def()
	return
}

var __typeRefForStructWithUnions types.TypeRef

func (m StructWithUnions) TypeRef() types.TypeRef {
	return __typeRefForStructWithUnions
}

func init() {
	__typeRefForStructWithUnions = types.MakeTypeRef(__genPackageInFile_struct_with_unions_CachedRef, 0)
	types.RegisterFromValFunction(__typeRefForStructWithUnions, func(v types.Value) types.Value {
		return StructWithUnionsFromVal(v)
	})
}

func StructWithUnionsFromVal(val types.Value) StructWithUnions {
	// TODO: Do we still need FromVal?
	if val, ok := val.(StructWithUnions); ok {
		return val
	}
	// TODO: Validate here
	return StructWithUnions{val.(types.Map), &ref.Ref{}}
}

func (s StructWithUnions) InternalImplementation() types.Map {
	return s.m
}

func (s StructWithUnions) Equals(other types.Value) bool {
	if other, ok := other.(StructWithUnions); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s StructWithUnions) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructWithUnions) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s StructWithUnions) A() __unionOfBOfFloat64AndCOfString {
	return s.m.Get(types.NewString("a")).(__unionOfBOfFloat64AndCOfString)
}

func (s StructWithUnions) SetA(val __unionOfBOfFloat64AndCOfString) StructWithUnions {
	return StructWithUnions{s.m.Set(types.NewString("a"), val), &ref.Ref{}}
}

func (s StructWithUnions) D() __unionOfEOfFloat64AndFOfString {
	return s.m.Get(types.NewString("d")).(__unionOfEOfFloat64AndFOfString)
}

func (s StructWithUnions) SetD(val __unionOfEOfFloat64AndFOfString) StructWithUnions {
	return StructWithUnions{s.m.Set(types.NewString("d"), val), &ref.Ref{}}
}

// __unionOfBOfFloat64AndCOfString

type __unionOfBOfFloat64AndCOfString struct {
	m   types.Map
	ref *ref.Ref
}

func New__unionOfBOfFloat64AndCOfString() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{types.NewMap(
		types.NewString("$unionIndex"), types.UInt32(0),
		types.NewString("$unionValue"), types.Float64(0),
	), &ref.Ref{}}
}

type __unionOfBOfFloat64AndCOfStringDef struct {
	__unionIndex uint32
	__unionValue interface{}
}

func (def __unionOfBOfFloat64AndCOfStringDef) New() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{
		types.NewMap(
			types.NewString("$unionIndex"), types.UInt32(def.__unionIndex),
			types.NewString("$unionValue"), def.__unionDefToValue(),
		), &ref.Ref{}}
}

func (s __unionOfBOfFloat64AndCOfString) Def() (d __unionOfBOfFloat64AndCOfStringDef) {
	d.__unionIndex = uint32(s.m.Get(types.NewString("$unionIndex")).(types.UInt32))
	d.__unionValue = s.__unionValueToDef()
	return
}

func (def __unionOfBOfFloat64AndCOfStringDef) __unionDefToValue() types.Value {
	switch def.__unionIndex {
	case 0:
		return types.Float64(def.__unionValue.(float64))
	case 1:
		return types.NewString(def.__unionValue.(string))
	}
	panic("unreachable")
}

func (s __unionOfBOfFloat64AndCOfString) __unionValueToDef() interface{} {
	switch uint32(s.m.Get(types.NewString("$unionIndex")).(types.UInt32)) {
	case 0:
		return float64(s.m.Get(types.NewString("$unionValue")).(types.Float64))
	case 1:
		return s.m.Get(types.NewString("$unionValue")).(types.String).String()
	}
	panic("unreachable")
}

var __typeRefFor__unionOfBOfFloat64AndCOfString types.TypeRef

func (m __unionOfBOfFloat64AndCOfString) TypeRef() types.TypeRef {
	return __typeRefFor__unionOfBOfFloat64AndCOfString
}

func init() {
	__typeRefFor__unionOfBOfFloat64AndCOfString = types.MakeTypeRef(__genPackageInFile_struct_with_unions_CachedRef, 1)
	types.RegisterFromValFunction(__typeRefFor__unionOfBOfFloat64AndCOfString, func(v types.Value) types.Value {
		return __unionOfBOfFloat64AndCOfStringFromVal(v)
	})
}

func __unionOfBOfFloat64AndCOfStringFromVal(val types.Value) __unionOfBOfFloat64AndCOfString {
	// TODO: Do we still need FromVal?
	if val, ok := val.(__unionOfBOfFloat64AndCOfString); ok {
		return val
	}
	// TODO: Validate here
	return __unionOfBOfFloat64AndCOfString{val.(types.Map), &ref.Ref{}}
}

func (s __unionOfBOfFloat64AndCOfString) InternalImplementation() types.Map {
	return s.m
}

func (s __unionOfBOfFloat64AndCOfString) Equals(other types.Value) bool {
	if other, ok := other.(__unionOfBOfFloat64AndCOfString); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s __unionOfBOfFloat64AndCOfString) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s __unionOfBOfFloat64AndCOfString) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s __unionOfBOfFloat64AndCOfString) B() (val float64, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 0 {
		return
	}
	return float64(s.m.Get(types.NewString("$unionValue")).(types.Float64)), true
}

func (s __unionOfBOfFloat64AndCOfString) SetB(val float64) __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{s.m.Set(types.NewString("$unionIndex"), types.UInt32(0)).Set(types.NewString("$unionValue"), types.Float64(val)), &ref.Ref{}}
}

func (def __unionOfBOfFloat64AndCOfStringDef) B() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(float64), true
}

func (def __unionOfBOfFloat64AndCOfStringDef) SetB(val float64) __unionOfBOfFloat64AndCOfStringDef {
	def.__unionIndex = 0
	def.__unionValue = val
	return def
}

func (s __unionOfBOfFloat64AndCOfString) C() (val string, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 1 {
		return
	}
	return s.m.Get(types.NewString("$unionValue")).(types.String).String(), true
}

func (s __unionOfBOfFloat64AndCOfString) SetC(val string) __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{s.m.Set(types.NewString("$unionIndex"), types.UInt32(1)).Set(types.NewString("$unionValue"), types.NewString(val)), &ref.Ref{}}
}

func (def __unionOfBOfFloat64AndCOfStringDef) C() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(string), true
}

func (def __unionOfBOfFloat64AndCOfStringDef) SetC(val string) __unionOfBOfFloat64AndCOfStringDef {
	def.__unionIndex = 1
	def.__unionValue = val
	return def
}

// __unionOfEOfFloat64AndFOfString

type __unionOfEOfFloat64AndFOfString struct {
	m   types.Map
	ref *ref.Ref
}

func New__unionOfEOfFloat64AndFOfString() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{types.NewMap(
		types.NewString("$unionIndex"), types.UInt32(0),
		types.NewString("$unionValue"), types.Float64(0),
	), &ref.Ref{}}
}

type __unionOfEOfFloat64AndFOfStringDef struct {
	__unionIndex uint32
	__unionValue interface{}
}

func (def __unionOfEOfFloat64AndFOfStringDef) New() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{
		types.NewMap(
			types.NewString("$unionIndex"), types.UInt32(def.__unionIndex),
			types.NewString("$unionValue"), def.__unionDefToValue(),
		), &ref.Ref{}}
}

func (s __unionOfEOfFloat64AndFOfString) Def() (d __unionOfEOfFloat64AndFOfStringDef) {
	d.__unionIndex = uint32(s.m.Get(types.NewString("$unionIndex")).(types.UInt32))
	d.__unionValue = s.__unionValueToDef()
	return
}

func (def __unionOfEOfFloat64AndFOfStringDef) __unionDefToValue() types.Value {
	switch def.__unionIndex {
	case 0:
		return types.Float64(def.__unionValue.(float64))
	case 1:
		return types.NewString(def.__unionValue.(string))
	}
	panic("unreachable")
}

func (s __unionOfEOfFloat64AndFOfString) __unionValueToDef() interface{} {
	switch uint32(s.m.Get(types.NewString("$unionIndex")).(types.UInt32)) {
	case 0:
		return float64(s.m.Get(types.NewString("$unionValue")).(types.Float64))
	case 1:
		return s.m.Get(types.NewString("$unionValue")).(types.String).String()
	}
	panic("unreachable")
}

var __typeRefFor__unionOfEOfFloat64AndFOfString types.TypeRef

func (m __unionOfEOfFloat64AndFOfString) TypeRef() types.TypeRef {
	return __typeRefFor__unionOfEOfFloat64AndFOfString
}

func init() {
	__typeRefFor__unionOfEOfFloat64AndFOfString = types.MakeTypeRef(__genPackageInFile_struct_with_unions_CachedRef, 2)
	types.RegisterFromValFunction(__typeRefFor__unionOfEOfFloat64AndFOfString, func(v types.Value) types.Value {
		return __unionOfEOfFloat64AndFOfStringFromVal(v)
	})
}

func __unionOfEOfFloat64AndFOfStringFromVal(val types.Value) __unionOfEOfFloat64AndFOfString {
	// TODO: Do we still need FromVal?
	if val, ok := val.(__unionOfEOfFloat64AndFOfString); ok {
		return val
	}
	// TODO: Validate here
	return __unionOfEOfFloat64AndFOfString{val.(types.Map), &ref.Ref{}}
}

func (s __unionOfEOfFloat64AndFOfString) InternalImplementation() types.Map {
	return s.m
}

func (s __unionOfEOfFloat64AndFOfString) Equals(other types.Value) bool {
	if other, ok := other.(__unionOfEOfFloat64AndFOfString); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s __unionOfEOfFloat64AndFOfString) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s __unionOfEOfFloat64AndFOfString) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s __unionOfEOfFloat64AndFOfString) E() (val float64, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 0 {
		return
	}
	return float64(s.m.Get(types.NewString("$unionValue")).(types.Float64)), true
}

func (s __unionOfEOfFloat64AndFOfString) SetE(val float64) __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{s.m.Set(types.NewString("$unionIndex"), types.UInt32(0)).Set(types.NewString("$unionValue"), types.Float64(val)), &ref.Ref{}}
}

func (def __unionOfEOfFloat64AndFOfStringDef) E() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(float64), true
}

func (def __unionOfEOfFloat64AndFOfStringDef) SetE(val float64) __unionOfEOfFloat64AndFOfStringDef {
	def.__unionIndex = 0
	def.__unionValue = val
	return def
}

func (s __unionOfEOfFloat64AndFOfString) F() (val string, ok bool) {
	if s.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 1 {
		return
	}
	return s.m.Get(types.NewString("$unionValue")).(types.String).String(), true
}

func (s __unionOfEOfFloat64AndFOfString) SetF(val string) __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{s.m.Set(types.NewString("$unionIndex"), types.UInt32(1)).Set(types.NewString("$unionValue"), types.NewString(val)), &ref.Ref{}}
}

func (def __unionOfEOfFloat64AndFOfStringDef) F() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(string), true
}

func (def __unionOfEOfFloat64AndFOfStringDef) SetF(val string) __unionOfEOfFloat64AndFOfStringDef {
	def.__unionIndex = 1
	def.__unionValue = val
	return def
}