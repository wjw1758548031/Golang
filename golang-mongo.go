package main

import (
	"bytes"
	"encoding/xml"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/mongodb/mongo-go-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"net/url"
	"reflect"
	//"reflect"
	//"bytes"
	"crypto/tls"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	//"strconv"
	"context"
	"strings"
	"time"
)

// HTTPSClient HTTPS客户端结构
type HTTPSClient struct {
	http.Client
}

type Dishes struct {
	DishesId   string
	DishesName string
	Count      int64
	Price      float64
}

var client http.Client

func init() {
	config := new(tls.Config)
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		TLSClientConfig: config,
	}
	client = http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"CMBSDKPGK"`
	INFO    []INFO   `xml:"INFO"`
}

type INFO struct {
	XMLName xml.Name `xml:"INFO"`
	DATTYP  string   `xml:"DATTYP"`
}

var client1 *mongo.Client

func main() {
	Init()
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	router := app.Party("/", crs).AllowMethods(iris.MethodOptions)
	router = app.Party("/api/v1")
	router.Post("/insert", InsertAndSelect)

	app.Run(iris.Addr(":" + "4015"))

}

var i int

func InsertAndSelect(ctx iris.Context) {
	i += 1
	data1, err := client1.Database("test").Collection("test").InsertOne(context.Background(), bson.M{"SS1": i, "SS2": i, "SS3": i})
	if err != nil {
		fmt.Println("err1", err)
	}
	fmt.Println("data1:", data1)
	var collOpts []*options.CollectionOptions
	//Secondary  读操作只在从节点
	collOpts = []*options.CollectionOptions{&options.CollectionOptions{ReadPreference: readpref.Secondary()}}
	data, err := client1.Database("test").Collection("test", collOpts...).Aggregate(context.Background(), []bson.M{})
	if err != nil {
		fmt.Println("err2", err)
	}
	if err = data.Err(); err != nil {
		fmt.Println("err3", data.Err())
	}
	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		var doc interface{}
		if err = data.Decode(&doc); err != nil {
			return
		}
		fmt.Println("doc:", doc)
	}
}

func Init() {
	i = 0
	dbUrl, err := url.Parse("mongodb://47.101.201.47/test?authSource=admin")
	if err != nil {
		panic(err)
	}
	//fmt.Println(dbUrl)
	dbname := dbUrl.Path[1:]
	fmt.Println("当前数据库", dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@47.101.201.47:27018,47.101.201.47:27019,47.101.201.47:27020/test?replicaSet=mmm"))
	if err != nil {
		panic(err)
	}
	client1 = client
}

func Ceshi() {
	str := `
<?xml version="1.0" encoding="GBK"?><CMBSDKPGK><INFO><DATTYP>2</DATTYP><ERRMSG></ERRMSG><FUNNAM>DCPAYMNT</FUNNAM><LGNNAM>银企直连测试用户98</LGNNAM><RETCOD>0</RETCOD></INFO><NTQPAYRQZ><ERRCOD>NCB4241</ERRCOD><ERRTXT>NCB4241 -业务参考号重复</ERRTXT><REQSTS>FIN</REQSTS><RTNFLG>F</RTNFLG><SQRNBR>0000000000</SQRNBR><YURREF>APP060928001255</YURREF></NTQPAYRQZ></CMBSDKPGK>
	`

	decoder := xml.NewDecoder(bytes.NewReader([]byte(str)))
	decoder.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		return charset.NewReaderLabel(c, i)
	}
	result := &Recurlyservers{}
	decoder.Decode(result)

	//xml.Unmarshal([]byte(str), &result)
	fmt.Println("xmlRe:", result)
}

//支付
func Pay() {
	data := make(map[string]interface{}, 0)
	INFO := make(map[string]interface{}, 0)
	INFO["FUNNAM"] = "DCPAYMNT"
	INFO["DATTYP"] = "2"
	INFO["LGNNAM"] = "银企直连测试用户98"
	data["INFO"] = INFO
	SDKPAYRQX := make(map[string]interface{}, 0)
	SDKPAYRQX["BUSCOD"] = "N02031"
	data["SDKPAYRQX"] = SDKPAYRQX
	DCOPDPAYX := make(map[string]interface{}, 0)
	DCOPDPAYX["YURREF"] = "APP060928001255"
	DCOPDPAYX["DBTACC"] = "755915677710908"
	DCOPDPAYX["DBTBBK"] = "75"
	DCOPDPAYX["TRSAMT"] = "1.01"
	DCOPDPAYX["CCYNBR"] = "10"
	DCOPDPAYX["STLCHN"] = "N"
	DCOPDPAYX["NUSAGE"] = "测试"
	DCOPDPAYX["BNKFLG"] = "Y"
	DCOPDPAYX["CRTACC"] = "6225885910000108"
	DCOPDPAYX["CRTNAM"] = "6225885910000108"
	DCOPDPAYX["CRTBNK"] = "招商银行"
	data["DCOPDPAYX"] = DCOPDPAYX
	xmlStr := MapToXml(data, "CMBSDKPGK")

	str, err := Utf8ToGbk([]byte(xmlStr))
	if err != nil {
		fmt.Println("err:", err)
	}
	xmlStr = string(str)
	fmt.Println("xmlStr:", xmlStr)

	re, err := client.Post("http://192.168.88.98:8089", "application/x-www-form-urlencoded", strings.NewReader(xmlStr))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	content, err := ioutil.ReadAll(re.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(content)
	strs := string(decodeBytes)
	fmt.Println("strs:", string(strs))
}

//查询余额 2.2
func SelectBalance() {
	data := make(map[string]interface{}, 0)
	INFO := make(map[string]interface{}, 0)
	INFO["FUNNAM"] = "GetAccInfo"
	INFO["DATTYP"] = "2"
	INFO["LGNNAM"] = "银企直连测试用户98"
	data["INFO"] = INFO
	SDKACINFX := make(map[string]interface{}, 0)
	SDKACINFX["BBKNBR"] = "75"
	SDKACINFX["ACCNBR"] = "755915677710908"
	data["SDKACINFX"] = SDKACINFX
	xmlStr := MapToXml(data, "CMBSDKPGK")

	/*xmlStr := `<?xml version="1.0" encoding = "GBK"?>
	<CMBSDKPGK>
	<INFO>
	<FUNNAM>GetAccInfo</FUNNAM>
	<DATTYP>2</DATTYP>
	<LGNNAM>银企直连测试用户98</LGNNAM>
	</INFO>
	<SDKACINFX>
	<BBKNBR>75</BBKNBR>
	<ACCNBR>755915677710908</ACCNBR>
	</SDKACINFX>
	</CMBSDKPGK>`*/

	str, err := Utf8ToGbk([]byte(xmlStr))
	if err != nil {
		fmt.Println("err:", err)
	}
	xmlStr = string(str)
	fmt.Println("xmlStr:", xmlStr)

	re, err := client.Post("http://192.168.88.98:8089", "application/x-www-form-urlencoded", strings.NewReader(xmlStr))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	content, err := ioutil.ReadAll(re.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(content)
	strs := string(decodeBytes)
	fmt.Println("strs:", string(strs))
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func MapToXml(data map[string]interface{}, sdk string) string {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding = "GBK"?>`))
	buf.WriteString(fmt.Sprintf("\n<%s>", sdk))
	for k, v := range data {
		str := bytes.NewBufferString("")
		switch reflect.TypeOf(v).String() {
		case "map[string]interface {}":
			for kOne, vOne := range v.(map[string]interface{}) {
				str.WriteString(fmt.Sprintf("\n<%s>%s</%s>", kOne, vOne, kOne))
			}
		}
		buf.WriteString(fmt.Sprintf("\n<%s>%s\n</%s>", k, fmt.Sprintf("%s", str), k))
	}
	buf.WriteString(fmt.Sprintf("\n</%s>", sdk))
	return fmt.Sprintf("%s", buf)
}
