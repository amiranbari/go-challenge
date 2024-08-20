package userservice

import (
	"context"
	"fmt"
	userparam "github.com/amiranbari/challenge/param/user"
	richerror "github.com/amiranbari/challenge/rich_error"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"time"
)

func (s Service) GetAll(ctx context.Context, req userparam.GetAllRequest) (userparam.GetAllResponse, error) {

	const op = "userservice.GetAll"
	if fieldErrors, vErr := s.vld.ValidateGetAll(req); vErr != nil {
		return userparam.GetAllResponse{FieldErrors: fieldErrors}, richerror.New(op).WithErr(vErr)
	}

	users, err := s.repo.GetAllUsers(ctx, req.Filter)
	if err != nil {
		return userparam.GetAllResponse{}, richerror.New(op).WithErr(err)
	}

	f := excelize.NewFile()

	sheetName := "Users"
	f.NewSheet(sheetName)

	headers := []string{"Name", "Family"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	for i, user := range users {
		rowNum := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNum), user.FirstName)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNum), user.LastName)
	}

	dir := "exports"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return userparam.GetAllResponse{Users: users}, richerror.New(op).WithErr(err)
	}

	fileName := fmt.Sprintf("users_%d.xlsx", time.Now().Unix())
	filePath := filepath.Join(dir, fileName)

	if err := f.SaveAs(filePath); err != nil {
		return userparam.GetAllResponse{Users: users}, richerror.New(op).WithErr(err)
	}

	downloadLink := fmt.Sprintf("http://localhost:1314/exports/%s", fileName)

	return userparam.GetAllResponse{
		Link:  downloadLink,
		Users: users,
	}, nil
}
