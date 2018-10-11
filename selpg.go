package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"os/exec"

	flag "github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage int
	endPage int
	inputFileName string
	printDst string
	pageLength int
	delimited bool
}

var progname string

func initArgs(args *selpgArgs) {
	flag.Usage = Usage
	flag.IntVarP(&args.startPage, "startPage", "s", -1, "sepcify start page")
	flag.IntVarP(&args.endPage,"endPage", "e", -1, "specify end page")
	flag.IntVarP(&args.pageLength,"page_len", "l", -1, "specify pageLength of line per page")
	flag.BoolVarP(&args.delimited,"page_split", "f", false, "specify if using delimiter")
	flag.StringVarP(&args.printDst,"print_dst", "d", "", "specify the printer")
	flag.Parse()
}

func processArgs(args *selpgArgs) {
	
	if len(os.Args) <= 3 {
		fmt.Fprintf(os.Stderr, "\n Error: \n %s: arguments are not enough\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}


	if args.startPage < 0 && args.endPage < 0 {
		fmt.Fprintf(os.Stderr, "\n Error: \n %s: arguments are not enough\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "\n Error: \n %s: the first arg must be -s which means the start page\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	endIndex := 2
	if len(os.Args[1]) == 2 {
		endIndex = 3
	}

	if os.Args[endIndex][0] != '-' || os.Args[endIndex][1] != 'e' {
		fmt.Fprintf(os.Stderr, "\n Error: \n %s: the second arg must be -e which means the end page\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if args.startPage > args.endPage || args.startPage < 0 || args.endPage < 0 {
		fmt.Fprintf(os.Stderr, "\n Error: \n %s: Invalid arguments\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if args.delimited == false {
		if args.pageLength == -1 {
			args.pageLength = 72
		}
	}

	if args.delimited == true && args.pageLength != -1 {
		fmt.Fprintf(os.Stderr, "\n Error: \n %s: delimited and pageLength can't coexist.\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

}

func processInput(args *selpgArgs) {
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd

	if args.printDst != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	if flag.NArg() > 0 {
		args.inputFileName = flag.Arg(0)
		output, err := os.Open(args.inputFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := bufio.NewReader(output)
		if args.delimited {
			for pageNum := 0; pageNum <= args.endPage; pageNum++ {
				line, err := reader.ReadString('\f')
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				printOrWrite(args, string(line), stdin)
			}
		} else {
			count := 0
	for {
		line, _, err := reader.ReadLine()
		if err != io.EOF && err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}
		if count / args.pageLength >= args.startPage {
			if count / args.pageLength >= args.endPage {
				break
			} else {
				printOrWrite(args, string(line), stdin)
			}
		}
		count++
	}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
	count := 0
	target := ""
	for scanner.Scan() {
		line := scanner.Text()
		line += "\n"
		if count / args.pageLength >= args.startPage {
			if count / args.pageLength < args.endPage {
				target += line
			}
		}
		count++
	}
	printOrWrite(args, string(target), stdin)
	}

	if args.printDst != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Usage()  {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tselpg -s=num -e=num [options] [filename]\n\n")
	fmt.Printf("The arguments are:\n\n")
	fmt.Printf("\t-s=num\tStart of Page <pageLength>.\n")
	fmt.Printf("\t-e=num\tEnd of Page <pageLength>.\n")
	fmt.Printf("\t-l=num\t[options]Specify the number of line per page.Default is 72.\n")
	fmt.Printf("\t-d=lp number\t[options]Using cat to test.\n")
	fmt.Printf("\t-f\t\t[options]Specify that the pages are sperated by \\f.\n")
	fmt.Printf("\t[filename]\t[options]Read input from the file.\n\n")
	fmt.Printf("If no file specified, %s will read input from stdin. Control-D to end.\n\n", progname)
}



func printOrWrite(args *selpgArgs, line string, stdin io.WriteCloser) {
	if args.printDst != "" {
		stdin.Write([]byte(line + "\n"))
	} else {
		fmt.Println(line)
	}
}

func main() {
	progname = os.Args[0]
	var args selpgArgs
	initArgs(&args)
	processArgs(&args)
	processInput(&args)
}
