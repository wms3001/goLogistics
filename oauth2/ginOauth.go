package oauth2

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-oauth2/oauth2/v4"
	//"github.com/go-oauth2/oauth2/v4/errors"
	//"github.com/go-oauth2/oauth2/v4/manage"
	//"github.com/go-oauth2/oauth2/v4/models"
	//"github.com/go-oauth2/oauth2/v4/server"
	//"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gopkg.in/go-oauth2/mongo.v3"
	oredis "gopkg.in/go-oauth2/redis.v3"

	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"logistics/model"
	"os"
	"time"
)

var (
	gServer             *server.Server
	gClient             *store.ClientStore
	gManage             *manage.Manager
	baseResponse        *model.Base
	credentialsResponse *model.Credentials
)

/**
统一token返回格式
*/
func SetExtensionFields(ti oauth2.TokenInfo) map[string]interface{} {
	data := map[string]interface{}{
		"code":    1,
		"message": "success",
	}
	return data
}

func init() {
	work, _ := os.Getwd()
	viper.SetConfigName("oauth")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/conf")
	viper.ReadInConfig()
	storeType := viper.GetString("oauth.tokenStore")
	accessTokenExp := time.Duration(viper.GetInt("oauth.accessTokenExp"))
	refreshTokenExp := time.Duration(viper.GetInt("oauth.refreshTokenExp"))
	var isGenerateRefresh bool
	if viper.GetInt("oauth.isGenerateRefresh") == 1 {
		isGenerateRefresh = true
	} else {
		isGenerateRefresh = false
	}
	gManage = manage.NewDefaultManager()
	switch storeType {
	case "mongo":
		viper.SetConfigName("mongo")
		viper.SetConfigType("yml")
		viper.AddConfigPath(work + "/conf")
		viper.ReadInConfig()
		mHost := viper.GetString("mongo.host")
		mPort := viper.GetString("mongo.port")
		mUser := viper.GetString("mongo.user")
		mPass := viper.GetString("mongo.pass")
		mDb := viper.GetString("mongo.dbName")
		gManage.MapTokenStorage(
			mongo.NewTokenStore(
				mongo.NewConfig("mongodb://"+mUser+":"+mPass+"@"+mHost+":"+mPort+"/"+mDb, mDb),
				mongo.NewDefaultTokenConfig(),
			),
		)
	case "memory":
		gManage.MustTokenStorage(store.NewMemoryTokenStore())
	case "redis":
		viper.SetConfigName("redis")
		viper.SetConfigType("yml")
		viper.AddConfigPath(work + "/conf")
		viper.ReadInConfig()
		rHost := viper.GetString("redis.host")
		rPort := viper.GetString("redis.port")
		rDb := viper.GetInt("redis.db")
		rPass := viper.GetString("redis.pass")
		gManage.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
			Addr:     rHost + ":" + rPort,
			DB:       rDb,
			Password: rPass,
		}, ""))
	case "file":
		gManage.MustTokenStorage(store.NewFileTokenStore("token.db"))
	}

	gClient = store.NewClientStore()
	gManage.MapClientStorage(gClient)
	gServer = server.NewDefaultServer(gManage)
	gServer.SetAllowGetAccessRequest(true)
	gServer.SetClientInfoHandler(server.ClientFormHandler)
	var cfg = &manage.Config{
		AccessTokenExp:    time.Hour * accessTokenExp,
		RefreshTokenExp:   time.Hour * 24 * refreshTokenExp,
		IsGenerateRefresh: isGenerateRefresh,
	}
	gManage.SetAuthorizeCodeTokenCfg(cfg)
	gManage.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)
	gManage.SetClientTokenCfg(cfg)
	gServer.SetExtensionFieldsHandler(SetExtensionFields)
	gServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	gServer.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
}

func TokenRequest(c *gin.Context) {
	gServer.HandleTokenRequest(c.Writer, c.Request)
}

func Credentials(c *gin.Context) {
	clientId := uuid.New().String()[:16]
	clientSecret := uuid.New().String()[:16]
	err := gClient.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost:9094",
	})
	if err != nil {
		baseResponse = &model.Base{}
		baseResponse.Code = 1000
		baseResponse.Message = err.Error()
		c.JSON(500, baseResponse)
		c.Abort()
	}
	credentialsResponse = &model.Credentials{}
	credentialsResponse.Code = 1
	credentialsResponse.Message = "success"
	credentialsResponse.ClientId = clientId
	credentialsResponse.ClientSecret = clientSecret
	c.JSON(200, credentialsResponse)
}

/**
权限验证中间件
*/
func AuthValidate(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := gServer.ValidationBearerToken(c.Request)
		if err != nil {
			baseResponse = &model.Base{}
			baseResponse.Code = 1001
			baseResponse.Message = err.Error()
			c.JSON(401, baseResponse)
			c.Abort()
			return
		} else {
			c.Next()
		}

	}
}
