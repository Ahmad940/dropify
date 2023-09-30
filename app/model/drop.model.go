package model

type Drop struct {
	ID string `json:"id" gorm:"primaryKey; type:varchar; not null; unique"`

	Name string `json:"name" gorm:"type:varchar; varchar; not null; index" validate:"required"`

	Repository string `json:"repository" gorm:"type:varchar; not null" validate:"required"`
	Branch     string `json:"branch" gorm:"type:varchar; not null" validate:"required"`

	Prefix         string `json:"prefix" gorm:"type:varchar; not null" validate:"required"`
	BuildCommand   string `json:"build_command" gorm:"type:varchar; not null" validate:"required"`
	InstallCommand string `json:"install_command" gorm:"type:varchar; not null" validate:"required"`
	StartCommand   string `json:"start_command" gorm:"type:varchar; not null" validate:"required"`

	AutoDeploy bool `json:"auto_deploy" gorm:"not null; default:false"`

	Gorm
}
