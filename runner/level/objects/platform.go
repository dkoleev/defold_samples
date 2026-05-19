components {
  id: "platform"
  component: "/src/platform.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"geometry\"\n"
  "mask: \"hero\"\n"
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
  "  data: 185.31258\n"
  "  data: 68.996796\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "spikes_left"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 391.0\n"
  "  y: 153.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -210.0
    z: -0.1
  }
}
embedded_components {
  id: "spikes_right"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 194.0
    z: -0.1
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "spikes_down_left"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -87.0
    y: -88.0
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "spikes_down_right"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 71.0
    y: -88.0
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "danger_edges"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"danger\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -207.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 200.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -8.0\n"
  "      y: -89.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  data: 19.796175\n"
  "  data: 58.78456\n"
  "  data: 10.0\n"
  "  data: 18.415113\n"
  "  data: 58.438198\n"
  "  data: 10.0\n"
  "  data: 137.35507\n"
  "  data: 23.62492\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/level/objects/coin.go\"\n"
  ""
}
