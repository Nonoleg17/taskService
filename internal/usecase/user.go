package usecase

import (
	"context"
	"github.com/google/uuid"
	"strconv"
	"taskService/internal/entity"
	"taskService/pkg/util"
	"time"
)

const (
	secretKey = "secret"
)

type UserCase struct {
	userRepo    UserRepo
	sessionRepo SessionRepo
	timeout     time.Duration
}

func NewUserCase(ur UserRepo, sr SessionRepo) *UserCase {
	return &UserCase{
		userRepo:    ur,
		sessionRepo: sr,
		timeout:     time.Duration(2) * time.Second,
	}
}

func (uc *UserCase) CreateUser(c context.Context, req *entity.CreateUserReq) (*entity.CreateUserRes, error) {
	ctx, cansel := context.WithTimeout(c, uc.timeout)
	defer cansel()

	hashedPass, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPass,
	}
	res, err := uc.userRepo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	resp := &entity.CreateUserRes{
		ID:       strconv.Itoa(int(res.ID)),
		Username: res.Username,
		Email:    res.Email,
	}
	return resp, nil

}

func (uc *UserCase) Login(c context.Context, req *entity.LoginUserReq) (*entity.LoginUserRes, error) {
	ctx, cansel := context.WithTimeout(c, uc.timeout)
	defer cansel()
	user, err := uc.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &entity.LoginUserRes{}, err
	}
	token := uuid.NewString()
	session, err := uc.sessionRepo.Set(token, user.Username)
	if err != nil {
		return &entity.LoginUserRes{}, err
	}
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
	//	ID:       strconv.Itoa(int(user.ID)),
	//	Username: user.Username,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		Issuer:    strconv.Itoa(int(user.ID)),
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	//	},
	//})
	//signedStr, err := token.SignedString([]byte(secretKey))
	//if err != nil {
	//	return &entity.LoginUserRes{}, err
	//}
	//return &entity.LoginUserRes{
	//	AccessToken: signedStr,
	//	Username:    user.Username,
	//	ID:          strconv.Itoa(int(user.ID)),
	//}, nil
}
