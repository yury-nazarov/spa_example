package ui

// Цель пакета ui обслуживать артифакты frontend

import "embed"

// StaticFiles Встраиваемая файловая система

//go:embed all:dist
var StaticFiles embed.FS
