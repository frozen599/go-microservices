package controllers

import (
	"net/http"
	"time"

	"github.com/frozen599/gin-microserices/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func NewRecipe(c *gin.Context) {
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	models.Recipes = append(models.Recipes, recipe)
	c.JSON(http.StatusOK, &recipe)
}

func GetAllRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, &models.Recipes)
}

func UpdateRecipe(c *gin.Context) {
	id := c.Param("id")
	var recipe models.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1
	for i := 0; i < len(models.Recipes); i++ {
		if models.Recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
	}

	recipe = models.Recipes[index]
	c.JSON(http.StatusOK, &recipe)
}

func DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	index := -1

	for i := 0; i < len(models.Recipes); i++ {
		if models.Recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
	}

	models.Recipes = append(models.Recipes[:index], models.Recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted",
	})
}
