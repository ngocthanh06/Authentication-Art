package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/ngocthanh06/authentication/internal/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

var initJwtCmd = &cobra.Command{
	Use:     "init-jwt-key",
	Aliases: []string{"init-jwt-key"},
	Short:   "Init jwt key",
	Long:    "Init jwt key if env haven't exists",
	Run: func(cmd *cobra.Command, args []string) {
		generateJwtKey()
	},
}

func init() {
	rootCmd.AddCommand(initJwtCmd)
}

func generateJwtKey() {
	if len(config.EnvKey.JwtKey) == 0 {
		key := make([]byte, 32)

		if _, err := rand.Read(key); err != nil {
			panic(err)
		}
		jwtKey := hex.EncodeToString(key)

		os.Setenv("JWT_KEY", jwtKey)

		envFileContent, err := ioutil.ReadFile(".env")

		if err != nil && !os.IsNotExist(err) {
			log.Fatalf("Error readling .env file: %v", err)

			return
		}

		content := string(envFileContent)
		newLine := fmt.Sprintf("JWT_KEY=%s", jwtKey)

		if strings.Contains(content, "JWT_KEY=") {
			content = strings.Replace(content, "JWT_KEY=", newLine, 1)
		} else {
			content = content + "\n" + newLine
		}

		err = ioutil.WriteFile(".env", []byte(content), 0644)
		if err != nil {
			log.Fatalf("Error writing to .env file: %v", err)

			return
		}

		return
	}
}
