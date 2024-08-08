package rest

import "github.com/gofiber/fiber/v3"

func UnmarshallMultipart(ctx fiber.Ctx) (map[string]any, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}

	values := make(map[string]any)
	for key, value := range form.Value {
		length := len(value)
		switch {
		case length == 0:
			values[key] = nil
		case length == 1:
			var v any = value[0]
			if value[0] == "null" {
				v = nil
			}
			values[key] = v
		case length > 1:
			var vs []any
			for _, v := range value {
				vs = append(vs, v)
			}
			values[key] = vs
		}
	}
	return values, nil
}
