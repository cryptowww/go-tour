# a cobra application

## create the project

```shell

mkdir pdfprocess

cd pdfprocess

go mod init aviva-cofco/pdf

cobra init

```
## test the project

```shell

go run main.go

```

## add subcommands

```shell
# a subcommand to split pdf

cobra add split

# a subcommand to merge pdfs

cobra add merge

```

## add icon for application

ref :[add icon](https://www.jianshu.com/p/77092cfa02a0)


## build

go build -o pdfp.exe
