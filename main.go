package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/KarolosLykos/hackertea/internal/cache"
	"github.com/KarolosLykos/hackertea/internal/client"
	"github.com/KarolosLykos/hackertea/internal/constants"
	"github.com/KarolosLykos/hackertea/internal/hn"
	"github.com/KarolosLykos/hackertea/internal/tui/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ctx := context.Background()

	c := client.New(constants.BaseURL, &http.Client{Timeout: 10 * time.Second})
	memCache := cache.New()
	hnClient := hn.New(c, memCache)

	m, err := model.New(ctx, hnClient)
	if err != nil {
		fmt.Println("Error creating model: ", err)
		os.Exit(1)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err = p.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
