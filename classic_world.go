package mc_classic_parser

import "fmt"

type ClassicWorld struct {
    Blocks []int8
    CloudColor int32 // make optional
    CreateTime int64
    CreativeMode bool // make optional
    Creator string
    Depth int32
    Entities []ClassicEntity // empty if no entities
    FogColor int32 // make optional
    GrowTrees bool // make optional
    Height int32
    Name string
    Player ClassicPlayer // make optional
    RotSpawn float32
    SkyColor int32 // make optional
    WaterLevel int32 // make optional
    Width int32
    XSpawn int32
    YSpawn int32
    ZSpawn int32
}

// Returns the water level of the world or -1 & error if not present
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