package app

type PointService interface {
	CreatePoint() error
}

type PointApp struct {
	PointService PointService
}

func NewPointApplication(pointService PointService) *PointApp {
	return &PointApp{
		PointService: pointService,
	}
}

func (a *PointApp) CreatePoint() error {
	return a.PointService.CreatePoint()
}
