local skill_type = require("src.skill_type")

local M = {
    skill_effects = {
        [skill_type.bullet_damage] = 1.2,
        [skill_type.bullet_interval] = 0.9,
        [skill_type.hp_regen_interval] = 0.9,
        [skill_type.bomb_damage] = 1.3,
        [skill_type.bomb_interval] = 0.85,
        [skill_type.max_hp] = 1.1,
        [skill_type.hp_regen] = 1.3,
        [skill_type.fire_range] = 1.1
    }
}

return M