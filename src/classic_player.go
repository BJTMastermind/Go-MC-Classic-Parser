package mcclassicparser

type ClassicPlayer struct {
    /*extends*/ ClassicEntity
}

func (player *ClassicPlayer) GetAirSupply() int32 {
    return player.airSupply
}

func (player *ClassicPlayer) GetArrowCount() int32 {
    return player.arrows
}

func (player *ClassicPlayer) GetAttackTime() int32 {
    return player.attackTime
}

func (player *ClassicPlayer) GetDeathTime() int32 {
    return player.deathTime
}

func (player *ClassicPlayer) GetFallDistance() float32 {
    return player.fallDistance
}

func (player *ClassicPlayer) GetHeath() int32 {
    return player.health
}

func (player *ClassicPlayer) GetHurtTime() int32 {
    return player.hurtTime
}

func (player *ClassicPlayer) GetInventory() map[string]any {
    inventory := map[string]any{
        "ids": player.inventory["slots"],
        "count": player.inventory["count"],
    }
    return inventory
}

func (player *ClassicPlayer) GetPosition() []float32 {
    pos := []float32{player.x, player.y, player.z}
    return pos
}

func (player *ClassicPlayer) GetRotation() []float32 {
    rot := []float32{player.xRot, player.yRot}
    return rot
}

func (player *ClassicPlayer) GetScore() int32 {
    return player.score
}