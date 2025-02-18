package service

import (
	"backend/internal/collection"
	"backend/internal/entity"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) UpdateQuiz(id primitive.ObjectID, name string, questions []entity.QuizQuestion) error {
	quiz, err := s.quizCollection.GetQuizById(id)
	if err != nil {
		return err
	}

	if quiz == nil {
		return errors.New("quiz not found")
	}

	quiz.Name = name
	quiz.Questions = questions
	return s.quizCollection.UpdateQuiz(*quiz)
}

func (s QuizService) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	return s.quizCollection.GetQuizById(id)
}

func (s QuizService) GetQuizzes() (*[]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
