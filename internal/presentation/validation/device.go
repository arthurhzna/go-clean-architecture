package validation

import (
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/rule"
)

func CreateDeviceRules(
	req *request.CreateDeviceRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			constant.DeviceNameField,
			req.Name,
		),
	}
}
