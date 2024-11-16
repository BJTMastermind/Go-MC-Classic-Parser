package mc_classic_parser

import "fmt"

type ClassicPlayer struct {
    /*extends*/ ClassicEntity
    Arrows int32
    Bob float32
    Inventory map[string]any // map[string]any{"slots": []int32, "selected": int32, "count": []int32}
    OBob float32
    Score int32
    UserType int8
}

func (player ClassicPlayer) ToString() string {
    format := fmt.Sprintf("%+v\n", player)
    return format
}