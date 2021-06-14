/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/HungryPandaHome/airdrop/pkg/db"
	"github.com/HungryPandaHome/airdrop/pkg/db/models"
	"github.com/HungryPandaHome/airdrop/pkg/defs"
	"github.com/HungryPandaHome/airdrop/pkg/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PageDate ...
type PageData struct {
	TimeForNext time.Duration
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		cache[strings.Split(name, ".")[0]] = ts
	}

	return cache, nil
}

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start airdrop daemon",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		privateKey, err := crypto.HexToECDSA(viper.GetString("priv_key"))
		cobra.CheckErr(err)
		adminAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

		client, err := ethclient.Dial(viper.GetString("rpc"))
		cobra.CheckErr(err)

		pgUri := viper.GetString("db_url")
		dbCon, err := db.Open(pgUri)
		cobra.CheckErr(err)
		defer func() {
			_ = dbCon.Close()
		}()
		// migrate schema ...
		cobra.CheckErr(db.Migrate(dbCon, pgUri, db.CurrentSchemaVersion))

		cache, err := newTemplateCache("templates")
		cobra.CheckErr(err)
		e := echo.New()
		e.Static("/static", "static")

		e.GET("/", func(c echo.Context) error {
			// render template based on contract state
			contract, err := (models.ContractModel{DB: dbCon}).ReadContext(c.Request().Context())
			if err != nil && !errors.Is(err, models.ErrNotFound) {
				return err
			}
			if contract == nil || contract.Address == nil {
				return cache[defs.StageUnknown.String()].Execute(c.Response().Writer, PageData{})
			}
			stage, err := utils.GetStage(
				c.Request().Context(),
				client,
				adminAddress,
				contract.Address)
			if err != nil {
				return fmt.Errorf("query contract status: %s", err)
			}
			return cache[stage.String()].Execute(c.Response().Writer, PageData{})
		})
		e.GET("/success/:id", func(c echo.Context) error {
			// render template based on contract state
			// 1 registration is not started yet
			// 2 registartion is in process
			// 3 claming is in process
			// 4 ended
			return cache["nearby"].Execute(c.Response().Writer, PageData{})
		})
		e.GET("/sorry/:id", func(c echo.Context) error {
			// render template based on contract state
			// 1 registration is not started yet
			// 2 registartion is in process
			// 3 claming is in process
			// 4 ended
			return cache["nearby"].Execute(c.Response().Writer, PageData{})
		})
		e.POST("/regiter/twitter", func(c echo.Context) error {
			// TODO: verify input
			// TODO: verify post format
			// TODO: genrate sign
			// TODO: save sign into db
			// TODO: render page with sign
			return nil
		})
		e.POST("/regiter/facebook", func(c echo.Context) error {
			// TODO: verify input
			// TODO: verify post format
			// TODO: genrate sign
			// TODO: save sign into db
			// TODO: render page with sign
			return nil
		})
		e.POST("/regiter/reddit", func(c echo.Context) error {
			// TODO: verify input
			// TODO: verify post format
			// TODO: genrate sign
			// TODO: save sign into db
			// TODO: render page with sign
			return nil
		})
		e.Logger.Fatal(e.Start(viper.GetString("bind")))
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
