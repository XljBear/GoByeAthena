package structs

type MasterInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Uid            int         `json:"uid"`
		CanJoin        bool        `json:"canJoin"`
		AthenaId       int         `json:"athenaId"`
		FaceImg        string      `json:"faceImg"`
		NickName       string      `json:"nickName"`
		Grade          int         `json:"grade"`
		NowExp         int         `json:"nowExp"`
		CurrentExp     int         `json:"currentExp"`
		MaxExp         int         `json:"maxExp"`
		Strength       int         `json:"strength"`
		RoleBone       string      `json:"roleBone"`
		RoleImg        string      `json:"roleImg"`
		CallName       string      `json:"callName"`
		GreetText      interface{} `json:"greetText"`
		BirthdayText   interface{} `json:"birthdayText"`
		SignTaskId     interface{} `json:"signTaskId"`
		Amount         int         `json:"amount"`
		TaskListId     string      `json:"taskListId"`
		PrizeCenterUrl string      `json:"prizeCenterUrl"`
		GameRuleUrl    string      `json:"gameRuleUrl"`
		FeedbackUrl    string      `json:"feedbackUrl"`
		QuestionUrl    string      `json:"questionUrl"`
		ShopUrl        string      `json:"shopUrl"`
		BackpackUrl    string      `json:"backpackUrl"`
		AdventureBone  string      `json:"adventureBone"`
		StrengthPerDay interface{} `json:"strengthPerDay"`
		StoreGuideTips string      `json:"storeGuideTips"`
		MinStrength    int         `json:"minStrength"`
		BubbleInfo     struct {
			Type        int         `json:"type"`
			Text        string      `json:"text"`
			EndTime     interface{} `json:"endTime"`
			CurrentTime interface{} `json:"currentTime"`
		} `json:"bubbleInfo"`
	} `json:"data"`
	Errtag  int    `json:"errtag"`
	Errno   int    `json:"errno"`
	Msg     string `json:"msg"`
	ShowMsg string `json:"showMsg"`
}

type AdventureInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Ids     []int `json:"ids"`
		PlaceId int   `json:"placeId"`
	} `json:"data"`
	Errtag  int    `json:"errtag"`
	Errno   int    `json:"errno"`
	Msg     string `json:"msg"`
	ShowMsg string `json:"showMsg"`
}
