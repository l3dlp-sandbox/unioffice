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

package zippkg ;import (_g "archive/zip";_c "bytes";_ff "encoding/xml";_a "fmt";_gbd "github.com/unidoc/unioffice";_dc "github.com/unidoc/unioffice/algo";_gb "github.com/unidoc/unioffice/common/tempstorage";_dfa "github.com/unidoc/unioffice/schema/soo/pkg/relationships";
_f "io";_df "path";_gg "sort";_dd "strings";_e "time";);

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_g .File ,path string )(string ,error ){_ddd ,_egd :=_gb .TempFile (path ,"\u007a\u007a");if _egd !=nil {return "",_egd ;};defer _ddd .Close ();_ee ,_egd :=f .Open ();if _egd !=nil {return "",_egd ;};defer _ee .Close ();_ ,_egd =_f .Copy (_ddd ,_ee );
if _egd !=nil {return "",_egd ;};return _ddd .Name (),nil ;};var _ggc =[]byte {'/','>'};

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_ae map[string ]Target ;_cc map[*_dfa .Relationships ]string ;_fgc []Target ;_ec OnNewRelationshipFunc ;_gc map[string ]struct{};_fgf map[string ]int ;};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_g .Writer ,filename string ,v interface{})error {_ceb :=&_g .FileHeader {};_ceb .Method =_g .Deflate ;_ceb .Name =filename ;_ceb .SetModTime (_e .Now ());_ge ,_ega :=z .CreateHeader (_ceb );if _ega !=nil {return _a .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_ega );
};_ ,_ega =_ge .Write ([]byte (XMLHeader ));if _ega !=nil {return _a .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_ega );
};if _ega =_ff .NewEncoder (SelfClosingWriter {_ge }).Encode (v );_ega !=nil {return _a .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_ega );};_ ,_ega =_ge .Write (_fd );return _ega ;
};func (_afd SelfClosingWriter )Write (b []byte )(int ,error ){_debc :=0;_fag :=0;for _daf :=0;_daf < len (b )-2;_daf ++{if b [_daf ]=='>'&&b [_daf +1]=='<'&&b [_daf +2]=='/'{_ccb :=[]byte {};_gd :=_daf ;for _cad :=_daf ;_cad >=0;_cad --{if b [_cad ]==' '{_gd =_cad ;
}else if b [_cad ]=='<'{_ccb =b [_cad +1:_gd ];break ;};};_fb :=[]byte {};for _dee :=_daf +3;_dee < len (b );_dee ++{if b [_dee ]=='>'{_fb =b [_daf +3:_dee ];break ;};};if !_c .Equal (_ccb ,_fb ){continue ;};_ed ,_abc :=_afd .W .Write (b [_debc :_daf ]);
if _abc !=nil {return _fag +_ed ,_abc ;};_fag +=_ed ;_ ,_abc =_afd .W .Write (_ggc );if _abc !=nil {return _fag ,_abc ;};_fag +=3;for _bga :=_daf +2;_bga < len (b )&&b [_bga ]!='>';_bga ++{_fag ++;_debc =_bga +2;_daf =_debc ;};};};_gae ,_gef :=_afd .W .Write (b [_debc :]);
return _gae +_fag ,_gef ;};

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _f .Writer ;};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_g .Writer ,zipPath ,storagePath string )error {_dg ,_efg :=z .Create (zipPath );if _efg !=nil {return _a .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_efg );
};_eag ,_efg :=_gb .Open (storagePath );if _efg !=nil {return _a .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_efg );};defer _eag .Close ();_ ,_efg =_f .Copy (_dg ,_eag );return _efg ;
};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_ga :=_dd .Split (path ,"\u002f");_ab :=_dd .Join (_ga [0:len (_ga )-1],"\u002f");_abd :=_ga [len (_ga )-1];_ab +="\u002f_\u0072\u0065\u006c\u0073\u002f";_abd +="\u002e\u0072\u0065l\u0073";return _ab +_abd ;};


// Decode loops decoding targets registered with AddTarget and calling th
func (_bc *DecodeMap )Decode (files []*_g .File )error {_db :=1;for _db > 0{for len (_bc ._fgc )> 0{_gf :=_bc ._fgc [0];_bc ._fgc =_bc ._fgc [1:];_ecf :=_gf .Ifc .(*_dfa .Relationships );for _ ,_dfg :=range _ecf .Relationship {_bd :=_bc ._cc [_ecf ];_dfg .TargetAttr =_dd .TrimPrefix (_dfg .TargetAttr ,"\u002f");
if _dd .IndexByte (_bd ,'/')> -1{_bcb :=_bd [:_dd .IndexByte (_bd ,'/')+1];if _dd .HasPrefix (_dfg .TargetAttr ,_bcb ){_bd ="";};};if _dd .HasPrefix (_dfg .TargetAttr ,_bd ){_bd ="";};_bc ._ec (_bc ,_bd +_dfg .TargetAttr ,_dfg .TypeAttr ,files ,_dfg ,_gf );
};};for _fac ,_bf :=range files {if _bf ==nil {continue ;};if _bb ,_dcf :=_bc ._ae [_bf .Name ];_dcf {delete (_bc ._ae ,_bf .Name );if _ba :=Decode (_bf ,_bb .Ifc );_ba !=nil {return _ba ;};files [_fac ]=nil ;if _bab ,_ce :=_bb .Ifc .(*_dfa .Relationships );
_ce {_bc ._fgc =append (_bc ._fgc ,_bb );_dba ,_ :=_df .Split (_df .Clean (_bf .Name +"\u002f\u002e\u002e\u002f"));_bc ._cc [_bab ]=_dba ;_db ++;};};};_db --;};return nil ;};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";


// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_g .Writer ,zipPath string ,data []byte )error {_ecg ,_cca :=z .Create (zipPath );if _cca !=nil {return _a .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_cca );
};_ ,_cca =_f .Copy (_ecg ,_c .NewReader (data ));return _cca ;};type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};func (_b *DecodeMap )IndexFor (path string )int {return _b ._fgf [path ]};

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_fa *DecodeMap ,_gbdb ,_de string ,_ef []*_g .File ,_fg *_dfa .Relationship ,_ac Target )error ;func MarshalXMLByTypeIndex (z *_g .Writer ,dt _gbd .DocType ,typ string ,idx int ,v interface{})error {_bg :=_gbd .AbsoluteFilename (dt ,typ ,idx );
return MarshalXML (z ,_bg ,v );};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_af *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_af ._ec =fn };func (_cg *DecodeMap )RecordIndex (path string ,idx int ){_cg ._fgf [path ]=idx };

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_g .File ,dest interface{})error {_bfg ,_cea :=f .Open ();if _cea !=nil {return _a .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_cea );};defer _bfg .Close ();_bcbe :=_ff .NewDecoder (_bfg );
if _gbb :=_bcbe .Decode (dest );_gbb !=nil {return _a .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_gbb );};if _dcff ,_ecd :=dest .(*_dfa .Relationships );_ecd {for _ea ,_bfga :=range _dcff .Relationship {switch _bfga .TypeAttr {case _gbd .OfficeDocumentTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .OfficeDocumentType ;
case _gbd .StylesTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .StylesType ;case _gbd .ThemeTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .ThemeType ;case _gbd .ControlTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .ControlType ;
case _gbd .SettingsTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .SettingsType ;case _gbd .ImageTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .ImageType ;case _gbd .CommentsTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .CommentsType ;
case _gbd .ThumbnailTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .ThumbnailType ;case _gbd .DrawingTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .DrawingType ;case _gbd .ChartTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .ChartType ;
case _gbd .ExtendedPropertiesTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .ExtendedPropertiesType ;case _gbd .CustomXMLTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .CustomXMLType ;case _gbd .WorksheetTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .WorksheetType ;
case _gbd .SharedStringsTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .SharedStringsType ;case _gbd .TableTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .TableType ;case _gbd .HeaderTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .HeaderType ;
case _gbd .FooterTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .FooterType ;case _gbd .NumberingTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .NumberingType ;case _gbd .FontTableTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .FontTableType ;
case _gbd .WebSettingsTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .WebSettingsType ;case _gbd .FootNotesTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .FootNotesType ;case _gbd .EndNotesTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .EndNotesType ;
case _gbd .SlideTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .SlideType ;case _gbd .VMLDrawingTypeStrict :_dcff .Relationship [_ea ].TypeAttr =_gbd .VMLDrawingType ;};};_gg .Slice (_dcff .Relationship ,func (_efb ,_da int )bool {_ded :=_dcff .Relationship [_efb ];
_efd :=_dcff .Relationship [_da ];return _dc .NaturalLess (_ded .IdAttr ,_efd .IdAttr );});};return nil ;};var _fd =[]byte {'\r','\n'};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_aeb *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _aeb ._ae ==nil {_aeb ._ae =make (map[string ]Target );_aeb ._cc =make (map[*_dfa .Relationships ]string );_aeb ._gc =make (map[string ]struct{});
_aeb ._fgf =make (map[string ]int );};if _df .IsAbs (filePath ){filePath =_dd .TrimPrefix (filePath ,"\u002f");};_fe :=_df .Clean (filePath );if _ ,_eg :=_aeb ._gc [_fe ];_eg {return false ;};_aeb ._gc [_fe ]=struct{}{};_aeb ._ae [_fe ]=Target {Path :_fe ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };
return true ;};func MarshalXMLByType (z *_g .Writer ,dt _gbd .DocType ,typ string ,v interface{})error {_fdf :=_gbd .AbsoluteFilename (dt ,typ ,0);return MarshalXML (z ,_fdf ,v );};