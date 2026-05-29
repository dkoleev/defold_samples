local state = require("src.game_state")

local M = {}

function M.on_input(action_id, action)
    if action_id == hash("change_game_speed") and action.pressed then
        local factor = state.on_pause and 1 or 0
        state.on_pause = not state.on_pause
        msg.post("/bootstrap#level_1_proxy", "set_time_step", { factor = factor, mode = 0 })
    end
end

return M
