package account

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/dj-yacine-flutter/y-z-a/pb/uspb"
	"github.com/dj-yacine-flutter/y-z-a/utils"
)

type Account struct {
	uscl uspb.UserServiceClient
}

type UserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Sign(uscl uspb.UserServiceClient) error {

	var access bool
	deviceHash := generateDeviceHash()
	account := &Account{
		uscl: uscl,
	}

	accountFile := filepath.Join(utils.ConfigPath, "account.json")
	if _, err := os.Stat(accountFile); os.IsNotExist(err) {

		fmt.Println("To login press [Y] - To create a account press [N]: ")
		var choice string
		if _, err := fmt.Scanln(&choice); err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		choice = strings.TrimSpace(strings.ToLower(choice))
		if choice == "n" {
			access = true
		} else if choice == "y" {
			access = false
		} else {
			log.Fatal("you cannot use the app without account.")
		}

		err = os.WriteFile(accountFile, nil, 0644)
		if err != nil {
			log.Fatal("failed to create account.json")
		}
	}

	if access {
		return account.createAccount(accountFile, deviceHash)
	}

	if err := account.loginAccount(accountFile, deviceHash); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("There is no user with this data!! Do you want to create an account? [Y/n]: ")

			var choice string
			if _, err := fmt.Scanln(&choice); err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}

			choice = strings.TrimSpace(strings.ToLower(choice))
			if choice == "n" {
				fmt.Println("Operation canceled.")
				return nil
			}

			return account.createAccount(accountFile, deviceHash)
		} else if errors.Is(err, ErrPermissionDenied) {
			log.Fatal("Your device isn't allowed to use this account.")
		} else {
			return err
		}
	}

	return nil
}

func generateDeviceHash() string {
	// Get IP address
	ip := getIPAddress()

	// Get MAC address
	mac := getMACAddress()

	// Get hard drive serial number
	serial := getHardDriveSerial()

	// Concatenate and hash the values
	hashInput := fmt.Sprintf("%s:%s:%s", ip, mac, serial)
	hash := md5.Sum([]byte(hashInput))

	return fmt.Sprintf("%x", hash)
}

func getIPAddress() string {
	var ip string
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
					break
				}
			}
		}
		if ip != "" {
			break
		}
	}
	return ip
}

func getMACAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			return iface.HardwareAddr.String()
		}
	}
	return ""
}

func getHardDriveSerial() string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("wmic", "diskdrive", "get", "SerialNumber")
	} else {
		cmd = exec.Command("lsblk", "-o", "SERIAL", "-n")
	}
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	serial := strings.TrimSpace(string(output))
	return serial
}
