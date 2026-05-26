components {
  id: "enemy_ship"
  component: "/main/enemy_ship.script"
  properties {
    id: "ship_type"
    value: "yellow_ship"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"island\"\n"
  "mask: \"bullet\"\n"
  "mask: \"fire_range\"\n"
  "mask: \"bomb\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 100.0\n"
  "  data: 40.0\n"
  "  data: 0.5\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"yellow_ship\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 419.0\n"
  "  y: 229.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.6
  }
  scale {
    x: 0.3
    y: 0.3
  }
}
