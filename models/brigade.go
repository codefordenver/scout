package models

import "github.com/jinzhu/gorm"

type Brigade struct {
	// Discord config
	gorm.Model
	ID                        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name                      string `gorm:"not null; unique"`
	DisplayName               string `gorm:"not null;unique"`
	GuildID                   string `gorm:"not null;unique"`
	ActiveProjectCategoryID   string `gorm:"not null;unique"`
	InactiveProjectCategoryID string `gorm:"not null;unique"`
	NewUserRole               string `gorm:"not null;unique"`
	OnboardingRole            string `gorm:"not null;unique"`
	MemberRole                string `gorm:"not null;unique"`
	OnboardingInviteCode      string `gorm:"not null;unique"`
	OnboardingInviteCount     int    `gorm:"not null"`
	CodeOfConductMessageID    string `gorm:"not null;unique"`
	// GDrive Config
	AgendaFolderID string `gorm:"not null;unique"`
	TimezoneString string `gorm:"not null"`
	// Github Config
	GithubOrganization string `gorm:"not null;unique"`
	IssueEmoji         string `gorm:"not null"`
}
