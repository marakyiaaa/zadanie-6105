package handlers

import (
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TenderHandler struct {
	service service.TenderService
}

func NewTenderHandler(service service.TenderService) *TenderHandler {
	return &TenderHandler{service: service}
}

func (h *TenderHandler) CreateTender(c *gin.Context) {
	var tender models.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateTender(c.Request.Context(), &tender); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tender)
}

func (h *TenderHandler) PublishTender(c *gin.Context) {
	tenderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tender ID"})
		return
	}
	if err := h.service.PublishTender(c.Request.Context(), tenderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "published"})
}

func (h *TenderHandler) CloseTender(c *gin.Context) {
	tenderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tender ID"})
		return
	}
	if err := h.service.CloseTender(c.Request.Context(), tenderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "closed"})
}

func (h *TenderHandler) UpdateTender(c *gin.Context) {
	var tender models.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateTender(c.Request.Context(), &tender); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tender)
}

func (h *TenderHandler) GetTenderByID(c *gin.Context) {
	tenderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tender ID"})
		return
	}
	tender, err := h.service.GetTenderByID(c.Request.Context(), tenderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tender)
}
