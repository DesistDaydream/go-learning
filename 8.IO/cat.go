package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	flag.Parse()
	// NArg返回解析flag之后剩余的参数的个数
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cat: 无法打开文件%s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(f)
		f.Close()
	}
}



func cat(f *os.File) {
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, err := f.Read(buf[:]); true {
        case nr < 0:
            fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
            os.Exit(1)
        case nr == 0: // EOF
            return
        case nr > 0:
            if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
                fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
            }
        }
    }		
}