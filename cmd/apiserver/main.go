package main

import (
	"github.com/ividernvi/algohub/internal/apiserver"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(apiserver.Run())
}
