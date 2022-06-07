package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"strconv"
)

type requestSections struct {
	Id                 uint64  `json:"id"`
	Number             uint64  `json:"number" binding:"required"`
	CurrentTemperature float32 `json:"current_temperature" binding:"required"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    uint32  `json:"current_capacity" binding:"required"`
	MinimumCapacity    uint32  `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    uint32  `json:"maximum_capacity" binding:"required"`
	WarehouseId        uint64  `json:"warehouse_id" binding:"required"`
	ProductTypeId      uint64  `json:"product_type_id" binding:"required"`
}

type sectionController struct {
	sectionService sections.SectionService
}

func NewController(s sections.SectionService) *sectionController {
	return &sectionController{
		sectionService: s,
	}
}

func (c *sectionController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		sections, err := c.sectionService.GetAll()

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, sections, ""))
	}
}

func (c *sectionController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		section, err := c.sectionService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, section, ""))
	}
}

func (c *sectionController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestSections

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		sections, err := c.sectionService.GetAll()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, v := range sections {
			if v.Number == req.Number {
				ctx.JSON(409, web.NewResponse(409, nil, "Section number already existis"))
				return
			}
		}

		addedSection, err := c.sectionService.Create(c.sectionService.GetNextId(), req.Number, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseId, req.ProductTypeId)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, addedSection, ""))
	}

}

func (c *sectionController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req requestSections

		err = ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = c.sectionService.Get(uint64(id))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		updatedSection, err := c.sectionService.Update(uint64(id), req.Number, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseId, req.ProductTypeId)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, updatedSection, ""))
	}

}

func (c *sectionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = c.sectionService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.sectionService.Delete(uint64(id))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(204, web.NewResponse(204, nil, ""))
	}
}
