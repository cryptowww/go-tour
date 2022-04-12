/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"github.com/unidoc/unipdf/v3/model"
)

var (
	infile string
	outdir string
	breakpage  []string
)

// pdf split -i d:/a.pdf -o d:/ -b "1-2,3-4,5-6"
// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "split a PDF file into more by page",
	Long: `You can use split command to split a PDF like this:
split
	-i	PDF文件的完整路径，必须

	-o	默认为-i文件同目录，也可以指定目录
	
	-b	pdf文件拆分方式，支持以下方式：

		1. 每几页拆分，如每3页、每1页、每5页，传参方式为3b、1b、5b

		2. 指定需要拆出第几页，如第3页、第1页、第5页，传参方式为3、1、5
		
		3. 拆分多个页码范围，如1-5页、5-7页、8-11页，传参方式为1-5,5-7,8-11，注意逗号中间不要有空格等

例：
	$ pdf.exe split -i d:\a\b\c.pdf -b 3b

	输出：d:\a\b\c\目录下存放拆分文件，每3页一个pdf

	$ pdf.exe split -i d:\a\b\c.pdf -b 3

	输出：d:\a\b\c\目录下存放拆分文件，第3页被单独拆分为一个pdf


	$ pdf.exe split -i d:\a\b\c.pdf -b 1-3,5-9

	输出：d:\a\b\c\目录下存放拆分文件，1-3页、5-9页各拆分两个pdf文件

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("split Done!")
		//fmt.Println(infile)
		//fmt.Println(outdir)
		//fmt.Println(breakpage)
		//for _, p := range breakpage {
		//		fmt.Println(p)
		//}

		split(infile, outdir, breakpage)
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// splitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// splitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	splitCmd.Flags().StringVarP(&infile, "infile", "i", "", "the full path of the PDF to split(required)")
	splitCmd.Flags().StringVarP(&outdir, "outdir", "o", "", "the dir to save the splited PDF files(default as the infile directory)")
	splitCmd.Flags().StringSliceVarP(&breakpage, "breakpage", "b", nil, "define the type of breaking pages(required)")
	splitCmd.MarkFlagRequired("infile")
	splitCmd.MarkFlagRequired("breakpage")
}
