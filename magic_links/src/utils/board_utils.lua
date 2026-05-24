local M = {}

--
-- filter (function, table) -> table
-- Returns a new table containing only the elements of the input table for which the function returns true
-- e.g. filter(function(x) return x > 5 end, {1, 6, 3, 8}) -> {6, 8}
--
function M.filter(func, tbl)
    local result = {}

    for i, v in ipairs(tbl) do
        if func(v) then
            table.insert(result, v)
        end
    end

    return result
end

function M.in_blocklist(blocks, block)
    for i, b in pairs(blocks) do
        if b.id == block then
            return true
        end
    end
    return false
end

return M
