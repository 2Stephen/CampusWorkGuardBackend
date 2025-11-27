package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/big"
	"time"
)

func SendCode(params dto.SendCodeRequest) error {
	switch params.Role {
	case "login":
		return sendLoginCode(params.Email)
	case "register":
		return sendRegisterCode(params.Email)
	default:
		return errors.New("非法角色类型")
	}
}

func generateCode() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n.Int64())
}

func sendLoginCode(email string) error {
	// 查询库中是否含有该用户
	if !repository.IsUserExistByEmail(email) {
		return errors.New("用户不存在")
	}
	// 发送登录验证码逻辑
	// 先检查是否有未过期的验证码
	_, err := utils.RedisGet("login_code:" + email)
	if !errors.Is(err, redis.Nil) {
		// 限制同一邮箱一分钟之内只能发一次
		duration, err := utils.RedisTTL("register_code:" + email)
		if err != nil {
			log.Println(err)
			return err
		}
		if duration > 4*time.Minute {
			return errors.New("验证码发送频繁，请稍后重试")
		}
		// 删除旧key
		err = utils.RedisDel("login_code:" + email)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	code := generateCode()
	err = utils.RedisSet("login_code:"+email, code, 5*time.Minute)
	if err != nil {
		log.Println(err)
		return err
	}
	// 发送邮件
	err = utils.SendEmailCode(email, code)
	if err != nil {
		log.Printf("验证码发送失败: %+v", err)
		return err
	}
	return nil
}

func sendRegisterCode(email string) error {
	// 查询库中是否含有该用户
	if repository.IsUserExistByEmail(email) {
		return errors.New("用户已存在")
	}
	// 发送注册验证码逻辑
	// 发送登录验证码逻辑
	// 先检查是否有未过期的验证码
	_, err := utils.RedisGet("register_code:" + email)
	if !errors.Is(err, redis.Nil) {
		// 限制同一邮箱一分钟之内只能发一次
		duration, err := utils.RedisTTL("register_code:" + email)
		if err != nil {
			log.Println(err)
			return err
		}
		if duration > 4*time.Minute {
			return errors.New("验证码发送频繁，请稍后重试")
		}
		// 删除旧key
		err = utils.RedisDel("register_code:" + email)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	code := generateCode()
	err = utils.RedisSet("register_code:"+email, code, 5*time.Minute)
	if err != nil {
		log.Println(err)
		return err
	}
	// 发送邮件
	err = utils.SendEmailCode(email, code)
	if err != nil {
		log.Printf("验证码发送失败: %+v", err)
		return err
	}
	return nil
}
