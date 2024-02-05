components {
  id: "card"
  component: "/Scripts/card.script"
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
  "default_animation: \"card\"\n"
  "material: \"/Shaders/outline.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "repeat_text"
  type: "label"
  data: "size {\n"
  "  x: 50.0\n"
  "  y: 12.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.902\n"
  "  y: 0.902\n"
  "  z: 0.902\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"X\"\n"
  "font: \"/Fonts/moonchild/moonchild.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 10.0
    y: 34.0
    z: 1.0
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
embedded_components {
  id: "repeat_count"
  type: "label"
  data: "size {\n"
  "  x: 50.0\n"
  "  y: 12.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.902\n"
  "  y: 0.902\n"
  "  z: 0.902\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"0\\n"
  "\"\n"
  "  \"\"\n"
  "font: \"/Fonts/moonchild/moonchild.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 25.0
    y: 34.0
    z: 1.0
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
embedded_components {
  id: "card_type"
  type: "label"
  data: "size {\n"
  "  x: 50.0\n"
  "  y: 12.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.902\n"
  "  y: 0.902\n"
  "  z: 0.902\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"0\\n"
  "\"\n"
  "  \"\"\n"
  "font: \"/Fonts/crunch_chips/crunch_chips.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 0.0
    y: -2.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.45
    y: 0.45
    z: 1.0
  }
}
embedded_components {
  id: "shadow"
  type: "sprite"
  data: "tile_set: \"/Atlases/BoardgameIcons.atlas\"\n"
  "default_animation: \"card\"\n"
  "material: \"/Shaders/shadow.material\"\n"
  "blend_mode: BLEND_MODE_MULT\n"
  ""
  position {
    x: 8.0
    y: -8.0
    z: 0.99
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
