package main

import (
	"context"
	"embed"

	"starlight/internal/auth"
	"starlight/internal/db"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	//config goes here
	app := NewApp()
	db := db.NewDB()
	auth := auth.NewAuth(app.ctx, db)

	defer db.SQLite.Close()

	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Starlight",
		Width:  1024,
		Height: 768,
		// Width:  1920,
		// Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
		},
		// OnStartup: app.startup,
		Bind: []interface{}{
			app,
			auth,
			db,
		},
		// OnShutdown: ,
		// OnBeforeClose: ,
		// CSSDragProperty: ,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
