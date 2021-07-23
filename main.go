package main

import (
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/kelseyhightower/envconfig"
	"regexp"

	"log"
)

type Config struct {
	VkToken string
}

func main1() {
	var cfg Config
	err := envconfig.Process("adapterbot", &cfg)
	if err != nil {
		log.Fatalf("process environment config: %v", err)
	}

	vk := api.NewVK(cfg.VkToken)

	//client, err := sheets.NewService(ctx)

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(group)
}

type SpreadsheetInfo struct {
	SpreadsheetId string
	SheetId       string
}

func getSpreadsheetInfoFromUrl(url string) (*SpreadsheetInfo, error) {
	regex, err := regexp.Compile("/\\d/([1-9A-z-]+)/.*gid=(\\d+)")
	if err != nil {
		return nil, fmt.Errorf("constructing regex: %w", err)
	}

	g := regex.FindStringSubmatch(url)
	if g == nil {
		return nil, errors.New("invalid url")
	}

	return &SpreadsheetInfo{
		SpreadsheetId: g[0],
		SheetId:       g[1],
	}, nil
}
