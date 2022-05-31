package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"

	"myapp/internal/dao"
	"myapp/internal/model"
	"myapp/internal/mongodb"
	"myapp/internal/response"
	"myapp/internal/util/helper"
	"myapp/internal/util/pstring"
	"myapp/internal/util/ptime"
	"myapp/internal/util/query"
	apimodel "myapp/pkg/admin/model/api"
	"myapp/pkg/admin/model/response"

	"go.mongodb.org/mongo-driver/bson"
)

type Staff struct{}

func (s Staff) Create(ctx context.Context, payload apimodel.StaffCreate) (result *responsemodel.ResponseCreate, err error) {
	var (
		roleDAO  = dao.Role{}
		staffDAO = dao.Staff{}
	)
	//find Role
	roleID, isValid := mongodb.NewIDFromString(payload.Role)
	if !isValid {
		return nil, errors.New(response.CommonInvalidRole)
	}
	if role := roleDAO.FindByID(ctx, roleID); role.ID.IsZero() {
		return nil, errors.New(response.CommonNotFoundRole)
	}
	//check email Ecited or not
	if s.isEmailExisted(ctx, payload.Email) {
		return nil, errors.New(response.CommonExistedEmail)
	}

	//Generate password
	passwordLength := 8
	password := pstring.GenerateDefaultPassword(passwordLength)

	//new Staff
	staff := payload.NewStaffRaw(roleID, helper.HashPassword(password))
	if err = staffDAO.InsertOne(ctx, staff); err != nil {
		return nil, errors.New(response.CommonErrorWhenHandler)
	}
	//go s.sendPasswordToEmail(payload.Name,payload.Email,password)
	return &responsemodel.ResponseCreate{ID: staff.ID.Hex()}, nil
}

/*
func (s Staff)sendPasswordToEmail(toName, toAddress, password string)error{
	var (envVars =config.GetENV()
		contentHTML bytes.Buffer
	)
	html,err :=htmltemplate.New("createStaff.html").Parse(constant.TemplateHTMLNewUser)
	if err!=nil{
		logger.Error("New template CreateStaff html Fail !",logger.LogData{
			"err":err.Error(),
		})
		return errors.New(response.CommonErrorWhenHandler)
	}
}
*/

func (s Staff) ChangeStatus(ctx context.Context, staff model.Staff) (*responsemodel.ResponseUpdate, error) {
	var (
		d = dao.Staff{}
	)
	if err := d.UpdateByID(ctx, staff.ID, bson.M{
		"$set": bson.M{
			"active": !staff.Active,
		},
	}); err != nil {

		return nil, errors.New(response.CommonErrorWhenHandler)
	}
	return &responsemodel.ResponseUpdate{ID: staff.ID.Hex()}, nil
}

//ALL..
func (s Staff) All(ctx context.Context, q query.Query) (r responsemodel.ResponseList) {
	var (
		wg sync.WaitGroup
		d  = dao.Staff{}
	)
	r.List = make([]responsemodel.Staff, 0)

	cond := bson.M{}
	q.SetDefaultLimit()
	q.AssignKeyword(cond)
	q.AssignActive(cond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := d.FindByCondition(ctx, cond, q.GetFindOptionsUsingPage())
		r.List = s.getListResponse(ctx, docs)
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		r.Total = d.CountByCondition(ctx, cond)
	}()

	wg.Wait()

	r.Limit = q.Limit
	return r
}

func (s Staff) getListResponse(ctx context.Context, staffs []model.Staff) []responsemodel.Staff {
	var (
		wg sync.WaitGroup
	)
	result := make([]responsemodel.Staff, len(staffs))

	wg.Add(len(staffs))
	for i, value := range staffs {
		go func(index int, staff model.Staff) {
			defer wg.Done()
			result[index] = s.getResponse(ctx, staff)
		}(i, value)
	}
	wg.Wait()

	return result
}
func (s Staff) getResponse(ctx context.Context, staff model.Staff) responsemodel.Staff {
	res := responsemodel.Staff{
		ID:        staff.ID.Hex(),
		Name:      staff.Name,
		Phone:     staff.Phone,
		Active:    staff.Active,
		Avatar:    staff.Avatar,
		IsRoot:    staff.IsRoot,
		Email:     staff.Email,
		CreatedAt: ptime.TimeResponseInit(staff.CreatedAt),
		UpdatedAt: ptime.TimeResponseInit(staff.UpdatedAt),
	}
	if staff.Role != nil {
		role := dao.Role{}.FindByID(ctx, *staff.Role)
		res.Role = &responsemodel.RoleShort{
			ID:   role.ID.Hex(),
			Name: role.Name,
		}
	}
	return res
}

//Update...
func (s Staff) Update(ctx context.Context, staff model.Staff, body apimodel.StaffUpdate) (result *responsemodel.ResponseUpdate, err error) {
	var (
		d       = dao.Staff{}
		roleDAO = dao.Role{}
	)

	//func Role

	roleID, isValid := mongodb.NewIDFromString(body.Role)
	if !isValid {
		return nil, errors.New(response.CommonInvalidRole)
	}

	if role := roleDAO.FindByID(ctx, roleID); role.ID.IsZero() {
		return nil, errors.New(response.CommonNotFoundRole)
	}

	if staff.Email != body.Email {
		if s.isEmailExisted(ctx, body.Email) {
			return nil, errors.New(response.CommonExistedEmail)
		}
	}

	//update
	payload := bson.M{
		"$set": bson.M{
			"name":         body.Name,
			"searchString": mongodb.NonAccentVietnamese(body.Name),
			"phone":        body.Name,
			"email":        body.Email,
			"role":         roleID,
			"avatar":       body.Avatar,
		},
	}
	if err = d.UpdateByID(ctx, staff.ID, payload); err != nil {
		return nil, errors.New(response.CommonErrorWhenHandler)
	}
	return &responsemodel.ResponseUpdate{ID: staff.ID.Hex()}, nil
}

//GetMe...
func (s Staff) GetMe(ctx context.Context, staff model.Staff) (result responsemodel.Staff, err error) {
	return s.getResponse(ctx, staff), nil
}

func (s Staff) isPhoneNumberExisted(ctx context.Context, phone string) bool {
	var (
		d = dao.Staff{}
	)
	//Find
	cond := bson.M{
		"phone": phone,
	}
	total := d.CountByCondition(ctx, cond)
	return total != 0
}

func (s Staff) isEmailExisted(ctx context.Context, email string) bool {
	var (
		d = dao.Staff{}
	)
	//find
	cond := bson.M{
		"email": email,
	}
	total := d.CountByCondition(ctx, cond)

	return total != 0
}

//InitRootStaff...
func (s Staff) InitRootStaff() {
	var (
		ctx = context.Background()
		d   = dao.Staff{}
		now = ptime.Now()
	)

	staff := d.FindOne(ctx, bson.M{"isRoot": true})

	if !staff.ID.IsZero() {
		return
	}

	_ = d.InsertOne(ctx, model.Staff{
		ID:             mongodb.NewObjectID(),
		Name:           "root",
		Email:          "nhly123123456789@gmail.com",
		SearchString:   mongodb.NonAccentVietnamese("Root"),
		Active:         true,
		IsRoot:         true,
		HashedPassword: helper.HashPassword("12345678"),
		CreatedAt:      now,
		UpdatedAt:      now,
	})
}

func (s Staff) LoginWithEmail(ctx context.Context, payload apimodel.StaffLoginWithEmail) (result *responsemodel.StaffLogin, err error) {
	var (
		d = dao.Staff{}
	)

	// Find staff
	staff := d.FindOne(ctx, bson.M{
		"email": payload.Email,
	})
	if staff.ID.IsZero() {
		err = errors.New(response.CommonNotFoundStaff)
		return
	}

	// Check Password
	if !helper.CheckPasswordHash(payload.Password, staff.HashedPassword) {
		err = errors.New(response.CommonIncorrectPassword)
		return
	}

	if !staff.Active {
		return nil, errors.New(response.CommonBannedUser)
	}

	// Generate auth token
	authToken := GenerateAuthToken(staff)

	return &responsemodel.StaffLogin{ID: staff.ID.Hex(), Token: authToken}, nil
}

// GenerateAuthToken ...
func GenerateAuthToken(staff model.Staff) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": staff.ID,
		"exp": time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte("AUTH_SECRET")) //config.GetENV().AuthSecret
	return tokenString
}
