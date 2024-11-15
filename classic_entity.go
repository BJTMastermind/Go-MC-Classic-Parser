package mc_classic_parser

import "fmt"

type ClassicEntity struct {
    AirSupply int32
    AllowAlpha bool
    AnimStep float32
    AnimStepO float32
    AttackTime int32
    Bb map[string]float32 // mapping same as original AABB class. map keys range from x0 to z1
    BbHeight float32
    BbWidth float32
    BobStrength float32
    Collision bool
    Dead bool
    DeathScore int32
    DeathTime int32
    FallDistance float32
    FootSize float32
    HasHair bool
    Health int32
    HeightOffset float32
    HorizontalCollision bool
    Hovered bool
    HurtDir float32
    HurtDuration int32
    HurtTime int32
    InvulnerableDuration int32
    InvulnerableTime int32
    LastHealth int32
    MakeStepSound bool
    ModelName string
    NoPhysics bool
    ORun float32
    OTilt float32
    OnGround bool
    Pushthrough float32
    Removed bool
    RenderOffset float32
    Rot float32
    RotA float32
    RotOffs float32
    Run float32
    Slide bool
    Speed float32
    TextureId int32
    TextureName string
    TickCount int32
    Tilt float32
    TimeOffs float32
    WalkDist float32
    WalkDistO float32
    X float32
    XOld float32
    XRot float32
    XRotO float32
    Xd float32
    Xo float32
    Y float32
    YBodyRot float32
    YBodyRotO float32
    YOld float32
    YRot float32
    YRotO float32
    YSlideOffset float32
    Yd float32
    Yo float32
    Z float32
    ZOld float32
    Zd float32
    Zo float32
}

func (player ClassicEntity) ToString() string {
    format := fmt.Sprintf("%+v\n", player)
    return format
}
