components {
  id: "platform"
  component: "/src/platform.script"
}
embedded_components {
  id: "sprite_main"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -175.0
  }
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
  "      x: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 362.12244\n"
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
    x: -379.0
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
    x: 372.0
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
    x: -250.0
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
    x: -82.0
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
  "      x: -376.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 363.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 2.0\n"
  "      y: -86.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  data: 19.796175\n"
  "  data: 58.78456\n"
  "  data: 10.0\n"
  "  data: 14.729368\n"
  "  data: 49.206894\n"
  "  data: 10.0\n"
  "  data: 310.9739\n"
  "  data: 23.62492\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite_main_2"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
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
    x: 175.0
  }
}
embedded_components {
  id: "spikes_down_left1"
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
    x: 82.0
    y: -88.0
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "spikes_down_left2"
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
    x: 253.0
    y: -88.0
    z: -0.1
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/level/objects/coin.go\"\n"
  ""
}
