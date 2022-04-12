/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"strings"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge more PDF files into one",
	Long: `You can merge more than one PDF file into one.
	
	merge pdf1 pdf2 pdf3

	例子：
	
	pdf.exe merge d:\a\b\c.pdf c:\e\f.pdf

	输出：合并后的文件为d:\a\b\c-combine.pdf
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("merge Done!")
		if len(args) < 2 {
			fmt.Println("Please input more than one pdf file.")
			return
		}
		
		var inFiles []string = []string{}

		// 拼接参数
		for _, arg := range args {
			//fmt.Printf("%d, %s \n",i,arg)
			
			_, err := os.Stat(arg)
			if err != nil || os.IsNotExist(err) {
				fmt.Printf("文件：%s 不存在\n",arg)
				return
			}
			if !strings.HasSuffix(arg, ".pdf") {
				//|| !strings.HasSuffix(arg, ".pdf") {
				fmt.Println("请输入pdf文件")
				return
			}
			inFiles = append(inFiles, arg)
		}
		merge(inFiles)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	
//	loadLicense()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
