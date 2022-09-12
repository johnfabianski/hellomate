package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "hellomate",
	Long: "Root command",
}

var HelloCmd = &cobra.Command{
	Use:   "hello <name>",
	Short: "Say hello to someone",
	Long:  `Say hello to someone`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args)
		fmt.Println("hello " + args[0])
		// cmd.Help()
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return []string{"steve", "john"}, cobra.ShellCompDirectiveNoFileComp
	},
}

// var HelloCmd = &cobra.Command{
// 	Use:   "hello <name>",
// 	Short: "Get user data",
// 	Args:  cobra.ExactValidArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Printf("Hello %s!!!\n", args[0])
// 	},
// 	ValidArgsFunction: UserGet,
// }

func getNames() []string {
	file, err := os.Open("common-names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return names
}

func UserGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println("UserGet called")
	return getNames(), cobra.ShellCompDirectiveNoFileComp //
}
func init() {
	RootCmd.AddCommand(HelloCmd) //, cmd.CompletionCmd)
}
func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// git tag -a 1.0.8 -m "shell completion"
// goreleaser --rm-dist

// So for some reason, the get names thing doesn't work either. Maybe something to do with the ValidArgsFunction
