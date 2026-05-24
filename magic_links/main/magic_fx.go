components {
  id: "magic_fx"
  component: "/main/magic_fx.script"
}
embedded_components {
  id: "light"
  type: "sprite"
  data: "default_animation: \"magic-sphere_layer3\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "magic"
  type: "sprite"
  data: "default_animation: \"magic-sphere_layer2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.01
  }
}
