package horizon

func initTangCoreInfo(app *App) {
	app.UpdateTangCoreInfo()
	return
}

func init() {
	appInit.Add("tangCoreInfo", initTangCoreInfo, "app-context", "log")
}
