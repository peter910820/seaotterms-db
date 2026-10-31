package discordbot

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Member struct {
	UserID     string    `gorm:"primaryKey"`
	ServerID   string    `gorm:"primaryKey"`
	UserName   string    `gorm:"not null"`
	Level      uint      `gorm:"not null;default:1"`
	Exp        uint      `gorm:"not null;default:0"` // 該等級的經驗值，加上LevelUpExp才是該成員的所有經驗值
	LevelUpExp uint      `gorm:"not null;default:5"`
	JoinAt     time.Time `gorm:"not null"`
	UpdatedAt  time.Time
}

func GetServerTopMembersByExp(db *gorm.DB, serverID string, limit int) ([]Member, error) {
	var memberData []Member
	err := db.Where("server_id = ?", serverID).Order("exp DESC").Limit(limit).Find(&memberData).Error
	if err != nil {
		return nil, err
	}

	return memberData, nil
}

func QueryAllMembers(db *gorm.DB) ([]Member, error) {
	var UserData []Member

	err := db.Find(&UserData).Error
	if err != nil {
		return nil, err
	}
	return UserData, nil
}

// query single member fo database use userID
func QueryMemberByUserID(db *gorm.DB, userID string) (Member, error) {
	var memberData Member
	err := db.Where("user_id = ?", userID).First(&memberData).Error
	if err != nil {
		return memberData, err
	}
	return memberData, nil
}

func CreateMember(db *gorm.DB, member Member) error {
	data := Member{
		UserID:   member.UserID,
		ServerID: member.ServerID,
		UserName: member.UserName,
		JoinAt:   member.JoinAt,
	}

	err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateMemberLevel(db *gorm.DB, userID string, member Member) error {
	data := Member{
		Level:      member.Level,
		Exp:        member.Exp,
		LevelUpExp: member.LevelUpExp,
		UpdatedAt:  member.UpdatedAt,
	}

	err := db.Model(&Member{}).Where("user_id = ?", userID).
		Select("level", "exp", "level_up_exp", "updated_at").Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}
