package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/output"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/parameter"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/response"
	"github.com/haradayoshitsugucz/purple-server/api/service"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/constant"
	"github.com/haradayoshitsugucz/purple-server/infra/dao"
	"github.com/haradayoshitsugucz/purple-server/logger"
	"go.uber.org/zap"
)

type ProductController interface {
	GetProduct(config config.Config) http.HandlerFunc
	GetProductsByName(conf config.Config) http.HandlerFunc
	AddProduct(conf config.Config) http.HandlerFunc
	EditProduct(conf config.Config) http.HandlerFunc
	DeleteProduct(conf config.Config) http.HandlerFunc
}

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewProductController(
	productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		productService: productService,
	}
}

func (c *ProductControllerImpl) GetProduct(conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// パスパラメーター
		productID, err := strconv.ParseInt(chi.URLParam(r, "product_id"), 10, 64)
		if err != nil {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		// validate
		param := parameter.Product{ProductID: productID}
		if ok, err := param.Valid(); !ok {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		// GetProduct
		product, err := c.productService.GetProduct(param.ProductID)
		if err != nil {
			response.ErrorWith(w, http.StatusInternalServerError, constant.Err500001, err.Error())
			return
		}
		if ok, err := product.Product.Empty(); ok {
			response.ErrorWith(w, http.StatusNotFound, constant.Err404001, err.Error())
			return
		}

		response.SuccessWith(w, output.ToProduct(product))
		return
	}
}

func (c *ProductControllerImpl) GetProductsByName(conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// クエリーパラメーター
		query := r.URL.Query()
		param := parameter.NewProductsByName(
			query.Get("name"),
			query.Get("offset"),
			query.Get("limit"),
		)

		// validate
		if ok, err := param.Valid(); !ok {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		products, total, err := c.productService.GetProductsByName(param.Name, param.Offset, param.Limit)
		if err != nil {
			response.ErrorWith(w, http.StatusInternalServerError, constant.Err500001, err.Error())
			return
		}
		out := output.ToProducts(products)
		response.SuccessWith(w, output.ToPager(total, param.Offset, param.Limit, out))

		return
	}
}

func (c *ProductControllerImpl) AddProduct(conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		param := &parameter.AddProduct{}
		err := bindJson(r, param)
		if err != nil {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		if ok, err := param.Valid(); !ok {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		productID, err := c.productService.AddProduct(param.Name, param.BrandID)

		logger.GetLogger().Debug("AddProduct", zap.Int64("productID", productID))

		if _, ok := err.(*dao.DuplicateError); ok {
			response.ErrorWith(w, http.StatusConflict, constant.Err409001, err.Error())
			return
		}
		if err != nil {
			response.ErrorWith(w, http.StatusInternalServerError, constant.Err500001, err.Error())
			return
		}

		response.SuccessWith(w, map[string]int64{"product_id": productID})
		return
	}
}

func (c *ProductControllerImpl) EditProduct(conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		param := &parameter.EditProduct{}
		err := bindJson(r, param)
		if err != nil {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		// パスパラメーター
		productID, err := strconv.ParseInt(chi.URLParam(r, "product_id"), 10, 64)
		if err != nil {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}
		param.ProductID = productID

		if ok, err := param.Valid(); !ok {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		_, err = c.productService.EditProduct(param.ProductID, param.Name, param.BrandID)
		if err != nil {
			response.ErrorWith(w, http.StatusInternalServerError, constant.Err500001, err.Error())
			return
		}

		response.Success(w)
		return
	}
}

func (c *ProductControllerImpl) DeleteProduct(conf config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// パスパラメーター
		productID, err := strconv.ParseInt(chi.URLParam(r, "product_id"), 10, 64)
		if err != nil {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		// validate
		param := parameter.Product{ProductID: productID}
		if ok, err := param.Valid(); !ok {
			response.ErrorWith(w, http.StatusBadRequest, constant.Err400001, err.Error())
			return
		}

		_, err = c.productService.DeleteProduct(param.ProductID)
		if err != nil {
			response.ErrorWith(w, http.StatusInternalServerError, constant.Err500001, err.Error())
			return
		}

		response.Success(w)
		return
	}
}
