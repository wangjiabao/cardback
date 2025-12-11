package biz

import (
	"bytes"
	pb "cardbinance/api/user/v1"
	"cardbinance/internal/pkg/middleware/auth"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Admin struct {
	ID       int64
	Password string
	Account  string
	Type     string
}

type CardTwo struct {
	ID               uint64
	UserId           uint64
	FirstName        string
	LastName         string
	Email            string
	CountryCode      string
	Phone            string
	City             string
	Country          string
	Street           string
	PostalCode       string
	BirthDate        string
	PhoneCountryCode string
	State            string
	Status           uint64
	CardId           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type User struct {
	ID               uint64
	Address          string
	Card             string
	CardNumber       string
	CardOrderId      string
	CardAmount       float64
	Amount           float64
	AmountTwo        uint64
	MyTotalAmount    uint64
	IsDelete         uint64
	Vip              uint64
	FirstName        string
	LastName         string
	BirthDate        string
	Email            string
	CountryCode      string
	PhoneCountryCode string
	Phone            string
	City             string
	Country          string
	Street           string
	PostalCode       string
	CardUserId       string
	Gender           string
	IdCard           string
	IdType           string
	State            string
	ProductId        string
	MaxCardQuota     uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	VipTwo           uint64
	VipThree         uint64
	CardTwo          uint64
	CanVip           uint64
	UserCount        uint64
}

type UserRecommend struct {
	ID            uint64
	UserId        uint64
	RecommendCode string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Config struct {
	ID      uint64
	KeyName string
	Name    string
	Value   string
}

type Withdraw struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	RelAmount float64
	Status    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Card struct {
	ID                  uint64
	CardID              string
	AccountID           string
	CardholderID        string
	BalanceID           string
	BudgetID            string
	ReferenceID         string
	UserName            string
	Currency            string
	Bin                 string
	Status              string
	CardMode            string
	Label               string
	CardLastFour        string
	InterlaceCreateTime int64
	UserId              int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Reward struct {
	ID        uint64
	UserId    uint64
	Amount    float64
	Reason    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Address   string
	One       uint64
}

type EthUserRecord struct {
	ID        int64
	UserId    int64
	Hash      string
	Amount    string
	AmountTwo uint64
	Last      int64
	CreatedAt time.Time
}

type UserRepo interface {
	SetNonceByAddress(ctx context.Context, wallet string) (int64, error)
	GetAndDeleteWalletTimestamp(ctx context.Context, wallet string) (string, error)
	GetConfigByKeys(keys ...string) ([]*Config, error)
	GetUserByAddress(address string) (*User, error)
	GetUserByCard(card string) (*User, error)
	GetUsersStatusDoing() ([]*User, error)
	GetUserByCardUserId(cardUserId string) (*User, error)
	GetUserById(userId uint64) (*User, error)
	GetUserRecommendByUserId(userId uint64) (*UserRecommend, error)
	CreateUser(ctx context.Context, uc *User) (*User, error)
	CreateUserRecommend(ctx context.Context, userId uint64, recommendUser *UserRecommend) (*UserRecommend, error)
	GetUserRecommendByCode(code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(code string) ([]*UserRecommend, error)
	GetUserByUserIds(userIds ...uint64) (map[uint64]*User, error)
	CreateCard(ctx context.Context, userId uint64, user *User) error
	HasCardByCardID(ctx context.Context, cardID string) (bool, error)
	GetAllUsers() ([]*User, error)
	UpdateCard(ctx context.Context, userId uint64, cardOrderId, card string) error
	UpdateCardNo(ctx context.Context, userId uint64, amount float64) error
	UpdateCardSucces(ctx context.Context, userId uint64, cardNum string) error
	CreateCardRecommend(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error
	CreateCardRecommendTwo(ctx context.Context, userId uint64, amount float64, vip uint64, address string) error
	GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error)
	AmountTo(ctx context.Context, userId, toUserId uint64, toAddress string, amount float64) error
	Withdraw(ctx context.Context, userId uint64, amount, amountRel float64, address string) error
	GetUserRewardByUserIdPage(ctx context.Context, b *Pagination, userId uint64, reason uint64) ([]*Reward, error, int64)
	SetVip(ctx context.Context, userId uint64, vip uint64) error
	GetUsersOpenCard() ([]*User, error)
	GetUsersOpenCardStatusDoing() ([]*User, error)
	GetEthUserRecordLast() (int64, error)
	GetUserByAddresses(Addresses ...string) (map[string]*User, error)
	GetUserRecommends() ([]*UserRecommend, error)
	CreateEthUserRecordListByHash(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
	UpdateUserMyTotalAmountAdd(ctx context.Context, userId uint64, amount uint64) error
	UpdateWithdraw(ctx context.Context, id uint64, status string) (*Withdraw, error)
	InsertCardRecord(ctx context.Context, userId, recordType uint64, remark string, code string, opt string) error
	UpdateCardTwo(ctx context.Context, id uint64) error
	GetUserCardTwo() ([]*Reward, error)
	GetUsers(b *Pagination, address string) ([]*User, error, int64)
	GetAdminByAccount(ctx context.Context, account string, password string) (*Admin, error)
	SetCanVip(ctx context.Context, userId uint64, lock uint64) (bool, error)
	SetVipThree(ctx context.Context, userId uint64, vipThree uint64) (bool, error)
	SetUserCount(ctx context.Context, userId uint64) (bool, error)
	GetConfigs() ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
	UpdateUserInfo(ctx context.Context, userId uint64, user *User) error
	CreateCardOne(ctx context.Context, userId uint64, in *Card) error
	CreateCardNew(ctx context.Context, userId, id uint64, in *Card) error
	GetCardPage(ctx context.Context, b *Pagination, accountId, status string) ([]*Card, error, int64)
	GetLatestCard(ctx context.Context) (*Card, error)
	GetCardTwoStatusOne() ([]*CardTwo, error)
}

type UserUseCase struct {
	repo UserRepo
	tx   Transaction
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		tx:   tx,
		log:  log.NewHelper(logger),
	}
}

type Pagination struct {
	PageNum  int
	PageSize int
}

// 后台

func (uuc *UserUseCase) GetEthUserRecordLast() (int64, error) {
	return uuc.repo.GetEthUserRecordLast()
}
func (uuc *UserUseCase) GetUserByAddress(Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(Addresses...)
}

func (uuc *UserUseCase) DepositNew(ctx context.Context, userId uint64, amount uint64, eth *EthUserRecord, system bool) error {
	// 推荐人
	var (
		err error
	)

	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		// 充值记录
		if !system {
			_, err = uuc.repo.CreateEthUserRecordListByHash(ctx, &EthUserRecord{
				Hash:      eth.Hash,
				UserId:    eth.UserId,
				Amount:    eth.Amount,
				AmountTwo: amount,
				Last:      eth.Last,
			})
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资3", userId, amount)
		return err
	}

	// 推荐人
	var (
		userRecommend       *UserRecommend
		tmpRecommendUserIds []string
	)
	userRecommend, err = uuc.repo.GetUserRecommendByUserId(userId)
	if nil != err {
		return err
	}
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	totalTmp := len(tmpRecommendUserIds) - 1
	for i := totalTmp; i >= 0; i-- {
		tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		if 0 >= tmpUserId {
			continue
		}

		// 增加业绩
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateUserMyTotalAmountAdd(ctx, tmpUserId, amount)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("遍历业绩：", err, tmpUserId, eth)
			continue
		}
	}

	return nil
}

var lockHandle sync.Mutex

func (uuc *UserUseCase) OpenCardHandle(ctx context.Context) error {
	lockHandle.Lock()
	defer lockHandle.Unlock()

	var (
		userOpenCard []*User
		err          error
	)

	userOpenCard, err = uuc.repo.GetUsersOpenCard()
	if nil != err {
		return err
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	//var (
	//	products          *CardProductListResponse
	//	productIdUse      string
	//	productIdUseInt64 uint64
	//	maxCardQuota      int
	//)
	//products, err = GetCardProducts()
	//if nil == products || nil != err {
	//	fmt.Println("产品信息错误1")
	//	return nil
	//}
	//
	//for _, v := range products.Rows {
	//	if 0 < len(v.ProductId) && "ENABLED" == v.ProductStatus {
	//		productIdUse = v.ProductId
	//		maxCardQuota = v.MaxCardQuota
	//		productIdUseInt64, err = strconv.ParseUint(productIdUse, 10, 64)
	//		if nil != err {
	//			fmt.Println("产品信息错误2")
	//			return nil
	//		}
	//		fmt.Println("当前选择产品信息", productIdUse, maxCardQuota, v)
	//		break
	//	}
	//}
	//
	//if 0 >= maxCardQuota {
	//	fmt.Println("产品信息错误3")
	//	return nil
	//}
	//
	//if 0 >= productIdUseInt64 {
	//	fmt.Println("产品信息错误4")
	//	return nil
	//}

	for _, user := range userOpenCard {
		//var (
		//	resCreatCardholder *CreateCardholderResponse
		//)
		//resCreatCardholder, err = CreateCardholderRequest(productIdUseInt64, user)
		//if nil == resCreatCardholder || 200 != resCreatCardholder.Code || err != nil {
		//	fmt.Println("持卡人订单创建失败", user, resCreatCardholder, err)
		//	continue
		//}
		//if 0 > len(resCreatCardholder.Data.HolderID) {
		//	fmt.Println("持卡人订单信息错误", user, resCreatCardholder, err)
		//	continue
		//}
		//fmt.Println("持卡人信息", user, resCreatCardholder)
		//

		var (
			holderId          uint64
			productIdUseInt64 uint64
			resCreatCard      *CreateCardResponse
			openRes           = true
		)
		if 5 > len(user.CardUserId) {
			fmt.Println("持卡人id空", user)
			openRes = false
		}
		holderId, err = strconv.ParseUint(user.CardUserId, 10, 64)
		if nil != err {
			fmt.Println("持卡人错误2")
			openRes = false
		}
		if 0 >= holderId {
			fmt.Println("持卡人错误3")
			openRes = false
		}
		if 5 > len(user.CardUserId) {
			fmt.Println("持卡人id空", user)
			openRes = false
		}

		if 0 >= user.MaxCardQuota {
			fmt.Println("最大额度错误", user)
			openRes = false
		}

		if 5 > len(user.ProductId) {
			fmt.Println("productid空", user)
			openRes = false
		}
		productIdUseInt64, err = strconv.ParseUint(user.ProductId, 10, 64)
		if nil != err {
			fmt.Println("产品信息错误1")
			openRes = false
		}
		if 0 >= productIdUseInt64 {
			fmt.Println("产品信息错误2")
			openRes = false
		}

		if !openRes {
			fmt.Println("回滚了用户", user)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}

			continue
		}

		//
		var (
			resHolder *QueryCardHolderResponse
		)

		resHolder, err = QueryCardHolderWithSign(holderId, productIdUseInt64)
		if nil == resHolder || err != nil || 200 != resHolder.Code {
			fmt.Println(user, err, "持卡人信息请求错误", resHolder)
			continue
		}

		if "active" == resHolder.Data.Status {

		} else if "pending" == resHolder.Data.Status {
			continue
		} else {
			fmt.Println(user, err, "持卡人创建失败", resHolder)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		resCreatCard, err = CreateCardRequestWithSign(0, holderId, productIdUseInt64)
		if nil == resCreatCard || 200 != resCreatCard.Code || err != nil {
			fmt.Println("开卡订单创建失败", user, resCreatCard, err)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}
		fmt.Println("开卡信息：", user, resCreatCard)

		if 0 >= len(resCreatCard.Data.CardID) || 0 >= len(resCreatCard.Data.CardOrderID) {
			fmt.Println("开卡订单信息错误", resCreatCard, err)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateCard(ctx, user.ID, resCreatCard.Data.CardOrderID, resCreatCard.Data.CardID)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "开卡后，写入mysql错误", err, user, resCreatCard)
			return nil
		}
	}

	return nil
}

var cardStatusLockHandle sync.Mutex

func (uuc *UserUseCase) CardStatusHandle(ctx context.Context) error {
	cardStatusLockHandle.Lock()
	defer cardStatusLockHandle.Unlock()

	var (
		userOpenCard []*User
		err          error
	)

	userOpenCard, err = uuc.repo.GetUsersOpenCardStatusDoing()
	if nil != err {
		return err
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = uuc.repo.GetAllUsers()
	if nil == users {
		fmt.Println("用户无")
		return nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	for _, user := range userOpenCard {
		// 查询状态。成功分红
		var (
			resCard *CardInfoResponse
		)
		if 2 >= len(user.Card) {
			continue
		}

		resCard, err = GetCardInfoRequestWithSign(user.Card)
		if nil == resCard || 200 != resCard.Code || err != nil {
			fmt.Println(resCard, err)
			continue
		}

		if "ACTIVE" == resCard.Data.CardStatus {
			fmt.Println("开卡状态，激活：", resCard, user.ID)
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.UpdateCardSucces(ctx, user.ID, resCard.Data.Pan)
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err，开卡成功", err, user.ID)
				continue
			}
		} else if "PENDING" == resCard.Data.CardStatus || "PROGRESS" == resCard.Data.CardStatus {
			fmt.Println("开卡状态，待处理：", resCard, user.ID)
			continue
		} else {
			fmt.Println("开卡状态，失败：", resCard, user.ID)
			backAmount := float64(10)
			if 0 < user.VipTwo {
				backAmount = float64(30)
			}
			err = uuc.backCard(ctx, user.ID, backAmount)
			if nil != err {
				fmt.Println("回滚了用户失败", user, err)
			}
			continue
		}

		// 分红
		var (
			userRecommend *UserRecommend
		)
		tmpRecommendUserIds := make([]string, 0)
		// 推荐
		userRecommend, err = uuc.repo.GetUserRecommendByUserId(user.ID)
		if nil == userRecommend {
			fmt.Println(err, "信息错误", err, user)
			return nil
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		tmpTopVip := uint64(10)
		if 30 == user.VipTwo {
			tmpTopVip = 30
		}
		totalTmp := len(tmpRecommendUserIds) - 1
		lastVip := uint64(0)
		for i := totalTmp; i >= 0; i-- {
			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("开卡遍历，信息缺失：", tmpUserId)
				continue
			}

			if usersMap[tmpUserId].VipTwo != user.VipTwo {
				fmt.Println("开卡遍历，信息缺失，不是一个vip区域：", usersMap[tmpUserId], user)
				continue
			}

			if tmpTopVip < usersMap[tmpUserId].Vip {
				fmt.Println("开卡遍历，vip信息设置错误：", usersMap[tmpUserId], lastVip)
				break
			}

			// 小于等于上一个级别，跳过
			if usersMap[tmpUserId].Vip <= lastVip {
				continue
			}

			tmpAmount := usersMap[tmpUserId].Vip - lastVip // 极差
			lastVip = usersMap[tmpUserId].Vip

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.CreateCardRecommend(ctx, tmpUserId, float64(tmpAmount), usersMap[tmpUserId].Vip, user.Address)
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err reward", err, user, usersMap[tmpUserId])
			}
		}
	}

	return nil
}

var cardTwoStatusLockHandle sync.Mutex

func (uuc *UserUseCase) CardTwoStatusHandle(ctx context.Context) error {
	cardTwoStatusLockHandle.Lock()
	defer cardTwoStatusLockHandle.Unlock()

	var (
		userOpenCard []*Reward
		err          error
	)

	var (
		configs       []*Config
		vipThreeThree uint64
		vipThreeTwo   uint64
		vipThreeOne   uint64
	)

	// 配置
	configs, err = uuc.repo.GetConfigByKeys("vip_three_three", "vip_three_two", "vip_three_one")
	if nil != configs {
		for _, vConfig := range configs {
			if "vip_three_three" == vConfig.KeyName {
				vipThreeThree, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "vip_three_two" == vConfig.KeyName {
				vipThreeTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "vip_three_one" == vConfig.KeyName {
				vipThreeOne, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
		}
	}

	userOpenCard, err = uuc.repo.GetUserCardTwo()
	if nil != err {
		return err
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	var (
		users    []*User
		usersMap map[uint64]*User
	)
	users, err = uuc.repo.GetAllUsers()
	if nil == users {
		fmt.Println("开卡2，用户无")
		return nil
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	if 0 >= len(userOpenCard) {
		return nil
	}

	for _, userCard := range userOpenCard {
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateCardTwo(ctx, userCard.ID)
			if err != nil {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward 2", err, userCard)
			continue
		}

		if _, ok := usersMap[userCard.UserId]; !ok {
			fmt.Println("开卡2，信息缺失：", userCard)
			continue
		}
		user := usersMap[userCard.UserId]

		// 分红
		var (
			userRecommend *UserRecommend
		)
		tmpRecommendUserIds := make([]string, 0)
		// 推荐
		userRecommend, err = uuc.repo.GetUserRecommendByUserId(user.ID)
		if nil == userRecommend {
			fmt.Println(err, "开卡2，信息错误", err, user)
			return nil
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		tmpTopVip := uint64(3)
		totalTmp := len(tmpRecommendUserIds) - 1
		lastVip := uint64(0)
		lastAmount := uint64(0)
		for i := totalTmp; i >= 0; i-- {
			if vipThreeThree <= lastAmount {
				break
			}

			tmpUserId, _ := strconv.ParseUint(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("开卡2遍历，信息缺失：", tmpUserId)
				continue
			}

			if tmpTopVip < usersMap[tmpUserId].VipThree {
				fmt.Println("开卡2遍历，vip信息设置错误：", usersMap[tmpUserId], lastVip)
				break
			}

			// 小于等于上一个级别，跳过
			if usersMap[tmpUserId].VipThree <= lastVip {
				continue
			}
			lastVip = usersMap[tmpUserId].VipThree

			// 奖励
			tmpAmount := uint64(0)
			if 1 == usersMap[tmpUserId].VipThree {
				if vipThreeOne <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误1：", usersMap[tmpUserId], lastVip, vipThreeOne, lastAmount)
					continue
				}

				tmpAmount = vipThreeOne - lastAmount
				lastAmount = vipThreeOne
			} else if 2 == usersMap[tmpUserId].VipThree {
				if vipThreeTwo <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误2：", usersMap[tmpUserId], lastVip, vipThreeTwo, lastAmount)
					continue
				}

				tmpAmount = vipThreeTwo - lastAmount
				lastAmount = vipThreeTwo
			} else if 3 == usersMap[tmpUserId].VipThree {
				if vipThreeThree <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误3：", usersMap[tmpUserId], lastVip, vipThreeThree, lastAmount)
					continue
				}

				tmpAmount = vipThreeThree - lastAmount
				lastAmount = vipThreeThree
			} else {
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				err = uuc.repo.CreateCardRecommendTwo(ctx, tmpUserId, float64(tmpAmount), usersMap[tmpUserId].Vip, user.Address)
				if err != nil {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err reward 2", err, user, usersMap[tmpUserId])
			}
		}
	}

	return nil
}

func (uuc *UserUseCase) backCard(ctx context.Context, userId uint64, amount float64) error {
	var (
		err error
	)
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.repo.UpdateCardNo(ctx, userId, amount)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println("err")
		return err
	}

	return nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error) {
	return uuc.repo.GetWithdrawPassOrRewardedFirst(ctx)
}

func (uuc *UserUseCase) GetUserByUserIds(userIds ...uint64) (map[uint64]*User, error) {
	return uuc.repo.GetUserByUserIds(userIds...)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id uint64) (*Withdraw, error) {
	return uuc.repo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id uint64) (*Withdraw, error) {
	return uuc.repo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest, ca string) (*pb.AdminLoginReply, error) {
	var (
		admin *Admin
		err   error
	)

	res := &pb.AdminLoginReply{}
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	fmt.Println(password)
	admin, err = uuc.repo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil != err {
		return res, err
	}

	claims := auth.CustomClaims{
		UserId:   uint64(admin.ID),
		UserType: "admin",
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(48 * time.Hour)), // 2天过期
			Issuer:    "game",
		},
	}

	token, err := auth.CreateToken(claims, ca)
	if err != nil {
		return nil, err
	}
	res.Token = token
	return res, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *pb.AdminRewardListRequest) (*pb.AdminRewardListReply, error) {
	var (
		userSearch  *User
		userId      uint64 = 0
		userRewards []*Reward
		users       map[uint64]*User
		userIdsMap  map[uint64]uint64
		userIds     []uint64
		err         error
		count       int64
	)
	res := &pb.AdminRewardListReply{
		Rewards: make([]*pb.AdminRewardListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(req.Address)
		if nil != err || nil == userSearch {
			return res, nil
		}

		userId = userSearch.ID
	}

	userRewards, err, count = uuc.repo.GetUserRewardByUserIdPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, req.Reason)
	if nil != err {
		return res, nil
	}
	res.Count = uint64(count)

	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		res.Rewards = append(res.Rewards, &pb.AdminRewardListReply_List{
			CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:     fmt.Sprintf("%.2f", vUserReward.Amount),
			Address:    tmpUser,
			Reason:     vUserReward.Reason,
			AddressTwo: vUserReward.Address,
			One:        vUserReward.One,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *pb.AdminUserListRequest) (*pb.AdminUserListReply, error) {
	var (
		users   []*User
		userIds []uint64
		count   int64
		err     error
	)

	res := &pb.AdminUserListReply{
		Users: make([]*pb.AdminUserListReply_UserList, 0),
	}

	users, err, count = uuc.repo.GetUsers(&Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Address)
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		myLowUser         map[uint64][]*UserRecommend
		userRecommendsMap map[uint64]*UserRecommend
	)

	myLowUser = make(map[uint64][]*UserRecommend, 0)
	userRecommendsMap = make(map[uint64]*UserRecommend, 0)

	userRecommends, err = uuc.repo.GetUserRecommends()
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId uint64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		usersAll []*User
		usersMap map[uint64]*User
	)
	usersAll, err = uuc.repo.GetAllUsers()
	if nil == usersAll {
		return nil, nil
	}
	usersMap = make(map[uint64]*User, 0)

	for _, vUsers := range usersAll {
		usersMap[vUsers.ID] = vUsers
	}

	for _, vUsers := range users {
		// 推荐人
		var (
			userRecommend *UserRecommend
		)

		addressMyRecommend := ""
		if _, ok := userRecommendsMap[vUsers.ID]; ok {
			userRecommend = userRecommendsMap[vUsers.ID]

			if nil != userRecommend && "" != userRecommend.RecommendCode {
				var (
					tmpRecommendUserIds   []string
					myUserRecommendUserId uint64
				)
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
				if 2 <= len(tmpRecommendUserIds) {
					myUserRecommendUserId, _ = strconv.ParseUint(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
				}

				if 0 < myUserRecommendUserId {
					if _, ok2 := usersMap[myUserRecommendUserId]; ok2 {
						addressMyRecommend = usersMap[myUserRecommendUserId].Address
					}
				}
			}
		}

		lenUsers := uint64(0)
		if _, ok := myLowUser[vUsers.ID]; ok {
			lenUsers = uint64(len(myLowUser[vUsers.ID]))
		}

		res.Users = append(res.Users, &pb.AdminUserListReply_UserList{
			UserId:             vUsers.ID,
			CreatedAt:          vUsers.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:            vUsers.Address,
			Amount:             fmt.Sprintf("%.2f", vUsers.Amount),
			Vip:                vUsers.Vip,
			CanVip:             vUsers.CanVip,
			VipThree:           vUsers.VipThree,
			MyRecommendAddress: addressMyRecommend,
			HistoryRecommend:   lenUsers,
			MyTotalAmount:      vUsers.MyTotalAmount,
			Card:               vUsers.Card,
			CardNumber:         vUsers.CardNumber,
			CardOrderId:        vUsers.CardOrderId,
			UserCount:          vUsers.UserCount,
			VipTwo:             vUsers.VipTwo,
			CardTwo:            vUsers.CardTwo,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) UpdateUserInfoTo(ctx transporthttp.Context) error {
	var (
		err       error
		accountId = interlaceAccountId
		bins      []*InterlaceCardBin
	)
	// 取原生 *http.Request，后面要用它的 Context 和 FormFile
	r := ctx.Request()

	// 1. 普通表单字段
	userIdS := r.FormValue("userId")
	email := r.FormValue("email")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	dob := r.FormValue("dob")
	gender := r.FormValue("gender")
	nationality := "CN"
	nationalid := r.FormValue("idCard")
	idType := "CN-RIC"
	phoneNumber := r.FormValue("phoneNumber")
	phoneCountryCode := r.FormValue("phoneCountryCode")
	addressLine1 := r.FormValue("addressLine")
	city := r.FormValue("city")
	state := r.FormValue("state")
	country := r.FormValue("country")
	postalCode := r.FormValue("postalCode")

	userId, err := strconv.ParseUint(userIdS, 10, 64)
	if err != nil {
		return err
	}
	var (
		user *User
	)
	user, err = uuc.repo.GetUserById(userId)
	if err != nil {
		return err
	}

	if "do" != user.CardOrderId {
		if err != nil {
			return nil
		}
	}

	////////////////////////////////////////////////////////////////////////////////////////////
	// 2. 取上传文件 (form-data: file)
	file, header, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// 3. 读取文件内容到内存（小文件 OK，大文件可以后面再做限制）
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// 文件名 & Content-Type
	fileName := header.Filename
	mimeType := header.Header.Get("Content-Type") // image/jpeg / image/png 等
	////////////////////////////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////////////////////////////
	// 2. 取上传文件 (form-data: file)
	fileTwo, headerTwo, err := r.FormFile("fileTwo")
	if err != nil {
		return err
	}
	defer fileTwo.Close()

	// 3. 读取文件内容到内存（小文件 OK，大文件可以后面再做限制）
	fileBytesTwo, err := io.ReadAll(fileTwo)
	if err != nil {
		return err
	}

	// 文件名 & Content-Type
	fileNameTwo := headerTwo.Filename
	mimeTypeTwo := headerTwo.Header.Get("Content-Type") // image/jpeg / image/png 等
	////////////////////////////////////////////////////////////////////////////////////////////

	// 3. 拿 accountId（你前面已经实现了 InterlaceGetFirstAccountID）
	//accountId, err = InterlaceGetFirstAccountID(r.Context())
	//if err != nil {
	//	return err
	//}

	// 4. 调用你写好的上传函数（注意：传的是 r.Context()，不是 &context.Context()）
	fileID, err := InterlaceUploadFile(r.Context(), accountId, fileName, mimeType, fileBytes)
	if err != nil {
		return err
	}

	// 4. 调用你写好的上传函数（注意：传的是 r.Context()，不是 &context.Context()）
	fileIDTwo, err := InterlaceUploadFile(r.Context(), accountId, fileNameTwo, mimeTypeTwo, fileBytesTwo)
	if err != nil {
		return err
	}

	bins, err = InterlaceListAvailableBins(ctx, accountId)
	if nil != err {
		fmt.Println(err)
		return err
	}

	for _, v := range bins {
		// 3. 地址
		addr := InterlaceAddress{
			AddressLine1: addressLine1,
			City:         city,
			State:        state,
			Country:      country,
			PostalCode:   postalCode,
		}

		// 4. 创建持卡人（只填必需字段）
		var (
			cardholderId string
		)
		cardholderId, err = InterlaceCreateCardholderMOR(
			ctx,
			v.ID,
			accountId,
			email,
			firstName,
			lastName,
			dob,
			gender,
			nationality,
			nationalid,
			idType,
			addr,
			fileID,
			fileIDTwo,
			phoneNumber,
			phoneCountryCode,
		)
		if nil != err {
			fmt.Println(err, v, cardholderId)
			continue
		}

		if 0 <= len(cardholderId) {
			fmt.Println("持卡人申请错误:", v, cardholderId)
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateUserInfo(ctx, userId, &User{
				ID:               0,
				FirstName:        firstName,
				LastName:         lastName,
				BirthDate:        dob,
				CountryCode:      nationality,
				Phone:            phoneNumber,
				City:             city,
				Country:          country,
				Street:           addressLine1,
				PostalCode:       postalCode,
				Gender:           gender,
				IdCard:           nationalid,
				IdType:           idType,
				State:            state,
				PhoneCountryCode: phoneCountryCode,
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			return err
		}

	}

	return nil

}

func (uuc *UserUseCase) UpdateAllCard(ctx context.Context, req *pb.UpdateAllCardRequest) (*pb.UpdateAllCardReply, error) {
	var (
		cardTwo []*CardTwo
		err     error
	)

	cardTwo, err = uuc.repo.GetCardTwoStatusOne()
	if nil != err {
		fmt.Println("update all card error:", err)
		return nil, err
	}

	// 把第一页也放到统一处理逻辑里
	for _, v := range cardTwo {
		if 10 > len(v.CardId) {
			fmt.Println("cardTwo card error:", cardTwo)
			continue
		}

		var cards []*InterlaceCard

		cards, _, errTwo := InterlaceListCards(ctx, &InterlaceListCardsReq{
			AccountId: interlaceAccountId,
			CardId:    v.CardId,
			Page:      1,
			Limit:     10,
		})
		if nil == cards || errTwo != nil {
			fmt.Println("InterlaceListCards page", "error:", errTwo)
			// 拉失败这一页就跳过，继续后面的
			continue
		}

		if len(cards) == 0 {
			continue
		}

		for _, ic := range cards {
			// 只保留 ACTIVE
			if ic.Status != "ACTIVE" {
				continue
			}

			if ic.CardMode != "PHYSICAL_CARD" {
				fmt.Println("模式错误", ic, v)
				continue
			}

			var tmpCreateTime int64
			tmpCreateTime, err = strconv.ParseInt(ic.CreateTime, 10, 64)
			if 0 >= tmpCreateTime || nil != err {
				fmt.Println("InterlaceListCards create time", ic, "error:", err)
				// 出错就整体结束本次同步，避免老数据乱插
				break
			}

			// 组装 biz.Card 对象
			card := &Card{
				CardID:              ic.ID,
				AccountID:           ic.AccountID,
				CardholderID:        ic.CardholderID,
				BalanceID:           ic.BalanceID,
				BudgetID:            ic.BudgetID,
				ReferenceID:         ic.ReferenceID,
				UserName:            ic.UserName,
				Currency:            ic.Currency,
				Bin:                 ic.Bin,
				Status:              ic.Status,
				CardMode:            ic.CardMode,
				Label:               ic.Label,
				CardLastFour:        ic.CardLastFour,
				InterlaceCreateTime: tmpCreateTime, // 毫秒时间戳
			}

			if errThree := uuc.repo.CreateCardNew(ctx, v.UserId, v.ID, card); errThree != nil {
				fmt.Println("CreateCard error, cardID =", ic.ID, "err =", err)
				// 这条失败就算了，不影响其它
				continue
			}
		}
	}

	return &pb.UpdateAllCardReply{}, nil
}

func (uuc *UserUseCase) UpdateAllCardTwo(ctx context.Context, req *pb.UpdateAllCardRequest) (*pb.UpdateAllCardReply, error) {
	var (
		users []*User
		err   error
	)

	users, err = uuc.repo.GetUsersStatusDoing()
	if nil != err {
		fmt.Println("update all card error:", err)
		return nil, err
	}

	// 把第一页也放到统一处理逻辑里
	for _, v := range users {
		if 10 > len(v.CardNumber) {
			fmt.Println("cardOne card error:", v)
			continue
		}

		var cards []*InterlaceCard

		cards, _, errTwo := InterlaceListCards(ctx, &InterlaceListCardsReq{
			AccountId: interlaceAccountId,
			CardId:    v.CardNumber,
			Page:      1,
			Limit:     10,
		})
		if nil == cards || errTwo != nil {
			fmt.Println("InterlaceListCards page", "error:", errTwo)
			// 拉失败这一页就跳过，继续后面的
			continue
		}

		if len(cards) == 0 {
			continue
		}

		for _, ic := range cards {
			// 只保留 ACTIVE
			if ic.Status != "ACTIVE" {
				continue
			}

			if ic.CardMode != "VIRTUAL_CARD" {
				fmt.Println("模式错误", ic, v)
				continue
			}

			var tmpCreateTime int64
			tmpCreateTime, err = strconv.ParseInt(ic.CreateTime, 10, 64)
			if 0 >= tmpCreateTime || nil != err {
				fmt.Println("InterlaceListCards create time", ic, "error:", err)
				// 出错就整体结束本次同步，避免老数据乱插
				break
			}

			// 组装 biz.Card 对象
			card := &Card{
				CardID:              ic.ID,
				AccountID:           ic.AccountID,
				CardholderID:        ic.CardholderID,
				BalanceID:           ic.BalanceID,
				BudgetID:            ic.BudgetID,
				ReferenceID:         ic.ReferenceID,
				UserName:            ic.UserName,
				Currency:            ic.Currency,
				Bin:                 ic.Bin,
				Status:              ic.Status,
				CardMode:            ic.CardMode,
				Label:               ic.Label,
				CardLastFour:        ic.CardLastFour,
				InterlaceCreateTime: tmpCreateTime, // 毫秒时间戳
			}

			if errThree := uuc.repo.CreateCardOne(ctx, v.ID, card); errThree != nil {
				fmt.Println("CreateCardOne error, cardID =", ic.ID, "err =", err)
				// 这条失败就算了，不影响其它
				continue
			}
		}
	}

	return &pb.UpdateAllCardReply{}, nil
}

func (uuc *UserUseCase) UpdateCanVip(ctx context.Context, req *pb.UpdateCanVipRequest) (*pb.UpdateCanVipReply, error) {
	var (
		err  error
		lock uint64
	)

	res := &pb.UpdateCanVipReply{}

	if 1 == req.SendBody.CanVip {
		lock = 1
	} else {
		lock = 0
	}

	_, err = uuc.repo.SetCanVip(ctx, req.SendBody.UserId, lock)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) SetVipThree(ctx context.Context, req *pb.SetVipThreeRequest) (*pb.SetVipThreeReply, error) {
	var (
		err  error
		lock uint64
	)

	res := &pb.SetVipThreeReply{}

	if 1 == req.SendBody.VipThree {
		lock = 1
	} else if 2 == req.SendBody.VipThree {
		lock = 2
	} else if 3 == req.SendBody.VipThree {
		lock = 3
	} else {
		lock = 0
	}

	_, err = uuc.repo.SetVipThree(ctx, req.SendBody.UserId, lock)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *pb.AdminConfigUpdateRequest) (*pb.AdminConfigUpdateReply, error) {
	var (
		err error
	)

	res := &pb.AdminConfigUpdateReply{}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = uuc.repo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *pb.AdminConfigRequest) (*pb.AdminConfigReply, error) {
	var (
		configs []*Config
	)

	res := &pb.AdminConfigReply{
		Config: make([]*pb.AdminConfigReply_List, 0),
	}

	configs, _ = uuc.repo.GetConfigs()
	if nil == configs {
		return res, nil
	}

	for _, v := range configs {
		res.Config = append(res.Config, &pb.AdminConfigReply_List{
			Id:    int64(v.ID),
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) SetUserCount(ctx context.Context, req *pb.SetUserCountRequest) (*pb.SetUserCountReply, error) {
	var (
		err error
	)

	res := &pb.SetUserCountReply{}

	_, err = uuc.repo.SetUserCount(ctx, req.SendBody.UserId)
	if nil != err {
		return res, err
	}

	return res, nil
}

type CardUserHandle struct {
	MerchantId string `json:"merchantId"`
	HolderId   string `json:"holderId"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

type CardCreateData struct {
	MerchantId string `json:"merchantId"`
	//ReferenceCode string `json:"referenceCode"`
	Remark     string `json:"remark"`
	CardId     string `json:"cardId"`
	CardNumber string `json:"cardNumber"`
	//Opt string `json:"opt"`
}

type RechargeData struct {
	MerchantId string `json:"merchantId"`
	//ReferenceCode string `json:"referenceCode"`
	//Opt string `json:"opt"`
	Remark     string `json:"remark"`
	CardId     string `json:"cardId"`
	CardNumber string `json:"cardNumber"`
}

func (uuc *UserUseCase) CallBackHandleOne(ctx context.Context, r *CardUserHandle) error {
	fmt.Println("结果：", r)
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByCardUserId(r.HolderId)
	if nil != err {
		fmt.Println("回调，不存在用户", r, err)
		return nil
	}

	err = uuc.repo.InsertCardRecord(ctx, user.ID, 1, r.Remark, "", "")
	if nil != err {
		fmt.Println("回调，新增失败", r, err)
		return nil
	}

	return nil
}

func (uuc *UserUseCase) CallBackHandleTwo(ctx context.Context, r *CardCreateData) error {
	fmt.Println("结果：", r)
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByCard(r.CardId)
	if nil != err {
		fmt.Println("回调，不存在用户", r, err)
		return nil
	}

	err = uuc.repo.InsertCardRecord(ctx, user.ID, 2, r.Remark, "", "")
	if nil != err {
		fmt.Println("回调，新增失败", r, err)
		return nil
	}

	return nil
}

func (uuc *UserUseCase) CallBackHandleThree(ctx context.Context, r *RechargeData) error {
	fmt.Println("结果：", r)
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByCard(r.CardId)
	if nil != err {
		fmt.Println("回调，不存在用户", r, err)
		return nil
	}

	err = uuc.repo.InsertCardRecord(ctx, user.ID, 3, r.Remark, "", "")
	if nil != err {
		fmt.Println("回调，新增失败", r, err)
		return nil
	}

	return nil
}

func GenerateSign(params map[string]interface{}, signKey string) string {
	// 1. 排除 sign 字段
	var keys []string
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接 key + value 字符串
	var sb strings.Builder
	sb.WriteString(signKey)

	for _, k := range keys {
		sb.WriteString(k)
		value := params[k]

		var strValue string
		switch v := value.(type) {
		case string:
			strValue = v
		case float64, int, int64, bool:
			strValue = fmt.Sprintf("%v", v)
		default:
			// map、slice 等复杂类型用 JSON 编码
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				strValue = ""
			} else {
				strValue = string(jsonBytes)
			}
		}
		sb.WriteString(strValue)
	}

	signString := sb.String()
	//fmt.Println("md5前字符串", signString)

	// 3. 进行 MD5 加密
	hash := md5.Sum([]byte(signString))
	return hex.EncodeToString(hash[:])
}

type CreateCardResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CardID      string `json:"cardId"`
		CardOrderID string `json:"cardOrderId"`
		CreateTime  string `json:"createTime"`
		CardStatus  string `json:"cardStatus"`
		OrderStatus string `json:"orderStatus"`
	} `json:"data"`
}

func CreateCardRequestWithSign(cardAmount uint64, cardholderId uint64, cardProductId uint64) (*CreateCardResponse, error) {
	//url := "https://test-api.ispay.com/dev-api/vcc/api/v1/cards/create"
	//url := "https://www.ispay.com/prod-api/vcc/api/v1/cards/create"
	baseUrl := "http://120.79.173.55:9102/prod-api/vcc/api/v1/cards/create"

	reqBody := map[string]interface{}{
		"merchantId":    "322338",
		"cardCurrency":  "USD",
		"cardAmount":    cardAmount,
		"cardholderId":  cardholderId,
		"cardProductId": cardProductId,
		"cardSpendRule": map[string]interface{}{
			"dailyLimit":   250000,
			"monthlyLimit": 1000000,
		},
		"cardRiskControl": map[string]interface{}{
			"allowedMerchants": []string{"ONLINE"},
			"blockedCountries": []string{},
		},
	}

	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")
	// 请求体（包括嵌套结构）
	reqBody["sign"] = sign

	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Language", "zh_CN")

	//fmt.Println("请求报文:", string(jsonData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		errTwo := Body.Close()
		if errTwo != nil {

		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	fmt.Println("响应报文:", string(body)) // ← 打印响应内容

	var result CreateCardResponse
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("开卡，JSON 解析失败:", err)
		return nil, err
	}

	return &result, nil
}

type CardInfoResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CardID     string `json:"cardId"`
		Pan        string `json:"pan"`
		CardStatus string `json:"cardStatus"`
		Holder     struct {
			HolderID string `json:"holderId"`
		} `json:"holder"`
	} `json:"data"`
}

func GetCardInfoRequestWithSign(cardId string) (*CardInfoResponse, error) {
	baseUrl := "http://120.79.173.55:9102/prod-api/vcc/api/v1/cards/info"
	//baseUrl := "https://www.ispay.com/prod-api/vcc/api/v1/cards/info"

	reqBody := map[string]interface{}{
		"merchantId": "322338",
		"cardId":     cardId, // 如果需要传 cardId，根据实际接口文档添加
	}

	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")
	reqBody["sign"] = sign

	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Language", "zh_CN")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		errTwo := Body.Close()
		if errTwo != nil {

		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", string(body))
	}

	//fmt.Println("响应报文:", string(body))

	var result CardInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("卡信息 JSON 解析失败:", err)
		return nil, err
	}

	return &result, nil
}

type CardHolderData struct {
	HolderId    string `json:"holderId"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birthDate"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
	Status      string `json:"status"`
}

type QueryCardHolderResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data CardHolderData `json:"data"`
}

func QueryCardHolderWithSign(holderId uint64, productId uint64) (*QueryCardHolderResponse, error) {
	baseUrl := "https://www.ispay.com/prod-api/vcc/api/v1/cards/holders/query"

	// 请求体
	reqBody := map[string]interface{}{
		"holderId":   holderId,
		"merchantId": "322338",
		"productId":  productId,
	}

	// 生成签名
	sign := GenerateSign(reqBody, "j4gqNRcpTDJr50AP2xd9obKWZIKWbeo9")
	reqBody["sign"] = sign

	// 转 JSON
	jsonData, _ := json.Marshal(reqBody)

	// 打印调试
	//fmt.Println("签名:", sign)
	//fmt.Println("请求报文:", string(jsonData))

	// 创建 HTTP 请求
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Language", "zh_CN")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 读取响应
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	//fmt.Println("响应报文:", string(body))

	// 解析结果
	var result QueryCardHolderResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ================= Interlace 授权配置 & 缓存 =================

const (
	interlaceBaseURL      = "https://api-sandbox.interlace.money/open-api/v3"
	interlaceBaseURLV1    = "https://api-sandbox.interlace.money/open-api/v1"
	interlaceClientID     = "interlacedc0330757f216112"
	interlaceClientSecret = "c0d8019217ad4903bf09336320a4ddd9" // v3 的接口目前用不到 secret，但建议以后放到配置/环境变量
	interlaceAccountId    = "cb6c8028-c828-4596-a501-6fa3196af4d7"
)

// 缓存在当前进程里，如果你将来多实例部署/重启频繁，可以再扩展成 Redis 存储
type interlaceAuthCache struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     int64 // unix 秒，提前留一点余量
}

var (
	interlaceAuth    = &interlaceAuthCache{}
	interlaceAuthMux sync.Mutex
)

// GetInterlaceAccessToken 获取一个当前可用的 accessToken
// 1. 如果缓存里有且没过期，直接返回
// 2. 否则调用 GetCode + Generate Access Token 重新获取
func GetInterlaceAccessToken(ctx context.Context) (string, error) {
	interlaceAuthMux.Lock()
	defer interlaceAuthMux.Unlock()

	now := time.Now().Unix()
	// 缓存未过期，直接用（提前 60 秒过期，避免边界）
	if interlaceAuth.AccessToken != "" && now < interlaceAuth.ExpireAt-60 {
		return interlaceAuth.AccessToken, nil
	}

	// 这里可以先尝试用 refreshToken 刷新（如果你想用 refresh-token 接口）
	// 为了简单稳定，这里直接重新 Get Code + Access Token
	code, err := interlaceGetCode(ctx)
	if err != nil {
		return "", fmt.Errorf("get interlace code failed: %w", err)
	}

	accessToken, refreshToken, expiresIn, err := interlaceGenerateAccessToken(ctx, code)
	if err != nil {
		return "", fmt.Errorf("generate interlace access token failed: %w", err)
	}

	interlaceAuth.AccessToken = accessToken
	interlaceAuth.RefreshToken = refreshToken
	interlaceAuth.ExpireAt = time.Now().Unix() + expiresIn

	return accessToken, nil
}

// Get a code 响应结构
type interlaceGetCodeResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Timestamp int64  `json:"timestamp"`
		Code      string `json:"code"`
	} `json:"data"`
}

func interlaceGetCode(ctx context.Context) (string, error) {
	urlStr := fmt.Sprintf("%s/oauth/authorize?clientId=%s", interlaceBaseURL, interlaceClientID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("interlace get code http %d: %s", resp.StatusCode, string(body))
	}

	var result interlaceGetCodeResp
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("interlace get code unmarshal: %w", err)
	}

	if result.Code != "000000" {
		return "", fmt.Errorf("interlace get code failed: code=%s msg=%s", result.Code, result.Message)
	}
	if result.Data.Code == "" {
		return "", fmt.Errorf("interlace get code success but orderId empty")
	}

	return result.Data.Code, nil
}

// Generate an access token 响应结构
type interlaceAccessTokenResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int64  `json:"expiresIn"` // 有效期秒数，比如 86400
		Timestamp    int64  `json:"timestamp"`
	} `json:"data"`
}

func interlaceGenerateAccessToken(ctx context.Context, code string) (accessToken, refreshToken string, expiresIn int64, err error) {
	urlStr := fmt.Sprintf("%s/oauth/access-token", interlaceBaseURL)

	reqBody := map[string]interface{}{
		"clientId": interlaceClientID,
		"code":     code,
	}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, bytes.NewReader(jsonData))
	if err != nil {
		return "", "", 0, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", 0, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", 0, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", "", 0, fmt.Errorf("interlace access-token http %d: %s", resp.StatusCode, string(body))
	}

	var result interlaceAccessTokenResp
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", 0, fmt.Errorf("interlace access-token unmarshal: %w", err)
	}

	if result.Code != "000000" {
		return "", "", 0, fmt.Errorf("interlace access-token failed: code=%s msg=%s", result.Code, result.Message)
	}
	if result.Data.AccessToken == "" {
		return "", "", 0, fmt.Errorf("interlace access-token success but accessToken empty")
	}

	return result.Data.AccessToken, result.Data.RefreshToken, result.Data.ExpiresIn, nil
}

func InterlaceCreateCardholder(ctx context.Context, token string, user *User) (string, error) {
	urlStr := interlaceBaseURL + "/cardholders"

	reqBody := map[string]interface{}{
		"programType": "BUSINESS USE - MOR", // 你用的是商户代收付 Mor 模式
		// "binId": ...,
		"name": map[string]interface{}{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
		},
		"email": user.Email,
		// 其它字段按文档补
	}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("create cardholder http %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Code    string          `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if result.Code != "000000" {
		return "", fmt.Errorf("create cardholder failed: %s", result.Message)
	}

	// 解析真实的 cardholderId 字段（按 Cardholder 文档来）
	var data struct {
		CardholderID string `json:"cardholderId"`
	}
	if err := json.Unmarshal(result.Data, &data); err != nil {
		return "", err
	}
	if data.CardholderID == "" {
		return "", fmt.Errorf("cardholderId empty")
	}

	return data.CardholderID, nil
}

// InterlaceCardBin 用来接住 List all available card BINs 返回里的单个 BIN 信息。
// 接口文档保证至少有 binId 字段，其他字段我们先按常见命名猜一猜，多出来没关系。
type InterlaceCardBin struct {
	ID                  string   `json:"id"`
	Bin                 string   `json:"bin"`
	Type                int      `json:"type"` // 0/1
	Currencies          []string `json:"currencies"`
	Network             string   `json:"network"`
	SupportPhysicalCard bool     `json:"supportPhysicalCard"`

	Verification struct {
		Avs     bool `json:"avs"`
		ThreeDs bool `json:"threeDs"`
	} `json:"verification"`

	PurchaseLimit struct {
		Day      string `json:"day"`
		Single   string `json:"single"`
		Lifetime string `json:"lifetime"`
	} `json:"purchaseLimit"`
}

// InterlaceListAvailableBins 使用 x-access-token + accountId 获取可用 BIN
func InterlaceListAvailableBins(ctx context.Context, accountId string) ([]*InterlaceCardBin, error) {
	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		fmt.Println("获取access token错误")
		return nil, err
	}

	base := interlaceBaseURL + "/card/bins"
	q := url.Values{}
	q.Set("accountId", accountId)
	urlStr := base + "?" + q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-access-token", accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("interlace list card bins http %d: %s", resp.StatusCode, string(body))
	}

	//fmt.Println("interlace list card bins body:", string(body))

	var outer struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List  []InterlaceCardBin `json:"list"`
			Total string             `json:"total"` // 注意这里
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &outer); err != nil {
		return nil, fmt.Errorf("list card bins unmarshal: %w", err)
	}
	if outer.Code != "000000" {
		return nil, fmt.Errorf("list card bins failed: code=%s msg=%s", outer.Code, outer.Message)
	}

	bins := make([]*InterlaceCardBin, 0, len(outer.Data.List))
	for i := range outer.Data.List {
		b := outer.Data.List[i]
		bins = append(bins, &b)
	}
	return bins, nil
}

// InterlaceGetFirstAccountID 调用 v1 /accounts，返回一个可用的 accountId
// 当前返回示例：{"code":0,"message":"ok","data":{"data":[{...}],"pageTotal":1,"total":1}}
func InterlaceGetFirstAccountID(ctx context.Context) (string, error) {
	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		fmt.Println("获取access token错误", err)
		return "", err
	}

	urlStr := interlaceBaseURLV1 + "/accounts"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-access-token", accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//fmt.Println("interlace v1 list accounts status:", resp.StatusCode)
	//fmt.Println("interlace v1 list accounts body:", string(body))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("interlace v1 accounts http %d: %s", resp.StatusCode, string(body))
	}

	// 按真实结构定义一个类型
	type accountItem struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Status string `json:"status"`
		Name   string `json:"name"`
	}

	type accountsResp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Data      []accountItem `json:"data"`
			PageTotal int           `json:"pageTotal"`
			Total     int           `json:"total"`
		} `json:"data"`
	}

	var res accountsResp
	if err := json.Unmarshal(body, &res); err != nil {
		return "", fmt.Errorf("accounts unmarshal error: %w", err)
	}

	// code == 0 是成功
	if res.Code != 0 {
		return "", fmt.Errorf("accounts api failed: code=%d msg=%s", res.Code, res.Message)
	}

	if len(res.Data.Data) == 0 {
		return "", fmt.Errorf("accounts list empty")
	}

	// 找一个 Active 的账号（你目前看到就是 ApiClient/Active 那个）
	for _, acc := range res.Data.Data {
		if acc.Status == "Active" && acc.ID != "" {
			//fmt.Println("use accountId:", acc.ID, "type:", acc.Type, "name:", acc.Name)
			return acc.ID, nil
		}
	}

	return "", fmt.Errorf("no active account found")
}

// InterlaceCreateConsumerCardholder
// 按照 MoR Consumer 模式，只用「必需字段」创建 cardholder，返回 cardholderId。
func InterlaceCreateConsumerCardholder(ctx context.Context, u *User, bin *InterlaceCardBin) (string, error) {
	if u == nil {
		return "", fmt.Errorf("user is nil")
	}
	if bin == nil || bin.ID == "" {
		return "", fmt.Errorf("bin is nil or bin.ID empty")
	}

	// 1) 拿 OAuth accessToken（Bearer，用于 /v3/cardholders）
	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		fmt.Println("InterlaceCreateConsumerCardholder: 获取 access token 错误:", err)
		return "", fmt.Errorf("get access token failed: %w", err)
	}

	// 2) 处理国家 / 区号（libphonenumber 要求 countryCode 是区号，不是 "CN" 这种）
	// 你现在 user.CountryCode 字段存的是 "CN"/"HK" 这一类，为了先跑通，这里简单映射一下。
	nationality := u.CountryCode
	if nationality == "" {
		nationality = "CN"
	}
	countryOfResidence := nationality

	phoneCountryCode := u.CountryCode
	if phoneCountryCode == "" || phoneCountryCode == "CN" {
		phoneCountryCode = "+86"
	}

	// 3) 构建请求体 —— 字段名严格对齐官方文档
	reqBody := map[string]interface{}{
		"programType": "CONSUMER USE - MOR", // MoR Consumer 模式固定值

		"binId": bin.ID, // 从 /card/bins 返回的某个 list[i].id

		"name": map[string]interface{}{
			"firstName": u.FirstName, // 你的 User 里已有
			"lastName":  u.LastName,
		},

		"email": u.Email,

		"phone": map[string]interface{}{
			"countryCode": phoneCountryCode, // 例如 "+86"
			"phoneNumber": u.Phone,          // 例如 "13077000000"
		},

		"dateOfBirth": u.BirthDate, // "1983-10-10"

		"nationality":        nationality,        // "CN" / "HK" / "US"...
		"countryOfResidence": countryOfResidence, // 一般和 nationality 一致

		"address": map[string]interface{}{
			"country":    nationality, // ISO2国家码
			"state":      "",          // 先留空，有需要你用省份填
			"city":       u.City,
			"street":     u.Street,
			"postalCode": u.PostalCode,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal cardholder body error: %w", err)
	}

	// 4) 发送 HTTP 请求
	urlStr := interlaceBaseURL + "/cardholders" // v3
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("interlace create cardholder status:", resp.StatusCode)
	fmt.Println("interlace create cardholder body:", string(respBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("create cardholder http %d: %s", resp.StatusCode, string(respBody))
	}

	// 5) 解析响应，只关心 code + data.id
	var outer struct {
		Code    string          `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(respBody, &outer); err != nil {
		return "", fmt.Errorf("create cardholder unmarshal resp: %w", err)
	}
	if outer.Code != "000000" {
		return "", fmt.Errorf("create cardholder failed: code=%s msg=%s", outer.Code, outer.Message)
	}

	// data 里至少会有 id
	var data struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(outer.Data, &data); err != nil {
		fmt.Println("create cardholder data raw:", string(outer.Data))
		return "", fmt.Errorf("create cardholder data unmarshal: %w", err)
	}
	if data.ID == "" {
		return "", fmt.Errorf("create cardholder success but id empty")
	}

	fmt.Println("interlace cardholderId:", data.ID)
	return data.ID, nil
}

type InterlaceAddress struct {
	AddressLine1 string
	City         string
	State        string
	Country      string // ISO2, 如 "US"/"CN"
	PostalCode   string
}

// InterlaceCreateCardholderMOR
// 只用必填字段创建 MoR Consumer 持卡人，返回 cardholderId。
func InterlaceCreateCardholderMOR(
	ctx context.Context,
	binId string,
	accountId string,
	email string,
	firstName string,
	lastName string,
	dob string, // YYYY-MM-DD
	gender string, // "M" / "F"
	nationality string, // ISO2, e.g. "CN"
	nationalId string,
	idType string, // "CN-RIC" / "PASSPORT" / ...
	addr InterlaceAddress,
	idFrontId string,
	selfie string,
	phoneNumber string, // 不带国家码
	phoneCountryCode string,
) (string, error) {

	if binId == "" || accountId == "" {
		return "", fmt.Errorf("binId/accountId required")
	}
	if email == "" || firstName == "" || lastName == "" || dob == "" || gender == "" {
		return "", fmt.Errorf("email/firstName/lastName/dob/gender required")
	}
	if nationality == "" || nationalId == "" || idType == "" {
		return "", fmt.Errorf("nationality/nationalId/idType required")
	}
	if addr.AddressLine1 == "" || addr.City == "" || addr.State == "" || addr.Country == "" || addr.PostalCode == "" {
		return "", fmt.Errorf("address fields required")
	}
	if idFrontId == "" || selfie == "" || phoneNumber == "" {
		return "", fmt.Errorf("idFrontId/selfie/phoneNumber required")
	}

	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		return "", fmt.Errorf("get access token failed: %w", err)
	}

	body := map[string]interface{}{
		"binId":         binId,
		"accountId":     accountId,
		"businessModel": "B2C_MOR",

		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
		"dob":       dob,
		"gender":    gender,

		"nationality": nationality,
		"nationalId":  nationalId,
		"idType":      idType,

		"address": map[string]interface{}{
			"addressLine1": addr.AddressLine1,
			"city":         addr.City,
			"state":        addr.State,
			"country":      addr.Country,
			"postalCode":   addr.PostalCode,
		},

		"idFrontId":        idFrontId,
		"selfie":           selfie,
		"phoneNumber":      phoneNumber,
		"phoneCountryCode": phoneCountryCode,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal cardholder body error: %w", err)
	}

	urlStr := interlaceBaseURL + "/cardholders"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// Create cardholder 文档用的是 x-access-token
	req.Header.Set("x-access-token", accessToken)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//fmt.Println("interlace create cardholder status:", resp.StatusCode)
	fmt.Println("interlace create cardholder body:", string(respBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("create cardholder http %d: %s", resp.StatusCode, string(respBody))
	}

	var outer struct {
		Code    string          `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(respBody, &outer); err != nil {
		return "", fmt.Errorf("create cardholder unmarshal resp: %w", err)
	}
	if outer.Code != "000000" && outer.Code != "0" {
		return "", fmt.Errorf("create cardholder failed: code=%s msg=%s", outer.Code, outer.Message)
	}

	var data struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(outer.Data, &data); err != nil {
		fmt.Println("create cardholder data raw:", string(outer.Data))
		return "", fmt.Errorf("create cardholder data unmarshal: %w", err)
	}
	if data.ID == "" {
		return "", fmt.Errorf("create cardholder success but id empty")
	}

	//fmt.Println("cardholderId:", data.ID)
	return data.ID, nil
}

// 上传一个文件到 Interlace，返回 fileId（用于 idFrontId / selfie 等）
func InterlaceUploadFile(ctx context.Context, accountId, fileName, mimeType string, fileData []byte) (string, error) {
	if accountId == "" {
		return "", fmt.Errorf("accountId required")
	}
	if len(fileData) == 0 {
		return "", fmt.Errorf("fileData is empty")
	}

	// 1) 拿 accessToken（后面用 x-access-token）
	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		return "", fmt.Errorf("get access token failed: %w", err)
	}

	// 2) 构造 multipart/form-data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// files 字段（必填）
	part, err := writer.CreateFormFile("files", fileName)
	if err != nil {
		return "", fmt.Errorf("create form file error: %w", err)
	}
	if _, err := part.Write(fileData); err != nil {
		return "", fmt.Errorf("write file data error: %w", err)
	}

	// accountId 字段（必填）
	if err := writer.WriteField("accountId", accountId); err != nil {
		return "", fmt.Errorf("write accountId field error: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("multipart writer close error: %w", err)
	}

	// 3) 发请求
	urlStr := interlaceBaseURL + "/files/upload"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-access-token", accessToken)

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//fmt.Println("interlace upload file status:", resp.StatusCode)
	//fmt.Println("interlace upload file body:", string(bodyBytes))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("upload file http %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// 4) 解析响应：{code, message, data: [...] } 或 {code, message, data:{id:...}}
	var outer struct {
		Code    string          `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(bodyBytes, &outer); err != nil {
		return "", fmt.Errorf("upload file unmarshal resp: %w", err)
	}
	if outer.Code != "000000" && outer.Code != "0" {
		return "", fmt.Errorf("upload file failed: code=%s msg=%s", outer.Code, outer.Message)
	}

	// data 可能是数组，也可能是对象，两种都试一下
	var arr []struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(outer.Data, &arr); err == nil && len(arr) > 0 && arr[0].ID != "" {
		return arr[0].ID, nil
	}

	var single struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(outer.Data, &single); err == nil && single.ID != "" {
		return single.ID, nil
	}

	// 文档只写 code/message 的话，这里可能拿不到 id，就先抛错
	return "", fmt.Errorf("upload file success but file id not found in data")
}

// 列表请求入参（你业务内部用）
type InterlaceListCardsReq struct {
	AccountId string // 必填

	CardId       string // 可选
	BudgetId     string // 可选
	CardholderId string // 可选
	Label        string // 可选
	ReferenceId  string // 可选

	Limit int // 1-100，默认 10
	Page  int // >=1，默认 1
}

// 单张卡片信息（输出）
type InterlaceCard struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Status    string `json:"status"`   // INACTIVE, CONTROL, ACTIVE, PENDING, FROZEN
	Currency  string `json:"currency"` // 货币代码
	Bin       string `json:"bin"`

	UserName     string `json:"userName"`
	CreateTime   string `json:"createTime"`
	CardLastFour string `json:"cardLastFour"`

	BillingAddress *InterlaceBillingAddress `json:"billingAddress"`

	Label        string `json:"label"`
	BalanceID    string `json:"balanceId"`
	BudgetID     string `json:"budgetId"`
	CardholderID string `json:"cardholderId"`
	ReferenceID  string `json:"referenceId"`

	CardMode string `json:"cardMode"` // PHYSICAL_CARD / VIRTUAL_CARD

	TransactionLimits []InterlaceTransactionLimit `json:"transactionLimits"`
}

// 账单地址
type InterlaceBillingAddress struct {
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	Country      string `json:"country,omitempty"`
}

// 单个额度限制
type InterlaceTransactionLimit struct {
	Type     string `json:"type"`     // DAY/WEEK/MONTH/QUARTER/YEAR/LIFETIME/TRANSACTION/NA
	Value    string `json:"value"`    // 金额（字符串）
	Currency string `json:"currency"` // 货币
}

// InterlaceListCards 使用 x-access-token + accountId 获取卡片列表
func InterlaceListCards(ctx context.Context, in *InterlaceListCardsReq) ([]*InterlaceCard, string, error) {
	if in == nil {
		return nil, "", fmt.Errorf("list cards req is nil")
	}
	if in.AccountId == "" {
		return nil, "", fmt.Errorf("accountId is required")
	}

	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		fmt.Println("获取access token错误")
		return nil, "", err
	}

	base := interlaceBaseURL + "/card-list"

	q := url.Values{}
	q.Set("accountId", in.AccountId)

	if in.CardId != "" {
		q.Set("cardId", in.CardId)
	}
	if in.BudgetId != "" {
		q.Set("budgetId", in.BudgetId)
	}
	if in.CardholderId != "" {
		q.Set("cardholderId", in.CardholderId)
	}
	if in.Label != "" {
		q.Set("label", in.Label)
	}
	if in.ReferenceId != "" {
		q.Set("referenceId", in.ReferenceId)
	}

	limit := in.Limit
	if limit <= 0 {
		limit = 10
	}
	page := in.Page
	if page <= 0 {
		page = 1
	}
	q.Set("limit", fmt.Sprintf("%d", limit))
	q.Set("page", fmt.Sprintf("%d", page))

	urlStr := base + "?" + q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-access-token", accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, "", fmt.Errorf("interlace list cards http %d: %s", resp.StatusCode, string(body))
	}

	// 外层通用结构：code/message/data
	var outer struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List  []InterlaceCard `json:"list"`
			Total string          `json:"total"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &outer); err != nil {
		return nil, "", fmt.Errorf("list cards unmarshal: %w", err)
	}
	if outer.Code != "000000" {
		return nil, "", fmt.Errorf("list cards failed: code=%s msg=%s", outer.Code, outer.Message)
	}

	cards := make([]*InterlaceCard, 0, len(outer.Data.List))
	for i := range outer.Data.List {
		c := outer.Data.List[i]
		cards = append(cards, &c)
	}

	return cards, outer.Data.Total, nil
}
