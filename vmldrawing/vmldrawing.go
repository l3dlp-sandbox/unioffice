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

package vmldrawing ;import (_b "encoding/xml";_e "fmt";_bf "github.com/unidoc/unioffice";_be "github.com/unidoc/unioffice/common/logger";_db "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes";_dd "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/office/excel";
_cc "github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml";_a "strconv";_c "strings";);

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape (col ,row int64 )*_cc .Shape {_gf :=_cc .NewShape ();_gf .IdAttr =_bf .String (_e .Sprintf ("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064",col ,row ));_gf .TypeAttr =_bf .String ("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032");
_gf .StyleAttr =_bf .String ("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e");
_gf .FillcolorAttr =_bf .String ("\u0023f\u0062\u0066\u0036\u0064\u0036");_gf .StrokecolorAttr =_bf .String ("\u0023e\u0064\u0065\u0061\u0061\u0031");_cd :=_cc .NewEG_ShapeElements ();_cd .Fill =_cc .NewFill ();_cd .Fill .Color2Attr =_bf .String ("\u0023f\u0062\u0066\u0065\u0038\u0032");
_cd .Fill .AngleAttr =_bf .Float64 (-180);_cd .Fill .TypeAttr =_cc .ST_FillTypeGradient ;_cd .Fill .Fill =_cc .NewOfcFill ();_cd .Fill .Fill .ExtAttr =_cc .ST_ExtView ;_cd .Fill .Fill .TypeAttr =_cc .OfcST_FillTypeGradientUnscaled ;_gf .EG_ShapeElements =append (_gf .EG_ShapeElements ,_cd );
_ad :=_cc .NewEG_ShapeElements ();_ad .Shadow =_cc .NewShadow ();_ad .Shadow .OnAttr =_db .ST_TrueFalseT ;_ad .Shadow .ObscuredAttr =_db .ST_TrueFalseT ;_gf .EG_ShapeElements =append (_gf .EG_ShapeElements ,_ad );_gg :=_cc .NewEG_ShapeElements ();_gg .Path =_cc .NewPath ();
_gg .Path .ConnecttypeAttr =_cc .OfcST_ConnectTypeNone ;_gf .EG_ShapeElements =append (_gf .EG_ShapeElements ,_gg );_af :=_cc .NewEG_ShapeElements ();_af .Textbox =_cc .NewTextbox ();_af .Textbox .StyleAttr =_bf .String ("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f");
_gf .EG_ShapeElements =append (_gf .EG_ShapeElements ,_af );_ed :=_cc .NewEG_ShapeElements ();_ed .ClientData =_dd .NewClientData ();_ed .ClientData .ObjectTypeAttr =_dd .ST_ObjectTypeNote ;_ed .ClientData .MoveWithCells =_db .ST_TrueFalseBlankT ;_ed .ClientData .SizeWithCells =_db .ST_TrueFalseBlankT ;
_ed .ClientData .Anchor =_bf .String ("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033");_ed .ClientData .AutoFill =_db .ST_TrueFalseBlankFalse ;_ed .ClientData .Row =_bf .Int64 (row );
_ed .ClientData .Column =_bf .Int64 (col );_gf .EG_ShapeElements =append (_gf .EG_ShapeElements ,_ed );return _gf ;};

// Width return width of shape.
func (_fg *ShapeStyle )Width ()float64 {return _fg ._fag };const (ShapeStylePositionAbsolute ="\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065";ShapeStylePositionRelative ="\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065";);

// ToString formatting ShapeStyle to string.
func (_dge *ShapeStyle )String ()string {_ebc :="";_ebc +=_e .Sprintf ("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0025\u0073\u003b",_dge ._gad );_ebc +=_e .Sprintf ("\u006da\u0072g\u0069\u006e\u002d\u006c\u0065\u0066\u0074\u003a\u0025\u0064\u003b",int64 (_dge ._bc ));
_ebc +=_e .Sprintf ("\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006fp\u003a\u0025\u0064\u003b",int64 (_dge ._dce ));_ebc +=_e .Sprintf ("w\u0069\u0064\u0074\u0068\u003a\u0025\u0064\u0070\u0074\u003b",int64 (_dge ._fag ));_ebc +=_e .Sprintf ("\u0068\u0065\u0069g\u0068\u0074\u003a\u0025\u0064\u0070\u0074\u003b",int64 (_dge ._ebe ));
_ebc +=_e .Sprintf ("z\u002d\u0069\u006e\u0064\u0065\u0078\u003a\u0025\u0064\u003b",_dge ._ccd );_ebc +=_e .Sprintf ("m\u0073\u006f\u002d\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069z\u006f\u006e\u0074a\u006c:\u0025\u0073\u003b",_dge ._ec );
_ebc +=_e .Sprintf ("\u006d\u0073o-\u0070\u006f\u0073i\u0074\u0069\u006f\u006e-ho\u0072iz\u006f\u006e\u0074\u0061\u006c\u002d\u0072el\u0061\u0074\u0069\u0076\u0065\u003a\u0025s\u003b",_dge ._dag );_ebc +=_e .Sprintf ("\u006ds\u006f\u002d\u0070\u006fs\u0069\u0074\u0069\u006f\u006e-\u0076e\u0072t\u0069\u0063\u0061\u006c\u003a\u0025\u0073;",_dge ._fe );
_ebc +=_e .Sprintf ("\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076e\u0072t\u0069c\u0061l\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065\u003a\u0025\u0073\u003b",_dge ._gc );return _ebc ;};

// MSOPositionVerticalRelative get `mso-position-vertical-relative` attribute of shape style.
func (_faab *ShapeStyle )MSOPositionVerticalRelative ()string {return _faab ._gc };func NewContainer ()*Container {return &Container {}};

// Height return height of shape.
func (_gef *ShapeStyle )Height ()float64 {return _gef ._ebe };

// SetFontSize sets text's fontSize.
func (_gb *TextpathStyle )SetFontSize (fontSize int64 ){_gb ._acd =fontSize };

// NewShapeStyle accept value of string style attribute in v:shape and format it to generate ShapeStyle.
func NewShapeStyle (style string )ShapeStyle {_bd :=ShapeStyle {_fag :0,_ebe :0};_bea :=_c .Split (style ,"\u003b");for _ ,_eda :=range _bea {_bfd :=_c .Split (_eda ,"\u003a");if len (_bfd )!=2{continue ;};var _ge error ;switch _bfd [0]{case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":_bd ._gad =_bfd [1];
break ;case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070":_bd ._dce ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);break ;case "m\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074":_bd ._bc ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);
break ;case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0062\u006f\u0074\u0074\u006f\u006d":_bd ._cg ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);break ;case "\u006d\u0061\u0072g\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074":_bd ._dca ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);
break ;case "\u0074\u006f\u0070":_bd ._bca ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);break ;case "\u006c\u0065\u0066\u0074":_bd ._eg ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);break ;case "\u0062\u006f\u0074\u0074\u006f\u006d":_bd ._cda ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);
break ;case "\u0072\u0069\u0067h\u0074":_bd ._ccfb ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);break ;case "\u0077\u0069\u0064t\u0068":_bd ._fag ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);break ;case "\u0068\u0065\u0069\u0067\u0068\u0074":_bd ._ebe ,_ge =_a .ParseFloat (_c .ReplaceAll (_bfd [1],"\u0070\u0074",""),64);
break ;case "\u007a-\u0069\u006e\u0064\u0065\u0078":_bd ._ccd ,_ge =_a .ParseInt (_bfd [1],10,64);break ;case "\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c":_bd ._ec =_bfd [1];
break ;case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006fr\u0069z\u006f\u006e\u0074\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":_bd ._dag =_bfd [1];break ;case "m\u0073\u006f\u002d\u0070os\u0069t\u0069\u006f\u006e\u002d\u0076e\u0072\u0074\u0069\u0063\u0061\u006c":_bd ._fe =_bfd [1];
break ;case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069o\u006e\u002d\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":_bd ._gc =_bfd [1];break ;};if _ge !=nil {_be .Log .Debug ("\u0055n\u0061\u0062l\u0065\u0020\u0074o\u0020\u0070\u0061\u0072\u0073\u0065\u0020s\u0074\u0079\u006c\u0065\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0073 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_bfd [0],_bfd [1]);
};};return _bd ;};

// IsBold returns true if text is bold.
func (_afe *TextpathStyle )IsBold ()bool {return _afe ._fb };

// SetWidth set width of shape.
func (_gd *ShapeStyle )SetWidth (width float64 ){_gd ._fag =width };

// Bottom get bottom attribute of shape style.
func (_ffg *ShapeStyle )Bottom ()float64 {return _ffg ._cda };

// Right get right attribute of shape style.
func (_cga *ShapeStyle )Right ()float64 {return _cga ._ccfb };

// ToString generate string of TextpathStyle.
func (_eaa *TextpathStyle )String ()string {_ddc :="";_ddc +=_e .Sprintf ("\u0066o\u006et\u002d\u0066\u0061\u006d\u0069\u006c\u0079\u003a\u0025\u0073\u003b",_eaa ._bdc );_ddc +=_e .Sprintf ("\u0066o\u006et\u002d\u0073\u0069\u007a\u0065\u003a\u0025\u0064\u0070\u0074\u003b",_eaa ._acd );
if _eaa ._bde {_ddc +=_e .Sprintf ("\u0066o\u006et\u002d\u0073\u0074\u0079\u006ce\u003a\u0069t\u0061\u006c\u0069\u0063\u003b");};if _eaa ._fb {_ddc +=_e .Sprintf ("\u0066\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074\u003ab\u006f\u006c\u0064\u003b");
};return _ddc ;};func (_dc *Container )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0076"},Value :"\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"});
start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u006f"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"});
start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"});
start .Name .Local ="\u0078\u006d\u006c";e .EncodeToken (start );if _dc .Layout !=nil {_ga :=_b .StartElement {Name :_b .Name {Local :"\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}};e .EncodeElement (_dc .Layout ,_ga );
};if _dc .ShapeType !=nil {_cb :=_b .StartElement {Name :_b .Name {Local :"v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}};e .EncodeElement (_dc .ShapeType ,_cb );};for _ ,_fa :=range _dc .Shape {_cf :=_b .StartElement {Name :_b .Name {Local :"\u0076:\u0073\u0068\u0061\u0070\u0065"}};
e .EncodeElement (_fa ,_cf );};return e .EncodeToken (_b .EndElement {Name :start .Name });};

// FontSize returns fontSize of the text.
func (_gda *TextpathStyle )FontSize ()int64 {return _gda ._acd };

// SetHeight set height of shape.
func (_fgb *ShapeStyle )SetHeight (height float64 ){_fgb ._ebe =height };

// ShapeStyle is style attribute of v:shape element.
type ShapeStyle struct{_gad string ;_dce float64 ;_bc float64 ;_cg float64 ;_dca float64 ;_bca float64 ;_eg float64 ;_cda float64 ;_ccfb float64 ;_fag float64 ;_ebe float64 ;_ccd int64 ;_ec string ;_dag string ;_fe string ;_gc string ;};

// TextpathStyle is style attribute of element v:textpath.
type TextpathStyle struct{_bdc string ;_acd int64 ;_fb bool ;_bde bool ;};

// SetItalic sets text to italic.
func (_ebcc *TextpathStyle )SetItalic (italic bool ){_ebcc ._bde =italic };

// FontFamily returns fontFamily of the text.
func (_egd *TextpathStyle )FontFamily ()string {return _egd ._bdc };

// SetBold sets text to bold.
func (_gdf *TextpathStyle )SetBold (bold bool ){_gdf ._fb =bold };

// SetFontFamily sets text's fontFamily.
func (_bg *TextpathStyle )SetFontFamily (fontFamily string ){_bg ._bdc =fontFamily };

// NewTextpathStyle accept value of string style attribute of element v:textpath and format it to generate TextpathStyle.
func NewTextpathStyle (style string )TextpathStyle {_eca :=TextpathStyle {_bdc :"\u0022C\u0061\u006c\u0069\u0062\u0072\u0069\"",_acd :44,_fb :false ,_bde :false };_fea :=_c .Split (style ,"\u003b");for _ ,_dgeg :=range _fea {_bdd :=_c .Split (_dgeg ,"\u003a");
if len (_bdd )!=2{continue ;};switch _bdd [0]{case "f\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069\u006c\u0079":_eca ._bdc =_bdd [1];break ;case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":_eca ._acd ,_ =_a .ParseInt (_c .ReplaceAll (_bdd [1],"\u0070\u0074",""),10,64);
break ;case "f\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074":_eca ._fb =_bdd [1]=="\u0062\u006f\u006c\u0064";break ;case "\u0066\u006f\u006e\u0074\u002d\u0073\u0074\u0079\u006c\u0065":_eca ._bde =_bdd [1]=="\u0069\u0074\u0061\u006c\u0069\u0063";
break ;};};return _eca ;};

// Position get position attribute of shape style.
func (_ef *ShapeStyle )Position ()string {return _ef ._gad };func (_ccf *Container )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_ccf .Shape =nil ;_fd :for {_eb ,_dg :=d .Token ();if _dg !=nil {return _dg ;};switch _aa :=_eb .(type ){case _b .StartElement :switch _aa .Name .Local {case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":_ccf .Layout =_cc .NewOfcShapelayout ();
if _bfc :=d .DecodeElement (_ccf .Layout ,&_aa );_bfc !=nil {return _bfc ;};case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":_ccf .ShapeType =_cc .NewShapetype ();if _ab :=d .DecodeElement (_ccf .ShapeType ,&_aa );_ab !=nil {return _ab ;};case "\u0073\u0068\u0061p\u0065":_ac :=_cc .NewShape ();
if _da :=d .DecodeElement (_ac ,&_aa );_da !=nil {return _da ;};_ccf .Shape =append (_ccf .Shape ,_ac );};case _b .EndElement :break _fd ;};};return nil ;};

// CreateFormula creates F element for typeFormulas.
func CreateFormula (s string )*_cc .CT_F {_gac :=_cc .NewCT_F ();_gac .EqnAttr =&s ;return _gac };

// Margins get margin top, left, bottom, and right of shape style.
func (_ff *ShapeStyle )Margins ()(float64 ,float64 ,float64 ,float64 ){return _ff ._dce ,_ff ._bc ,_ff ._cg ,_ff ._dca ;};

// IsItalic returns true if text is italic.
func (_bcc *TextpathStyle )IsItalic ()bool {return _bcc ._bde };type Container struct{Layout *_cc .OfcShapelayout ;ShapeType *_cc .Shapetype ;Shape []*_cc .Shape ;};

// Left get left attribute of shape style.
func (_ggd *ShapeStyle )Left ()float64 {return _ggd ._eg };

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing ()*Container {_f :=NewContainer ();_f .Layout =_cc .NewOfcShapelayout ();_f .Layout .ExtAttr =_cc .ST_ExtEdit ;_f .Layout .Idmap =_cc .NewOfcCT_IdMap ();_f .Layout .Idmap .DataAttr =_bf .String ("\u0031");_f .Layout .Idmap .ExtAttr =_cc .ST_ExtEdit ;
_f .ShapeType =_cc .NewShapetype ();_f .ShapeType .IdAttr =_bf .String ("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032");_f .ShapeType .CoordsizeAttr =_bf .String ("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030");_f .ShapeType .SptAttr =_bf .Float32 (202);
_f .ShapeType .PathAttr =_bf .String ("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065");_g :=_cc .NewEG_ShapeElements ();_f .ShapeType .EG_ShapeElements =append (_f .ShapeType .EG_ShapeElements ,_g );
_g .Path =_cc .NewPath ();_g .Path .GradientshapeokAttr =_db .ST_TrueFalseT ;_g .Path .ConnecttypeAttr =_cc .OfcST_ConnectTypeRect ;return _f ;};

// MSOPositionHorizontalRelative get `mso-position-horizontal-relative` attribute of shape style.
func (_faa *ShapeStyle )MSOPositionHorizontalRelative ()string {return _faa ._dag };

// Top get top attribute of shape style.
func (_ea *ShapeStyle )Top ()float64 {return _ea ._bca };