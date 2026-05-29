local state = require("src.game_state")
local game_speed_type = require("src.game_speed_type")
local sound_manager = require("src.sound_manager")

local M = {}

function M.init_game()
    state.init_state()
    msg.post("/bootstrap#level_1_proxy", "set_time_step", { factor = 1, mode = 0 })
end

function M.restart_game()
    state.init_state()
    msg.post("/bootstrap#level_1_proxy", "set_time_step", { factor = 1, mode = 0 })
end

function M.on_input(action_id, action)
    if action_id == hash("pause_game") and action.pressed then
        local factor = state.on_pause and 1 or 0
        state.on_pause = not state.on_pause
        msg.post("/bootstrap#level_1_proxy", "set_time_step", { factor = factor, mode = 0 })
    elseif action_id == hash("change_game_speed") and action.pressed then
        if state.on_pause then
            return
        end
        
        local target_speed = game_speed_type.normal
        if state.game_speed == game_speed_type.slow then
            target_speed = game_speed_type.normal
        elseif state.game_speed == game_speed_type.normal then
            target_speed = game_speed_type.fast
        else
            target_speed = game_speed_type.slow
        end

        state.game_speed = target_speed

        msg.post("/bootstrap#level_1_proxy", "set_time_step", { factor = state.game_speed, mode = 0 })        
    elseif action_id == hash("switch_music") and action.pressed then
        state.game_settings.music_on = not state.game_settings.music_on
        sound_manager.set_music(state.game_settings.music_on)
    elseif action_id == hash("switch_sfx") and action.pressed then
        state.game_settings.sound_fx_on = not state.game_settings.sound_fx_on
        sound_manager.set_sfx(state.game_settings.sound_fx_on)
    end
end

return M
