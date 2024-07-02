package service

import (
	"context"
	"mime/multipart"
	"mywebmall/dao"
	"mywebmall/model"
	"mywebmall/pkg/e"
	"mywebmall/pkg/util"
	"mywebmall/serializer"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"` //多张图片选第一张作为展示图
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

// 创建商品
func (service *ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	code := e.Success
	var boss *model.User
	var err error
	// 批量检查上传的文件是否符合图片格式
	for _, fh := range files {
		ok, _ := CheckPhotoType(fh)
		if !ok {
			code = e.ErrorFileType
			return serializer.Response{
				Status: e.ErrorFileType,
				Msg:    e.GetMsg(e.ErrorFileType),
				Data:   "上传文件类型错误!请只上传jpeg,jpg,png的图片!",
			}
		}
	}

	userDao := dao.NewUserDao(ctx)
	// 获取创建商品的用户
	boss, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("ProductService Create func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        boss.ID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	// 先创建这个商品
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("ProductService Create func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 以第一张图作为封面图
	_, fileType := CheckPhotoType(files[0])
	tmp, err := files[0].Open()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("ProductService Create func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 再设置商品的图片并进行更新
	pid := strconv.Itoa(int(product.ID))
	// 在对应的用户下面创建对应的商品图片文件夹，专门存这个商品的
	path, err := UploadProductToLocalStatic(tmp, uid, service.Name+pid, "0", fileType)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("ProductService Create func error,image upload error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 对商品图路径更新
	product.ImgPath = path
	err = productDao.UpdateProduct(product.ID, product)
	if err != nil {
		code = e.ErrorDataBase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	for i, file := range files {
		if i == 0 {
			continue
		}
		num := strconv.Itoa(i)
		pIDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ := file.Open()
		_, ft := CheckPhotoType(file)
		// 上传商品图，根据的是第几张图片，这样就不会重复了
		tpath, err := UploadProductToLocalStatic(tmp, uid, service.Name+pid, num, ft)
		if err != nil {
			code = e.ErrorProductImgUpload
			util.LogrusObj.Infoln("ProductService Create func error,image upload error", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductId: product.ID,
			ImgPath:   tpath,
		}
		// 在数据库更新
		err = pIDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.ErrorDataBase
			util.LogrusObj.Infoln("ProductService Create func error:", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	util.LogrusObj.Infoln("New product:", serializer.BuildProduct(product))
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

// 商品列表
func (service *ProductService) List(ctx context.Context) serializer.Response {
	code := e.Success
	var products []*model.Product
	// page参数为0则初始化一下
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if service.PageNum == 0 {
		service.PageNum = 1
	}
	// 商品检索条件
	condition := make(map[string]interface{})
	// 有分类则显示该分类的商品
	if service.CategoryId != 0 {
		condition["category_id"] = int(service.CategoryId)
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("ProductService List func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 不知道为什么用多goroutine的模式
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		tmpPDao := dao.NewProductDaoByDB(productDao.DB)
		products, _ = tmpPDao.ListProductByCondition(condition, service.BasePage)
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// 搜索商品
func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if service.PageNum == 0 {
		service.PageNum = 1
	}
	productDao := dao.NewProductDao(ctx)
	products, err := productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("ProductService Search func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(len(products)))
}

// 展示某个商品
func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pid, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pid))
	if err != nil {
		code = e.ErrorDataBase
		util.LogrusObj.Infoln("ProductService Show func error:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
