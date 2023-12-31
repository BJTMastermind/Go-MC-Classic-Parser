package mcclassicparser

type ClassicEntity struct {
    Ai map[string]any
    airSupply int32
    AllowAlpha bool
    AnimStep float32
    AnimStepO float32
    arrows int32
    attackTime int32
    Bb map[string]any
    BbHeight float32
    BbWidth float32
    // blockMap map[string]any
    Bob float32
    BobStrength float32
    Collision bool
    Dead bool
    DeathScore int32
    deathTime int32
    fallDistance float32
    FootSize float32
    HasHair bool
    health int32
    HeightOffset float32
    HorizontalCollision bool
    Hovered bool
    HurtDir float32
    HurtDuration int32
    hurtTime int32
    inventory map[string]any
    InvulnerableDuration int32
    InvulnerableTime int32
    LastHealth int32
    // Level ClassicWorld
    MakeStepSound bool
    Map map[string]any
    ModelName string
    NextStep int32
    NoPhysics bool
    OBob float32
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
    score int32
    Slide bool
    Speed float32
    TextureId int32
    TextureName string
    TickCount int32
    Tilt float32
    TimeOffs float32
    UserType int8
    WalkDist float32
    WalkDistO float32
    x float32
    XOld float32
    xRot float32
    XRotO float32
    Xd float32
    Xo float32
    y float32
    YBodyRot float32
    YBodyRotO float32
    YOld float32
    yRot float32
    YRotO float32
    YSlideOffset float32
    Yd float32
    Yo float32
    z float32
    ZOld float32
    Zd float32
    Zo float32
}
