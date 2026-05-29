local state = require("src.game_state")

local M = {}

function M.play_music()
    local theme = math.random(2) == 1 and "sound:/music#main_theme_1" or "sound:/music#main_theme_2"
    sound.play(theme, {
        gain = 0.4, pan = 0, speed = 1
    })
end

function M.set_music(active)
    sound.set_group_gain("music", active and 1 or 0 )
end

function M.set_sfx(active)
    sound.set_group_gain("sfx", active and 1 or 0 )
end

function M.click_sfx()
    sound.play("sound:/sfx#click")
end

function M.explosion_sfx()
    sound.play("sound:/sfx#explosion_big")
end

function M.explosion_big_sfx()
    sound.play("sound:/sfx#explosion_small")
end

return M