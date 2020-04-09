package main

import (
	"bufio"
	"fmt"
	"github.com/jawher/mow.cli"
	"io"
	"io/ioutil"
	"log"
	"os"
	"playgo/advexamples/iocookbook/internal/composereaders"
	"playgo/advexamples/iocookbook/internal/composewriters"
	"playgo/advexamples/iocookbook/internal/implementreader"
	"playgo/advexamples/iocookbook/internal/implwriter"
	"playgo/advexamples/iocookbook/internal/usingreader"
	"strings"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("Golang IO Cookbook", "demo for Using IO READER and Writer interface")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("ir", "implementing reader", cli.ActionCommand(implementreader.Main))
	app.Command("cr", "Composing io.Readers", cli.ActionCommand(composereaders.Main))
	app.Command("dw", "Directly using io.Writer", cli.ActionCommand(func() {
		writer := os.Stdout
		writer.Write([]byte("hello there\n"))
	}))
	app.Command("iw", "Implementing io.Writer", cli.ActionCommand(implwriter.Main))
	app.Command("cw", "IComposing io.Writers", cli.ActionCommand(composewriters.Main))
	app.Command("ic", "io.Copy", cli.ActionCommand(func() {
		reader := strings.NewReader("this is the stuff I'm reading\n")
		writer := os.Stdout
		// io.Copy uses a 32kb buffer and siphons data from the reader through the buffer and to the writer.
		n, err := io.Copy(writer, reader)
		fmt.Printf("%d bytes written\n", n)
		if err != nil {
			log.Fatal(err)
		}
	}))

	app.Command("uw", "augment  Writer : 增强 变换在writer端", cli.ActionCommand(func() {
		reader := strings.NewReader("this is the stuff I'm reading\n")
		originalWriter := os.Stdout

		augmentedWriter := composewriters.UpcaseWriter(composewriters.BangWriter(originalWriter))
		_, err := io.Copy(augmentedWriter, reader)
		if err != nil {
			log.Fatal(err)
		}
	}))
	app.Command("ur", "augment  Reader : 增强 变换发生在reader端", cli.ActionCommand(func() {
		originalReader := strings.NewReader("this is the stuff I'm reading\n")
		writer := os.Stdout

		augmentedReader := composereaders.UpcaseReader(composereaders.BangReader(originalReader))
		_, err := io.Copy(writer, augmentedReader)
		if err != nil {
			log.Fatal(err)
		}
	}))
	app.Command("pipe", "io.Pipe", cli.ActionCommand(func() {
		originalReader := strings.NewReader("this is the stuff I'm reading\n")
		originalWriter := os.Stdout

		pipeReader, pipeWriter := io.Pipe()

		go func() {
			defer pipeWriter.Close()
			_, err := io.Copy(composewriters.UpcaseWriter(pipeWriter), originalReader)
			if err != nil {
				log.Fatal(err)
			}
		}()

		defer pipeReader.Close()
		_, err := io.Copy(originalWriter, composereaders.EncryptReader(pipeReader))
		if err != nil {
			log.Fatal(err)
		}
		// output: 'pdeo<eo<pda<opqbb<eCi<na]`ejc&' (notably not uppercased)
	}))

	app.Command("tr", "teeReader： fanout 扇出现象", cli.ActionCommand(func() {
		reader := strings.NewReader("look at me\n")

		file, err := os.OpenFile("file.encrypted", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		teeReader := io.TeeReader(reader, composewriters.EncryptWriter(file))

		_, err = io.Copy(os.Stdout, teeReader)
		if err != nil {
			log.Fatal(err)
		}
		// output: this is the stuff I'm reading\n
		// file.encrypted's contents: ����<��<���<�����<eC�<��}����&
	}))

	app.Command("ce", "Complex example", cli.ActionCommand(func() {
		originalReader := strings.NewReader("look at me\n")

		file, err := os.OpenFile("file.encrypted", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		pipeReader, pipeWriter := io.Pipe()
		teeReader := io.TeeReader(originalReader, composewriters.UpcaseWriter(os.Stdout))

		go func() {
			defer pipeWriter.Close()

			_, err = io.Copy(pipeWriter, teeReader)
			if err != nil {
				log.Fatal(err)
			}
		}()

		defer pipeReader.Close()
		_, err = io.Copy(file, composereaders.EncryptReader(pipeReader))
		if err != nil {
			log.Fatal(err)
		}
	}))

	app.Command("sc", "helpers: bufio.Scanner", cli.ActionCommand(func() {
		originalReader := strings.NewReader("the internet\nis a strange\nplace")

		scanner := bufio.NewScanner(originalReader)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			token := scanner.Text()
			fmt.Println(token)
		}
	}))

	app.Command("ws", "helpers: io.WriteString", cli.ActionCommand(func() {
		n, err := io.WriteString(os.Stdout, "test\n")
		fmt.Printf("%d bytes written\n", n)
		if err != nil {
			log.Fatal(err)
		}
	}))

	app.Command("rf", "files :ioutil.ReadFile ", cli.ActionCommand(func() {
		bytes, err := ioutil.ReadFile("file")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bytes)
	}))
	app.Command("wf", "files : ioutil.WriteFile", cli.ActionCommand(func() {
		err := ioutil.WriteFile("file", []byte("test"), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}))
	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================
func basics() {
	usingreader.Main()
}
