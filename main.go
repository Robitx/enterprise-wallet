package main

/*
 * Begins all the services required by the GUI wallet
 * 		- WSAPI for wallet
 *		- Webserver
 * Requires for all functionality
 *		- Factomd Instance
 */

import (
	"fmt"

	"github.com/FactomProject/enterprise-wallet/wallet"
	"github.com/FactomProject/factomd/util"
)

var MasterWallet *wallet.WalletDB

func close() {
	fmt.Println("Shutting down gracefully...")
	if MasterWallet == nil {
		return
	}

	err := MasterWallet.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Complete shut down.")
}

// InitiateWalletAndWeb initiates and serves the guiwallet. If databases are given, they will be attempted to be loaded
// and will be created if they are not found.
func InitiateWalletAndWeb(guiDBStr string, walDBStr string, txDBStr string, port int, walletdPort int, v1Import bool, v1Path string, factomdLocFlag string) {
	fmt.Println("--------- Initiating GUIWallet ----------")

	filename := util.ConfigFilename() //file name and path to factomd.conf file
	cfg := util.ReadConfig(filename)

	// Ports
	walletPort := walletdPort
	factomdLocation := cfg.Walletd.FactomdLocation
	if factomdLocFlag != "" {
		factomdLocation = factomdLocFlag
	}

	controlPanelPort := cfg.App.ControlPanelPort
	if cfg.App.ControlPanelSetting == "disabled" {
		controlPanelPort = -1
	}

	wallet.WalletBoltV1Path = v1Path

	var (
		walletDB, guiDB, txDB int
	)

	// DB Types
	switch guiDBStr { // Holds names associated with addresses for gui. Also holds settings
	case "Map":
		guiDB = wallet.MAP
	case "Bolt":
		guiDB = wallet.BOLT
	case "LDB":
		guiDB = wallet.LDB
	}
	switch walDBStr { // WalletDB is DB used by wallet wsapi
	case "Map":
		walletDB = wallet.MAP
	case "Bolt":
		walletDB = wallet.BOLT
	case "LDB":
		walletDB = wallet.LDB
	}
	switch txDBStr { // Holds transactions cache
	case "Map":
		txDB = wallet.MAP
	case "Bolt":
		txDB = wallet.BOLT
	case "LDB":
		txDB = wallet.LDB
	}

	// Start Walletd
	fmt.Printf("Wallet DB using %s, GUI DB using %s, TX DB using %s\n", IntToStringDBType(walletDB), IntToStringDBType(guiDB), IntToStringDBType(txDB))

	// Can adjust starting variables
	// This will also start wallet wsapi
	wal, err := wallet.StartWallet(walletPort, factomdLocation, walletDB, guiDB, txDB, v1Import)
	if err != nil {
		panic("Error in starting wallet: " + err.Error())
	}

	MasterWallet = wal

	// Start Settings
	MasterSettings = new(SettingsStruct)
	data, err := MasterWallet.GUIlDB.Get([]byte("gui-wallet"), []byte("settings"), MasterSettings)
	if err != nil || data == nil {
		err = MasterWallet.GUIlDB.Put([]byte("gui-wallet"), []byte("settings"), MasterSettings)
		if err != nil {
			panic("Error in loading settings: " + err.Error())
		}

		MasterSettings.FactomdLocation = factomdLocation

		// Default dark
		MasterSettings.DarkTheme = true
		MasterSettings.Theme = "darkTheme"
	} else {
		MasterSettings = data.(*SettingsStruct)
		// If we have a custom config file, or a custom flag, we will overwrite the settings.
		// This is so we can still trump the settings in the GUI
		if factomdLocation != "localhost:8088" {
			MasterSettings.FactomdLocation = factomdLocation
		}
		// Here is the first override of the factomd location from the GUI settings.
		// You can see abover, this value will be overwritten by any config or flag
		factomdLocation = MasterSettings.FactomdLocation
	}

	MasterSettings.SetFactomdLocation(factomdLocation)

	MasterSettings.ControlPanelPort = controlPanelPort
	// We always need to load transactions, even if in database. So let's start as not synced
	MasterSettings.Synced = false

	// For Testing adds random addresses
	if ADD_RANDOM_ADDRESSES {
		addRandomAddresses()
	}
	//

	ServeWallet(port)
}

func addRandomAddresses() {
	for i := 0; i < 5; i++ {
		MasterWallet.GenerateEntryCreditAddress("AddedForTesting")
	}

	for i := 0; i < 5; i++ {
		MasterWallet.GenerateFactoidAddress("AddedForTesting")
	}

	MasterWallet.AddAddress("Sand", "Fs3E9gV6DXsYzf7Fqx1fVBQPQXV695eP3k5XbmHEZVRLkMdD9qCK")
}

func IntToStringDBType(t int) string {
	switch t {
	case wallet.MAP:
		return "Map"
	case wallet.LDB:
		return "LDB"
	case wallet.BOLT:
		return "Bolt"
	}
	return "[No DB Type Found]"
}
