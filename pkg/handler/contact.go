package handler

import (
	"net/http"
	"strconv"

	phonebook "github.com/DanYesmagulov/go-phone-book"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createContact(c *gin.Context) {
	var input phonebook.Contact

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Contacts.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllContactsResponse struct {
	Data []phonebook.Contact `json:"data"`
}

func (h *Handler) getAllContacts(c *gin.Context) {
	contacts, err := h.services.Contacts.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllContactsResponse{
		Data: contacts,
	})
}

func (h *Handler) getContactById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильный параметр id")
		return
	}
	contact, err := h.services.Contacts.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *Handler) updateContactById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильный параметр id")
		return
	}

	var input phonebook.UpdateContact

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Contacts.UpdateById(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

func (h *Handler) deleteContactById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильный параметр id")
		return
	}

	err = h.services.Contacts.DeleteById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
