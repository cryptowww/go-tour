package cmd

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/model"
	"path/filepath"
	"strings"
	"os"
	"strconv"
)

func loadLicense(){
	err := license.SetMeteredKey("ba6324ae854ba42f199aa841056d67abcb9d80c980421041f9935319ba2b0c74")
	if err != nil {
		panic(err)
	}
	//fmt.Println("License loaded successfull.")
}


func split(infile string, outdir string, breakpage []string) {

	//fmt.Println(infile)
	//fmt.Println(outdir)
	//fmt.Println(breakpage)

	inPath := filepath.Dir(infile)
	//fmt.Println(inPath)

	inName := filepath.Base(infile)
	//fmt.Println(inName)

	inFileName := strings.Split(inName,".")[0]
	//fmt.Println(inFileName)

	// 新建子目录
	if outdir == "" {
		outdir = filepath.Join(inPath,inFileName)
	}
	//fmt.Println(outdir)
	
	if err := os.MkdirAll(outdir, 0666); err != nil {
		fmt.Println("mkdir err")
		return
	}
	
	// 获取PDF页数
	pdfReader, f, err := model.NewPdfReaderFromFile(infile, nil)
	if err != nil {
		return 
	}
	defer f.Close()

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return
	}

	// 循环拆分
	for i, str := range breakpage {
		//fmt.Printf("分页，第%d, %s", i, str)

		// 多个分页
		if strings.Index(str, "-") != -1 {
			idxs := strings.Split(str, "-")
			star,err := strconv.Atoi(idxs[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			end,err := strconv.Atoi(idxs[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			if star > end {
				fmt.Println("页码错误")
				return
			}
			splitPdf(infile, filepath.Join(outdir, inFileName+"-"+str+".pdf"), star, end)
		} else {
			if len(breakpage) >1 {
				nump, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(err)
					return
				}
				splitPdf(infile, filepath.Join(outdir, inFileName+"-"+str+".pdf"), nump, nump)
			} else if len(breakpage) == 1 {
				// 每几页拆分一个，如1b，2b，3b
				if strings.Index(str,"b") != -1 {
					idxs := strings.Split(str, "b")
					nump, err := strconv.Atoi(idxs[0])
					if err != nil {
						fmt.Println(err)
						return
					}
					// 每几页
					if nump == 1{
						for n := 1; n <= numPages; i++ {
							splitPdf(infile, filepath.Join(outdir, inFileName+"-"+strconv.Itoa(n)+".pdf"), n, n)
						}
					} else {
						for n := 1; n <= (numPages/nump + 1); n++ {
							start := nump * (n - 1) + 1
							end := nump * n
							if end > numPages {
								end = numPages
							}
							if start > end {
								return
							}
							splitPdf(infile, filepath.Join(outdir, inFileName+"-"+strconv.Itoa(start)+"-"+strconv.Itoa(end)+".pdf"), start, end)
						}
					}
				} else {
					nump, err := strconv.Atoi(str)
					if err != nil {
						fmt.Println(err)
						return
					}
					if nump > numPages {
						fmt.Println("页码超届")
						return
					}
					splitPdf(infile, filepath.Join(outdir, inFileName+"-"+str+".pdf"), nump, nump)

				}
			}
		}

	}
}

// 实际拆分函数
func splitPdf(infile string, outfile string, pageFrom int, pageTo int) error {
	pdfWriter := model.NewPdfWriter()
	pdfReader, f, err := model.NewPdfReaderFromFile(infile, nil)
	if err != nil {
		return err
	}
	defer f.Close()

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	// 如果超页，则取最大页
	if numPages < pageTo {
		//return err
		pageTo = numPages
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

	fWrite, err := os.Create(outfile)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}

// 合并pdf文件
func merge(inFiles []string) {

	outFile := strings.Replace(inFiles[0], ".pdf", "-combine.pdf", -1)
	err := mergePdf(inFiles, outFile)
	if err != nil {
		fmt.Println(err)
	}
}

func mergePdf(inputPaths []string, outputPath string) error {
	pdfWriter := model.NewPdfWriter()

	for _, inputPath := range inputPaths {
		pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
		if err != nil {
			return err
		}
		defer f.Close()

		numPages, err := pdfReader.GetNumPages()
		if err != nil {
			return err
		}

		for i := 0; i < numPages; i++ {
			pageNum := i + 1

			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				return err
			}

			err = pdfWriter.AddPage(page)
			if err != nil {
				return err
			}
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
}


func rotate(inFile string, degrees int64) {
	outFile := strings.Replace(inFile, ".pdf", "-rotate.pdf", -1)
	err := rotatePdf(inFile, degrees, outFile)
	if err != nil {
		fmt.Println(err)
	}
}

// Rotate all pages by degrees.
func rotatePdf(inputPath string, degrees int64, outputPath string) error {
	pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfWriter, err := pdfReader.ToWriter(&model.ReaderToWriterOpts{})
	if err != nil {
		return nil
	}

	// Rotate all page degrees.
	err = pdfWriter.SetRotation(degrees)
	if err != nil {
		return nil
	}

	pdfWriter.WriteToFile(outputPath)

	return err
}
