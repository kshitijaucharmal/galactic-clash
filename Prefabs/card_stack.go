components {
  id: "click_game_object"
  component: "/Scripts/click_game_object.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/Atlases/BoardgameIcons.atlas\"\n"
  "default_animation: \"cards_stack_high\"\n"
  "material: \"/Shaders/outline.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "card_factory"
  type: "factory"
  data: "prototype: \"/Prefabs/card.go\"\n"
  "load_dynamically: false\n"
  "dynamic_prototype: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "shadow"
  type: "sprite"
  data: "tile_set: \"/Atlases/BoardgameIcons.atlas\"\n"
  "default_animation: \"cards_stack_high\"\n"
  "material: \"/Shaders/shadow.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -1.0
    y: -12.0
    z: -0.99
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "icon"
  type: "sprite"
  data: "tile_set: \"/Atlases/BoardgameIcons.atlas\"\n"
  "default_animation: \"card_add\"\n"
  "material: \"/Shaders/outline.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 40.0
    y: 68.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.5
    y: 0.5
    z: 1.0
  }
}
