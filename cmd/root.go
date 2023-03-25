/*
Copyright © 2023 Natsuki Kirigakure <natsuki@kirigakure.net>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	regexpPattern string
	inputFile     string
	outputFile    = "result.csv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "regexport",
	Short: "regexport is a tool to export csv from a file using a regular expression.",
	Long:  `regexport is a tool to export csv from a file using a regular expression.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		Regex(regexpPattern, inputFile, outputFile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.regexport.yaml)")
	rootCmd.PersistentFlags().StringVarP(&regexpPattern, "pattern", "p", "", "regular expression pattern")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "input file")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "result.csv", "output file")

	check(rootCmd.MarkPersistentFlagFilename("input"))
	check(rootCmd.MarkPersistentFlagFilename("output"))

	check(rootCmd.MarkPersistentFlagRequired("pattern"))
	check(rootCmd.MarkPersistentFlagRequired("input"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")
}
