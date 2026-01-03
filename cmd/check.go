package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/dcastro0/netpulse/internal/monitor"
	"github.com/spf13/cobra"
)

var (
	timeout int
	file    string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Verifica a saúde de sites via URL única ou arquivo CSV",
	Run: func(cmd *cobra.Command, args []string) {
		var urls []string

		if file != "" {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println("Erro ao abrir arquivo:", err)
				return
			}
			defer f.Close()

			reader := csv.NewReader(f)
			records, err := reader.ReadAll()
			if err != nil {
				fmt.Println("Erro ao ler CSV:", err)
				return
			}
			for _, row := range records {
				if len(row) > 0 {
					urls = append(urls, row[0])
				}
			}
		} else if len(args) > 0 {
			urls = append(urls, args[0])
		} else {
			fmt.Println("Por favor, forneça uma URL ou use a flag --file")
			return
		}

		resultsChan := make(chan monitor.Result, len(urls))
		var wg sync.WaitGroup

		for _, u := range urls {
			wg.Add(1)
			go func(target string) {
				defer wg.Done()
				resultsChan <- monitor.Check(target, timeout)
			}(u)
		}

		go func() {
			wg.Wait()
			close(resultsChan)
		}()

		var results []monitor.Result
		for res := range resultsChan {
			results = append(results, res)
		}

		sort.Slice(results, func(i, j int) bool {
			return results[i].Duration < results[j].Duration
		})

		renderTable(results)
	},
}

func renderTable(results []monitor.Result) {
	rows := [][]string{}

	styleSuccess := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#04B575")). 
		Align(lipgloss.Center).
		Width(10)

	styleError := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#FF4672")). 
		Align(lipgloss.Center).
		Width(10)
	
	styleWarning := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#FFB86C")). 
		Align(lipgloss.Center).
		Width(10)

	for _, r := range results {
		var statusBadge string
		var statusText string

		if r.Err != nil {
			statusText = "ERROR"
			statusBadge = styleError.Render(statusText)
		} else {
			statusText = fmt.Sprintf("%d", r.Status)
			if r.Status >= 200 && r.Status < 300 {
				statusBadge = styleSuccess.Render(statusText)
			} else if r.Status >= 400 && r.Status < 500 {
				statusBadge = styleWarning.Render(statusText)
			} else {
				statusBadge = styleError.Render(statusText)
			}
		}

		rows = append(rows, []string{
			statusBadge,
			r.URL,
			fmt.Sprintf("%v", r.Duration),
		})
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers("STATUS", "URL", "LATÊNCIA").
		Rows(rows...)

	fmt.Println(t)
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "Timeout da requisição em segundos")
	checkCmd.Flags().StringVarP(&file, "file", "f", "", "Caminho para arquivo CSV com URLs")
}