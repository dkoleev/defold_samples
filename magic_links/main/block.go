components {
  id: "block"
  component: "/main/block.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"orange\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "cover"
  type: "sprite"
  data: "default_animation: \"magic-sphere_layer4\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.03
  }
}
embedded_components {
  id: "factory_fx"
  type: "factory"
  data: "prototype: \"/main/magic_fx.go\"\n"
  ""
}
