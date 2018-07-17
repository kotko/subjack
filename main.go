package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kotko/subjack/subjack"
)

func main() {
	o := subjack.Options{}

	flag.StringVar(&o.Wordlist, "w", "", "Path to wordlist.")
	flag.IntVar(&o.Threads, "t", 10, "Number of concurrent threads (Default: 10).")
	flag.IntVar(&o.Timeout, "timeout", 10, "Seconds to wait before connection timeout (Default: 10).")
	flag.BoolVar(&o.Ssl, "ssl", false, "Force HTTPS connections (May increase accuracy (Default: http://).")
	flag.BoolVar(&o.All, "a", false, "Find those hidden gems by sending requests to every URL. (Default: Requests are only sent to URLs with identified CNAMEs).")
	flag.BoolVar(&o.Verbose, "v", false, "Display more information per each request.")
	flag.StringVar(&o.Output, "o", "", "Output results to file.")

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	subjack.Process(&o)
}
