package parse

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cast"
)

type ServerConfig struct {
	Difficulty                           string  `json:"Difficulty"`
	DayTimeSpeedRate                     float64 `json:"DayTimeSpeedRate"`
	NightTimeSpeedRate                   float64 `json:"NightTimeSpeedRate"`
	ExpRate                              float64 `json:"ExpRate"`
	PalCaptureRate                       float64 `json:"PalCaptureRate"`
	PalSpawnNumRate                      float64 `json:"PalSpawnNumRate"`
	PalDamageRateAttack                  float64 `json:"PalDamageRateAttack"`
	PalDamageRateDefense                 float64 `json:"PalDamageRateDefense"`
	PlayerDamageRateAttack               float64 `json:"PlayerDamageRateAttack"`
	PlayerDamageRateDefense              float64 `json:"PlayerDamageRateDefense"`
	PlayerStomachDecreaseRate            float64 `json:"PlayerStomachDecreaseRate"`
	PlayerStaminaDecreaseRate            float64 `json:"PlayerStaminaDecreaseRate"`
	PlayerAutoHPRegeneRate               float64 `json:"PlayerAutoHPRegeneRate"`
	PlayerAutoHpRegeneRateInSleep        float64 `json:"PlayerAutoHpRegeneRateInSleep"`
	PalStomachDecreaseRate               float64 `json:"PalStomachDecreaseRate"`
	PalStaminaDecreaseRate               float64 `json:"PalStaminaDecreaseRate"`
	PalAutoHPRegeneRate                  float64 `json:"PalAutoHPRegeneRate"`
	PalAutoHpRegeneRateInSleep           float64 `json:"PalAutoHpRegeneRateInSleep"`
	BuildObjectDamageRate                float64 `json:"BuildObjectDamageRate"`
	BuildObjectDeteriorationDamageRate   float64 `json:"BuildObjectDeteriorationDamageRate"`
	CollectionDropRate                   float64 `json:"CollectionDropRate"`
	CollectionObjectHpRate               float64 `json:"CollectionObjectHpRate"`
	CollectionObjectRespawnSpeedRate     float64 `json:"CollectionObjectRespawnSpeedRate"`
	EnemyDropItemRate                    float64 `json:"EnemyDropItemRate"`
	DeathPenalty                         string  `json:"DeathPenalty"`
	BEnablePlayerToPlayerDamage          bool    `json:"bEnablePlayerToPlayerDamage"`
	BEnableFriendlyFire                  bool    `json:"bEnableFriendlyFire"`
	BEnableInvaderEnemy                  bool    `json:"bEnableInvaderEnemy"`
	BActiveUNKO                          bool    `json:"bActiveUNKO"`
	BEnableAimAssistPad                  bool    `json:"bEnableAimAssistPad"`
	BEnableAimAssistKeyboard             bool    `json:"bEnableAimAssistKeyboard"`
	DropItemMaxNum                       int     `json:"DropItemMaxNum"`
	DropItemMaxNumUNKO                   int     `json:"DropItemMaxNum_UNKO"`
	BaseCampMaxNum                       int     `json:"BaseCampMaxNum"`
	BaseCampWorkerMaxNum                 int     `json:"BaseCampWorkerMaxNum"`
	DropItemAliveMaxHours                float64 `json:"DropItemAliveMaxHours"`
	BAutoResetGuildNoOnlinePlayers       bool    `json:"bAutoResetGuildNoOnlinePlayers"`
	AutoResetGuildTimeNoOnlinePlayers    float64 `json:"AutoResetGuildTimeNoOnlinePlayers"`
	GuildPlayerMaxNum                    int     `json:"GuildPlayerMaxNum"`
	PalEggDefaultHatchingTime            float64 `json:"PalEggDefaultHatchingTime"`
	WorkSpeedRate                        float64 `json:"WorkSpeedRate"`
	BIsMultiplay                         bool    `json:"bIsMultiplay"`
	BIsPvP                               bool    `json:"bIsPvP"`
	BCanPickupOtherGuildDeathPenaltyDrop bool    `json:"bCanPickupOtherGuildDeathPenaltyDrop"`
	BEnableNonLoginPenalty               bool    `json:"bEnableNonLoginPenalty"`
	BEnableFastTravel                    bool    `json:"bEnableFastTravel"`
	BIsStartLocationSelectByMap          bool    `json:"bIsStartLocationSelectByMap"`
	BExistPlayerAfterLogout              bool    `json:"bExistPlayerAfterLogout"`
	BEnableDefenseOtherGuildPlayer       bool    `json:"bEnableDefenseOtherGuildPlayer"`
	CoopPlayerMaxNum                     int     `json:"CoopPlayerMaxNum"`
	ServerPlayerMaxNum                   int     `json:"ServerPlayerMaxNum"`
	ServerName                           string  `json:"ServerName"`
	ServerDescription                    string  `json:"ServerDescription"`
	AdminPassword                        string  `json:"AdminPassword"`
	ServerPassword                       string  `json:"ServerPassword"`
	PublicPort                           int     `json:"PublicPort"`
	PublicIP                             string  `json:"PublicIP"`
	RCONEnabled                          bool    `json:"RCONEnabled"`
	RCONPort                             int     `json:"RCONPort"`
	Region                               string  `json:"Region"`
	BUseAuth                             bool    `json:"bUseAuth"`
	BanListURL                           string  `json:"BanListURL"`
}

func String2Struct(input string) ServerConfig {
	var sc ServerConfig
	m, err := ParseConfigString(input)
	if err != nil {
		return sc
	}
	FillStructFromMap(m, &sc)
	return sc
}

// ParseConfigString parses the input string and returns a map of key-value pairs.
func ParseConfigString(input string) (map[string]string, error) {
	// 使用正则表达式提取键值对
	re := regexp.MustCompile(`([A-Za-z][A-Za-z0-9_]*)=([^,]+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	// 创建map存储键值对
	configMap := make(map[string]string)

	// 将键值对存储到map中
	for _, match := range matches {
		key := match[1]
		value := match[2]
		configMap[key] = value
	}

	return configMap, nil
}

// FillStructFromMap fills the fields of the given struct based on the provided map.
func FillStructFromMap(configMap map[string]string, config *ServerConfig) {
	// 遍历map，根据键值对填充结构体字段
	for key, value := range configMap {
		switch key {
		// float64 类型的字段
		case "Difficulty":
			config.Difficulty = value
		case "DayTimeSpeedRate":
			config.DayTimeSpeedRate = cast.ToFloat64(value)
		case "NightTimeSpeedRate":
			config.NightTimeSpeedRate = cast.ToFloat64(value)
		case "ExpRate":
			config.ExpRate = cast.ToFloat64(value)
		case "PalCaptureRate":
			config.PalCaptureRate = cast.ToFloat64(value)
		case "PalSpawnNumRate":
			config.PalSpawnNumRate = cast.ToFloat64(value)
		case "PalDamageRateAttack":
			config.PalDamageRateAttack = cast.ToFloat64(value)
		case "PalDamageRateDefense":
			config.PalDamageRateDefense = cast.ToFloat64(value)
		case "PlayerDamageRateAttack":
			config.PlayerDamageRateAttack = cast.ToFloat64(value)
		case "PlayerDamageRateDefense":
			config.PlayerDamageRateDefense = cast.ToFloat64(value)
		case "PlayerStomachDecreaseRate":
			config.PlayerStomachDecreaseRate = cast.ToFloat64(value)
		case "PlayerStaminaDecreaseRate":
			config.PlayerStaminaDecreaseRate = cast.ToFloat64(value)
		case "PlayerAutoHPRegeneRate":
			config.PlayerAutoHPRegeneRate = cast.ToFloat64(value)
		case "PlayerAutoHpRegeneRateInSleep":
			config.PlayerAutoHpRegeneRateInSleep = cast.ToFloat64(value)
		case "PalStomachDecreaseRate":
			config.PalStomachDecreaseRate = cast.ToFloat64(value)
		case "PalStaminaDecreaseRate":
			config.PalStaminaDecreaseRate = cast.ToFloat64(value)
		case "PalAutoHPRegeneRate":
			config.PalAutoHPRegeneRate = cast.ToFloat64(value)
		case "PalAutoHpRegeneRateInSleep":
			config.PalAutoHpRegeneRateInSleep = cast.ToFloat64(value)
		case "BuildObjectDamageRate":
			config.BuildObjectDamageRate = cast.ToFloat64(value)
		case "BuildObjectDeteriorationDamageRate":
			config.BuildObjectDeteriorationDamageRate = cast.ToFloat64(value)
		case "CollectionDropRate":
			config.CollectionDropRate = cast.ToFloat64(value)
		case "CollectionObjectHpRate":
			config.CollectionObjectHpRate = cast.ToFloat64(value)
		case "CollectionObjectRespawnSpeedRate":
			config.CollectionObjectRespawnSpeedRate = cast.ToFloat64(value)
		case "EnemyDropItemRate":
			config.EnemyDropItemRate = cast.ToFloat64(value)
		case "DropItemAliveMaxHours":
			config.DropItemAliveMaxHours = cast.ToFloat64(value)
		case "AutoResetGuildTimeNoOnlinePlayers":
			config.AutoResetGuildTimeNoOnlinePlayers = cast.ToFloat64(value)
		case "PalEggDefaultHatchingTime":
			config.PalEggDefaultHatchingTime = cast.ToFloat64(value)
		case "WorkSpeedRate":
			config.WorkSpeedRate = cast.ToFloat64(value)

		// int 类型的字段
		case "DropItemMaxNum":
			config.DropItemMaxNum = cast.ToInt(value)
		case "DropItemMaxNum_UNKO":
			config.DropItemMaxNumUNKO = cast.ToInt(value)
		case "BaseCampMaxNum":
			config.BaseCampMaxNum = cast.ToInt(value)
		case "BaseCampWorkerMaxNum":
			config.BaseCampWorkerMaxNum = cast.ToInt(value)
		case "GuildPlayerMaxNum":
			config.GuildPlayerMaxNum = cast.ToInt(value)
		case "CoopPlayerMaxNum":
			config.CoopPlayerMaxNum = cast.ToInt(value)
		case "ServerPlayerMaxNum":
			config.ServerPlayerMaxNum = cast.ToInt(value)
		case "PublicPort":
			config.PublicPort = cast.ToInt(value)
		case "RCONPort":
			config.RCONPort = cast.ToInt(value)

		// bool 类型的字段
		case "bEnablePlayerToPlayerDamage":
			config.BEnablePlayerToPlayerDamage = cast.ToBool(value)
		case "bEnableFriendlyFire":
			config.BEnableFriendlyFire = cast.ToBool(value)
		case "bEnableInvaderEnemy":
			config.BEnableInvaderEnemy = cast.ToBool(value)
		case "bActiveUNKO":
			config.BActiveUNKO = cast.ToBool(value)
		case "bEnableAimAssistPad":
			config.BEnableAimAssistPad = cast.ToBool(value)
		case "bEnableAimAssistKeyboard":
			config.BEnableAimAssistKeyboard = cast.ToBool(value)
		case "bAutoResetGuildNoOnlinePlayers":
			config.BAutoResetGuildNoOnlinePlayers = cast.ToBool(value)
		case "bIsMultiplay":
			config.BIsMultiplay = cast.ToBool(value)
		case "bIsPvP":
			config.BIsPvP = cast.ToBool(value)
		case "bCanPickupOtherGuildDeathPenaltyDrop":
			config.BCanPickupOtherGuildDeathPenaltyDrop = cast.ToBool(value)
		case "bEnableNonLoginPenalty":
			config.BEnableNonLoginPenalty = cast.ToBool(value)
		case "bEnableFastTravel":
			config.BEnableFastTravel = cast.ToBool(value)
		case "bIsStartLocationSelectByMap":
			config.BIsStartLocationSelectByMap = cast.ToBool(value)
		case "bExistPlayerAfterLogout":
			config.BExistPlayerAfterLogout = cast.ToBool(value)
		case "bEnableDefenseOtherGuildPlayer":
			config.BEnableDefenseOtherGuildPlayer = cast.ToBool(value)
		case "RCONEnabled":
			config.RCONEnabled = cast.ToBool(value)
		case "BUseAuth":
			config.BUseAuth = cast.ToBool(value)

		// string 类型的字段
		case "DeathPenalty":
			config.DeathPenalty = value
		case "ServerName":
			config.ServerName = value
		case "ServerDescription":
			config.ServerDescription = value
		case "AdminPassword":
			config.AdminPassword = value
		case "ServerPassword":
			config.ServerPassword = value
		case "PublicIP":
			config.PublicIP = value
		case "Region":
			config.Region = value
		case "BanListURL":
			if value == "" {
				config.BanListURL = "https://api.palworldgame.com/api/banlist.txt"
			} else {
				config.BanListURL = value
			}
		case "PalStaminaDecreaceRate":
			{
				config.PalStaminaDecreaseRate = cast.ToFloat64(value)
			}
		case "PlayerStaminaDecreaceRate":
			{
				config.PlayerStaminaDecreaseRate = cast.ToFloat64(value)
			}
		case "PlayerStomachDecreaceRate":
			config.PlayerStomachDecreaseRate = cast.ToFloat64(value)
		case "PalStomachDecreaceRate":
			config.PalStomachDecreaseRate = cast.ToFloat64(value)
		case "bUseAuth":
			config.BUseAuth = cast.ToBool(value)
		default:
			log.Error("Unknown field:", key)
		}
	}
}

// Struct2String converts the ServerConfig struct to a string representation.
func Struct2String(config *ServerConfig) string {
	// 使用字符串拼接构建格式化后的字符串
	result := fmt.Sprintf("Difficulty=%s,DayTimeSpeedRate=%f,NightTimeSpeedRate=%f,ExpRate=%f,PalCaptureRate=%f,PalSpawnNumRate=%f,PalDamageRateAttack=%f,PalDamageRateDefense=%f,PlayerDamageRateAttack=%f,PlayerDamageRateDefense=%f,PlayerStomachDecreaseRate=%f,PlayerStaminaDecreaseRate=%f,PlayerAutoHPRegeneRate=%f,PlayerAutoHpRegeneRateInSleep=%f,PalStomachDecreaseRate=%f,PalStaminaDecreaseRate=%f,PalAutoHPRegeneRate=%f,PalAutoHpRegeneRateInSleep=%f,BuildObjectDamageRate=%f,BuildObjectDeteriorationDamageRate=%f,CollectionDropRate=%f,CollectionObjectHpRate=%f,CollectionObjectRespawnSpeedRate=%f,EnemyDropItemRate=%f,DeathPenalty=%s,bEnablePlayerToPlayerDamage=%t,bEnableFriendlyFire=%t,bEnableInvaderEnemy=%t,bActiveUNKO=%t,bEnableAimAssistPad=%t,bEnableAimAssistKeyboard=%t,DropItemMaxNum=%d,DropItemMaxNum_UNKO=%d,BaseCampMaxNum=%d,BaseCampWorkerMaxNum=%d,DropItemAliveMaxHours=%f,bAutoResetGuildNoOnlinePlayers=%t,AutoResetGuildTimeNoOnlinePlayers=%f,GuildPlayerMaxNum=%d,PalEggDefaultHatchingTime=%f,WorkSpeedRate=%f,bIsMultiplay=%t,bIsPvP=%t,bCanPickupOtherGuildDeathPenaltyDrop=%t,bEnableNonLoginPenalty=%t,bEnableFastTravel=%t,bIsStartLocationSelectByMap=%t,bExistPlayerAfterLogout=%t,bEnableDefenseOtherGuildPlayer=%t,CoopPlayerMaxNum=%d,ServerPlayerMaxNum=%d,ServerName=%s,ServerDescription=%s,AdminPassword=%s,ServerPassword=%s,PublicPort=%d,PublicIP=%s,RCONEnabled=%t,RCONPort=%d,Region=%s,bUseAuth=%t,BanListURL=%s",
		config.Difficulty, config.DayTimeSpeedRate, config.NightTimeSpeedRate, config.ExpRate, config.PalCaptureRate, config.PalSpawnNumRate, config.PalDamageRateAttack, config.PalDamageRateDefense, config.PlayerDamageRateAttack, config.PlayerDamageRateDefense, config.PlayerStomachDecreaseRate, config.PlayerStaminaDecreaseRate, config.PlayerAutoHPRegeneRate, config.PlayerAutoHpRegeneRateInSleep, config.PalStomachDecreaseRate, config.PalStaminaDecreaseRate, config.PalAutoHPRegeneRate, config.PalAutoHpRegeneRateInSleep, config.BuildObjectDamageRate, config.BuildObjectDeteriorationDamageRate, config.CollectionDropRate, config.CollectionObjectHpRate, config.CollectionObjectRespawnSpeedRate, config.EnemyDropItemRate, config.DeathPenalty, config.BEnablePlayerToPlayerDamage, config.BEnableFriendlyFire, config.BEnableInvaderEnemy, config.BActiveUNKO, config.BEnableAimAssistPad, config.BEnableAimAssistKeyboard, config.DropItemMaxNum, config.DropItemMaxNumUNKO, config.BaseCampMaxNum, config.BaseCampWorkerMaxNum, config.DropItemAliveMaxHours, config.BAutoResetGuildNoOnlinePlayers, config.AutoResetGuildTimeNoOnlinePlayers, config.GuildPlayerMaxNum, config.PalEggDefaultHatchingTime, config.WorkSpeedRate, config.BIsMultiplay, config.BIsPvP, config.BCanPickupOtherGuildDeathPenaltyDrop, config.BEnableNonLoginPenalty, config.BEnableFastTravel, config.BIsStartLocationSelectByMap, config.BExistPlayerAfterLogout, config.BEnableDefenseOtherGuildPlayer, config.CoopPlayerMaxNum, config.ServerPlayerMaxNum, config.ServerName, config.ServerDescription, config.AdminPassword, config.ServerPassword, config.PublicPort, config.PublicIP, config.RCONEnabled, config.RCONPort, config.Region, config.BUseAuth, config.BanListURL)

	// 替换布尔类型的 true/false 为字符串格式
	result = strings.ReplaceAll(result, "true", "True")
	result = strings.ReplaceAll(result, "false", "False")
	result = "(" + result + ")"
	return result
}
