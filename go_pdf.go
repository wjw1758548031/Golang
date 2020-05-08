package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	"log"
	"os"
	//pdfi "github.com/wjw1758548031/go_pdf"
	// "github.com/unidoc/unidoc/pdf/model"
	//"github.com/unidoc/unipdf"
	//pdf "github.com/wjw1758548031/go-pdf/unipdf"
	//可以去对应的github.com/unidoc/unipdf是否有跟新
	"github.com/wjw1758548031/go_pdf/extractor" //这里考入的包是自己导入的
	pdf "github.com/wjw1758548031/go_pdf/model" //这里考入的包是自己导入的
	//"github.com/unidoc/unipdf/extractor"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	//自己去下载一个字体，可以显示中文的，就可以了
	err := pdf.AddTTFFont("wts11", "D:/goprj/src/test/test1/t/HYXuJingXingKaiW.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		panic(err.Error())
	}
	//里面还有很多的方法，比如灰色之类的点进去去看有什么设置，不做过得的介绍.
	pdf.SetX(100) //设置位置
	pdf.SetY(100)
	pdf.SetTextColor(100, 100, 0) //文字色調
	pdf.Text("wwwwwww:www.baidu.com")
	pdf.Cell(nil, "1111ssss我闭目在经殿的香雾中， 蓦然听见你颂经中的真言；")

	//htmlStr := newHtml() //new一个html模板

	pdf.WritePdf("hello1.pdf")

	err = outputPdfText("D:/goprj/src/test/hello1.pdf")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}

// outputPdfText 输出pdf文件内容到终端
func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	fmt.Printf("Total Pages:%d\n", numPages)

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Println("\"%s\"\n", text)
	}

	return nil
}

//pdf "github.com/wjw1758548031/go_pdf/model" 引用得包

//Golang 切割pdf 文件并取出前三页进行存储
/*func pdfSplit(){
	err := splitPdf("D:/goprj/src/test/hello1.pdf", "D:/goprj/src/test", 1, 1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func splitPdf(inputPath string, outputPath string, pageFrom int, pageTo int) error {
	pdfWriter := pdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	if numPages < pageTo {
		return err
	}

	for i := pageFrom; i <= pageTo; i++ {
		pageNum := i

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			return err
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}*/
