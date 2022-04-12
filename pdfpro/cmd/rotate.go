/*
Copyright © 2022 James <jameschuh@126.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"github.com/spf13/cobra"
)

// rotateCmd represents the rotate command
var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "A brief description of your command",
	Long: `You can rotate a pdf like this:
	
	rotate a.pdf 90

	例子：

	$ pdf.exe rotate d:\a\b\c.pdf 90

	输出：d:\a\b\c-rotate.pdf
.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rotate Done")

		if len(args) != 2 {
			fmt.Println("请输入正确的参数")
			os.Exit(1)
		}

		_, err := os.Stat(args[0])
		if err != nil || os.IsNotExist(err) {
			fmt.Printf("文件：%s 不存在\n",args[0])
			return
		}
		if !strings.HasSuffix(args[0], ".pdf") {
			//|| !strings.HasSuffix(arg, ".pdf") {
			fmt.Println("请输入pdf文件")
			return
		}


		degrees, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Printf("Invalid degrees: %v\n", err)
			os.Exit(1)
		}
		if degrees%90 != 0 {
			fmt.Printf("Degrees needs to be a multiple of 90\n")
			os.Exit(1)
		}

		rotate(args[0], degrees)
	},
}

func init() {
	rootCmd.AddCommand(rotateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rotateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rotateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
