package main

import (
  "chain/config"
  "chain/domain"
  "chain/presentation"
  "fmt"
  "os"
  "strconv"
  "time"
)

func printUsage() {
  fmt.Println("Usage: chain <size> or chain web")
}

func runAsCli(args []string) {
  startTime := time.Now()
  size, err := strconv.Atoi(args[0])

  if err != nil {
    fmt.Println("Size should be an integer")
    return
  }

  defer func() {
    finish := domain.GetTimeTracker("Mined "+strconv.Itoa(size+1)+" blocks (including genesis)", &startTime)
    fmt.Println(finish())
  }()

  blockchain := domain.CreateChain(size)

  fmt.Println(domain.StringifyChain(blockchain))
}

func runAsApi () {
  blockChain := domain.CreateChain(11)
  presentation.StartServer(config.NewAppConfig(), blockChain)
}

func main() {
  args := os.Args[1:]

  if len(args) != 1 {
    printUsage()
    return
  }

  if args[0] == "web" {
    runAsApi()
    return
  }

  runAsCli(args)
}
