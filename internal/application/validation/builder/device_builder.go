package builder

import (
	"github.com/arthurhzna/go-clean-architecture/internal/application/dto/request"

	"github.com/arthurhzna/go-clean-architecture/internal/application/validation/field"
	"github.com/arthurhzna/go-clean-architecture/internal/application/validation/rule"
)

func FindDeviceByIDRules(
	req request.FindDeviceByIDRequest,
) []rule.Rule {

	return []rule.Rule{
		rule.RequiredInt64(
			field.DeviceIdField,
			req.ID,
		),
	}
}
