/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"../../client"
	"log"
	"net/rpc"

	"github.com/spf13/cobra"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := rpc.DialHTTP("tcp", "127.0.0.1:12039")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			log.Fatal(err)
		}
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal(err)
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		flags := &client.Args{
			Addr:         addr,
			RoomId:       id,
			RoomPassword: password,
			ClientName:   name,
		}
		joinCall := c.Go("Intermediate.Join", flags, &client.Reply{}, nil)
		replyCall := <-joinCall.Done // will be equal to divCall
		log.Println(replyCall)
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
	joinCmd.Flags().String("addr", "127.0.0.1:8080", "Address of the websocket")
	joinCmd.Flags().String("id", "default room", "Id of the room")
	joinCmd.Flags().String("password", "default password", "Password of the room")
	joinCmd.Flags().String("name", "default name", "Name of the client")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// joinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// joinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
