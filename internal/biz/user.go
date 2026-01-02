package biz

import (
	"bytes"
	pb "cardbinance/api/user/v1"
	"cardbinance/internal/pkg/middleware/auth"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	imapid "github.com/ProtonMail/go-imap-id"
	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/go-kratos/kratos/v2/log"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"html"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
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
	CardAmount       float64
	IdCard           string
	Gender           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CardOrder struct {
	ID        uint64
	Last      uint64
	Code      string
	Card      string
	Time      *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
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
	CardTwoNumber    string
	LockCard         uint64
	LockCardTwo      uint64
	ChangeCard       uint64
	ChangeCardTwo    uint64
	Pic              string
	PicTwo           string
	CardNumberRel    string
	CardNumberRelTwo string
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
	GetUserByUserIdsTwo(userIds []uint64) (map[uint64]*User, error)
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
	GetUsers(b *Pagination, address string, cardTwo uint64, cardOrderId string, lockCard uint64, changeCard uint64) ([]*User, error, int64)
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
	UpdateUserDoing(ctx context.Context, userId uint64, cardNumber, cardNumberRel string, cardAmount float64) error
	UpdateCardStatus(ctx context.Context, id, userId uint64, cardNumber, cardNumberRel string, cardAmount float64) error
	GetCardTwos(b *Pagination, userId uint64, status uint64, cardId string) ([]*CardTwo, error, int64)
	GetCardTwoById(id uint64) (*CardTwo, error)
	GetCardOrder() (*CardOrder, error)
	CreateCardOrder(ctx context.Context, in *CardOrder) error
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
		vipThreeFive  uint64
		vipThreeFour  uint64
		vipThreeThree uint64
		vipThreeTwo   uint64
		vipThreeOne   uint64
		vipThreeSix   uint64
		vipThreeSeven uint64
	)

	// 配置
	configs, err = uuc.repo.GetConfigByKeys("new_vip_three_seven", "new_vip_three_six", "new_vip_three_five", "new_vip_three_four", "new_vip_three_one", "new_vip_three_two", "new_vip_three_three")
	if nil != configs {
		for _, vConfig := range configs {
			if "new_vip_three_five" == vConfig.KeyName {
				vipThreeFive, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "new_vip_three_four" == vConfig.KeyName {
				vipThreeFour, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "new_vip_three_three" == vConfig.KeyName {
				vipThreeThree, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "new_vip_three_two" == vConfig.KeyName {
				vipThreeTwo, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "new_vip_three_one" == vConfig.KeyName {
				vipThreeOne, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "new_vip_three_seven" == vConfig.KeyName {
				vipThreeSeven, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "new_vip_three_six" == vConfig.KeyName {
				vipThreeSix, _ = strconv.ParseUint(vConfig.Value, 10, 64)
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

		tmpTopVip := uint64(5)
		totalTmp := len(tmpRecommendUserIds) - 1
		lastVip := uint64(0)
		lastAmount := uint64(0)
		for i := totalTmp; i >= 0; i-- {
			if vipThreeSeven <= lastAmount {
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
			} else if 4 == usersMap[tmpUserId].VipThree {
				if vipThreeFour <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误4：", usersMap[tmpUserId], lastVip, vipThreeFour, lastAmount)
					continue
				}

				tmpAmount = vipThreeFour - lastAmount
				lastAmount = vipThreeFour
			} else if 5 == usersMap[tmpUserId].VipThree {
				if vipThreeFive <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误5：", usersMap[tmpUserId], lastVip, vipThreeFive, lastAmount)
					continue
				}

				tmpAmount = vipThreeFive - lastAmount
				lastAmount = vipThreeFive
			} else if 6 == usersMap[tmpUserId].VipThree {
				if vipThreeSix <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误6：", usersMap[tmpUserId], lastVip, vipThreeSix, lastAmount)
					continue
				}

				tmpAmount = vipThreeSix - lastAmount
				lastAmount = vipThreeSix
			} else if 7 == usersMap[tmpUserId].VipThree {
				if vipThreeSeven <= lastAmount {
					fmt.Println("开卡2遍历，vip奖励信息设置错误7：", usersMap[tmpUserId], lastVip, vipThreeSeven, lastAmount)
					continue
				}

				tmpAmount = vipThreeSeven - lastAmount
				lastAmount = vipThreeSeven
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

func (uuc *UserUseCase) AllInfo(ctx context.Context, req *pb.AllInfoRequest) (*pb.AllInfoReply, error) {
	return &pb.AllInfoReply{
		TotalUser:          0,
		TodayUser:          0,
		TotalDeposit:       0,
		TodayDeposit:       0,
		ToAmount:           0,
		FeeAmount:          0,
		CardTotal:          0,
		CardTwoTotal:       0,
		CardRewardTotal:    0,
		CardRewardTwoTotal: 0,
		TodayWithdraw:      0,
		TotalWithdraw:      0,
		BalanceAll:         0,
	}, nil
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

func (uuc *UserUseCase) EmailGet(ctx context.Context, req *pb.EmailGetRequest) (*pb.EmailGetReply, error) {
	var (
		lastUid  uint32
		res      []NewMailParsed
		err      error
		cardCode *CardOrder
		last     uint32 = 1765876719
	)
	cardCode, err = uuc.repo.GetCardOrder()
	if err != nil {
		return nil, err
	}

	if nil != cardCode {
		last = uint32(cardCode.Last)
	}

	lastUid, res, err = FetchNewBindOtpMailsSyncV1(ctx, "", "", last, 20)
	if err == nil && res != nil {
		for _, v := range res {
			if lastUid <= last {
				continue
			}

			_ = uuc.repo.CreateCardOrder(ctx, &CardOrder{
				Last: uint64(lastUid),
				Code: v.Parsed.OTP,
				Card: v.Parsed.CardMasked,
				Time: v.Parsed.MailTime,
			})
		}
	}

	return nil, err
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
	}, req.Address, req.CardTwo, req.CardOrderId, req.LockCard, req.ChangeCard)
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
			CardNumber:         vUsers.CardNumber,
			CardTwoNumber:      vUsers.CardTwoNumber,
			CardOrderId:        vUsers.CardOrderId,
			CardTwo:            vUsers.CardTwo,
			PicTwo:             vUsers.PicTwo,
			Pic:                vUsers.Pic,
			ChangeCard:         vUsers.ChangeCard,
			ChangeCardTwo:      vUsers.ChangeCardTwo,
			LockCard:           vUsers.LockCard,
			LockCardTwo:        vUsers.LockCardTwo,
			CardNumberRel:      vUsers.CardNumberRel,
			CardNumberRelTwo:   vUsers.CardNumberRelTwo,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminUserBind(ctx context.Context, req *pb.AdminUserBindRequest) (*pb.AdminUserBindReply, error) {
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByAddress(req.SendBody.Address)
	if nil != err {
		return &pb.AdminUserBindReply{}, err
	}

	if 0 >= len(req.SendBody.CardId) {
		return &pb.AdminUserBindReply{}, err
	}

	if 0 >= len(req.SendBody.CarNum) {
		return &pb.AdminUserBindReply{}, err
	}

	if errThree := uuc.repo.UpdateUserDoing(ctx, user.ID, req.SendBody.CardId, req.SendBody.CarNum, req.SendBody.Amount); errThree != nil {
		fmt.Println("AdminUserBind", "err =", err)
		// 这条失败就算了，不影响其它
		return &pb.AdminUserBindReply{}, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminUserBindTwo(ctx context.Context, req *pb.AdminUserBindTwoRequest) (*pb.AdminUserBindTwoReply, error) {
	var (
		cardTwo *CardTwo
		err     error
	)
	cardTwo, err = uuc.repo.GetCardTwoById(req.SendBody.Id)
	if nil != err || nil == cardTwo {
		return nil, err
	}

	if 0 >= len(req.SendBody.CardId) {
		return &pb.AdminUserBindTwoReply{}, err
	}

	if 0 >= len(req.SendBody.CarNum) {
		return &pb.AdminUserBindTwoReply{}, err
	}

	if errThree := uuc.repo.UpdateCardStatus(ctx, req.SendBody.Id, cardTwo.UserId, req.SendBody.CardId, req.SendBody.CarNum, req.SendBody.Amount); errThree != nil {
		fmt.Println("AdminUserBindTwo", "err =", err)
		// 这条失败就算了，不影响其它
		return &pb.AdminUserBindTwoReply{}, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminCardTwoList(ctx context.Context, req *pb.AdminCardTwoRequest) (*pb.AdminCardTwoReply, error) {
	var (
		cards     []*CardTwo
		userIds   []uint64
		usersMap  map[uint64]*User
		count     int64
		err       error
		tmpUserId uint64
		sUser     *User
	)

	res := &pb.AdminCardTwoReply{
		Users: make([]*pb.AdminCardTwoReply_EntityCardUser, 0),
	}

	if 0 < len(req.Address) {
		sUser, err = uuc.repo.GetUserByAddress(req.Address)
		if nil != err || nil == sUser {
			return res, nil
		}

		tmpUserId = sUser.ID
	}

	cards, err, count = uuc.repo.GetCardTwos(&Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, tmpUserId, req.Status, "")
	if nil != err {
		return res, nil
	}
	res.Count = count

	if len(cards) < 0 {
		return res, nil
	}

	userIds = make([]uint64, 0)
	for _, vUsers := range cards {
		userIds = append(userIds, vUsers.UserId)
	}

	usersMap, err = uuc.repo.GetUserByUserIdsTwo(userIds)
	if nil != err {
		return res, nil
	}

	for _, vUsers := range cards {
		addressTmp := ""
		if _, ok := usersMap[vUsers.UserId]; ok {
			addressTmp = usersMap[vUsers.UserId].Address
		}

		res.Users = append(res.Users, &pb.AdminCardTwoReply_EntityCardUser{
			Id:               vUsers.ID,
			UserId:           vUsers.UserId,
			Address:          addressTmp,
			CreatedAt:        vUsers.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			FirstName:        vUsers.FirstName,
			LastName:         vUsers.LastName,
			Email:            vUsers.Email,
			CountryCode:      vUsers.CountryCode,
			Phone:            vUsers.Phone,
			City:             vUsers.City,
			Country:          vUsers.Country,
			Street:           vUsers.Street,
			PostalCode:       vUsers.PostalCode,
			State:            vUsers.State,
			BirthDate:        vUsers.BirthDate,
			PhoneCountryCode: vUsers.PhoneCountryCode,
			CardId:           vUsers.CardId,
			Status:           vUsers.Status,
			IdCard:           vUsers.IdCard,
			Gender:           vUsers.Gender,
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

			if 0.01 < v.CardAmount {
				// 划转出去
				data, errThree := InterlaceCardTransferOut(ctx, &InterlaceCardTransferOutReq{
					AccountId:           interlaceAccountId,
					CardId:              v.CardId,
					ClientTransactionId: fmt.Sprintf("out-%d", time.Now().UnixNano()),
					Amount:              fmt.Sprintf("%.2f", v.CardAmount), // 字符串
				})
				if nil == data || errThree != nil {
					fmt.Println("InterlaceCardTransferOut error:", err)
					continue
				}

				if 3 != data.Type {
					fmt.Println("out err", v, data)
					continue
				}

				if "CLOSED" != data.Status {
					fmt.Println("out status err", v, data)
				}

				if "FAIL" == data.Status {
					fmt.Println("out status fail err", v, data)
					continue
				}
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

	var (
		usersAll []*User
		usersMap map[uint64]*User
	)
	usersAll, err = uuc.repo.GetAllUsers()
	if nil == usersAll || nil != err {
		fmt.Println("用户无")
		return nil, err
	}

	usersMap = make(map[uint64]*User, 0)
	for _, vUsers := range usersAll {
		usersMap[vUsers.ID] = vUsers
	}

	// 把第一页也放到统一处理逻辑里
	for _, v := range users {
		user := v
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

			if 0.01 < v.CardAmount {
				// 划转出去
				data, errThree := InterlaceCardTransferOut(ctx, &InterlaceCardTransferOutReq{
					AccountId:           interlaceAccountId,
					CardId:              v.CardNumber,
					ClientTransactionId: fmt.Sprintf("out-%d", time.Now().UnixNano()),
					Amount:              fmt.Sprintf("%.2f", v.CardAmount), // 字符串
				})
				if nil == data || errThree != nil {
					fmt.Println("InterlaceCardTransferOut error:", err)
					continue
				}

				if 3 != data.Type {
					fmt.Println("out err", v, data)
					continue
				}

				if "CLOSED" != data.Status {
					fmt.Println("out status err", v, data)
				}

				if "FAIL" == data.Status {
					fmt.Println("out status fail err", v, data)
					continue
				}
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

			// 分红
			var (
				userRecommend *UserRecommend
			)
			tmpRecommendUserIds := make([]string, 0)
			// 推荐
			userRecommend, err = uuc.repo.GetUserRecommendByUserId(user.ID)
			if nil == userRecommend {
				fmt.Println(err, "信息错误", err, user)
				continue
			}
			if "" != userRecommend.RecommendCode {
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
			}

			tmpTopVip := uint64(15)
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
	} else if 4 == req.SendBody.VipThree {
		lock = 4
	} else if 5 == req.SendBody.VipThree {
		lock = 5
	} else if 6 == req.SendBody.VipThree {
		lock = 6
	} else if 7 == req.SendBody.VipThree {
		lock = 7
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

	accessToken, refreshToken, expiresIn, t, err := interlaceGenerateAccessToken(ctx, code)
	if 0 >= len(accessToken) || err != nil {
		return "", fmt.Errorf("generate interlace access token failed: %w", err)
	}

	interlaceAuth.AccessToken = accessToken
	interlaceAuth.RefreshToken = refreshToken
	interlaceAuth.ExpireAt = t + expiresIn

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

func interlaceGenerateAccessToken(ctx context.Context, code string) (accessToken, refreshToken string, expiresIn, t int64, err error) {
	urlStr := fmt.Sprintf("%s/oauth/access-token", interlaceBaseURL)

	reqBody := map[string]interface{}{
		"clientId": interlaceClientID,
		"code":     code,
	}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, bytes.NewReader(jsonData))
	if err != nil {
		return "", "", 0, 0, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", 0, 0, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", 0, 0, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", "", 0, 0, fmt.Errorf("interlace access-token http %d: %s", resp.StatusCode, string(body))
	}

	var result interlaceAccessTokenResp
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", 0, 0, fmt.Errorf("interlace access-token unmarshal: %w", err)
	}

	if result.Code != "000000" {
		return "", "", 0, 0, fmt.Errorf("interlace access-token failed: code=%s msg=%s", result.Code, result.Message)
	}
	if result.Data.AccessToken == "" {
		return "", "", 0, 0, fmt.Errorf("interlace access-token success but accessToken empty")
	}

	return result.Data.AccessToken, result.Data.RefreshToken, result.Data.ExpiresIn, result.Data.Timestamp, nil
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
		fmt.Println("at", interlaceAuth, time.Now().Unix())
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

// InterlaceCardTransferOutReq 划转请求
type InterlaceCardTransferOutReq struct {
	AccountId           string `json:"accountId"`           // 账户 UUID
	CardId              string `json:"cardId"`              // 卡 UUID
	ClientTransactionId string `json:"clientTransactionId"` // 自定义交易 ID
	Amount              string `json:"amount"`              // 划转金额（字符串）
}

// 手续费明细
type InterlaceFeeDetail struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	FeeType  string `json:"feeType"` // 0:Platform Settlement Fee, 1:Apple Pay Fee, ...
}

// 划转结果 data 部分
type InterlaceCardTransferOutData struct {
	ID                       string               `json:"id"`
	AccountId                string               `json:"accountId"`
	CardId                   string               `json:"cardId"`
	CardholderId             string               `json:"cardholderId"`
	CardTransactionId        string               `json:"cardTransactionId"`
	Currency                 string               `json:"currency"`
	Amount                   string               `json:"amount"`
	Fee                      string               `json:"fee"`
	FeeDetails               []InterlaceFeeDetail `json:"feeDetails"`
	ClientTransactionId      string               `json:"clientTransactionId"`
	RelatedCardTransactionId string               `json:"relatedCardTransactionId"`
	TransactionDisplayId     string               `json:"transactionDisplayId"`
	Type                     int32                `json:"type"`   // 0:Credit,1:Consumption,2:TransferIn,3:TransferOut...
	Status                   string               `json:"status"` // CLOSED,PENDING,FAIL

	MerchantName    string `json:"merchantName"`
	Mcc             string `json:"mcc"`
	MccCategory     string `json:"mccCategory"`
	MerchantCity    string `json:"merchantCity"`
	MerchantCountry string `json:"merchantCountry"`
	MerchantState   string `json:"merchantState"`
	MerchantZipcode string `json:"merchantZipcode"`
	MerchantMid     string `json:"merchantMid"`

	TransactionTime     string `json:"transactionTime"`
	TransactionCurrency string `json:"transactionCurrency"`
	TransactionAmount   string `json:"transactionAmount"`
	CreateTime          string `json:"createTime"`
	Remark              string `json:"remark"`
	Detail              string `json:"detail"`
}

// 外层响应
type InterlaceCardTransferOutResp struct {
	Code    string                       `json:"code"`
	Message string                       `json:"message"`
	Data    InterlaceCardTransferOutData `json:"data"`
}

// InterlaceCardTransferOut 预付卡划转出到 Quantum 账户
func InterlaceCardTransferOut(ctx context.Context, in *InterlaceCardTransferOutReq) (*InterlaceCardTransferOutData, error) {
	if in == nil {
		return nil, fmt.Errorf("transfer out req is nil")
	}
	if in.AccountId == "" {
		return nil, fmt.Errorf("accountId is required")
	}
	if in.CardId == "" {
		return nil, fmt.Errorf("cardId is required")
	}
	if in.ClientTransactionId == "" {
		return nil, fmt.Errorf("clientTransactionId is required")
	}
	if in.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}

	accessToken, err := GetInterlaceAccessToken(ctx)
	if err != nil || accessToken == "" {
		fmt.Println("获取access token错误")
		return nil, err
	}

	// baseURL 建议为: https://api-sandbox.interlace.money/open-api/v3
	base := interlaceBaseURL + "/cards/transfer-out"

	bodyBytes, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("marshal transfer out body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println("transfer-out resp:", string(respBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("interlace transfer out http %d: %s", resp.StatusCode, string(respBody))
	}

	var outer InterlaceCardTransferOutResp
	if err := json.Unmarshal(respBody, &outer); err != nil {
		return nil, fmt.Errorf("transfer out unmarshal: %w", err)
	}
	if outer.Code != "000000" {
		return nil, fmt.Errorf("transfer out failed: code=%s msg=%s", outer.Code, outer.Message)
	}

	return &outer.Data, nil
}

/*************** 解析：卡号 + OTP + TTL + 时间 ***************/

type BindOtpMail struct {
	CardMasked string     // 例如 49387519xxxxxx0945（保留 x）
	CardDigits string     // 例如 493875190000000945（如果邮件给的是纯数字）
	CardLast4  string     // 例如 0945
	OTP        string     // 6位验证码
	TTLMinutes int        // 有效期分钟数（解析不到为 0）
	MailTime   *time.Time // 邮件正文里出现的时间（解析不到为 nil）
}

var (
	// OTP：优先找“验证码/OTP/verification code”附近的6位数字
	reOTPNear = regexp.MustCompile(`(?i)(验证码|otp|verification\s*code)[^0-9]{0,40}(\d{6})`)
	// OTP：兜底找任意独立6位数字
	reOTPAny = regexp.MustCompile(`\b(\d{6})\b`)

	reCardMasked = regexp.MustCompile(`(?i)(?:your\s+card|card|您的?卡|卡号|卡)\s*[:：]?\s*([0-9]{6,12}[xX\*]{2,12}[0-9]{2,6})`)
	reCardDigits = regexp.MustCompile(`(?i)(?:your\s+card|card|您的?卡|卡号|卡)\s*[:：]?\s*([0-9]{12,19})`)

	// 后四位：用来从卡号里提取最后4位
	reLast4 = regexp.MustCompile(`([0-9]{4})\b`)

	// 有效期：有效期:7分钟 / 有效期 7 分钟
	reTTLMin = regexp.MustCompile(`有效期[^0-9]{0,10}(\d{1,3})\s*分钟`)

	// 邮件正文时间：2025-12-11 15:14
	reTime = regexp.MustCompile(`\b(20\d{2}-\d{2}-\d{2}\s+\d{2}:\d{2})\b`)

	// --- 轻量 HTML 清洗（Go regexp 不支持 \1 反向引用，所以拆开） ---
	reStripScript = regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	reStripStyle  = regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	reBr          = regexp.MustCompile(`(?is)<br\s*/?>`)
	rePclose      = regexp.MustCompile(`(?is)</p>`)
	reTag         = regexp.MustCompile(`(?is)<[^>]+>`)
)

// ParseBindOtpMail 从 subject+body 中解析：卡号(掩码/纯数字)、后四位、OTP、TTL分钟、正文时间
func ParseBindOtpMail(subject, body string) BindOtpMail {
	raw := compactSpaces(subject + "\n" + body)
	out := BindOtpMail{}

	// 1) 卡号
	if m := reCardMasked.FindStringSubmatch(raw); len(m) >= 2 {
		out.CardMasked = normalizeMask(m[1])
		out.CardLast4 = extractLast4(out.CardMasked)
	} else if m := reCardDigits.FindStringSubmatch(raw); len(m) >= 2 {
		out.CardDigits = m[1]
		out.CardLast4 = extractLast4(out.CardDigits)
	}

	// 2) OTP
	if m := reOTPNear.FindStringSubmatch(raw); len(m) >= 3 {
		out.OTP = m[2]
	} else {
		out.OTP = findAnyOTPExcludingCard(raw, out.CardMasked, out.CardDigits)
	}

	// 3) 有效期分钟
	if m := reTTLMin.FindStringSubmatch(raw); len(m) >= 2 {
		out.TTLMinutes = atoiSafe(m[1])
	}

	// 4) 邮件正文时间
	if tm := parseMailTime(raw); tm != nil {
		out.MailTime = tm
	}

	return out
}

// findAnyOTPExcludingCard 从全文找6位数字，并尽量排除卡号片段/后四位相关误判
func findAnyOTPExcludingCard(raw, cardMasked, cardDigits string) string {
	matches := reOTPAny.FindAllStringSubmatch(raw, -1)

	last4 := extractLast4(cardMasked)
	if last4 == "" {
		last4 = extractLast4(cardDigits)
	}

	for _, mm := range matches {
		if len(mm) < 2 {
			continue
		}
		code := mm[1]

		// 如果 code 是卡号的一部分（纯数字卡号里包含这6位），跳过
		if cardDigits != "" && strings.Contains(cardDigits, code) {
			continue
		}

		// 避免把“含后四位的6位”误判为 OTP（例如 XX0945）
		if last4 != "" && len(code) == 6 && code[2:] == last4 {
			continue
		}

		return code
	}
	return ""
}

func parseMailTime(raw string) *time.Time {
	m := reTime.FindStringSubmatch(raw)
	if len(m) < 2 {
		return nil
	}
	tm, err := time.ParseInLocation("2006-01-02 15:04", m[1], time.Local)
	if err != nil {
		return nil
	}
	return &tm
}

// extractLast4 取“最后一个”4位数字（避免拿到开头那段 4 位）
func extractLast4(s string) string {
	if s == "" {
		return ""
	}
	all := reLast4.FindAllStringSubmatch(s, -1)
	if n := len(all); n > 0 && len(all[n-1]) >= 2 {
		return all[n-1][1]
	}
	d := onlyDigits(s)
	if len(d) >= 4 {
		return d[len(d)-4:]
	}
	return ""
}

func compactSpaces(s string) string {
	// HTML实体解码（&nbsp; 等）
	s = html.UnescapeString(s)

	// 轻量去 script/style
	s = reStripScript.ReplaceAllString(s, " ")
	s = reStripStyle.ReplaceAllString(s, " ")

	// 处理换行
	s = reBr.ReplaceAllString(s, "\n")
	s = rePclose.ReplaceAllString(s, "\n")

	// 去掉所有标签
	s = reTag.ReplaceAllString(s, " ")

	// 统一空白
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return strings.TrimSpace(s)
}

// normalizeMask 把 X/* 统一成 x，并去空格
func normalizeMask(card string) string {
	c := strings.TrimSpace(card)
	c = strings.ReplaceAll(c, " ", "")
	c = strings.Map(func(r rune) rune {
		if r == 'X' || r == '*' {
			return 'x'
		}
		return r
	}, c)
	return c
}

func onlyDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func atoiSafe(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			break
		}
		n = n*10 + int(r-'0')
	}
	return n
}

/*************** 同步增量拉取（无 goroutine） ***************/
type NewMailParsed struct {
	UID          uint32
	Subject      string
	From         string
	InternalDate time.Time
	Parsed       BindOtpMail
}

func FetchNewBindOtpMailsSyncV1(ctx context.Context, email, authCode string, lastUID uint32, limit int) (uint32, []NewMailParsed, error) {
	const (
		addr       = "imap.163.com:993"
		serverName = "imap.163.com"
		mailbox    = "INBOX"

		targetFrom = "noreply@email.interlace.money"

		totalTimeout = 40 * time.Second

		dialTimeout      = 12 * time.Second
		handshakeTimeout = 12 * time.Second

		maxRawBytes = int64(5 << 20) // 5MB
	)

	if limit <= 0 {
		limit = 50
	}

	// ✅ 不用外层 request ctx（它可能很短导致 dial deadline exceeded）
	runCtx, cancel := context.WithTimeout(context.Background(), totalTimeout)
	defer cancel()

	// 1) Dial + TLS Handshake（强制 tcp4，避免 IPv6 卡住）
	c, err := dialIMAPTLS(runCtx, "tcp4", addr, serverName, dialTimeout, handshakeTimeout)
	if err != nil {
		return lastUID, nil, fmt.Errorf("dial tls: %w", err)
	}

	// ✅ 只允许关闭一次，避免重复 close 导致 noisy 的 “use of closed network connection”
	var once sync.Once
	closeOnce := func() {
		once.Do(func() {
			forceCloseIMAP(c)
		})
	}
	defer closeOnce()

	// 2) Login
	if err := runWithCtx(runCtx, closeOnce, func() error {
		return c.Login(email, authCode)
	}); err != nil {
		return lastUID, nil, fmt.Errorf("login: %w", err)
	}

	// 3) 163 常见需要 IMAP ID（忽略错误）
	_, _ = imapid.NewClient(c).ID(imapid.ID{
		"name":    "kratos-card-service",
		"version": "1.0.0",
		"vendor":  "your-company",
	})

	// 4) Select INBOX（只读）
	if err := runWithCtx(runCtx, closeOnce, func() error {
		_, e := c.Select(mailbox, true)
		return e
	}); err != nil {
		return lastUID, nil, fmt.Errorf("select %s: %w", mailbox, err)
	}

	// 5) UID 增量搜索：UID (lastUID+1):*
	criteria := imap.NewSearchCriteria()
	criteria.Uid = new(imap.SeqSet)
	criteria.Uid.AddRange(lastUID+1, 0)

	var uidsAll []uint32
	if err := runWithCtx(runCtx, closeOnce, func() error {
		var e error
		uidsAll, e = c.UidSearch(criteria)
		return e
	}); err != nil {
		return lastUID, nil, fmt.Errorf("uid search: %w", err)
	}
	if len(uidsAll) == 0 {
		return lastUID, nil, nil
	}

	sort.Slice(uidsAll, func(i, j int) bool { return uidsAll[i] < uidsAll[j] })

	// ✅ 关键：本轮最大 UID（不受 fetch 成功与否影响）
	maxUID := uidsAll[len(uidsAll)-1]

	// ✅ fetch 用窗口（你只追最新，漏本轮无所谓）
	scanLimit := limit * 10
	if scanLimit < 200 {
		scanLimit = 200
	}
	if scanLimit > 1000 {
		scanLimit = 1000
	}
	uidsFetch := uidsAll
	if len(uidsFetch) > scanLimit {
		uidsFetch = uidsFetch[len(uidsFetch)-scanLimit:]
	}

	uidset := new(imap.SeqSet)
	uidset.AddNum(uidsFetch...)

	// 用 BODY[] 拿到整封 RFC822，再自己解 MIME
	section := &imap.BodySectionName{Peek: true}
	fetchItems := []imap.FetchItem{
		imap.FetchUid,
		imap.FetchEnvelope,
		imap.FetchInternalDate,
		section.FetchItem(),
	}

	// 6) Fetch：done + select，保证不死等
	msgCh := make(chan *imap.Message, 10)
	done := make(chan error, 1)

	go func() {
		done <- c.UidFetch(uidset, fetchItems, msgCh) // 结束会 close(msgCh)
	}()

	out := make([]NewMailParsed, 0, limit)
	msgChOpen := true

	for msgChOpen {
		select {
		case msg, ok := <-msgCh:
			if !ok {
				msgChOpen = false
				continue
			}
			if msg == nil {
				continue
			}

			subject := ""
			from := ""
			if msg.Envelope != nil {
				subject = msg.Envelope.Subject

				for _, a := range msg.Envelope.From {
					if a == nil {
						continue
					}
					mb := a.MailboxName
					hs := a.HostName
					if mb == "" || hs == "" {
						continue
					}
					addr2 := mb + "@" + hs
					if from == "" {
						from = addr2
					}
					if strings.EqualFold(addr2, targetFrom) {
						from = addr2
						break
					}
				}
			}

			// ✅ 只记录目标发件人
			if !strings.EqualFold(from, targetFrom) {
				continue
			}

			bodyText := ""
			if r := msg.GetBody(section); r != nil {
				rawRFC822, _ := io.ReadAll(io.LimitReader(r, maxRawBytes))
				bodyText = ExtractDecodedMailBody(rawRFC822)
				if bodyText == "" {
					bodyText = string(rawRFC822)
				}
			}

			parsed := ParseBindOtpMail(subject, bodyText)

			// 你自己调试打印
			// fmt.Println(parsed)

			if len(out) < limit {
				out = append(out, NewMailParsed{
					UID:          msg.Uid,
					Subject:      subject,
					From:         from,
					InternalDate: msg.InternalDate,
					Parsed:       parsed,
				})
			}

		case err := <-done:
			if err != nil {
				// ✅ 即使 fetch 失败，也返回 maxUID（本轮最大）
				return maxUID, out, fmt.Errorf("uid fetch: %w", err)
			}
			msgChOpen = false

		case <-runCtx.Done():
			// ✅ 超时：断开连接打断阻塞；仍返回 maxUID
			closeOnce()
			return maxUID, out, runCtx.Err()
		}
	}

	return maxUID, out, nil
}

// ExtractDecodedMailBody：从 RFC822 原始邮件中提取“已解码”的正文
// 优先 text/plain；没有则用 text/html
func ExtractDecodedMailBody(rawRFC822 []byte) string {
	mr, err := mail.CreateReader(bytes.NewReader(rawRFC822))
	if err != nil {
		return ""
	}

	var plain, htmlBody string

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		ih, ok := p.Header.(*mail.InlineHeader)
		if !ok {
			continue
		}

		ct, _, _ := ih.ContentType()
		ct = strings.ToLower(ct)

		b, _ := io.ReadAll(p.Body)

		if strings.HasPrefix(ct, "text/plain") && plain == "" {
			plain = string(b)
		}
		if strings.HasPrefix(ct, "text/html") && htmlBody == "" {
			htmlBody = string(b)
		}
	}

	if plain != "" {
		return plain
	}
	if htmlBody != "" {
		return htmlBody
	}
	return ""
}

// ---- helpers ----

func dialIMAPTLS(ctx context.Context, network, addr, serverName string, dialTimeout, handshakeTimeout time.Duration) (*client.Client, error) {
	d := &net.Dialer{Timeout: dialTimeout}
	rawConn, err := d.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	tlsConn := tls.Client(rawConn, &tls.Config{
		ServerName: serverName,
		MinVersion: tls.VersionTLS12,
	})

	_ = tlsConn.SetDeadline(time.Now().Add(handshakeTimeout))
	if err := tlsConn.Handshake(); err != nil {
		_ = tlsConn.Close()
		return nil, err
	}
	_ = tlsConn.SetDeadline(time.Time{})

	return client.New(tlsConn)
}

// runWithCtx：让 IMAP 命令在 ctx 超时/取消时可被打断（只关闭一次）
func runWithCtx(ctx context.Context, closeFn func(), fn func() error) error {
	done := make(chan error, 1)
	go func() { done <- fn() }()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		closeFn()
		return ctx.Err()
	}
}

func forceCloseIMAP(c *client.Client) {
	if c == nil {
		return
	}
	// 有些版本有 Terminate()，没有就 Logout()
	v := reflect.ValueOf(c)
	m := v.MethodByName("Terminate")
	if m.IsValid() && m.Type().NumIn() == 0 {
		m.Call(nil)
		return
	}
	_ = c.Logout()
}
