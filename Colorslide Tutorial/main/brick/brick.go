components {
  id: "script"
  component: "/src/game_objects/brick.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"blue\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/bricks.atlas\"\n"
  "}\n"
  ""
}
