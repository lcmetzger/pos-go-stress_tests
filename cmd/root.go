package cmd

import (
	"os"

	"github.com/lcmetzger/stress_test/internal/report"
	"github.com/lcmetzger/stress_test/internal/req"
	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "st",
	Short: "utilize para realizar testes de stress de um serviço http",
	Long: `
Para realizar testes de stress de um serviço http, basta informar os parâmetros requeridos,
sendo obrigatório informar a url a ser chamada através do parâmetro --url.

Demais parâmetos são opcionais e serão assumidos com o valor 1 (um), caso não sejam informados.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			cmd.Help()
			os.Exit(1)
		}

		report := report.NewReport()
		report.Url = url
		report.Concurrency = concurrency
		report.Requests = requests

		req.MakeRequests(report)
		report.Results()

	},
}

// func makeRequest() error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Status:", resp.StatusCode)

// 	return nil
// }

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "url do serviço http a ser requisitado.")
	rootCmd.PersistentFlags().IntVarP(&requests, "requests", "r", 1, "número total de requests.")
	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 1, "número de chamadas simultâneas.")
}
