package data

import (
	"cardbinance/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID               uint64    `gorm:"primarykey;type:int"`
	Address          string    `gorm:"type:varchar(100);default:'no'"`
	Card             string    `gorm:"type:varchar(100);not null;default:'no'"`
	CardOrderId      string    `gorm:"type:varchar(100);not null;default:'no'"`
	CardNumber       string    `gorm:"type:varchar(100);not null;default:'no'"`
	CardTwoNumber    string    `gorm:"type:varchar(100);not null;default:'no'"`
	CardAmount       float64   `gorm:"type:decimal(65,20);not null"`
	Amount           float64   `gorm:"type:decimal(65,20)"`
	IsDelete         uint64    `gorm:"type:int"`
	Vip              uint64    `gorm:"type:int"`
	MyTotalAmount    uint64    `gorm:"type:bigint"`
	AmountTwo        uint64    `gorm:"type:bigint"`
	CardUserId       string    `gorm:"type:varchar(45);not null;default:'0'"`
	FirstName        string    `gorm:"type:varchar(45);not null;default:'no'"`
	LastName         string    `gorm:"type:varchar(45);not null;default:'no'"`
	BirthDate        string    `gorm:"type:varchar(45);not null;default:'no'"`
	Email            string    `gorm:"type:varchar(100);not null;default:'no'"`
	CountryCode      string    `gorm:"type:varchar(45);not null;default:'no'"`
	Phone            string    `gorm:"type:varchar(45);not null;default:'no'"`
	City             string    `gorm:"type:varchar(100);not null;default:'no'"`
	Country          string    `gorm:"type:varchar(100);not null;default:'no'"`
	Street           string    `gorm:"type:varchar(100);not null;default:'no'"`
	PostalCode       string    `gorm:"type:varchar(45);not null;default:'no'"`
	MaxCardQuota     uint64    `gorm:"type:bigint"`
	ProductId        string    `gorm:"type:varchar(45);not null;default:'0'"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
	VipTwo           uint64    `gorm:"type:int"`
	VipThree         uint64    `gorm:"type:int"`
	CardTwo          uint64    `gorm:"type:int"`
	CanVip           uint64    `gorm:"type:int"`
	UserCount        uint64    `gorm:"type:int"`
	LockCard         uint64    `gorm:"type:int"`
	LockCardTwo      uint64    `gorm:"type:int"`
	ChangeCard       uint64    `gorm:"type:int"`
	ChangeCardTwo    uint64    `gorm:"type:int"`
	Pic              string    `gorm:"type:varchar(45);not null;default:'no'"`
	PicTwo           string    `gorm:"type:varchar(45);not null;default:'no'"`
	CardNumberRel    string    `gorm:"type:varchar(100);not null;default:'no'"`
	CardNumberRelTwo string    `gorm:"type:varchar(100);not null;default:'no'"`
}

type CardTwo struct {
	ID               uint64    `gorm:"primarykey;type:int"`
	UserId           uint64    `gorm:"type:int;not null"`
	FirstName        string    `gorm:"type:varchar(45);not null;default:'no'"`
	LastName         string    `gorm:"type:varchar(45);not null;default:'no'"`
	Email            string    `gorm:"type:varchar(100);not null;default:'no'"`
	CountryCode      string    `gorm:"type:varchar(45);not null;default:'no'"`
	Phone            string    `gorm:"type:varchar(45);not null;default:'no'"`
	City             string    `gorm:"type:varchar(100);not null;default:'no'"`
	Country          string    `gorm:"type:varchar(100);not null;default:'no'"`
	Street           string    `gorm:"type:varchar(100);not null;default:'no'"`
	PostalCode       string    `gorm:"type:varchar(45);not null;default:'no'"`
	BirthDate        string    `gorm:"type:varchar(45);not null;default:'no'"`
	PhoneCountryCode string    `gorm:"type:varchar(45);not null;default:'no'"`
	State            string    `gorm:"type:varchar(45);not null;default:'no'"`
	Status           uint64    `gorm:"type:int"`
	CardId           string    `gorm:"type:varchar(100);not null;default:'no'"`
	CardAmount       float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
	IdCard           string    `gorm:"type:varchar(45);not null;default:'no'"`
	Gender           string    `gorm:"type:varchar(45);not null;default:'no'"`
}

type Admin struct {
	ID       int64  `gorm:"primarykey;type:int"`
	Account  string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Type     string `gorm:"type:varchar(40);not null"`
}

type UserRecommend struct {
	ID            uint64    `gorm:"primarykey;type:int"`
	UserId        uint64    `gorm:"type:int;not null"`
	RecommendCode string    `gorm:"type:varchar(10000);not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type Config struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(45);not null"`
	KeyName   string    `gorm:"type:varchar(45);not null"`
	Value     string    `gorm:"type:varchar(1000);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Reward struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	Reason    uint64    `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
	Address   string    `gorm:"type:varchar(100);not null"`
	One       uint64    `gorm:"type:int;not null"`
}

type CardRecord struct {
	ID         uint64    `gorm:"primarykey;type:int"`
	UserId     uint64    `gorm:"type:int;not null"`
	RecordType uint64    `gorm:"type:int;not null"`
	Remark     string    `gorm:"type:varchar(500);not null"`
	Code       string    `gorm:"type:varchar(100);not null"`
	Opt        string    `gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

type Withdraw struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	RelAmount float64   `gorm:"type:decimal(65,20);not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	Address   string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type EthUserRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Hash      string    `gorm:"type:varchar(100);not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Amount    string    `gorm:"type:varchar(45);not null"`
	AmountTwo uint64    `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
	Last      int64     `gorm:"type:int;not null"`
}

type CardOrder struct {
	ID        uint64     `gorm:"primarykey;type:int"`
	Last      uint64     `gorm:"type:int;not null"`                       // createTime(ms)
	Code      string     `gorm:"type:varchar(100);not null;default:'no'"` // referenceId
	Card      string     `gorm:"type:varchar(100);not null;default:'no'"` // referenceId
	Time      *time.Time `gorm:"type:datetime;not null"`
	CreatedAt time.Time  `gorm:"type:datetime;not null"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null"`
}

type Card struct {
	ID uint64 `gorm:"primarykey;type:int"`

	CardID       string `gorm:"type:varchar(100);not null"`             // interlace 的 id
	AccountID    string `gorm:"type:varchar(100);not null"`             // accountId
	CardholderID string `gorm:"type:varchar(100);not null;default:'0'"` // cardholderId
	BalanceID    string `gorm:"type:varchar(100);not null;default:'0'"` // balanceId
	BudgetID     string `gorm:"type:varchar(100);not null;default:'0'"` // budgetId
	ReferenceID  string `gorm:"type:varchar(100);not null;default:'0'"` // referenceId

	UserName string `gorm:"type:varchar(100);not null;default:'no'"` // userName
	Currency string `gorm:"type:varchar(20);not null;default:'no'"`  // currency
	Bin      string `gorm:"type:varchar(20);not null;default:'no'"`  // bin

	Status   string `gorm:"type:varchar(45);not null;default:'PENDING'"`      // 状态
	CardMode string `gorm:"type:varchar(45);not null;default:'VIRTUAL_CARD'"` // 模式
	Label    string `gorm:"type:varchar(100);not null;default:'no'"`          // 标签

	CardLastFour string `gorm:"type:varchar(10);not null;default:'no'"` // 后四位

	// Interlace 的创建时间，用毫秒时间戳存
	InterlaceCreateTime int64 `gorm:"type:bigint;not null"` // createTime(ms)
	UserId              int64 `gorm:"type:int;not null"`    // createTime(ms)

	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
	IsDelete  uint32    `gorm:"type:int;not null;default:0"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *UserRepo) SetNonceByAddress(ctx context.Context, wallet string) (int64, error) {
	key := "wallet:" + wallet

	val, err := u.data.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// 设置键值，60 秒后自动过期
		timestamp := time.Now().Unix()
		return timestamp, u.data.rdb.Set(ctx, key, timestamp, 60*time.Second).Err()

	} else if err != nil {
		return -1, err
	}

	// 转换为 int64 时间戳
	t, errThree := strconv.ParseInt(val, 10, 64)
	if errThree != nil {
		return 0, errThree
	}

	return t, nil
}

// GetAndDeleteWalletTimestamp 获取并删除，确保只用一次（无并发可用）
func (u *UserRepo) GetAndDeleteWalletTimestamp(ctx context.Context, wallet string) (string, error) {
	key := "wallet:" + wallet

	val, err := u.data.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	// 删除
	if errTwo := u.data.rdb.Del(ctx, key).Err(); errTwo != nil {
		return "", errTwo
	}

	return val, nil
}

func (u *UserRepo) GetUserByAddress(address string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where("address=?", address).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		CardAmount:    user.CardAmount,
		MyTotalAmount: user.MyTotalAmount,
		AmountTwo:     user.AmountTwo,
		IsDelete:      user.IsDelete,
		Vip:           user.Vip,
		ID:            user.ID,
		Address:       user.Address,
		Card:          user.Card,
		Amount:        user.Amount,
		CardNumber:    user.CardNumber,
		CardOrderId:   user.CardOrderId,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetUserByCardUserId(cardUserId string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where("card_user_id=?", cardUserId).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		CardAmount:    user.CardAmount,
		MyTotalAmount: user.MyTotalAmount,
		AmountTwo:     user.AmountTwo,
		IsDelete:      user.IsDelete,
		Vip:           user.Vip,
		ID:            user.ID,
		Address:       user.Address,
		Card:          user.Card,
		Amount:        user.Amount,
		CardNumber:    user.CardNumber,
		CardOrderId:   user.CardOrderId,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetUserByCard(card string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where("card=?", card).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		CardAmount:    user.CardAmount,
		MyTotalAmount: user.MyTotalAmount,
		AmountTwo:     user.AmountTwo,
		IsDelete:      user.IsDelete,
		Vip:           user.Vip,
		ID:            user.ID,
		Address:       user.Address,
		Card:          user.Card,
		Amount:        user.Amount,
		CardNumber:    user.CardNumber,
		CardOrderId:   user.CardOrderId,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetUserById(userId uint64) (*biz.User, error) {
	var user User
	if err := u.data.db.Where("id=?", userId).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		CardAmount:       user.CardAmount,
		MyTotalAmount:    user.MyTotalAmount,
		AmountTwo:        user.AmountTwo,
		IsDelete:         user.IsDelete,
		Vip:              user.Vip,
		ID:               user.ID,
		Address:          user.Address,
		Card:             user.Card,
		Amount:           user.Amount,
		CardNumber:       user.CardNumber,
		CardOrderId:      user.CardOrderId,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
		CardNumberRelTwo: user.CardNumberRelTwo,
	}, nil
}

// GetUserRecommendByUserId .
func (u *UserRepo) GetUserRecommendByUserId(userId uint64) (*biz.UserRecommend, error) {
	var userRecommend UserRecommend
	if err := u.data.db.Where("user_id=?", userId).Table("user_recommend").First(&userRecommend).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return &biz.UserRecommend{
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
	}, nil
}

// CreateUser .
func (u *UserRepo) CreateUser(ctx context.Context, uc *biz.User) (*biz.User, error) {
	var user User
	user.Address = uc.Address
	user.Card = "no"
	user.CardNumber = "no"
	user.CardOrderId = "no"
	if 0 < uc.Vip {
		user.Vip = uc.Vip
	}

	res := u.data.DB(ctx).Table("user").Create(&user)
	if res.Error != nil || 0 >= res.RowsAffected {
		return nil, errors.New(500, "CREATE_USER_ERROR", "用户创建失败")
	}

	return &biz.User{
		CardAmount:    user.CardAmount,
		MyTotalAmount: user.MyTotalAmount,
		AmountTwo:     user.AmountTwo,
		IsDelete:      user.IsDelete,
		Vip:           user.Vip,
		ID:            user.ID,
		Address:       user.Address,
		Card:          user.Card,
		Amount:        user.Amount,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		CardNumber:    user.CardNumber,
		CardOrderId:   user.CardOrderId,
	}, nil
}

// CreateUserRecommend .
func (u *UserRepo) CreateUserRecommend(ctx context.Context, userId uint64, recommendUser *biz.UserRecommend) (*biz.UserRecommend, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatUint(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var userRecommend UserRecommend
	userRecommend.UserId = userId
	userRecommend.RecommendCode = tmpRecommendCode

	res := u.data.DB(ctx).Table("user_recommend").Create(&userRecommend)
	if res.Error != nil || 0 >= res.RowsAffected {
		return nil, errors.New(500, "CREATE_USER_RECOMMEND_ERROR", "用户推荐关系创建失败")
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
	}, nil
}

// GetUserRecommendByCode .
func (u *UserRepo) GetUserRecommendByCode(code string) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)
	res := make([]*biz.UserRecommend, 0)

	instance := u.data.db.Table("user_recommend").Where("recommend_code=?", code)
	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
		})
	}

	return res, nil
}

// GetUserRecommendLikeCode .
func (u *UserRepo) GetUserRecommendLikeCode(code string) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)
	res := make([]*biz.UserRecommend, 0)

	instance := u.data.db.Table("user_recommend").Where("recommend_code Like ?", code+"%")
	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
		})
	}

	return res, nil
}

// UpdateWithdraw .
func (u *UserRepo) UpdateWithdraw(ctx context.Context, id uint64, status string) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.Status = status
	res := u.data.DB(ctx).Table("withdraw").Where("id=?", id).Updates(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Withdraw{
		ID:        withdraw.ID,
		UserId:    withdraw.UserId,
		Amount:    withdraw.Amount,
		RelAmount: withdraw.RelAmount,
		Status:    withdraw.Status,
		Address:   withdraw.Address,
		CreatedAt: withdraw.CreatedAt,
		UpdatedAt: withdraw.UpdatedAt,
	}, nil
}

// GetUserByUserIds .
func (u *UserRepo) GetUserByUserIds(userIds ...uint64) (map[uint64]*biz.User, error) {
	var users []*User

	res := make(map[uint64]*biz.User, 0)
	if err := u.data.db.Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res[user.ID] = &biz.User{
			CardAmount:    user.CardAmount,
			MyTotalAmount: user.MyTotalAmount,
			AmountTwo:     user.AmountTwo,
			IsDelete:      user.IsDelete,
			Vip:           user.Vip,
			ID:            user.ID,
			Address:       user.Address,
			Card:          user.Card,
			Amount:        user.Amount,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			CardNumber:    user.CardNumber,
			CardOrderId:   user.CardOrderId,
		}
	}

	return res, nil
}

// GetUserByUserIdsTwo .
func (u *UserRepo) GetUserByUserIdsTwo(userIds []uint64) (map[uint64]*biz.User, error) {
	var users []*User

	res := make(map[uint64]*biz.User, 0)
	if err := u.data.db.Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res[user.ID] = &biz.User{
			CardAmount:    user.CardAmount,
			MyTotalAmount: user.MyTotalAmount,
			AmountTwo:     user.AmountTwo,
			IsDelete:      user.IsDelete,
			Vip:           user.Vip,
			ID:            user.ID,
			Address:       user.Address,
			Card:          user.Card,
			Amount:        user.Amount,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			CardNumber:    user.CardNumber,
			CardOrderId:   user.CardOrderId,
		}
	}

	return res, nil
}

// GetConfigByKeys .
func (u *UserRepo) GetConfigByKeys(keys ...string) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := u.data.db.Where("key_name IN (?)", keys).Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// HasCardByCardID 判断数据库中是否已经存在该 card_id
func (u *UserRepo) HasCardByCardID(ctx context.Context, cardID string) (bool, error) {
	if cardID == "" {
		return false, nil
	}

	var c Card
	err := u.data.DB(ctx).
		Table("card").
		Select("id").
		Where("card_id = ?", cardID).
		Take(&c).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.New(500, "CARD_ERROR", err.Error())
	}

	return true, nil
}

// CreateCard .
func (u *UserRepo) CreateCard(ctx context.Context, userId uint64, user *biz.User) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("amount>=?", user.Amount).Where("card_order_id=?", "no").
		Updates(map[string]interface{}{
			"amount":        gorm.Expr("amount - ?", user.Amount),
			"card_order_id": "do",
			"first_name":    user.FirstName,
			"last_name":     user.LastName,
			"birth_date":    user.BirthDate,
			"email":         user.Email,
			"phone":         user.Phone,
			"country_code":  user.CountryCode,
			"country":       user.Country,
			"city":          user.City,
			"street":        user.Street,
			"postal_code":   user.PostalCode,
			"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = user.Amount
	reward.Reason = 3 // 给我分红的理由
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	return nil
}

// SetVip .
func (u *UserRepo) SetVip(ctx context.Context, userId uint64, vip uint64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"vip":        vip,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateCard .
func (u *UserRepo) UpdateCard(ctx context.Context, userId uint64, cardOrderId, card string) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("card_order_id=?", "do").
		Updates(map[string]interface{}{
			"card_order_id": cardOrderId,
			"card":          card,
			"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateCardNo .
func (u *UserRepo) UpdateCardNo(ctx context.Context, userId uint64, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_order_id": "no",
			"card":          "no",
			"amount":        gorm.Expr("amount + ?", amount),
			"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = 10
	reward.Reason = 7 // 给我分红的理由
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	return nil
}

// UpdateCardSuccess .
func (u *UserRepo) UpdateCardSucces(ctx context.Context, userId uint64, cardNum string) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_number": cardNum,
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// GetAllUsers .
func (u *UserRepo) GetAllUsers() ([]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, user := range users {
		res = append(res, &biz.User{
			CardAmount:    user.CardAmount,
			MyTotalAmount: user.MyTotalAmount,
			AmountTwo:     user.AmountTwo,
			IsDelete:      user.IsDelete,
			Vip:           user.Vip,
			ID:            user.ID,
			Address:       user.Address,
			Card:          user.Card,
			Amount:        user.Amount,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			CardNumber:    user.CardNumber,
			CardOrderId:   user.CardOrderId,
			VipTwo:        user.VipTwo,
			VipThree:      user.VipThree,
		})
	}
	return res, nil
}

// GetUserRecommends .
func (u *UserRepo) GetUserRecommends() ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := u.data.db.Table("user_recommend").Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
		})
	}

	return res, nil
}

// GetUsersOpenCard .
func (u *UserRepo) GetUsersOpenCard() ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	if err := u.data.db.Table("user").Where("card_order_id=?", "do").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			CardAmount:    user.CardAmount,
			MyTotalAmount: user.MyTotalAmount,
			AmountTwo:     user.AmountTwo,
			IsDelete:      user.IsDelete,
			Vip:           user.Vip,
			ID:            user.ID,
			Address:       user.Address,
			Card:          user.Card,
			Amount:        user.Amount,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			CardNumber:    user.CardNumber,
			CardOrderId:   user.CardOrderId,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			BirthDate:     user.BirthDate,
			Email:         user.Email,
			CountryCode:   user.CountryCode,
			Phone:         user.Phone,
			City:          user.City,
			Country:       user.Country,
			Street:        user.Street,
			PostalCode:    user.PostalCode,
			CardUserId:    user.CardUserId,
			MaxCardQuota:  user.MaxCardQuota,
			ProductId:     user.ProductId,
			VipTwo:        user.VipTwo,
		})
	}
	return res, nil
}

// GetUsersOpenCardStatusDoing .
func (u *UserRepo) GetUsersOpenCardStatusDoing() ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	if err := u.data.db.Table("user").Where("card!=?", "no").Where("card_number=?", "no").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			CardAmount:    user.CardAmount,
			MyTotalAmount: user.MyTotalAmount,
			AmountTwo:     user.AmountTwo,
			IsDelete:      user.IsDelete,
			Vip:           user.Vip,
			ID:            user.ID,
			Address:       user.Address,
			Card:          user.Card,
			Amount:        user.Amount,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			CardNumber:    user.CardNumber,
			CardOrderId:   user.CardOrderId,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			BirthDate:     user.BirthDate,
			Email:         user.Email,
			CountryCode:   user.CountryCode,
			Phone:         user.Phone,
			City:          user.City,
			Country:       user.Country,
			Street:        user.Street,
			PostalCode:    user.PostalCode,
			CardUserId:    user.CardUserId,
			MaxCardQuota:  user.MaxCardQuota,
			ProductId:     user.ProductId,
			VipTwo:        user.VipTwo,
		})
	}
	return res, nil
}

// GetWithdrawPassOrRewardedFirst .
func (u *UserRepo) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*biz.Withdraw, error) {
	var withdraw *Withdraw
	if err := u.data.db.Table("withdraw").Where("status=?", "rewarded").First(&withdraw).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	return &biz.Withdraw{
		ID:        withdraw.ID,
		UserId:    withdraw.UserId,
		Amount:    withdraw.Amount,
		RelAmount: withdraw.RelAmount,
		Status:    withdraw.Status,
		Address:   withdraw.Address,
		CreatedAt: withdraw.CreatedAt,
		UpdatedAt: withdraw.UpdatedAt,
	}, nil
}

// CreateCardRecommend .
func (u *UserRepo) CreateCardRecommend(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("vip=?", vip).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount + ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}
	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = amount
	reward.One = vip
	reward.Reason = 6 // 给我分红的理由
	reward.Address = address
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	return nil
}

// CreateCardRecommendTwo .
func (u *UserRepo) CreateCardRecommendTwo(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("vip=?", vip).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount + ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}
	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = amount
	reward.One = vip
	reward.Reason = 11 // 给我分红的理由
	reward.Address = address
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	return nil
}

// UpdateCardTwo .
func (u *UserRepo) UpdateCardTwo(ctx context.Context, id uint64) error {
	res := u.data.DB(ctx).Table("reward").Where("id=?", id).
		Updates(map[string]interface{}{
			"one":        1,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_REWARD_CARD_ERROR", "信息修改失败")
	}

	return nil
}

// AmountTo .
func (u *UserRepo) AmountTo(ctx context.Context, userId, toUserId uint64, toAddress string, amount float64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("amount>=?", amount).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount - ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	resTwo := u.data.DB(ctx).Table("user").Where("id=?", toUserId).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount + ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = amount
	reward.Reason = 5 // 给我分红的理由
	reward.Address = toAddress
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	return nil
}

// Withdraw .
func (u *UserRepo) Withdraw(ctx context.Context, userId uint64, amount, amountRel float64, address string) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("amount>=?", amount).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount - ?", amount),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	var withdraw Withdraw
	withdraw.UserId = userId
	withdraw.Amount = amount
	withdraw.RelAmount = amountRel
	withdraw.Status = "rewarded"
	withdraw.Address = address
	resTwo := u.data.DB(ctx).Table("withdraw").Create(&withdraw)
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录创建失败")
	}

	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = amount
	reward.Reason = 2 // 给我分红的理由
	reward.Address = address
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	return nil
}

// GetUserRewardByUserIdPage .
func (u *UserRepo) GetUserRewardByUserIdPage(ctx context.Context, b *biz.Pagination, userId uint64, reason uint64) ([]*biz.Reward, error, int64) {
	var (
		count   int64
		rewards []*Reward
	)

	res := make([]*biz.Reward, 0)

	instance := u.data.db.Table("reward").Order("id desc")
	if 0 < userId {
		instance = instance.Where("user_id", userId)
	}

	if 0 < reason {
		instance = instance.Where("reason=?", reason)
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:        reward.ID,
			UserId:    reward.UserId,
			Amount:    reward.Amount,
			Reason:    reward.Reason,
			CreatedAt: reward.CreatedAt,
			Address:   reward.Address,
			One:       reward.One,
			UpdatedAt: reward.UpdatedAt,
		})
	}

	return res, nil, count
}

func (u *UserRepo) GetEthUserRecordLast() (int64, error) {
	var ethUserRecord *EthUserRecord
	if err := u.data.db.Table("eth_user_record").Order("last desc").First(&ethUserRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}

		return -1, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return ethUserRecord.Last, nil
}

// GetUserByAddresses .
func (u *UserRepo) GetUserByAddresses(addresses ...string) (map[string]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Where("address IN (?)", addresses).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make(map[string]*biz.User, 0)
	for _, item := range users {
		res[item.Address] = &biz.User{
			ID:      item.ID,
			Address: item.Address,
		}
	}
	return res, nil
}

func (u *UserRepo) CreateEthUserRecordListByHash(ctx context.Context, r *biz.EthUserRecord) (*biz.EthUserRecord, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", r.UserId).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount + ?", float64(r.AmountTwo)),
			"amount_two": gorm.Expr("amount_two + ?", r.AmountTwo),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return nil, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}
	var (
		reward Reward
	)
	reward.UserId = uint64(r.UserId)
	reward.Amount = float64(r.AmountTwo)
	reward.Reason = 1 // 给我分红的理由
	resInsert := u.data.DB(ctx).Table("reward").Create(&reward)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return nil, errors.New(500, "CREATE_LOCATION_ERROR", "信息创建失败")
	}

	var ethUserRecord EthUserRecord
	ethUserRecord.UserId = r.UserId
	ethUserRecord.Hash = r.Hash
	ethUserRecord.Amount = r.Amount
	ethUserRecord.AmountTwo = r.AmountTwo
	ethUserRecord.Last = r.Last

	resTwo := u.data.DB(ctx).Table("eth_user_record").Create(&ethUserRecord)
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return nil, errors.New(500, "CREATE_ETH_USER_RECORD_ERROR", "以太坊交易信息创建失败")
	}

	return &biz.EthUserRecord{
		ID:        ethUserRecord.ID,
		UserId:    ethUserRecord.UserId,
		Hash:      ethUserRecord.Hash,
		Amount:    ethUserRecord.Amount,
		AmountTwo: ethUserRecord.AmountTwo,
		Last:      ethUserRecord.Last,
	}, nil
}

// UpdateUserMyTotalAmountAdd .
func (u *UserRepo) UpdateUserMyTotalAmountAdd(ctx context.Context, userId uint64, amount uint64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"my_total_amount": gorm.Expr("my_total_amount + ?", amount),
			"updated_at":      time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// InsertCardRecord .
func (u *UserRepo) InsertCardRecord(ctx context.Context, userId, recordType uint64, remark string, code string, opt string) error {
	var (
		record CardRecord
	)

	record.UserId = userId
	record.Remark = remark
	record.RecordType = recordType
	record.Code = code
	record.Opt = opt
	resInsert := u.data.DB(ctx).Table("card_record").Create(&record)
	if resInsert.Error != nil || 0 >= resInsert.RowsAffected {
		return errors.New(500, "CREATE_CARD_RECORD_ERROR", "信息创建失败")
	}

	return nil
}

// GetUserCardTwo .
func (u *UserRepo) GetUserCardTwo() ([]*biz.Reward, error) {
	var (
		rewards []*Reward
	)

	res := make([]*biz.Reward, 0)

	instance := u.data.db.Table("reward").Order("id asc")
	instance = instance.Where("reason=?", 9).Where("one=?", 0)
	if err := instance.Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:        reward.ID,
			UserId:    reward.UserId,
			Amount:    reward.Amount,
			Reason:    reward.Reason,
			CreatedAt: reward.CreatedAt,
			Address:   reward.Address,
			One:       reward.One,
			UpdatedAt: reward.UpdatedAt,
		})
	}

	return res, nil
}

// GetUsers .
func (u *UserRepo) GetUsers(b *biz.Pagination, address string, cardTwo uint64, cardOrderId string, lockCard uint64, changeCard uint64) ([]*biz.User, error, int64) {
	var (
		users []*User
		count int64
	)
	instance := u.data.db.Table("user")
	if "" != address {
		instance = instance.Where("address=?", address)
	}

	if "all" != cardOrderId {
		instance = instance.Where("card_order_id=?", cardOrderId)
	}

	if 3 != cardTwo {
		instance = instance.Where("card_two=?", cardTwo)
	}

	if 1 == lockCard {
		instance = instance.Where("lock_card=?", 1)
	} else if 2 == lockCard {
		instance = instance.Where("lock_card_two=?", 1)
	}

	if 1 == changeCard {
		instance = instance.Where("change_card=?", 1)
	} else if 2 == changeCard {
		instance = instance.Where("change_card_two=?", 1)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found"), 0
		}

		return nil, errors.New(500, "USER ERROR", err.Error()), 0
	}

	res := make([]*biz.User, 0)
	for _, user := range users {
		res = append(res, &biz.User{
			ID:               user.ID,
			Address:          user.Address,
			Card:             user.Card,
			CardNumber:       user.CardNumber,
			CardOrderId:      user.CardOrderId,
			CardAmount:       user.CardAmount,
			Amount:           user.Amount,
			AmountTwo:        user.AmountTwo,
			MyTotalAmount:    user.MyTotalAmount,
			IsDelete:         user.IsDelete,
			Vip:              user.Vip,
			FirstName:        user.FirstName,
			LastName:         user.LastName,
			BirthDate:        user.BirthDate,
			Email:            user.Email,
			CountryCode:      user.CountryCode,
			Phone:            user.Phone,
			City:             user.City,
			Country:          user.Country,
			Street:           user.Street,
			PostalCode:       user.PostalCode,
			CardUserId:       user.CardUserId,
			ProductId:        user.ProductId,
			MaxCardQuota:     user.MaxCardQuota,
			CreatedAt:        user.CreatedAt,
			UpdatedAt:        user.UpdatedAt,
			VipTwo:           user.VipTwo,
			VipThree:         user.VipThree,
			CanVip:           user.CanVip,
			CardTwo:          user.CardTwo,
			UserCount:        user.UserCount,
			CardTwoNumber:    user.CardTwoNumber,
			Pic:              user.Pic,
			PicTwo:           user.PicTwo,
			LockCard:         user.LockCard,
			LockCardTwo:      user.LockCardTwo,
			ChangeCardTwo:    user.ChangeCardTwo,
			ChangeCard:       user.ChangeCard,
			CardNumberRel:    user.CardNumberRel,
			CardNumberRelTwo: user.CardNumberRelTwo,
		})
	}
	return res, nil, count
}

// GetAdminByAccount .
func (u *UserRepo) GetAdminByAccount(ctx context.Context, account string, password string) (*biz.Admin, error) {
	var admin Admin
	if err := u.data.db.Where("account=? and password=?", account, password).Table("admin").First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ADMIN_NOT_FOUND", "admin not found")
		}

		return nil, errors.New(500, "ADMIN ERROR", err.Error())
	}

	return &biz.Admin{
		ID:       admin.ID,
		Password: admin.Password,
		Account:  admin.Account,
		Type:     admin.Type,
	}, nil
}

func (u *UserRepo) SetCanVip(ctx context.Context, userId uint64, lock uint64) (bool, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Updates(map[string]interface{}{"can_vip": lock})
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_ERROR", "用户修改失败")
	}

	return true, nil
}

func (u *UserRepo) SetVipThree(ctx context.Context, userId uint64, vipThree uint64) (bool, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Updates(map[string]interface{}{"vip_three": vipThree})
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_ERROR", "用户修改失败")
	}

	return true, nil
}

func (u *UserRepo) SetUserCount(ctx context.Context, userId uint64) (bool, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Updates(map[string]interface{}{"user_count": 0})
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_ERROR", "用户修改失败")
	}

	return true, nil
}

// GetConfigs .
func (u *UserRepo) GetConfigs() ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := u.data.db.Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("CONFIG_NOT_FOUND", "config not found")
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// UpdateConfig .
func (u *UserRepo) UpdateConfig(ctx context.Context, id int64, value string) (bool, error) {
	var config Config
	config.Value = value

	res := u.data.DB(ctx).Table("config").Where("id=?", id).Updates(&config)
	if res.Error != nil {
		return false, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return true, nil
}

// UpdateUserInfo .
func (u *UserRepo) UpdateUserInfo(ctx context.Context, userId uint64, user *biz.User) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Where("card_order_id=?", "do").
		Updates(map[string]interface{}{
			"first_name":         user.FirstName,
			"last_name":          user.LastName,
			"phone":              user.Phone,
			"country_code":       user.CountryCode,
			"birth_date":         user.BirthDate,
			"country":            user.Country,
			"street":             user.Street,
			"postal_code":        user.PostalCode,
			"gender":             user.Gender,
			"id_card":            user.IdCard,
			"id_type":            user.IdType,
			"state":              user.State,
			"city":               user.City,
			"phone_country_code": user.PhoneCountryCode,
			"card_order_id":      "upload",
			"updated_at":         time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// GetCardByCardId z
func (u *UserRepo) GetCardByCardId(ctx context.Context, cardId string) (*biz.Card, error) {
	var c Card

	instance := u.data.DB(ctx).Table("card").Where("card_id=?", cardId)

	if err := instance.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有就返回 nil, nil，或者你习惯的 NotFound 错误
			return nil, nil
		}
		return nil, errors.New(500, "CARD_ERROR", err.Error())
	}

	res := &biz.Card{
		ID:                  c.ID,
		CardID:              c.CardID,
		AccountID:           c.AccountID,
		CardholderID:        c.CardholderID,
		BalanceID:           c.BalanceID,
		BudgetID:            c.BudgetID,
		ReferenceID:         c.ReferenceID,
		UserName:            c.UserName,
		Currency:            c.Currency,
		Bin:                 c.Bin,
		Status:              c.Status,
		CardMode:            c.CardMode,
		Label:               c.Label,
		CardLastFour:        c.CardLastFour,
		InterlaceCreateTime: c.InterlaceCreateTime,
		CreatedAt:           c.CreatedAt,
		UpdatedAt:           c.UpdatedAt,
		UserId:              c.UserId,
	}

	return res, nil
}

// GetNoBindCardV 按 InterlaceCreateTime 倒序取最新一条
func (u *UserRepo) GetNoBindCardV(ctx context.Context) (*biz.Card, error) {
	var c Card

	instance := u.data.DB(ctx).Table("card").
		Where("user_id=?", 0).
		Where("card_mode=?", "VIRTUAL_CARD").
		Order("id asc")

	if err := instance.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有就返回 nil, nil，或者你习惯的 NotFound 错误
			return nil, nil
		}
		return nil, errors.New(500, "CARD_ERROR", err.Error())
	}

	res := &biz.Card{
		ID:                  c.ID,
		CardID:              c.CardID,
		AccountID:           c.AccountID,
		CardholderID:        c.CardholderID,
		BalanceID:           c.BalanceID,
		BudgetID:            c.BudgetID,
		ReferenceID:         c.ReferenceID,
		UserName:            c.UserName,
		Currency:            c.Currency,
		Bin:                 c.Bin,
		Status:              c.Status,
		CardMode:            c.CardMode,
		Label:               c.Label,
		CardLastFour:        c.CardLastFour,
		InterlaceCreateTime: c.InterlaceCreateTime,
		CreatedAt:           c.CreatedAt,
		UpdatedAt:           c.UpdatedAt,
		UserId:              c.UserId,
	}

	return res, nil
}

// GetLatestCard 按 InterlaceCreateTime 倒序取最新一条
func (u *UserRepo) GetLatestCard(ctx context.Context) (*biz.Card, error) {
	var c Card

	instance := u.data.DB(ctx).Table("card").
		Order("id DESC").
		Limit(1)

	if err := instance.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有就返回 nil, nil，或者你习惯的 NotFound 错误
			return nil, nil
		}
		return nil, errors.New(500, "CARD_ERROR", err.Error())
	}

	res := &biz.Card{
		ID:                  c.ID,
		CardID:              c.CardID,
		AccountID:           c.AccountID,
		CardholderID:        c.CardholderID,
		BalanceID:           c.BalanceID,
		BudgetID:            c.BudgetID,
		ReferenceID:         c.ReferenceID,
		UserName:            c.UserName,
		Currency:            c.Currency,
		Bin:                 c.Bin,
		Status:              c.Status,
		CardMode:            c.CardMode,
		Label:               c.Label,
		CardLastFour:        c.CardLastFour,
		InterlaceCreateTime: c.InterlaceCreateTime,
		CreatedAt:           c.CreatedAt,
		UpdatedAt:           c.UpdatedAt,
		UserId:              c.UserId,
	}

	return res, nil
}

// GetCardPage 分页查询卡片
func (u *UserRepo) GetCardPage(ctx context.Context, b *biz.Pagination, accountId, status string) ([]*biz.Card, error, int64) {
	var (
		count int64
		list  []*Card
	)

	res := make([]*biz.Card, 0)

	instance := u.data.db.Table("card").
		Order("id DESC")

	if accountId != "" {
		instance = instance.Where("account_id = ?", accountId)
	}
	if status != "" {
		instance = instance.Where("status = ?", status)
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&list).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("CARD_NOT_FOUND", "card not found"), 0
		}
		return nil, errors.New(500, "CARD_ERROR", err.Error()), 0
	}

	for _, c := range list {
		res = append(res, &biz.Card{
			ID:                  c.ID,
			CardID:              c.CardID,
			AccountID:           c.AccountID,
			CardholderID:        c.CardholderID,
			BalanceID:           c.BalanceID,
			BudgetID:            c.BudgetID,
			ReferenceID:         c.ReferenceID,
			UserName:            c.UserName,
			Currency:            c.Currency,
			Bin:                 c.Bin,
			Status:              c.Status,
			CardMode:            c.CardMode,
			Label:               c.Label,
			CardLastFour:        c.CardLastFour,
			InterlaceCreateTime: c.InterlaceCreateTime,
			UserId:              c.UserId,
			CreatedAt:           c.CreatedAt,
			UpdatedAt:           c.UpdatedAt,
		})
	}

	return res, nil, count
}

// CreateCardNew 创建一条卡片记录
func (u *UserRepo) CreateCardNew(ctx context.Context, userId, id uint64, in *biz.Card, isNew bool) error {
	res := u.data.DB(ctx).Table("card_two").Where("id=?", id).
		Updates(map[string]interface{}{
			"status":     2,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "CreateCardNew", "用户信息修改失败")
	}

	resTwo := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_two_number": in.CardID,
			"card_two":        2,
			"lock_card_two":   0,
			"change_card_two": 0,
			"updated_at":      time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "CreateCardNew", "用户信息修改失败")
	}

	if isNew {
		var c Card

		c.CardID = in.CardID
		c.AccountID = in.AccountID
		c.CardholderID = in.CardholderID
		c.BalanceID = in.BalanceID
		c.BudgetID = in.BudgetID
		c.ReferenceID = in.ReferenceID

		c.UserName = in.UserName
		c.Currency = in.Currency
		c.Bin = in.Bin
		c.Status = in.Status
		c.CardMode = in.CardMode
		c.Label = in.Label
		c.CardLastFour = in.CardLastFour

		c.InterlaceCreateTime = in.InterlaceCreateTime
		c.UserId = int64(userId)

		resInsert := u.data.DB(ctx).Table("card").Create(&c)
		if resInsert.Error != nil || resInsert.RowsAffected <= 0 {
			return errors.New(500, "CREATE_CARD_ERROR", "卡片信息创建失败")
		}
	} else {
		resThree := u.data.DB(ctx).Table("card").Where("card_id=?", in.CardID).
			Updates(map[string]interface{}{
				"user_id":    userId,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if resThree.Error != nil || 0 >= resThree.RowsAffected {
			return errors.New(500, "CreateCardOne", "用户信息修改失败")
		}
	}

	return nil
}

// UpdateUserDoing 创建一条卡片记录
func (u *UserRepo) UpdateUserDoing(ctx context.Context, userId uint64, cardNumber, cardNumberRel string, cardAmount float64) error {
	resTwo := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_order_id":   "doing",
			"card_number":     cardNumber,
			"card_amount":     cardAmount,
			"card_number_rel": cardNumberRel,
			"updated_at":      time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "CreateCardOne", "用户信息修改失败")
	}

	return nil
}

// UpdateUserDone 创建一条卡片记录
func (u *UserRepo) UpdateUserDone(ctx context.Context, userId uint64, cardId string, cardAmount float64) error {
	resTwo := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_order_id": "success",
			"lock_card":     0,
			"card_number":   cardId,
			"change_card":   0,
			"card_amount":   cardAmount,
			"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UpdateUserDone", "用户信息修改失败")
	}

	resThree := u.data.DB(ctx).Table("card").Where("card_id=?", cardId).
		Updates(map[string]interface{}{
			"user_id":    userId,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resThree.Error != nil || 0 >= resThree.RowsAffected {
		return errors.New(500, "UpdateUserDone", "用户信息修改失败")
	}

	return nil
}

// UpdateCardStatus 创建一条卡片记录
func (u *UserRepo) UpdateCardStatus(ctx context.Context, id, userId uint64, cardNumber, cardNumberRel string, cardAmount float64) error {
	res := u.data.DB(ctx).Table("card_two").Where("id=?", id).
		Updates(map[string]interface{}{
			"card_amount": cardAmount,
			"card_id":     cardNumber,
			"status":      1,
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UpdateCardStatus", "用户信息修改失败")
	}

	resTwo := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_number_rel_two": cardNumberRel,
			"updated_at":          time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "CreateCardOne", "用户信息修改失败")
	}

	return nil
}

// CreateCardOnly 创建一条卡片记录
func (u *UserRepo) CreateCardOnly(ctx context.Context, in *biz.Card) error {
	var c Card

	c.CardID = in.CardID
	c.AccountID = in.AccountID
	c.CardholderID = in.CardholderID
	c.BalanceID = in.BalanceID
	c.BudgetID = in.BudgetID
	c.ReferenceID = in.ReferenceID

	c.UserName = in.UserName
	c.Currency = in.Currency
	c.Bin = in.Bin
	c.Status = in.Status
	c.CardMode = in.CardMode
	c.Label = in.Label
	c.CardLastFour = in.CardLastFour

	c.InterlaceCreateTime = in.InterlaceCreateTime
	c.UserId = int64(0)

	resInsert := u.data.DB(ctx).Table("card").Create(&c)
	if resInsert.Error != nil || resInsert.RowsAffected <= 0 {
		return errors.New(500, "CREATE_CARD_ERROR", "卡片信息创建失败")
	}

	return nil
}

// CreateCardOne 创建一条卡片记录
func (u *UserRepo) CreateCardOne(ctx context.Context, userId uint64, in *biz.Card, isNew bool) error {
	resTwo := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"card_order_id": "success",
			"lock_card":     0,
			"change_card":   0,
			"updated_at":    time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "CreateCardOne", "用户信息修改失败")
	}

	if isNew {
		var c Card

		c.CardID = in.CardID
		c.AccountID = in.AccountID
		c.CardholderID = in.CardholderID
		c.BalanceID = in.BalanceID
		c.BudgetID = in.BudgetID
		c.ReferenceID = in.ReferenceID

		c.UserName = in.UserName
		c.Currency = in.Currency
		c.Bin = in.Bin
		c.Status = in.Status
		c.CardMode = in.CardMode
		c.Label = in.Label
		c.CardLastFour = in.CardLastFour

		c.InterlaceCreateTime = in.InterlaceCreateTime
		c.UserId = int64(userId)

		resInsert := u.data.DB(ctx).Table("card").Create(&c)
		if resInsert.Error != nil || resInsert.RowsAffected <= 0 {
			return errors.New(500, "CREATE_CARD_ERROR", "卡片信息创建失败")
		}
	} else {
		resThree := u.data.DB(ctx).Table("card").Where("card_id=?", in.CardID).
			Updates(map[string]interface{}{
				"user_id":    userId,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if resThree.Error != nil || 0 >= resThree.RowsAffected {
			return errors.New(500, "CreateCardOne", "用户信息修改失败")
		}
	}

	return nil
}

// GetCardTwoById .
func (u *UserRepo) GetCardTwoById(id uint64) (*biz.CardTwo, error) {
	var (
		c *CardTwo
	)

	// 按 id 升序，你可以按需要改成 desc
	instance := u.data.db.Table("card_two").Where("id = ?", id)

	if err := instance.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 跟 GetConfigs 一样风格，返回 NotFound 错误
			return nil, errors.NotFound("CARD_TWO_NOT_FOUND", "card_two not found")
		}

		return nil, errors.New(500, "CARD_TWO_ERROR", err.Error())
	}

	return &biz.CardTwo{
		ID:               c.ID,
		UserId:           c.UserId,
		FirstName:        c.FirstName,
		LastName:         c.LastName,
		Email:            c.Email,
		CountryCode:      c.CountryCode,
		Phone:            c.Phone,
		City:             c.City,
		Country:          c.Country,
		Street:           c.Street,
		PostalCode:       c.PostalCode,
		BirthDate:        c.BirthDate,
		PhoneCountryCode: c.PhoneCountryCode,
		State:            c.State,
		Status:           c.Status,
		CardId:           c.CardId,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
		CardAmount:       c.CardAmount,
	}, nil
}

// GetCardOrder .
func (u *UserRepo) GetCardOrder() (*biz.CardOrder, error) {
	var (
		c *CardOrder
	)

	// 按 id 升序，你可以按需要改成 desc
	instance := u.data.db.Table("card_code").Order("last desc")

	if err := instance.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 跟 GetConfigs 一样风格，返回 NotFound 错误
			return nil, nil
		}

		return nil, errors.New(500, "CARD_ORDER_ERROR", err.Error())
	}

	return &biz.CardOrder{
		ID:        c.ID,
		Last:      c.Last,
		Code:      c.Code,
		Card:      c.Card,
		Time:      c.Time,
		CreatedAt: c.CreatedAt,
	}, nil
}

// CreateCardOrder
func (u *UserRepo) CreateCardOrder(ctx context.Context, in *biz.CardOrder) error {
	var c CardOrder

	c.Last = in.Last
	c.Card = in.Card
	c.Code = in.Code
	c.Time = in.Time

	resInsert := u.data.DB(ctx).Table("card_code").Create(&c)
	if resInsert.Error != nil || resInsert.RowsAffected <= 0 {
		return errors.New(500, "CREATE_CARD_CODE_ERROR", "卡片信息创建失败")
	}

	return nil
}

// GetCardTwoStatusOne .
// 查询 status = 0 的 card_two 记录
func (u *UserRepo) GetCardTwoStatusOne() ([]*biz.CardTwo, error) {
	var (
		cardTwos []*CardTwo
	)

	res := make([]*biz.CardTwo, 0)

	// 按 id 升序，你可以按需要改成 desc
	instance := u.data.db.Table("card_two").Where("status = ?", 1).Order("id asc")

	if err := instance.Find(&cardTwos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 跟 GetConfigs 一样风格，返回 NotFound 错误
			return res, errors.NotFound("CARD_TWO_NOT_FOUND", "card_two not found")
		}

		return nil, errors.New(500, "CARD_TWO_ERROR", err.Error())
	}

	for _, c := range cardTwos {
		res = append(res, &biz.CardTwo{
			ID:               c.ID,
			UserId:           c.UserId,
			FirstName:        c.FirstName,
			LastName:         c.LastName,
			Email:            c.Email,
			CountryCode:      c.CountryCode,
			Phone:            c.Phone,
			City:             c.City,
			Country:          c.Country,
			Street:           c.Street,
			PostalCode:       c.PostalCode,
			BirthDate:        c.BirthDate,
			PhoneCountryCode: c.PhoneCountryCode,
			State:            c.State,
			Status:           c.Status,
			CardId:           c.CardId,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
			CardAmount:       c.CardAmount,
		})
	}

	return res, nil
}

// GetUsersStatusDoing .
func (u *UserRepo) GetUsersStatusDoing() ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	if err := u.data.db.Table("user").Where("card_order_id=?", "doing").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, user := range users {
		res = append(res, &biz.User{
			CardAmount:    user.CardAmount,
			MyTotalAmount: user.MyTotalAmount,
			AmountTwo:     user.AmountTwo,
			IsDelete:      user.IsDelete,
			Vip:           user.Vip,
			ID:            user.ID,
			Address:       user.Address,
			Card:          user.Card,
			Amount:        user.Amount,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
			CardNumber:    user.CardNumber,
			CardOrderId:   user.CardOrderId,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			BirthDate:     user.BirthDate,
			Email:         user.Email,
			CountryCode:   user.CountryCode,
			Phone:         user.Phone,
			City:          user.City,
			Country:       user.Country,
			Street:        user.Street,
			PostalCode:    user.PostalCode,
			CardUserId:    user.CardUserId,
			MaxCardQuota:  user.MaxCardQuota,
			ProductId:     user.ProductId,
			VipTwo:        user.VipTwo,
		})
	}
	return res, nil
}

// GetCardTwos .
// 分页查询实体卡用户信息（card_two 表）
func (u *UserRepo) GetCardTwos(b *biz.Pagination, userId uint64, status uint64, cardId string) ([]*biz.CardTwo, error, int64) {
	var (
		cardTwos []*CardTwo
		count    int64
	)

	// 基础查询
	instance := u.data.db.Table("card_two")

	if userId > 0 {
		instance = instance.Where("user_id = ?", userId)
	}

	instance = instance.Where("status = ?", status)

	if cardId != "" {
		instance = instance.Where("card_id = ?", cardId)
	}

	// 统计总数
	instance = instance.Count(&count)

	// 分页 + 排序
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).
		Order("id desc").
		Find(&cardTwos).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("CARD_TWO_NOT_FOUND", "card_two not found"), 0
		}

		return nil, errors.New(500, "CARD_TWO_ERROR", err.Error()), 0
	}

	// 转成 biz 层对象
	res := make([]*biz.CardTwo, 0, len(cardTwos))
	for _, c := range cardTwos {
		res = append(res, &biz.CardTwo{
			ID:               c.ID,
			UserId:           c.UserId,
			FirstName:        c.FirstName,
			LastName:         c.LastName,
			Email:            c.Email,
			CountryCode:      c.CountryCode,
			Phone:            c.Phone,
			City:             c.City,
			Country:          c.Country,
			Street:           c.Street,
			PostalCode:       c.PostalCode,
			BirthDate:        c.BirthDate,
			PhoneCountryCode: c.PhoneCountryCode,
			State:            c.State,
			Status:           c.Status,
			CardId:           c.CardId,
			CardAmount:       c.CardAmount,
			CreatedAt:        c.CreatedAt,
			UpdatedAt:        c.UpdatedAt,
			IdCard:           c.IdCard,
			Gender:           c.Gender,
		})
	}

	return res, nil, count
}
