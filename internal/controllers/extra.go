package controllers

import (
	"sync"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/Katsusan/centaur/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	rootUser *models.User
	rootOnce sync.Once
)

//GetRootUser will return models.User who owns root permission.
func GetRootUser() *models.User {
	rootOnce.Do(func() {
		superadmin := config.GetGlobalConfig().SuperAdmin()
		hashedpwd, err := bcrypt.GenerateFromPassword([]byte(superadmin.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln("failed to generate bcrypted password, err=", err)
		}
		rootUser = &models.User{
			Username: superadmin.UserName,
			RealName: superadmin.RealName,
			Password: string(hashedpwd),
		}
	})

	return rootUser
}
