package controllers

import (
	"bufio"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize"
	_ "image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type WenJianLiu struct {
	ApiController
}

/////////////////重点////////////
///https://xuri.me/excelize/zh-hans/cell.html#SetCellStyle  这是对exce进行详细操作，比如样式图表等

// @Description wenjianliu 读文件*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliutxtt [post]
func (this *WenJianLiu) Wenjianliutxtt() {
	//response.setHeader("Access-Control-Allow-Origin", "*");
	fmt.Println("进入文件流1")
	b, err := ioutil.ReadFile("D:/goprj/src/scs/controllers/wenjianliu.txt")
	//b, err := ioutil.ReadFile("../scs/controllers/wenjianliu.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// @Description wenjianliu 读文件*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliutxtOpen [post]
func (this *WenJianLiu) WenjianliutxtOpen() {
	fmt.Println("进入文件流2")
	f, err := os.Open("../scs/controllers/wenjianliu.txt")
	//f,err := os.Open("D:/goprj/src/scs/controllers/wenjianliu.txt")
	check(err)
	b1 := make([]byte, 100)
	//读取的数据写入b1里面,n1为存储的具体数据，，b1为[]byte类型数据
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//第一个参数是从什么位置开始，第二个参数未知，最后整个f都变了值
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 100)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(4, 0)
	check(err)
	b3 := make([]byte, 100)
	n3, err := io.ReadAtLeast(f, b3, 1)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	//从第0个位置到第五个位置
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}

// @Description wenjianliu 写文件*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuCreate [post]
func (this *WenJianLiu) WenjianliuCreate() {
	//response.setHeader("Access-Control-Allow-Origin", "*");
	//如果已有文件，则默认打开文件
	fmt.Println("进入文件流2")
	b, err := ioutil.ReadFile("D:/goprj/src/scs/controllers/wenjianliu.txt")
	//b, err := ioutil.ReadFile("../scs/controllers/wenjianliu.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}

// @Description wenjianliu 写文件*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuWrite [post]
func (this *WenJianLiu) WenjianliuWrite() {
	//response.setHeader("Access-Control-Allow-Origin", "*");
	fmt.Println("进入文件流3")
	d1 := []byte(" 往后增加")
	//D:/goprj/src/scs/controllers/wenjianliu.txt
	//向文件中写数据,如果文件不存在,将以 perm 权限创建文件。
	err := ioutil.WriteFile("../scs/controllers/wenjianliu.txt", d1, 0644)
	check(err)
}

// @Description wenjianliu c创建*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuWriteCreateA [post]
func (this *WenJianLiu) WenjianliuWriteCreateA() {
	//response.setHeader("Access-Control-Allow-Origin", "*");
	//如果文件存在，默认打开
	f, err := os.Create("../scs/controllers/wenjianliuCreate.txt")
	check(err)
	defer f.Close()
	//d2 := []byte{115, 111, 109, 101, 10}
	d2 := []byte("王建文")
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
}

//-------------------------以上都是文件,一下是exce

// @Description wenjianliu c查询exce*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuexce [post]
func (this *WenJianLiu) Wenjianliuexce() {
	fmt.Println("-----------wenjianliuexce--------")
	excelFileName := "D:/goprj/src/scs/controllers/test_write.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	fmt.Println("-----------zzzzzzzzzzzz--------")

	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
	fmt.Println("----------结束---------")
}

// @Description wenjianliu 创建exce*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuCreatexce [post]
func (this *WenJianLiu) WenjianliuCreatexce() {
	fmt.Println("-----------wenjianliuexce--------")
	//如果新增的路径文件已有，则会覆盖

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "姓名1"
	cell = row.AddCell()
	cell.Value = "年龄"

	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"

	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"

	err = file.Save("D:/goprj/src/scs/controllers/test_write.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("----------结束---------")
}

// @Description Filepath 获取执行文件在哪个位置执行*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /filepath [post]
func (this *WenJianLiu) Filepath() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	//获取执行文件在哪个位置执行
	fmt.Println(dir)
	fmt.Println("-----------zzzzzzzzzzzz--------")
}

// @Description wenjianliu 修改exce*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuUpdateexce [post]
func (this *WenJianLiu) WenjianliuUpdateexce() {
	fmt.Println("-----------wenjianliuexce--------")
	//Save可以当做替换，如果目录上有，则会替换，所以可以当做修改
	excelFileName := "D:/goprj/src/scs/controllers/test_write.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	first := xlFile.Sheets[0]
	//又增加了一行和列
	row := first.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "铁锤"
	cell = row.AddCell()
	cell.Value = "99"

	//指定修改第几行第几列，当然也可能指定第几行，新增列，都可以的
	row1 := first.Rows[1]
	row1.SetHeightCM(1)
	cell1 := row1.Cells[0]
	cell1.Value = "铁锤12"
	cell1 = row.Cells[1]
	cell1.Value = "9912"

	err = xlFile.Save(excelFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("----------结束---------")
}

//---------------------------以下是用excelize对exce进行操作---------------------

// @Description wenjianliu 使用excelize读取exce*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuExcelizeReadexce [post]
func (this *WenJianLiu) WenjianliuExcelizeReadexce() {
	fmt.Println("-----------wenjianliuexce--------")

	xlsx, err := excelize.OpenFile("D:/goprj/src/scs/controllers/test_write.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	//定位，直接定位到B2的位置，获取数据
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	//获取表格
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	fmt.Println("----------结束---------")
}

// @Description wenjianliu 使用excelize创建exce*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuExcelizeCreatexce [post]
func (this *WenJianLiu) WenjianliuExcelizeCreatexce() {
	fmt.Println("-----------wenjianliuexce--------")
	//如果创建的exce在该目录已经有了，则会进行覆盖
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	//定位到表格的对应位置进行增加
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "年龄")
	xlsx.SetCellValue("Sheet1", "A2", "狗子")
	xlsx.SetCellValue("Sheet1", "B2", "18888")
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("D:/goprj/src/scs/controllers/test_write2.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------结束---------")
}

// @Description wenjianliu 使用excelize插入图表*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuExcelizeAddIconexce [post]
func (this *WenJianLiu) WenjianliuExcelizeAddIconexce() {
	fmt.Println("-----------wenjianliuexce--------")
	//如果目录有该文件，则会进行覆盖
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	xlsx := excelize.NewFile()
	//Sheet1表格中的k插入什么值
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	/* type 图形类型
		area 未知   空白             | 2D area chart
	 areaStacked  未知   空白        | 2D stacked area chart
	 areaPercentStacked   | 2D 100% stacked area chart
	 area3D               | 3D area chart
	 area3DStacked        | 3D stacked area chart
	 area3DPercentStacked | 3D 100% stacked area chart
	 bar                  | 2D clustered bar chart			横行图
	 barStacked           | 2D stacked bar chart
	 barPercentStacked    | 2D 100% stacked bar chart
	 bar3DClustered       | 3D clustered bar chart
	 bar3DStacked         | 3D stacked bar chart
	 bar3DPercentStacked  | 3D 100% stacked bar chart
	 col                  | 2D clustered column chart
	 colStacked           | 2D stacked column chart
	 colPercentStacked    | 2D 100% stacked column chart
	 col3DClustered       | 3D clustered column chart		树状图
	 col3D                | 3D column chart					树状图
	 col3DStacked         | 3D stacked column chart
	 col3DPercentStacked  | 3D 100% stacked column chart
	 doughnut             | doughnut chart					圆形图
	 line                 | line chart						圆形图
	 pie                  | pie chart
	 pie3D                | 3D pie chart					圆形图
	 radar                | radar chart						一个箭头上的图
	 scatter              | scatter chart					·点状图
	*/

	//插入表格 type":"col3DClustered"必填 name 名称 values值（树状图有多高）
	xlsx.AddChart("Sheet1", "E1", `{"type":"col3DPercentStacked","series":[{"name":"111","categories":"222","values":"333"},{"name":"444","categories":"555","values":"666"},{"name":"777","categories":"888","values":"999"}],"title":{"name":"主题！"}}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("D:/goprj/src/scs/controllers/test_write2.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------结束---------")
}

// @Description wenjianliu 使用excelize插入图片*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuExcelizeInsertImageexce [post]
func (this *WenJianLiu) WenjianliuExcelizeInsertImageexce() {
	fmt.Println("-----------wenjianliuexce--------")

	//必须要读取到该文件，否则错误
	xlsx, err := excelize.OpenFile("D:/goprj/src/scs/controllers/test_write3.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	err = xlsx.AddPicture("Sheet1", "A2", "D:/goprj/src/scs/controllers/1547449232(1).png", "")
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling. x_scale 本图片的宽占百分比  y_scale高
	err = xlsx.AddPicture("Sheet1", "D2", "D:/goprj/src/scs/controllers/1547449236(1).png", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	err = xlsx.AddPicture("Sheet1", "H2", "D:/goprj/src/scs/controllers/1547449239(1).png", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
	if err != nil {
		fmt.Println(err)
	}
	// Save the xlsx file with the origin path.
	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------结束---------")
}

// @Description wenjianliu 使用excelize读取exce并修改其中字体的样式*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /wenjianliuExcelizeReadFontStyle [post]
func (this *WenJianLiu) WenjianliuExcelizeReadFontStyle() {
	fmt.Println("-----------wenjianliuexce--------")

	xlsx, err := excelize.OpenFile("D:/goprj/src/scs/controllers/test_write.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置好的样式
	style, err := xlsx.NewStyle(`{"font":{"bold":true,"italic":true,"family":"Berlin Sans FB Demi","size":20,"color":"#777777"}}`)
	if err != nil {
		fmt.Println(err)
	}
	// 从A3到B3全部变成这种样式
	xlsx.SetCellStyle("Sheet1", "A3", "B3", style)
	xlsx.Save()
	fmt.Println("----------结束---------")
}
