local skill_type = require("src.skill_type")
local game_utils = require("src.utils.table_utils")
local balance = require("src.game_balance")

local M = {
    game_settings = {},
    current_level = "level_1",
    player_stats = {},
    skill_levels = {},
    skill_costs = {},
    on_pause = false
}

function M.init_state()
    M.player_stats, M.skill_costs, M.skill_levels = {}, {}, {}
    M.player_stats = game_utils.copy_table(balance.base_player_stats)
    M.skill_costs = game_utils.copy_table(balance.base_skill_costs)
    M.skill_levels = game_utils.copy_table(balance.base_skill_levels)
    M.game_settings = game_utils.copy_table(balance.default_game_settings)

end

return M
