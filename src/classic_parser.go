package mcclassicparser

import (
    "bytes"
    "compress/gzip"
    "encoding/binary"
    "errors"
    "fmt"
    "os"
    "reflect"

    "github.com/jkeys089/jserial"
)

type ClassicParser struct {
    // for testing
    // Data map[string]any
}

func (cp *ClassicParser) Parse(filepath string) (*ClassicWorld, error) {
    // Check that given file is a gzipped file
    gzbytes, _ := os.ReadFile(filepath)

    if gzMagic := binary.BigEndian.Uint16(gzbytes[0:2]); gzMagic != 0x1f8b {
        return nil, errors.New("Not a GZIP file.")
    }

    // Get uncompress file size
    uncompressedSize := binary.LittleEndian.Uint32(gzbytes[len(gzbytes) - 4:])

    // Decompress Gzip
    gzReader, _ := gzip.NewReader(bytes.NewBuffer(gzbytes))
    defer gzReader.Close()

    // Parse Minecraft Classic world into ClassicWorld struct
    cw := ClassicWorld{}

    var magic int32
    var version byte
    var java_object = make([]byte, uncompressedSize - 5)
    binary.Read(gzReader, binary.BigEndian, &magic)
    binary.Read(gzReader, binary.BigEndian, &version)
    binary.Read(gzReader, binary.BigEndian, &java_object)

    if magic != 0x271bb788 || version != 0x02 {
        return nil, errors.New("Not a vaild Minecraft Classic 0.0.14a - 0.30 world save!")
    }

    fmt.Println("Found a vaild Minecraft Classic 0.0.14a - 0.30 world save!")

    objects, _ := jserial.ParseSerializedObject(java_object)
    data := objects[0].(map[string]any)
    // cp.Data = data

    cw.Blocks = cp.parseBlocks(data["blocks"].([]any))

    cw.CloudColor = cp.getOrDefault(data["cloudColor"], (int32)(16777215)).(int32)
    cw.CreateTime = cp.getOrDefault(data["createTime"], (int64)(0)).(int64)

    cw.CreativeMode = cp.getOrDefault(data["creativeMode"], (bool)(false)).(bool)
    cw.Creator = cp.getOrDefault(data["creator"], (string)("")).(string)
    cw.Depth = cp.getOrDefault(data["depth"], (int32)(0)).(int32)
    cw.Entities = cp.parseEntities(data)

    cw.FogColor = cp.getOrDefault(data["fogColor"], (int32)(16777215)).(int32)

    cw.GrowTrees = cp.getOrDefault(data["growTrees"], (bool)(false)).(bool)
    cw.Height = cp.getOrDefault(data["height"], (int32)(0)).(int32)
    cw.Name = cp.getOrDefault(data["name"], (string)("")).(string)

    cw.Player = cp.parsePlayer(cp.getOrDefault(data["player"], (map[string]any)(nil)).(map[string]any))
    cw.RotSpawn = cp.getOrDefault(data["rotSpawn"], (float32)(0)).(float32)

    cw.SkyColor = cp.getOrDefault(data["skyColor"], (int32)(10079487)).(int32)

    cw.WaterLevel = cp.getOrDefault(data["waterLevel"], (int32)(cw.Height / 2)).(int32)
    cw.Width = cp.getOrDefault(data["width"], (int32)(0)).(int32)
    cw.XSpawn = cp.getOrDefault(data["xSpawn"], (int32)(0)).(int32)
    cw.YSpawn = cp.getOrDefault(data["ySpawn"], (int32)(0)).(int32)
    cw.ZSpawn = cp.getOrDefault(data["zSpawn"], (int32)(0)).(int32)

    return &cw, nil
}

func (cp *ClassicParser) instanceof(data any, typePtr any) bool {
    return reflect.TypeOf(data) == reflect.TypeOf(typePtr)
}

func (cp *ClassicParser) getOrDefault(data any, defaultValue any) any {
    if ok := cp.instanceof(data, defaultValue); ok {
        return data
    }
    return defaultValue
}

func (cp *ClassicParser) toInt32Array(arr []any) []int32 {
    var out []int32
    for i := 0; i < len(arr); i++ {
        out = append(out, arr[i].(int32))
    }
    return out
}

func (cp *ClassicParser) parseBlocks(blocks []any) []int8 {
    var out []int8
    for i := 0; i < len(blocks); i++ {
        out = append(out, blocks[i].(int8))
    }
    return out
}

func (cp *ClassicParser) parseEntities(data map[string]any) []ClassicEntity {
    var entities []ClassicEntity
    var entitiesArray []any

    if data == nil {
        return entities
    }

    if ok := cp.instanceof(data["entities"], (map[string]any)(nil)); ok {
        entitiesArray = data["entities"].(map[string]any)["value"].([]any)
    } else if ok := cp.instanceof(data["blockMap"], (map[string]any)(nil)); ok {
        entitiesArray = data["blockMap"].(map[string]any)["all"].(map[string]any)["value"].([]any)
    } else {
        return entities
    }

    if len(entitiesArray) == 0 {
        return entities
    }

    for _, entityAny := range entitiesArray {
        entityData := entityAny.(map[string]any)
        var entity ClassicEntity

        entity.AirSupply = cp.getOrDefault(entityData["airSupply"], (int32)(0)).(int32)
        entity.AllowAlpha = cp.getOrDefault(entityData["allowAlpha"], (bool)(false)).(bool)
        entity.AnimStep = cp.getOrDefault(entityData["animStep"], (float32)(0)).(float32)
        entity.AnimStepO = cp.getOrDefault(entityData["animStepO"], (float32)(0)).(float32)
        entity.AttackTime = cp.getOrDefault(entityData["attackTime"], (int32)(0)).(int32)

        bb := make(map[string]float32)
        bb["x0"] = entityData["bb"].(map[string]any)["x0"].(float32)
        bb["y0"] = entityData["bb"].(map[string]any)["y0"].(float32)
        bb["z0"] = entityData["bb"].(map[string]any)["z0"].(float32)
        bb["x1"] = entityData["bb"].(map[string]any)["x1"].(float32)
        bb["y1"] = entityData["bb"].(map[string]any)["y1"].(float32)
        bb["z1"] = entityData["bb"].(map[string]any)["z1"].(float32)
        entity.Bb = bb

        entity.BbHeight = cp.getOrDefault(entityData["bbHeight"], (float32)(0)).(float32)
        entity.BbWidth = cp.getOrDefault(entityData["bbWidth"], (float32)(0)).(float32)
        entity.BobStrength = cp.getOrDefault(entityData["bobStrength"], (float32)(0)).(float32)
        entity.Collision = cp.getOrDefault(entityData["collision"], (bool)(false)).(bool)
        entity.Dead = cp.getOrDefault(entityData["dead"], (bool)(false)).(bool)
        entity.DeathScore = cp.getOrDefault(entityData["deathScore"], (int32)(0)).(int32)
        entity.DeathTime = cp.getOrDefault(entityData["deathTime"], (int32)(0)).(int32)
        entity.FallDistance = cp.getOrDefault(entityData["fallDistance"], (float32)(0)).(float32)
        entity.FootSize = cp.getOrDefault(entityData["footSize"], (float32)(0)).(float32)
        entity.HasHair = cp.getOrDefault(entityData["hasHair"], (bool)(false)).(bool)
        entity.Health = cp.getOrDefault(entityData["health"], (int32)(0)).(int32)
        entity.HeightOffset = cp.getOrDefault(entityData["heightOffset"], (float32)(0)).(float32)
        entity.HorizontalCollision = cp.getOrDefault(entityData["horizontalCollision"], (bool)(false)).(bool)

        entity.Hovered = cp.getOrDefault(data["hovered"], (bool)(false)).(bool)
        entity.HurtDir = cp.getOrDefault(entityData["hurtDir"], (float32)(0)).(float32)
        entity.HurtDuration = cp.getOrDefault(entityData["hurtDuration"], (int32)(0)).(int32)
        entity.HurtTime = cp.getOrDefault(entityData["hurtTime"], (int32)(0)).(int32)
        entity.InvulnerableDuration = cp.getOrDefault(entityData["invulnerableDuration"], (int32)(0)).(int32)
        entity.InvulnerableTime = cp.getOrDefault(entityData["invulnerableTime"], (int32)(0)).(int32)
        entity.LastHealth = cp.getOrDefault(entityData["lastHealth"], (int32)(0)).(int32)
        entity.MakeStepSound = cp.getOrDefault(entityData["makeStepSound"], (bool)(false)).(bool)
        entity.ModelName = cp.getOrDefault(entityData["modelName"], (string)("")).(string)

        entity.NoPhysics = cp.getOrDefault(data["noPhysics"], (bool)(false)).(bool)
        entity.ORun = cp.getOrDefault(entityData["oRun"], (float32)(0)).(float32)
        entity.OTilt = cp.getOrDefault(entityData["oTilt"], (float32)(0)).(float32)
        entity.OnGround = cp.getOrDefault(entityData["onGround"], (bool)(false)).(bool)

        entity.Pushthrough = cp.getOrDefault(data["pushthrough"], (float32)(0)).(float32)
        entity.Removed = cp.getOrDefault(entityData["removed"], (bool)(false)).(bool)

        entity.RenderOffset = cp.getOrDefault(data["renderOffset"], (float32)(0)).(float32)
        entity.Rot = cp.getOrDefault(entityData["rot"], (float32)(0)).(float32)
        entity.RotA = cp.getOrDefault(entityData["rotA"], (float32)(0)).(float32)
        entity.RotOffs = cp.getOrDefault(entityData["rotOffs"], (float32)(0)).(float32)
        entity.Run = cp.getOrDefault(entityData["run"], (float32)(0)).(float32)
        entity.Slide = cp.getOrDefault(entityData["slide"], (bool)(false)).(bool)
        entity.Speed = cp.getOrDefault(entityData["speed"], (float32)(0)).(float32)
        entity.TextureId = cp.getOrDefault(entityData["textureId"], (int32)(0)).(int32)
        entity.TextureName = cp.getOrDefault(entityData["textureName"], (string)("/char.png")).(string)
        entity.TickCount = cp.getOrDefault(entityData["tickCount"], (int32)(0)).(int32)
        entity.Tilt = cp.getOrDefault(entityData["tilt"], (float32)(0)).(float32)
        entity.TimeOffs = cp.getOrDefault(entityData["timeOffs"], (float32)(0)).(float32)
        entity.WalkDist = cp.getOrDefault(entityData["walkDist"], (float32)(0)).(float32)
        entity.WalkDistO = cp.getOrDefault(entityData["walkDistO"], (float32)(0)).(float32)
        entity.X = cp.getOrDefault(entityData["x"], (float32)(0)).(float32)
        entity.XOld = cp.getOrDefault(entityData["xOld"], (float32)(0)).(float32)
        entity.XRot = cp.getOrDefault(entityData["xRot"], (float32)(0)).(float32)
        entity.XRotO = cp.getOrDefault(entityData["xRotO"], (float32)(0)).(float32)
        entity.Xd = cp.getOrDefault(entityData["xd"], (float32)(0)).(float32)
        entity.Xo = cp.getOrDefault(entityData["xo"], (float32)(0)).(float32)
        entity.Y = cp.getOrDefault(entityData["y"], (float32)(0)).(float32)
        entity.YBodyRot = cp.getOrDefault(entityData["yBodyRot"], (float32)(0)).(float32)
        entity.YBodyRotO = cp.getOrDefault(entityData["yBodyRotO"], (float32)(0)).(float32)
        entity.YOld = cp.getOrDefault(entityData["yOld"], (float32)(0)).(float32)
        entity.YRot = cp.getOrDefault(entityData["yRot"], (float32)(0)).(float32)
        entity.YRotO = cp.getOrDefault(entityData["yRotO"], (float32)(0)).(float32)
        entity.YSlideOffset = cp.getOrDefault(entityData["ySlideOffset"], (float32)(0)).(float32)
        entity.Yd = cp.getOrDefault(entityData["yd"], (float32)(0)).(float32)
        entity.Yo = cp.getOrDefault(entityData["yo"], (float32)(0)).(float32)
        entity.Z = cp.getOrDefault(entityData["z"], (float32)(0)).(float32)
        entity.ZOld = cp.getOrDefault(entityData["zOld"], (float32)(0)).(float32)
        entity.Zd = cp.getOrDefault(entityData["zd"], (float32)(0)).(float32)
        entity.Zo = cp.getOrDefault(entityData["zo"], (float32)(0)).(float32)

        entities = append(entities, entity)
    }
    return entities
}

func (cp *ClassicParser) parsePlayer(data map[string]any) ClassicPlayer {
    var player ClassicPlayer

    if data == nil {
        return player
    }

    player.AirSupply = cp.getOrDefault(data["airSupply"], (int32)(0)).(int32)
    player.AllowAlpha = cp.getOrDefault(data["allowAlpha"], (bool)(false)).(bool)
    player.AnimStep = cp.getOrDefault(data["animStep"], (float32)(0)).(float32)
    player.AnimStepO = cp.getOrDefault(data["animStepO"], (float32)(0)).(float32)
    player.Arrows = cp.getOrDefault(data["arrows"], (int32)(0)).(int32)
    player.AttackTime = cp.getOrDefault(data["attackTime"], (int32)(0)).(int32)

    bb := make(map[string]float32)
    bb["x0"] = data["bb"].(map[string]any)["x0"].(float32)
    bb["y0"] = data["bb"].(map[string]any)["y0"].(float32)
    bb["z0"] = data["bb"].(map[string]any)["z0"].(float32)
    bb["x1"] = data["bb"].(map[string]any)["x1"].(float32)
    bb["y1"] = data["bb"].(map[string]any)["y1"].(float32)
    bb["z1"] = data["bb"].(map[string]any)["z1"].(float32)
    player.Bb = bb

    player.BbHeight = cp.getOrDefault(data["bbHeight"], (float32)(0)).(float32)
    player.BbWidth = cp.getOrDefault(data["bbWidth"], (float32)(0)).(float32)
    player.Bob = cp.getOrDefault(data["bob"], (float32)(0)).(float32)
    player.BobStrength = cp.getOrDefault(data["bobStrength"], (float32)(0)).(float32)
    player.Collision = cp.getOrDefault(data["collision"], (bool)(false)).(bool)
    player.Dead = cp.getOrDefault(data["dead"], (bool)(false)).(bool)
    player.DeathScore = cp.getOrDefault(data["deathScore"], (int32)(0)).(int32)
    player.DeathTime = cp.getOrDefault(data["deathTime"], (int32)(0)).(int32)
    player.FallDistance = cp.getOrDefault(data["fallDistance"], (float32)(0)).(float32)
    player.FootSize = cp.getOrDefault(data["footSize"], (float32)(0)).(float32)
    player.HasHair = cp.getOrDefault(data["hasHair"], (bool)(false)).(bool)
    player.Health = cp.getOrDefault(data["health"], (int32)(0)).(int32)
    player.HeightOffset = cp.getOrDefault(data["heightOffset"], (float32)(0)).(float32)
    player.HorizontalCollision = cp.getOrDefault(data["horizontalCollision"], (bool)(false)).(bool)

    player.Hovered = cp.getOrDefault(data["hovered"], (bool)(false)).(bool)
    player.HurtDir = cp.getOrDefault(data["hurtDir"], (float32)(0)).(float32)
    player.HurtDuration = cp.getOrDefault(data["hurtDuration"], (int32)(0)).(int32)
    player.HurtTime = cp.getOrDefault(data["hurtTime"], (int32)(0)).(int32)

    inventory := make(map[string]any)
    inventory["slots"] = cp.toInt32Array(data["inventory"].(map[string]any)["slots"].([]any))
    inventory["count"] = cp.toInt32Array(data["inventory"].(map[string]any)["count"].([]any))
    inventory["selected"] = cp.getOrDefault(data["inventory"].(map[string]any)["selected"], (int32)(0)).(int32)
    player.Inventory = inventory

    player.InvulnerableDuration = cp.getOrDefault(data["invulnerableDuration"], (int32)(0)).(int32)
    player.InvulnerableTime = cp.getOrDefault(data["invulnerableTime"], (int32)(0)).(int32)
    player.LastHealth = cp.getOrDefault(data["lastHealth"], (int32)(0)).(int32)
    player.MakeStepSound = cp.getOrDefault(data["makeStepSound"], (bool)(false)).(bool)
    player.ModelName = cp.getOrDefault(data["modelName"], (string)("")).(string)

    player.NoPhysics = cp.getOrDefault(data["noPhysics"], (bool)(false)).(bool)
    player.OBob = cp.getOrDefault(data["oBob"], (float32)(0)).(float32)
    player.ORun = cp.getOrDefault(data["oRun"], (float32)(0)).(float32)
    player.OTilt = cp.getOrDefault(data["oTilt"], (float32)(0)).(float32)
    player.OnGround = cp.getOrDefault(data["onGround"], (bool)(false)).(bool)

    player.Pushthrough = cp.getOrDefault(data["pushthrough"], (float32)(0)).(float32)
    player.Removed = cp.getOrDefault(data["removed"], (bool)(false)).(bool)

    player.RenderOffset = cp.getOrDefault(data["renderOffset"], (float32)(0)).(float32)
    player.Rot = cp.getOrDefault(data["rot"], (float32)(0)).(float32)
    player.RotA = cp.getOrDefault(data["rotA"], (float32)(0)).(float32)
    player.RotOffs = cp.getOrDefault(data["rotOffs"], (float32)(0)).(float32)
    player.Run = cp.getOrDefault(data["run"], (float32)(0)).(float32)
    player.Score = cp.getOrDefault(data["score"], (int32)(0)).(int32)
    player.Slide = cp.getOrDefault(data["slide"], (bool)(false)).(bool)
    player.Speed = cp.getOrDefault(data["speed"], (float32)(0)).(float32)
    player.TextureId = cp.getOrDefault(data["textureId"], (int32)(0)).(int32)
    player.TextureName = cp.getOrDefault(data["textureName"], (string)("/char.png")).(string)
    player.TickCount = cp.getOrDefault(data["tickCount"], (int32)(0)).(int32)
    player.Tilt = cp.getOrDefault(data["tilt"], (float32)(0)).(float32)
    player.TimeOffs = cp.getOrDefault(data["timeOffs"], (float32)(0)).(float32)
    player.UserType = cp.getOrDefault(data["userType"], (int8)(0)).(int8)
    player.WalkDist = cp.getOrDefault(data["walkDist"], (float32)(0)).(float32)
    player.WalkDistO = cp.getOrDefault(data["walkDistO"], (float32)(0)).(float32)
    player.X = cp.getOrDefault(data["x"], (float32)(0)).(float32)
    player.XOld = cp.getOrDefault(data["xOld"], (float32)(0)).(float32)
    player.XRot = cp.getOrDefault(data["xRot"], (float32)(0)).(float32)
    player.XRotO = cp.getOrDefault(data["xRotO"], (float32)(0)).(float32)
    player.Xd = cp.getOrDefault(data["xd"], (float32)(0)).(float32)
    player.Xo = cp.getOrDefault(data["xo"], (float32)(0)).(float32)
    player.Y = cp.getOrDefault(data["y"], (float32)(0)).(float32)
    player.YBodyRot = cp.getOrDefault(data["yBodyRot"], (float32)(0)).(float32)
    player.YBodyRotO = cp.getOrDefault(data["yBodyRotO"], (float32)(0)).(float32)
    player.YOld = cp.getOrDefault(data["yOld"], (float32)(0)).(float32)
    player.YRot = cp.getOrDefault(data["yRot"], (float32)(0)).(float32)
    player.YRotO = cp.getOrDefault(data["yRotO"], (float32)(0)).(float32)
    player.YSlideOffset = cp.getOrDefault(data["ySlideOffset"], (float32)(0)).(float32)
    player.Yd = cp.getOrDefault(data["yd"], (float32)(0)).(float32)
    player.Yo = cp.getOrDefault(data["yo"], (float32)(0)).(float32)
    player.Z = cp.getOrDefault(data["z"], (float32)(0)).(float32)
    player.ZOld = cp.getOrDefault(data["zOld"], (float32)(0)).(float32)
    player.Zd = cp.getOrDefault(data["zd"], (float32)(0)).(float32)
    player.Zo = cp.getOrDefault(data["zo"], (float32)(0)).(float32)

    return player
}
