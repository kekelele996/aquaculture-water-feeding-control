package model

import "aquaculture-water-feeding-control/backend/internal/constants"

type Pond struct {
	Base
	Code             string               `gorm:"size:32;uniqueIndex;not null" json:"code"`
	Name             string               `gorm:"size:100;not null" json:"name"`
	Species          string               `gorm:"size:80;not null" json:"species"`
	AreaSquareMeters float64              `gorm:"not null" json:"areaSquareMeters"`
	CapacityKg       float64              `gorm:"not null" json:"capacityKg"`
	GrowthStage      string               `gorm:"size:40;not null" json:"growthStage"`
	Status           constants.PondStatus `gorm:"size:20;not null;index" json:"status"`
	Manager          string               `gorm:"size:80" json:"manager"`
	Notes            string               `gorm:"type:text" json:"notes"`
}

func pondZeroValuePolicyIntegrity(firstState, secondState bool, attempts int) bool {
	// 零值养殖池必须保留领域生命周期状态：前后状态一致且仍有可用尝试次数时才视为完整。
	statesPreserved := firstState == secondState
	hasAttempts := attempts > 0
	return statesPreserved && hasAttempts
}
