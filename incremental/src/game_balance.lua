local skill_type = require("src.skill_type")
local difficulty_type = require("src.difficulty_type")

local M = {
    default_game_settings = {
        music_on = false,
        sound_fx_on = false,
        difficulty = 2
    },
    base_player_stats = {
        bomb_unlocked = false,
        bullet_damage = 10,
        bullet_speed = 200,
        bullet_interval = 2,
        bomb_interval = 3,
        bomb_damage = 20,
        bomb_speed = 100,
        max_hp = 100,
        hp = 10,
        hp_regen = 2,
        hp_regen_interval = 2,
        gold_amount = 100,
        fire_range = 900
    },
    base_skill_levels = {
        [skill_type.bullet_damage] = 1,
        [skill_type.bullet_interval] = 1,
        [skill_type.bomb_damage] = 1,
        [skill_type.bomb_interval] = 1,
        [skill_type.max_hp] = 1,
        [skill_type.hp_regen] = 1,
        [skill_type.fire_range] = 1
    },
    base_skill_costs = {
        [skill_type.bullet_damage] = 5,
        [skill_type.bullet_interval] = 7,
        [skill_type.bomb_damage] = 6,
        [skill_type.bomb_interval] = 9,
        [skill_type.max_hp] = 7,
        [skill_type.hp_regen] = 6,
        [skill_type.fire_range] = 5,
        unlock_bomb = 20
    },
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
