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
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/HungryPandaHome/airdrop/pkg/db"
	"github.com/HungryPandaHome/airdrop/pkg/defs"
	"github.com/HungryPandaHome/airdrop/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	fb "github.com/huandu/facebook/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// PageDate ...
type PageData struct {
	FacebookSupport  bool
	NicknameRequired bool
	TwitterSupport   bool
	TimeForNext      time.Duration
}

// WalletBind ...
type WalletBind struct {
	Wallet string `param:"wallet" query:"wallet" form:"wallet" json:"wallet"`
}

// OauthBind ...
type OauthBind struct {
	Code  string `param:"code" query:"code"`
	State string `param:"state" query:"state"`
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

		facebookOauth := &oauth2.Config{
			ClientID:     viper.GetString("fb_app_id"),
			ClientSecret: viper.GetString("fb_app_secret"),
			RedirectURL:  strings.TrimRight(viper.GetString("base_url"), "/") + "/callback",
			Scopes:       []string{"public_profile"},
			Endpoint:     facebook.Endpoint,
		}

		cache, err := newTemplateCache("templates")
		cobra.CheckErr(err)
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Static("/static", "static")

		e.GET("/", func(c echo.Context) error {
			_ = adminAddress
			_ = client
			// render template based on contract state
			// contract, err := (models.ContractModel{DB: dbCon}).ReadContext(c.Request().Context())
			// if err != nil && !errors.Is(err, models.ErrNotFound) {
			// 	return err
			// }
			// if contract == nil || contract.Address == nil {
			// 	return cache[defs.StageUnknown.String()].Execute(c.Response().Writer, PageData{})
			// }
			// stage, err := utils.GetStage(
			// 	c.Request().Context(),
			// 	client,
			// 	adminAddress,
			// 	contract.Address)
			// if err != nil {
			// 	return fmt.Errorf("query contract status: %s", err)
			// }
			return cache[defs.StageRegistration.String()].Execute(c.Response().Writer,
				PageData{
					FacebookSupport: true,
				})
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
		e.POST("/register/twitter", func(c echo.Context) error {
			// TODO: verify input
			// TODO: verify post format
			// TODO: genrate sign
			// TODO: save sign into db
			// TODO: render page with sign
			return nil
		})
		e.POST("/register/facebook", func(c echo.Context) error {
			u := new(WalletBind)
			if err := c.Bind(u); err != nil {
				return err
			}
			// if !common.IsHexAddress(u.Wallet) {
			// 	err = fmt.Errorf("invalid wallet format")
			// 	c.Logger().Error(err)
			// 	return err
			// }

			data := []byte{byte(defs.FacebookSource)}
			data = append(data, common.HexToAddress(u.Wallet).Bytes()...)
			sign, err := utils.SignWith(data,
				[]byte(viper.GetString("fb_app_secret")))
			if err != nil {
				return err
			}
			data = append(data, sign...)
			theUrl := facebookOauth.AuthCodeURL(base64.URLEncoding.
				Strict().EncodeToString(data))
			return c.Redirect(http.StatusPermanentRedirect, theUrl)
		})
		e.POST("/register/reddit", func(c echo.Context) error {
			// TODO: verify input
			// TODO: verify post format
			// TODO: genrate sign
			// TODO: save sign into db
			// TODO: render page with sign
			return nil
		})
		// handle facebook logons
		e.GET("/callback", func(c echo.Context) error {
			a := new(OauthBind)
			if err := c.Bind(a); err != nil {
				return err
			}
			stateBytes, err := base64.URLEncoding.
				Strict().DecodeString(a.State)
			if err != nil {
				return err
			}
			source := defs.SocialSource(stateBytes[0])

			var secret []byte
			switch source {
			case defs.FacebookSource:
				secret = []byte(viper.GetString("fb_app_secret"))
			}

			if !utils.VerifyHMACSign(
				stateBytes[:21],
				stateBytes[21:],
				secret) {
				return fmt.Errorf("failed to verify signature")
			}

			var userID string
			switch source {
			case defs.FacebookSource:
				token, err := facebookOauth.Exchange(c.Request().Context(), a.Code)
				if err != nil {
					return err
				}
				fbClient := facebookOauth.Client(c.Request().Context(), token)
				session := &fb.Session{
					Version:    "v11.0",
					HttpClient: fbClient,
				}

				// Use session.
				res, err := session.Get("/me", nil)
				if err != nil {
					return err
				}
				id, ok := res["id"]
				if !ok {
					return fmt.Errorf("unexpected facebook response format")
				}
				userID = fmt.Sprint(id)
			default:
				return fmt.Errorf("unsupported social source")
			}
			_ = userID

			// TODO: insert into datbase
			// walletBytes := stateBytes[1:21]
			// signBytes := stateBytes[22:]

			return cache["success"].Execute(c.Response().Writer, PageData{})
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
