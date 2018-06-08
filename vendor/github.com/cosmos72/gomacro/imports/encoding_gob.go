// this file was generated by gomacro command: import _b "encoding/gob"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"encoding/gob"
)

// reflection: allow interpreted code to import "encoding/gob"
func init() {
	Packages["encoding/gob"] = Package{
	Binds: map[string]Value{
		"NewDecoder":	ValueOf(gob.NewDecoder),
		"NewEncoder":	ValueOf(gob.NewEncoder),
		"Register":	ValueOf(gob.Register),
		"RegisterName":	ValueOf(gob.RegisterName),
	}, Types: map[string]Type{
		"CommonType":	TypeOf((*gob.CommonType)(nil)).Elem(),
		"Decoder":	TypeOf((*gob.Decoder)(nil)).Elem(),
		"Encoder":	TypeOf((*gob.Encoder)(nil)).Elem(),
		"GobDecoder":	TypeOf((*gob.GobDecoder)(nil)).Elem(),
		"GobEncoder":	TypeOf((*gob.GobEncoder)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"GobDecoder":	TypeOf((*P_encoding_gob_GobDecoder)(nil)).Elem(),
		"GobEncoder":	TypeOf((*P_encoding_gob_GobEncoder)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for encoding/gob.GobDecoder ---------------
type P_encoding_gob_GobDecoder struct {
	Object	interface{}
	GobDecode_	func(interface{}, []byte) error
}
func (P *P_encoding_gob_GobDecoder) GobDecode(unnamed0 []byte) error {
	return P.GobDecode_(P.Object, unnamed0)
}

// --------------- proxy for encoding/gob.GobEncoder ---------------
type P_encoding_gob_GobEncoder struct {
	Object	interface{}
	GobEncode_	func(interface{}) ([]byte, error)
}
func (P *P_encoding_gob_GobEncoder) GobEncode() ([]byte, error) {
	return P.GobEncode_(P.Object)
}
