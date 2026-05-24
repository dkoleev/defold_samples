-- main/colors.lua
local M = {}

M.type = {
    ORANGE = hash("orange"),
    PINK   = hash("pink"),
    BLUE   = hash("blue"),
    YELLOW = hash("yellow"),
    GREEN  = hash("green"),
    MAGIC = hash("magic")
}

-- A separate array of just the hashes for counting and random picking
M.list = {
    M.type.ORANGE,
    M.type.PINK,
    M.type.BLUE,
    M.type.YELLOW,
    M.type.GREEN
}

function M.get_random_color()
    return M.list[math.random(#M.list)]
end

return M
