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

package elements ;import (_g "encoding/xml";_a "fmt";_d "github.com/unidoc/unioffice";);func NewElementsGroup ()*ElementsGroup {_fd :=&ElementsGroup {};return _fd };func (_dgbd *SimpleLiteral )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewElementContainer ()*ElementContainer {_da :=&ElementContainer {};return _da };func (_baeb *ElementsGroupChoice )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _baeb .Any !=nil {_be :=_g .StartElement {Name :_g .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};for _ ,_bc :=range _baeb .Any {e .EncodeElement (_bc ,_be );};};return nil ;};type SimpleLiteral struct{};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_cd *ElementContainer )ValidateWithPath (path string )error {for _ab ,_df :=range _cd .Choice {if _gb :=_df .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_ab ));_gb !=nil {return _gb ;};};return nil ;};

// Validate validates the ElementsGroupChoice and its children
func (_agg *ElementsGroupChoice )Validate ()error {return _agg .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};func NewAny ()*Any {_ce :=&Any {};_ce .SimpleLiteral =*NewSimpleLiteral ();return _ce };func (_b *Any )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {return _b .SimpleLiteral .MarshalXML (e ,start );};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_bg *SimpleLiteral )ValidateWithPath (path string )error {return nil };type ElementsGroup struct{Choice []*ElementsGroupChoice ;};func NewElementsGroupChoice ()*ElementsGroupChoice {_fb :=&ElementsGroupChoice {};return _fb };type Any struct{SimpleLiteral };func NewSimpleLiteral ()*SimpleLiteral {_ca :=&SimpleLiteral {};return _ca };type ElementsGroupChoice struct{Any []*Any ;};func (_ed *ElementsGroup )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_ec :for {_eaa ,_ega :=d .Token ();if _ega !=nil {return _ega ;};switch _ef :=_eaa .(type ){case _g .StartElement :switch _ef .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_ag :=NewElementsGroupChoice ();if _fg :=d .DecodeElement (&_ag .Any ,&_ef );_fg !=nil {return _fg ;};_ed .Choice =append (_ed .Choice ,_ag );default:_d .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_ef .Name );if _gcf :=d .Skip ();_gcf !=nil {return _gcf ;};};case _g .EndElement :break _ec ;case _g .CharData :};};return nil ;};

// Validate validates the SimpleLiteral and its children
func (_adb *SimpleLiteral )Validate ()error {return _adb .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};func (_f *Any )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_f .SimpleLiteral =*NewSimpleLiteral ();for {_ad ,_ba :=d .Token ();if _ba !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_ba );};if _e ,_cc :=_ad .(_g .EndElement );_cc &&_e .Name ==start .Name {break ;};};return nil ;};func (_de *SimpleLiteral )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for {_dfe ,_ac :=d .Token ();if _ac !=nil {return _a .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_ac );};if _gfb ,_ffb :=_dfe .(_g .EndElement );_ffb &&_gfb .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_cb *Any )ValidateWithPath (path string )error {if _ae :=_cb .SimpleLiteral .ValidateWithPath (path );_ae !=nil {return _ae ;};return nil ;};

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_egae *ElementsGroup )ValidateWithPath (path string )error {for _fc ,_fge :=range _egae .Choice {if _dfc :=_fge .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_fc ));_dfc !=nil {return _dfc ;};};return nil ;};func (_cbd *ElementContainer )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _cbd .Choice !=nil {for _ ,_ea :=range _cbd .Choice {_ea .MarshalXML (e ,_g .StartElement {});};};e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_beb *ElementsGroupChoice )ValidateWithPath (path string )error {for _bca ,_dc :=range _beb .Any {if _ga :=_dc .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_bca ));_ga !=nil {return _ga ;};};return nil ;};type ElementContainer struct{Choice []*ElementsGroupChoice ;};

// Validate validates the Any and its children
func (_cg *Any )Validate ()error {return _cg .ValidateWithPath ("\u0041\u006e\u0079")};

// Validate validates the ElementsGroup and its children
func (_ecg *ElementsGroup )Validate ()error {return _ecg .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};func (_aa *ElementContainer )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_eag :for {_eg ,_fe :=d .Token ();if _fe !=nil {return _fe ;};switch _gc :=_eg .(type ){case _g .StartElement :switch _gc .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_gf :=NewElementsGroupChoice ();if _bf :=d .DecodeElement (&_gf .Any ,&_gc );_bf !=nil {return _bf ;};_aa .Choice =append (_aa .Choice ,_gf );default:_d .Log ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_gc .Name );if _bae :=d .Skip ();_bae !=nil {return _bae ;};};case _g .EndElement :break _eag ;case _g .CharData :};};return nil ;};

// Validate validates the ElementContainer and its children
func (_ge *ElementContainer )Validate ()error {return _ge .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};func (_ee *ElementsGroup )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _ee .Choice !=nil {for _ ,_af :=range _ee .Choice {_af .MarshalXML (e ,_g .StartElement {});};};return nil ;};func (_age *ElementsGroupChoice )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_dg :for {_dgb ,_gea :=d .Token ();if _gea !=nil {return _gea ;};switch _cbf :=_dgb .(type ){case _g .StartElement :switch _cbf .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cda :=NewAny ();if _ccc :=d .DecodeElement (_cda ,&_cbf );_ccc !=nil {return _ccc ;};_age .Any =append (_age .Any ,_cda );default:_d .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_cbf .Name );if _fdb :=d .Skip ();_fdb !=nil {return _fdb ;};};case _g .EndElement :break _dg ;case _g .CharData :};};return nil ;};func init (){_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );};