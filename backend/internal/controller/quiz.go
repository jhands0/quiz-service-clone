package controller

import (
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizController struct {
	quizService *service.QuizService
}

func Quiz(quizService *service.QuizService) QuizController {
	return QuizController{
		quizService: quizService,
	}
}

func (c QuizController) GetQuizById(ctx *fiber.Ctx) error {
	quizIdStr := ctx.Params("quizId")
	quizId, err := primitive.ObjectIDFromHex(quizIdStr)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	quiz, err := c.quizService.GetQuizById(quizId)
	if err != nil {
		return err
	}

	if quiz == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(quiz)
}

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	quizzes, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}

	return ctx.JSON(quizzes)
}
