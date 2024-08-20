package userservice

import (
	"context"
	"errors"
	"fmt"
	userparam "github.com/amiranbari/challenge/param/user"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"time"
)

func (s Service) GetAll(ctx context.Context, req userparam.GetAllRequest) (userparam.GetAllResponse, error) {

	users, err := s.repo.GetAllUsers(ctx, req.Filter)
	if err != nil {
		return userparam.GetAllResponse{}, errors.New("error in export excel")
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
		return userparam.GetAllResponse{Users: users}, err
	}

	fileName := fmt.Sprintf("users_%d.xlsx", time.Now().Unix())
	filePath := filepath.Join(dir, fileName)

	if err := f.SaveAs(filePath); err != nil {
		return userparam.GetAllResponse{Users: users}, err
	}

	downloadLink := fmt.Sprintf("http://localhost:1314/exports/%s", fileName)

	return userparam.GetAllResponse{
		Link:  downloadLink,
		Users: users,
	}, nil
}
