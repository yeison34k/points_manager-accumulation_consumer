package app

import (
	"fmt"
)

type PointHandler struct {
	PointApp *PointApp
}

func NewPointHandler(pointApp *PointApp) *PointHandler {
	return &PointHandler{
		PointApp: pointApp,
	}
}

func (h *PointHandler) HandlePointCreation() error {
	err := h.PointApp.CreatePoint()
	if err != nil {
		return fmt.Errorf("failed to create point: %w", err)
	}

	return nil
}
