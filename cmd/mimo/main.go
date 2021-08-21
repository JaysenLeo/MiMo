package main

import (
	"encoding/json"
	"time"

	//"encoding/json"
	"flag"
	"fmt"
	"github.com/asticode/go-astikit"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"log"
	"os"
	//"time"

	"github.com/asticode/go-astilectron"
	//bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

var (
	mainWindowHeight = 1080
	mainWindowWidth  = 1920
)

// Constants
const htmlAbout = `Welcome on <b>Astilectron</b> demo!<br>
This is using the bootstrap and the bundler.`

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	fs    = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	debug = fs.Bool("d", true, "enables the debug mode")
	w     *astilectron.Window
)

func main() {
	//// Set logger
	//l := log.New(log.Writer(), log.Prefix(), log.Flags())
	//
	//// Create astilectron
	//a, err := astilectron.New(l, astilectron.Options{
	//	AppName:           "Test",
	//	BaseDirectoryPath: "example",
	//})
	//if err != nil {
	//	l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	//}
	//defer a.Close()
	//
	//// Handle signals
	//a.HandleSignals()
	//a.SetProvisioner(astilectron.NewDisembedderProvisioner())
	//
	//// Start
	//if err = a.Start(); err != nil {
	//	l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	//}
	//
	//// New window
	//var w *astilectron.Window
	//if w, err = a.NewWindow("./static/main/index.html", &astilectron.WindowOptions{
	//	Center: astikit.BoolPtr(true),
	//	Height: astikit.IntPtr(mainWindowHeight),
	//	Width:  astikit.IntPtr(mainWindowWidth),
	//}); err != nil {
	//	l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	//}
	//
	//// Create windows
	//if err = w.Create(); err != nil {
	//	l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	//}
	////w.OpenDevTools()
	//// Blocking pattern
	//a.Wait()
	//// Create logger
	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Parse flags
	_ = fs.Parse(os.Args[1:])

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		//Asset:    Asset,
		//AssetDir: AssetDir,
		ResourcesPath: "main",
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "icon.icns",
			AppIconDefaultPath: "icon.png",
			SingleInstance:     true,
			//SkipSetup:          true,
			DataDirectoryPath: "resources",
			BaseDirectoryPath:  "resources",
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: l,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("About"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
							// Unmarshal payload
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								l.Println(fmt.Errorf("unmarshaling payload failed: %w", err))
								return
							}
							l.Printf("About modal has been displayed and payload is %s!\n", s)
						}); err != nil {
							l.Println(fmt.Errorf("sending about event failed: %w", err))
						}
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(5 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					l.Println(fmt.Errorf("sending check.out.menu event failed: %w", err))
				}
			}()
			return nil
		},
		//RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage: "index.html",
			//MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				Center: astikit.BoolPtr(true),
				Height: astikit.IntPtr(mainWindowHeight),
				Width:  astikit.IntPtr(mainWindowWidth),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}
}
