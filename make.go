// Build script for producing multiple binaries out of a single project
// Has fallen into disrepair - has been commented out in the interests of
// getting the project going again

package main

// import (
//	"bytes"
//	"os"
//	"os/exec"
//	"log"
//	"flag"
//	"strings"
// )

// var (
//	targets = flag.String("targets", "", "Optional comma-separated list of targets (i.e go packages) to build and install. '*' builds everything.  Empty builds defaults for this platform. Example: abcd/cmd/abcv, abcd/cmd/abcd")
//	verbose = flag.Bool("v", false, "Verbose mode")
// )

// func main() {
//	// Flags/logging
//	log.SetFlags(0)
//	flag.Parse()

//	buildAll := false

//	targs := []string{
//		"abcd/cmd/abcd",
//		"abcd/cmd/abcv",
//	}

//	switch *targets {
//		case "*":
//			buildAll = true
//		default:
//			if t := strings.Split(*targets, ","); len(t) != 0 {
//				targs = t
//			}
//	}

//	// Juggle argv
//	args := append(baseArgs, targs...)

//	cmd := exec.Command("go", args...)
//	cmd.Env = append(cleanGoEnv(),
//		"GOPATH="+buildGoPath,
//	)
//	if *verbose {
//		log.Printf("Running go %q with Env %q", args, cmd.Env)
//	}
//	var output bytes.Buffer
//	if *quiet {
//		cmd.Stdout = &output
//		cmd.Stderr = &output
//	} else {
//		cmd.Stdout = os.Stdout
//		cmd.Stderr = os.Stderr
//	}
//	if *verbose {
//		log.Printf("Running go install of main binaries with args %s", cmd.Args)
//	}
//	if err := cmd.Run(); err != nil {
//		log.Fatalf("Error building main binaries: %v\n%s", err, output.String())
//	}
//	if *verbose {
//		log.Printf("Success. Binaries are in %s", "/bin") // TODO: Make dynamic
//	}
// }
