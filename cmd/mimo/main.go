package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	"github.com/gin-gonic/gin"
	"log"
	"mimo/internal/app/riro/cfg"
	"mimo/internal/app/riro/service/index"
	"os"
	"time"
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

	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Parse flags
	_ = fs.Parse(os.Args[1:])
	go RunRestfulServer()
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
			MessageHandler: HandleMessages,
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
type KnowledgeDetailParams struct {
	Kid   int64 `uri:"kid" binding:"required"`
}
func RunRestfulServer() {
	r := gin.Default()
	r.GET("/knowledge/:kid", func(c *gin.Context) {
		var knowledgeDetailParams KnowledgeDetailParams
		if err := c.ShouldBindUri(&knowledgeDetailParams); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		riro := index.NewRiRo(func() cfg.Option {
			opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository" })
			return *opt
		})
		c.JSON(200, gin.H{
			"Data": riro.ListAllLocalOneKnowledgeFile(knowledgeDetailParams.Kid),
			"Status": "knowledge.GotInfoOK",
		})
	})
	_ = r.Run(":8088")
}

// handleMessages handles messages
func HandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	p, _:=m.Payload.MarshalJSON()
	fmt.Println(m.Name, string(p))
	switch m.Name {
	case "hello":
		// Unmarshal payload
		return struct{
			Data string `json:"data"`
		}{ Data: "world" }, nil
	case "knowledge":
		riro := index.NewRiRo(func() cfg.Option {
			opt := cfg.NewRiRoOption([]string{"E:\\DevHub\\MiMo\\repository" })
			return *opt
		})
		return struct{
			Data []map[string]interface{} `json:"data"`
		}{ Data: riro.ListAllLocalKnowledge() } , nil
	}
	return
}

