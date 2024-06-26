package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DeployEvent(event string) {
	// Get the message
	runtime.EventsEmit(a.ctx, "event", event)
	SendFromParam(a.ctx, "Hello from param")
	// Handle other things
}

func SendFromParam(ctx context.Context, message string) {
	runtime.EventsEmit(ctx, "event", message)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
