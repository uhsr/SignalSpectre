// cmd/signalspectre/main.go
package main

import (
"flag"
"log"
"os"

"signalspectre/internal/signalspectre"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := signalspectre.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
