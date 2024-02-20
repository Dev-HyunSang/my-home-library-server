package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/dev-hyunsang/my-home-library-server/auth"
	"github.com/dev-hyunsang/my-home-library-server/db"
	"github.com/dev-hyunsang/my-home-library-server/dto"
	"github.com/dev-hyunsang/my-home-library-server/ent/book"
	"github.com/dev-hyunsang/my-home-library-server/ent/user"
	"github.com/dev-hyunsang/my-home-library-server/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateBookHandler(ctx *fiber.Ctx) error {
	req := new(dto.RequestUpdateBook)
	if err := ctx.BodyParser(req); err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusBadRequest,
				Message:    "올바르지 않은 요청 방식으로 요청을 했어요. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusUnauthorized,
				Message:    "로그인을 하지 않은 상태로 요청을 주셨어요. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID, err := auth.FetchAuth(token)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusUnauthorized,
				Message:    "로그인 유효시간이 지났어요. 다시 로그인을 하신 다음 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	logger.Info(fmt.Sprintf("Requested User By %s", userUUID))

	client, err := db.ConnectMySQL()
	if err != nil {
		logger.Error(err.Error())
	}

	userData, err := client.User.Query().Where(user.ID(userUUID)).First(ctx.Context())
	if err != nil {
		logger.Error(err.Error())
	}

	updatedBook, err := client.Book.Create().
		SetID(uuid.New()).
		SetUser(userData).
		SetTitle(req.Title).
		SetSubtitle(req.SubTitle).
		SetPublisher(req.Publisher).
		SetPublishingCompany(req.PublishingCompany).
		SetMemo(req.Memo).
		SetTotalPage(req.TotalPage).
		SetCurrentPage(req.CurrentPage).
		SetCreatedAt(time.Now()).
		SetEditedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusInternalServerError,
				Message:    "데이터베이스에 사용자의 정보를 저장하던 도중 오류가 발생했습니다. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseUpdateBook{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "성공적으로 책을 등록 했어요! 열심히 읽어봐요!",
		},
		Data:        updatedBook,
		ResponsedAt: time.Now(),
	})
}

func GetBookHandler(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusUnauthorized,
				Message:    "로그인을 하지 않은 상태로 요청을 주셨어요. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		logger.Info(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	parseBookID, err := uuid.Parse(bookID)
	if err != nil {
		logger.Error(err.Error())
	}

	result, err := client.Book.Query().
		Where(book.ID(parseBookID)).
		Where(book.UserUUID(token.UserID)).
		First(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusInternalServerError,
				Message:    "사용자의 소중한 정보를 가져오던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseGetBook{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "성공적으로 사용자의 소중한 책 정보를 불러왔어요.",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func GetBooksHandler(ctx *fiber.Ctx) error {
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusUnauthorized,
				Message:    "로그인을 하지 않은 상태로 요청을 주셨어요. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	result, err := client.Book.Query().Where(book.UserUUID(token.UserID)).All(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusInternalServerError,
				Message:    "사용자의 정보를 찾던 도중 동일한 사용자의 정보를 찾을 수 없어요. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseGetBooks{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "성공적으로 소중한 사용자의 책 데이터를 불러왔어요!",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func EditBookHandler(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	req := new(dto.RequestUpdateBook)
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusUnauthorized,
				Message:    "로그인을 하지 않은 상태로 요청을 주셨어요. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusBadRequest,
				Message:    "올바르지 않은 요청입니다. 확인 후 다시 요청해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	parseBookID, err := uuid.Parse(bookID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{})
	}

	client.Book.Update().Where(book.UserUUID(token.UserID)).Where(book.ID(parseBookID)).
		SetTitle(req.Title).
		SetSubtitle(req.SubTitle).
		SetPublisher(req.Publisher).
		SetPublishingCompany(req.PublishingCompany).
		SetMemo(req.Memo).
		SetEditedAt(time.Now()).Save(context.Background())

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseUpdateBook{})
}

func DeleteBookHandler(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	authToken := ctx.GetReqHeaders()["Authorization"][0]

	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusUnauthorized,
				Message:    "로그인을 하지 않은 상태로 요청을 주셨어요. 확인 후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusInternalServerError,
				Message:    "데이터베이스에 연결을 시도할 수 없습니다. 잠시후 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	parseBookUUID, err := uuid.Parse(bookID)
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusInternalServerError,
				Message:    "소중한 정보를 처리하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	result, err := client.Book.Delete().
		Where(book.ID(parseBookUUID)).
		Where(book.UserUUID(token.UserID)). // 요청하는 사용자의 UUID를 확인함. 보안요소
		Exec(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseErr{
			Status: dto.ErrStatus{
				Code:       fiber.StatusInternalServerError,
				Message:    "사용자의 소중한 정보를 제거하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
				ErrMessage: err.Error(),
			},
		})
	}

	logger.Info(fmt.Sprintf("Deleted By %d", result))

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseBook{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Message: "책을 성공적으로 삭제했어요.",
		},
		ResponsedAt: time.Now(),
	})
}
