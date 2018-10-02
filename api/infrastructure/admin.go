package infrastructure

import (
	"api/domain"
	"api/service"
	"log"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
)

func defineUserMetaInfo(user *admin.Resource) {
	user.NewAttrs("FirstName", "LastName", "NickName", "Email", "Role", "Password")
	user.EditAttrs("FirstName", "LastName", "NickName", "Email", "InvitationToken", "Role")
	user.IndexAttrs("FirstName", "LastName", "NickName", "Email", "Role")
	user.Meta(&admin.Meta{
		Name:   "Password",
		Type:   "password",
		Valuer: func(interface{}, *qor.Context) interface{} { return "" },
		Setter: encryptPassword,
	})
	user.Action(&admin.Action{
		Name:    "ResetPassword",
		Handler: generateResetPasswordToken,
		Visible: tokenExist,
		Modes:   []string{"edit", "show", "collection", "menu_item"},
	})
}

func tokenExist(record interface{}, context *admin.Context) bool {
	user := record.(*domain.User)
	if user.ResetPasswordToken == "" {
		return true
	}
	return false
}

func generateResetPasswordToken(argument *admin.ActionArgument) error {
	for _, record := range argument.FindSelectedRecords() {
		token, err := service.GenerateToken()
		if err != nil {
			log.Fatal(err)
			return err
		}
		argument.Context.DB.Model(record.(*domain.User)).Update("ResetPasswordToken", token)
	}
	return nil
}

func encryptPassword(record interface{}, metaValue *resource.MetaValue, context *qor.Context) {
	if password := utils.ToString(metaValue.Value); password != "" {
		passwordHash := service.ToHash(password)
		record.(*domain.User).EncryptedPassword = passwordHash
	}
}
