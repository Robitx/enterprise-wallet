package templates

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"addressBook.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xecWAo\xe2:\x10>ï\xb0\xfc\xaa\xa7\xf6\x90Fz\xc7>\x13\x89\xa2ݽT\xbd\x94\xfbj\x88\a\xb0H\xec\xc8\x1e\xe8V\x88\xff\xber\x9cP7\x98v\xb5{\xa8\xaa\xed\tg\xe6\xf3\xe7o<\xf3\x01\xd9\xef%.\x95F\xc6AJ\x8b\xce\xdd\x1a\xb3\xe1\x87\xc3x\xbc\xdf_4\xb0¹\xa2\n\xd9̈́\xf1i\x00\xb0\x1e\xb1\xdf_Ԡ\xf4\xac\x02\xe7Z@ǐ-\x9e\x01\xb4ƺ\xdd}\xfd\x80DJ\xaf\xdc\xf5܇B\x16JR;\xbc\x87\x9dG\xfc\x17bJ/\x8d\u007f\xac7Sk\xe1\x89E\":\xb6h[$\xe0x@\xd0NX7\x15\x102ޯ\xe6\xa6ᬥ?\x1c\xc6B\xaa\x1dSr\xc2W[r\xbc\x183Ƙ\xd8V\xac\xf4\\\x13N\xb0p\x9cI \xc8\xfc\xb2E\x92\x05\xed\xfc\xd1FgQ\xde4>\xe2&\\\"6\xdf+\xa57J\xafn\xc8n\xb1\xa3m\xa9+\x15Sg\xd4֣\\\x16J\xe1\x85\x00\xb6\xb6\xb8\x9c\xf0\u007f\x96P\x92Q2\xeb.\x13\x1dg`\x15d\x0e+,\t[!\x9e\xfbk\xc0\xb1i\x8f\x139\x14\"\xaf\xd4\x1b\xa7\xc6g\x95\x16\xa5\xa2\xe8\xa8\xe2\x8b&\xfb\xc4fm\xfc\x8f\xa8\xf1\a\xa1\xd5P\xc5\xe43\xa3\tJbw\xca\xd1KJ\x91o\xabn\xe5\x1b\x133\x97F\x13j\x8a\xbaч\x12-\x89\x04\x0ey\x1a\xd0XEW\xde\xf6\xf4\xf4\xb2\x9f\x19Z\x16\x82\x85oT\n\x9a\xb5\xb9\xc1\x86\x91\xa05\x82\x1c\x04G\x82\xec0䑅\xc8i\x1d\x12\xa3Q\x1f\x82\xf68g,e\x1aj/I\xa8A(S\xa5Ѽ/o\tl\t\x99\xcfv\xa3\xb2VR\xa2\xee\aE\xe4\xaa`\xff\xea\x85k\xfe\xbfW\xe5\xc6\x13\x84\xcb\xef\x8f\x1ej\x8a\x048D\x89\xf2\xa5\x82\x10\xfb]\t\x0f\x88\x92}C\x8d\x16\b%\xbb\x14\xae\x01=\xa0!\xb0gh<\xb8\xb8\xfae\xf9\v\xa8@\x97\xc3+\xec\xa3\xe7JȠ6[\xfdf%\xb7\x81\xe6\x9c\x18\x91\x9f\xf4\xdc\xc3N\x86#\x8c\xd9\xc2ȧD<O$D\xde\x0e\xde \x18\x8d\xbb\xc6\xc7\xd9|\xca\x13t\xbd;5>\xf6s|\xac\xff\xf8\xaca\x97\x95\x95*7\xcf5\xc3P\x81T\xbb\xc8j\x83Ǥ\xf3\x82\xdfN\xbfp\xce\xd9m\x88L\xba-\xecI\xdfi\xc8\xd9t\xa2ۘh\xdc\t\xe2\xdd\xec\xf8\x9a\x92\x8f\xe4\xcb\xd7\xeax?\x83Fn\xb2I\xe3\xfd\xe5NM\xfd~\x9f\xf3\xea)\xf6ӭ\x9f>\xf8\x18>\xe8\x96\xddG\xea\xbd\xe1\xd6\x10\x99\x9a3~g@\x1e\xff\x12_^\x857\x1c\xd4\xf2p\x18\xff\f\x00\x00\xff\xffVH\x12\xf8F\r\x00\x00",
		hash:  "6f8113fb8963552488892113449964683f009446d0b293e061feeff5b664f1c9",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  3398,
	},
	"create-entry-credits.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xdcUMk\xdb@\x10\xbd\xfbW,C(\xedA1\xbe\xb6\x92\xc0\x11)\x14J/\xc9\x1f\x98hG\xd6\xe2ծ\xd9\x1d\xa9\x18\xa3\xff^\xb4\x92\x9c\x95\x92CL\x0f\x85\x9el\xcdǛ7\xf3\x9e\xd0\xe5\"\xa9R\x86\x04\x94\x8e\x90\xe9Ѱ;\x17\x8e\xa4b\x0f}\xbf\xd9\\.w'<гbM\xe2k&\xe0\xa1=\x8bP%\xa2\xb2\xcb\xe5\xaeAe\n\x8dއ\xaa\x11-\xa1\xa10)\x17\x85\\S\x13\xa0\ue7c8Y\x99\x83\xbf\u007f\x1eBc\x16KV\x1d\xfd\xc2n\xa8؍1e*;<6ǽsx\x16\x11\xa3\t-j\x8b\x88\\\a\x04\x18\xa6椑I\xc0\xfc\xefٞ@\x04\xf4\xbeߤRuB\xc9\f\x0e-{\xc87B\b\x11b倕\x81\xb3\xbf\xa7\xe8:SZ\xdd6\xc6G\xd9P\x81\xc2`\x97\x94Z\x95\xc7\fص\x04\xa2vTe\xb0\x85\xb9\xf3\xa5e\xb6F\x94\xdazJ\xc6\a\x10\x12\x19\x93\x10\x12\xe8\x14&\x1a_HgP\x84\x88\xa3\x8ePC\x9e\xfa\x13\x9a1_+)\xc9L#\xf2O\xac\x1a\xf2\xdf\xd2\xedP\x90\xa7[\\\xb1\xaawaI\xb6G2IM(Ɂ\xe8P\xb7\x94\xc1\x0e\xf27\xea\xa6\xdbz\xb7\x84\x18$\x1c\xba\u007f\xacT\x81\xc7\x02\x04\x04\xc5Ũ8\b\xa3t$sa\x95)\xacagu߯0#u\x1c\x1a?\xe8iM2GA\xbcΌZӭT\xdd$U\xf47Vm\xbcW\xd8\x19\xa5t\xe4\xfd\x8b\xb5\xc7\xe9\xcac6\xbaxb\xcdR\xb1\x90B\xa3\x1a\f|\x94ɠBI\x892or\xb6\xe5)i[\x8e\xadR\xef\xf2'\xd2T\xb2؏\f\x967M+뚕J\xb3\x17'\xca\xe4\x93Y\xf8e\xdd\xebҋ\xd6\xc0\x8c\xed\xe1\xa0i\xb5\xf6\xd2x\x9e\x8cL*,\xd9*9M\xb8\xda\xf0\xbd\xf1s2\xdfK\xf9\xbâC\xba]\xae\x93^=>\xbe+7[]\xf0\xf9D3\xe3\xf5\xfe\x1f{\v\"r#̇-S\x11%\xff\xd06Em\x87K|\x1f\x15\xba\xdd>\x11\xfd\xbf\xb2Л3\xdcf\xa3\xf7h\\\xad4\xed\xf8\xbf\xb8i\xfay\xefs\xf3`\x99m\x03\x02~Z\x94\xfb\xf9\x1eOddaMG\x8e?\u007f\x19?\x91dd\xdf\xff\t\x00\x00\xff\xff̓Lx\x9a\a\x00\x00",
		hash:  "b8282ebb635408a95e416dc00a6394eab05caf485013590f8f81ba26d83676d1",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  1946,
	},
	"edit-address-entry-credits.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xc4V\xcdn\xe36\x10\xbe\xe7)\x06Ģ7\xc6\xcdb[lSI@\xe2\xe6T X`\xf3\x02crl\x11\xa6H\x81\xa4\xd4U\r\xbf{ARJ\x1d[N\xdc`۞l}\xfc43\xdf\xfc\x89\xbb\x9d\xa4\xb52\x04\x8c\xa4\n\x1c\xa5t\xe4='\x13\xdc\xc0\x85\x8b\xa0g\xfb\xfd\xd5\xd5n\xf7\xa1\xc5\r=\xa9\xa0\tnK`\x0fR\x05x\x88<X&\x1e\xdc\xe5\x97#}\xb7\xfbР2K\x8d\xde'\xb6'#\xf9\x1aE\xb0JN\x8cPS\x93l]\u007f\xa5\x10\x94\xd9\xf8\xeb\xa7\b\xe5S\x14A\xf5\xf4\x88}d|̘2k\x1b\x1f\x9b\xed\x9ds8\xc0AH\xa3\xb5\x83\xd7\x0e\"xv\x90\xcc\x04jZ\x8d\x81\x80M\xff\x9el\xcb Y\xdf\xef\xaf\n\xa9zP\xb2d\x9b.xV]\x01\x00$LD[%s\xf6\x8f\x11=>\x11Vw\x8d\xf1\a\xa7\x89\x81`\xb0\xe7B+\xb1-Yp\x1d1\xa8\x1d\xadK6&\xec\xde\xda-\x9bl\xac\xba\x10\xac\x01\xa1\xad'\x9e\x1f\x18H\f\xc8\x13\x04\xe8\x14r\x8d+\xd2%[&\xc4QO\xa8YU\xf8\x16M>\xaf\x95\x94dFg\xd5\x0fA5\xe4\u007f-\x16\x91P\x15\v<\x8a\xaf\xbe\xa9\xce\x16\xb3X\xd47G\xf4\xb5u\xcdK\xe8\xf5\f\x9dc5$U\xd7\xf0\xcf0\x9f\xb6\x17\xaf%\xc1\u0557n\xa5\x95\x80-\r\xb7\xc5\"C\xf3^\x16R\xf5\x17\a\xf0\xe9\xf2\x00\x1e\x95\xd8\x1al\xe8]\xee\xcf\xc1\xffj\xe2ZGU\xa1Lۅ\xd4\xd2\xd3|\xaf\x15i\xc9 \f-\x95,з\xc0@*\x8f+MrjѨ\xb3dmJ9\xdf\xd2\xc0\xa0G\xddQ\xc9v\xbb\xeb\xb19\xf6\xfb\xe7\xb6m\xc9\t2ᗟXU,\xa2\xd3\xff\xaa0\xa7\xdab\xe0/\xa5e)\xa8\x15\xfaC\x15\x8f\x18w\r\xfb\x9f\x8bX+I|m\x1d\xf7\rjͭ\xd1\xc3\x1b\xbas3&\xcd+\xd4h\x04qaM@eȱ\xea>C\xb7\x90\xf7\xc1\x01\x8bU?\x8e;\x00\x1e\x96\xdfu\x80r\xe87\x1f/)\x18\xa6\x90<\xf6\x94*\xc5E\x8dfC\xc7\xfb\x8f\xbe\xb5h$IV}Ş`\x99H\xfedw\xbd\xb3Z\xb5{\xb3\x84Q\x8a\xc6\xd6\xd3\x05\xb5\xd4\xe86\xc4?\xc1\x98\x8f\x9fߑ\x0f\xa9|\xabq\xe0\xadS=\x06\xca\x03w.%\xbfe2|\xc9d\xf8\x9d\x86\u007f\x9a\x99y\t\x9f\xdf'\xe1h\xc9\x1ch\x98\x16\xcd\xcc8\xbeP\xea\be\xec\xfci\xf9\x8cC\xfa}\xd6\xcb\f\\,N\xbfbo\xcf\xf0̷\x1e\xa2$\xeeԦ\x0e\xe7\x1a%\x17X\xd8v\xe0\xc1\xc6k@\xbb\xb2\xe8\xe4\xc9\x17?\x12\xc6\x1dƪ\xa5m\a\b\x16\x96\x13}\xb6\xc0\xb3\xcaN\xa1\xe9BC\xceY\xc7\xff\xb4\xe6\xefq\x13\xa8\xb5\xed\x02\xa0&\x17\xb2\x98\x98hr\f|\x184=w歉\xafUw\x06\x92\x15\xa8у\x15\xa2s$_s\xe9;!\xe2R\x9eu:\x1e^\xe0\xf65\x89\a\x8f\xe3\xdf\xf1g\xf6\xaawoC\xb0\r\x03\x96/\xa1d\xe4~\xffW\x00\x00\x00\xff\xff~\\E\xc6\x05\v\x00\x00",
		hash:  "5553bfe979027b367155e218340b13648b02173c60c55fa467580a3d19c51bfa",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  2821,
	},
	"edit-address-external.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xc4V\xcdn\xe36\x10\xbe\xfb)\xa6Ģ7FȢ\x05\xda\xd4\x12\x90u\xf7P\xa0\b\nl^`L\x8e-\xc2\x14)\x90#mTC\xef^\x88\x92R\xc7Q\n\xc3\xd8fO\xb6\x86\xdf\xfc|ߐC\x1e\x8f\x9av\xc6\x11\b҆%j\x1d(FIOL\xc1\xa1\x15}\xbfZ\x1d\x8f\x1fj\xdcӣaKp\x97\x83\xf8\xac\r\xc3\xe7\t\x02\xf7\xa3\xcf\x00=\x1e?Th\xdc\xc6b\x8c\t\x19\xc9i\xb9C\xc5\xde\xe8\x19\xc1%U)\xce\xcd\x17b6n\x1fo\x1e\aӸ\x8a\x8aMK\x0f\xd8\x0e\x88\x8f\xa3\u0378\x9d\x1f>\xab\xc3}\b\xd8\xc1I9S\xb4\x13\xb7\x93\n\x9e\x13\xa40LUm\x91\t\xc4\xfc\xef\xd1\xd7\x02R\xf4\xbe_\xad\xb5i\xc1\xe8\\\xec\x1b\x8e\xa2X\x01\x00$\x9b\x1ab\xe5\"\xf8\xaf\x93\xf5|Ey\xdbT.\x9e\xac&\x04\x82\xc3V*k\xd4!\x17\x1c\x1a\x12P\x06\xda\xe5b\x12\xec\x93\xf7\x879\x04l\x1bf\xef@Y\x1fI\x8e\x1f\x0242\xcad\x02\f\x06\xa5\xc5-\xd9\\l\x92%PKhE\xb1\x8e5\xbaq\xbd4Z\x93\x9b\x92\x15?\xb2\xa9(\xfe\xb6\xce\x06@\xb1\xce\xf0\xac\xbe\xf2\xb6H\x8d\x9cʁ?\x1cl\xbccT\f\u007f\x9a\xc8물=\xf3\xd8\xf9P\xbd4\xfd\xb7Ho\xa1*Ҧ\xa9\xe4/\xb0\xac\xdc\v\xb7Ĺ\xf8\xab\xd9Z\xa3\xe0@\xdd\xdd:\x1bM\xcbY2mڋ\v\xf8\xe9\xf2\x02\x1e\x8c:8\xac\xe8\xaa\xf4o\x99\xffW\xe1\xea@\xc5ڸ\xbaᴫ瓽3d\xb5\x00\xeej\xca\x05\xd3\x13\v\xd0&\xe2֒\x9ew\xe9\xc03\x17u\x92\\\x1e\xa8\x13Тm(\x17\xc7\xe3ʹY\xfa^\xcc5\xd5\x14\x149\xfe\xf5gQ\xac\xb3!\xe9{5\xe65\xb7\xa1\xf0\x97\xd4F*h\r\xc6S\x16\x0f8\x8c\x1b\xf1\x9d\x9bX\x1aMr烌\x15Z+\xbd\xb3\xdd\t\xefo\"\xe3\x18\xf9\xf6\xe3%zb\xd22bKIH\xa9Jt{z\xee\xf34\xa1\xe8\xa9F\xa7I\x8b\xe2\v\xb6\x04\x9b\x04\x8a\xaf\xa6˕b\x96a\xc1\xf8\x83\x94W\xc8|\x11\xe3q g_\xd1Zb\xb9oLv\xce\x17-\x05\x16\xc5\xefd\x89\tv\xc1W\xcf\x03s\x18\xe0\xd7\x11\a)\xbf\xc5N\xba\xb8\xa9:U?\xdf\xee\xefB\xf1\xe5͑\xbd\xbe:.l\xeb\x02_\x18ζ\ff_\xf2[\"\x8d\xb4\x95\xaf;\xc9~\xb8\x83\xeb\xadǠϙ'\xc0\xacJ\xb1\xf1u\a\xeca3\xc3\x17\x99/\xd2[l\xea\xf3\x8b\x82B\xf0A\xfe\xedݿ\xa7I\xa1\xb5\xbe\xe1Q\xfb\x91\xd00E)\b\x88\xdcYʅ6\xb1\xb6\xd8ݹ\xc1\xad\xb8w\x90\xa2@\x89\x11\xbcRM \xbdTɜ26J\r#q1\xe9\xb4xA\xda\xd5\x02\xeb\xe9\xef\xf4\xb3\xf8\xb0\xfa\xe4\x99}%@\x8cO>r\xba\xef\xff\t\x00\x00\xff\xff\xf87`+j\n\x00\x00",
		hash:  "b5dcad57bd00e1fc8228987970275b235546c2aad1b45c20d7ccc38a28771ecc",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  2666,
	},
	"edit-address-factoid.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xc4V\xcdn\xe36\x10\xbe\xfb)\x06Ģ7\xc6\xc9b[lSI@\xe2v/\x05\x82\x056/0&\xc7\x16a\x8a$HJ]\xd5\xf0\xbb\x17\x14\xa5Ա\x95\x8d\x1blۓ\xad\x8f\x9ff\xe6\x9b?j\xbf\x97\xb4Q\x86\x80\x91T\x91\xa3\x94\x9eB\xe0\x1b\x14\xd1*\xc9\x0e\x87\xc5b\xbf\u007f\xe7pK\x8f*j\x82\xdb\x12\xd8oRE\xf8\x94\x19p\x97\xdfH\xcc\xfd\xfe]\x83ʬ4\x860\x10\x03\x199\x99\x9a\x18\xb1\xa6f0s\xf5\x85bTf\x1b\xae\x1e\x13\x94OQD\xd5\xd1\x03v\x89\xf1>c\xcallzlvw\xdec\x0fGь֎^;\x8a\xe0\xc9\xc1`&R\xe34F\x026\xfd{\xb4\x8e\xc1`\xfdpX\x14Ru\xa0dɶm\f\xacZ\x00\x00\f\x98H\xb6J\xe6\xed\x1f#zz\"\xacn\x1b\x13\x8eN\a\x06\x82\xc1\x8e\v\xadĮdѷĠ\xf6\xb4)٘\xb0{kwl\xb2\xb1nc\xb4\x06\x84\xb6\x81x~` 1\"\x1f @\xaf\x90k\\\x93.\xd9j@<u\x84\x9aUEph\xf2y\xad\xa4$3:\xab~\x88\xaa\xa1\xf0K\xb1L\x84\xaaX\xe2I|\xf5M5W\xc7bYߜ07\xd67ϡo'\xe7%VCR\xb5\r\xff\b\xf3\x19{\xf6ڠ\xb5\xfaܮ\xb5\x12\xb0\xa3\xfe\xb6Xfh\xde\xcbR\xaa\xee\xe2\x00>\\\x1e\xc0\x83\x12;\x83\r\xbd\xc9\xfdK\xf0\xbf\x9a8\xe7\xa9*\x94qm\x1c\xba\xf9i\x9e\x15i\xc9 \xf6\x8eJ\x16\xe9kd U\xc0\xb5&9ug\xd2Y27\xa4\x9c\xef\xa8gСn\xa9d\xfb\xfd\xd5\xd8\x1c\x87\xc3S\xc7:\xf2\x82L\xfc\xf9GV\x15\xcb\xe4\xf4\xbf*̹\xb6\x14\xf8siY\nj\x85\xe1X\xc5\x03\xa65\xc3\xfe\xe7\"\xd6J\x12\xdfX\xcfC\x83Zskt\xff\x8a\xee\xb1\x19\xefQ\xa3\x11t\vy\xe6S\n\xd6\x19b\xd5\xf5\xd5\xf5\xf58\xea\xf0i\xf5\xf8]\xa7%\xc7y\xf3\xfe\x92\xea\xe0\x10V\xc0\x8e\x86\xb2pQ\xa3\xd9\xd2鞣\xaf\x0e\x8d$ɪ/\xd8\x11\xac\x06R8\xdbQo,M\xed_\xadW\x92\xa2\xd1\x05\xba\xa0p\x1a\xfd\x96\xf8\a\x18\xf3\xf1\xd3\x1b\xf2!Up\x1a{\xee\xbc\xea0R\x9e\xae\x97R\xf2k&\xc3\xe7L\x86ߩ\xff\xa7\x99\x99\x97\xf0\xf1m\x12N6ʑ\x86\xb9\xad2n\x91c\xa1\x9eP\xa6.\x9f\x16\xcd8\x90\xdfg\x95\xcc\xc0\xc5\xf2\xfc\xc6z}^g\xaetH\x92\xb8W\xdb:\xbe\xd4'\xb9\xbeº\x9eG\x9bn{\xb7\xb6\xe8\xe5\xd9Ş\b\xe3\xbeb\xd5ʺ\x1e\xa2\x85\xd5D\x9f\xadגּsh\xfan!\xef\xad\xe7\u007fZ\xf3\xf7\xb4\t\xd4ڶ\x11P\x93\x8fYLJ4y\x06!\xf6\x9a\x9e\x1a\xf3֤ת;\x03\x83\x15\xa81\x80\x15\xa2\xf5$\xbf\xe52\xb4B\xa4\x05<\xebt<\xbc\xc0\xedbF\xde\xf8w\xfc\x99\xfbz\xbb\xb71چ\x01\xcbߕd\xe4\xe1\xf0W\x00\x00\x00\xff\xff\xado*\v\xcd\n\x00\x00",
		hash:  "e1ed78d7b673596fd4e8322d0849a9003c146439576675ee72d8c6fd55b5d1d5",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  2765,
	},
	"import-export-transaction.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xdcV_\x8f\xa36\x10\u007f\xdfO1\xb5V\xa5}\xf0E\xe9c\vHۨ+\x9dT\xf5%Q_+\a\x0f`\x05ld\x1b\xf6N(\u07fd\xc26Y\x93p\xba\xbd\xf4a\xab{\x02f\xc6\xf3\xef7\x9e\x1f\xe3ȱ\x14\x12\x81\x88\xb6S\xdaR\xfc\xe4\x1eV3iXa\x85\x92\xe4|~x\x18\xc7ǎUx\x10\xb6A\xf85\x03\xf2љo\xfep\xe6pX\x9a\x8f\xe3c˄\xdc5\xcc\x18gmPrZ\xb2\xc2*\xc1M\xb0\xb05\xb6\xceׇ=Z+de>\x1c&\x91\xd7N\xce\x06\xfc\x8b\r\x93\xc5\xd6˄,\xd5\xf4ٞ\x9e\xb4f\x9f!J)x\x8b\x8eE\x19\\\x0287\x16ۮa\x16\x81\xcco\a\xd5\x11p\xde\xcf燔\x8b\x01\x04\xcfH\xd5[C\xf2\a\x00\x00'+&_\x19\xd1\xea%H\xaf5\x85j\xfaV\x9aH\xeb,\x18H6Т\x11\xc5)#V\xf7H\xa0\xd6XfdC\xe6\x93\xc7\xdeZ%\xa1h\x94A\xea?\bpf\x19u\"`Z0ڰ#6\x19\xd99\x89\xc6\x01YC\xf2\xd4tLz}-8G\x19B\xe4?ZѢ\xf9-\xddL\x06y\xbaaWY\xd5[W\xa4U'\x94\xb4F\xc6Q\x13\x18X\xd3cF~!\xf9\x17\xe15\xe9\xa6\xde^\xf9Zv\a\xd6\xfb\xe0,\x85\xecz\xeb\x02\xf7]\xa3\x18GNK\xd1 \x01\xfb\xb9\xc3,\x99\xde\x13\x90\xac\xc5,\xe9\r\xeag\xf7\x1d\\\xbb\xb3\xb4Ҫ\xefh)\xb0\xe1P\v\x8ekQ~\xa0t\x0e\xe5\xfd\x9a\xfe\xd8\n{\xf1\xecB\xffs\xb42\t\x15\aїC\x91\x1c(]\x89\xc4\\-\xe1\xea\xf8J\x16\x90\u038d\x84R\xab\x16&\x83[$6\\\fK\xd18>r1\xfc\xae\xacU\xed4\xef\x9d\x16\xd2\x02Ik\x9dOw\xe7릑\x8cD\xe8$Z\xbd$w\xb8\x88\x11N\x02\xbaw\xf9Y\xa2R\xd4X\x9c\x8e\xeaS2511\xa2\x92\xf1ڙѺ\x95\a̦9\xbf3\vw\x95\xa0Tz\xc5}\xbe\x17\x95\x84H\x92n\x9c\xf9]}s\xd8ށٽ\xe7\xfeWX\xbbˑ\xb4\xec\x84\x14\xa5\x15\x1a\x970\x86\x18\xfe\xa2\\P\xdd&\xf9\xdf\x02_\xe2\x953]\x99wi\xfe\xf5A\xb7,?^\x91\x10y\xde\x1d\b\x90@n$v5M\xe8M\xf8\x88~^K\xa4\xb3\x94\xc0k\x94\xe8h\xb4$\xe2\xd7x\xf1zBp\xfb\x88q\xaeј\xa3R\xa7@#^\x1bQ\nUrIINŤh\x99\xcbGȌ\x94\x8c#\x15\xf2F\xa7z\x1b\x94\xaa\xb71\x17\xd6\xdb|\x8f\r\x16\x16\x9e|\x06K\xaeHK\xa5\xdb\x15\xea\x88RFCgf\xfbڒtG]fVUU\x83We/\x995\xfe\xff\b\x11.<\xbb\x16~V\xe6O\x9c\xbf\x16\xb3\xc8!\xdd,\xcbI/$\xee\u007f\x06\xbe\x99\xcb\xfdV\xbc\x10ǲط\xd1|\x94\x9cw\xf3\xe6\x91)\x11\xe9\xfb\x8fͳG\xe8\xedc\x13\xa5\xfd\x9fF\xe7\xa6\xfco\x1b\x9f\xb54.#\xb4\xabՄ\xf2\xf72E\xe1\xb1\xf6\x1f\xed\xd7\x1e\x01\xf2\xa7b\xfci\xee\xc7\x1e%\xdf)9\xa0\xb6?\xfd\xec\xff\xfdQ\xf2\xf3\xf9\xdf\x00\x00\x00\xff\xffB\xe4\x12\x82{\f\x00\x00",
		hash:  "5a485185b7f0315be7801f7faf4baed3d8651fbf39991faf37c92c7958eccbcc",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  3195,
	},
	"index.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xa4TAk\x1b=\x10\xbd\xfbW\f\"\x87\xef;l\x96BNamHL\n\x81RB\xe3?0\x91\xc6\xf6\x10\xad\xb4Hc7\xee\xe2\xff^\xb4Z;\xb2\xeb@CO^\xcd\xcc{\xf3\xc6O\x9a\xbe7\xb4dG\xa0\xd8\x19z{\xc2\x15\xa9\xfd~2\xe9\xfb\xab\x0eW\xb4`\xb1\x04\xb7SP\x8b\x80.\xa2\x16\xf6.\xa6\x8a\xbe\xbfj\x91\xdd\xdcb\x8cC\x81\xfcY kj\a\xf4\xf53\x89\xb0[\xc5\xebE\n\xe5l\xaa\xdd\xd2wܦ\x8a/9\xc6n\xe9ӱ}\xbd\v\x01wP\x88\x18\xd9\nX!\xe0\xd8 k\x17j;\x8bB\xa0\x0e_\v\xdf)\x18\xe8\xf7\xfbIcx\vl\xa6j\xb5\x91\xa8f\x13\x00\x80&v膠\xf6N\x82\xb7\x1d:\xb2U\xe7\x83(\x88\xb2\xb34U\x86cgqw\xeb\xbc#5\xeb\xfb\xf7\xb1\xe6\x19\xf3\x940O>\xc8~\xdfԉp\xe4\x16|\xb1\x94\xbf\xf3\xf9ś\xddЬ\xf8\xd7*\xcbQ\xd4{\xd5PyPj=\x1av\xab*\x89Cv\x14\x14\xe84y\xceP(\x12\xa7\fG\x96\xd3\xf2$\xe08\xd7\x05\xc8\x00[ߜ\xf4n)\xc6t?f\x8f\x8e\x85\xd1\u0098\x80\x16w \xf8J\x10}K \xdcRS\xafo> =\x9b\xe8l\x0e5kj\xc3ۿ\xc0\xbe\xcb\xf9\x00q!|\x16j\xea\xc1\x88Ѥ\xbap\xe9؉B\xf0\xa1\xfa\x95\f?\b\xd5h\xad\xdf\b\xa0\xa5  \xf4&\x95&'ɒ\x8b\xf7\xe4\xce\xc1\xc0\x02k\x8c\xe0\xb5\xde\x042\x85\x90\xd2\x1cG?狻\u008e\x06\x0f\xa9⢤\x89\xb1\xa8\xd9س)-\xcf\x1a\x04\x87\xdbJ[֯\t\xbc!\x05\xeb@˩\x8a\xe4L\xb5D-\x9eMT\xb3gr\x06\xbe\x8e\xc7\xc4\xdbԖ?\xc3\x17H\x13o\xa9\xa0\xfc\x91#\xffĪ\x03\xa1PEN®ҁ\f\xa7\x97z\xbf\xd9\xc1C\n\xc1<\x87.S\xf7=/\x8b\x9d\xf3ئW\xfc\xf0\xd6\r\x0f\xf3\x13\"x\x00V4 \xab\x13\a2g\x9dI\xa1X\x8e\x1f)\"g\x8a\xdeM}0m\xbc\t\xe3ϥ\xd5u\xefE|\xab@}\xf3h\xca5\xfc\xdf\xffy\xcff\xea\xdf\x01\x00\x00\xff\xff\xee\xfd\xdaa\xca\x05\x00\x00",
		hash:  "19896855ed8167192e587146a4947ad2cc86cfd6c689e52592cb8a7128dd3843",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  1482,
	},
	"new-address.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xb4V͎\xdb6\x10\xbe\xfb)\x06D\xd0\x1b#l\x90\xb4\xe9V\x12\xe0\xe4P\x14\x05\xb6E\xbb/0&G\x16a\x8a\x14H\xca^\xd7\xf0\xbb\x17\xa4$\xffld\xaf\x11#'\x8b3Ï\xe3\xf9\xe6\x1br\xb7\x93T)C\xc0\fm8J\xe9\xc8{\xb6\xdf\xcff\xbbݻ\x16\x97\xf4\xac\x82&x,\x80=\xd1\x06\xe6ǀ\xdd\xee]\x83\xca|\xd5\xe8}\xf2{2\x92W(\x82Ur\x8c\b55i\xf7\xfb\u007f)\x04e\x96\xfe\xfds4\xf5^\x14A\xad\xe9\t\xd71\xe2CoS\xa6\xb2q٬\xe6\xce\xe1\x16N\x92\x18\xd0N\xb6\x9ddp8 \xc1\x04jZ\x8d\x81\x80\x8d_϶e\x90\xd0\xf7\xfbY.\xd5\x1a\x94,ز\v\x9e\x953\x00\x80d\x13\x11\xab`\xcen\x06\xebk\x8f\xb0\xbak\x8c?\xf1\xa6\b\x04\x83k.\xb4\x12\xab\x82\x05\xd7\x11\x83\xdaQU\xb0\xa1`_\xac]\xb1\x11cх`\r\bm=\xf1~\xc1@b@\x9eL\x80N!\u05f8 ]\xb0\xaf\xc9\xe2hM\xa8Y\x99\xfb\x16Mﯕ\x94d\x86\xc3ʟ\x82j\xc8\xff\x96g1\xa0\xcc3|\x95_\xfdP\x9eЗg\xf5ë\x80ʺ\xe6\xdct\xbd&\x97\xa2|\x83Z\xf3\x87\x0fАT]\xc3?\x83F\xb7$\xfe3LW\xee\f'\xfd\xe7\xf2w2\xe4\"wC7B\xe5l\xf3\x98g\xbd\xf7\xf2nO\x9aD\xe8i\x1d \xb8\xb7\x9d\x13\xc4\xc0`Cߚ/\x83%@\xdb\x06e\r\xacQwT0\x87F\xdaflpV\xfe\x93\xd6`h\x03\x83mL8\xcf\xfa\x9d\xdf\x03O\xe2\f\x99Lp[\x10\x8e\xa4\nw\xc1\xab\xa6\xb5.\x1c\x04^\xfe\x91֩\xb4\xd0:\xb5\x8e\xf5^\xd1\xf6\x1e\xec\x95UFU\xdbs\xec?{#\xb4\xb5CO\xdf\x05\x1f\a\x13\xbd\x04r\x06\xf5\xf1\x0f\xc4v\x16\xd6\x04\x8c\x94\x9bçV>\xbc}J\x9e\xf5\xcdr\xa1\x9f3\xa9\xd6\x13b\xb8`\xbe[#\xbf\f\x1a\xf9|\xbbF\xfaA\x10[ݓ\xe0m\xb7`\xe5\xdf=\x8b\xc3\b\x88dޠ\x99\x93\xac\x94i\xbb\xc0\x97\xcev-K\xc8C[\xf0\x15my\xefLEV\x86\xdcU\xe5䭣2O;\xa6q\x18\x84mK\x05\v\xf4\x12Fi\x9e\x04\xb1\x89\x8cx\xa5HKh\xc9\t2\xe1\xd7O \x95ǅ&9\"\xb6\x1a\x05\xd5VKr\x05\x9b\x9b)Y\xc2Fi\r\v\x8ar\xc2@\x92\x1d@\xca<\x8b9_i\x97I\xe6\xdfr]\xa1\xfc\xd3@\xf9\xc7\xdb)\u007fRb\x15ku\x03\xab\xc7ڛa\xd3\xe5£V\xe8_\xd7/\xda\xc0Vc\xe5&\x19!\xe7\xec\xa5F\xf8\xd1\n:\\.\xb5\x92\xc4+\xebx_`k\xf4\xf6\xa4\xa2\xf706\x1c\xf1\x11\x0e\xd4\xdd\xc0\x14\xa6\xa2\xa3\x94<\xd8qT-&\xae~zi\xd1H\x92\xac\x9cK\t\xc1\x8e\x173\xc4w\xc27\xb7\xf7\x1dU\x8d\xe9$\xa2\xf8\u007f\xd6\xd0!\r\x81Z\xdb.\x00jr\x01b;\xf0\xa8+r\f|\xd8j*\x98T\xbeո}4q[97\x90P\xa0\x8e}!D\xe7H\xbeu\xac\xef\x84 \xef\xa7\x0f\x1e\x9c7\x1c}\x91\xa8\x12\xe6\x10/\x81\x83\xbaяʆ\x8d\n5\x84\x9a\xa0\xb2Zۍ2K\x90\x14Pi\xff8\x924\xa6\xa7\x95YŇUd\xb9\xfc\xab%s\x00T\xfd\x14\t\xb8ȳ\xde\x1d\x99\xb9\xd2>OQ\x9dp\x1c\xcdQ`\xacL\xe6a*\xcf.\xce\xcb\xe3\xd3lr\x12M\x9c\x9ag\xe7o\xb6\x93\x90\xe1s\xf8\x99z\r\u007f\xb1!؆\x01\xeb\xdf\xe9d\xe4~\xff\u007f\x00\x00\x00\xff\xff\xc6\t\xa2\x9a\v\f\x00\x00",
		hash:  "19a501f82649f41e765e58bab39a12368ddcdbb527f1dd902c094527d752623f",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  3083,
	},
	"notFoundError.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xffl\x8f\xc1j\xc30\f\x86\xef~\n!z\xee\xd8u\xb4\x85ml\xc7\\\x96\x170\xb3\x9a\x8a\xc5Rp\xd4\xc00~\xf7a\xa7\x1b\x19\xf4f\xeb\xd3\xff\xfdv\u0381\xce,\x04(j\xefz\x95\xf0\x96\x92&,Ź\x9cw\x93\x1f\xa8g\x1b\t\x9e\x8e\x80\xddm\xa5Ҝwѳ\xbc\x8e~\x9e\x1b\x94\xff\xd0.\x14[j\xffAf,ü\xef\xebh\xa5\xfe\xd3x\xa1\xce/u\xe3q\x9d\xb1\x9c\xb5^\xe3\xd7sJ\xfe\x1b6\xe57\xdb&\xb6)\xff+h\x1a\xa38\x8d\xde\b\xf0\xf7\xd4\xeb\x84\xd0쥸C\xe0\x058\x1cq\xb8ڌ'ש]X\x06hOw\x87\x87\xc0\xcb\xe9\xae\xe5E\xcd4\"\xe0\xfa?\x92P\xcaO\x00\x00\x00\xff\xff<\x19\x9e\xcb=\x01\x00\x00",
		hash:  "94eccbdbc58c9c9314082dc1b73a3b46e52c84a8f9861c851c2f02ed49d634de",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  317,
	},
	"receive-factoids.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xd4VMO\xe48\x13\xbe\xf3+\xea\xb5Ы\x99\x83;\xd3,\x8bX6\x89\x04h\xd0\x1c\x16\xc42 \xedm嶫\x13/\x8e\x1dlw\x00E\xfd\xdfWq>\xe8i\x12\xa6ᶇV'U\xe5'\x8f\x9f\xfa\xb0\xebZ\xe0Rj\x04b\x91\xa3\xac\x90.\x19\xf7F\nG\xd6뽽\xba\xde/Y\x86\xb7\xd2+\x84\x93\x04\xc8M\x1b\x05\x17\x1bQu\xbd_0\xa9\xcf\x15s.\x04\x8dA\xd5\xf5\xbeϱ\b(\xb3\xef\xe8\xbdԙ\x9b\xdd6\xa6\xd6˸\x97\x15^\xb1\xaa\x89\x98\a\x1b\xecK\xbd4\xcd{q\u007fj-{\x86\r6\x1d\xdcƺ\r\x16\xc3\x17\x02\x8eǢT\xcc#\x90\xfe\xe9֔\xa4E_\xaf\xf7b!+\x90\"!\xd9\xca;\x92\xee\x01\x00\x04\x1bo\xb0\x12b\xcdcg\xdd\xf6p\xa3V\x85v\x1b\xde\x10\xc1@\xb3\x8ar%\xf9}B\xbc]!\x81\xdc\xe22!\x11\xe9W.V\xde\x1b\r\\\x19\x87\xb4}! \x98g4\x98\x80Yɨb\vT\t9\x0f\x16\x8b\x152E\xd2ؕL\xb7\xfe\\\n\x81\xba\xfbD\xfa\u007f/\vt\xbf\xc7Q\x13\x90\xc6\x11\xdbb\x95\xcf\xd3\xed\xecE_\xb5\xb7\xcfpnQH\xef\xe2(\x9fo\xad\x99Ta,\xc2\x15L):?\x00\xc5l\x86\xf4\x10zy\x82\xb8\x0f\x96\x1b\x81#\x10\x01\xe6\u007f\x94B,\x8b\f\x9c\xe5\t\x91E\x16=ؿ\xdb:\x92:\x9b\x95:\x1b\xa4{\xb0tp\x904^\xd8\xf0\x03JG\xd8EBV#\xe6 l\x8f\x97K\x81-E\x87\n\xb9GA\x99\x10\x16\x9d\xa3M\x81\x904\x8eB\xfc\xbb6\u007f\f\xe3\xb51,]\x1a[\x8c\xbb^\x18v\xf9\x92:\x83\x8e\xd1\t\xc4\xe1C\xe9t\x8d\x9d\xb6\x91g\xc6ܓ\xf4\x92i\x96!t6tMU\xc4Q\a1\xb5\xad\xba\x96K\x98uk\xd6\xebi\x92}\xdf\f\xe9\x18\x84[\xca'\x14\x044+p\xc4=\xa4\xb2d\xe2\xe0\v\x99\x96!|\xa5e[׳+\xd6\xcc\n\b\x19/-\xa6\x9f\xea\xfa\x85\xe5\xe78jl\xfd\xe6ކ\f\xd5fJ/\x8d\x86\x8a\xa9\x15&\xa4\x9bUs\x92\xf6O\xf0\xe9\xe2\xf4\xe0@~;\\\x1d\xf3\xea\xf6+\xe3\xbf\x1d^g\xd9\xf5ё\xf5\x97G\xf3\xab\xef\xbf\x14w7.\xff\x8b_\u007f;\xba{\xbc\xfb\xf3\xd7\xcb\xd2<\xf1\xfc\xf8s\x1c\xb5\xc8\xe3%\xf9\x93ҬkT\x0e\xdf\x12\xbd-\xd2qݧ\x15\xff/\b\xd2\xeelT\x13-&$\x89\xa3\xe9Fj{\f\xe0\xed\x02~\x19o\xf0\xce\x1e\xdej\xd63\xa6\x98\xe6x2\xd9W\uf6ec;\r\x9b\x02\x85\\\x15\xf4\xb8c|\xb4\x03\xce6\x96\xd4\xe5\xca\xd3̚U\xb9\xe3\xea\x80\x10\x96\x85*\\\xb4\xfb&\xe0\x9fKL\x88\xc7'OF\xc0\xe9R\xa2\x12 \xa4c\v\x85\x82\x06\x17\x81R1\x8e\xb9Q\x02mB\xbe\x90\xc1ߟj\xbb3\n\xe7\xe2\x06!\xda\xf0\x19e\x12\xf2\xd3\xcc\xf5pR\xee\xa6\xd8x\xbb~4\xec\xf5\x1d\x02\x1aݨ\x95Y\xeea;\xbd\x87\xefM/\xeb\xc7\x03\xe5\xa6|\xa6\xde4\aE\xb90̊W7\x90&`\x98\x12\xe7\xa6|\x06o\xe0\xbc\x0f\u007fu\x8d\xf8\xe0\xae\u007f\x12\xf2\x86{\xaa\xc3G\x96l\x996^\xbb\xc7\xeeo\xf4Jxf\xbc7\x05\x01\xd2]|_\x9f\x81\xb1\xe3V\x96~\xa3Σ\u007fX\xc5Z+I\x1f\xa5\x16\xe6qf\xb42L@\x02\u007f\x18&.\x9as\xb0Cinf!4\xdd\xfba\xca\u007f\x00\xf6\x06\xf9p\xa4\xff\b\x1b\x06e\xf7\xffo\x00\x00\x00\xff\xff\x8d\xf9Ҝ\xdc\v\x00\x00",
		hash:  "665f39932d19c7e9a29f79e33a447e5c480d967818b4f043948102828daa4425",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  3036,
	},
	"send-factoids.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xc4Uϊ\xdb>\x10>'O!\x86\xe5ǯ\ao\x9akk\x1bR\xc3B\xa1\xf4\xb2y\x81Yk\x1c\x8b\xc8R\x90d\x97\xc5\xf8\u074be9\x91\x93P6\x87\xd2S\xec\xf9\xfb\xcd|\xdf\xc4}ϩ\x12\x8a\x18XR<\xa9\xb0tZp\vð^\xf7\xfd\xd3\t\x0f\xb4\x17N\x12\xfb\x921x%\xc5\xd9K\x14\xd2\xf7O\r\nUH\xb4\xd6G\xdc\x14\xe9\xfb'WS\xe3\xf3\x9f_\xc99\xa1\x0e\xf6y?\x9a&/\x96Nt\xf4\x13\xbb1b;ل\xaa\xf4\xf8\xda\x1cw\xc6\xe0;\x8b`\x84jQZ\x84\xe0\xdc\xc0\x97qԜ$:b0?\xed\xf5\t\x98\xaf>\f딋\x8e\t\x9e\xc1\xa1u\x16\xf25c\x8cy[9\xd6\xca\xc0\xe8_\xc1z\xed)\xb5l\x1be#\xaf\x8f@\xa6\xb0KJ)\xcac\x06δ\x04\xac6Te\xb0\x819\xf3\xaduN+VJm)\x99^\x80qt\x98x\x13C#0\x91\xf8F2\x83\xc2[\fu\x84\x12\xf2ԞPM\xfeZpN*\xb4\xc8\xffs\xa2!\xfb5\u074c\x01y\xba\xc1+T\xf5\xd6\x0f\xe9\xf4\x91TR\x13r2\xc0:\x94-e\xf0\x19\xf2\x05\xa5\xe9\xa6\xde.\xd3G\xfa\xc6\xcc\xefW\x8c\xc0K\xb1\a\x06\x81i`JȈ\xdeB\vUh化\xc3pU/bŠ\xb2#\x8fZ%\xb3\x15إ_\x94\x9an\xb8\xe8\x02E\xd1c\xccִ'?+rn\xc8\xda7\xad\x8fa\xbb\x937\xdat\xa2Ւ)\xefB%\x1a\xf4x\x84ʠBN\x89P7>ݺ\xe0ԭ\v\"X\xa5\xf56\u007f%I\xa5c\xbb\xa9\xfde\x99\xab\xb4Ҧ\xb9\xe2eV_\x00K6\x99\xa9^\xc6]\xc6]\xad|\x92G\xe3\xf4\xe1 \xe9jԥ\xc8\xe2S\f\xb5ϒ\xbb\xd7xv\xe6;\xce/3\\\xba\xa7\x9b\xcb\x14\xab\xf4\xac\xe4\xe9\"\x1e\x164s\xef'\x9a\xb1\xc2y\xc0\x8f\xa9|\x064e\u007fX\x17\x15Q\xf2\x0f\xb4\x11\xce0/j=. \x1cۭL|\xdc\x1f\x94\x12\xc1\xff\x80Z\x16\xe9\v\xcdܬ\xe11\xdd܃q\xd6N\x98\xf1V>\x13\xba\xe5t\u007fED\xe7\xe2\x0fHi\x02wGM\xe1\xe7\u07b7\xe4\x9bvN7\xc0\xe0\x87F\xbe\x9b\xf71\xfe\x9d\x16Zud\xdc\xff\x9f\xa6\xef\x1f)>\f\xbf\x03\x00\x00\xff\xff\x8a\xa1[\xecg\a\x00\x00",
		hash:  "bccd721432c12ae01318a6e0f0da8884c8bd1cf7bb034f219564817448eb19b3",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  1895,
	},
	"settings.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xb4XO\x8fۺ\x11?{?Ŕ\b\xea\x04\x88\xd6IsYll\x01\x89\x9b\x00A\x8b\xf4\xb0{/\xc6\xe4\xd8\"L\x91\x02I\xd9\xeb\x1a\xfe\xee\x05\xff\xc8+\xdb\xf2\"o_\xdeI\xe2\xccp\xfeqf~\x94\xf6{AK\xa9\t\x98#\xef\xa5^9v8\xdc\xdc\xec\xf7o\x1a\\ѣ\xf4\x8a\xe0~\x06\xec\xa1\xc7\xdd\xef\xdf\xd4(\xf5\\\xa1s\x91\xe9N\x99\xbe\xa2:\xee\xba\xedv\xdd>\x06R\xe2\"\xf7rC?q\x13$>%\x9a\xd4K\x13\x96\xf5\xfa\x8b\xb5\xb8\x83\x9e\U0006cb77\xadg\xfch \xf9\xec\xa9n\x14z\x02ֽ=\x9a\x86AT\u007f8\xdcL\x85܀\x143\xb6j\xbdc\xe5\r\x00\xc0\xb4\xfaX~\x11\x1bԜ\x04|'\xf4\xad%7\x9dT\x1f3{il\x9d^\xe3\xb2\xd3@\xd6\x1a[\xfc\xcfhb\xc0\x83+3\xc6Q)\xd3z@Eփ\xa7'_pҞ,\x03\xe7w\x8afLH\xd7(\xdc\xddk\xa3\xe93+\xbfh\x88j\xa0B\a\x86\xf3֒\x98N\x84\xdc\f\xd8s-\xe7\xe4ܰ\xc5\xcc<\xb5\xb9\xdf\xcb%\xdc>$\xd6\xe1\x00\xfb=)G\x87Ð3A\x9a\xb48\x1c\xca.\x9f\xf0\x80\x9bao\xb2mk\xb6\xec\x99s\xce\xe5F\xb5\xb5vg\x12Q\xaa\xb9\xa4E\xba\xd4M\xeb\xc1\xef\x1a\x9a1^\x11_/\xcc\x13\x8b\xb1s#u\xc1\x8d\xf6\xd6(\x06\x1ak:\xa7mP\xb54c\u07b6t\f\xbc\xab\xbc\xb9\x91z\x9e\x04\x0f\a\x88\x9aIt\xe1\x0e\xbb\xa2pA\n\x96ƞ\xd9)\xbfi\\(\x82>\x11\u07ba\x86\xb8\\J\x0e\xc6B\xdd*/\x1bE\x90\xa2A!,9G.(\x03M[\xf0\x16\xb5\v\xa5l\xf4\xbb\xe9$\x1a\xba\xe2\xc4\xc2\xfe\xe1Dɺ1\xd6\x17\xf4\x14\x1e]\xa6Έ/\xa5\xeaG\x14\xfd\x16%_\x91\xabSK]\xb2p!\x95\xf4;\xf0\x06\x92\xc0$\t\xf4S\xe1~{*\x92\x8dbM;\xd7%\xe2\x84\xf4R\x1a\xfeE\xbbW\xe7\xa0od \x039\xf4\xc6\xcaM\x18SA췇.Ю\xe3\xb8\xed\x02O\x84\xb0~)\xea\u007fv\xdb^\x11\xf5\xb3\xc9.\xe6@\x818\xb9\u007f{|\xbcu\xde\xd4ߑ{S\x8b\xe3<8%^\x06\xcaMݠ\xa5\ao\xe3x{\x8e;\xef\xf9\xb7\xe1\x18J\x11\x982\x1cUe\x9c\xbf\xbf\xfbpw\xc7\xfa\xa33\xe7\x05~}|\x9c\xb8U\xce\xe3\x12\xf2\x1a:\x9bWS4\x9d\fL\xcb#&,\x93\x9aBe5q$\xa1\xd4a\xfa\xa79<\xb6f\vY\xecGL\xe8\x9f\xceE%\x05u\xf1\x8f\xc7W2\xd0C\x02W\xa3R\xc5\xc7\u007f@MB\xb6uq\aױ\xe1\x19#,\x95\xb9\x02\x86\x02e\xb9.\x02\xde\x1d\xc10\x8a\x17+kڦXJR\x82A\x8dO\x8a\xf4\xcaW3\xf6\xe9\x03\x83F!\xa7\xca(Av\xc6\xf6\xfb\xaba\x1f\x0e\xac\x9cN\x82\x0f\xc3ѝ\x82\xe2\v\xe4s\xf8\xfc\v\xd0t\xa8\u007f\xa6M9\xc5tk\xc0\r\x15\xbcB\xbd\"wLԢ\xf5\xdehV\x06\x84\x87ybN'X^\x14\xdbk\xdd\x1f\xbd\xec\xfb蔟\xef0[\xb4Z\xeaU\x96\x1b\x8dF\xa3iS~E\xbe\x96z\x05m\x03;\xd3ZآR䁣\x86\xc6\x1aO\xdc\a:,\xad\xa9A\x19\xe7\xc0,SsI\xe1\x00\xb5\x80o\xda\xdb\x1d\xcc-\t\xe9\xdd-\xa4\xb1\x1e4Fu\x8eH\xc0V*\x05\v\xe4\xeb`\x05\xf5\xae\a\xda+\xd2dѓH\x16|%]\xe7B\xd8\xfa\x1e*\xb3\xa5\r٤C\x1b?\xa4\a|\x85\x1e\xb6d\t~\xfe\xe7\xf1R'E]\xb7\xf0c\x19c\xa9©\x84\xed\vT\xe1R\xea\xc0\xe8\x9eK\xc1ș\x8a~f\x92[AO\x8d;آ\xf6\x01uj\xb3\xa18\x8c\xc3\x02\x8f\xea\x06\x15Eg\x8e\xa5\x90\xcfa\x8emh\x8c{H7\x84\x90\xc0\x18\xb2\xa5\xd8Si'o\xad%\x9d|\x80E8\xd2$\x12F\x12\xa0R\xbd(\xa4>\xf1\xfa\xed\xb6\x92\xbcJ\xf2\xbe\"\r\v\x02W\x99\xad\x06t\xbf\x10{0\xf8\xee\xc2\xe7\xdc\x02\x19\x90\x83\xccY\a\x80#n\xb4@\xbbce*\fx\b\xae{\x03K\xa9(\xb4\x04\x1c\x15\x8ez\xe3\xa8\xfb\xda)\xdaF\x19\x14$\x8a \x9f\x87\xd28\xbc\x8f\x13*\x8d[G\xf6{\\_\x1bRq\x9e\xbe4\n1c\xba\xc7p\xb1B-\x18ą7\xab\x95z\xbeޅ\xf8\nK\x1bBu\f3\xa9\x0eq\xf4\xc2\xf8[Q\x9c\x80\xebص\x8bZ\xfa\xa3\xc71\xa4\xff.\xbc\x1eg\x10ͤ\xeb!\xb0\x12\x8a\xa2o\x02O\xb3\x94=L9\xbaz\x02\xa9\xb2ҩv\xe9?\x99H\xa3\xd1\xd9\x00\x1a]\x9fO\xd3I\xfal\xbb\xb9\x18T9C\xbd\xdb\xf2I\xe6bj\xd3\"\xbdse\x1c\x15\x01[\x95\xe4\xeb\xeeF\x11Y\xa8e\x9dpW\xea\x80Q\x82\n\xa9/x\xa6\xf5\x99iZ\xdf;\xe8\xf0\xdd97z)m\x9d\xbb\xea\xf9\xa3s\xe0\xc33\xcf\xf5y\xf0\"\xb4\xdf8\xe7+T\xec8u\x8eِ\xddZ\xe9\a\xda\xf1\x16\x1e\xe3\xec\nb5\xae)\xb6c`\x90\xe8\xf5W\x1b\xafn\xb1\x01\x16\x14G\x19\th\x9b\xf7\xe0L\xda\xe5Z\x1b\x95\x1f\xc7\\4d\x94\xc8=OKc\xc3ב\xf6R\xb7R\xafn/!\xe5|\xbd(\u007f\xd26Fq?\x9d,r\xc3=_\x02\xaf\x1cT\x11\x91\xbf\x9cN\xe49b5\x03(\xfa\xab\xfd\x92\x8b\xb2C\xa2k\xb6y:\xb5c\xc9\x06\xe7\x87\xc0\xff\x15\xa6㯃\xeb\x86\x03\x1e(V\xce\xe3\xf3\xa2\xfa\xcf\xfeTd\x95\x1d\xcc\xc6:\xce\xd0\xdf+m@+\xb1\x88\xd7\xcf\x19\x9bGJ\xe7Z\xbacu\xb7\x85\xd3\xd8\\\x13`$l\xad\xa4\x10\xa4sc\x94\u007f\xf7\xb2&\xf7y:\t\x02}璚\xae;\xa3\xe7\xf91\xf4\xd7\xe6\xab\xf1\xde\xd4\fX\xfa\xa1\x14/\x9b7\xff\x0f\x00\x00\xff\xffy\xe7\x96\xf7\xaa\x12\x00\x00",
		hash:  "44afa43e7928c749b44f3a523b5303d74e0b458c3b85b89a0ff36a7183e33b2e",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  4778,
	},
	"templateBottom.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xecXMo\xe36\x10\xbd\xe7WL\x89E\xd7^@\x96\\\xa0(\x1aK\x02\xba\xbb\rZ\xa0@\v$\xeduA\x8b#\x9b\x0eE\xaa\xe4(vV\xf0\u007f/\xf4\xe5\xafȎ\x9d\xa6@\x0f\xd1I\xe6<>\xce\f\xdf(3)K\x81\xa9\xd4\b\x8c0\xcb\x15'\xfch\x88L\xc6\xd6\xeb+8xB\xdfaB\xd2\xe8\xb8ǔqٳ^\x96\xfe\x87\xb2\x94)$&˹\xc5[\xb2R\xcf\x1c\x8c\x80I-p\xc5\xd6\xeb\x0f~\xdfQ\xdfx\x1e\xdc\x18\v5\n\xfe\xe03\x04\xcf\xeb9W\xc8\aH\x14w.b\x16\x1f\x90+\x06RD\x8c,\xd7\xee3\x12\x97\xca1\x10\x9c\xb8ט\x9b\xf7D\x19\x87\x9e\xd1^\xa2dr_\xc1\vla\\ˌWAzRG,\xe5\x02=\xa9\x9f\xd8LA\xad\xd1\x14Ğ\xfaU\xfb6\x1f\xc7w\x95\x1f\xbcN\x1a\xb4\xee\x84\xfe||d\x03\xf1\xa9\xc2.\x1cI\x98ɯ(\x8e\xb07\x1b\xe6\xc8\xc5q{\x83\xb1\xa7\x01\rHl\xd3\xe6\x89\xdaQ\x8fV\xf2\xd4\xe1\xbbO\xf0/\x9f\xe7\x1d\xf4\xe9\xb98\xfdS\x81\x86\xfe3\xa9\nij\xc4\xe3+\xa5\xb2\xbbA\x9d\x17\xe4\xceLa\xe8\xc8\x1a=\x8b\u007f\xad7]\x87~\xfb;\x9c\xda3\t\x1a\xf5\x1c^\xa3\xf3.r\xe3\xec8\x0fb\x8eo~\xba\xfcڟ\xbfվ\x93\xc6\xd5V\xb8\xf9tw\xd9\xf6\xd3\x029@V\xa9\x8c\xffsU\xbeDT\xa6\xa0\x17\xa8\xea\x86'd\xa4\x80ߛݯ)\xaf\xcb\x1cz\xd3\xd7&\aݍ\x92!\xae\xd8\xe5~\xdeU\xfb\xdeJ\xe8\xf2\x12\xba$\xe1]\x01\xd5Ɇ\xee\xcf9\x8am\x05A\xe8r\xae\x9b¨@\x1em@,\x0eB\xbf\xb2\xc6U\xae\xffW\xb9\x88\xbb\xc0>sB\xb8EM\xfd\x115\xa5\xce\tY\x1c\xfc\xe8\ac\xff\xbb`\xfc=p\x82\xe0\x87\xeb \xb8\x0e\xba\x00_\xc5\xf7\xaa\xe9\xbb8\x80\x8f\xca$\xf7\xf0\v\xca\xd9|\x1bÎ\xf3\xd3\xca\xee\xcdk;\x8b!8\xd3\xd3\xdevs\aq\xbc]8\xa9\xfd0\x8fC^{\xa7\xa4\xbeg0\xb7\x98F\x8cuҜ\x16DF3 ngH\x11\xfb2U\\߳\xf8/\x89K\x90\x1a~^\xe5\xcaX\xb4\xa1\xcf7<&\xe1\xca{\x11\xdbo\xd5\xd6}N??\xe2w\xc3\xd517\rtǿ\xed\xa9\x81[\xc9=ŧ\xa8\"\xf6\xa9^\xe9\xdarz\xccq\xe3҉\xd4\xd6ګi\xe6R\b\xd4mw\x1e\u007fK2C7i\xf5v$\xf3\r}\xdf|\"\xe4C\xdc?f\xf4\xdet3\xb8\xa0\x16=\x03J\xcf\x1cԷT\x8f0r\x05\xa9\xb1p\x9bc\"S\x99\xc0\x9f\x0e\xad\xdb?2t\x89\x959\xc52\x85A\x95$\x93BfD\xa1\x10\xa2(\x82\xf7f\xba\xc0\x84\xde\x0f\xa1\\J-\xccr\xd4\x19[\xd4d\x83\x86B7Ü\x98\xacC\xbfe\xbd:<\a\x9cM\"65K\xb4_\xaa\xa9\xcch\xd4\xe4\xfc\xc5\xdf\x05\xdaG_HG\xed\xfb(\x93z\xb4p,\xder\xf5R-\x9c\x9f\x19\x81V˯\xf6\x1c\xfc\x93\xa3\x97sNM\x97\xba\xf3z\xee\xe9O\xd8RSh\xd1LiN\x12\xba&\xa4\xed\xea\x05a\xf1\x05_\x9d\v\xcd\xf3c\xc8\xfa\xab\xd6·\x05\xdb,8\xb6a(\xcb\xd1z\xbd\xcf\xd1'\x93S\x14\\\xa9\x03'\xfa\xd5x\xb6\x1c\xf7\x147\xdc\nmo}r\\k\xfbu\xf5\xae\xe5\x1b\x8e\x94\xe1b\x90\x16\xba.\x99\xc1\x10ʞ\xfa\x1b\xf5\xfc[\xe0\xdd@\x98\xa4\xc8P\xd3p\xb4\xbd\xd1\xc1p\xb2\x87\\\xef\xfc\u07bf\x8f\xd0o>ۡ?\xa7L\xc5Wm\x85\xff\x13\x00\x00\xff\xff\vr\xaf4\x05\x11\x00\x00",
		hash:  "88199435749c44cc9f84b95eb81170a2376727dcea647b26be761130ecb8ad1e",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  4357,
	},
	"templateTop.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xb4WO\x8f۶\x12?ǟb\x1e\xf1NAd\xad\x93\xe0!XH\x02^\x83\x04襗l\xdbc0&\xc7\x12\xbb\x14\xa9\x92#e\rC\u07fd\xa0\xfe\xd8ڍ\xedݶ\xc8\x1eV\x16\x87\xf3\xef\xf7\x1b͐\x87\x83\xa2\x9d\xb6\x04\x82\xa9n\f2ݹF\xf4\xfdju8\xa4\xaf\xe1=h\xabY\xa3\x81\xc0\xe8Y\xdb\x12:\xf4\x1a\xb7\x86\x020ޓ\x85\x9dw54\x18\x02)\xd0\x16\xd0{\xdc\x03;\b\xc4\xd06\xd0`I\xafӾ_\x1d\x0e\xff\x8d\xbf\xef4\x1b\x82\xdb\x1c\xb4U\xf4\x00k\xb8\x19e\\Q\xfdh}3\xae\xa3d\xdd\xd1/\xd8-eoa\x14֨\xedG\x83!,\x85\xef&a \x8e\x01?\x92\xbd\x8f\x99e\xffQN\xf2\xbe!\xa8\xb86\xc5*\x8b\x0f\x90\xd1N.\xacK\xfe\b\x02\f\xda2\x17dE\xb1\x02\x00\xc8*B5\xfe\x1c^kb\x04Y\xa1\x0fĹhy\x97|\x10O\xc5\x15s\x93П\xad\xeer\U00050d18HW7\xc8zkH\x80t\x96\xc9r.4\xe5\xa4J\xfaN\xdbbM\xb9\xe84}k\x9c\xe7\x85\xc27\xad\xb8\xca\x15uZR2\xbc\xbc\x99IJ\x82DC\xf9f}\xb34g\xb4\xbd\aO&\x17\xa1r\x9eeˠ\xa5\xb3\x02\"\x04\xb9\xd05\x96\x94>$\xe3Z\xe5i\x17\xd7\xca4\xbe\xa7;\xec\xe2s\xad\xa5[\x9a\xe4Hb\xf1\x19%\xbb\x1a>Y&\xdfx\x1d\b~Gc\x88\x87\xcaI`Ix\xdf\xc7\x1a\xc8\xd2Q\xf1ll\xbc7\x14*\"\x9e\x83\x90!\xa4\xd84k\x19\x82\xf8\x1b*\x8e+\xf2\v\xa5,=q\x97m\x9d\xda\xcfL\xcfE\xd7\xf7K\xf3\x81$kgA\xab\\\xec<֑\xabq\xbfw\xdf\x00\x8d.m\x12\xd8\x13\xcbj\xa1\xf6Hu\xdaoh\xc7\x1f\x9d\x81P\xa31\xc9\xe6-Ԥt['\xef@:\xd3\xd66<\xd1?\xd6\x19\xf9\xef\x05\x83P\xd7%\x04/G~v\x03\xfa\xc9εVa\xf4\xfb\xb5F\u007f\xbf\x0e]y\x8c8t%\x18W:\x01h8\x17\x13_\x9f\x8f\x1a\u209fjS\x9cH\xcd\xd2js&\xd0\xf4R\xa4\x99\xc5\xee\x82\xdd\xd6\f\xb0Z\xec\x12\xa3\x03_p?\xb1<le\x8f6\xe0\x00kH,v\x02\x0e\a\xbd\x83\xe1K\xf2\xf4\xb3\xe5\x00\x8b\x0e\xb1\x81\xbe\x9fs\x1fW\xe3~\xb2\xaa\xef\x8b\f\xa7\x1aI\x8f\xf0,\x8d\v\x88QI\xa3\xe5}\x14\xb4$\x8aL\x17\x8f\x11\xb7\xd8}]\xea<\xc5z\x82\xf9ni\xb6\xc8R],W\xb2\x14\x8b,5\xfa\xf9\xdcQ)O!$[\xe7\xee\x9f\xcf\xfd\xedKr\xff\xffh\xf2'\xe7\ue3c1/\xdd\f \x90\xa1\x9a,\xbf\x14\x91\xa5\xfe\x05D&\xb70\xf8\x1d\x10Y\xae\xbc\x1c\x91\xb9\xa3?\x8fƻ\x97\xa0\xf1e2w\x8ax\x9e\x18\xff\x00\x86Y\xf7\x02\x04G_C\xfa\xf3\xdb\xf5Գ\xb45羼\xb3\x1fX\xa6t7\x82\xb4\xb72\t\x8cܞ\x12\xab\xb4\xa2d\xe7|2\xf6\"g\xcd^\x14\xabW\xaf^e\xc3\xc2\xe2k\xbf\x85\xdfȇ\xd8ź\x9b\xf5f\xbdY\xdfd\xe9\xb8'\xdb\xfaI\xa5A{\xf2\xb4E\u007ftcpK\x06АgQ\x8c\xb3\x00\xbe쭼\x85\x93\x92q\xa8\x92\x86\xbc\x8c\xd8\x16\xd1z\x83\xb6\x98\x1eg\x92U\xfaI\xb2Y:\xf5\xd9\xeb\xdd\xd7\xeb\xb2:\xdb~\xe9\xa1A\xab\xae\xf5\xe0'\x96\xb6h\xd0J\n\x97\xbae\x04\xfe4\"\xae5\xb5\xc5\xce1\xaa\xff]\t\xe3\x05!\xc10\x03\xb4\n\xa0H~x\xc6\xc68A\x9eT\xeed\xe8\xebl\xe8B\xf5~\x9e\xc4c\xf5\xc6*\x18\xfeN\xacN\xfa\xc9dO\x147\xeb\x99\xd9\xe9\u007f\xcc\xf7\xdcք=j\xa3m9\x16\xc3Xj\x17\xaa\xe1\x05e\xf0L\xf1\xfcX:\xa4'\xa59\xfc\x1b\"&\x13\x17x\xf8d\xd9\xef\xe1\xe3\xecf \xe3YWG\x92H.\xf8\xf9\xf1\x18_\x10]1\xfa\xe8\xf0\x15\xbb\v\xa9D[\xa5%\xb2[t\x19\x87*^D$\x1a\xe3Z\x9e\xfa\xcd8\x0e\x8e\x87\xfe\xf5\x97A\xbd\xefa8%\xe6B\xe9\xd0\x18\xdc\xdfZg\xc7Y`\x02\xf5\xfdq&,ja\xac\x01`z\xe0$v)\xf2\xb1\x99y\xabmy\vw\xd5|\xca\x05\x1d \xb06\x06怸\"\xd8\x1a'\xefe\x85ڮ\xe1W\xcbڀ\x1ev\xeeZc\xf60\xe6\xf4\x06\xe6\x96\x02\xb1\x13\xed\x88\x02Ը\a\xeb\x18\xb6\x04\xd2yO\x92\xd7#~\xd7\xe0\x8aן\xc5i\xf6x\x1bzt\xa2=\v\xae\xda[\xac\xb5L\xa6[\x85(V\x13\x12\xab\xbf\x02\x00\x00\xff\xff\xe3\x19\xbd;\x11\x0e\x00\x00",
		hash:  "9e093fc8cadf04d051479f8ff070c0fb24943fd840de98c86621130f23bbab95",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  3601,
	},
	"transaction-template.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff̙͎\xdb6\x10\xc7\xcfݧ\x98\x12E/\x05c;I\xd14\xb1\x05l\x16I\xd1C\x83\x001z\xa7űM\x84\"\x05\x92Z\xafk\xe8\xdd\v\x89\x92,\xcb\xfa\xb2\xd7[4\x97\xac(r\x86\xa4\xfe\xf3\x9b!}8p\\\v\x85@\x9caʲ\xd0\t\xad\xa8\xc3(\x96\xcc!Iӻ\xbb\xc3\xe1'\xb62\xf0~\x01Bq|\x82W0MӬu\xb7\xd5\x12\xeb\xed3\xdf\xce\xc5\xe3G휎\xea\xef^\xfbw\xa1\x16*\xd4\xca\x19-\xebo\xdfd\x8e\xe6km\xa2\xe0\xee\x0e\x00`\xfe#\xa5\xf0\xd9\xe8\b(\r|\v\x17\x8f \xf8\x82d&ha\x83\xc0\xe1 \xd6P7\x9b\xa6p8\xa0\xb4\x98\xa6\x10Jf\xed\xf9\bT<M\xbd\xd5\xcar\xd1\xd5\xe8\x1d9\xbei\xbe\r\xb5L\"e\xc1FLJ:{\r\x11r\x91D\xf4-Hf6H\xdf6\xc6\xe6\xe3%[\xa1\f\xbe\xa1\xe2\xb06:z?\x9f\xf8\x96S/\x13.\x1e\xbb\x1d7\x1d\xbe+\x1c\xbeksx\xba\x1cp\xf8\xe4\xa8\x11\x9b\xadk\xe9<\xe4\xe8M\xb92(\xd6\xdea#\xb7\xc3\xf2\xef\xa3\x109r*T\x9c8\xbaJ\x9cӊ\x94\xe6\x8b\xc7\x00\xfe@\a_\xf2\x8e\xf0g\xd6q>a\x1d\x93;ߘ\xb6y{o\x1b\xa3\x93\xb8\xfc:ӳ\xcd\x1a\xb3\x86\xdcP\xbe\x0eo\xb2X\r\x8bt\xa2\x1c\x01\xb7\x8fqA\xb2-%-\xae\xe9Z\xa0\xe4\x04\f2\xae\x95\xdc/\x883\t\x12P,\xc2\x05Y#\x12xd2\xc1\x05\xa1\x94Ҿiؘ\xa96\a\xb9tH\xf0\xf9a9\x9fd}\x9e\xb5i~\xa3\xaao=+\xf6ivܧ\x0eCmj=mj>6B,\xdf\xe0\xccy\xbe6\xdb\x13r\x99\x82\xadP\x1b\x89\x85\xa2\xa6~p\xbdm \b\x9a\x9a\xfem\xa4\x1e:\xf4Ej\xeaX\xb3\xd0i\xc1)\xe3ܠ\xb59d\x98Ph\xfa\xbeml08\xd3Y\xc3ҩҼ~\xf2\x8e\xb3n\xddA\x8c&D\xe5~\xff\x15\xb8\xb0l%\xcb($\x10K\x16\xe2VK\x8efA\x1e\xb6Z[\x84\xc2#T\x1e\xcbA\xc1|\x92Mq(\u038bq+\xad\xbfWaΙc\xd4\xe9\xcdFz\xbd\xd3Z\xa7։\xfb\x81P\xfc\xd7\x12J\rOU\x00MI0\x17\xa5\xc55\x835\xa3\xf9\xfb`>\x11\xc1\xa54\xe9j>\x17Ѵ\x89\xfc7W\x8a(\xb8\xfb!\xfbw\x96a\x98\fa+8\x92`\xda\x13ŭ\xa8\xaa$t\x11\xab\xeaڢ*\x89\x1ab\xb9ύ\x81^\x97z\xb1\xff9\xb6F\u007f\x9d\x1a\xc9Fg+\x16Ǩ8U\xb8+c\xe5$U\x01>\xc5LeiJ\xe1\xee\xabLl\x8bF\x83\x9f\xd5\xca\xc6\x1fZEw\x11)\x8b?\x8f\x05\xd0R\x9f\x96?\xadEJK\x81\xd2\xe4i\xad\x04q\xfa\xbc\x009\x9b\x02\x1c\xa7p\xefc\x0f~\x81\xfb\xbf\x96\x9d\xb3\xa9x\xae\x13\xd7\x00z\a\xcc}G:;\xa1\xb9o$\xe3\xeb\xa01(\xef\xc5x1\x8dK9\xdedx\xbb\x996\x88\xfb\x9e\xa3(ވ\xc4\xe5>F(\xcb\xee4\xad\xb8݇뱨~1L\xcf.\xc6\xf4p\xc4\xdc\x02\xcd\x03X\xbe\x02\xc9=j\xb8\x82ǅLz\x81|\xd4Bg\xf2\x19\xe0\xb1?ڥ\xa9?m=,'\x9f\x1e\xb2\x10\xefb\xf4\xb5\x9ff4\x97ϙ\\\x10\xe1vP\x1e\xa4\xef\xf1\xf4\x89x\x03ޝ\xf0-;\x01\xbc\x00\xdc2\x0f\xb5b\xab\xb3(\x1e\x97-Z2G\x99\x04\x9c\x86\xafl\x9f\xedL\xfb)\xf6\x19\xa76R\xad\xe2\x165u\x8b\x9d\xdb¸\xa3\x86\xfeߕ\xceg\x86\xca\xf3\xe7\xec\x85\xeb\xe7\x97`\xb6W\\\xaf\xfa\xba\xb9>\xc0\xec\xfa\xf5W~P\x1fG\xeb\xf2\xcct\xd5I\xff\x19\xe5\xf2MQ|\x05 \xb7\xa6F\xcao\xc9\xcai\xc7\xe4\x85\xc5ꥼ;\xb2N\x18\xeb\xa8ul\x83E\x14\f\xad\xe2d\xb8\xc5P+\xde\x18\x0f\xd6\xed\xb3\xc8\xe3\xc2ƒ\xed\xdf+\xad\x9a\xac\xf6W\x8d\xd5\xedf\x9a\xdeu\x847>\xc5\xda8Z\x93T\xf3\x02\f\xb8\xde)\xa9Y\xd5\x19y\xbd\xfb+\xf7\xe4H\xf0)\u007f\x03\xcbc{{\xc1䝮\x8cf<d\xb6\xdfoc\x95\x90-\xf3\x03\t>\x96\x83{\xbd\x95\xf7\xaa]S\xb0Y\xd6F\xe5\x84\xc1\xbeI\xf8\x83Ȁ#\xc5{\xf6\x97\x8b\xdeU\x06\x9f\xb8\xe8_I\x9f<\xd0\x18m\xe8?\xd9\xf7\xaf\x95\x80R'\x0e\x98D\xe3\xfc\x85j\x96\x15д\xca\xe6\x03\t\xee\x15\xe4f`\xcb,\xe80L\f\xf2^I&a\x98\xe5\xbaV\xaf\xc5\xcbA\xbf\xad\xd7u\xf3v`\x97W\xb5\xbb\x9e}\xacʬ\x0eo_pw\xc9&\xf7\\\tvݧW\xd9\x00P\xf1\xf6C\xed2\x03Oǝ\xfa\xe8\\Б\ar\xa8=/\x13\xd8\x02\x8dE\x19X=\xf5'\x87+\x13\xc3(twS\xac\xfe\xd3M\xd1|\x12\xf1\xb7\xa3\xfa\x05\x9a\x18]!\x14\x92\x8e\xd8w\x1cfPM\xda\xc7\xfbĿE\x8f\xa0\ag:x\xc8\x19\xfc.\x9ey\xf3\x89\xff\x19\xac|\xfe7\x00\x00\xff\xff\xfc\x97\x94^\xa2\x1b\x00\x00",
		hash:  "2087ed1b9a24d7b92ec728dd8d9e8e96e613ec77bee9ffa59cbab9cc422f48af",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1484602929, 0),
		size:  7074,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f, ok := staticFiles[strings.TrimPrefix(req.URL.Path, "/")]
	if !ok {
		NotFound(rw, req)
		return
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	} else {
		return gzip.NewReader(strings.NewReader(f.data))
	}
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}

// Slower than Open as it must cycle through every element in map. Open all files that match glob.
func OpenGlob(name string) ([]io.ReadCloser, error) {
	readers := make([]io.ReadCloser, 0)
	for file := range staticFiles {
		matches, err := path.Match(name, file)
		if err != nil {
			continue
		}
		if matches {
			reader, err := Open(file)
			if err == nil && reader != nil {
				readers = append(readers, reader)
			}
		}
	}
	if len(readers) == 0 {
		return nil, fmt.Errorf("No assets found that match.")
	}
	return readers, nil
}
