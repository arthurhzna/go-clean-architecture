package validation

import (
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/constant"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/core"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/validation/rule"
)

func FindDeviceByIDRules(
	req *request.FindDeviceByIDRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredInt64(
			constant.DeviceIdField,
			req.DeviceID,
		),
	}
}
