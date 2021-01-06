//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore ;import (_f "encoding/hex";_e "errors";_c "fmt";_cb "github.com/unidoc/unioffice/common/tempstorage";_a "io";_be "io/ioutil";_ae "math/rand";_g "sync";);func _cff (_bd int )(string ,error ){_ggd :=make ([]byte ,_bd );if _ ,_ga :=_ae .Read (_ggd );_ga !=nil {return "",_ga ;};return _f .EncodeToString (_ggd ),nil ;};type memFile struct{_d *memDataCell ;_cg int64 ;};

// Open returns tempstorage File object by name
func (_dg *memStorage )Open (path string )(_cb .File ,error ){_af ,_gfe :=_dg ._egg .Load (path );if !_gfe {return nil ,_e .New (_c .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_d :_af .(*memDataCell )},nil ;};

// TempDir creates a name for a new temp directory using a pattern argument
func (_ba *memStorage )TempDir (pattern string )(string ,error ){return _fb (pattern ),nil };

// Add reads a file from a disk and adds it to the storage
func (_de *memStorage )Add (path string )error {_ ,_gfc :=_de ._egg .Load (path );if _gfc {return nil ;};_gg ,_ab :=_be .ReadFile (path );if _ab !=nil {return _ab ;};_de ._egg .Store (path ,&memDataCell {_gc :path ,_ec :_gg ,_ac :int64 (len (_gg ))});return nil ;};

// TempFile creates a new empty file in the storage and returns it
func (_bb *memStorage )TempFile (dir ,pattern string )(_cb .File ,error ){_cf :=dir +"\u002f"+_fb (pattern );_ce :=&memDataCell {_gc :_cf ,_ec :[]byte {}};_gcd :=&memFile {_d :_ce };_bb ._egg .Store (_cf ,_ce );return _gcd ,nil ;};

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_gf *memFile )Read (p []byte )(int ,error ){_fg :=_gf ._cg ;_gb :=_gf ._d ._ac ;_ad :=int64 (len (p ));if _ad > _gb {_ad =_gb ;p =p [:_ad ];};if _fg >=_gb {return 0,_a .EOF ;};_fa :=_fg +_ad ;if _fa >=_gb {_fa =_gb ;};_ee :=copy (p ,_gf ._d ._ec [_fg :_fa ]);_gf ._cg =_fa ;return _ee ,nil ;};

// RemoveAll removes all files according to the dir argument prefix
func (_aea *memStorage )RemoveAll (dir string )error {_aea ._egg .Range (func (_dc ,_cbf interface{})bool {_aea ._egg .Delete (_dc );return true });return nil ;};type memDataCell struct{_gc string ;_ec []byte ;_ac int64 ;};func _fb (_fbe string )string {_eef ,_ :=_cff (6);return _fbe +_eef };

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_eg *memFile )Write (p []byte )(int ,error ){_eg ._d ._ec =append (_eg ._d ._ec ,p ...);_eg ._d ._ac +=int64 (len (p ));return len (p ),nil ;};

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_acg :=memStorage {_egg :_g .Map {}};_cb .SetAsStorage (&_acg )};

// Close is not applicable in this implementation
func (_fe *memFile )Close ()error {return nil };type memStorage struct{_egg _g .Map };

// Name returns the filename of the underlying memDataCell
func (_bc *memFile )Name ()string {return _bc ._d ._gc };