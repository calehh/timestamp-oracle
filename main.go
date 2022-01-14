package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/mux"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	OriginCommandHelpTemplate = `{{.Name}}{{if .Subcommands}} command{{end}}{{if .Flags}} [command options]{{end}} {{.ArgsUsage}}
{{if .Description}}{{.Description}}
{{end}}{{if .Subcommands}}
SUBCOMMANDS:
  {{range .Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}{{end}}{{if .Flags}}
OPTIONS:
{{range $.Flags}}   {{.}}
{{end}}
{{end}}`
)

var app *cli.App

var (
	portFlag = cli.StringFlag{
		Name:  "port",
		Usage: "restful rpc port",
		Value: "3333",
	}

	fileFlag = cli.StringFlag{
		Name:  "private",
		Usage: "private key file path",
		Value: "./private.txt",
	}
)

var PrivateKeyStr = "db96ee03632eb01fed47f061d2ba383ad8231844d51fc94d9a938d2476e7fd53"

func init() {
	app = cli.NewApp()
	app.Version = "v1.0.0"
	app.Action = Start
	app.Flags = []cli.Flag{
		portFlag,
		fileFlag,
	}

	cli.CommandHelpTemplate = OriginCommandHelpTemplate
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type res struct {
	Timestamp string `json:"timestamp"`
	Sig       []byte `json:"sig"`
}

func Start(ctx *cli.Context) {
	log.Info("starting oracle")
	if ctx.IsSet(fileFlag.Name) {
		log.Info("loading key")
		filePath := ctx.String(fileFlag.Name)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		} else {
			PrivateKeyStr = string(b)
		}
	}

	pri, err := crypto.HexToECDSA(PrivateKeyStr)
	if err != nil {
		panic(err)
	}
	log.Info("auth address", "address", crypto.PubkeyToAddress(pri.PublicKey).String())

	log.Info("start restful server")
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		timestamp, sig, err := SignTimeStamp()
		if err != nil {
			_, e := writer.Write([]byte{})
			if e != nil {
				log.Error("write response err", "err", err)
			}
			return
		}
		re, err := json.Marshal(res{
			Timestamp: timestamp,
			Sig:       sig,
		})
		if err != nil {
			log.Error("marshal response err", "err", err)
		}
		_, e := writer.Write(re)
		if e != nil {
			log.Error("write response err", "err", err)
		}
	})

	address := ""
	if ctx.IsSet(portFlag.Name) {
		address = "0.0.0.0:" + ctx.String(portFlag.Name)
	} else {
		address = "0.0.0.0:8547"
	}
	log.Info("address", "url", address)
	err = http.ListenAndServe(address, r)
	if err != nil {
		utils.Fatalf("http listen err", "err", err)
	}
	waitToExit()
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	if !signal.Ignored(syscall.SIGHUP) {
		signal.Notify(sc, syscall.SIGHUP)
	}
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sc {
			fmt.Printf("received exit signal:%v", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}

func SignTimeStamp() (string, []byte, error) {
	uintTy, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return "", nil, err
	}

	argument := abi.Arguments{
		{
			Type: uintTy,
		},
	}

	timestamp := big.NewInt(time.Now().Unix())

	msg, err := argument.Pack(timestamp)
	if err != nil {
		return "", nil, err
	}

	msgHash := crypto.Keccak256(msg)

	privKey, err := crypto.HexToECDSA(PrivateKeyStr)
	if err != nil {
		return "", nil, err
	}

	sigRaw, err := crypto.Sign(msgHash, privKey)
	if err != nil {
		return "", nil, err
	}

	sig := append(sigRaw[:64], uint8(int(sigRaw[64]))+27)

	return timestamp.String(), sig, nil
}
