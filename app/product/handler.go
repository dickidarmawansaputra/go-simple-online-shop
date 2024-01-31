package product

import (
	"net/http"

	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/infrafiber"
	"github.com/dickidarmawansaputra/go-simple-online-shop/infra/response"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	var req = ProductRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("invalid payload"),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create product success"),
	).Send(ctx)
}

func (h handler) GetListProduct(ctx *fiber.Ctx) error {
	var req = ListProductRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	products, err := h.svc.ListProduct(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("invalid payload"),
		).Send(ctx)
	}

	productListResponse := NewProductListResponseFromEntity(products)

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get list products success"),
		infrafiber.WithPayload(productListResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetProductDetail(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku")
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	product, err := h.svc.ProductDetail(ctx.UserContext(), sku)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorNotFound
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("invalid payload"),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get product detail success"),
		infrafiber.WithPayload(product.ToProductDetailResponse()),
	).Send(ctx)
}
