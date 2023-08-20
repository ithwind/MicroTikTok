package logic

import (
	"MicroTikTok/common/cryptx"
	"MicroTikTok/user/model/dao/model"
	"MicroTikTok/user/rpc/internal/svc"
	"MicroTikTok/user/rpc/rpc/user"
	"context"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
//默认黑色背景

func GenerateRandomBackgroundImage() ([]byte, error) {

	const width, height = 100, 100

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 设置黑色背景
	drawColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{drawColor}, image.ZP, draw.Src)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			color := randomColor()
			img.Set(x, y, color)
		}
	}

	var buf bytes.Buffer
	err := png.Encode(&buf, img)

	return buf.Bytes(), err

}

// 生成随机颜色
func randomColor() color.Color {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

// 随机生成头像
func GenerateRandomAvatar() ([]byte, error) {

	// 设置图像尺寸
	const width, height = 100, 100

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 填充随机颜色
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, randomColor())
		}
	}

	//编码成PNG文件并返回
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	return buf.Bytes(), err
}
*/
//四位随机数

func GenerateRandomCode() string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	var str string
	for i := 0; i < 4; i++ {
		num := r.Intn(10)
		str += strconv.Itoa(num)
	}

	return str
}

// 检验邮箱格式
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w-]+@([\w-]+\.)+[\w-]{2,4}`)
	return emailRegex.MatchString(email)
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	name := GenerateRandomCode()
	/*avatarBytes, _ := GenerateRandomAvatar()
	backgroundimagebytes, _ := GenerateRandomBackgroundImage()
	backgroundimage := base64.StdEncoding.EncodeToString(backgroundimagebytes)
	avatar := base64.StdEncoding.EncodeToString(avatarBytes)*/
	table := l.svcCtx.BkModel.User
	_, err := table.WithContext(l.ctx).Where(table.Username.Eq(in.Username)).Debug().First()
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}
	if err == gorm.ErrRecordNotFound {
		newUser := model.User{
			Username:        in.Username,
			Password:        cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			Name:            name,
			Signature:       "该用户什么都没留下",
			Avatar:          "avatar",
			BackgroundImage: "backgroundimage",
		}
		if !isValidEmail(in.Username) {
			return nil, status.Error(500, err.Error())
		}
		_, err := table.WithContext(l.ctx).Create(&newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		result, err := table.WithContext(l.ctx).Last(&newUser)
		newUser.ID = result.ID
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		statusMsg := "注册成功"
		return &user.RegisterResponse{
			UserId:     int64(newUser.ID),
			StatusCode: 0,
			StatusMsg:  &statusMsg,
		}, nil
	}
	return nil, status.Error(500, err.Error())
}
