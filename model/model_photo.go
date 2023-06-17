package model

type Photo struct {
	GormModel
	Title    string `gorm:"not null;size:150" json:"title" form:"title" valid:"required~Your title is required"`
	Caption  string `gorm:"size:150" json:"caption" form:"caption" `
	PhotoUrl string `gorm:"not null;size:150" json:"photo_url" form:"photo_url" valid:"required~Your photo_url is required"`
	UserID   int    `gorm:"index" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user"`
}

type PhotoInput struct {
	Title    string `gorm:"not null;size:150" json:"title" form:"title" valid:"required~Your title is required" binding:"required"`
	Caption  string `gorm:"size:150" json:"caption" form:"caption"  binding:"required"`
	PhotoUrl string `gorm:"not null;size:150" json:"photo_url" form:"photo_url" valid:"required~Your photo_url is required" binding:"required"`
}

type GetPhotoInput struct {
	ID int `uri:"id" binding:"required"`
}
