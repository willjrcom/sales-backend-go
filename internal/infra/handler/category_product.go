package handlerimpl

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/willjrcom/sales-backend-go/bootstrap/handler"
	entitydto "github.com/willjrcom/sales-backend-go/internal/infra/dto/entity"
	productdto "github.com/willjrcom/sales-backend-go/internal/infra/dto/product"
	categoryproductusecases "github.com/willjrcom/sales-backend-go/internal/usecases/category_product"
	jsonpkg "github.com/willjrcom/sales-backend-go/pkg/json"
)

type handlerCategoryProductImpl struct {
	pcs *categoryproductusecases.Service
}

func NewHandlerCategoryProduct(productService *categoryproductusecases.Service) *handler.Handler {
	c := chi.NewRouter()

	h := &handlerCategoryProductImpl{
		pcs: productService,
	}

	c.With().Group(func(c chi.Router) {
		c.Post("/new", h.handlerRegisterCategoryProduct)
		c.Put("/update/name/{id}", h.handlerUpdateCategoryProductName)
		c.Put("/update/sizes/{id}", h.handlerUpdateCategoryProductSizes)
		c.Delete("/delete/{id}", h.handlerDeleteCategoryProduct)
		c.Get("/{id}", h.handlerGetCategoryProduct)
		c.Get("/all", h.handlerGetAllCategoryProducts)
	})

	return handler.NewHandler("/category", c)
}

func (h *handlerCategoryProductImpl) handlerRegisterCategoryProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category := &productdto.RegisterCategoryProductInput{}
	jsonpkg.ParseBody(r, category)

	id, err := h.pcs.RegisterCategoryProduct(ctx, category)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("new product: " + id.String()))
}

func (h *handlerCategoryProductImpl) handlerUpdateCategoryProductName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	dtoId := &entitydto.IdRequest{ID: uuid.MustParse(id)}

	category := &productdto.UpdateCategoryProductNameInput{}
	jsonpkg.ParseBody(r, category)

	err := h.pcs.UpdateCategoryProductName(ctx, dtoId, category)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("update product name"))
}

func (h *handlerCategoryProductImpl) handlerUpdateCategoryProductSizes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	dtoId := &entitydto.IdRequest{ID: uuid.MustParse(id)}

	category := &productdto.UpdateCategoryProductSizesInput{}
	jsonpkg.ParseBody(r, category)

	err := h.pcs.UpdateCategoryProductSizes(ctx, dtoId, category)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("update product sizes"))
}

func (h *handlerCategoryProductImpl) handlerDeleteCategoryProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	dtoId := &entitydto.IdRequest{ID: uuid.MustParse(id)}

	err := h.pcs.DeleteCategoryProductById(ctx, dtoId)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("delete category"))
}

func (h *handlerCategoryProductImpl) handlerGetCategoryProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	dtoId := &entitydto.IdRequest{ID: uuid.MustParse(id)}

	category, err := h.pcs.GetCategoryProductById(ctx, dtoId)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	text, err := json.Marshal(category)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(text)
}

func (h *handlerCategoryProductImpl) handlerGetAllCategoryProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	categories, err := h.pcs.GetAllCategoryProduct(ctx)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	text, err := json.Marshal(categories)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(text)
}
