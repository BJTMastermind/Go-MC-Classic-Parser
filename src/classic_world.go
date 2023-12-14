package mcclassicparser

import "fmt"

type ClassicWorld struct {
    Blocks []int8
    CreateTime int64
    Creator string
    Depth int32
    Entities []ClassicEntity
    Height int32
    Name string
    networkMode bool
    RotSpawn float32
    tickCount int32
    unprocessed int32
    Width int32
    XSpawn int32
    YSpawn int32
    ZSpawn int32
}

// Returns the water level of the world or -1 & error if not present
// func (world *ClassicWorld) GetWaterLevel() (int32, error) {
//     var waterLevel int32
//     waterLevel, ok := world.Data["waterLevel"].(int32)
//     if !ok {
//         return -1, errors.New("waterLevel not found")
//     }
//     return waterLevel, nil
// }

func (world *ClassicWorld) GetWorldSpawn() []int32 {
    spawn := []int32{
        world.XSpawn,
        world.YSpawn,
        world.ZSpawn,
    }
    return spawn
}

// Returns a ClassicPlayer if present otherwise returns an error
// func (world *ClassicWorld) GetPlayer() (*ClassicPlayer, error) {
//     playerdata, ok := world.Data["player"].(map[string]any)
//     if ok {
//         player := new(ClassicPlayer)
//         player.Ai = playerdata["ai"].(map[string]any)
//         player.airSupply = playerdata["airSupply"].(int32)
//         player.AllowAlpha = playerdata["allowAlpha"].(bool)
//         player.AnimStep = playerdata["animStep"].(float32)
//         player.AnimStepO = playerdata["animStepO"].(float32)
//         player.arrows = playerdata["arrows"].(int32)
//         player.attackTime = playerdata["attackTime"].(int32)
//         player.Bb = playerdata["bb"].(map[string]any)
//         player.BbHeight = playerdata["bbHeight"].(float32)
//         player.BbWidth = playerdata["bbWidth"].(float32)
//         // player.blockMap = playerdata["blockMap"].(map[string]any)
//         player.Bob = playerdata["bob"].(float32)
//         player.BobStrength = playerdata["bobStrength"].(float32)
//         player.Collision = playerdata["collision"].(bool)
//         player.Dead = playerdata["dead"].(bool)
//         player.DeathScore = playerdata["deathScore"].(int32)
//         player.deathTime = playerdata["deathTime"].(int32)
//         player.fallDistance = playerdata["fallDistance"].(float32)
//         player.FootSize = playerdata["footSize"].(float32)
//         player.HasHair = playerdata["hasHair"].(bool)
//         player.health = playerdata["health"].(int32)
//         player.HeightOffset = playerdata["heightOffset"].(float32)
//         player.HorizontalCollision = playerdata["horizontalCollision"].(bool)
//         // player.Hovered = playerdata["hovered"].(bool)
//         player.HurtDir = playerdata["hurtDir"].(float32)
//         player.HurtDuration = playerdata["hurtDuration"].(int32)
//         player.hurtTime = playerdata["hurtTime"].(int32)
//         player.inventory = playerdata["inventory"].(map[string]any)
//         player.InvulnerableDuration = playerdata["invulnerableDuration"].(int32)
//         player.InvulnerableTime = playerdata["invulnerableTime"].(int32)
//         player.LastHealth = playerdata["lastHealth"].(int32)
//         // player.Level = *world
//         player.MakeStepSound = playerdata["makeStepSound"].(bool)
//         player.Map = playerdata
//         player.ModelName = playerdata["modelName"].(string)
//         player.NextStep = playerdata["nextStep"].(int32)
//         // player.NoPhysics = playerdata["noPhysics"].(bool)
//         player.OBob = playerdata["oBob"].(float32)
//         player.ORun = playerdata["oRun"].(float32)
//         player.OTilt = playerdata["oTilt"].(float32)
//         player.OnGround = playerdata["onGround"].(bool)
//         // player.Pushthrough = playerdata["pushthrough"].(float32)
//         player.Removed = playerdata["removed"].(bool)
//         // player.RenderOffset = playerdata["renderOffset"].(float32)
//         player.Rot = playerdata["rot"].(float32)
//         player.RotA = playerdata["rotA"].(float32)
//         player.RotOffs = playerdata["rotOffs"].(float32)
//         player.Run = playerdata["run"].(float32)
//         player.score = playerdata["score"].(int32)
//         player.Slide = playerdata["slide"].(bool)
//         player.Speed = playerdata["speed"].(float32)
//         player.TextureId = playerdata["textureId"].(int32)
//         player.TextureName = playerdata["textureName"].(string)
//         player.TickCount = playerdata["tickCount"].(int32)
//         player.Tilt = playerdata["tilt"].(float32)
//         player.TimeOffs = playerdata["timeOffs"].(float32)
//         player.UserType = playerdata["userType"].(int8)
//         player.WalkDist = playerdata["walkDist"].(float32)
//         player.WalkDistO = playerdata["walkDistO"].(float32)
//         player.x = playerdata["x"].(float32)
//         player.XOld = playerdata["xOld"].(float32)
//         player.xRot = playerdata["xRot"].(float32)
//         player.XRotO = playerdata["xRotO"].(float32)
//         player.Xd = playerdata["xd"].(float32)
//         player.Xo = playerdata["xo"].(float32)
//         player.y = playerdata["y"].(float32)
//         player.YBodyRot = playerdata["yBodyRot"].(float32)
//         player.YBodyRotO = playerdata["yBodyRotO"].(float32)
//         player.YOld = playerdata["yOld"].(float32)
//         player.yRot = playerdata["yRot"].(float32)
//         player.YRotO = playerdata["yRotO"].(float32)
//         player.YSlideOffset = playerdata["ySlideOffset"].(float32)
//         player.Yd = playerdata["yd"].(float32)
//         player.Yo = playerdata["yo"].(float32)
//         player.z = playerdata["z"].(float32)
//         player.ZOld = playerdata["zOld"].(float32)
//         player.Zd = playerdata["zd"].(float32)
//         player.Zo = playerdata["zo"].(float32)

//         return player, nil
//     }
//     return nil, errors.New("No player found")
// }

func (world *ClassicWorld) ToString() string {
    blocksLen := len(world.Blocks)
    blocksSuffix := "entry"
    if len(world.Blocks) > 1 || len(world.Blocks) == 0 {
        blocksSuffix = "entries"
    }

    entitiesLen := len(world.Entities)
    entitiesSuffix := "entry"
    if len(world.Entities) > 1 || len(world.Entities) == 0 {
        entitiesSuffix = "entries"
    }

    s := fmt.Sprintf("ClassicWorld: {\n    Blocks: [%d %s]" +
        "\n    CreateTime: %dL" +
        "\n    Creator: \"%s\"" +
        "\n    Depth: %d" +
        "\n    Entities: [%d %s]" +
        "\n    Height: %d" +
        "\n    Name: \"%s\"" +
        "\n    networkMod: %t" +
        "\n    RotSpawn: %fF" +
        "\n    tickCount: %d" +
        "\n    unprocessed: %d" +
        "\n    Width: %d" +
        "\n    XSpawn: %d" +
        "\n    YSpawn: %d" +
        "\n    ZSpawn: %d\n}",
        blocksLen,
        blocksSuffix,
        world.CreateTime,
        world.Creator,
        world.Depth,
        entitiesLen,
        entitiesSuffix,
        world.Height,
        world.Name,
        world.networkMode,
        world.RotSpawn,
        world.tickCount,
        world.unprocessed,
        world.Width,
        world.XSpawn,
        world.YSpawn,
        world.ZSpawn)
    return s
}