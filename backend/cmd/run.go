package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"pimview.thelabshack.com/pkg/mqtt"
	"pimview.thelabshack.com/pkg/publisher"
	"pimview.thelabshack.com/pkg/subscriber"
	"time"
)

//var (
//	log = log.NewLogger()
//)

func RunPlugin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run pimview plugin",
	}

	cmd.AddCommand(RunPublisher())
	cmd.AddCommand(RunSubscriber())

	return cmd
}

func RunPublisher() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pub",
		Short: "Run Pimview Test Publisher",
		Run: func(cmd *cobra.Command, args []string) {

			ticker := time.NewTicker(10 * time.Second)
			done := make(chan struct{})

			client := mqtt.GetClient("pimviewpub1")

			for {

				publisher.Run(client)

				select {
				case <-done:
					return
				case <-ticker.C:
					log.Println("redo")
				}
			}
		},
	}

	return cmd
}

func RunSubscriber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sub",
		Short: "Run Pimview Test Publisher",
		Run: func(cmd *cobra.Command, args []string) {
			//ticker := time.NewTicker(30 * time.Second)
			//done := make(chan struct{})
			//
			//for {
			subscriber.Run()

			//	select {
			//	case <-done:
			//		return
			//	case <-ticker.C:
			//		log.Println("redo")
			//	}
			//}
		},
	}

	return cmd
}
