package mcclassicparser

import (
    "bytes"
    "compress/gzip"
    "encoding/binary"
    "fmt"
    "os"

    "github.com/jkeys089/jserial"
)

type ClassicParser struct {
    Data map[string]any
}

func (cp *ClassicParser) Parse(filepath string) ClassicWorld {
    // Get uncompress file size
    gzbytes, _ := os.ReadFile(filepath)
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

    if magic == 0x271bb788 && version == 0x02 {
        fmt.Println("Found vaild Minecraft Classic world save!")

        objects, _ := jserial.ParseSerializedObject(java_object)
        data := objects[0].(map[string]any)
        cp.Data = data

        cw.Blocks = cp.parseBlocks(data["blocks"].([]any))
        cw.CreateTime = data["createTime"].(int64)
        cw.Creator = data["creator"].(string)
        cw.Depth = data["depth"].(int32)
        // cw.Entities = cp.parseEntities(data)
        cw.Height = data["height"].(int32)
        cw.Name = data["name"].(string)
        // cw.networkMode = data["networkMode"].(bool)
        cw.RotSpawn = data["rotSpawn"].(float32)
        cw.tickCount = data["tickCount"].(int32)
        cw.unprocessed = data["unprocessed"].(int32)
        cw.Width = data["width"].(int32)
        cw.XSpawn = data["xSpawn"].(int32)
        cw.YSpawn = data["ySpawn"].(int32)
        cw.ZSpawn = data["zSpawn"].(int32)
    } else {
        fmt.Println("Invaild Minecraft Classic world save!")
    }
    return cw
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
    entitiesArray, ok := data["entities"].(map[string]any)["value"].([]any)
    if !ok {
        entitiesArray, ok = data["blockMap"].(map[string]any)["all"].(map[string]any)["value"].([]any)
        if !ok {
            return entities
        }
    }

    if len(entitiesArray) == 0 {
        return entities
    }

    for _, entityAny := range entitiesArray {
        entityData := entityAny.(map[string]any)
        var entity ClassicEntity

        entity.Ai = entityData["ai"].(map[string]any)
        entity.airSupply = entityData["airSupply"].(int32)
        entity.AllowAlpha = entityData["allowAlpha"].(bool)
        entity.AnimStep = entityData["animStep"].(float32)
        entity.AnimStepO = entityData["animStepO"].(float32)
        entity.arrows = entityData["arrows"].(int32)
        entity.attackTime = entityData["attackTime"].(int32)
        entity.Bb = entityData["bb"].(map[string]any)
        entity.BbHeight = entityData["bbHeight"].(float32)
        entity.BbWidth = entityData["bbWidth"].(float32)
        // entity.blockMap = entityData["blockMap"].(map[string]any)
        entity.Bob = entityData["bob"].(float32)
        entity.BobStrength = entityData["bobStrength"].(float32)
        entity.Collision = entityData["collision"].(bool)
        entity.Dead = entityData["dead"].(bool)
        entity.DeathScore = entityData["deathScore"].(int32)
        entity.deathTime = entityData["deathTime"].(int32)
        entity.fallDistance = entityData["fallDistance"].(float32)
        entity.FootSize = entityData["footSize"].(float32)
        entity.HasHair = entityData["hasHair"].(bool)
        entity.health = entityData["health"].(int32)
        entity.HeightOffset = entityData["heightOffset"].(float32)
        entity.HorizontalCollision = entityData["horizontalCollision"].(bool)
        entity.Hovered = entityData["hovered"].(bool)
        entity.HurtDir = entityData["hurtDir"].(float32)
        entity.HurtDuration = entityData["hurtDuration"].(int32)
        entity.hurtTime = entityData["hurtTime"].(int32)
        entity.inventory = entityData["inventory"].(map[string]any)
        entity.InvulnerableDuration = entityData["invulnerableDuration"].(int32)
        entity.InvulnerableTime = entityData["invulnerableTime"].(int32)
        entity.LastHealth = entityData["lastHealth"].(int32)
        // entity.Level = entityData["level"].(ClassicWorld)
        entity.MakeStepSound = entityData["makeStepSound"].(bool)
        entity.Map = entityData["map"].(map[string]any)
        entity.ModelName = entityData["modelName"].(string)
        entity.NextStep = entityData["NextStep"].(int32)
        entity.NoPhysics = entityData["noPhysics"].(bool)
        entity.OBob = entityData["oBob"].(float32)
        entity.ORun = entityData["oRun"].(float32)
        entity.OTilt = entityData["oTilt"].(float32)
        entity.OnGround = entityData["onGround"].(bool)
        entity.Pushthrough = entityData["pushthrough"].(float32)
        entity.Removed = entityData["removed"].(bool)
        entity.RenderOffset = entityData["renderOffset"].(float32)
        entity.Rot = entityData["rot"].(float32)
        entity.RotA = entityData["rotA"].(float32)
        entity.RotOffs = entityData["rotOffs"].(float32)
        entity.Run = entityData["run"].(float32)
        entity.score = entityData["score"].(int32)
        entity.Slide = entityData["slide"].(bool)
        entity.Speed = entityData["speed"].(float32)
        entity.TextureId = entityData["textureId"].(int32)
        entity.TextureName = entityData["textureName"].(string)
        entity.TickCount = entityData["tickCount"].(int32)
        entity.Tilt = entityData["tilt"].(float32)
        entity.TimeOffs = entityData["timeOffs"].(float32)
        entity.UserType = entityData["userType"].(int8)
        entity.WalkDist = entityData["walkDist"].(float32)
        entity.WalkDistO = entityData["walkDistO"].(float32)
        entity.x = entityData["x"].(float32)
        entity.XOld = entityData["xOld"].(float32)
        entity.xRot = entityData["xRot"].(float32)
        entity.XRotO = entityData["xRotO"].(float32)
        entity.Xd = entityData["xd"].(float32)
        entity.Xo = entityData["xo"].(float32)
        entity.y = entityData["y"].(float32)
        entity.YBodyRot = entityData["yBodyRot"].(float32)
        entity.YBodyRotO = entityData["yBodyRotO"].(float32)
        entity.YOld = entityData["yOld"].(float32)
        entity.yRot = entityData["yRot"].(float32)
        entity.YRotO = entityData["yRotO"].(float32)
        entity.YSlideOffset = entityData["ySlideOffset"].(float32)
        entity.Yd = entityData["yd"].(float32)
        entity.Yo = entityData["yo"].(float32)
        entity.z = entityData["z"].(float32)
        entity.ZOld = entityData["zOld"].(float32)
        entity.Zd = entityData["zd"].(float32)
        entity.Zo = entityData["zo"].(float32)

        entities = append(entities, entity)
    }
    return entities
}

// Returns a map[string]any of the entities in the world or nil if not present
// func (cp *ClassicParser) parseEntities(data map[string]any) map[string]any {
//     entities, ok := data["entities"].(map[string]any)
//     if !ok {
//         entities, ok = data["blockMap"].(map[string]any)
//         if !ok {
//             // TODO: Return empty value instead to stop crashing
//             return nil
//         }
//     }
//     return entities
// }