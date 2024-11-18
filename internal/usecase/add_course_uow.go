package usecase

import (
	"context"

	"github.com/Genarodaniel/go-uow/internal/entity"
	"github.com/Genarodaniel/go-uow/internal/repository"
	"github.com/Genarodaniel/go-uow/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		//tudo que colocar aqui será feito um begin e rollback;
		category := entity.Category{
			Name: input.CategoryName,
		}

		err := a.getCategoryRepository(ctx).Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		err = a.getCourseRepository(ctx).Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CourseRepositoryInterface)
}
