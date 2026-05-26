local log = require("log.log")
local logger = log.get_logger("enemy_ship")

local M = {}

function M.print_table(t)
    for key, value in pairs(t) do
        logger:debug(tostring(key) .. " : " .. tostring(value))
        -- print(key, " : ", value)
    end
end

function M.remove_value_from_table(t, value)
    for key, cur_value in pairs(t) do
        if t[key] ~= nil and cur_value == value then
            table.remove(t, key)
            return true
        end
    end

    return false
end

return M