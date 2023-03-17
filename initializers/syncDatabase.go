package initializers

import "github.com/twk-mn/goJWT/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}) // Generates an empty Table, and no need to use "Initializers." as we are already there..
}
