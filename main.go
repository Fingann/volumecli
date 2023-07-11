package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"volumecli/volume"
)

var mixer string
var element string

func init() {
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(setVolumeCmd)
	rootCmd.AddCommand(muteCmd)
	rootCmd.AddCommand(unmuteCmd)
	rootCmd.AddCommand(toggleCmd)
}

var rootCmd = &cobra.Command{
	Use:   "volumecli",
	Short: "A CLI for controlling your computer's volume",
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&mixer, "mixer", "m", "default", "Mixer name")
	rootCmd.PersistentFlags().StringVarP(&element, "element", "e", "Master", "Element name")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Increase volume by two percent",
	Run: func(cmd *cobra.Command, args []string) {
		handler, err := volume.NewVolumeHandler(mixer, element)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer handler.Close()

		volume, err := handler.GetVolume()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = handler.SetVolume(volume + 2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Volume increased")
	},
}
var setVolumeCmd = &cobra.Command{
	Use:   "set [volume]",
	Short: "Set the volume to a specific value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wantedVolume, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: volume must be an integer")
			return
		}

		handler, err := volume.NewVolumeHandler(mixer, element)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer handler.Close()

		err = handler.SetVolume(wantedVolume)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Volume set to", wantedVolume)
	},
}
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Decrease volume by two percent",
	Run: func(cmd *cobra.Command, args []string) {
		handler, err := volume.NewVolumeHandler(mixer, element)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer handler.Close()

		volume, err := handler.GetVolume()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = handler.SetVolume(volume - 2)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Volume decreased")
	},
}

var muteCmd = &cobra.Command{
	Use:   "mute",
	Short: "Mute the volume",
	Run: func(cmd *cobra.Command, args []string) {
		handler, err := volume.NewVolumeHandler(mixer, element)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer handler.Close()

		err = handler.SetMute(true)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Volume muted")
	},
}

var unmuteCmd = &cobra.Command{
	Use:   "unmute",
	Short: "Unmute the volume",
	Run: func(cmd *cobra.Command, args []string) {
		handler, err := volume.NewVolumeHandler(mixer, element)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer handler.Close()

		err = handler.SetMute(false)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Volume unmuted")
	},
}

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle mute/unmute",
	Run: func(cmd *cobra.Command, args []string) {
		handler, err := volume.NewVolumeHandler(mixer, element)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer handler.Close()

		muted, err := handler.IsMuted()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = handler.SetMute(!muted)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if muted {
			fmt.Println("Volume unmuted")
		} else {
			fmt.Println("Volume muted")
		}
	},
}
