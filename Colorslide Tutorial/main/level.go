components {
  id: "level"
  component: "/src/game_objects/level.script"
}
components {
  id: "gui"
  component: "/main/gui/level.gui"
}
embedded_components {
  id: "brickfactory"
  type: "collectionfactory"
  data: "prototype: \"/main/brick/brick.collection\"\n"
  ""
}
