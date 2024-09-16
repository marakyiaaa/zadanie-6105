package handlers

import (
	"main/internal/models"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProposalHandler struct {
	service service.ProposalService
}

func NewProposalHandler(service service.ProposalService) *ProposalHandler {
	return &ProposalHandler{service: service}
}

func (h *ProposalHandler) CreateProposal(c *gin.Context) {
	var proposal models.Proposal
	if err := c.ShouldBindJSON(&proposal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProposal(c.Request.Context(), &proposal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, proposal)
}

func (h *ProposalHandler) GetProposalByID(c *gin.Context) {
	proposalID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid proposal ID"})
		return
	}
	proposal, err := h.service.GetProposalByID(c.Request.Context(), proposalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, proposal)
}

func (h *ProposalHandler) UpdateProposal(c *gin.Context) {
	var proposal models.Proposal
	if err := c.ShouldBindJSON(&proposal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateProposal(c.Request.Context(), &proposal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, proposal)
}

func (h *ProposalHandler) DeleteProposal(c *gin.Context) {
	proposalID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid proposal ID"})
		return
	}
	if err := h.service.DeleteProposal(c.Request.Context(), proposalID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *ProposalHandler) ListProposalsByTenderID(c *gin.Context) {
	tenderID, err := strconv.ParseInt(c.Param("tenderID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tender ID"})
		return
	}
	proposals, err := h.service.ListProposalsByTenderID(c.Request.Context(), tenderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, proposals)
}
