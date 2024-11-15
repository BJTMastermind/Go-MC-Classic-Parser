package mc_classic_parser

import "fmt"

type ClassicWorld struct {
    Blocks []byte
    CloudColor int32
    CreateTime int64
    CreativeMode bool
    Creator string
    Depth int32
    Entities []ClassicEntity
    FogColor int32
    GrowTrees bool
    Height int32
    Name string
    Player ClassicPlayer
    RotSpawn float32
    SkyColor int32
    WaterLevel int32
    Width int32
    XSpawn int32
    YSpawn int32
    ZSpawn int32
}

// Returns the water level of the world as a float32
func (world *ClassicWorld) GetWaterLevel() float32 {
    return float32(world.WaterLevel)
}

func (world *ClassicWorld) GetWorldSpawn() (int32, int32, int32) {
    return world.XSpawn, world.YSpawn, world.ZSpawn
}

func (player ClassicWorld) ToString() string {
    format := fmt.Sprintf("%+v\n", player)
    return format
}