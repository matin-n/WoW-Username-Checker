package main

import (
	"github.com/micmonay/keybd_event"
	"io/ioutil"
	"strings"
	"time"
)


func main() {

	time.Sleep(5 * time.Second)

  // Read usernames.txt which contains list of usernames to check for
	content, _ := ioutil.ReadFile("usernames.txt")

  // Each username within usernames.txt should be split by a new line
	lines := strings.Split(string(content), "\n")

  // Read thru the usernames.txt and check for username availability
	for i :=0; i < len(lines); i++{
		checkUsername(lines[i])
		time.Sleep(3 * time.Second)
	}

}


func checkUsername(input string) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(input); i++ {
		switch string(input[i]){
		case "a":
			kb.SetKeys(keybd_event.VK_A)
			kb.Launching()
		case "b":
			kb.SetKeys(keybd_event.VK_B)
			kb.Launching()
		case "c":
			kb.SetKeys(keybd_event.VK_C)
			kb.Launching()
		case "d":
			kb.SetKeys(keybd_event.VK_D)
			kb.Launching()
		case "e":
			kb.SetKeys(keybd_event.VK_E)
			kb.Launching()
		case "f":
			kb.SetKeys(keybd_event.VK_F)
			kb.Launching()
		case "g":
			kb.SetKeys(keybd_event.VK_G)
			kb.Launching()
		case "h":
			kb.SetKeys(keybd_event.VK_H)
			kb.Launching()
		case "i":
			kb.SetKeys(keybd_event.VK_I)
			kb.Launching()
		case "j":
			kb.SetKeys(keybd_event.VK_J)
			kb.Launching()
		case "k":
			kb.SetKeys(keybd_event.VK_K)
			kb.Launching()
		case "l":
			kb.SetKeys(keybd_event.VK_L)
			kb.Launching()
		case "m":
			kb.SetKeys(keybd_event.VK_M)
			kb.Launching()
		case "n":
			kb.SetKeys(keybd_event.VK_N)
			kb.Launching()
		case "o":
			kb.SetKeys(keybd_event.VK_O)
			kb.Launching()
		case "p":
			kb.SetKeys(keybd_event.VK_P)
			kb.Launching()
		case "q":
			kb.SetKeys(keybd_event.VK_Q)
			kb.Launching()
		case "r":
			kb.SetKeys(keybd_event.VK_R)
			kb.Launching()
		case "s":
			kb.SetKeys(keybd_event.VK_S)
			kb.Launching()
		case "t":
			kb.SetKeys(keybd_event.VK_T)
			kb.Launching()
		case "u":
			kb.SetKeys(keybd_event.VK_U)
			kb.Launching()
		case "v":
			kb.SetKeys(keybd_event.VK_V)
			kb.Launching()
		case "w":
			kb.SetKeys(keybd_event.VK_W)
			kb.Launching()
		case "x":
			kb.SetKeys(keybd_event.VK_X)
			kb.Launching()
		case "y":
			kb.SetKeys(keybd_event.VK_Y)
			kb.Launching()
		case "z":
			kb.SetKeys(keybd_event.VK_Z)
			kb.Launching()
		}
	}

	time.Sleep(2 * time.Second)
	kb.SetKeys(keybd_event.VK_ENTER)
	kb.Launching()

	time.Sleep(2 * time.Second)
	kb.SetKeys(keybd_event.VK_ENTER)
	kb.Launching()

	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_A)
	kb.SetKeys(keybd_event.VK_BACKSPACE)
	kb.Launching()
	time.Sleep(1 * time.Second)
}
