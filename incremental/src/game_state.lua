local skill_type = require("src.skill_type")
local game_utils = require("src.utils.table_utils")
local balance = require("src.game_balance")
local game_speed_type = require("src.game_speed_type")
local game_difficulty_type = require("src.difficulty_type")

local M = {
    game_settings = {},
    current_level = "level_1",
    player_stats = {},
    skill_levels = {},
    skill_costs = {},
    on_pause = false,
    game_speed = game_speed_type.normal,
    game_difficulty = game_difficulty_type.normal
}

function M.init_state()
    M.player_stats, M.skill_costs, M.skill_levels = {}, {}, {}
    M.player_stats = game_utils.copy_table(balance.base_player_stats)
    M.skill_costs = game_utils.copy_table(balance.base_skill_costs)
    M.skill_levels = game_utils.copy_table(balance.base_skill_levels)
    M.game_settings = game_utils.copy_table(balance.default_game_settings)
    M.on_pause = false
    M.game_speed = game_speed_type.normal
end

return M
