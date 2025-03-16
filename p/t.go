package p

const (
	P_GameEnter_Jackaroo int32 = iota + 1
	P_GameWin_Jackaroo
	P_GameEnter_Baloot
	P_GameWin_Baloot
	P_GameEnter_Domino
	P_GameWin_Domino
	P_GameEnter_MonsterCrush
	P_GameWin_MonsterCrush
	P_GameEnter_Ludo
	P_GameWin_Ludo
	P_GameEnter_Uno
	P_GameWin_Uno
	P_GameEnter_Carrom
	P_GameWin_Carrom

	P_Rechage         = 501 // 充值
	P_SpecialOfferAdd = 502 // specialoffer每日奖励
	P_GameLevelUp     = 503 // 游戏升级奖励
)

const (
	G_Jackaroo     = 1
	G_Baloot       = 2
	G_Domino       = 3
	G_MonsterCrush = 4
	G_Ludo         = 5
	G_Uno          = 6
	G_Carrom       = 7
)

var (
	G2P_Enter = map[int32]int32{
		G_Jackaroo:     P_GameEnter_Jackaroo,
		G_Baloot:       P_GameEnter_Baloot,
		G_Domino:       P_GameEnter_Domino,
		G_MonsterCrush: P_GameEnter_MonsterCrush,
		G_Ludo:         P_GameEnter_Ludo,
		G_Uno:          P_GameEnter_Uno,
		G_Carrom:       P_GameEnter_Carrom,
	}

	G2P_Win = map[int32]int32{
		G_Jackaroo:     P_GameWin_Jackaroo,
		G_Baloot:       P_GameWin_Baloot,
		G_Domino:       P_GameWin_Domino,
		G_MonsterCrush: P_GameWin_MonsterCrush,
		G_Ludo:         P_GameWin_Ludo,
		G_Uno:          P_GameWin_Uno,
		G_Carrom:       P_GameWin_Carrom,
	}
)
